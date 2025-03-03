package main

import (
	"github.com/dedbee/Calcserv_Go/internal/agent"
	"github.com/dedbee/Calcserv_Go/internal/application"
)

func main() {
	app := application.New()
	go app.RunServer()

	agent.StartAgent()
}
