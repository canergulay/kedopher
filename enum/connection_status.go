package enum

type ConnectionStatus int

const (
    ConnectionInitial ConnectionStatus = iota

	ConnectionWaitingForIceExchange

	ConnectionEstablished
)