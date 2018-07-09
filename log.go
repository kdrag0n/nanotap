package main

import (
	"fmt"
	"os"
	"time"

	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
)

var log zerolog.Logger

func init() {
	dwr := diode.NewWriter(os.Stdout, 1000, 10*time.Millisecond, func(missed int) {
		fmt.Printf("Dropped %d messages", missed)
	})

	log = zerolog.New(dwr)
	if isatty.IsTerminal(os.Stdout.Fd()) {
		log = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
}
