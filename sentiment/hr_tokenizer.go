package sentiment

import (
	"regexp"
	"strings"
)

type hrTokenizer struct{}

func NewHrTokenizer() Tokenizer {
	return &hrTokenizer{}
}

func (t *hrTokenizer) Tokenize(text string) (words map[string]int64) {
	words = make(map[string]int64)
	//log.Println(splitTokenRxHr)
	for _, w := range splitTokenRxHr.Split(text, -1) {
		if len(w) > 1 {
			words[strings.ToLower(w)]++
		}
	}
	return
}

var hrStopWords = `|a|ako|ali|bi|bih|bila|bili|bilo|bio|bismo|biste|biti|bumo|da|do|duž|ga|hoće|hoćemo|hoćete|hoćeš|hoću|i|iako|ih|ili|iz|ja|je|jedna|jedne|jedno|jer|jesam|jesi|jesmo|jest|jeste|jesu|jim|joj|još|ju|kada|kako|kao|koja|koje|koji|kojima|koju|kroz|li|me|mene|meni|mi|mimo|moj|moja|moje|mu|na|nad|nakon|nam|nama|nas|naš|naša|naše|našeg|ne|nego|neka|neki|nekog|neku|nema|netko|neće|nećemo|nećete|nećeš|neću|nešto|ni|nije|nikoga|nikoje|nikoju|nisam|nisi|nismo|niste|nisu|njega|njegov|njegova|njegovo|njemu|njezin|njezina|njezino|njih|njihov|njihova|njihovo|njim|njima|njoj|nju|no|o|od|odmah|on|ona|oni|ono|ova|pa|pak|po|pod|pored|prije|s|sa|sam|samo|se|sebe|sebi|si|smo|ste|su|sve|svi|svog|svoj|svoja|svoje|svom|ta|tada|taj|tako|te|tebe|tebi|ti|to|toj|tome|tu|tvoj|tvoja|tvoje|u|uz|vam|vama|vas|vaš|vaša|vaše|već|vi|vrlo|za|zar|će|ćemo|ćete|ćeš|ću|što`
var splitTokenRxHr = regexp.MustCompile(`[^\w]+` + hrStopWords)
