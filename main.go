package main

import (
	"canny/pkg/log"
	"canny/pkg/setting"
	"canny/routers"
	"fmt"
	"net/http"
)

// init Initialise project configs
func init() {
	log.Setup()
	setting.Setup()
}

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
