package server

import (
	"fmt"

	"github.com/canergulay/go-betternews-signaling/model"
	"github.com/canergulay/go-betternews-signaling/model/dto"
	"github.com/canergulay/go-betternews-signaling/model/enum"
	"github.com/sirupsen/logrus"
)

func (w wsServer) handleIceCandidate(message dto.Message,user *model.User){
	fmt.Println("bura tetik")
	iceCandidate,ok := message.GetBodyAsIceCandidate()
	if !ok {
		logrus.Errorf("unable to get body as ice candidate for user %s",user.ID)
		return
	}

	connection := w.connectionHub.GetConnectionById(model.ID(iceCandidate.ConnectionID))
	if connection == nil {
		logrus.Errorf("unable to get connection for user %s",user.ID)
		return
	}

	connection.AddUserToCandidateSentUsers(user.ID)
	user.SetStatus(enum.UserSentIceCandidate)


	// send iceCandidate to all users in the connection except the candidate owner
	for _, userID := range connection.Users {
		if userID == user.ID{
			logrus.Info("userID is same, skipping")
			continue
		}

		user := w.userHub.GetUserById(userID)
		if user == nil {
			logrus.Warnf("unable to find userID %s",userID)
			continue
		}
	
		w.sendMessage(user.Conn,dto.Message{
			Type: enum.ICE_CANDIDATES,
			Body: iceCandidate,
		})
		logrus.Infof("ice candidates are sent to user %s from user %s",userID,user.ID)
	}

	w.checkForCandidateSentUsersAndTriggerCall(connection)
}

func (w wsServer) checkForCandidateSentUsersAndTriggerCall(connection *model.Connection) {
	if !connection.IsAllUsersSentIceCandidates() {
		return
	}

	for _,userID := range connection.CandidateSentUsers {
		user := w.userHub.GetUserById(userID)
		if user == nil {
			// todo log
			continue
		}

		w.sendMessage(user.Conn,dto.Message{
			Type: enum.START_CALL,
			Body: dto.StartCall{
				ConnectionID: connection.ID,
			},
		})
	} 
}