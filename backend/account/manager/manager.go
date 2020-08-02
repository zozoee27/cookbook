package accountManager

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/zozoee27/cookbook/backend/account"
)

func RegisterAccount (w http.ResponseWriter, r *http.Request) {

    log.Print("Recieved /account/register command")

    var account account.Account;

    body, _:= ioutil.ReadAll(r.Body)
    err := json.Unmarshal(body, &account)

    if err != nil {
        log.Print("AccountManager::RegisterAccount: Cannot unmarshal json")
        return
    }

    Username := "username: " + account.Username 
    firstname := "firstName: " + account.FirstName 
    lastname := "lastName: " + account.LastName 
    email := "email: " + account.Email 
    password := "password: " + account.Password 

    log.Print(Username)
    log.Print(firstname)
    log.Print(lastname)
    log.Print(email)
    log.Print(password)
}
