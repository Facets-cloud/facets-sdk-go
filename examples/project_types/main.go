package main

import (
	"fmt"
	"log"
	"os"

	runtime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Facets-cloud/facets-sdk-go/facets/client"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_project_type_controller"
)

func main() {
	facetsClient, auth := newClient()

	// List all project types
	listProjectTypes(facetsClient, auth)

	// Get a specific project type
	getProjectType(facetsClient, auth, "project-type-id")
}

// listProjectTypes retrieves all available project types.
func listProjectTypes(c *client.Facets, auth runtime.ClientAuthInfoWriter) {
	params := ui_project_type_controller.NewGetAllProjectTypesParams()
	result, err := c.UIProjectTypeController.GetAllProjectTypes(params, auth)
	if err != nil {
		log.Fatalf("Failed to list project types: %v", err)
	}

	fmt.Printf("Project Types (%d):\n", len(result.Payload))
	for _, pt := range result.Payload {
		fmt.Printf("  - %+v\n", pt)
	}
}

// getProjectType retrieves a specific project type by ID.
func getProjectType(c *client.Facets, auth runtime.ClientAuthInfoWriter, id string) {
	params := ui_project_type_controller.NewGetProjectTypeByIDParams()
	params.ID = id

	result, err := c.UIProjectTypeController.GetProjectTypeByID(params, auth)
	if err != nil {
		log.Fatalf("Failed to get project type: %v", err)
	}

	fmt.Printf("\nProject Type Details:\n")
	fmt.Printf("  %+v\n", result.Payload)
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
