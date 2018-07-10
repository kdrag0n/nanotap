package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

var (
	// flag only arguments
	argConfigPath    string
	argWriteConfig   bool
	argProbeMode     bool
	argRequireConfig bool

	// config overrides
	argInputDevice *string
	argMaxFingers  *uint32
)

func ParseArgs() {
	// flag only arguments
	flag.StringVarP(&argConfigPath, "config", "c", "config.toml", "path to the toml config file")
	flag.BoolVarP(&argWriteConfig, "create-config", "w", false, "create a config.toml file with default values")
	flag.BoolVarP(&argProbeMode, "probe", "p", false, "print information of all input devices and exit")
	flag.BoolVarP(&argRequireConfig, "require-config", "r", false, "exit if config reading fails, use defaults otherwise")

	// config overrides
	cflags := flag.NewFlagSet("Config overrides", flag.ExitOnError)
	argInputDevice = cflags.StringP("input-device", "i", DefaultConfig.InputDevice, "input device to read events from, 'auto' to auto-detect")
	argMaxFingers = cflags.Uint32P("max-fingers", "f", DefaultConfig.MaxFingers, "maximum number of fingers to support")

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Nanotap: Fluid Gestures for Android, by @kdrag0n\n\nUsage:")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nConfiguration override flags:")
		cflags.PrintDefaults()
	}

	flag.Parse()
	cflags.Parse(os.Args)
}
