package app

import (
	"avitotest/config"
	"avitotest/internal/database"
	"avitotest/internal/repo"
	"avitotest/internal/service"
	"avitotest/internal/transport/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func Run(configPath string) {

	// Configuration
	cfg, err := config.NewConfig(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	base, err := database.Connect(cfg)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	// Repositories
	repositories := repo.NewRepositories(base)

	// Services dependencies
	deps := service.ServicesDependencies{
		Repos: repositories,
	}
	services := service.NewServices(deps)

	// Handlers
	handlers := handlers.NewHandler(services)
	handlers.DinamicSegmentRoutes(r)

	http.Handle("/", r)
	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Fatal(server.ListenAndServe())
}
