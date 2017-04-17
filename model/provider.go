package model

import (
	"fmt"
)

// Provider needs to be implemented for each 3rd party provider
// e.g. Comentarismo, Facebook, Twitter, Youtube, etc...
type Provider interface {
	Name() string
	GetType() string
	SetID([]string) error
	SetLang(lang string) error
	SetReport(*Report, CommentList)
	//BeginAuth(state string) (Session, error)
	//UnmarshalSession(string) (Session, error)
	//FetchUser(Session) (User, error)
	GetComments() CommentList
	GetCommentsChan(chan *Comment, chan int)
	GetMetadata() bool
	GetPageID() Provider
	Debug(bool)
}

// Providers is list of known/available providers.
type Providers map[string]Provider

var providers = Providers{}

// UseProviders sets a list of available providers for use by Moody.
func UseProviders(viders ...Provider) {
	for _, provider := range viders {
		providers[provider.Name()] = provider
	}
}

// GetProviders returns a list of all the providers currently in use.
func GetProviders() Providers {
	return providers
}

// GetProvider returns a previously created provider. If Moody has not
// been told to use the named provider it will return an error.
func GetProvider(name string) (Provider, error) {
	provider := providers[name]
	if provider == nil {
		return nil, fmt.Errorf("no provider for %s exists", name)
	}
	//log.Println(provider.Name())
	return provider, nil
}

// ClearProviders will remove all providers currently in use.
// This is useful, mostly, for testing purposes.
func ClearProviders() {
	providers = Providers{}
}
