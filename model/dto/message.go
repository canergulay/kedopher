package dto

import (
	"encoding/json"
	"errors"

	"github.com/canergulay/go-betternews-signaling/model/enum"
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

func (m Message) GetBodyAsTriggerIceCandidates() (TriggerIceCandidates, bool) {
	parsed,ok := m.Body.(TriggerIceCandidates)
	return parsed,ok
}



func (m *Message) UnmarshalJSON(data []byte) error {
	type Alias Message
	aux := &struct {
		*Alias
		Body json.RawMessage `json:"body"`
	}{
		Alias: (*Alias)(m),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	switch m.Type {
	case enum.INIT_CALL:
	case enum.OFFER:
		var body Offer
		if err := json.Unmarshal(aux.Body, &body); err != nil {
			return err
		}
		m.Body = body
	case enum.ANSWER:
		var body Answer
		if err := json.Unmarshal(aux.Body, &body); err != nil {
			return err
		}
		m.Body = body
	case enum.ICE_CANDIDATES:
		var body IceCandidates
		if err := json.Unmarshal(aux.Body, &body); err != nil {
			return err
		}
		m.Body = body
	case enum.TRIGGER_ICE_CANDIDATE:
		var body TriggerIceCandidates
		if err := json.Unmarshal(aux.Body, &body); err != nil {
			return err
		}
		m.Body = body
	default:
		return errors.New("unknown message type")
	}
	return nil
}