package main

import (
	"log"

	_ "github.com/NayanPahuja/fam-bcknd-test/cmd/docs"
	"github.com/NayanPahuja/fam-bcknd-test/config"
	"github.com/NayanPahuja/fam-bcknd-test/db"
	"github.com/NayanPahuja/fam-bcknd-test/internal/routes"
	"github.com/NayanPahuja/fam-bcknd-test/internal/temporal"
	"github.com/NayanPahuja/fam-bcknd-test/internal/workflows"
	"github.com/gofiber/fiber/v2"
	fiberSwagger "github.com/swaggo/fiber-swagger"
)

func main() {
	// Initialize Config
	cfg := config.Envs

	// Initialize Database
	db.InitDB()

	// Initialize Fiber App
	app := fiber.New()
	routes.RegisterRoutes(app)
	app.Get("/swagger/*", fiberSwagger.WrapHandler)
	// Initialize Temporal Client and Worker
	temporalClient := temporal.Init(cfg.TemporalHost)
	defer temporalClient.Close()

	// Start Temporal Worker
	go temporal.StartWorker(temporalClient)

	go func() {
		log.Println("Triggering Temporal Workflow...")
		input := workflows.YouTubeWorkflowInput{
			SearchQuery: "Happy New Year",
		}
		temporal.TriggerWorkflow(temporalClient, input)
	}()

	// Start HTTP Server
	port := ":8080"
	log.Printf("Server running on http://localhost%s", port)
	log.Println(app.Stack())
	log.Fatal(app.Listen(port))
}
