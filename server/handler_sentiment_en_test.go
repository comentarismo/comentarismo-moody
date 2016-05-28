package server_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/drewolson/testflight"
	"testing"
	"log"
	"comentarismo-moody/server"
)

//$ SHIELD_DEBUG=true go test  handler_sentiment_en_test.go -v
func TestEnSentimentPost(t *testing.T) {

	server.InitProviders()
	testflight.WithServer(server.InitRouting(), func(r *testflight.Requester) {
		response := r.Post("/sentiment?lang=en", testflight.FORM_ENCODED, "text=Yep it's good. But I had trouble with the interval sound not working and the sound varying. So I switched to Zenso which is basic but works as a timer with intervals which is all I want.");

		log.Println(response.Body)
		assert.Equal(t, 200, response.StatusCode)
		assert.True(t,len(response.Body) > 0)

	})
}