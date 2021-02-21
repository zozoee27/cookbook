package acceptanceTests

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jimlawless/whereami"

	"github.com/zozoee27/cookbook/backend/app"
	"github.com/zozoee27/cookbook/backend/entity"
	"github.com/zozoee27/cookbook/backend/testutil"
)

var application app.App

func TestMain(m *testing.M) {
	application.Initialize("CookbookDB_Test")

	code := m.Run()
	clearUserDB()
	os.Exit(code)
	application.StopApplication()
}

func clearUserDB() {
	err := application.UserService.ClearAllEntries()
	if err != nil {
		log.Fatal(err)
	}
}

func executeJsonRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()

	req.Header.Set("Content-Type", "application/json")
	application.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected Response [%d] Actual [%d]\n", expected, actual)
	}
}

func getUserInfoFromDatabase(t *testing.T, username string) *entity.User {

	result, err := application.UserService.FindUserFromCollection(username)

	if err != nil {
		t.Errorf("Could not find user %s", username)
	}
	return result
}

func TestRegisterUserWithValidInfo(t *testing.T) {
	clearUserDB()

	// Prepare account info
	accountInfo := &entity.User{
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

	testutil.CompareUserEntity(t, result, accountInfo, "Account information different", whereami.WhereAmI())
}

func TestRegisterUserWithMissingInfo(t *testing.T) {
	var tests = []struct {
		name         string
		accountInfo  entity.User
		expectedCode int
	}{
		{"Empty account", entity.User{}, http.StatusBadRequest},
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
