package main

import (
	"fmt"

	evdev "github.com/kdrag0n/golang-evdev"
)

func ProbeInputDevice(verbosity uint) (path string, err error) {
	devices, err := evdev.ListInputDevices("/dev/input/event*")
	if err != nil {
		return
	}

devloop:
	for _, dev := range devices {
		log.Debug().Str("path", dev.Fn).Str("name", dev.Name).Str("phys", dev.Phys).Msg("Found device")

		for cType, cCodes := range dev.Capabilities {
			if verbosity > 0 {
				log.Debug().Int("type", cType.Type).Str("name", cType.Name).Msg("Found capability")
				if verbosity > 1 {
					fmt.Print("Codes: ")
					for _, cCode := range cCodes {
						fmt.Printf("%d%s ", cCode.Code, cCode.Name)
					}
					fmt.Print("\n")
				}
			}

			if cType.Type == 3 {
				path = dev.Fn
				log.Info().Str("device", path).Msg("Using device")
				break devloop
			}
		}
	}

	return
}
