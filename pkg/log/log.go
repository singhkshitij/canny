package log

import (
	"canny/pkg/config"
	"go.uber.org/zap"
)

var (
	Logger *zap.SugaredLogger
	zapper *zap.Logger
)

// Setup initialise logger
func Setup() {
	env := config.Cfg().String("app.release.env")

	if env == "prod" {
		zapper, _ = zap.NewProduction()
	} else {
		zapper, _ = zap.NewDevelopment()
	}
	Logger = zapper.Sugar()

}

func Shutdown() {
	defer func(zapper *zap.Logger) {
		err := zapper.Sync()
		if err != nil {
			Logger.Fatal("Failed to shutdown zapper : %v", err)
		}
	}(zapper)
}
