package main

// libraries used
import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edwinnduti/currency-changer/lib"
	"github.com/edwinnduti/currency-changer/model"
	"github.com/gorilla/mux"
)

// test main function
func TestMain(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/", lib.PostHandler).Methods("POST", "OPTIONS")
	curr := []byte(`{"To":"KSH", "From": "GHS", "Cash": 100.00}`)
	request, err := http.NewRequest(
		"POST",
		"/",
		bytes.NewBuffer(curr),
	)

	// handle request error
	if err != nil {
		t.Error(err)
	}

	// specify the required result
	var want = model.Currency{
		Type:   "NGN",
		Amount: 100.00,
	}

	// empty struct to decode to
	var got model.Currency

	// decode request body to got struct
	err = json.NewDecoder(request.Body).Decode(&got)
	// handle test error
	if err != nil {
		t.Error(err)
	}

	// compare structs
	if got != want {
		t.Errorf("Got %v , wanted %v", got, want)
	}

	// register writer
	w := httptest.NewRecorder()
	r.ServeHTTP(w, request)
	if w.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d ", w.Code)
	}

}
