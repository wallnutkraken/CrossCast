package main

import "github.com/satori/go.uuid"

type Device struct {
	CurrentPodcastURL string
	UUID string
	PlaybackPosition int
}

func NewDevide() Device {
	return Device{UUID: uuid.NewV4().String()}
}
