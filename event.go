package main

import (
	"os"
	"time"
	"unsafe"
)

func DecodeRawEvent(r RawEvent) (ev Event) {
	ev.Time = time.Unix(r.Seconds, r.Microseconds*1000)

	switch r.Type {
	case RawEvKey:
		switch r.Code {
		case RawBtnTouch:
			ev.FingerDown = r.Value == 1
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
		event := DecodeRawEvent(rawEvent)

		ch <- event
	}
}
