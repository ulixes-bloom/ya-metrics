package main

import (
	"context"
	"database/sql"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/ulixes-bloom/ya-metrics/internal/server/api"
	"github.com/ulixes-bloom/ya-metrics/internal/server/config"
	"github.com/ulixes-bloom/ya-metrics/internal/server/service"
	"github.com/ulixes-bloom/ya-metrics/internal/server/storage/memory"
	"github.com/ulixes-bloom/ya-metrics/internal/server/storage/pg"
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
		log.Fatal().Err(err).Msg("unable to parse config")
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

	var storage service.Storage
	if conf.DatabaseDSN != "" {
		db, err := sql.Open("pgx", conf.DatabaseDSN)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}
		ps, err := pg.NewStorage(ctx, db)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}
		storage = ps
	} else {
		ms, err := memory.NewStorage(ctx, conf)
		if err != nil {
			log.Fatal().Msg(err.Error())
		}
		storage = ms
	}

	err = api.New(conf, storage).Run(ctx)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
}
