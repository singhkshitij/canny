package log

import "go.uber.org/zap"

var Logger *zap.SugaredLogger

// Setup initialise logger
func Setup() {
	zapper, _ := zap.NewProduction()
	defer zapper.Sync()
	Logger = zapper.Sugar()
}
