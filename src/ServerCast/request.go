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

type LoggedInRequest struct {
	AccessToken string `json:"token"`
}

type CreateDeviceRequest struct {
	LoggedInRequest
	DeviceName string `json:"device_name"`
}

func (j LoginRequest) ToJSON() ([]byte, error) {
	return json.Marshal(j)
}