package main

import (
	"fmt"
	"github.com/Achmadqizwini/SportKai/config"
	"github.com/Achmadqizwini/SportKai/factory"
	"github.com/Achmadqizwini/SportKai/utils/database"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	cfg := config.GetConfig()
	db := database.InitDB(cfg)

	r := mux.NewRouter()

	factory.InitFactory(r, db)

	// Start the server
	port := fmt.Sprintf(":%d", cfg.AppConfig.AppPort)
	fmt.Printf("Server is running on port %s\n", port)
	http.ListenAndServe(port, r)
}
