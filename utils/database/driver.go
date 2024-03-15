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
		connStr := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s`,
			cfg.DBconfig.DB_USERNAME,
			cfg.DBconfig.DB_PASSWORD,
			cfg.DBconfig.DB_HOST,
			cfg.DBconfig.DB_PORT,
			cfg.DBconfig.DB_NAME,
		)
		db, err = sql.Open(driverMySQL, connStr)

	case driverPostgreSQL:
		connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
			cfg.DBconfig.DB_HOST,
			cfg.DBconfig.DB_USERNAME,
			cfg.DBconfig.DB_PASSWORD,
			cfg.DBconfig.DB_NAME,
			cfg.DBconfig.DB_PORT,
		)
		db, err = sql.Open("postgres", connStr)
	default:
		return nil, fmt.Errorf("unsupported database driver: %s", cfg.DBconfig.DB_DRIVER)

	}

	if err != nil {
		return nil, fmt.Errorf("error opening connection: %v", err)
	}

	errPing := db.Ping()
	if errPing != nil {
		db.Close()
		return nil, fmt.Errorf("error connecting to database: %v", errPing)
	}

	fmt.Println("Database Connection Success")
	return db, nil

}
