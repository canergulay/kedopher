package dto

import "github.com/canergulay/go-betternews-signaling/model/enum"

type UserStatus struct {
	Status enum.UserStatus `json:"status"`
}