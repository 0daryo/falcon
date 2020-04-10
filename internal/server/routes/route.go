package route

import (
	"fmt"
	"net/http"

	"github.com/apex/log"
	"github.com/go-chi/chi"
)

// Routing ... define routing
func Routing(r chi.Router) {

	// recover
	r.Use(recoverAndServe)

	// access control
	// Ping
	r.Get("/", ping)

	http.Handle("/", r)

}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func recoverAndServe(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Error(fmt.Sprintf("%w", err))
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
