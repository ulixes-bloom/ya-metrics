package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	grpcclient "github.com/ulixes-bloom/ya-metrics/internal/agent/client/grpc"
	httpclient "github.com/ulixes-bloom/ya-metrics/internal/agent/client/http"
	"github.com/ulixes-bloom/ya-metrics/internal/agent/config"
	"github.com/ulixes-bloom/ya-metrics/internal/agent/memory"
)

var (
	buildVersion string = "N/A"
	buildDate    string = "N/A"
	buildCommit  string = "N/A"
)

func main() {
	log.Info().Msgf("Build version: %s\nBuild date: %s\nBuild commit: %s\n", buildVersion, buildDate, buildCommit)

	conf, err := config.Parse()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}

	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	defer stop()

	logLvl, err := zerolog.ParseLevel(conf.LogLvl)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to parse log level")
	}
	zerolog.SetGlobalLevel(logLvl)

	ms := memory.NewStorage()

	switch conf.Protocol {
	case "http":
		cl, err := httpclient.New(conf, ms)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}

		cl.Run(ctx)
	case "grpc":
		cl, err := grpcclient.New(conf, ms)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}

		cl.Run(ctx)
	default:
		log.Fatal().Msgf("unknown client protocol %s", conf.Protocol)
	}
}
