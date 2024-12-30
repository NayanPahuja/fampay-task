package temporal

import (
	"log"

	"go.temporal.io/sdk/client"
)

// Init initializes the Temporal client.
func Init(temporalHost string) client.Client {
	// Connect to Temporal
	temporalClient, err := client.Dial(client.Options{
		HostPort: temporalHost,
	})
	if err != nil {
		log.Fatalf("Unable to connect to Temporal: %v", err)
	}

	log.Println("Successfully connected to Temporal!")
	return temporalClient
}
