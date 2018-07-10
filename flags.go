package main

import (
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

func ParseArgs() {
	// flag only arguments
	flag.StringVar(&argConfigPath, "config", "config.toml", "path to the toml config file")
	flag.BoolVar(&argWriteConfig, "create-config", false, "create a config.toml file with default values")
	flag.BoolVar(&argProbeMode, "probe", false, "print information of all input devices and exit")
	flag.BoolVar(&argRequireConfig, "require-config", false, "exit if config reading fails, use defaults otherwise")

	// config overrides
	flag.StringVar(&argInputDevice, "input-device", DefaultConfig.InputDevice, "input device to read events from, 'auto' to auto-detect")
	flag.Uint32Var(&argMaxFingers, "max-fingers", DefaultConfig.MaxFingers, "maximum number of fingers to support")

	flag.Parse()
}
