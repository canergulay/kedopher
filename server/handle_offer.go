package server

import (
	"github.com/canergulay/go-betternews-signaling/model"
	"github.com/canergulay/go-betternews-signaling/model/dto"
	"github.com/canergulay/go-betternews-signaling/model/enum"
	"github.com/sirupsen/logrus"
)

func (w wsServer) handleOffer(message dto.Message,user *model.User){
	logrus.Infof("offer for user: %s",user.ID)
	offer,ok := message.GetBodyAsOffer()
	if !ok {
		logrus.Errorf("unable to get body as offer for user %s, body: %v",user.ID,message.Body)
		return
	}

	logrus.Infof("offer for user is ok %s",user.ID)

	connection := w.connectionHub.GetConnectionById(offer.ConnectionID)
	if connection == nil {
		logrus.Warnf("unable to find connection %s for user %s",offer.ConnectionID,user.ID)
		return
	}

	for _,userID := range connection.Users {
		if userID == user.ID{
			continue
		}

		peer := w.userHub.GetUserById(userID)
		if peer == nil {
			logrus.Infof("unable to find peer for user %s, peer: ",user.ID,userID)
			continue
		}

		w.sendMessage(peer.Conn,dto.Message{
			Type: enum.OFFER,
			Body: offer,
		})
	}
}