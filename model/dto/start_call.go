package dto

import "github.com/canergulay/go-betternews-signaling/model"

type StartCall struct {
	ConnectionID model.ID `json:"connectionId"`
}