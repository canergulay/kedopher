package server

import (
	"github.com/canergulay/go-betternews-signaling/enum"
	"github.com/canergulay/go-betternews-signaling/model"
	"github.com/canergulay/go-betternews-signaling/model/dto"
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

	if err := connection.AddUserToAcceptedUsers(user.ID); err!=nil {
		// todo log
	}

	if connection.IsAllUsersAccepted(){
		w.triggerIceCandidatesExchangeForConnectionUsers(connection)
	}
}

func (w wsServer) triggerIceCandidatesExchangeForConnectionUsers(connection *model.Connection){
	defer connection.SetStatus(enum.ConnectionWaitingForIceExchange)
	
	for _,userID := range connection.AcceptedUsers {
		user := w.userHub.GetUserById(userID)
		if user == nil {
			// todo log
		}

		w.sendMessage(user.Conn,dto.Message{
			Type: enum.TRIGGER_ICE_CANDIDATES,
			Body: dto.TriggerIceCandidatesDTO{
				ConnectionID: string(connection.ID),
			},
		})

		user.SetStatus(enum.UserNotifiedForIce)
	}
}