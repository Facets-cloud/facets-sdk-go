package main

import (
	"fmt"
	"log"
	"os"

	runtime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Facets-cloud/facets-sdk-go/facets/client"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_stack_controller"
)

func main() {
	facetsClient, auth := newClient()

	// List all stacks
	listAllStacks(facetsClient, auth)

	// Get a specific stack by name
	getStack(facetsClient, auth, "my-stack")
}

// listAllStacks retrieves all stacks in the account.
func listAllStacks(c *client.Facets, auth runtime.ClientAuthInfoWriter) {
	params := ui_stack_controller.NewGetStacksParams()
	result, err := c.UIStackController.GetStacks(params, auth)
	if err != nil {
		log.Fatalf("Failed to list stacks: %v", err)
	}

	fmt.Printf("Found %d stacks:\n", len(result.Payload))
	for _, stack := range result.Payload {
		fmt.Printf("  - Name: %s | Cloud: %s | VCS: %s\n",
			stack.Name, stack.PrimaryCloud, stack.Vcs)
	}
}

// getStack retrieves a single stack by name.
func getStack(c *client.Facets, auth runtime.ClientAuthInfoWriter, stackName string) {
	params := ui_stack_controller.NewGetStackParams()
	params.StackName = stackName

	result, err := c.UIStackController.GetStack(params, auth)
	if err != nil {
		log.Fatalf("Failed to get stack %q: %v", stackName, err)
	}

	stack := result.Payload
	fmt.Printf("\nStack Details:\n")
	fmt.Printf("  Name:        %s\n", stack.Name)
	fmt.Printf("  ID:          %s\n", stack.ID)
	fmt.Printf("  Cloud:       %s\n", stack.PrimaryCloud)
	fmt.Printf("  VCS:         %s\n", stack.Vcs)
	fmt.Printf("  VCS URL:     %s\n", stack.VcsURL)
	fmt.Printf("  Description: %s\n", stack.Description)
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
