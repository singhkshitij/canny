package scheduler

import (
	"canny/pkg/log"
	"github.com/go-co-op/gocron"
	"time"
)

var client *gocron.Scheduler

type fn func()

func Setup() {
	client = gocron.NewScheduler(time.UTC)
	client.TagsUnique()
	client.StartAsync()
}

func Add(t int, f fn) {
	client.Every(t).Hours().Tag("price").Do(f)
	log.Logger.Info("Price cache refresh job added !")
}
