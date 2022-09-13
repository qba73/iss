![Go](https://github.com/qba73/iss/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/qba73/iss)](https://goreportcard.com/report/github.com/qba73/iss)
![GitHub](https://img.shields.io/github/license/qba73/iss)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/qba73/iss)

# ISS
ISS is a Go library for the the [International Space Station](https://en.wikipedia.org/wiki/International_Space_Station) API. It allows to get station's lat and long coordinates.

## Using the Go library

Import the library using
```go
import github.com/qba73/iss
```

## Creating a client
Create a new client object by calling ```iss.New()```
```go
client := iss.New()
```

## Retrieving ISS coordinates
```go
client := iss.New()
client.GetPosition()
```


