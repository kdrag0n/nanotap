package main

import (
	"os"

	"github.com/mattn/go-isatty"
	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

var log zerolog.Logger

func init() {
	isTerminal := isatty.IsTerminal(os.Stdout.Fd())
	if isTerminal {
		zlog.Logger = zlog.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	dwr := NewDiodeWriter(os.Stdout, 1024, func(missed int) {
		zlog.Warn().Int("count", missed).Msg("Dropped messages")
	})

	log = zerolog.New(dwr).With().Timestamp().Logger()
	if isTerminal {
		log = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})
	}

	if isTerminal {
		zerolog.TimeFieldFormat = "Mon Jan 2 15:04:05"
	} else {
		zerolog.TimeFieldFormat = ""
	}
}
