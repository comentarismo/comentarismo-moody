package model

import "time"

type Report struct {
	ID                     string `schema:"id" gorethink:"id,omitempty" json:"id,omitempty"`
	URL                    string `schema:"url" gorethink:"url,omitempty" json:"url,omitempty"`
	Type                   string `schema:"type" gorethink:"type,omitempty" json:"type,omitempty"`
	Title                  string `schema:"title" gorethink:"title,omitempty" json:"title,omitempty"`
	Date                   time.Time `schema:"date" gorethink:"date,omitempty" json:"date,omitempty"`
	PublishedAt            string `schema:"publishedat" gorethink:"publishedat,omitempty" json:"publishedat,omitempty"`
	TotalComments          uint64 `schema:"totalcomments" gorethink:"totalcomments,omitempty" json:"totalcomments,omitempty"`
	CollectedComments      uint64 `schema:"collectedcomments" gorethink:"collectedcomments,omitempty" json:"collectedcomments,omitempty"`
	CommentCoveragePercent float64 `schema:"commentcoveragepercent" gorethink:"commentcoveragepercent,omitempty" json:"commentcoveragepercent,omitempty"`
	CommentAvgPerDay       float64 `schema:"commentavgperday" gorethink:"commentavgperday,omitempty" json:"commentavgperday,omitempty"`
	Keywords               map[string]int `schema:"keywords" gorethink:"keywords,omitempty" json:"keywords,omitempty"`
	Sentiment              []SentimentTag `schema:"sentiment" gorethink:"sentiment,omitempty" json:"sentiment,omitempty"`
	SentimentList          map[string][]*Comment `schema:"sentimentlist" gorethink:"sentimentlist,omitempty" json:"sentimentlist,omitempty"`
	Metadata               Provider `schema:"metadata" gorethink:"metadata,omitempty" json:"metadata,omitempty"`
	//SampleComments         []*Comment `schema:"samplecomments" gorethink:"samplecomments,omitempty" json:"samplecomments,omitempty"`
	TopComments            []*Comment `schema:"topcomments" gorethink:"topcomments,omitempty" json:"topcomments,omitempty"`
}

type YoutubeReport struct {
	ID                     string `schema:"id" gorethink:"id,omitempty" json:"id,omitempty"`
	URL                    string `schema:"url" gorethink:"url,omitempty" json:"url,omitempty"`
	Type                   string `schema:"type" gorethink:"type,omitempty" json:"type,omitempty"`
	Title                  string `schema:"title" gorethink:"title,omitempty" json:"title,omitempty"`
	PublishedAt            string `schema:"publishedat" gorethink:"publishedat,omitempty" json:"publishedat,omitempty"`
	TotalComments          uint64 `schema:"totalcomments" gorethink:"totalcomments,omitempty" json:"totalcomments,omitempty"`
	CollectedComments      uint64 `schema:"collectedcomments" gorethink:"collectedcomments,omitempty" json:"collectedcomments,omitempty"`
	CommentCoveragePercent float64 `schema:"commentcoveragepercent" gorethink:"commentcoveragepercent,omitempty" json:"commentcoveragepercent,omitempty"`
	CommentAvgPerDay       float64 `schema:"commentavgperday" gorethink:"commentavgperday,omitempty" json:"commentavgperday,omitempty"`
	Keywords               map[string]int `schema:"keywords" gorethink:"keywords,omitempty" json:"keywords,omitempty"`
	Sentiment              []SentimentTag `schema:"sentiment" gorethink:"sentiment,omitempty" json:"sentiment,omitempty"`
	SentimentList          map[string][]*Comment `schema:"sentimentlist" gorethink:"sentimentlist,omitempty" json:"sentimentlist,omitempty"`
	Metadata               Metadata `schema:"metadata" gorethink:"metadata,omitempty" json:"metadata,omitempty"`
	//SampleComments         []*Comment `schema:"samplecomments" gorethink:"samplecomments,omitempty" json:"samplecomments,omitempty"`
	TopComments            []*Comment `schema:"topcomments" gorethink:"topcomments,omitempty" json:"topcomments,omitempty"`
}

type Metadata struct {
	ClientKey     string `schema:"clientkey" gorethink:"clientkey,omitempty" json:"clientkey,omitempty"`
	Secret        string `schema:"secret" gorethink:"secret,omitempty" json:"secret,omitempty"`
	ID            string `schema:"id" gorethink:"id,omitempty" json:"id,omitempty"`
	Language      string `schema:"language" gorethink:"language,omitempty" json:"language,omitempty"`
	Title         string `schema:"title" gorethink:"title,omitempty" json:"title,omitempty"`
	VideoViews    uint64 `schema:"videoviews" gorethink:"videoviews,omitempty" json:"videoviews,omitempty"`
	ChannelID     string `schema:"channelid" gorethink:"channelid,omitempty" json:"channelid,omitempty"`
	ChannelTitle  string `schema:"channeltitle" gorethink:"channeltitle,omitempty" json:"channeltitle,omitempty"`
	TotalComments uint64 `schema:"totalcomments" gorethink:"totalcomments,omitempty" json:"totalcomments,omitempty"`
	Thumbnail     string `schema:"thumbnail" gorethink:"thumbnail,omitempty" json:"thumbnail,omitempty"`
	PublishedAt   string `schema:"publishedat" gorethink:"publishedat,omitempty" json:"publishedat,omitempty"`
}

// SentimentTag is a list entry of the tag and the percent of comments that were classified with that tag.
type SentimentTag struct {
	Name    string `schema:"name" gorethink:"url,omitempty" json:"name,omitempty"`
	Percent float64 `schema:"percent" gorethink:"percent,omitempty" json:"percent,omitempty"`
}

// Comment is the distilled comment dataset
type Comment struct {
	ID              string `schema:"id" gorethink:"id,omitempty" json:"id,omitempty"`
	Published       string `schema:"published" gorethink:"published,omitempty" json:"published,omitempty"`
	Title           string `schema:"title" gorethink:"title,omitempty" json:"title,omitempty"`
	Content         string `schema:"content" gorethink:"content,omitempty" json:"content,omitempty"`
	AuthorName      string `schema:"authorname" gorethink:"authorname,omitempty" json:"authorname,omitempty"`
	Sentiment       string `schema:"sentiment" gorethink:"sentiment,omitempty" json:"sentiment,omitempty"`
	Tag             string `schema:"tag" gorethink:"tag,omitempty" json:"tag,omitempty"`
	SentimentScores map[string]string `schema:"sentimentscores" gorethink:"sentimentscores,omitempty" json:"sentimentscores,omitempty"`
	Keywords        []string `schema:"keywords" gorethink:"keywords,omitempty" json:"keywords,omitempty"`
	Likes           int64  `schema:"likes" gorethink:"likes,omitempty" json:"likes,omitempty"`
	Language        string `schema:"language" gorethink:"language,omitempty" json:"language,omitempty"`
}

type CommentList struct {
	Comments []*Comment
}