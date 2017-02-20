# API definitions

CrossCast uses a very basic RESTful HTTP API.

## API object definitions

These are the objects that can be sent to the server in JSON form to interact with the API.
The description of each API Call will mention which Request or response
objects are used for the call.

### Request objects

#### RegisterRequest
```json
{ 
  "username" : "unique_username",
  "password" : "password123"
}
```
The username `username` here is the unique username the user wants to
register with. `password` is the password the user wants to use for the
account.

#### LoginRequest

Object is the same as [RegisterRequest](#RegisterRequest), except it's
used for logging in.

#### LoggedInRequest
```json
{
  "token" : "unique-token"
}
```
`LoggedInRequest` is used for or in most requests that the user needs to be logged in for.
Tokens expire after 24 hours, and re logging in creates a new token.

#### CreateDeviceRequest
```json
{
  "token" : "unique-token",
  "device_name" : "name for device"
}
```
`device_name` is not unique.

#### SetElapsedTimeRequest
```json
{
  "token" : "unique-token",
  "elapsed_time" : 60
}
```
`elapsed_time` is an integer that represents seconds of elapsed time for the
current file.

#### ChangePodcastRequest
```json
  "token" : "unique-token",
  "url" : "http://www.podcast-example.com/example/podcast.mp3",
  "elapsed_time" : 60
```

Almost the same as [SetElapsedTimeRequest](#SetElapsedTimeRequest), except
for `url`, which is the URL for the currently played file for the
device.

### Response objects

#### GenericResponse

```json
{
  "OK" : true,
  "detail" : "Here's your login token",
  "value" : {
    "token" : "d459c89a-6fc0-4369-ad10-cf1e1263f8d9"
  }
}
```

`GenericResponse` is the base response for pretty much every single API call.
`OK` tells the caller whether the request was successful,
`detail` may give additional detail (may be omitted) and `value` can
be any kind of JSON object, and is usually where the data the caller
is looking for will be.

The rest of this section will be the definition of objects that will
be retrieved from the `value` element.

#### LoginResponse

See example for [GenericResponse](#GenericResponse).

#### Device

```json
{
  "name" : "device name",
  "current_podcast_url" : "http://www.podcast-example.com/example/podcast.mp3",
  "uuid" : "baabca25-7341-42dd-88c0-c9f9a9874217",
  "elapsed_seconds" : 554
}
```

## Available methods 
`api_url` is the URL that the API runs on.
All API calls require JSON objects to be in the request body.
All API calls use `POST` unless otherwise specified.

All methods will be represented by a table like this:

|URL|Request Object|Response object|
|---|---|---|
||   |   |

`URL` being the URL that the method needs to be called on,
`Request Object` being the object that the user of the API must send to
the given `URL` as a JSON object. `Response Object` being the type
of the object that the API will respond with in the `value` field of
[GenericResponse](#GenericResponse), which may be omitted depending on 
the method.

### Login

|URL|Request Object|Response object|
|---|---|---|
|api_url/login|[LoginRequest](#LoginRequest)|[LoginResponse](#LoginResponse)|

### Register

|URL|Request Object|Response object|
|---|---|---|
|api_url/register|[RegisterRequest](#RegisterRequest)|N/A|

### NewDevice

|URL|Request Object|Response object|
|---|---|---|
|api_url/device/add|[CreateDeviceRequest](#CreateDeviceRequest)|[Device](#Device)|

### SetElapsedTime

|URL|Request Object|Response object|
|---|---|---|
|api_url/device/uuid/elapsed|[SetElapsedTimeRequest](#SetElapsedTimeRequest)|N/A|

Where `uuid` is the UUID of the device to be affected.

### SetPodcastInfo

|URL|Request Object|Response object|
|---|---|---|
|api_url/device/uuid/podcast|[ChangePodcastRequest](#ChangePodcastRequest)|N/A|

Where `uuid` is the UUID of the device to be affected.

### GetDeviceInfo

|URL|Request Object|Response object|
|---|---|---|
|api_url/device/uuid|[LoggedInRequest](#LoggedInRequest)|[Device](#Device)|

Where `uuid` is the UUID of the device to be affected.

### GetDevices

|URL|Request Object|Response object|
|---|---|---|
|api_url/device|[LoggedInRequest](#LoggedInRequest)|Array of [Device](#Device)|

