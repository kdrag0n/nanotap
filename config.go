package main

import (
	"io/ioutil"

	toml "github.com/pelletier/go-toml"
)

var DefaultConfig = Config{
	InputDevice: "/dev/input/event0",
	MaxFingers:  10,
}

type Config struct {
	InputDevice string `toml:"input"`
	MaxFingers  uint32 `toml:"max_fingers"`
}

func ParseConfig(doc []byte) (config Config, err error) {
	config = DefaultConfig

	err = toml.Unmarshal(doc, &config)
	return
}

func LoadConfigFile(path string) (config Config, err error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	config, err = ParseConfig(bytes)
	return
}
