package temporal

import (
	"context"
	"log"

	"github.com/NayanPahuja/fam-bcknd-test/internal/workflows"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

// StartWorker starts the Temporal worker.
func StartWorker(temporalClient client.Client) {
	// Create a new worker
	w := worker.New(temporalClient, "YouTubeTaskQueue", worker.Options{})

	// Register the workflow and activities with the worker
	w.RegisterWorkflow(workflows.YouTubeFetchWorkflow)
	w.RegisterActivity(workflows.YouTubeActivity)

	// Start the worker
	log.Println("Starting Temporal Worker...")
	if err := w.Run(worker.InterruptCh()); err != nil {
		log.Fatalf("Worker failed to start: %v", err)
	}
	log.Println("Temporal Worker stopped.")
}

// TriggerWorkflow triggers the Temporal workflow.
func TriggerWorkflow(c client.Client, input workflows.YouTubeWorkflowInput) {
	workflowOptions := client.StartWorkflowOptions{
		ID:        "youtube_workflow",
		TaskQueue: "YouTubeTaskQueue",
	}

	_, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflows.YouTubeFetchWorkflow, input)
	if err != nil {
		log.Fatalf("Unable to start workflow: %v", err)
	}
	log.Println("Workflow started successfully!")
}
