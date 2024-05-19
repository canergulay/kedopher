package server

import (
	"github.com/canergulay/go-betternews-signaling/model"
	"github.com/canergulay/go-betternews-signaling/model/dto"
	"github.com/sirupsen/logrus"
)

func (w wsServer) handleUserStatus(message dto.Message,user *model.User){
	body,ok := message.GetBoyAsUserStatus()
	if !ok {
		logrus.Errorf("unable to get body as user status %v",message)
		return
	}

	user.SetStatus(body.Status)
}
