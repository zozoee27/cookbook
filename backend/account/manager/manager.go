package accountManager

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/zozoee27/cookbook/backend/account"
	"github.com/zozoee27/cookbook/backend/database"

	"go.mongodb.org/mongo-driver/mongo"
)

type AccountManager struct {
	Database     *mongo.Database
	UserDatabase *database.UserDatabase
}

func Initialize(db *mongo.Database) *AccountManager {
	return &AccountManager{
		Database:     db,
		UserDatabase: database.InitializeUserDatabase(db.Collection("users"))}
}

func (m *AccountManager) RegisterAccount(w http.ResponseWriter, r *http.Request) {
	log.Print("Recieved /account/register command")

	var account account.Account

	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &account)

	if err != nil {
		log.Print("AccountManager::RegisterAccount: Cannot unmarshal json")
		return
	}

	err = m.UserDatabase.AddUserToCollection(account)
	if err != nil {
		log.Print("Could not add user to collection: ", err)
	}
}

func (m *AccountManager) ClearAllEntries() error {
	return m.UserDatabase.ClearAllEntries()
}
