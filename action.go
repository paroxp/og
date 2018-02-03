package main

import (
	"fmt"

	"github.com/atotto/clipboard"
	"github.com/gen2brain/beeep"
)

type Action struct {
	Response
}

func (a *Action) Distribute() (*Response, error) {
	switch a.Type {
	case "clipboard":
		if _, ok := a.Body.(string); !ok {
			return nil, fmt.Errorf("clipboard: string body expected")
		}
		err := clipboard.WriteAll(a.Body.(string))
		return &Response{Type: a.Type, Action: a.Action, Message: "copied"}, err
	// case "media":
	//
	// 	return nil, nil
	case "notification":
		if _, ok := a.Body.(string); !ok {
			return nil, fmt.Errorf("notification: string body expected")
		}
		err := beeep.Notify(a.Body.(string), a.Message, a.Meta["image"])
		return &Response{Type: a.Type, Action: a.Action, Message: "notified"}, err
	case "volume":
		v, err := actionToVolume(a)
		if err != nil {
			return nil, err
		}

		return v.AdjustVolume()
	default:
		return nil, fmt.Errorf("unrecognised action: %s", a.Type)
	}
}
