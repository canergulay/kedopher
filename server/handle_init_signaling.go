package server

import (
	"math/rand"

	"github.com/canergulay/go-betternews-signaling/model"
	"github.com/canergulay/go-betternews-signaling/model/dto"
	"github.com/canergulay/go-betternews-signaling/model/enum"
	"github.com/sirupsen/logrus"
)

func (ws wsServer) initSignalingForUsers(users []model.ID) {
	logrus.Infof("initializing signaling for users %v",users)

	connection := model.NewConnection(users)

	randomUserToStartSignaling := users[rand.Intn(len(users))]

	logrus.Infof("random user selected for the signaling initialization, userID: %s",randomUserToStartSignaling)

	ws.connectionHub.AddNewConnection(&connection)
	user := ws.userHub.GetUserById(randomUserToStartSignaling)
	if user == nil {
		logrus.Warn("unable to find user for signal initializing")
		
	}
	
	ws.sendMessage(user.Conn,dto.Message{
		Type: enum.OFFER_START,
		Body:dto.Offer{
			ConnectionID: connection.ID,
		},
	},)
}