package enum

type UserStatus int

const (
    UserIdle UserStatus = iota

    UserWaiting

    UserNotifiedForIce

    UserSentIceCandidate

    UserMatched
)