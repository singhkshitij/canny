package config

import (
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"log"
	"strings"
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

	// Load environment variables and merge into the loaded config.
	// "CANNY" is the prefix to filter the env vars by.
	// "." is the delimiter used to represent the key hierarchy in env vars.
	// The (optional, or can be nil) function can be used to transform
	// the env var names, for instance, to lowercase them.
	//
	// For example, env vars: CANNY_TYPE and CANNY_PARENT1_CHILD1_NAME
	// will be merged into the "type" and the nested "parent1.child1.name"
	// keys in the config file here as we lowercase the key, 
	// replace `_` with `.` and strip the CANNY_ prefix so that 
	// only "parent1.child1.name" remains.
	client.Load(env.Provider("CANNY_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "CANNY_")), "_", ".", -1)
	}), nil)
}

func Cfg() *koanf.Koanf {
	return client
}

