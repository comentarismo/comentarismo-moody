package vinevideo_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"comentarismo-moody/providers/facebook"
	"comentarismo-moody/model"
)

func Test_New(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	provider := facebookProvider()
	a.Equal(provider.ClientKey, os.Getenv("FACEBOOK_KEY"))
	a.Equal(provider.Secret, os.Getenv("FACEBOOK_SECRET"))
}

func Test_Implements_Provider(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	a.Implements((*model.Provider)(nil), facebookProvider())
}


func facebookProvider() *facebook.Provider {
	return facebook.New(os.Getenv("FACEBOOK_KEY"), os.Getenv("FACEBOOK_SECRET"), "/foo", "email")
}
