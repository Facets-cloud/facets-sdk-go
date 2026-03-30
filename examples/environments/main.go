package main

import (
	"fmt"
	"log"
	"os"

	runtime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Facets-cloud/facets-sdk-go/facets/client"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_environment_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/models"
)

func main() {
	facetsClient, auth := newClient()

	// Get environment details by cluster ID
	getEnvironment(facetsClient, auth, "my-cluster-id")

	// Create a new environment
	createEnvironment(facetsClient, auth)
}

// getEnvironment retrieves environment details for a given cluster ID.
func getEnvironment(c *client.Facets, auth runtime.ClientAuthInfoWriter, clusterID string) {
	params := ui_environment_controller.NewGetEnvironmentParams()
	params.ClusterID = clusterID

	result, err := c.UIEnvironmentController.GetEnvironment(params, auth)
	if err != nil {
		log.Fatalf("Failed to get environment: %v", err)
	}

	env := result.Payload
	fmt.Printf("Environment Details:\n")
	fmt.Printf("  ID:             %s\n", env.ID)
	fmt.Printf("  Name:           %s\n", stringVal(env.Name))
	fmt.Printf("  Stack:          %s\n", stringVal(env.StackName))
	fmt.Printf("  Release Stream: %s\n", stringVal(env.ReleaseStream))
	fmt.Printf("  Cloud:          %s\n", env.Cloud)
	fmt.Printf("  Cluster State:  %s\n", env.ClusterState)
}

// createEnvironment creates a new environment.
func createEnvironment(c *client.Facets, auth runtime.ClientAuthInfoWriter) {
	params := ui_environment_controller.NewCreateEnvironmentParams()
	params.Body = &models.EnvironmentRequest{
		ClusterName:   "staging",
		StackName:     "my-stack",
		ReleaseStream: "default",
		Cloud:         "AWS",
	}

	result, err := c.UIEnvironmentController.CreateEnvironment(params, auth)
	if err != nil {
		log.Fatalf("Failed to create environment: %v", err)
	}

	env := result.Payload
	fmt.Printf("\nCreated environment: %s (ID: %s)\n",
		stringVal(env.Name), env.ID)
}

func stringVal(s *string) string {
	if s == nil {
		return "<nil>"
	}
	return *s
}

func newClient() (*client.Facets, runtime.ClientAuthInfoWriter) {
	host := os.Getenv("FACETS_API_HOST")
	username := os.Getenv("FACETS_USERNAME")
	token := os.Getenv("FACETS_API_TOKEN")

	if host == "" || username == "" || token == "" {
		log.Fatal("Set FACETS_API_HOST, FACETS_USERNAME, and FACETS_API_TOKEN environment variables")
	}

	transport := httptransport.New(host, "/", []string{"https"})
	auth := httptransport.BasicAuth(username, token)
	return client.New(transport, strfmt.Default), auth
}
