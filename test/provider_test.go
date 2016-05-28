package test

import (
	"testing"

	"comentarismo-moody/model"
	"comentarismo-moody/providers/facebook"
	"comentarismo-moody/providers/youtube"
	"comentarismo-moody/providers/faux"
	"github.com/stretchr/testify/assert"
	"comentarismo-moody/providers/instagram"
	"comentarismo-moody/providers/vinevideo"
)

var (
	KEY string
	SECRET string
)


func Test_UseProviders(t *testing.T) {
	a := assert.New(t)

	//create facebook provider
	f := &facebook.Provider{}
	model.UseProviders(f)

	//create facebook provider
	y := &youtube.Provider{}
	model.UseProviders(y)

	//create test  provider
	fb := &faux.Provider{}
	model.UseProviders(fb)

	a.Equal(len(model.GetProviders()), 3)
	a.Equal(model.GetProviders()[f.Name()], f)
	a.Equal(model.GetProviders()[fb.Name()], fb)
	a.Equal(model.GetProviders()[y.Name()], y)
	model.ClearProviders()
}

func Test_GetProvider(t *testing.T) {
	a := assert.New(t)

	//create test  provider
	fa := &faux.Provider{}
	model.UseProviders(fa)

	//create facebook provider
	fb := &facebook.Provider{}
	model.UseProviders(fb)

	//create facebook provider
	y := &youtube.Provider{}
	model.UseProviders(y)

	//create facebook provider
	insta := &instagram.Provider{}
	model.UseProviders(insta)

	//create facebook provider
	vine := &vinevideo.Provider{}
	model.UseProviders(vine)

	p, err := model.GetProvider(fa.Name())
	a.NoError(err)
	a.Equal(p, fa)


	fbp, err := model.GetProvider(fb.Name())
	a.NoError(err)
	a.Equal(fbp, fb)

	yt, err := model.GetProvider(y.Name())
	a.NoError(err)
	a.Equal(yt, y)

	instag, err := model.GetProvider(insta.Name())
	a.NoError(err)
	a.Equal(instag, insta)

	vinevideo, err := model.GetProvider(vine.Name())
	a.NoError(err)
	a.Equal(vine, vinevideo)

	p, err = model.GetProvider("unknown")
	a.Error(err)
	a.Equal(err.Error(), "no provider for unknown exists")
	model.ClearProviders()
}
