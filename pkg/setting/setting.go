package setting

import "time"

type Server struct {
	Port        int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

// Setup TODO: values should come from config library
func Setup() {
	ServerSetting.Port = 8090
	ServerSetting.ReadTimeout = 10 * time.Second
	ServerSetting.WriteTimeout = 10 * time.Second
}
