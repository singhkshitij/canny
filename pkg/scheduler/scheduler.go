package scheduler

import (
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

func Add(t int, f fn, tag string) {
	client.Every(t).Hours().Tag(tag).Do(f)
}
