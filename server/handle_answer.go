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
		logrus.Errorf("unable to get body as answer for user %s",user.ID)
		return
	}

	user.SetSdp(answer.Sdp)
	user.SetStatus(enum.UserWaiting)

	connection := w.connectionHub.GetConnectionById(answer.ConnectionID)
	if connection == nil {
		logrus.Errorf("unable to get connection %s",answer.ConnectionID)
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
}