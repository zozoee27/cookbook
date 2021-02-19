// +build unit

package util

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jimlawless/whereami"

	"github.com/zozoee27/cookbook/backend/testutil"
)

func TestRespondWithError(t *testing.T) {
	str := "Error message"
	status := http.StatusBadRequest
	w := httptest.NewRecorder()
	body := map[string]string{"error": str}

	RespondWithError(w, status, str)

	expectedJson, _ := json.Marshal(body)
	checkResponse(t, w, status, expectedJson, whereami.WhereAmI())
}

func TestRespondWithJson(t *testing.T) {
	type responseJson struct {
		Message string
		Number  int
	}

	status := http.StatusAlreadyReported
	body := &responseJson{
		Message: "TEST1",
		Number:  12345,
	}

	w := httptest.NewRecorder()
	RespondWithJson(w, status, body)

	expectedJson, _ := json.Marshal(body)
	checkResponse(t, w, status, expectedJson, whereami.WhereAmI())
}

func TestResponseWithCode(t *testing.T) {
	status := http.StatusTooManyRequests

	w := httptest.NewRecorder()
	RespondWithCode(w, status)

	checkResponseCode(t, w, status, whereami.WhereAmI())
}

func checkResponseCode(t *testing.T, w *httptest.ResponseRecorder, expectedCode int, where string) {
	testutil.CompareInt(t, w.Code, expectedCode, "Response Code Is Wrong", where)
}

func checkResponse(t *testing.T, w *httptest.ResponseRecorder, expectedCode int, expectedJson []byte, where string) {

	checkResponseCode(t, w, expectedCode, where)

	testutil.CompareString(t,
		w.HeaderMap.Get("Content-Type"),
		"application/json",
		"Content type of response is not JSON",
		where)

	testutil.CompareByteArray(t, w.Body.Bytes(), expectedJson, "Response body is different", where)
}
