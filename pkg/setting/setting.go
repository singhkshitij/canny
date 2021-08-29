package setting

import (
	"canny/pkg/config"
	"time"
)

type Server struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

func Setup() {
	readTimeout := config.Cfg().Int("app.server.readTimeout")
	writeTimeout := config.Cfg().Int("app.server.writeTimeout")

	ServerSetting.Port = config.Cfg().Int("app.server.port")
	ServerSetting.ReadTimeout = time.Duration(readTimeout) * time.Second
	ServerSetting.WriteTimeout = time.Duration(writeTimeout) * time.Second
}
