package sentiment

import (
	"github.com/garyburd/redigo/redis"
	"log"
	"strconv"
)

func NewRedisStore(addr, password string, logger *log.Logger, prefix string) Store {
	return &RedisModel{
		addr:       addr,
		password:   password,
		sumKey:     "sentiment:sum",
		classKey:   "sentiment:class",
		classesKey: "sentiment:classes",
		logger:     logger,
		prefix:     prefix,
	}
}

func (rs *RedisModel) conn() (conn redis.Conn, err error) {
	if rs.redis == nil {
		c, err2 := redis.Dial("tcp", rs.addr)
		if err2 != nil {
			Debug("Error: conn() redis.Dial -> ", rs.addr, err2)
			return nil, err2
		}
		if rs.logger != nil {
			c = redis.NewLoggingConn(c, rs.logger, rs.prefix)
		}
		if rs.password != "" {
			_, authErr := redis.String(c.Do("AUTH", rs.password))
			if authErr != nil {
				Debug("Error: conn() AUTH -> ",authErr)
				err = authErr
				return
			}
		}
		rs.redis = c
	}
	return rs.redis, nil
}

func (rs *RedisModel) Classes() (a []string, err error) {
	c, err := rs.conn()
	if err != nil {
		Debug("Error: Classes rs.conn() -> ",err)
		return
	}
	classes, err := redis.Strings(c.Do("SMEMBERS", rs.classesKey))
	if err != nil {
		Debug("Error: Classes redis.Strings(c.Do(SMEMBERS)) -> ",err)
		return
	}
	for _, v := range classes {
		a = append(a, string(v))
	}
	return
}

func (rs *RedisModel) AddClass(class string) (err error) {
	c, err := rs.conn()
	if err != nil {
		Debug("Error: AddClass rs.conn() -> ",err)
		return
	}
	if class == "" {
		panic("invalid class: " + class)
	}
	_, err = c.Do("SADD", rs.classesKey, class)
	if err != nil {
		Debug("Error: AddClass c.Do(SADD) -> ",err)
		return
	}
	return err
}

func (rs *RedisModel) ClassWordCounts(class string, words []string) (mc map[string]int64, err error) {
	c, err := rs.conn()
	if err != nil {
		Debug("Error: ClassWordCounts rs.conn() -> ",err)
		return
	}
	//begin transaction
	c.Send("MULTI")

	key := rs.classKey + ":" + class
	Debug("ClassWordCounts key ->",key)
	args := make([]interface{}, 0, len(words)+1)
	args = append(args, key)
	for _, v := range words {
		args = append(args, v)
	}
	Debug("ClassWordCounts args ->",args)

	c.Send("HMGET", args...)
	values, err := redis.Values(c.Do("EXEC"))
	if err != nil {
		Debug("Error: ClassWordCounts redis.Values(c.Do(EXEC)) -> ",err)
		return
	}

	Debug("ClassWordCounts values ->", values)

	if len(values) > 0 {
		if x, ok := values[0].([]interface{}); ok {
			values = x
		}
	}

	var i int64

	mc = make(map[string]int64)
	for len(values) > 0 {
		var count int64
		//// Scan copies from values to the values pointed at by count.
		values, err = redis.Scan(values, &count)
		if err != nil {
			Debug("Error: ClassWordCounts redis.Scan -> ",err)
			return
		}
		mc[words[i]] = count
		i++
	}

	Debug("ClassWordCounts mc -> ",mc)

	return
}

func (rs *RedisModel) IncrementClassWordCounts(m map[string]map[string]int64) (err error) {
	c, err := rs.conn()
	if err != nil {
		Debug("Error: IncrementClassWordCounts() rs.conn() -> ",err)
		return
	}
	type tuple struct {
		word string
		d    int64
	}
	decrTuples := make(map[string][]*tuple, len(m))

	// Apply only positive increments
	c.Send("MULTI")
	for class, words := range m {
		for word, d := range words {
			if d > 0 {
				c.Send("HINCRBY", rs.classKey+":"+class, word, d)
				c.Send("HINCRBY", rs.sumKey, class, d)
			} else {
				decrTuples[class] = append(decrTuples[class], &tuple{
					word: word,
					d:    d,
				})
			}
		}
	}
	_, err = redis.Values(c.Do("EXEC"))
	if err != nil {
		return
	}

	// If we decrement something, we have to check if we are
	// about to drop below 0 and adjust the value accordingly.
	//
	// 2 trips per value, TODO: optimize
	//
	for class, paths := range decrTuples {
		key := rs.classKey + ":" + class

		// Build HMGET params
		hmget := make([]interface{}, 0, len(paths))
		hmget = append(hmget, key)
		for _, path := range paths {
			Debug("mget -> ",path.word)
			hmget = append(hmget, path.word)
		}
		Debug("IncrementClassWordCounts mget all -> ",hmget)

		values, err2 := redis.Strings(c.Do("HMGET", hmget...))
		if err2 != nil {
			Debug("Error: IncrementClassWordCounts redis.Strings(c.Do(HMGET)) -> ",err2)
			return
		}

		c.Send("MULTI")
		for i, v := range values {
			path := paths[i]
			x, err2 := strconv.ParseInt(v, 10, 64)
			if err2 != nil {
				Debug("Error: IncrementClassWordCounts() strconv.ParseInt -> ",err2)
				panic(err2)
			}
			if x != 0 {
				d := path.d
				if (x + d) < 0 {
					d = x * -1
				}
				c.Send("HINCRBY", key, path.word, d)
				c.Send("HINCRBY", rs.sumKey, class, d)
			}
		}
		_, err2 = c.Do("EXEC")
		if err2 != nil {
			Debug("Error: IncrementClassWordCounts() c.Do(EXEC) -> ",err2)
			return err2
		}
	}
	return
}


func (rs *RedisModel) TotalClassWordCounts() (m map[string]int64, err error) {
	c, err := rs.conn()
	if err != nil {
		Debug("Error: TotalClassWordCounts() rs.conn() -> ",err)
		return
	}
	values, err := redis.Values(c.Do("HGETALL", rs.sumKey))
	if err != nil {
		Debug("Error: TotalClassWordCounts() redis.Values(c.Do(HGETALL)) -> ",err)
		return
	}
	m = make(map[string]int64)
	for len(values) > 0 {
		var class string
		var count int64
		values, err = redis.Scan(values, &class, &count)
		if err != nil {
			Debug("Error: TotalClassWordCounts() redis.Scan -> ",err)
			return
		}
		m[class] = count
	}
	return
}


func (rs *RedisModel) Reset() (err error) {
	c, err := rs.conn()
	if err != nil {
		return
	}
	a, err := redis.Strings(c.Do("KEYS", "sentiment:*"))
	if err != nil {
		Debug("Error: Reset() redis.Strings(c.Do(KEYS)) -> ",err)
		return
	}
	c.Send("MULTI")
	for _, key := range a {
		c.Send("DEL", key)
	}
	_, err = c.Do("EXEC")
	if err != nil {
		Debug("Error: Reset() c.Do(EXEC) -> ",err)
	}
	return
}
