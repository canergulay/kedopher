package server

import (
	"github.com/canergulay/go-betternews-signaling/enum"
	"github.com/canergulay/go-betternews-signaling/model"
	"github.com/canergulay/go-betternews-signaling/model/dto"
	"github.com/sirupsen/logrus"
)

func (ws wsServer) initSignalingForUsers(users []model.ID) {
	connection := model.NewConnection(users)

	ws.connectionHub.AddNewConnection(&connection)

	for _, userID := range users {
		user := ws.userHub.GetUserById(userID)
		if user == nil {
			logrus.Warn("unable to find user for signal initializing")
			continue
		}

		ws.sendMessage(user.Conn,dto.Message{
			Type: enum.OFFER_START,
			Body:dto.Offer{
				ConnectionID: string(connection.ID),
			},
		},)
	}
}