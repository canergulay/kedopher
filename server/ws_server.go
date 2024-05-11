package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/canergulay/go-betternews-signaling/connectionhub"
	"github.com/canergulay/go-betternews-signaling/enum"
	"github.com/canergulay/go-betternews-signaling/model"
	"github.com/canergulay/go-betternews-signaling/model/dto"
	"github.com/canergulay/go-betternews-signaling/userhub"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type WsServer interface {
	HandleWebsocketConnections(w http.ResponseWriter, r *http.Request)
}

type wsServer struct {
	userHub *userhub.UserHub
	connectionHub *connectionhub.ConnectionHub
}

func NewWsServer(userHub *userhub.UserHub, connectionhub *connectionhub.ConnectionHub) wsServer {
	return wsServer{
		userHub: userHub,
		connectionHub: connectionhub,
	}
}


var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}


func (ws wsServer) HandleWebsocketConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logrus.Errorf("unable to upgrade connection to websocket, error: %s",err)
		return
	}
	defer conn.Close()

	user := &model.User{
		ID:model.ID(uuid.New().String()),
		Conn: conn,
		Status: enum.UserIdle,
	}

	logrus.Infof("connection set for user %d",user.ID)
	
	ws.userHub.AddNewUser(user)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			logrus.Infof("unable to receive msg for user %d, error: %v",user.ID,err)
			ws.userHub.DeleteUserByID(user.ID)
			break
		}

		logrus.WithField("message",msg).Infof("message received for user %s",user.ID)

		var message dto.Message
		err = json.Unmarshal(msg,&message)
		if err != nil {

			logrus.WithField("message",msg).Warnf("unable to parse message for user %s",user.ID)
			continue
		}
		fmt.Println("able to marshall",message)
		switch message.Type {
		case enum.INIT_SIGNALING:
			ws.initSignalingForUser(user)
		case enum.OFFER:
			ws.handleOffer(message,user)
		case enum.ANSWER:
			ws.handleAnswer(message,user)
		case enum.SEND_ICE_CANDIDATE:
			ws.handleIceCandidate(message,user)
		}
	}
}


func (ws wsServer) sendMessage(conn *websocket.Conn, msg interface{}) {
	if conn == nil {
		return
	}

	message, err := json.Marshal(msg)
	if err != nil {
		log.Printf("Failed to marshal message: %v", err)
		return
	}
	if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
		log.Printf("Failed to send message: %v", err)
	}
}
