package main

import (
	"os"
	"time"

	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	zlog "github.com/rs/zerolog/log"
)

var log zerolog.Logger

func init() {
	isTerminal := isatty.IsTerminal(os.Stdout.Fd())
	if isTerminal {
		zlog.Logger = zlog.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	dwr := diode.NewWriter(os.Stdout, 1024, 10*time.Millisecond, func(missed int) {
		zlog.Warn().Int("count", missed).Msg("Dropped messages")
	})

	log = zerolog.New(dwr)
	if isTerminal {
		log = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}
}
