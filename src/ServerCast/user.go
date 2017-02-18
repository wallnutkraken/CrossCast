package main

import (
	"golang.org/x/crypto/bcrypt"
	"encoding/hex"
)

type User struct {
	Username string
	Password string
	Devices *Devices
	Subscriptions []PodcastFeed
}
// LoginValid verifies if the given password is valid for this user
func (u User) LoginValid(password string) (bool, error) {
	passwordBytes, err := hex.DecodeString(u.Password)
	if err != nil {
		return false, err
	}

	err = bcrypt.CompareHashAndPassword(passwordBytes, []byte(password))
	return err == nil, err
}

// ChangePassword changes the password for this user and hashes it with the current salt
func (u *User) ChangePassword(password string) error {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = hex.EncodeToString(result)
	return nil
}

type Users []*User

func FindUser(username string) (*User, error) {
	for _, user := range users {
		if user.Username == username {
			return user, nil
		}
	}
	return nil, ErrInvalidLogin
}

func Register(user User) error {
	if _, err := FindUser(user.Username); err == nil {
		return ErrUserAlreadyExists
	}

	/* Call changepassword to hash the password */
	user.ChangePassword(user.Password)

	users = append(users, &user)
	return nil
}