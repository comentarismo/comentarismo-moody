package server_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/drewolson/testflight"
	"testing"
	"log"
	"comentarismo-moody/server"
)

//$ SHIELD_DEBUG=true go test  handler_sentiment_en_test.go -v
func TestItSentimentPost(t *testing.T) {

	server.InitProviders()
	testflight.WithServer(server.InitRouting(), func(r *testflight.Requester) {
		response := r.Post("/sentiment?lang=it", testflight.FORM_ENCODED, "text=ogni giorno una tremenda caz..ta, tra un pò, sicuramente  prima delle elzioni amministrative, direte che avete visto e parlato con padre pio, il quale vi dirà che come al solito siete i più onesti, i migliori ecc. ecc.");

		log.Println(response.Body)
		assert.Equal(t, 200, response.StatusCode)
		assert.True(t,len(response.Body) > 0)

	})
}