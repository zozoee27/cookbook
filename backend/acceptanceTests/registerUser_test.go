package acceptanceTests

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/zozoee27/cookbook/backend/account"
	"github.com/zozoee27/cookbook/backend/application"

	"go.mongodb.org/mongo-driver/bson"
)

var app application.App

func TestMain(m *testing.M) {
	app.Initialize("CookbookDB_Test")

	code := m.Run()
	clearUserDB()
	os.Exit(code)
	app.StopApplication()
}

func clearUserDB() {
	err := app.AccountManager.ClearAllEntries()
	if err != nil {
		log.Fatal(err)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected Response [%d] Actual [%d]\n", expected, actual)
	}
}

func getUserInfoFromDatabase(t *testing.T, username string) account.Account {
	var result account.Account

	filter := bson.D{{"username", username}}

	err := app.AccountManager.UserDatabase.UserCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		t.Errorf("Could not find user %s", username)
	}
	return result
}

func TestRegisterUser(t *testing.T) {
	clearUserDB()

	// Prepare account info
	accountInfo := account.Account{
		Username:  "ButtersButtons",
		Email:     "butters@buttons.com",
		FirstName: "Butters",
		LastName:  "Buttons",
		Password:  "Password123",
	}

	body := new(bytes.Buffer)
	json.NewEncoder(body).Encode(accountInfo)

	request, _ := http.NewRequest("POST", "/account/register", body)

	response := executeRequest(request)

	checkResponseCode(t, http.StatusOK, response.Code)

	result := getUserInfoFromDatabase(t, "ButtersButtons")

	if result != accountInfo {
		t.Errorf("Account information different")
	}
}
