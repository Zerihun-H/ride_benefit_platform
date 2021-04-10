package auth

import (
	"fmt"
	"os"
	"rideBenefit/internal/constant/model"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

// ValidateRefreshToken validates a refresh token.
func ValidateRefreshToken(login *model.LoginModel) (bool, string, error) {
	return false, "", nil
}

// GenerateAccessToken generates an access token from the given payload
func GenerateAccessToken(user model.User) (string, error) {

	// Create the token
	claims := model.AccessTokenClaims{
		UserID: uint64(user.ID),
		RoleID: strconv.Itoa(int(user.RoleID)),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.NewTime(float64(jwt.Now().Add(time.Minute * 5).Unix())),
			IssuedAt:  jwt.NewTime(float64(jwt.Now().Unix()))},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	return tokenString, err
}

// ValidateAccessToken checks if a token is valid or not
func ValidateAccessToken(tokenString string) (bool, jwt.MapClaims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		return false, jwt.MapClaims{}, err
	}

	if !token.Valid {
		return false, jwt.MapClaims{}, err

	}

	claims, isClaims := token.Claims.(jwt.MapClaims)
	if !isClaims {

		fmt.Println(err)
		return false, jwt.MapClaims{}, err

	}

	return true, claims, err
}
