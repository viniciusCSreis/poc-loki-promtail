package main

import (
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

	for {
		time.Sleep(time.Second * 1)
		log.Info().Msg("Hello 1")
	}
}
