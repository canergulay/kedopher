package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/canergulay/go-betternews-signaling/connectionhub"
	"github.com/canergulay/go-betternews-signaling/model"
	"github.com/canergulay/go-betternews-signaling/model/dto"
	"github.com/canergulay/go-betternews-signaling/model/enum"
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
	ws := wsServer{
		userHub: userHub,
		connectionHub: connectionhub,
	}

	go ws.waitingUsersProcessor()

	return ws 
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

	logrus.Infof("connection set for user %s",user.ID)
	
	ws.userHub.AddNewUser(user)

	go ws.sendMessage(user.Conn, dto.Message{
		Type: enum.ONLINE_USERS_COUNT,
		Body: dto.OnlineUsersCount{
			Count: len(ws.userHub.Users),
		},
	})

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			logrus.Infof("unable to receive msg for user %s, error: %v",user.ID,err)
			ws.userHub.DeleteUserByID(user.ID)
			break
		}

		var message dto.Message
		fmt.Println(string(msg))
		err = message.UnmarshalJSON(msg)
		if err != nil {
			logrus.WithField("message",msg).WithError(err).Warnf("unable to parse message for user %s",user.ID)
			continue
		}

		logrus.WithField("message",string(message.Type)).Infof("message received for user %s",user.ID)

		switch message.Type {
		case enum.INIT_CALL:
			ws.handleInitCall(user)
		case enum.OFFER:
			ws.handleOffer(message,user)
		case enum.ANSWER:
			ws.handleAnswer(message,user)
		case enum.TRIGGER_ICE_CANDIDATE:
			ws.handleTriggerIceCandidates(message,user)
		case enum.ICE_CANDIDATES:
			ws.handleIceCandidate(message,user)
		case enum.USER_STATUS:
			ws.handleUserStatus(message,user)
		case enum.KILL_CONNECTION:
			ws.handleKillConnection(message,user)
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
