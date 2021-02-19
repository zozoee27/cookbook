// +build acceptance

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

	"github.com/zozoee27/cookbook/backend/app"
	"github.com/zozoee27/cookbook/backend/entity"

	"go.mongodb.org/mongo-driver/bson"
)

var appplication app.App

func TestMain(m *testing.M) {
	appplication.Initialize("CookbookDB_Test")

	code := m.Run()
	clearUserDB()
	os.Exit(code)
	appplication.StopApplication()
}

func clearUserDB() {
	//	err := app.AccountManager.ClearAllEntries()
	if err != nil {
		log.Fatal(err)
	}
}

func executeJsonRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	req.Header.Set("Content-Type", "application/json")
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

func TestRegisterUserWithValidInfo(t *testing.T) {
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

	response := executeJsonRequest(request)

	checkResponseCode(t, http.StatusCreated, response.Code)

	result := getUserInfoFromDatabase(t, "ButtersButtons")

	if result != accountInfo {
		t.Errorf("Account information different")
	}
}

func TestRegisterUserWithMissingInfo(t *testing.T) {
	var tests = []struct {
		name         string
		accountInfo  account.Account
		expectedCode int
	}{
		{"Empty account", account.Account{}, http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			body := new(bytes.Buffer)
			json.NewEncoder(body).Encode(tt.accountInfo)

			request, _ := http.NewRequest("POST", "/account/register", body)
			response := executeJsonRequest(request)
			checkResponseCode(t, tt.expectedCode, response.Code)
		})
	}

}
