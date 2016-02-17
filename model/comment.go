package model

import (
	"math"
	"math/rand"
	"sort"
)

// Comment is the distilled comment dataset
type Comment struct {
	ID         string
	Published  string
	Title      string
	Content    string
	AuthorName string
	Sentiment  string
	Likes      int64
}

type CommentList struct {
	Comments []*Comment
}

// Comment methods
func (comment *Comment) GetSentiment() string {
	if comment.Sentiment == "" {
		comment.Sentiment = GetSentiment(comment.Content)
	}

	return comment.Sentiment
}

// CommentList methods

func (commentList *CommentList) IsEmpty() bool {
	if len(commentList.Comments) == 0 {
		return true
	}

	return false
}

func (commentList *CommentList) GetTotal() uint64 {
	return uint64(len(commentList.Comments))
}

func (commentList *CommentList) GetKeywords() map[string]int {
	return GetKeywords(commentList.Comments)
}

func (commentList *CommentList) GetSentimentSummary() []SentimentTag {
	tags := map[string]int{}

	for _, comment := range commentList.Comments {
		tag := comment.GetSentiment()
		tags[tag]++
	}

	summary := []SentimentTag{}

	for tag, count := range tags {
		st := SentimentTag{
			Name:    tag,
			Percent: math.Ceil((float64(count) / float64(len(commentList.Comments))) * float64(100)),
		}

		summary = append(summary, st)
	}

	return summary
}

func (commentList *CommentList) GetRandom(count int) []*Comment {
	seed := rand.NewSource(42)
	rnum := rand.New(seed)

	resp := []*Comment{}

	for i := 0; i < count; i++ {
		resp = append(resp, commentList.Comments[rnum.Intn(len(commentList.Comments))])
	}

	return resp
}

type ByLikes []*Comment

func (a ByLikes) Len() int           { return len(a) }
func (a ByLikes) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByLikes) Less(i, j int) bool { return a[i].Likes > a[j].Likes }

func (commentList *CommentList) GetMostLiked(count int) []*Comment {
	bl := ByLikes(commentList.Comments)
	sort.Sort(bl)

	resp := []*Comment{}

	for i := 0; i < count; i++ {
		resp = append(resp, bl[i])
	}
	return resp
}
