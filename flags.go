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
	flag.StringVarP(&argConfigPath, "config", "c", "config.toml", "path to the toml config file")
	flag.BoolVarP(&argWriteConfig, "create-config", "w", false, "create a config.toml file with default values")
	flag.BoolVarP(&argProbeMode, "probe", "p", false, "print information of all input devices and exit")
	flag.BoolVarP(&argRequireConfig, "require-config", "r", false, "exit if config reading fails, use defaults otherwise")

	// config overrides
	flag.StringVarP(&argInputDevice, "input-device", "i", DefaultConfig.InputDevice, "input device to read events from, 'auto' to auto-detect")
	flag.Uint32VarP(&argMaxFingers, "max-fingers", "f", DefaultConfig.MaxFingers, "maximum number of fingers to support")

	flag.Parse()
}
