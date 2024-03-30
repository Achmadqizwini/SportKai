package database

import (
	"database/sql"
	"fmt"

	"github.com/Achmadqizwini/SportKai/config"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

const (
	driverMySQL      = "mysql"
	driverPostgreSQL = "postgres"
)

func InitDB(cfg *config.Config) (*sql.DB, error) {
	var db *sql.DB
	var err error

	switch cfg.DBconfig.DB_DRIVER {
	case driverMySQL:
		connStr := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True`,
			cfg.DBconfig.DB_USERNAME,
			cfg.DBconfig.DB_PASSWORD,
			cfg.DBconfig.DB_HOST,
			cfg.DBconfig.DB_PORT,
			cfg.DBconfig.DB_NAME,
		)
		db, err = sql.Open(driverMySQL, connStr)

	case driverPostgreSQL:
		connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cfg.DBconfig.DB_HOST,
			cfg.DBconfig.DB_PORT,
			cfg.DBconfig.DB_USERNAME,
			cfg.DBconfig.DB_PASSWORD,
			cfg.DBconfig.DB_NAME,
		)

		db, err = sql.Open("postgres", connStr)
	default:
		return nil, err

	}

	if err != nil {
		return nil, err
	}

	errPing := db.Ping()
	if errPing != nil {
		db.Close()
		return nil, errPing
	}
	return db, nil
}
