package model

import (
	"fmt"
	"slices"

	"github.com/canergulay/go-betternews-signaling/enum"
	"github.com/google/uuid"
)

type Connection struct {
	ID ID
	Users []ID
	AcceptedUsers []ID
	Status enum.ConnectionStatus
}

func NewConnection(users[]ID) Connection{
	return Connection{
		ID: ID(uuid.New().String()),
		Users: users,
		Status: enum.ConnectionInitial,
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

func (c Connection) IsAllUsersAccepted() bool {
	return len(c.Users) == len(c.AcceptedUsers)
}