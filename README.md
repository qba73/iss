# ISS

![Go](https://github.com/qba73/iss/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/qba73/iss)](https://goreportcard.com/report/github.com/qba73/iss)
![GitHub](https://img.shields.io/github/license/qba73/iss)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/qba73/iss)


ISS is a Go library for the the [International Space Station](https://en.wikipedia.org/wiki/International_Space_Station) API. It allows to get station's current lat and long coordinates.

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

## Retrieving ISS coordinates using client

```go
client := iss.New()
position, err := client.GetPosition()
fmt.Println(position)
// Output: {10.5489 1.3942}
```

## Retrieving ISS coordinates using functions

The ```iss``` package provides a high level functions for retrieving ISS coordinates.

```go
lat, long, err := iss.GetPosition()
fmt.Println(lat, long)
// Output: -8.0037 14.7139
```

```go
lat, long, _ := iss.GetPositionAsStrings()
fmt.Println(lat, long)
// Output: -11.6732 17.4279
```

## A complete example program

You can see an example program which retrieves the ISS coordinates in the [examples/demo](examples/demo/main.go) folder.
