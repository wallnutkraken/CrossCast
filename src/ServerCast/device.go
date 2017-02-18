package main

import "github.com/satori/go.uuid"

type Device struct {
	Name string `json:"name"`
	CurrentPodcastURL string `json:"current_podcast_url,omitempty"`
	UUID string `json:"uuid"`
	ElapsedSeconds int `json:"elapsed_seconds,omitempty"`
}

func NewDevide() Device {
	return Device{UUID: uuid.NewV4().String()}
}

type Devices struct {
	List []Device
}

func (d *Devices) Add(name string) Device {
	dev := Device{Name: name}
	dev.UUID = uuid.NewV4().String()
	d.List = append(d.List, dev)
	return dev
}