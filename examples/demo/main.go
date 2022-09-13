package main

import (
	"fmt"
	"log"

	"github.com/qba73/iss"
)

func main() {
	lat, long, err := iss.GetPosition()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(lat, long)
}
