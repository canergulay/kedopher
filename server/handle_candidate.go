package server

import (
	"github.com/canergulay/go-betternews-signaling/model"
	"github.com/canergulay/go-betternews-signaling/model/dto"
	"github.com/canergulay/go-betternews-signaling/model/enum"
)

func (w wsServer) handleIceCandidate(message dto.Message,user *model.User){
	iceCandidate,ok := message.GetBodyAsIceCandidate()
	if !ok {
		// todo log
		return
	}

	connection := w.connectionHub.GetConnectionById(model.ID(iceCandidate.ConnectionID))
	if connection == nil {
		// todo log
		return
	}

	connection.AddUserToCandidateSentUsers(user.ID)
	user.SetStatus(enum.UserSentIceCandidate)


	// send iceCandidate to all users in the connection except the candidate owner
	for _, userID := range connection.AcceptedUsers {
		if userID == user.ID{
			continue
		}

		user := w.userHub.GetUserById(userID)
		if user != nil {
			// todo log
			continue
		}

		w.sendMessage(user.Conn,dto.Message{
			Type: enum.ICE_CANDIDATES,
			Body: iceCandidate,
		})

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