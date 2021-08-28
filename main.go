package main

import (
	"canny/pkg/log"
	"canny/pkg/setting"
	"canny/routers"
	"canny/worker"
	"fmt"
	"net/http"
)

// init Initialise project configs
func init() {
	log.Setup()
	setting.Setup()
	worker.Setup()
	worker.InitialiseData()
}

// @title Canny documentation
// @version 1.0
// @description This is a sample server Canny server.
// @termsOfService http://swagger.io/terms/

// @contact.name Kshitij
// @contact.url https://github.com/singhkshitij
// @contact.email singh_kshitij@yahoo.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /
func main() {
	initRouter := routers.InitRouter()

	serverPort := fmt.Sprintf(":%d", setting.ServerSetting.Port)
	server := &http.Server{
		Addr:         serverPort,
		Handler:      initRouter,
		ReadTimeout:  setting.ServerSetting.ReadTimeout,
		WriteTimeout: setting.ServerSetting.WriteTimeout,
	}

	log.Logger.Infow("Starting server with configs : " , "port", serverPort)

	server.ListenAndServe()
}
