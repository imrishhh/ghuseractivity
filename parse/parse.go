// Package parse is used to parse the payload of events
package parse

import (
	"encoding/json"

	"github.com/imrishhh/ghuseractivity/response"
)

func ParsePayLoad(event response.Event) (any, error) {
	switch event.Type {
	case "CommitCommentEvent":
		var p response.CommitCommentEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "CreateEvent":
		var p response.CreateEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "DeleteEvent":
		var p response.DeleteEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "DiscussionEvent":
		var p response.DiscussionEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "ForkEvent":
		var p response.ForkEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "GollumEvent":
		var p response.GollumEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "IssueCommentEvent":
		var p response.IssueCommentEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "IssuesEvent":
		var p response.IssuesEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "MemberEvent":
		var p response.MemberEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "PublicEvent":
		var p response.PublicEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "PushEvent":
		var p response.PushEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "ReleaseEvent":
		var p response.ReleaseEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "WatchEvent":
		var p response.Event
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	default:
		return nil, nil
	}
}
