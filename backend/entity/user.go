package entity

import (
	"fmt"
)

type User struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
}

func (u *User) PrettyString() string {
	return fmt.Sprintf(`
		UserName:[%s]
		FirstName:[%s]
		LastName:[%s]
		Email:[%s]
		Password:[%s]]`, u.Username, u.FirstName, u.LastName, u.Email, u.Password)
}
