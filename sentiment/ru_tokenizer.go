package sentiment

import (
	"regexp"
	"strings"
)

type ruTokenizer struct{}

func NewRuTokenizer() Tokenizer {
	return &ruTokenizer{}
}

func (t *ruTokenizer) Tokenize(text string) (words map[string]int64) {
	words = make(map[string]int64)
	for _, w := range splitTokenRxRu.Split(text, -1) {
		if len(w) > 2 {
			words[strings.ToLower(w)]++
		}
	}
	return
}

var ruStopWords = `|а|алло|без|белый|близко|более|больше|большой|будем|будет|будете|будешь|будто|буду|будут|будь|бы|бывает|бывь|был|была|были|было|быть|в|важная|важное|важные|важный|вам|вами|вас|ваш|ваша|ваше|ваши|вверх|вдали|вдруг|ведь|везде|вернуться|весь|вечер|взгляд|взять|вид|видеть|вместе|вниз|внизу|во|вода|война|вокруг|вон|вообще|вопрос|восемнадцатый|восемнадцать|восемь|восьмой|вот|впрочем|времени|время|все|всегда|всего|всем|всеми|всему|всех|всею|всю|всюду|вся|всё|второй|вы|выйти|г|где|главный|глаз|говорил|говорит|говорить|год|года|году|голова|голос|город|да|давать|давно|даже|далекий|далеко|дальше|даром|дать|два|двадцатый|двадцать|две|двенадцатый|двенадцать|дверь|двух|девятнадцатый|девятнадцать|девятый|девять|действительно|дел|делать|дело|день|деньги|десятый|десять|для|до|довольно|долго|должно|должный|дом|дорога|друг|другая|другие|других|друго|другое|другой|думать|душа|е|его|ее|ей|ему|если|есть|еще|ещё|ею|её|ж|ждать|же|жена|женщина|жизнь|жить|за|занят|занята|занято|заняты|затем|зато|зачем|здесь|земля|знать|значит|значить|и|идти|из|или|им|именно|иметь|ими|имя|иногда|их|к|каждая|каждое|каждые|каждый|кажется|казаться|как|какая|какой|кем|книга|когда|кого|ком|комната|кому|конец|конечно|которая|которого|которой|которые|который|которых|кроме|кругом|кто|куда|лежать|лет|ли|лицо|лишь|лучше|любить|люди|маленький|мало|мать|машина|между|меля|менее|меньше|меня|место|миллионов|мимо|минута|мир|мира|мне|много|многочисленная|многочисленное|многочисленные|многочисленный|мной|мною|мог|могут|мож|может|можно|можхо|мои|мой|мор|москва|мочь|моя|моё|мы|на|наверху|над|надо|назад|наиболее|найти|наконец|нам|нами|народ|нас|начала|начать|наш|наша|наше|наши|не|него|недавно|недалеко|нее|ней|некоторый|нельзя|нем|немного|нему|непрерывно|нередко|несколько|нет|нею|неё|ни|нибудь|ниже|низко|никакой|никогда|никто|никуда|ними|них|ничего|ничто|но|новый|нога|ночь|ну|нужно|нужный|нх|о|об|оба|обычно|один|одиннадцатый|одиннадцать|однажды|однако|одного|одной|оказаться|окно|около|он|она|они|оно|опять|особенно|остаться|от|ответить|отец|отовсюду|отсюда|очень|первый|перед|писать|плечо|по|под|подумать|пожалуйста|позже|пойти|пока|пол|получить|помнить|понимать|понять|пор|пора|после|последний|посмотреть|посреди|потом|потому|почему|почти|правда|прекрасно|при|про|просто|против|процентов|пятнадцатый|пятнадцать|пятый|пять|работа|работать|раз|разве|рано|раньше|ребенок|решить|россия|рука|русский|ряд|рядом|с|сам|сама|сами|самим|самими|самих|само|самого|самой|самом|самому|саму|самый|свет|свое|своего|своей|свои|своих|свой|свою|сделать|сеаой|себе|себя|сегодня|седьмой|сейчас|семнадцатый|семнадцать|семь|сидеть|сила|сих|сказал|сказала|сказать|сколько|слишком|слово|случай|смотреть|сначала|снова|со|собой|собою|советский|совсем|спасибо|спросить|сразу|стал|старый|стать|стол|сторона|стоять|страна|суть|считать|та|так|такая|также|таки|такие|такое|такой|там|твой|твоя|твоё|те|тебе|тебя|тем|теми|теперь|тех|то|тобой|тобою|товарищ|тогда|того|тоже|только|том|тому|тот|тою|третий|три|тринадцатый|тринадцать|ту|туда|тут|ты|тысяч|у|увидеть|уж|уже|улица|уметь|утро|хороший|хорошо|хотеть|хоть|хотя|хочешь|час|часто|часть|чаще|чего|человек|чем|чему|через|четвертый|четыре|четырнадцатый|четырнадцать|что|чтоб|чтобы|чуть|шестнадцатый|шестнадцать|шестой|шесть|эта|эти|этим|этими|этих|это|этого|этой|этом|этому|этот|эту|я`
var splitTokenRxRu = regexp.MustCompile(`[^\p{L}]+` + ruStopWords)
