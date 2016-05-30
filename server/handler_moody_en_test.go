package server_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/drewolson/testflight"
	"testing"
	"log"
	"comentarismo-moody/server"
)

func TestMoodyEnGet(t *testing.T) {

	server.InitProviders()
	testflight.WithServer(server.InitRouting(), func(r *testflight.Requester) {
		response := r.Get("/moody?vid=https://www.youtube.com/watch?v=HKU96i_Qh8Y&lang=en&refresh=true")
		//response := r.Get("/moody?vid=https://www.youtube.com/watch?v=HKU96i_Qh8Y")

		log.Println(len(response.Body))
		assert.Equal(t, 200, response.StatusCode)
		assert.True(t,len(response.Body) > 0)
	})
}
