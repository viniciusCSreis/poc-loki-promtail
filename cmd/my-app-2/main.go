package main

import (
	"errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
	"time"
)

var (
	logLevel = os.Getenv("LOG_LEVEL")
)

func main() {

	if logLevel == "" {
		logLevel = "info"
	}

	// zerolog
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	level, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		log.Fatal().Err(err).Msg("Fail to create consul client")
	}
	zerolog.SetGlobalLevel(level)

	log.Info().Msg("Starting")

	go func() {
		http.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
			writer.WriteHeader(http.StatusOK)
			_, _ = writer.Write([]byte("UP"))
		})
		log.Fatal().Err(http.ListenAndServe(":3000", nil)).Msg("Web server Down.")
	}()

	count := 0

	for {
		count++
		time.Sleep(time.Second * 1)

		switch count % 3 {
		case 1:
			log.Info().Msgf("Hello %d", count)
		case 2:
			log.Warn().Msgf("Ola %d", count)
		default:
			log.Error().Err(errors.New("some Error")).Msgf("Hola %d", count)
		}

	}
}
