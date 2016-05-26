package model

type Report struct {
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
	Metadata               Provider `schema:"metadata" gorethink:"metadata,omitempty" json:"metadata,omitempty"`
	//SampleComments         []*Comment `schema:"samplecomments" gorethink:"samplecomments,omitempty" json:"samplecomments,omitempty"`
	TopComments            []*Comment `schema:"topcomments" gorethink:"topcomments,omitempty" json:"topcomments,omitempty"`
}

