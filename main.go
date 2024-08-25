package main

import (
	"log"
	"net/http"

	"github.com/canergulay/go-betternews-signaling/connectionhub"
	"github.com/canergulay/go-betternews-signaling/server"
	"github.com/canergulay/go-betternews-signaling/userhub"
	"github.com/sirupsen/logrus"
)

func main() {
	userHub := userhub.NewUserHub()
	connectionHub := connectionhub.NewConnectionHub()

	wsServer := server.NewWsServer(&userHub, &connectionHub)
	go wsServer.OnlineUsersCountBroadcastProcessor()

	http.HandleFunc("/ws", wsServer.HandleWebsocketConnections)

	http.HandleFunc("/statistics", wsServer.HandleStatistics)

	logrus.Info("app is up")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
