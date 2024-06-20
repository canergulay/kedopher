package server

import (
	"fmt"
	"net/http"
)

func (ws wsServer) HandleStatistics (w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("active users: %d \n", len(ws.userHub.Users))))
	w.Write([]byte(fmt.Sprintf("active connections: %d \n", len(ws.connectionHub.Connections))))
}