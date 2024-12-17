package main

import (
	"github.com/dedbee/Calcserv_Go/internal/application"
)

func main() {
	app := application.New()
	//app.Run()
	app.RunServer()
}
