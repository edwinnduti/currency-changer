// the main function
package main

// libraries to use
import (
	"os"
	"log"
	"net/http"
	"github.com/edwinnduti/currency-changer/middleware"
	"github.com/edwinnduti/currency-changer/router"
	"github.com/gorilla/handlers"
)

// Main function
func main() {

	// register router
	r := router.Router()

	// get the port
	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "3000"
	}

	// set CORS
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"POST", "OPTIONS"})

	// establish a logging mechanism
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	// define server
	server := &http.Server{
		Handler: handlers.CORS(originsOk, methodsOk)(middleware.Logging(logger)(r)),
		Addr:    ":" + Port,
	}

	// print and run server
	log.Printf("Listening on PORT: %s", Port)
	server.ListenAndServe()
}
