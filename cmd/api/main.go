package main

import (
	"api-http-hex-golang/cmd/api/bootstrap"
	"log"
)

func main() {
	if err := bootstrap.Run(); err != nil {
		log.Fatal()
	}
}
