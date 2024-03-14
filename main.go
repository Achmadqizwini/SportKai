package main

import (
	"fmt"
	"github.com/Achmadqizwini/SportKai/config"
	"github.com/Achmadqizwini/SportKai/factory"
	"github.com/Achmadqizwini/SportKai/utils/logger"
	"github.com/Achmadqizwini/SportKai/utils/database"
	"net/http"
	"github.com/joho/godotenv"

)

func main() {
	log := logger.NewLogger().Logger.With().Str("pkg", "main").Logger()

	if err := godotenv.Load(); err != nil {
		log.Fatal().Err(err).Msg("Failed to load env file")
		panic(err)
	}
	cfg := config.GetConfig()
	db := database.InitDB(cfg)

	r := http.NewServeMux()

	factory.InitFactory(r, db)

	// Start the server
	port := fmt.Sprintf(":%d", cfg.AppConfig.AppPort)
	fmt.Printf("Server is running on port %s\n", port)
	if err:= http.ListenAndServe(port, r); err != nil {
		log.Fatal().Err(err).Msg("Failed to start server")
	}
}
