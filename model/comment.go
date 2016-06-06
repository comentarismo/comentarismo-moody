package model

import (
	"math"
	"math/rand"
	"sort"
)

// Comment methods
func (comment *Comment) GetSentiment() string {
	var scores map[string]float64
	var logScores map[string]map[string]int64

	if comment.Sentiment == "" {
		//comment.SentimentScores,
		//comment.LogScores
		comment.Sentiment, scores, logScores = GetSentiment(comment.Language,comment.Content)
		scoresFinal := make(map[string]string,len(scores))
		var keywords []string
		for k,_ := range  scores {
			//log.Println("scores ",k)
			for k1,v1 := range logScores {
				if k == k1 {
					//log.Println("logScores ",k1,v1)
					for k2,v2 := range v1 {
						keywords = appendIfMissing(keywords,k2)
						if v2 > 0 {
							//log.Println("v2 ",v2)
							//log.Println("scoresFinal[k2] ",scoresFinal[k2])
							//log.Println("k ",k)
							scoresFinal[k2] = k
						}
					}


				}
			}

		}

		comment.SentimentScores = scoresFinal
		comment.Keywords = keywords

	}

	return comment.Sentiment
}

func appendIfMissing(slice []string, i string) []string {
	for _, ele := range slice {
		if ele == i {
			return slice
		}
	}
	return append(slice, i)
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

func  (commentList *CommentList) GetSentimentList() map[string][]*Comment{
	tags := map[string][]*Comment{}

	for _, comment := range commentList.Comments {
		tag := comment.GetSentiment()
		newComments := append(tags[tag], comment)
		tags[tag] = newComments
	}
	return tags
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
	if commentList == nil {
		return []*Comment{}
	}
	bl := ByLikes(commentList.Comments)
	if len(bl) == 0 {
		return []*Comment{}
	}

	sort.Sort(bl)

	resp := []*Comment{}

	for i := 0; i < count && i < len(bl); i++ {
		resp = append(resp, bl[i])
	}
	return resp
}
