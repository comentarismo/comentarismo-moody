package sentiment

import (
	"regexp"
	"strings"
)

type ptTokenizer struct {
}

func NewPtTokenizer() Tokenizer {
	return &ptTokenizer{}
}

func (t *ptTokenizer) Tokenize(text string) (words map[string]int64) {
	words = make(map[string]int64)
	for _, w := range splitTokenRxPt.Split(text, -1) {
		if len(w) > 2 {
			words[strings.ToLower(w)]++
		}
	}
	return
}

// Spamassassin stoplist
//
// http://wiki.apache.org/spamassassin/BayesStopList
//
var ptStopWords = `|último|é|acerca|agora|algmas|alguns|ali|ambos|antes|apontar|aquela|aquelas|aquele|aqueles|aqui|atrás|bem|bom|cada|caminho|cima|com|como|comprido|conhecido|corrente|das|debaixo|dentro|desde|desligado|deve|devem|deverá|direita|diz|dizer|dois|dos|e|ela|ele|eles|em|enquanto|então|está|estão|estado|estar|estará|este|estes|esteve|estive|estivemos|estiveram|eu|fará|faz|fazer|fazia|fez|fim|foi|fora|horas|iniciar|inicio|ir|irá|ista|iste|isto|ligado|maioria|maiorias|mais|mas|mesmo|meu|muito|muitos|nós|não|nome|nosso|novo|o|onde|os|ou|outro|para|parte|pegar|pelo|pessoas|pode|poderá|podia|por|porque|povo|promeiro|quê|qual|qualquer|quando|quem|quieto|são|saber|sem|ser|seu|somente|têm|tal|também|tem|tempo|tenho|tentar|tentaram|tente|tentei|teu|teve|tipo|tive|todos|trabalhar|trabalho|tu|um|uma|umas|uns|usa|usar|valor|veja|ver|verdade|verdadeiro|você`
var splitTokenRxPt = regexp.MustCompile(`[^\w]+` + ptStopWords)
