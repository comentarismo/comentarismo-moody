package server

import (
	"gopkg.in/redis.v3"
	"log"
	"time"
)

// An object implementing Store interface can be registered with SetCustomStore
// function to handle storage and retrieval of captcha ids and solutions for
// them, replacing the default memory store.
//
// It is the responsibility of an object to delete expired and used captchas
// when necessary (for example, the default memory store collects them in Set
// method after the certain amount of captchas has been stored.)
type Store interface {
	// Set sets the digits for the captcha id.
	Set(id string, digits []byte)

	// Get returns stored digits for the captcha id. Clear indicates
	// whether the captcha must be deleted from the store.
	Get(id string, clear bool) (digits []byte)
}

// memoryStore is an internal store for captcha ids and their values.
type memoryStore struct {
}

// NewMemoryStore returns a new standard memory store for captchas with the
// given collection threshold and expiration time (duration). The returned
// store must be registered with SetCustomStore to replace the default one.
func NewMemoryStore() Store {
	log.Println("captcha NewMemoryStore")
	s := new(memoryStore)
	return s
}

func (s *memoryStore) Set(id string, digits []byte) {
	log.Println("captcha Set")
	Client.Set(id, digits, time.Minute*5)
}

func (s *memoryStore) Get(id string, clear bool) (digits []byte) {
	log.Println("captcha Get")
	d, err := Client.Get(id).Result()
	if err == redis.Nil {
		log.Println("captcha " + id + " does not exists")
	} else if err != nil {
		log.Println(err)
	} else {
		digits = []byte(d)
		log.Println("captcha Serve from cache")
		return
	}
	return
}
