package main

import (
	"time"
	"github.com/satori/go.uuid"
)

type AccessToken struct {
	Owner string
	IssueDate int64
	Token string
}

func NewToken(owner string) AccessToken {
	return AccessToken{owner, time.Now().Unix(), uuid.NewV4().String()}
}

func (at AccessToken) Valid() bool {
	const twenty_four_hours int64 = 86400
	return time.Now().Unix() - at.IssueDate < twenty_four_hours
}

type TokenCollection []AccessToken

func (tc TokenCollection) New(username string) AccessToken {
	newToken := NewToken(username)
	for index, token := range tc {
		if token.Owner == username {
			tc[index] = newToken
		}
	}
	return newToken
}