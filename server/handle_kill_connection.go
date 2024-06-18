package server

import (
	"fmt"

	"github.com/canergulay/go-betternews-signaling/model"
	"github.com/canergulay/go-betternews-signaling/model/dto"
	"github.com/canergulay/go-betternews-signaling/model/enum"
	"github.com/sirupsen/logrus"
)

func (ws wsServer) handleKillConnection(message dto.Message,user *model.User) {
	fmt.Println("kill connection geldi !")
	killConnectionDTO,ok := message.GetBodyAsKillConnection()
	if !ok {
		logrus.Errorf("unable to get body as killConnectionDTO %v", message)
		return
	}

	connection := ws.connectionHub.GetConnectionById(killConnectionDTO.ConnectionID)
	if connection == nil {
		logrus.Errorf("unable to get connection with id %s", killConnectionDTO.ConnectionID)
		return 
	}


	for _,userID := range connection.Users{
		if userID == user.ID {
			continue
		}

		user := ws.userHub.GetUserById(userID)
		if user == nil {
			logrus.Warnf("unable to get user with id %s", userID)
			return
		}

		ws.sendMessage(user.Conn,dto.Message{
			Type: enum.KILL_CONNECTION,
			Body: dto.KillConnection{
				ConnectionID: killConnectionDTO.ConnectionID,
			},
		})

	}
}