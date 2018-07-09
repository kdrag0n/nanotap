package main

import (
	"fmt"
	"os"
	"time"
	"unsafe"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func DecodeEvent(r RawEvent) (ev Event) {
	ev.Time = time.Unix(r.Seconds, r.Microseconds*1000)

	switch r.Type {
	case RawEvKey:
		switch r.Code {
		case RawBtnTouch:
			if r.Value == 1 {
				ev.FingerDown = true
				ev.Type = EvFingerDown
			} else {
				ev.FingerDown = false
				ev.Type = EvFingerUp
			}
		}

	case RawEvAbs:
		switch r.Code {
		case RawAbsMtSlot:
			ev.Finger = r.Value
		case RawAbsMtPositionX:
			ev.X = r.Value
		case RawAbsMtPositionY:
			ev.Y = r.Value
		}

		if ev.Type == 0 {
			ev.Type = EvFingerMove
		}
	}

	return
}

func ReadEvents(f *os.File, ch chan Event) {
	buf := make([]byte, RawEventSize)

	for {
		_, err := f.Read(buf)
		check(err)

		bufDataPtr := *(*uintptr)(unsafe.Pointer(&buf))
		rawEvent := *(*RawEvent)(unsafe.Pointer(bufDataPtr))
		event := DecodeEvent(rawEvent)

		ch <- event
	}
}

func main() {
	f, err := os.Open("/dev/input/event0")
	check(err)

	eventChan := make(chan Event)
	go ReadEvents(f, eventChan)

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
