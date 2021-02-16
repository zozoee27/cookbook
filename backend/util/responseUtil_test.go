package util

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRespondWithError(t *testing.T) {
	str := "Error message"
	status := http.StatusBadRequest
	w := httptest.NewRecorder()
	body := map[string]string{"error": str}

	RespondWithError(w, status, str)

	expectedJson, _ := json.Marshal(body)
	checkResponse(t, w, status, expectedJson)
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
	checkResponse(t, w, status, expectedJson)
}

func TestResponseWithCode(t *testing.T) {
	status := http.StatusTooManyRequests

	w := httptest.NewRecorder()
	RespondWithCode(w, status)

	checkResponseCode(t, w, status)
}

func checkResponseCode(t *testing.T, w *httptest.ResponseRecorder, expectedCode int) {
	if w.Code != expectedCode {
		t.Errorf("Response Code Is wrong: Actual[%d] Expected[%d]", w.Code, expectedCode)
	}
}

func checkResponse(t *testing.T, w *httptest.ResponseRecorder, expectedCode int, expectedJson []byte) {
	if w.HeaderMap.Get("Content-Type") != "application/json" {
		t.Errorf("Content type of response is not JSON")
	}

	checkResponseCode(t, w, expectedCode)

	if w.Code != expectedCode {
		t.Errorf("Response Code Is wrong")
	}

	var actualBytes = w.Body.Bytes()

	if len(actualBytes) != len(expectedJson) {
		t.Errorf("Response length is different: Actual[%d] Expected[%d]", len(actualBytes), len(expectedJson))
		return
	}

	for i, v := range actualBytes {
		if v != expectedJson[i] {
			t.Errorf("Response body is wrong: \n  Actual[%s]\n  Expected[%s]\n]: ", actualBytes, expectedJson)
			return
		}
	}
}
