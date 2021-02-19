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
		UserName:[%s] \n
		FirstName:[%s] \n
		LastName:[%s] \n
		Email:[%s] \n
		Password:[%s]]\n`, u.Username, u.FirstName, u.LastName, u.Email, u.Password)
}
