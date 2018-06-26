package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/viper"
)

func GetFromConfig(cfg *viper.Viper) (*sql.DB, error) {
	return sql.Open(`sqlite3`, cfg.GetString(`DATABASE.PATH`))
}
