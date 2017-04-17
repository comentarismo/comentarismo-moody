package model

import "time"

type Metadata struct {
	ClientKey     string    `schema:"clientkey" gorethink:"clientkey,omitempty" json:"clientkey,omitempty"`
	Secret        string    `schema:"secret" gorethink:"secret,omitempty" json:"secret,omitempty"`
	ID            string    `schema:"id" gorethink:"id,omitempty" json:"id,omitempty"`
	Language      string    `schema:"language" gorethink:"language,omitempty" json:"language,omitempty"`
	Title         string    `schema:"title" gorethink:"title,omitempty" json:"title,omitempty"`
	VideoViews    uint64    `schema:"videoviews" gorethink:"videoviews,omitempty" json:"videoviews,omitempty"`
	ChannelID     string    `schema:"channelid" gorethink:"channelid,omitempty" json:"channelid,omitempty"`
	ChannelTitle  string    `schema:"channeltitle" gorethink:"channeltitle,omitempty" json:"channeltitle,omitempty"`
	TotalComments uint64    `schema:"totalcomments" gorethink:"totalcomments,omitempty" json:"totalcomments,omitempty"`
	Thumbnail     string    `schema:"thumbnail" gorethink:"thumbnail,omitempty" json:"thumbnail,omitempty"`
	PublishedAt   string    `schema:"publishedat" gorethink:"publishedat,omitempty" json:"publishedat,omitempty"`
	UpdatedAt     time.Time `schema:"updatedAt" gorethink:"updatedAt" json:"updateAt"`
}
