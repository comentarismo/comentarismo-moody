// Package facebook implements the OAuth2 protocol for authenticating users through Facebook.
// This package can be used as a reference implementation of an OAuth2 provider for
package youtube

import (
	"net/http"

	transport "github.com/google/google-api-go-client/googleapi/transport"
	youtube "google.golang.org/api/youtube/v3"

	"comentarismo-moody/model"
	"time"
	"strconv"
	"log"
)


func init(){

}
// New creates a new Facebook provider, and sets up important connection details.
// You should always call `facebook.New` to get a new Provider. Never try to create
// one manually.
func New(clientKey, secret, callbackURL string, scopes ...string) *Provider {
	if(clientKey == ""){
		clientKey = "AIzaSyBameiyxxJw0W27lydpPuPRocfvGza9gXM"
	}
	p := &Provider{
		ClientKey:   clientKey,
		Secret:      secret,
	}
	return p
}

// Provider is the implementation of `Provider` for accessing Facebook.
type Provider struct {
	ClientKey   string
	Secret      string
	ID            string
	Title         string
	VideoViews    uint64
	ChannelID     string
	ChannelTitle  string
	TotalComments uint64
	Thumbnail     string
	PublishedAt   string
}

// Name is the name used to retrieve this provider later.
func (p *Provider) Name() string {
	return "youtube"
}

func (p *Provider) SetID(urlParts []string) error {
	i := len(urlParts)-1
	p.ID = urlParts[i]
	return nil
}

func (p *Provider) SetReport(theReport *model.Report, comments model.CommentList) {
	theReport.Type = "YouTubeVideo"
	theReport.ID = p.ID
	theReport.Title = p.Title
	theReport.PublishedAt = p.PublishedAt
	theReport.TotalComments = p.TotalComments
	theReport.Metadata = p
	theReport.TopComments = comments.GetMostLiked(10)
}



// Debug is a no-op for the facebook package.
func (p *Provider) Debug(debug bool) {}

func (this *Provider) GetPageID() model.Provider {
	return this
}

// YouTubeGetCommentsV2 pulls the comments for a given YouTube video
func (ytv *Provider) GetComments() model.CommentList {
	videoID := ytv.ID
	log.Println("videoID: ",videoID)
	var comments = []*model.Comment{}

	client := &http.Client{
		Transport: &transport.APIKey{Key: ytv.ClientKey},
	}

	youtubeService, err := youtube.New(client)
	if err != nil {
		panic(err)
	}

	pageToken := ""
	for pageToken != "EOL" {
		results, err := youtubeService.CommentThreads.List("id,snippet,replies").TextFormat("plainText").MaxResults(100).VideoId(videoID).PageToken(pageToken).Do()

		if err != nil {
			panic(err)
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
						ID:         c.Id,
						Published:  c.Snippet.PublishedAt,
						Title:      "",
						Content:    c.Snippet.TextDisplay,
						AuthorName: c.Snippet.AuthorDisplayName,
						Likes:      c.Snippet.LikeCount,
					}

					comments = append(comments, thisComment)
				}
			}
		}

		pageToken = results.NextPageToken
		if pageToken == "" {
			pageToken = "EOL"
		}
	}

	return model.CommentList{Comments: comments}
}

// GetMetadata returns a subset of video information from the YouTube API
func (ytv *Provider) GetMetadata() bool {
	videoID := ytv.ID
	log.Println(videoID)

	client := &http.Client{
		Transport: &transport.APIKey{Key: ytv.ClientKey},
	}

	youtubeService, err := youtube.New(client)
	if err != nil {
		panic(err)
	}

	call := youtubeService.Videos.List("id,snippet,statistics").Id(videoID)
	resp, err := call.Do()
	if err != nil {
		panic(err)
	}

	if(resp == nil){
		log.Println("karai")
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

