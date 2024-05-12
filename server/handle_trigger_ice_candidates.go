package server

import (
	"github.com/canergulay/go-betternews-signaling/model"
	"github.com/canergulay/go-betternews-signaling/model/dto"
	"github.com/canergulay/go-betternews-signaling/model/enum"
	"github.com/sirupsen/logrus"
)

func (ws wsServer) handleTriggerIceCandidates(message dto.Message,user *model.User)  {
	triggerIceCandidates,ok := message.GetBodyAsTriggerIceCandidates()
	if !ok {
		// todo log
		return
	}

	logrus.Infof("triggering ice candidate exchange for the connection : %s",triggerIceCandidates.ConnectionID)

	connection := ws.connectionHub.GetConnectionById(model.ID(triggerIceCandidates.ConnectionID))
	if connection == nil {
		// todo log
		return
	}

	ws.triggerIceCandidatesExchangeForConnectionUsers(connection)
}

func (w wsServer) triggerIceCandidatesExchangeForConnectionUsers(connection *model.Connection){
	defer connection.SetStatus(enum.ConnectionWaitingForIceExchange)
	
	for _,userID := range connection.Users {
		user := w.userHub.GetUserById(userID)
		if user == nil {
			// todo log
			continue
		}

		w.sendMessage(user.Conn,dto.Message{
			Type: enum.TRIGGER_ICE_CANDIDATE,
			Body: dto.TriggerIceCandidates{
				ConnectionID: connection.ID,
			},
		})

		logrus.Infof("user notified for ice exchange, userID: %s",userID)

		user.SetStatus(enum.UserNotifiedForIce)
	}
}