package middleware

import (
	"fmt"
	"net/http"
	"os"
	"rideBenefit/internal/constant/model"
	"strconv"
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

		if len(splitAuthHeader) == 2 {
			if authHeader == "" {
				http.Error(w, "Invalid access  token", http.StatusBadRequest)
				return
			}
		}

		accessToken := strings.TrimSpace(splitAuthHeader[1])

		token, err := jwt.ParseWithClaims(accessToken, &model.AccessTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET")), nil
		})

		if err != nil {
			http.Error(w, "Invalid access token", http.StatusBadRequest)

			return
		}

		if !token.Valid {
			http.Error(w, "Invalid access token", http.StatusUnauthorized)
			return
		}

		claims, isClaims := token.Claims.(*model.AccessTokenClaims)
		if !isClaims {

			fmt.Println(err)
			http.Error(w, "Invalid access token", http.StatusUnauthorized)
			return

		}

		err = claims.Valid(jwt.DefaultValidationHelper)
		if err != nil {
			http.Error(w, "Invalid access token", http.StatusUnauthorized)
			return

		}

		r.Header.Set("rle", claims.RoleID)
		r.Header.Set("id", strconv.Itoa(int(claims.UserID)))

		next(w, r, ps)
	})
}


	