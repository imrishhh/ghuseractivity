// Package response holds json structure of the Github event api
package response

import "encoding/json"

type Event struct {
	ID        string          `json:"id"`
	Type      string          `json:"type"`
	Actor     Actor           `json:"actor"`
	Repo      Repo            `json:"repo"`
	Public    bool            `json:"public"`
	Payload   json.RawMessage `json:"payload"`
	CreatedAt string          `json:"created_at"`
	Org       *Org            `json:"org"`
}

type Org struct {
	ID         int    `json:"id"`
	Login      string `json:"login"`
	GravatarID string `json:"gravatar_id"`
	URL        string `json:"url"`
	AvatarURL  string `json:"avatar_url"`
}

type Repo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Actor struct {
	ID           int    `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarID   string `json:"gravatar_id"`
	URL          string `json:"url"`
	AvatarURL    string `json:"avatar_url"`
}

type CommitCommentEvent struct {
	Action  string `json:"action"`
	Comment string `json:"comment"`
}

type CreateEvent struct {
	Ref          string `json:"ref"`
	RefType      string `json:"ref_type"`
	FullRef      string `json:"full_ref"`
	MasterBranch string `json:"master_branch"`
	Description  string `json:"description"`
	PusherType   string `json:"pusher_type"`
}

type DeleteEvent struct {
	Ref        string `json:"ref"`
	RefType    string `json:"ref_type"`
	FullRef    string `json:"full_ref"`
	PusherType string `json:"pusher_type"`
}

type DiscussionEvent struct {
	Action     string          `json:"action"`
	Discussion json.RawMessage `json:"discussion"`
}

type ForkEvent struct {
	Action string          `json:"action"`
	Forkee json.RawMessage `json:"forkee"`
}

type GollumEvent struct {
	Pages []WikiPage `json:"pages"`
}

type WikiPage struct {
	PageName string  `json:"page_name"`
	Title    string  `json:"title"`
	Summary  *string `json:"summary"`
	Action   string  `json:"action"`
	SHA      string  `json:"sha"`
	HTMLURL  string  `json:"html_url"`
}

type IssueCommentEvent struct {
	Action string `json:"action"`
}

type IssuesEvent struct {
	Action string `json:"action"`
}

type MemberEvent struct {
	Action string `json:"action"`
}

type PublicEvent struct{}

type PullRequestEvent struct {
	Action string `json:"action"`
	Number int    `json:"number"`
}

type PullRequestReviewEvent struct {
	Action string `json:"action"`
}

type PullRequestReviewCommentEvent struct {
	Action string `json:"action"`
}

type PushEvent struct {
	RepositoryID int    `json:"repository_id"`
	PushID       int    `json:"push_id"`
	Ref          string `json:"ref"`
	Head         string `json:"head"`
	Before       string `json:"before"`
}

type ReleaseEvent struct {
	Action string `json:"action"`
}

type WatchEvent struct {
	Action string `json:"action"`
}
