// Package youtube implements the OAuth2 protocol for authenticating users through Youtube.
// This package can be used as a reference implementation of an OAuth2 provider
package youtube

import (
	"net/http"

	"github.com/google/google-api-go-client/googleapi/transport"
	youtube "google.golang.org/api/youtube/v3"

	"comentarismo-moody/model"
	"log"
	"os"
	"strconv"
	"time"
)

var YOUTUBE_REPORT_MAX_STR = os.Getenv("YOUTUBE_REPORT_MAX")
var YOUTUBE_REPORT_MAX = 3000

func init() {

	if YOUTUBE_REPORT_MAX_STR == "" {
		YOUTUBE_REPORT_MAX = 3000
	}
}

// New creates a new Youtube provider, and sets up important connection details.
// You should always call `youtube.New` to get a new Provider. Never try to create
// one manually.
func New(clientKey, secret, callbackURL string, scopes ...string) *Provider {
	p := &Provider{
		ClientKey: clientKey,
		Secret:    secret,
	}
	return p
}

// Provider is the implementation of `Provider` for accessing Youtube.
//Provider.Secret is not used for youtube but is here just in case
type Provider struct {
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

// Name is the name used to retrieve this provider later.
func (p *Provider) Name() string {
	return "youtube"
}

// Name is the name used to retrieve this report type later.
func (p *Provider) GetType() string {
	return "YouTubeVideo"
}

func (p *Provider) SetID(urlParts []string) error {
	i := len(urlParts) - 1
	p.ID = urlParts[i]
	return nil
}

func (p *Provider) SetLang(lang string) error {
	p.Language = lang
	return nil
}

func (p *Provider) SetReport(theReport *model.Report, comments model.CommentList) {
	theReport.Operator = p.GetType()
	theReport.ID = p.ID
	theReport.Title = p.Title
	theReport.PublishedAt = p.PublishedAt
	theReport.TotalComments = p.TotalComments
	theReport.Metadata = p
	theReport.UpdatedAt = time.Now()
	//clear key and secret fields
	theReport.TopComments = comments.GetMostLiked(10)
}

// Debug is a no-op for the package.
func (p *Provider) Debug(debug bool) {}

func (this *Provider) GetPageID() model.Provider {
	return this
}

// YouTubeGetCommentsV2 pulls the comments for a given YouTube video
func (ytv *Provider) GetComments() model.CommentList {
	videoID := ytv.ID
	log.Println("videoID,", videoID)
	var comments = []*model.Comment{}

	client := &http.Client{
		Transport: &transport.APIKey{Key: ytv.ClientKey},
	}

	youtubeService, err := youtube.New(client)
	if err != nil {
		log.Println("Error: GetComments when -> youtube.New, ", err)
		return model.CommentList{Comments: comments}
	}

	pageToken := ""
	for pageToken != "EOL" {
		results, err := youtubeService.CommentThreads.List("id,snippet,replies").TextFormat("plainText").MaxResults(100).VideoId(videoID).PageToken(pageToken).Do()

		if err != nil {
			log.Println("Error: GetComments when -> youtubeService.CommentThreads, ", videoID, pageToken, err)
			return model.CommentList{Comments: comments}
		}

		if len(results.Items) > 0 {
			for _, item := range results.Items {

				tempComments := []*youtube.Comment{}
				tempComments = append(tempComments, item.Snippet.TopLevelComment)

				if item.Replies != nil {
					for _, reply := range item.Replies.Comments {
						tempComments = append(tempComments, reply)
					}
				}

				for _, c := range tempComments {
					thisComment := &model.Comment{
						ID:                         c.Id,
						Published:                  c.Snippet.PublishedAt,
						Comment:                    c.Snippet.TextDisplay,
						Nick:                       c.Snippet.AuthorDisplayName,
						NickIcon:                   c.Snippet.AuthorProfileImageUrl,
						Likes:                      c.Snippet.LikeCount,
						ProfileURL:                 c.Snippet.AuthorChannelUrl,
						AuthorGoogleplusProfileUrl: c.Snippet.AuthorGoogleplusProfileUrl,
						ModerationStatus:           c.Snippet.ModerationStatus,
						ChannelId:                  c.Snippet.ChannelId,
						Language:                   ytv.Language,
						Operator:                   ytv.Name(),
						Type:                       ytv.GetType(),
					}

					comments = append(comments, thisComment)
				}
			}
		}

		pageToken = results.NextPageToken
		if pageToken == "" || len(comments) >= YOUTUBE_REPORT_MAX {
			pageToken = "EOL"
		}
	}

	return model.CommentList{Comments: comments}
}

// YouTubeGetCommentsV2 pulls the comments for a given YouTube video
func (ytv *Provider) GetCommentsChan(resultsChannel chan *model.Comment, countChannel chan int) {
	videoID := ytv.ID
	log.Println("videoID,", videoID)

	client := &http.Client{
		Transport: &transport.APIKey{Key: ytv.ClientKey},
	}

	youtubeService, err := youtube.New(client)
	if err != nil {
		log.Println("Error: GetComments when -> youtube.New, ", err)
		return
	}

	pageToken := ""
	//defer resultsChannel.Close()
	var lencomments int
	for pageToken != "EOL" {
		results, err := youtubeService.CommentThreads.List("id,snippet,replies").TextFormat("plainText").MaxResults(100).VideoId(videoID).PageToken(pageToken).Do()

		if err != nil {
			log.Println("Error: GetComments when -> youtubeService.CommentThreads, ", videoID, pageToken, err)
			return
		}
		pageToken = results.NextPageToken

		if len(results.Items) > 0 {
			for _, item := range results.Items {

				tempComments := []*youtube.Comment{}
				tempComments = append(tempComments, item.Snippet.TopLevelComment)

				if item.Replies != nil {
					for _, reply := range item.Replies.Comments {
						tempComments = append(tempComments, reply)
					}
				}

				for _, c := range tempComments {
					thisComment := &model.Comment{
						ID:         c.Id,
						Published:  c.Snippet.PublishedAt,
						Comment:    c.Snippet.TextDisplay,
						Nick:       c.Snippet.AuthorDisplayName,
						NickIcon:   c.Snippet.AuthorProfileImageUrl,
						ProfileURL: c.Snippet.AuthorChannelUrl,
						Likes:      c.Snippet.LikeCount,
						AuthorGoogleplusProfileUrl: c.Snippet.AuthorGoogleplusProfileUrl,
						ModerationStatus:           c.Snippet.ModerationStatus,
						ChannelId:                  c.Snippet.ChannelId,
						Language:                   ytv.Language,
						Operator:                   ytv.Name(),
						Type:                       ytv.GetType(),
					}

					resultsChannel <- thisComment
					countChannel <- 1
					lencomments = lencomments + 1

				}
			}
		}

		if pageToken == "" || lencomments >= YOUTUBE_REPORT_MAX {
			pageToken = "EOL"
		}
	}

	countChannel <- 0

	log.Println("Finished processing GetCommentsChan")

	return
}

// GetMetadata returns a subset of video information from the YouTube API
func (ytv *Provider) GetMetadata() bool {
	videoID := ytv.ID
	log.Println("GetMetadata, ", videoID)

	client := &http.Client{
		Transport: &transport.APIKey{Key: ytv.ClientKey},
	}

	youtubeService, err := youtube.New(client)
	if err != nil {
		log.Println("Error, GetMetadata, youtube.New(client)", err)
		return false
	}

	call := youtubeService.Videos.List("id,snippet,statistics").Id(videoID)
	resp, err := call.Do()
	if err != nil {
		log.Println("Error, youtubeService.Videos.List, ", videoID, err)
		return false
	}

	if resp == nil {
		log.Println("Error: youtubeService.Videos.List, resp is nil -> ", resp)
	}

	if len(resp.Items) > 0 {
		video := resp.Items[0]

		t, _ := time.Parse(time.RFC3339Nano, video.Snippet.PublishedAt)

		ytv.Title = video.Snippet.Title
		ytv.ChannelID = video.Snippet.ChannelId
		ytv.ChannelTitle = video.Snippet.ChannelTitle
		ytv.TotalComments = video.Statistics.CommentCount
		ytv.PublishedAt = strconv.FormatInt(t.Unix(), 10)
		ytv.VideoViews = video.Statistics.ViewCount
		ytv.Thumbnail = video.Snippet.Thumbnails.High.Url

		return true
	}

	return false
}
