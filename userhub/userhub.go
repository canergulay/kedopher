package userhub

import (
	"sync"

	"github.com/canergulay/go-betternews-signaling/enum"
	"github.com/canergulay/go-betternews-signaling/model"
)

type UserHub struct {
	Users map[model.ID]*model.User 
	WaitingUsersQueue chan model.ID
	mutex sync.Mutex
}

func NewUserHub () UserHub{
	return UserHub{
		Users: make(map[model.ID]*model.User),
		WaitingUsersQueue: make(chan model.ID, 1000),
		mutex: sync.Mutex{},
	}
}

func (u *UserHub) AddNewUser(user *model.User) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	u.Users[model.ID(user.ID)] = user
}

func (u *UserHub) DeleteUserByID(userID model.ID){
	delete(u.Users,userID)
}

func (u *UserHub) GetUserById(id model.ID) (*model.User) {
	u.mutex.Lock()
	defer u.mutex.Unlock()
	return u.Users[id]
}

func (u *UserHub) GetRandomAvailableUser(userID model.ID) (bool, *model.User){
	u.mutex.Lock()
	defer u.mutex.Unlock()

	for _,user := range u.Users{
		if userID != user.ID && (user.Status == enum.UserIdle || user.Status == enum.UserWaiting){
			return true,user
		}
	}

	return false,nil
}