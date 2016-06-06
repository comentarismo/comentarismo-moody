package lang

import (
	"time"
)

type DetectLanguage struct {
	Text     string         `gorethink:"text,omitempty" json:"text,omitempty"`
	Language string         `gorethink:"language,omitempty" json:"language,omitempty"`
	Error    string         `gorethink:"error,omitempty" json:"error,omitempty"`
	Date 	 time.Time	`gorethink:"date,omitempty" json:"date,omitempty"`
}

