package main

import (
	"log"

	"github.com/NayanPahuja/fam-bcknd-test/config"
	"github.com/NayanPahuja/fam-bcknd-test/db"
	"github.com/NayanPahuja/fam-bcknd-test/internal/routes"
	"github.com/NayanPahuja/fam-bcknd-test/internal/temporal"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize Config
	cfg := config.Envs

	// Initialize Database
	db.InitDB()

	// Initialize Fiber App
	app := fiber.New()
	routes.RegisterRoutes(app)

	// Initialize Temporal Client and Worker
	temporalClient := temporal.Init(cfg.TemporalHost)
	defer temporalClient.Close()

	// Start Temporal Worker
	go temporal.StartWorker(temporalClient)

	// Start HTTP Server
	port := ":3000"
	log.Printf("Server running on http://localhost%s", port)
	log.Fatal(app.Listen(port))
}
