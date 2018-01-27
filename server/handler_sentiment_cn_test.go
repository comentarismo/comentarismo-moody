package server_test

import (
	"comentarismo-moody/server"
	"github.com/drewolson/testflight"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

// SENTIMENT_DEBUG=true go test  handler_sentiment_cn_test.go -v
func TestCNSentimentPost(t *testing.T) {

	server.InitProviders()
	testflight.WithServer(server.InitRouting(), func(r *testflight.Requester) {
		response := r.Post("/sentiment?lang=cn", testflight.FORM_ENCODED, "text=南京现在也有很多人加入挖比特币行业，刚刚兴起的，都靠偷国家的电赚钱")

		log.Println("INFO: TEST RESULT: ", response.Body)
		assert.Equal(t, 200, response.StatusCode)
		assert.True(t, len(response.Body) > 0)

	})
}
