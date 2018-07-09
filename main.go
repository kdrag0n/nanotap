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
	f, err := os.Open("/dev/input/event0")
	check(err)

	eventChan := make(chan Event)
	go ReadEvents(10, f, eventChan)

	for {
		event := <-eventChan

		status := "down"
		if event.Type == EvFingerUp {
			status = "up"
		}

		var typeStr string
		switch event.Type {
		case EvFingerDown:
			typeStr = "FingerDown"
		case EvFingerUp:
			typeStr = "FingerUp"
		case EvFingerMove:
			typeStr = "FingerMove"
		}

		fmt.Printf("Finger %d %s @ (%d, %d); %s\n", event.Finger, status, event.X, event.Y, typeStr)
	}
}
