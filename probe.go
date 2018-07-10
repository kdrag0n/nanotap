package main

import (
	"fmt"

	evdev "github.com/gvalkov/golang-evdev"
)

func ProbeInputDevice(verbosity uint) (path string, err error) {
	devices, err := evdev.ListInputDevices("/dev/input/event*")
	if err != nil {
		return
	}

devloop:
	for _, dev := range devices {
		log.Info().Str("path", dev.Fn).Str("name", dev.Name).Str("phys", dev.Phys).Msg("Found device")
		if verbosity > 0 {
			log.Info().Uint16("bustype", dev.Bustype).Uint16("vendor", dev.Vendor).Uint16("product", dev.Product).Uint16("version", dev.Version).Int("evVersion", dev.EvdevVersion).Msg("ID")
		}

		for cType, cCodes := range dev.Capabilities {
			if verbosity > 0 {
				log.Printf("Capability: %d %s", cType.Type, cType.Name)
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
