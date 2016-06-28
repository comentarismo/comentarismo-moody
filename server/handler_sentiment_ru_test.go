package server_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/drewolson/testflight"
	"testing"
	"log"
	"comentarismo-moody/server"
)

//$ SHIELD_DEBUG=true go test  handler_sentiment_en_test.go -v
func TestRuSentimentPost(t *testing.T) {

	server.InitProviders()
	testflight.WithServer(server.InitRouting(), func(r *testflight.Requester) {
		response := r.Post("/sentiment?lang=ru", testflight.FORM_ENCODED, "text=Уже тошнит от таких новостей. Лучше бы деньги потраченные на заработную плату этих борцов с экстремизмом израсходовали на лечение больных детей, или сиротам еды бы приличной купили.");

		log.Println(response.Body)
		assert.Equal(t, 200, response.StatusCode)
		assert.True(t,len(response.Body) > 0)

	})
}