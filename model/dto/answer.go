package dto

import "github.com/canergulay/go-betternews-signaling/model"

type Answer struct {
	ConnectionID model.ID `json:"connectionID"`
	Sdp string `json:"sdp"`
	Type string `json:"type"`
}