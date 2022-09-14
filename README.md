[![GoDoc](https://godoc.org/github.com/qba73/iss?status.png)](http://godoc.org/github.com/qba73/iss)
![Go](https://github.com/qba73/iss/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/qba73/iss)](https://goreportcard.com/report/github.com/qba73/iss)
![GitHub](https://img.shields.io/github/license/qba73/iss)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/qba73/iss)

# ISS

ISS is a Go library for the the [International Space Station](https://en.wikipedia.org/wiki/International_Space_Station) API. It allows to get station's current lat and long coordinates.

## Using the Go library

Import the library using

```go
import github.com/qba73/iss
```

## Creating a client

Create a new client object by calling ```iss.New()```

```go
client, err := iss.New()
if err != nil {
  // handle error
}
```

If you  want to use your ```http.Client```

```go
httpClient := http.Client{}
issClient, err := iss.New(iss.WithHTTPClient(&httpClient))
if err != nil {
  // handle error
}

lat, long, err := issClient.GetPosition()
if err != nil {
  // handle error
}

```

## Retrieving ISS coordinates using client

```go
client, err := iss.New()
if err != nil {
  // handle error
}

position, err := client.GetPosition()
if err != nil {
  // handle error
}

fmt.Println(position)
// Output: {10.5489 1.3942}

```

## Retrieving ISS coordinates using functions

The ```iss``` package provides a high level functions for retrieving ISS coordinates.

```go
lat, long, err := iss.GetPosition()
if err != nil {
  // handle error
}
fmt.Println(lat, long)
// Output: -8.0037 14.7139

```

```go
lat, long, err := iss.GetPositionAsStrings()
if err != nil {
  // handle error
}
fmt.Println(lat, long)
// Output: -11.6732 17.4279

```

## A complete example program

You can see an example program which retrieves the ISS coordinates in the [examples/demo](examples/demo/main.go) folder.
