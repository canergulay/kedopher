package enum

type MessageType string

const (
    INIT_SIGNALING MessageType = "INIT_SIGNALING"

    TRIGGER_ICE_CANDIDATES MessageType = "TRIGGER_ICE_CANDIDATES"

    OFFER MessageType = "OFFER"

    OFFER_BACK MessageType = "OFFER_BACK"

	ANSWER MessageType = "ANSWER"

	CANDIDATE MessageType = "CANDIDATE"
)
