package test

import (
	"testing"

	"comentarismo-moody/model"
	"comentarismo-moody/providers/youtube"
	"github.com/stretchr/testify/assert"
)

func Test_Youtube_New(t *testing.T) {
	t.Parallel()
	a := assert.New(t)
	provider := youtubeProvider()
	a.Equal(provider.ClientKey, KEY)
	a.Equal(provider.Secret, SECRET)
}

func Test_Youtube_Implements_Provider(t *testing.T) {
	t.Parallel()
	a := assert.New(t)

	a.Implements((*model.Provider)(nil), youtubeProvider())
}

func youtubeProvider() *youtube.Provider {
	if KEY == "" {
		KEY = "asd"
	}
	return youtube.New(KEY, SECRET, "/foo", "email")
}
