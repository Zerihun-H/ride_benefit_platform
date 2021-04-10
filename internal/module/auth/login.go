package auth

import (
	"rideBenefit/internal/constant/model"

	"golang.org/x/crypto/bcrypt"
)

// Login validates the login credentials of a user and returns an access token
func (s *service) Login(login *model.LoginModel) (bool, string, error) {

	password := login.Password
	// Get user
	user, err := s.authPersist.GetUserByUsername(login.Username)
	if err != nil {
		return false, "", err
	}

	// Validate Password
	if valid := CheckPassword(user.Password, password); !valid {
		return false, "", err
	}

	// Create the user model to generate access token
	usr := model.User{
		ID:     user.ID,
		RoleID: user.RoleID,
	}

	// Issue access token
	accessToken, err := GenerateAccessToken(usr)
	if err != nil {
		return false, "", err
	}

	return true, accessToken, err
}

// CheckPassword checks if a given password is equal with its hash
func CheckPassword(hashedPassword, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return false
	}
	return true
}
