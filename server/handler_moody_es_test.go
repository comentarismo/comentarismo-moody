package server_test

import (
	"comentarismo-moody/server"
	"github.com/drewolson/testflight"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestMoodyEsGet(t *testing.T) {

	server.InitProviders()
	testflight.WithServer(server.InitRouting(), func(r *testflight.Requester) {
		response := r.Get("/moody?vid=https://www.youtube.com/watch?v=n8F55puGHIs&lang=es&refresh=true")
		//response := r.Get("/moody?vid=https://www.youtube.com/watch?v=HKU96i_Qh8Y")

		log.Println(len(response.Body))
		assert.Equal(t, 200, response.StatusCode)
		assert.True(t, len(response.Body) > 0)
	})
}
