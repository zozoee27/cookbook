package account

type Account struct {
    Username string `json:"username"`
    Email string `json:"email"`
    FirstName string `json:"firstname"`
    LastName string `json:"lastname"`
    Password string `json:"password"`
}
