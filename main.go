package main

import (
	"capi/app"
	"capi/logger"
)

// capi = nama module di go.mod // app = nama folder -> package
// this is called refactor in golang, quite similar to context mgt in reactjs
// task: get data customer

func main() {
	// log.Println("Starting application...")
	logger.Info("Starting application...")
	app.Start()
}
