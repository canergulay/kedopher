package server

import (
	"time"

	"github.com/canergulay/go-betternews-signaling/model/dto"
	"github.com/canergulay/go-betternews-signaling/model/enum"
)

var schedulePeriod = 15*time.Second

func (ws wsServer) OnlineUsersCountBroadcastProcessor() {
    ticker := time.NewTicker(schedulePeriod)
    defer ticker.Stop()

    for range ticker.C {
        onlineUsersCount := len(ws.userHub.Users)
        for _, user := range ws.userHub.Users {
            if user != nil && user.Conn != nil {
                ws.sendMessage(user.Conn, dto.Message{
                    Type: enum.ONLINE_USERS_COUNT,
                    Body: dto.OnlineUsersCount{
                        Count: onlineUsersCount,
                    },
                })
            }
        }
    }
}
