package main

import "encoding/json"

type GenericResponse struct {
	Success bool `json:"OK"`
	Message string `json:"detail,omitempty"`
	Value interface{} `json:"value,omitempty"`
}

type LoginResponse struct {
	AccessToken string
}

func ToJSON(object interface{}) ([]byte, error) {
	return json.Marshal(object);
}