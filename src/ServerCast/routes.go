package main

import "net/http"

type Route struct {
	Name string
	Method string
	Pattern string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Login",
		"POST",
		"/login",
		LoginHandler,
	},
	Route{
		"Register",
		"POST",
		"/register",
		RegisterHandler,
	},
	Route {
		"NewDevice",
		"POST",
		"/device/add",
		NewDeviceHandler,
	},
	Route{
		"SetElapsedTime",
		"POST",
		"/device/{uuid}/elapsed",
		SetElapsedTimeHandler,
	},
	Route{
		"SetPodcastInfo",
		"POST",
		"/device/{uuid}/podcast",
		SetPodcastHandler,
	},
	Route{
		"GetDeviceInfo",
		"POST",
		"/device/{uuid}",
		GetDeviceInfoHandler,
	},
	Route{
		"GetDevices",
		"POST",
		"/device",
		GetDevicesHandler,
	},
}