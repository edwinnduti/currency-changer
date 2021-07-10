package router

// libraries to be used
import(
	"github.com/edwinnduti/currency-changer/lib"
	"github.com/gorilla/mux"
)

// register routers
func Router() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	router.HandleFunc("/",lib.PostHandler).Queries("to", "{TO}", "amount", "{AMOUNT}", "from", "{FROM}").Methods("POST","OPTIONS")

	return router
}
