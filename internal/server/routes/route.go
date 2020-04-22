package route

import (
	"fmt"
	"net/http"
	"time"

	"github.com/apex/log"
	"github.com/go-chi/chi"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

var GetTokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	// headerのセット
	token := jwt.New(jwt.SigningMethodHS256)

	// claimsのセット
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = "54546557354"
	claims["name"] = "taro"
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// 電子署名
	tokenString, _ := token.SignedString([]byte("ARHzbmWBIvbOMxt7OwadxrryTS5lbFwcBQBuXypOulSdD71uyw=="))

	// JWTを返却
	w.Write([]byte(tokenString))
})

// JwtMiddleware check token
var JwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte("2pW1ARHzbmWBIvbOMxt7OwadxrryTS5lbFwcBQBuXypOulSdD71uyw=="), nil
	},
	SigningMethod: jwt.SigningMethodHS256,
})

// Routing ... define routing
func Routing(r chi.Router) {

	// recover
	r.Use(recoverAndServe)

	// access control
	// Ping
	r.Get("/", ping)
	if h, ok := JwtMiddleware.Handler(private).(http.HandlerFunc); !ok {
		panic("not assignable handler func")
	} else {
		r.Get("/private", h)
	}
	r.Get("/login", GetTokenHandler)

	http.Handle("/", r)

}

func ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

var private = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("jwt success"))
})

func recoverAndServe(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Error(fmt.Sprintf("%v", err))
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
