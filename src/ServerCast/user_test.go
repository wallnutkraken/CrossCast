package main

import "testing"

func init() {
	setup()
}

func TestCanRegister(t *testing.T) {
	userToRegister := User{"uname",
		"pass",
		Devices{},
		make([] PodcastFeed, 0)}
	err := Register(userToRegister)
	if err != nil {
		t.Fatal(err)
	}

	user, err := FindUser(userToRegister.Username)
	if err != nil {
		t.Fatal(err)
	}
	if userToRegister.Password == user.Password {
		t.Fatal("Password was not hashed")
	}
}

func TestCanRegisterDevice(t *testing.T) {
	u := User{"user",
		"pass",
		Devices{},
		make([] PodcastFeed, 0)}
	u.Devices.Add("TestDevice")

	if len(u.Devices.List) < 1 {
		t.Fatal("Device was not added")
	}
}

func TestUserIsReference(t *testing.T) {
	userToRegister := User{"uname2",
			       "pass",
			       Devices{},
			       make([] PodcastFeed, 0)}
	err := Register(userToRegister)
	if err != nil {
		t.Fatal(err)
	}

	user, err := FindUser(userToRegister.Username)
	if err != nil {
		t.Fatal(err)
	}
	user.Password = "123"
	anotherUser, err := FindUser(userToRegister.Username)
	if anotherUser.Password != user.Password {
		t.Fatal("Users are not handled by reference")
	}
}