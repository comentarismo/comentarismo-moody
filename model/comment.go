package model

import (
	"log"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

// commentaries_sentiment_report table
type Comment struct {
	ID        string    `schema:"id" gorethink:"id,omitempty" json:"id,omitempty"`
	Published string    `schema:"published" gorethink:"published,omitempty" json:"published,omitempty"`
	Title     string    `schema:"title" gorethink:"title,omitempty" json:"title,omitempty"`
	Comment   string    `schema:"comment" gorethink:"comment,omitempty" json:"comment,omitempty"`
	Nick      string    `schema:"nick" gorethink:"nick,omitempty" json:"nick,omitempty"`
	NickId    string    `schema:"nick_id" gorethink:"nick_id,omitempty" json:"nick_id,omitempty"`
	NickIcon  string    `schema:"nickicon" gorethink:"nickicon,omitempty" json:"nickicon,omitempty"`
	Keywords  []string  `schema:"keywords" gorethink:"keywords,omitempty" json:"keywords,omitempty"`
	Language  string    `schema:"language" gorethink:"language,omitempty" json:"language,omitempty"`
	Operator  string    `schema:"operator" gorethink:"operator,omitempty" json:"operator,omitempty"`
	Type      string    `schema:"type" gorethink:"type,omitempty" json:"type,omitempty"`
	UpdatedAt time.Time `schema:"updatedAt" gorethink:"updatedAt" json:"updateAt"`
	UUID      string    `schema:"uuid" gorethink:"uuid" json:"uuid"`

	Likes          []string `schema:"likes" gorethink:"likes" json:"likes"`
	DisLikes       []string `schema:"dislikes" gorethink:"dislikes" json:"dislikes"`
	OriginDislikes int      `schema:"origindislikes" gorethink:"origindislikes,omitempty" json:"origindislikes,omitempty"`
	OriginLikes    int      `schema:"originlikes" gorethink:"originlikes,omitempty" json:"originlikes,omitempty"`

	Sentiment       int               `schema:"sentiment" gorethink:"sentiment,omitempty" json:"sentiment,omitempty"`
	Tag             string            `schema:"tag" gorethink:"tag,omitempty" json:"tag,omitempty"`
	SentimentScores map[string]string `schema:"sentimentscores" gorethink:"sentimentscores,omitempty" json:"sentimentscores,omitempty"`

	// AuthorChannelUrl: Link to the author's YouTube channel, if any.
	//AuthorChannelUrl string `schema:"authorChannelUrl" gorethink:"authorChannelUrl,omitempty" json:"authorChannelUrl,omitempty"`
	ProfileURL string `schema:"profileURL" gorethink:"profileURL,omitempty" json:"profileURL,omitempty"`

	// AuthorGoogleplusProfileUrl: Link to the author's Google+ profile, if any.
	AuthorGoogleplusProfileUrl string `schema:"authorGoogleplusProfileUrl" gorethink:"authorGoogleplusProfileUrl,omitempty" json:"authorGoogleplusProfileUrl,omitempty"`

	// AuthorProfileImageUrl: The URL for the avatar of the user who posted the comment.

	// ModerationStatus: The comment's moderation status. Will not be set if
	// the comments were requested through the id filter.
	//
	// Possible values:
	//   "heldForReview"
	//   "likelySpam"
	//   "published"
	//   "rejected"
	ModerationStatus string `schema:"moderationStatus" gorethink:"moderationStatus,omitempty" json:"moderationStatus,omitempty"`

	// ChannelId: The id of the corresponding YouTube channel. In case of a
	// channel comment this is the channel the comment refers to. In case of
	// a video comment it's the video's channel.
	ChannelId string `schema:"channelId" gorethink:"channelId,omitempty" json:"channelId,omitempty"`
}

func (comment *Comment) FixID() {
	comment.ID = strings.Replace(comment.ID, ".", "-", -1)
}

// Comment methods
//
func (comment *Comment) GetSentiment() int {
	var scores map[string]float64
	var logScores map[string]map[string]int64

	if comment.Sentiment == 0 {
		//comment.SentimentScores,
		//comment.LogScores
		var sent string
		sent, comment.Tag, scores, logScores = GetSentiment(comment.Language, comment.Comment)
		scoresFinal := make(map[string]string, len(scores))
		var keywords []string
		for k := range scores {
			//log.Println("scores ",k)
			for k1, v1 := range logScores {
				if k == k1 {
					//log.Println("logScores ",k1,v1)
					for k2, v2 := range v1 {
						keywords = appendIfMissing(keywords, k2)
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

		sentimentInt, err := strconv.Atoi(comment.Tag)
		if err != nil {
			log.Println("GetSentiment, error: ", sent, err)
		} else {
			comment.Sentiment = sentimentInt
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

func (commentList *CommentList) GetSentimentScores() (ret []map[string]string) {
	for _, comment := range commentList.Comments {
		if len(comment.SentimentScores) > 0 {
			ret = append(ret, comment.SentimentScores)
		}
	}
	log.Println("GetSentimentScores -> ", len(ret))
	return
}

func (commentList *CommentList) GetSentimentList() map[string][]*Comment {
	tags := map[string][]*Comment{}

	for _, comment := range commentList.Comments {
		tag := comment.GetSentiment()
		//cleanup object leaving only the ID to be persisted
		c := &Comment{
			ID: comment.ID,
		}
		newComments := append(tags[strconv.Itoa(tag)], c)
		tags[strconv.Itoa(tag)] = newComments
	}
	return tags
}

func (commentList *CommentList) GetSentimentSummary() []SentimentTag {
	tags := map[int]int{}

	for _, comment := range commentList.Comments {
		tag := comment.GetSentiment()
		tags[tag]++
	}

	summary := []SentimentTag{}

	for tag, count := range tags {
		st := SentimentTag{
			Name:    strconv.Itoa(tag),
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
func (a ByLikes) Less(i, j int) bool { return a[i].OriginLikes > a[j].OriginDislikes }

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
