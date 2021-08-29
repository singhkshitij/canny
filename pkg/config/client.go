package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"log"
)

var client = koanf.New(".")

func Setup() {
	f := file.Provider("config/dev.yaml")
	client.Load(f, yaml.Parser())

	// Watch the file and get a callback on change. The callback can do whatever,
	// like re-load the configuration.
	// File provider always returns a nil `event`.
	f.Watch(func(event interface{}, err error) {
		if err != nil {
			log.Printf("config watch error: %v", err)
			return
		}
		// Throw away the old config and load a fresh copy.
		log.Println("Config changed. Reloading ...")
		client.Load(f, json.Parser())
	})
}

func Cfg() *koanf.Koanf {
	return client
}

