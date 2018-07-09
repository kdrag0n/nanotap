package main

import (
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	config, err := LoadConfigFile("config.toml")
	check(err)

	f, err := os.Open(config.InputDevice)
	check(err)

	eventChan := make(chan Event)
	go ReadEvents(config.MaxFingers, f, eventChan)

	for {
		event := <-eventChan

		var typeStr string
		switch event.Type {
		case EvFingerDown:
			typeStr = "down"
		case EvFingerUp:
			typeStr = "up"
		}

		fmt.Printf("Finger %d (%4d, %4d) %s\n", event.Finger, event.X, event.Y, typeStr)
	}
}
