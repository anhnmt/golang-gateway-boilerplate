package config

import (
	"github.com/spf13/viper"
)

func DbEnabled() bool {
	return viper.GetBool("DB_ENABLED")
}

func DbDebug() bool {
	return viper.GetBool("DB_DEBUG")
}

func DbMigration() bool {
	return viper.GetBool("DB_MIGRATION")
}

func DbPgbouncer() bool {
	return viper.GetBool("DB_PGBOUNCER")
}

func DbMaxOpenConns() int {
	return viper.GetInt("DB_MAX_OPEN_CONNS")
}

func DbMaxIdleConns() int {
	return viper.GetInt("DB_MAX_IDLE_CONNS")
}

func DbMaxLifetime() int {
	return viper.GetInt("DB_MAX_LIFETIME")
}

func DbUrl() string {
	return viper.GetString("DB_URL")
}
