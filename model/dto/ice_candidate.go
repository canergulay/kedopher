package dto

import "github.com/canergulay/go-betternews-signaling/model"


type IceCandidate struct {
	ConnectionID model.ID `json:"connectionId"`
	UserID model.ID `json:"userId"`

	Candidates []candidate `json:"candidates"` 
}


type candidate struct{
	Candidate string `json:"candidate"`
	SdpMid string `json:"sdpMid"`
	SdpMLineIndex int `json:"sdpMLineIndex"`
}