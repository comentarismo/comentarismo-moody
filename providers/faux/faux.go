// Package faux is used exclusive for testing purposes. I would strongly suggest you move along
// as there's nothing to see here.
package faux

import (
	"comentarismo-moody/model"
)

// Provider is used only for testing.
type Provider struct {
}

// Name is used only for testing.
func (p Provider) Name() string {
	return "faux"
}

func (p Provider) GetType() string {
	return "Faux"
}

func (p *Provider) SetLang(lang string) error {
	return nil
}

func (p Provider) Debug(debug bool) {}

// Debug is used only for testing.
func (p Provider) GetComments() model.CommentList {
	return model.CommentList{}
}

// Debug is used only for testing.
func (p Provider) GetCommentsChan(resultsChannel chan *model.Comment, countChannel chan int) {
	return
}

func (p Provider) SetID([]string) error {
	return nil
}

func (p Provider) SetReport(theReport *model.Report, comments model.CommentList) {

}

// Debug is used only for testing.
func (p Provider) GetMetadata() bool {
	return false
}

func (this *Provider) GetPageID() model.Provider {
	return this
}
