package lib

// libraries imported
import(
	"net/http"
	"encoding/json"
	"github.com/edwinnduti/currency-changer/model"
)

// convertion table
const(
	KSH_NGN = 3.80
	KSH_GHS = 0.055
	GHS_NGN = 0.014
)

// post handler function 
func PostHandler(w http.ResponseWriter, r *http.Request) {
	// empty money struct
        var money model.Money

        // decode incoming values to user empty struct
	err := json.NewDecoder(r.Body).Decode(&money)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		response := model.Response{
			Code : 500,
			Message: "DataParsingError",
		}

		// give response to client
		json.NewEncoder(w).Encode(response)
	}

	// amount to post to
	var amount float32

	// use switch case for conversion of types
	switch money.To {
	case "KSH":
		if money.From == "KSH" {
			amount = money.Cash
		}
		if money.From == "NGN" {
			amount = money.Cash/KSH_NGN
		}
		if money.From == "GHS" {
			amount = money.Cash/KSH_GHS
		}

		// response header
		w.WriteHeader(http.StatusOK)
		currency := model.Currency{
			Type : "GHS",
			Amount : amount,
		}
		// give response to client
		json.NewEncoder(w).Encode(currency)

	case "NGN":
		if money.From == "NGN" {
			amount = money.Cash
		}
		if money.From == "KSH" {
			amount = money.Cash * KSH_NGN
		}
		if money.From == "GHS" {
			amount = money.Cash/GHS_NGN
		}

		// response header
		w.WriteHeader(http.StatusOK)
		currency := model.Currency{
			Type : "GHS",
			Amount : amount,
		}
		// give response to client
		json.NewEncoder(w).Encode(currency)

	case "GHS":
		if money.From == "GHS" {
			amount = money.Cash
		}
		if money.From == "NGN" {
			amount = money.Cash * GHS_NGN
		}
		if money.From == "KSH" {
			amount = money.Cash/KSH_GHS
		}

		// response header
		w.WriteHeader(http.StatusOK)
		currency := model.Currency{
                        Type : "GHS",
                        Amount : amount,
                }

                // give response to client
                json.NewEncoder(w).Encode(currency)

	default:
		w.WriteHeader(http.StatusOK)
		response := model.Response{
			Code : 500,
			Message: "AmountIsNotParsed",
		}

		// give response to client
		json.NewEncoder(w).Encode(response)
	}
}
