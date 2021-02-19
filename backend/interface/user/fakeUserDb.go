package user

import (
	"github.com/zozoee27/cookbook/backend/entity"
)

type fakeUserDb struct {
	m map[string]*entity.User
	e error
}

func NewFakeUserDb() *fakeUserDb {
	var m = map[string]*entity.User{}
	return &fakeUserDb{
		m: m,
		e: nil,
	}
}

func NewFakeUserDbWithError(e error) *fakeUserDb {
	var m = map[string]*entity.User{}
	return &fakeUserDb{
		m: m,
		e: e,
	}
}

func (r *fakeUserDb) Insert(e *entity.User) error {
	r.m[e.Username] = e
	return r.e
}

func (r *fakeUserDb) FindOne(username string) (*entity.User, error) {
	result := r.m[username]
	if result == nil {
		return nil, r.e
	}

	return result, r.e
}

func (r *fakeUserDb) ClearAllEntries() error {
	r.m = map[string]*entity.User{}
	return r.e
}

func (r *fakeUserDb) Size() int {
	return len(r.m)
}
