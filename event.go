package main

import (
	"os"
	"time"
	"unsafe"
)

func DecodeRawEvent(r RawEvent, ev *Event) (err error) {
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
		default:
			err = ErrUnknownCode
			return
		}

	case RawEvAbs:
		switch r.Code {
		case RawAbsMtSlot:
			ev.Finger = r.Value
		case RawAbsMtPositionX:
			ev.X = r.Value
		case RawAbsMtPositionY:
			ev.Y = r.Value
		default:
			err = ErrUnknownCode
			return
		}

		if ev.Type == 0 {
			ev.Type = EvFingerMove
		}
	default:
		err = ErrUnknownType
	}

	return
}

func ReadEvents(f *os.File, ch chan Event) {
	buf := make([]byte, RawEventSize)
	var event Event

	for {
		_, err := f.Read(buf)
		check(err)

		bufDataPtr := *(*uintptr)(unsafe.Pointer(&buf))
		rawEvent := *(*RawEvent)(unsafe.Pointer(bufDataPtr))
		err = DecodeRawEvent(rawEvent, &event)
		if err != nil {
			continue
		}

		ch <- event
	}
}
