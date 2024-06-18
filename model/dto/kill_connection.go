package dto

import "github.com/canergulay/go-betternews-signaling/model"

type KillConnection struct {
	ConnectionID model.ID `json:"connectionId"`
}