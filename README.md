# Facets Go SDK

This is a Go client SDK for the Facets API, generated from the Swagger specification.

## Overview

This SDK provides a Go client for interacting with the Facets API. It was generated using the go-swagger tool from a Swagger 2.0 specification.

## Installation

To install the SDK in your Go project:

```bash
go get github.com/Facets-cloud/facets-sdk-go
```

## Usage

Here's a basic example of how to use the SDK:

```go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Facets-cloud/facets-sdk-go/facets/client"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_user_controller"
)

func main() {
	// Get API host and token from environment variables or configuration
	host := getEnvOrDefault("FACETS_API_HOST", "")
	token := getEnvOrDefault("FACETS_API_TOKEN", "")
    username := getEnvOrDefault("FACETS_USERNAME", "")

    if host == "" {
		log.Fatalf("API token is required. Set FACETS_API_HOST environment variable.")
	}
	if token == "" {
		log.Fatalf("API token is required. Set FACETS_API_TOKEN environment variable.")
	}
    if username == "" {
		log.Fatalf("API token is required. Set FACETS_USERNAME environment variable.")
	}

	// Create the transport
	transport := httptransport.New(host, "/", []string{"https"})

	// Set up authentication using Basic Auth
	auth := auth := httptransport.BasicAuth(username, token)


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
```

## Available Controllers

The SDK provides clients for various controllers in the Facets API, including:

- UIUserController - User management
- UIAccountsController - Account management
- UIArtifactsController - Artifact management
- UIDeploymentController - Deployment management
- UIStackController - Stack management
- UITeamController - Team management

And many more for cloud providers (AWS, GCP, Azure), Kubernetes, CI/CD, and other functionalities.

## Authentication

The SDK supports different authentication methods:

```go
// API Key authentication (Bearer token)
auth := httptransport.APIKeyAuth("Authorization", "header", "Bearer YOUR_API_TOKEN")

// Basic authentication
auth := httptransport.BasicAuth("username", "password")
```

## Models

The SDK includes Go struct representations for all API models in the `facets/models` package, providing type-safe interaction with the API.

## Custom Configuration

You can customize the client with different configurations:

```go
// Create a client with custom timeout
httpClient := &http.Client{
    Timeout: 30 * time.Second,
}
transport := httptransport.NewWithClient(
    host, 
    basePath, 
    []string{"https"}, 
    httpClient,
)

// Create a transport with custom configuration
cfg := client.DefaultTransportConfig().
    WithHost("custom-host.facets.cloud").
    WithBasePath("/api/v1").
    WithSchemes([]string{"https"})
transportCfg := client.NewHTTPClientWithConfig(nil, cfg)
```

## Examples

See the `examples/` directory for additional usage examples.

## License

[License information] 