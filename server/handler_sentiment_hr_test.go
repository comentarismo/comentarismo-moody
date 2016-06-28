package server_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/drewolson/testflight"
	"testing"
	"log"
	"comentarismo-moody/server"
)

//$ SHIELD_DEBUG=true go test  handler_sentiment_en_test.go -v
func TestHrSentimentPost(t *testing.T) {

	server.InitProviders()
	testflight.WithServer(server.InitRouting(), func(r *testflight.Requester) {
		response := r.Post("/sentiment?lang=hr", testflight.FORM_ENCODED, "text=Nazalost, ona vise nije pevacica, sada je samo devojka Carlija Corluke. Pametno, vise ce para tako zaraditi. A i nije da je Carli samo fudbaler, kakvi su ovi vasi tipa Vida i Modric, pored njih Carli izgleda odlicno. Franka je predobra, blago Carliju.");

		log.Println(response.Body)
		assert.Equal(t, 200, response.StatusCode)
		assert.True(t,len(response.Body) > 0)

	})
}