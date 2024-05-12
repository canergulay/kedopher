package dto

import "github.com/canergulay/go-betternews-signaling/model"


type IceCandidates struct {
	ConnectionID model.ID `json:"connectionId"`

	Candidates []candidate `json:"candidates"` 
}


type candidate struct{
	Candidate string `json:"candidate"`
	SdpMid string `json:"sdpMid"`
	SdpMLineIndex int `json:"sdpMLineIndex"`
}