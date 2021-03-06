package cache

import (
	"canny/pkg/log"
	cac "github.com/patrickmn/go-cache"
	"time"
)

var client *cac.Cache

func Setup() {
	client = cac.New(48*time.Hour, 24*time.Hour)
}

func Set(keyName string, value interface{}) {
	client.Set(keyName, value, cac.DefaultExpiration)
}

func Get(key string) interface{} {
	data, found := client.Get(key)
	if found != true {
		log.Logger.Errorf("Cache key lookup failed for key %s", key)
		return nil
	}
	log.Logger.Infof("Look succeded for key %s", key)
	return data
}
