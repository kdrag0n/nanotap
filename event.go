package main

import (
	"os"
	"time"
	"unsafe"
)

func DecodeRawEvent(r RawEvent, ev *Event) (err error) {
	ev.Time = time.Unix(r.Time.Sec, r.Time.Usec*1000)

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

func ReadEvents(maxFingers uint32, f *os.File, ch chan Event) {
	buf := make([]byte, RawEventSize)
	slots := make([]*Event, maxFingers)
	var currentSlot uint32

	for slot, _ := range slots {
		slots[slot] = &Event{
			Finger: uint32(slot),
		}
	}

	for {
		_, err := f.Read(buf)
		check(err)

		bufDataPtr := *(*uintptr)(unsafe.Pointer(&buf))
		rawEvent := *(*RawEvent)(unsafe.Pointer(bufDataPtr))
		event := slots[currentSlot]
		err = DecodeRawEvent(rawEvent, event)
		if err != nil {
			continue
		}

		currentSlot = event.Finger
		if currentSlot > maxFingers-1 {
			currentSlot = maxFingers - 1
		}

		ch <- *event
	}
}
