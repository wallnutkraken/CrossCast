package main

import "encoding/json"

type GenericResponse struct {
	Success bool `json:"OK"`
	Message string `json:"detail,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

func ToJSON(object interface{}) ([]byte, error) {
	return json.Marshal(object)
}