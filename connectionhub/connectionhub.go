package connectionhub

import (
	"sync"

	"github.com/canergulay/go-betternews-signaling/model"
)

type ConnectionHub struct {
	Connections map[model.ID]*model.Connection

	mutex sync.Mutex
}

func NewConnectionHub () ConnectionHub{
	return ConnectionHub{
		Connections: make(map[model.ID]*model.Connection),
		mutex: sync.Mutex{},
	}
}

func (u *ConnectionHub) AddNewConnection(user *model.Connection) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	u.Connections[model.ID(user.ID)] = user
}

func (u *ConnectionHub) GetConnectionById(id model.ID) (*model.Connection) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	return u.Connections[id]
}