package server_test

import (
	"comentarismo-moody/model"
	"comentarismo-moody/server"
	"encoding/json"
	"github.com/drewolson/testflight"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

// SENTIMENT_DEBUG=true go test  handler_sentiment_nl_test.go -v
func TestNLSentimentPost(t *testing.T) {

	server.InitProviders()
	testflight.WithServer(server.InitRouting(), func(r *testflight.Requester) {
		response := r.Post("/sentiment?lang=nl", testflight.FORM_ENCODED, "text=Super!!!!!! Super handig, ideaal.........nu kunnen we heerlijk spotify luisteren via onze versterker! Echt heel blij met dit handige kastje!")

		log.Println("TestNLSentimentPost Got Response, ", response.Body)
		assert.Equal(t, 200, response.StatusCode)
		assert.True(t, len(response.Body) > 0)

		sentimentReport := model.Comment{}
		err := json.Unmarshal(response.RawBody, &sentimentReport)

		// expect sentiment to be good
		assert.Equal(t, nil, err)
		assert.Equal(t, 3, sentimentReport.Sentiment)

	})
}
