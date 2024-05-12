package server

import (
	"github.com/canergulay/go-betternews-signaling/model"
	"github.com/canergulay/go-betternews-signaling/model/dto"
	"github.com/canergulay/go-betternews-signaling/model/enum"
	"github.com/sirupsen/logrus"
)

func (w wsServer) handleAnswer(msg dto.Message,user *model.User){
	answer,ok := msg.GetBodyAsAnswer()
	if !ok {
		// todo log
		return
	}

	user.SetSdp(answer.Sdp)
	user.SetStatus(enum.UserWaiting)

	connection := w.connectionHub.GetConnectionById(answer.ConnectionID)
	if connection == nil {
		// todo log
		return
	}

	for _,userID := range connection.Users {
		if userID == user.ID {
			continue
		}

		user := w.userHub.GetUserById(userID)
		if user == nil {
		logrus.Warnf("unable to find user for handle answer back message %s",userID)
		continue
		}

		w.sendMessage(user.Conn,dto.Message{
			Type: enum.ANSWER,
			Body: answer,
		})

	}



	// if err := connection.AddUserToAcceptedUsers(user.ID); err!=nil {
	// 	// todo log
	// }

	// if connection.IsAllUsersAccepted(){
	// 	w.triggerIceCandidatesExchangeForConnectionUsers(connection)
	// }
}