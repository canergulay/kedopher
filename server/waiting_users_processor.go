package server

import (
	"github.com/canergulay/go-betternews-signaling/model"
	"github.com/sirupsen/logrus"
)

const userCountForProcess = 2

func (ws wsServer) waitingUsersProcessor () {

	users := make([]model.ID,0,userCountForProcess)
	for userID := range ws.userHub.WaitingUsersQueue {
		logrus.Infof("user is dequeued %s",userID)
		users = append(users, userID)

		if len(users) == userCountForProcess{
			logrus.Infof("user count is at limit, cleaning process")
			ws.initSignalingForUsers(users)
			users = make([]model.ID,0,userCountForProcess)
		}
	}
		
}