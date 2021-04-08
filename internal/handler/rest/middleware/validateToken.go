package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/julienschmidt/httprouter"
)

func ValidateAccessToken(next httprouter.Handle) httprouter.Handle {
	return httprouter.Handle(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Invalid Authorization header", http.StatusBadRequest)
			return
		}

		splitAuthHeader := strings.Split(authHeader, "Bearer ")

		if len(splitAuthHeader) < 2 {
			if authHeader == "" {
				http.Error(w, "Invalid access  token", http.StatusBadRequest)
				return
			}
		}

		accessToken := splitAuthHeader[1]
		log.Println("In token validator")
		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// secret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return os.Getenv("SECRET"), nil
		})

		if err != nil {
			http.Error(w, "Invalid access token", http.StatusBadRequest)

			return
		}

		if !token.Valid {
			http.Error(w, "Invalid access token", http.StatusUnauthorized)

			return
		}

		next(w, r, ps)
	})
}
