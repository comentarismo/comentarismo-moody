package sentiment

import (
	"regexp"
	"strings"
)

type itTokenizer struct {}

func NewItTokenizer() Tokenizer {
	return &itTokenizer{}
}

func (t *itTokenizer) Tokenize(text string) (words map[string]int64) {
	words = make(map[string]int64)
	for _, w := range splitTokenRxIt.Split(text, -1) {
		if len(w) > 2 {
			words[strings.ToLower(w)]++
		}
	}
	return
}

var itStopWords = `|a|adesso|ai|al|alla|allo|allora|altre|altri|altro|anche|ancora|avere|aveva|avevano|ben|buono|che|chi|cinque|comprare|con|consecutivi|consecutivo|cosa|cui|da|del|della|dello|dentro|deve|devo|di|doppio|due|e|ecco|fare|fine|fino|fra|gente|giu|ha|hai|hanno|ho|il|indietro|invece|io|la|lavoro|le|lei|lo|loro|lui|lungo|ma|me|meglio|molta|molti|molto|nei|nella|no|noi|nome|nostro|nove|nuovi|nuovo|o|oltre|ora|otto|peggio|pero|persone|piu|poco|primo|promesso|qua|quarto|quasi|quattro|quello|questo|qui|quindi|quinto|rispetto|sara|secondo|sei|sembra|sembrava|senza|sette|sia|siamo|siete|solo|sono|sopra|soprattutto|sotto|stati|stato|stesso|su|subito|sul|sulla|tanto|te|tempo|terzo|tra|tre|triplo|ultimo|un|una|uno|va|vai|voi|volte|vostro`
var splitTokenRxIt = regexp.MustCompile(`[^\w]+` + itStopWords)
