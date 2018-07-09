package main

import (
	"syscall"
	"time"
	"unsafe"
)

type RawEvent struct {
	Time  syscall.Timeval
	Type  uint16
	Code  uint16
	Value uint32
}

var _sREvent RawEvent

const RawEventSize = unsafe.Sizeof(_sREvent)

const (
	RawEvKey           = 1
	RawEvAbs           = 3
	RawBtnTouch        = 330
	RawBtnToolFinger   = 325
	RawAbsMtTouchMajor = 48
	RawAbsMtTouchMinor = 49
	RawAbsMtPositionX  = 53
	RawAbsMtPositionY  = 54
	RawAbsMtTrackingID = 57
	RawAbsMtSlot       = 47
)

type Event struct {
	Finger     uint32
	Type       uint16
	Time       time.Time
	FingerDown bool
	X          uint32
	Y          uint32
}

const (
	EvFingerDown uint16 = iota
	EvFingerUp
	EvFingerMove
)
