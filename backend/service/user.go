package service

import (
	"log"

	"github.com/zozoee27/cookbook/backend/entity"
	"github.com/zozoee27/cookbook/backend/interface/user"
)

type User struct {
	db user.Database
}

func CreateUserService(d user.Database) *User {
	return &User{
		db: d,
	}
}

func (u *User) CreateUser(user *entity.User) error {
	err := u.db.Insert(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) FindUserFromCollection(username string) (*entity.User, error) {
	result, err := u.db.FindOne(username)

	if err != nil {
		log.Print("Find user error: ", err)
	}

	return result, err
}

func (u *User) ClearAllEntries() error {
	return u.db.ClearAllEntries()
}
