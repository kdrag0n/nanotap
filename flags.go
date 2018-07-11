package main

import (
	"errors"
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
	argInputDevice string
	argMaxFingers  uint32
)

func parseArgs() {
	// flag only arguments
	flag.StringVarP(&argConfigPath, "config", "c", "config.toml", "path to the toml config file")
	flag.BoolVarP(&argWriteConfig, "create-config", "w", false, "create a config.toml file with default values")
	flag.BoolVarP(&argProbeMode, "probe", "p", false, "print information of all input devices and exit")
	flag.BoolVarP(&argRequireConfig, "require-config", "r", false, "exit if config reading fails, use defaults otherwise")

	// config overrides
	def := DefaultConfig
	cflags := flag.NewFlagSet("Config overrides", flag.ExitOnError)
	cflags.StringVarP(&argInputDevice, "input-device", "i", "", fmt.Sprintf(`input device to read events from, 'auto' to auto-detect (default "%s")`, def.InputDevice))
	cflags.Uint32VarP(&argMaxFingers, "max-fingers", "f", 0, fmt.Sprintf(`maximum number of fingers to support (default %d)`, def.MaxFingers))

	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Nanotap: Fluid Gestures for Android by @kdrag0n\n\nUsage:")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "\nConfiguration override flags:")
		cflags.PrintDefaults()
	}

	cflags.VisitAll(func(fl *flag.Flag) {
		flagCopy := *fl
		flagCopy.Hidden = true

		flag.CommandLine.AddFlag(&flagCopy)
	})

	flag.ErrHelp = errors.New("")
	flag.Parse()
}

func processConfigOverrides(cfg *Config) (overrides uint) {
	if argInputDevice != "" {
		cfg.InputDevice = argInputDevice
		overrides++
	}

	if argMaxFingers != 0 {
		cfg.MaxFingers = argMaxFingers
		overrides++
	}

	return
}
