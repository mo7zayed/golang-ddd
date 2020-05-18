package main

import (
	"flag"
	"golang-ddd/infrastructure/bootstrap"
)

var port string

func init() {
	flag.StringVar(&port, "port", "5000", "The Port The Application Will Use To Be Served.")

	flag.Parse()
}

func main() {
	bootstrap.StartApplication(
		port,
	)
}
