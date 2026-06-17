// Package parse is used to parse the payload of events
package parse

import (
	"encoding/json"

	"github.com/imrishhh/ghuseractivity/model"
)

func ParsePayLoad(event model.Event) (any, error) {
	switch event.Type {
	case "CommitCommentEvent":
		var p model.CommitCommentEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "CreateEvent":
		var p model.CreateEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "DeleteEvent":
		var p model.DeleteEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "DiscussionEvent":
		var p model.DiscussionEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "ForkEvent":
		var p model.ForkEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "GollumEvent":
		var p model.GollumEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "IssueCommentEvent":
		var p model.IssueCommentEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "IssuesEvent":
		var p model.IssuesEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "MemberEvent":
		var p model.MemberEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "PublicEvent":
		var p model.PublicEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "PushEvent":
		var p model.PushEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "ReleaseEvent":
		var p model.ReleaseEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	case "WatchEvent":
		var p model.WatchEvent
		err := json.Unmarshal(event.Payload, &p)
		return p, err
	default:
		return nil, nil
	}
}
