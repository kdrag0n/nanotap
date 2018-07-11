package main

import (
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		log.Panic().Err(e).Msg("Fatal error")
	}
}

func checkMsg(e error, msg string) {
	if e != nil {
		log.Panic().Err(e).Msg(msg)
	}
}

func main() {
	parseArgs()

	if argWriteConfig {
		cfgBytes, err := Asset("config.toml")
		checkMsg(err, "Unable to load default config")

		log.Info().Msg("Writing default config to config.toml...")
		err = ioutil.WriteFile("config.toml", cfgBytes, 644)
		checkMsg(err, "Unable to write config")

		log.Info().Msg("Config written")
		return
	} else if argProbeMode {
		log.Info().Msg("Probing devices...")
		ProbeInputDevice(1)
		return
	}

	config, err := LoadConfigFile(argConfigPath)
	if err != nil {
		if argRequireConfig {
			log.Panic().Err(err).Msg("Unable to load config")
		}

		log.Warn().Err(err).Msg("Unable to load config, using defaults")
		config = DefaultConfig
	}

	overwritten := processConfigOverrides(&config)
	log.Debug().Uint("overwritten", overwritten).Msg("Loaded config")

	if config.InputDevice == "auto" {
		log.Info().Msg("Auto-detecting input device...")
		path, err := ProbeInputDevice(1)
		checkMsg(err, "Unable to find valid input device")

		config.InputDevice = path
	}

	f, err := os.Open(config.InputDevice)
	check(err)
	log.Debug().Str("device", config.InputDevice).Msg("Opened input device")

	eventChan := make(chan Event)
	go ReadEvents(config.MaxFingers, f, eventChan)
	log.Print("Started event reader")

	log.Info().Msg("Getting events")
	for {
		event := <-eventChan

		var typeStr string
		switch event.Type {
		case EvFingerDown:
			typeStr = "down"
		case EvFingerUp:
			typeStr = "up"
		case EvFingerMove:
			typeStr = "move"
		}

		log.Debug().Uint32("x", event.X).Uint32("y", event.Y).Str("action", typeStr).Msgf("Event @ (%4d, %4d)", event.X, event.Y)
	}
}
