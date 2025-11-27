package main

import (
	"fmt"

	"github.com/rs/zerolog"
	zlog "github.com/rs/zerolog/log"
)

func main() {
	fmt.Println("Logs são legáis:3")

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	zlog.Print("Isso é um log")

	zlog.Info().Str("Amo", "Logs").Msg("Isso também é um log")
}
