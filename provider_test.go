package main_test

import (
	"testing"

	model "comentarismo-moody/model"
	"comentarismo-moody/providers/facebook"
	"comentarismo-moody/providers/faux"
	assert "github.com/stretchr/testify/assert"
)

func Test_UseProviders(t *testing.T) {
	a := assert.New(t)

	//create test provider
	f := &facebook.Provider{}
	model.UseProviders(f)

	//create facebook provider
	fb := &faux.Provider{}
	model.UseProviders(fb)

	a.Equal(len(model.GetProviders()), 2)
	a.Equal(model.GetProviders()[f.Name()], f)
	a.Equal(model.GetProviders()[fb.Name()], fb)
	model.ClearProviders()
}

func Test_GetProvider(t *testing.T) {
	a := assert.New(t)

	provider := &faux.Provider{}
	model.UseProviders(provider)

	p, err := model.GetProvider(provider.Name())
	a.NoError(err)
	a.Equal(p, provider)

	p, err = model.GetProvider("unknown")
	a.Error(err)
	a.Equal(err.Error(), "no provider for unknown exists")
	model.ClearProviders()
}
