package main

import "encoding/json"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (j LoginRequest) ToJSON() ([]byte, error) {
	return json.Marshal(j);
}