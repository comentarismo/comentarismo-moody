package model

import "time"

//sentiment_report table
type Report struct {
	ID                     string                `schema:"id" gorethink:"id,omitempty" json:"id,omitempty"`
	URL                    string                `schema:"url" gorethink:"url,omitempty" json:"url,omitempty"`
	Type                   string                `schema:"type" gorethink:"type,omitempty" json:"type,omitempty"`
	Operator               string                `schema:"operator" gorethink:"operator,omitempty" json:"operator,omitempty"`
	Title                  string                `schema:"title" gorethink:"title,omitempty" json:"title,omitempty"`
	Date                   time.Time             `schema:"date" gorethink:"date,omitempty" json:"date,omitempty"`
	PublishedAt            string                `schema:"publishedat" gorethink:"publishedat,omitempty" json:"publishedat,omitempty"`
	TotalComments          uint64                `schema:"totalcomments" gorethink:"totalcomments,omitempty" json:"totalcomments,omitempty"`
	CollectedComments      uint64                `schema:"collectedcomments" gorethink:"collectedcomments,omitempty" json:"collectedcomments,omitempty"`
	CommentCoveragePercent float64               `schema:"commentcoveragepercent" gorethink:"commentcoveragepercent,omitempty" json:"commentcoveragepercent,omitempty"`
	CommentAvgPerDay       float64               `schema:"commentavgperday" gorethink:"commentavgperday,omitempty" json:"commentavgperday,omitempty"`
	Keywords               map[string]int        `schema:"keywords" gorethink:"keywords,omitempty" json:"keywords,omitempty"`
	SentimentScores        []map[string]string   `schema:"sentimentscores" gorethink:"sentimentscores,omitempty" json:"sentimentscores,omitempty"`
	Sentiment              []SentimentTag        `schema:"sentiment" gorethink:"sentiment,omitempty" json:"sentiment,omitempty"`
	SentimentList          map[string][]*Comment `schema:"sentimentlist" gorethink:"sentimentlist,omitempty" json:"sentimentlist,omitempty"`
	Metadata               Provider              `schema:"metadata" gorethink:"metadata,omitempty" json:"metadata,omitempty"`
	TopComments            []*Comment            `schema:"topcomments" gorethink:"topcomments,omitempty" json:"topcomments,omitempty"`
	UpdatedAt              time.Time             `schema:"updatedAt" gorethink:"updatedAt" json:"updateAt"`

	//SampleComments         []*Comment `schema:"samplecomments" gorethink:"samplecomments,omitempty" json:"samplecomments,omitempty"`
}

//type KeywordMap struct {
//	Sentiment string `schema:"tag" gorethink:"tag,omitempty" json:"tag,omitempty"`
//	Count     int    `schema:"count" gorethink:"count,omitempty" json:"count,omitempty"`
//	Keyword   string `schema:"keyword" gorethink:"keyword,omitempty" json:"keyword,omitempty"`
//}

// SentimentTag is a list entry of the tag and the percent of comments that were classified with that tag.
type SentimentTag struct {
	Name    string  `schema:"name" gorethink:"url,omitempty" json:"name,omitempty"`
	Percent float64 `schema:"percent" gorethink:"percent,omitempty" json:"percent,omitempty"`
}

type CommentList struct {
	Comments []*Comment
}
