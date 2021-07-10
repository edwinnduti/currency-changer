package main

// libraries used
import(
	"github.com/edwinnduti/currency-changer/lib"
	"github.com/gorilla/mux"
	"net/http/httptest"
	"net/http"
	"strings"
	"testing"
)

// test main function
func TestMain(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/",lib.PostHandler).Methods("POST", "OPTIONS")
	curr := `{"To":"KSH", "From": "GHS", "Cash": 100.00}`
	request, err := http.NewRequest(
		"POST",
		"/",
		strings.NewReader(curr),
	)

	// handle test error
	if err != nil{
		t.Error(err)
	}

	// register writer
	w := httptest.NewRecorder()
	r.ServeHTTP(w,request)
	if w.Code != 200 {
		t.Errorf("HTTP Status expected: 200, got: %d ",w.Code)
	}

}
