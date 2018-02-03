package main

import (
	"fmt"
	"strconv"

	volume "github.com/itchyny/volume-go"
)

type Volume struct {
	Action string
	Value  int
	Res    *Response
}

func actionToVolume(a *Action) (*Volume, error) {
	v := &Volume{
		Action: a.Action,
		Res:    &Response{Type: a.Type, Action: a.Action},
	}

	if a.Body != nil {
		if s, ok := a.Body.(string); ok {
			i, err := strconv.ParseInt(s, 10, 64)
			v.Value = int(i)
			if err != nil {
				return nil, err
			}
		} else if i, ok := a.Body.(int); ok {
			v.Value = i
		} else {
			return nil, fmt.Errorf("volume: either int or string body expected")
		}
	}

	return v, nil
}

func ifZero(value, def int) int {
	if value == 0 {
		return def
	}

	return value
}

func (v *Volume) AdjustVolume() (*Response, error) {
	switch v.Action {
	case "mute":
		v.Res.Message = "muted"
		return v.Res, volume.Mute()
	case "unmute":
		v.Res.Message = "unmuted"
		return v.Res, volume.Unmute()
	case "change":
		v.Res.Body = v.Value
		v.Res.Message = fmt.Sprintf("changed: %d", v.Value)
		return v.Res, volume.SetVolume(v.Value)
	case "increase":
		v.Res.Body = ifZero(v.Value, 10)
		v.Res.Message = fmt.Sprintf("increased by: %d", v.Res.Body)
		return v.Res, volume.IncreaseVolume(v.Res.Body.(int))
	case "decrease":
		v.Res.Body = ifZero(v.Value, 10)
		v.Res.Message = fmt.Sprintf("decreased by: %d", v.Res.Body)
		return v.Res, volume.IncreaseVolume(-v.Res.Body.(int))
	case "check-mute":
		muted, err := volume.GetMuted()
		v.Res.Body = muted
		v.Res.Message = fmt.Sprintf("muted: %t", muted)
		return v.Res, err
	case "check-volume":
		i, err := volume.GetVolume()
		v.Res.Body = i
		v.Res.Message = fmt.Sprintf("volume: %d", i)
		return v.Res, err
	default:
		return v.Res, fmt.Errorf("unrecognised volume action: %s", v.Action)
	}
}
