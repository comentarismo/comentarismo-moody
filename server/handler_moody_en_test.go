package server_test

import (
	"comentarismo-moody/model"
	"comentarismo-moody/server"
	"github.com/drewolson/testflight"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestMoodyEnGet(t *testing.T) {

	store := model.InitSentiment("en")
	err := store.Reset()
	assert.Equal(t, err, nil)

	server.InitProviders()
	testflight.WithServer(server.InitRouting(), func(r *testflight.Requester) {

		response := r.Get("/moody?vid=https://www.youtube.com/watch?v=HKU96i_Qh8Y&lang=en&refresh=true")
		//response := r.Get("/moody?vid=https://www.youtube.com/watch?v=HKU96i_Qh8Y")

		log.Println(len(response.Body))
		assert.Equal(t, 200, response.StatusCode)
		assert.True(t, len(response.Body) > 0)
	})
}
