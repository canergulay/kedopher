package dto

type Offer struct {
	ConnectionID string `json:"connectionID"`
	Sdp string `json:"sdp"`
	Type string `json:"type"`
}