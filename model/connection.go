package model

import (
	"fmt"
	"slices"
	"time"

	"github.com/canergulay/go-betternews-signaling/model/enum"
	"github.com/google/uuid"
)

type Connection struct {
	ID ID
	Users []ID
	AcceptedUsers []ID
	CandidateSentUsers []ID
	Status enum.ConnectionStatus
	CreatedAt time.Time
}

func NewConnection(users[]ID) Connection{
	return Connection{
		ID: ID(uuid.New().String()),
		Users: users,
		AcceptedUsers: make([]ID, 0, len(users)),
		CandidateSentUsers: make([]ID, 0, len(users)),
		Status: enum.ConnectionInitial,
		CreatedAt: time.Now(),
	}
}

func (c *Connection) SetStatus(status enum.ConnectionStatus){
	c.Status = status
}

func (c *Connection) AddUserToUsers(userID ID){
	c.Users = append(c.Users, userID)
}

func (c *Connection) AddUserToAcceptedUsers(userID ID) error {
	if !slices.Contains(c.Users,userID){
		return fmt.Errorf("user does not exist within the connection, userID: %s",userID)
	}

	c.AcceptedUsers = append(c.AcceptedUsers, userID)

	return nil
}

func (c *Connection) AddUserToCandidateSentUsers(userID ID) error {
	if !slices.Contains(c.Users,userID){
		return fmt.Errorf("user does not exist within the connection, userID: %s",userID)
	}

	c.CandidateSentUsers = append(c.CandidateSentUsers, userID)

	return nil
}


func (c Connection) IsAllUsersAccepted() bool {
	return len(c.Users) == len(c.AcceptedUsers)
}

func (c Connection) IsAllUsersSentIceCandidates() bool {
	return len(c.Users) == len(c.CandidateSentUsers)
}