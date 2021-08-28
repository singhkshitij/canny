package cache

import (
	"canny/pkg/alphavantage"
	"canny/pkg/log"
	"github.com/patrickmn/go-cache"
	"time"
)

var client *cache.Cache

func Setup() {
	client = cache.New(48*time.Hour, 24*time.Hour)
}

func Set(key string, value *alphavantage.DailyCurrencyDataResult){
	client.Set(key, value, cache.DefaultExpiration)
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
