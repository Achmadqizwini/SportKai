package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Achmadqizwini/SportKai/config"
	_ "github.com/go-sql-driver/mysql"
)

const (
	driverMySQL      = "mysql"
	driverPostgreSQL = "postgres"
)

func InitDB(cfg *config.Config) *sql.DB {

	switch cfg.DBconfig.DB_DRIVER {
	case driverMySQL:
		dbConf := fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s`,
			cfg.DBconfig.DB_USERNAME,
			cfg.DBconfig.DB_PASSWORD,
			cfg.DBconfig.DB_HOST,
			cfg.DBconfig.DB_PORT,
			cfg.DBconfig.DB_NAME,
		)
		dbConn, err := sql.Open(driverMySQL, dbConf)

		if err != nil {
			log.Fatal("Error open connection", err.Error())
			return nil
		}
		errPing := dbConn.Ping()
		if errPing != nil {
			log.Fatal("Error connect to db", errPing.Error())
		} else {
			fmt.Println("Database Connection Success ")
		}
		return dbConn
	}

	// case driverPostgreSQL:
	// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	// 		cfg.DBconfig.DBhost,
	// 		cfg.DBconfig.DBuser,
	// 		cfg.DBconfig.DBpassword,
	// 		cfg.DBconfig.DBname,
	// 		cfg.DBconfig.DBport,
	// 	)

	// 	dbConn, err := gorm.Open(postgres.Open(dsn), gormConfig)
	// 	if err != nil {
	// 		log.Error().Err(err).Str("dsn", dsn).Msg("failed to connect to database")

	// 		return nil, err
	// 	}

	// 	return dbConn, nil
	// }
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True", cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOST, cfg.DB_PORT, cfg.DB_NAME)
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatal("Cannot connect to DB")
	// }

	// migrateDB(db)

	return nil

}
