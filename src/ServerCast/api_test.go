package main

import (
	"testing"
	"net/http"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"github.com/satori/go.uuid"
)

func init() {
	go main()
}

func TestCanRegister(t *testing.T) {
	req := RegisterRequest{"emile", "password"}
	requestJSON, _ := ToJSON(req)

	r, err := http.NewRequest("POST", "http://localhost:8080/register", bytes.NewBuffer(requestJSON))
	if err != nil {
		t.Fatal(err)
	}
	r.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	response := GenericResponse{}
	bodyRead, _ := ioutil.ReadAll(resp.Body)
	r.Body.Close()
	err = json.Unmarshal(bodyRead, &response)
	if err != nil {
		t.Fatal(err)
	}

	if !response.Success {
		t.Fatal("Not successful", response)
	}
}

func TestCanLogin(t *testing.T) {
	req := LoginRequest{"emile", "password"}
	requestJSON, _ := ToJSON(req)

	r, err := http.NewRequest("POST", "http://localhost:8080/login", bytes.NewBuffer(requestJSON))
	if err != nil {
		t.Fatal(err)
	}
	r.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	type testResponse struct {
		Success bool `json:"OK"`
		Message string `json:"detail,omitempty"`
		Value LoginResponse `json:"value,omitempty"`
	}

	response := testResponse{}
	bodyRead, _ := ioutil.ReadAll(resp.Body)
	r.Body.Close()
	err = json.Unmarshal(bodyRead, &response)
	if err != nil {
		t.Fatal(err)
	}

	if !response.Success {
		t.Fatal("Not successful", string(bodyRead))
	}
	if _, err = uuid.FromString(response.Value.AccessToken); err != nil {
		t.Fatal("Invalid token UUID")
	}

}

func TestCanAddDevice(t *testing.T) {
	req :=  CreateDeviceRequest{LoggedInRequest{tokens.Tokens[0].Token},
		"DeviceName"}
	requestJSON, _ := ToJSON(req)

	r, err := http.NewRequest("POST", "http://localhost:8080/device/add", bytes.NewBuffer(requestJSON))
	if err != nil {
		t.Fatal(err)
	}
	r.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(r)
	if err != nil {
		t.Fatal(err)
	}
	type testResponse struct {
		Success bool `json:"OK"`
		Message string `json:"detail,omitempty"`
		Value Device `json:"value,omitempty"`
	}
	response := testResponse{}
	bodyRead, _ := ioutil.ReadAll(resp.Body)
	r.Body.Close()
	err = json.Unmarshal(bodyRead, &response)
	if err != nil {
		t.Fatal(err)
	}

	if !response.Success {
		t.Fatal("Not successful", response)
	}
	t.Log(response)
	if _, err := uuid.FromString(response.Value.UUID); err != nil {
		t.Fatal("Invalid device UUID")
	}
}