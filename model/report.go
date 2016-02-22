package model


type Report struct {
	ID                     string
	URL                    string
	Type                   string
	Title                  string
	PublishedAt            string
	TotalComments          uint64
	CollectedComments      uint64
	CommentCoveragePercent float64
	CommentAvgPerDay       float64
	Keywords               map[string]int

	Sentiment              []SentimentTag
	SentimentList		map[string][]*Comment
	Metadata               Provider
	SampleComments         []*Comment
	TopComments            []*Comment
}

