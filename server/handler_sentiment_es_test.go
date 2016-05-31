package server_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/drewolson/testflight"
	"testing"
	"log"
	"comentarismo-moody/server"
)

//$ SHIELD_DEBUG=true go test  handler_sentiment_en_test.go -v
func TestEsSentimentPost(t *testing.T) {

	server.InitProviders()
	testflight.WithServer(server.InitRouting(), func(r *testflight.Requester) {
		response := r.Post("/sentiment?lang=es", testflight.FORM_ENCODED, "text=Buenos periodistas en el paro y estos 'en chu fados' cobrando por este tipo de artículos......así va el país");

		log.Println(response.Body)
		assert.Equal(t, 200, response.StatusCode)
		assert.True(t,len(response.Body) > 0)

	})
}