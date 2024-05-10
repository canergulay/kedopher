package server

import (
	"time"

	"github.com/canergulay/go-betternews-signaling/enum"
	"github.com/canergulay/go-betternews-signaling/model"
	"github.com/canergulay/go-betternews-signaling/model/dto"
)

func (ws wsServer) initSignalingForUser(user *model.User) {
	defer user.SetStatus(enum.UserWaiting)
	ok,peer := ws.userHub.GetRandomAvailableUser(user.ID)
	if !ok {
		time.Sleep(time.Second*3)
	}

	connection := model.NewConnection([]model.ID{user.ID,peer.ID})

	ws.connectionHub.AddNewConnection(&connection)

	ws.sendMessage(peer.Conn,dto.Message{
		Type: enum.OFFER,
		Body:dto.Offer{
			ConnectionID: string(connection.ID),
		},
	},)

	ws.sendMessage(user.Conn,dto.Message{
		Type: enum.OFFER,
		Body: dto.Offer{
			ConnectionID: string(connection.ID),
		},
	},)
}