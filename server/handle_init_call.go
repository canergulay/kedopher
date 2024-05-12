package server

import (
	"github.com/canergulay/go-betternews-signaling/model"
	"github.com/canergulay/go-betternews-signaling/model/enum"
	"github.com/sirupsen/logrus"
)

func (ws wsServer) handleInitCall(user *model.User)  {
	if user.Status != enum.UserIdle {
		return
	}

	user.SetStatus(enum.UserWaiting)

	ws.userHub.WaitingUsersQueue <- user.ID

	logrus.Infof("user added to waiting users queue %s",user.ID)
}