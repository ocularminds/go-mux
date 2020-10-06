package main

import (
	"os"

	"ocularminds.com/mux"
)

func main() {
	a := mux.App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)
	a.Run(":8010")
}

//export APP_DB_USERNAME=postgres
//export APP_DB_PASSWORD=
//export APP_DB_NAME=postgres
