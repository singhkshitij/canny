package log

import (
	"canny/pkg/config"
	"go.uber.org/zap"
)

var Logger *zap.SugaredLogger

// Setup initialise logger
func Setup() {
	env := config.Cfg().String("model.release.env")
	var zapper *zap.Logger
	if env == "prod" {
		zapper, _ = zap.NewProduction()
	} else {
		zapper, _ = zap.NewDevelopment()
	}
	defer zapper.Sync()
	Logger = zapper.Sugar()

}
