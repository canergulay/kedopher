package model

import (
	"github.com/canergulay/go-betternews-signaling/model/enum"
	"github.com/gorilla/websocket"
)

type User struct {
	ID ID
	Conn *websocket.Conn
	Sdp string
	Status enum.UserStatus
	Connections []ID
}

func (u *User) SetSdp(sdp string){
	u.Sdp = sdp
}

func (u *User) SetStatus(status enum.UserStatus){
	u.Status = status
}

func (u *User) AddConnectionToUser(connectionID ID){
	
}