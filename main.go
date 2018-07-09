package main

import (
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
	config, err := LoadConfigFile("config.toml")
	if err != nil {
		log.Warn().Err(err).Msg("Unable to load config, using defaults")
		config = DefaultConfig
	}
	log.Print("Loaded config")

	if config.InputDevice == "auto" {
		log.Print("Auto-detecting input device...")
		path, err := ProbeInputDevice()
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

		log.Printf("Finger %d (%4d, %4d) %s", event.Finger, event.X, event.Y, typeStr)
	}
}
