package server

import (
	"fmt"

	"github.com/canergulay/go-betternews-signaling/model"
	"github.com/canergulay/go-betternews-signaling/model/dto"
)

func (w wsServer) handleOffer(message dto.Message,user *model.User){
	offer,ok := message.GetBodyAsOffer()
	if !ok {
		// todo log
		return
	}

	// todo logic

	fmt.Println(offer)
}