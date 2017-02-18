package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"io"
)

func bodyToObject(body io.ReadCloser, object interface{}) error {
	bodyRead, err := ioutil.ReadAll(body)
	if err != nil {
		return err
	}
	defer body.Close()
	return json.Unmarshal(bodyRead, object)
}

// Login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	req := LoginRequest{}
	err := bodyToObject(r.Body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	user, err := FindUser(req.Username)
	if err != nil {
		response, _ := ToJSON(GenericResponse{false, err.Error(), nil})
		w.Write(response)
		w.WriteHeader(http.StatusForbidden)
		return
	}
	valid, err := user.LoginValid(req.Password)
	if err != nil || !valid {
		response, _ := ToJSON(GenericResponse{false, "Invalid password.", nil})
		w.Write(response)
		w.WriteHeader(http.StatusForbidden)
		return
	}

	w.Write([]byte(tokens.New(user.Username).Token))
	w.WriteHeader(http.StatusOK)
	return
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	req := RegisterRequest{}
	err := bodyToObject(r.Body, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = FindUser(req.Username)
	if err == nil {
		response, _ := ToJSON(GenericResponse{false, "User already exists", nil})
		w.Write(response)
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}
	err = Register(User{
		req.Username,
		req.Password,
		Devices{},
		make([]PodcastFeed, 0)})
	if err != nil {
		response, _ := ToJSON(GenericResponse{false, err.Error() + err.Error(), nil})
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
		return
	}
	response, _ := ToJSON(GenericResponse{true, "", nil})
	w.Write(response)
	w.WriteHeader(http.StatusOK)
}

func NewDeviceHandler(w http.ResponseWriter, r *http.Request) {
	lir := CreateDeviceRequest{}
	err := bodyToObject(r.Body, &lir)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := tokens.FindUser(lir.AccessToken)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		response, _ := ToJSON(GenericResponse{false, err.Error(), nil})
		w.Write(response)
		return
	}
	dev := user.Devices.Add(lir.DeviceName)
	w.WriteHeader(http.StatusOK)
	response, _ := ToJSON(dev)
	w.Write(response)
}