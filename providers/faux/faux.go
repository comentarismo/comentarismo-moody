// Package faux is used exclusive for testing purposes. I would strongly suggest you move along
// as there's nothing to see here.
package faux

import (
	model "comentarismo-moody/model"
	"encoding/json"
	"strings"
)

// Provider is used only for testing.
type Provider struct {
}

// Session is used only for testing.
type Session struct {
	Name  string
	Email string
}

// Name is used only for testing.
func (p *Provider) Name() string {
	return "faux"
}

// BeginAuth is used only for testing.
func (p *Provider) BeginAuth(state string) (model.Session, error) {
	return &Session{}, nil
}

// FetchUser is used only for testing.
func (p *Provider) FetchUser(session model.Session) (model.User, error) {
	sess := session.(*Session)
	return model.User{
		Name:  sess.Name,
		Email: sess.Email,
	}, nil
}

// UnmarshalSession is used only for testing.
func (p *Provider) UnmarshalSession(data string) (model.Session, error) {
	sess := &Session{}
	err := json.NewDecoder(strings.NewReader(data)).Decode(sess)
	return sess, err
}

// Debug is used only for testing.
func (p *Provider) Debug(debug bool) {}

// Authorize is used only for testing.
func (p *Session) Authorize(provider model.Provider, params model.Params) (string, error) {
	return "", nil
}

// Marshal is used only for testing.
func (p *Session) Marshal() string {
	b, _ := json.Marshal(p)
	return string(b)
}

// GetAuthURL is used only for testing.
func (p *Session) GetAuthURL() (string, error) {
	return "http://example.com/auth/", nil
}
