package db

import (
	"context"
	"time"

	"ariga.io/entcache"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"

	"github.com/anhnmt/golang-gateway-boilerplate/ent"
	"github.com/anhnmt/golang-gateway-boilerplate/ent/migrate"
	"github.com/anhnmt/golang-gateway-boilerplate/internal/bootstrap/config"
)

const (
	defaultTimeout = 10 * time.Second
)

func NewDatabase(
	ctx context.Context,
	redis redis.UniversalClient,
) (*ent.Client, error) {
	if !config.DbEnabled() {
		return nil, nil
	}

	maxOpenConns := config.DbMaxOpenConns()
	if maxOpenConns == 0 {
		maxOpenConns = 15
	}

	maxIdleConns := config.DbMaxIdleConns()
	if maxIdleConns == 0 {
		maxIdleConns = 10
	}

	maxLifetime := config.DbMaxLifetime()
	if maxLifetime == 0 {
		maxLifetime = 5
	}

	pgbouncer := config.DbPgbouncer()

	log.Info().
		Bool("pgbouncer", pgbouncer).
		Int("max_open_conns", maxOpenConns).
		Int("max_idle_conns", maxIdleConns).
		Int("max_lifetime", maxLifetime).
		Msg("Connecting to DB")

	cfg, err := pgx.ParseConfig(config.DbUrl())
	if err != nil {
		return nil, err
	}

	if pgbouncer {
		cfg.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol
	}

	db := stdlib.OpenDB(*cfg)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetMaxOpenConns(maxOpenConns)
	db.SetConnMaxLifetime(time.Duration(maxLifetime) * time.Minute)

	newCtx, cancel := context.WithTimeout(ctx, defaultTimeout)
	defer cancel()

	err = db.PingContext(newCtx)
	if err != nil {
		return nil, err
	}

	drv := entsql.OpenDB(dialect.Postgres, db)

	// Decorates the sql.Driver with entcache.Driver.
	drvCache := entcache.NewDriver(
		drv,
		entcache.TTL(time.Second*5),
		entcache.Levels(
			entcache.NewRedis(redis),
		),
	)

	// Create an ent.Driver from `db`.
	client := ent.NewClient(ent.Driver(drvCache))

	if config.DbDebug() {
		client = client.Debug()
	}

	// Run the auto migration tool.
	if config.DbMigration() {
		if err = client.Schema.Create(
			entcache.Skip(newCtx),
			migrate.WithForeignKeys(false), // Disable foreign keys.
		); err != nil {
			log.Err(err).Msg("Failed creating schema resources")
			return nil, err
		}
	}

	log.Info().Msg("Connecting to DB successfully.")
	return client, nil
}
