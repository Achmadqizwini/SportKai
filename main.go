package main

import (
	"fmt"
	"github.com/Achmadqizwini/SportKai/config"
	"github.com/Achmadqizwini/SportKai/factory"
	"github.com/Achmadqizwini/SportKai/middlewares"
	"github.com/Achmadqizwini/SportKai/utils/database"
	"github.com/Achmadqizwini/SportKai/utils/logger"
	"github.com/joho/godotenv"
	// "github.com/rs/cors"
	"net/http"
)

func main() {
	// log := logger.NewLogger().Logger.With().Str("pkg", "main").Logger()
	log := logger.NewLogger().Logger

	if err := godotenv.Load(); err != nil {
		log.Fatal().Err(err).Msg("Failed to load env file")
		panic(err)
	}
	cfg := config.GetConfig()
	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
		panic(err)
	} else {
		log.Info().Msg("Database Connection Success")
	}

	mux := http.NewServeMux()

	handler := middlewares.CollectMiddleware(
		middlewares.Logging,
		middlewares.Cors,
	)

	factory.InitFactory(mux, db)

	// Start the server
	port := fmt.Sprintf(":%d", cfg.AppConfig.AppPort)
	server := http.Server{
		Addr:    port,
		Handler: handler(mux),
	}

	log.Info().Msgf("Server is running on port %s\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
