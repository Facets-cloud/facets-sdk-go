package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	// These imports will need to be updated based on your actual module name
	// when you publish this SDK to a repository
	"github.com/Facets-cloud/facets-sdk-go/facets/client"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_user_controller"
)

func main() {
	// Get API host and token from environment variables or configuration
	host := getEnvOrDefault("FACETS_API_HOST", "")
	token := getEnvOrDefault("FACETS_API_TOKEN", "")
	username := getEnvOrDefault("FACETS_USERNAME", "")

	if token == "" {
		log.Fatalf("API token is required. Set FACETS_API_TOKEN environment variable.")
	}
	if username == "" {
		log.Fatalf("API username is required. Set FACETS_USERNAME environment variable.")
	}
	if host == "" {
		log.Fatalf("API host is required. Set FACETS_API_HOST environment variable.")
	}

	// Create the transport
	transport := httptransport.New(host, "/", []string{"https"})

	// Set up authentication using Bearer token
	auth := httptransport.BasicAuth(username, token)

	// Create the Facets client
	facetsClient := client.New(transport, strfmt.Default)

	// Example: Get current user information
	params := ui_user_controller.NewGetCurrentUserUsingGETParams()
	result, err := facetsClient.UIUserController.GetCurrentUserUsingGET(params, auth)

	if err != nil {
		if apiErr, ok := err.(*runtime.APIError); ok {
			// Handle API-level error
			fmt.Printf("API error (code %d): %s\n", apiErr.Code, apiErr.Error())
		} else {
			// Handle other errors
			fmt.Printf("Error: %v\n", err)
		}
		os.Exit(1)
	}

	// Process successful response
	user := result.Payload
	fmt.Printf("Current user information:\n")
	fmt.Printf("  ID: %s\n", user.ID)
}

// Helper function to get environment variable with default fallback
func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
