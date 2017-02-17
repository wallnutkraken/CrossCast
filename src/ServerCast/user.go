package ServerCast

import "golang.org/x/crypto/bcrypt"

type User struct {
	Username string
	Password string
	Salt string
	Devices []Device
	Subscriptions []PodcastFeed
}

// LoginValid verifies if the given password is valid for this user
func (u User) LoginValid(password string) (bool, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password + u.Salt), bcrypt.DefaultCost)

	return string(result) == u.Password, err
}

// ChangePassword changes the password for this user and hashes it with the current salt
func (u *User) ChangePassword(password string) error {
	result, err := bcrypt.GenerateFromPassword([]byte(password + u.Salt), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(result)
	return nil
}
