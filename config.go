package main

//go:generate go-bindata -nomemcopy -nometadata -nocompress -o bindata.go config.toml

import (
	"io/ioutil"

	toml "github.com/pelletier/go-toml"
)

var DefaultConfig Config

func init() {
	bConfig, err := Asset("config.toml")
	checkMsg(err, "Unable to load default config")

	config, err := ParseConfig(bConfig)
	checkMsg(err, "Unable to parse default config")

	DefaultConfig = config
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
