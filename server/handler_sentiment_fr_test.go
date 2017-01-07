package server_test

import (
	"comentarismo-moody/server"
	"github.com/drewolson/testflight"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

//$ SHIELD_DEBUG=true go test  handler_sentiment_en_test.go -v
func TestFrSentimentPost(t *testing.T) {

	server.InitProviders()
	testflight.WithServer(server.InitRouting(), func(r *testflight.Requester) {
		response := r.Post("/sentiment?lang=fr", testflight.FORM_ENCODED, "text=Que voulez-vous faire ? On ne peut pas lutter contre les salaires chinois ! C'est trop facile d'accuser le gouvernement !")

		log.Println(response.Body)
		assert.Equal(t, 200, response.StatusCode)
		assert.True(t, len(response.Body) > 0)

	})
}
