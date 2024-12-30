package workflows

import (
	"time"

	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type YouTubeWorkflowInput struct {
	SearchQuery string
}

func YouTubeFetchWorkflow(ctx workflow.Context, input YouTubeWorkflowInput) error {
	logger := workflow.GetLogger(ctx)
	logger.Info("Starting YouTube Workflow", "SearchQuery", input.SearchQuery)

	options := workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second * 5,
			BackoffCoefficient: 2.0,
			MaximumAttempts:    3,
		},
	}
	ctx = workflow.WithActivityOptions(ctx, options)

	// Repeatedly fetch videos in a loop
	for {
		err := workflow.ExecuteActivity(ctx, YouTubeActivity, input.SearchQuery).Get(ctx, nil)
		if err != nil {
			workflow.GetLogger(ctx).Error("Failed to fetch videos", "error", err)
		}

		// Sleep for the defined interval (e.g., 10 seconds)
		workflow.Sleep(ctx, time.Second*10)
		logger.Info("Activity completed successfully!")
	}

}
