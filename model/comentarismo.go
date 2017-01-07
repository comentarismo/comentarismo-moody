package model

import (
	"os"
	"time"
)

var ComentarismoKey = os.Getenv("youtubekey")

type Comentarismo struct {
	ID          string   `gorethink:"id,omitempty" json:"id,omitempty"`
	Titleurlize string   `gorethink:"titleurlize,omitempty" json:"titleurlize,omitempty"`
	Countries   string   `gorethink:"countries,omitempty" json:"countries,omitempty"`
	Languages   string   `gorethink:"languages,omitempty" json:"languages,omitempty"`
	Operator    string   `gorethink:"operator,omitempty" json:"operator,omitempty"`
	Title       string   `gorethink:"title,omitempty" json:"title,omitempty"`
	Categories  string   `gorethink:"categories,omitempty" json:"categories,omitempty"`
	Genre       string   `gorethink:"genre,omitempty" json:"genre,omitempty"`
	Comment     string   `gorethink:"comment,omitempty" json:"comment,omitempty"`
	Likes       []string `gorethink:"likes" json:"likes"`
	DisLikes    []string `gorethink:"dislikes" json:"dislikes"`

	OriginDislikes int            `gorethink:"origindislikes,omitempty" json:"origindislikes,omitempty"`
	OriginLikes    int            `gorethink:"originlikes,omitempty" json:"originlikes,omitempty"`
	Nick           string         `gorethink:"nick,omitempty" json:"nick,omitempty"`
	NickId         string         `gorethink:"nick_id,omitempty" json:"nick_id,omitempty"`
	NickIcon       string         `gorethink:"nickicon,omitempty" json:"nickicon,omitempty"`
	Date           string         `gorethink:"date,omitempty" json:"date,omitempty"`
	Created        time.Time      `gorethink:"created,omitempty" json:"created,omitempty"`
	AvatarURL      string         `gorethink:"avatarurl,omitempty" json:"avatarurl,omitempty"`
	Replies        []Comentarismo `gorethink:"replies" json:"replies"`
}
