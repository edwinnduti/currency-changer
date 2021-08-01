package middleware

// libraries
import (
	"log"
	"net/http"
)

// logger function
func Logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				logger.Println(r.Method, r.URL.Path, r.RemoteAddr)
			}()
			next.ServeHTTP(w, r)
		})
	}
}
