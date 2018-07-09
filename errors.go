package main

import "errors"

var (
	ErrUnknownType = errors.New("Unknown event type")
	ErrUnknownCode = errors.New("Unknown event code")
)
