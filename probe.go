package main

import (
	"fmt"
	"strings"

	evdev "github.com/gvalkov/golang-evdev"
)

func ProbeInputDevice() (path string, err error) {
	devices, err := evdev.ListInputDevices("/dev/input/event*")
	if err != nil {
		return
	}

	for _, dev := range devices {
		log.Info().Str("path", dev.Fn).Str("name", dev.Name).Str("phys", dev.Phys).Msg("Found device")
		log.Info().Uint16("bustype", dev.Bustype).Uint16("vendor", dev.Vendor).Uint16("product", dev.Product).Uint16("version", dev.Version).Int("evVersion", dev.EvdevVersion).Msg("ID")

		for cType, cCodes := range dev.Capabilities {
			log.Printf("Capability: %d %s", cType.Type, cType.Name)
			fmt.Print("Codes: ")
			for _, cCode := range cCodes {
				fmt.Printf("%d%s ", cCode.Code, cCode.Name)
			}
			fmt.Print("\n")
		}

		if strings.Contains(dev.Name, "touch") {
			path = dev.Fn
			log.Warn().Str("device", path).Msg("Using device")
		}
	}

	return
}
