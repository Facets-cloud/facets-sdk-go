package main

import (
	"fmt"
	"log"
	"os"

	runtime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Facets-cloud/facets-sdk-go/facets/client"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_blueprint_designer_controller"
)

func main() {
	facetsClient, auth := newClient()

	stackName := "my-stack"
	branchName := "master"

	// List all resources in the blueprint
	listDesignerResources(facetsClient, auth, stackName, branchName)

	// List branches for a stack
	listBranches(facetsClient, auth, stackName)

	// Get autocomplete data for blueprint editing
	getAutocompleteData(facetsClient, auth, stackName)
}

// listDesignerResources retrieves all resources defined in a blueprint.
func listDesignerResources(c *client.Facets, auth runtime.ClientAuthInfoWriter, stackName, branchName string) {
	params := ui_blueprint_designer_controller.NewGetDesignerResourcesParams()
	params.StackName = stackName
	params.BranchName = branchName

	result, err := c.UIBlueprintDesignerController.GetDesignerResources(params, auth)
	if err != nil {
		log.Fatalf("Failed to list resources: %v", err)
	}

	fmt.Printf("Blueprint Resources for %s/%s (%d):\n", stackName, branchName, len(result.Payload))
	for _, res := range result.Payload {
		fmt.Printf("  - %+v\n", res)
	}
}

// listBranches retrieves all branches for a stack's blueprint.
func listBranches(c *client.Facets, auth runtime.ClientAuthInfoWriter, stackName string) {
	params := ui_blueprint_designer_controller.NewListBranchesParams()
	params.StackName = stackName

	result, err := c.UIBlueprintDesignerController.ListBranches(params, auth)
	if err != nil {
		log.Fatalf("Failed to list branches: %v", err)
	}

	fmt.Printf("\nBranches for %s (%d):\n", stackName, len(result.Payload))
	for _, branch := range result.Payload {
		fmt.Printf("  - %+v\n", branch)
	}
}

// getAutocompleteData retrieves autocomplete suggestions for blueprint editing.
func getAutocompleteData(c *client.Facets, auth runtime.ClientAuthInfoWriter, stackName string) {
	params := ui_blueprint_designer_controller.NewGetAutocompleteDataParams()
	params.StackName = stackName

	result, err := c.UIBlueprintDesignerController.GetAutocompleteData(params, auth)
	if err != nil {
		log.Fatalf("Failed to get autocomplete data: %v", err)
	}

	fmt.Printf("\nAutocomplete data retrieved: %+v\n", result.Payload)
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
