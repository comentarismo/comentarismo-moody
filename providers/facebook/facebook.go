// Package facebook implements the OAuth2 protocol for authenticating users through Facebook.
// This package can be used as a reference implementation of an OAuth2 provider for
package facebook

import (
	"encoding/json"
	"net/http"
	"net/url"

	model "comentarismo-moody/model"
	"errors"
)

// New creates a new Facebook provider, and sets up important connection details.
// You should always call `facebook.New` to get a new Provider. Never try to create
// one manually.
func New(clientKey, secret, callbackURL string, scopes ...string) *Provider {
	p := &Provider{
		ClientKey: clientKey,
		Secret:    secret,
	}
	return p
}

// Provider is the implementation of `Provider` for accessing Facebook.
type Provider struct {
	ClientKey     string
	Secret        string
	ID            string // id
	Language      string // lang
	PageName      string
	PageID        string
	Type          string // type
	Caption       string // message / description
	TotalLikes    uint64 // likes.summary.total_count
	TotalComments uint64 // comments.summary.total_count
	Thumbnail     string // picture
	PublishedAt   string // created_time
}

// Name is the name used to retrieve this provider later.
func (p *Provider) Name() string {
	return "facebook"
}

func (p *Provider) SetID(urlParts []string) error {
	p.ID = urlParts[len(urlParts)-1]
	p.PageName = urlParts[len(urlParts)-2]
	return nil
}

func (p *Provider) SetLang(lang string) error {
	p.Language = lang
	return nil
}

func (p *Provider) SetReport(theReport *model.Report, comments model.CommentList) {
	theReport.Type = "FacebookPost"
	theReport.ID = p.ID
	theReport.PublishedAt = p.PublishedAt
	theReport.TotalComments = p.TotalComments
	theReport.Metadata = p
}

// Debug is a no-op for the facebook package.
func (p *Provider) Debug(debug bool) {}

func (this *Provider) GetMetadata() bool {
	this.GetPageID()

	var respTyped postMetaResp
	resp, _ := fbRequest("/"+this.PageID+"_"+this.ID+"?fields=id,name,caption,description,picture,created_time,type,message,properties,insights,likes.limit(1).summary(true),comments.limit(1).summary(true)", this.ClientKey, this.Secret)

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(&respTyped)
	if err != nil {
		return false
	}

	this.Type = respTyped.Type
	this.Caption = respTyped.Message
	this.TotalLikes = respTyped.Likes.Summery.TotalCount
	this.TotalComments = respTyped.Comments.Summary.TotalCount
	this.Thumbnail = respTyped.Picture
	this.PublishedAt = respTyped.CreatedTime

	return true
}

func (this *Provider) GetComments() model.CommentList {
	this.GetPageID()

	var comments = []*model.Comment{}
	after := ""

	for {
		var respTyped postCommentListResp
		resp, _ := fbRequest("/"+this.PageID+"_"+this.ID+"/comments?limit=100&order=reverse_chronological&after="+after, this.ClientKey, this.Secret)

		defer resp.Body.Close()

		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&respTyped)

		if err == nil {
			for _, entry := range respTyped.Data {
				thisComment := &model.Comment{
					ID:         entry.ID,
					Published:  entry.CreatedOn,
					Content:    entry.Message,
					AuthorName: entry.From.Name,
				}

				comments = append(comments, thisComment)
			}

			if respTyped.Pagination.Cursors.After != "" {
				after = respTyped.Pagination.Cursors.After
			} else {
				break
			}
		}
	}

	return model.CommentList{Comments: comments}
}

func (this *Provider) GetPageID() model.Provider {
	if this.PageID != "" {
		return this
	}

	var respTyped pageNameResp
	resp, _ := fbRequest("/"+this.PageName, this.ClientKey, this.Secret)

	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	err := decoder.Decode(&respTyped)
	if err != nil {
		return this
	}

	this.PageID = respTyped.ID

	return this
}

func fbRequest(path, FacebookKey, FacebookSecret string) (*http.Response, error) {
	u, err := url.Parse(path)
	if err != nil {
		return nil, errors.New("FB request path invalid.")
	}

	u.Scheme = "https"
	u.Host = "graph.facebook.com"

	query := u.Query()
	query.Add("access_token", FacebookKey+"|"+FacebookSecret)
	u.RawQuery = query.Encode()

	response, err := http.Get(u.String())
	if err != nil {
		return nil, err
	} else {
		return response, nil
	}
}

type pageNameResp struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type postCommentFromResp struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type postCommentResp struct {
	From      postCommentFromResp `json:"from"`
	Message   string              `json:"message"`
	CreatedOn string              `json:"created_time"`
	ID        string              `json:"id"`
}

type postCommentPagination struct {
	Cursors struct {
		After  string `json:"after,omitempty"`
		Before string `json:"before,omitempty"`
	} `json:"cursors"`
}

type postCommentListResp struct {
	Data       []postCommentResp     `json:"data"`
	Pagination postCommentPagination `json:"paging"`
}

type postMetaProps struct {
	Name string `json:"name"`
	Text string `json:"text"`
}

type likesSummary struct {
	TotalCount uint64 `json:"total_count"`
	CanLike    bool   `json:"can_like"`
	HasLiked   bool   `json:"has_liked"`
}

type commentSummary struct {
	Order      string `json:"order"`
	TotalCount uint64 `json:"total_count"`
	CanComment bool   `json:"can_comment"`
}

type postMetaResp struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Picture     string          `json:"picture"`
	CreatedTime string          `json:"created_time"`
	Type        string          `json:"type"`
	Message     string          `json:"message"`
	Properties  []postMetaProps `json:"properties"`
	Likes       struct {
		Data []struct {
			ID string `json:"id"`
		} `json:"data"`
		Paging struct {
			Cursors struct {
				After  string `json:"after"`
				Before string `json:"before"`
			}
			Next string `json:"next"`
		} `json:"paging"`
		Summery likesSummary `json:"summary"`
	} `json:"likes"`
	Comments struct {
		Data   []postCommentResp
		Paging struct {
			Cursors struct {
				After  string `json:"after"`
				Before string `json:"before"`
			}
			Next string `json:"next"`
		} `json:"paging"`
		Summary commentSummary `json:"summary"`
	}
}
