package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"io"
	"github.com/gorilla/mux"
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
	w.Header().Add("Content-Type", "application/json")
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
	token := tokens.New(user.Username)

	response, _ := ToJSON(GenericResponse{
		true,
		"User is logged in, value contains access token",
		LoginResponse{token.Token}})
	w.Write(response)
	w.WriteHeader(http.StatusOK)
	return
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
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
		&Devices{},
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
	w.Header().Add("Content-Type", "application/json")
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
	response, _ := ToJSON(GenericResponse{true, "", dev})
	w.Write(response)
}

func SetElapsedTimeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)

	deviceUUID := vars["uuid"]
	req := SetElapsedTimeRequest{}
	bodyToObject(r.Body, &req)
	if req.ElapsedTime < 0 {
		w.WriteHeader(http.StatusBadRequest)
		response, _ := ToJSON(GenericResponse{false, ErrNegativeNumber.Error(), nil})
		w.Write(response)
		return
	}

	user, err := tokens.FindUser(req.AccessToken)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		response, _ := ToJSON(GenericResponse{false, err.Error(), nil})
		w.Write(response)
		return
	}

	dev, err := user.Devices.FindDevice(deviceUUID)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		response, _ := ToJSON(GenericResponse{false, err.Error(), nil})
		w.Write(response)
		return
	}

	dev.ElapsedSeconds = req.ElapsedTime
	response, _ := ToJSON(GenericResponse{true, "", nil})
	w.Write(response)
	w.WriteHeader(http.StatusOK)
}

func SetPodcastHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)

	deviceUUID := vars["uuid"]
	req := ChangePodcastRequest{}
	bodyToObject(r.Body, &req)
	if req.ElapsedTime < 0 {
		w.WriteHeader(http.StatusBadRequest)
		response, _ := ToJSON(GenericResponse{false, ErrNegativeNumber.Error(), nil})
		w.Write(response)
		return
	}

	user, err := tokens.FindUser(req.AccessToken)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		response, _ := ToJSON(GenericResponse{false, err.Error(), nil})
		w.Write(response)
		return
	}

	dev, err := user.Devices.FindDevice(deviceUUID)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		response, _ := ToJSON(GenericResponse{false, err.Error(), nil})
		w.Write(response)
		return
	}

	dev.ElapsedSeconds = req.ElapsedTime
	dev.CurrentPodcastURL = req.URL
	response, _ := ToJSON(GenericResponse{true, "", nil})
	w.Write(response)
	w.WriteHeader(http.StatusOK)
}

func GetDeviceInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	vars := mux.Vars(r)

	deviceUUID := vars["uuid"]
	req := LoggedInRequest{}
	bodyToObject(r.Body, &req)

	user, err := tokens.FindUser(req.AccessToken)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		response, _ := ToJSON(GenericResponse{false, err.Error(), nil})
		w.Write(response)
		return
	}

	dev, err := user.Devices.FindDevice(deviceUUID)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		response, _ := ToJSON(GenericResponse{false, err.Error(), nil})
		w.Write(response)
		return
	}

	response, _ := ToJSON(GenericResponse{true, "", dev})
	w.Write(response)
	w.WriteHeader(http.StatusOK)
}

func GetDevicesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	req := LoggedInRequest{}
	bodyToObject(r.Body, &req)

	user, err := tokens.FindUser(req.AccessToken)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		response, _ := ToJSON(GenericResponse{false, err.Error(), nil})
		w.Write(response)
		return
	}

	response, _ := ToJSON(GenericResponse{true, "", user.GetDevices()})
	w.Write(response)
	w.WriteHeader(http.StatusOK)
}