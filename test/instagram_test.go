package test

import (
	"testing"

	"comentarismo-moody/model"
	"comentarismo-moody/providers/facebook"
	"github.com/stretchr/testify/assert"
)

func Test_Instagram_New(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	provider := instagramProvider()
	a.Equal(provider.ClientKey, KEY)
	a.Equal(provider.Secret, SECRET)
}

func Test_Instagram_Implements_Provider(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	a.Implements((*model.Provider)(nil), instagramProvider())
}

func instagramProvider() *facebook.Provider {
	if KEY == "" {
		KEY = "asd"
	}
	if SECRET == "" {
		SECRET = "adsd"
	}
	return facebook.New(KEY, SECRET, "/foo", "email")
}
