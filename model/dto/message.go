package dto

import (
	"github.com/canergulay/go-betternews-signaling/enum"
)

type Message struct {
	Type enum.MessageType `json:"type"`
	Body any `json:"body"`
}


func (m Message) GetBodyAsAnswer() (Answer,bool){
	parsed,ok := m.Body.(Answer)
	return parsed,ok
}

func (m Message) GetBodyAsOffer() (Offer,bool){
	parsed,ok := m.Body.(Offer)
	return parsed,ok
}

func (m Message) GetBodyAsIceCandidate() (IceCandidates, bool) {
	parsed,ok := m.Body.(IceCandidates)
	return parsed,ok
}
