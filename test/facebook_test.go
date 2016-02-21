package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"comentarismo-moody/providers/facebook"
	"comentarismo-moody/model"
)

func Test_Facebook_New(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	provider := facebookProvider()
	a.Equal(provider.ClientKey, KEY)
	a.Equal(provider.Secret, SECRET)
}

func Test_Facebook_Implements_Provider(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	a.Implements((*model.Provider)(nil), facebookProvider())
}


func facebookProvider() *facebook.Provider {
	if KEY == "" {
		KEY = "asd"
	}
	if SECRET == "" {
		SECRET = "adsd"
	}
	return facebook.New(KEY, SECRET, "/foo", "email")
}
