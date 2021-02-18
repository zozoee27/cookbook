package user

import (
	"github.com/zozoee27/cookbook/backend/entity"
)

type Database interface {
	Insert(e *entity.User) error
	FindOne(username string) (*entity.User, error)

	ClearAllEntries() error
}
