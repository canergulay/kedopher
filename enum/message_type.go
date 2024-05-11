package enum

type MessageType string

const (
    INIT_CALL MessageType = "INIT_CALL"

    TRIGGER_ICE_CANDIDATES MessageType = "TRIGGER_ICE_CANDIDATES"

    OFFER_START MessageType = "OFFER_START"

    OFFER MessageType = "OFFER"

    OFFER_BACK MessageType = "OFFER_BACK"

	ANSWER MessageType = "ANSWER"

	SEND_ICE_CANDIDATE MessageType = "SEND_ICE_CANDIDATE"

	RECEIVE_ICE_CANDIDATE MessageType = "RECEIVE_ICE_CANDIDATE"

	START_CALL MessageType = "START_CALL"
)
