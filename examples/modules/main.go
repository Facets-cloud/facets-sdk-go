package main

import (
	"fmt"
	"log"
	"os"

	runtime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Facets-cloud/facets-sdk-go/facets/client"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/module_management"
)

func main() {
	facetsClient, auth := newClient()

	// List all modules
	listModules(facetsClient, auth)

	// Get a module by ID
	getModuleByID(facetsClient, auth, "module-id-here")

	// Get modules grouped by stack
	getGroupedModules(facetsClient, auth, "my-stack")
}

// listModules retrieves all available modules.
func listModules(c *client.Facets, auth runtime.ClientAuthInfoWriter) {
	params := module_management.NewGetAllModulesParams()

	result, err := c.ModuleManagement.GetAllModules(params, auth)
	if err != nil {
		log.Fatalf("Failed to list modules: %v", err)
	}

	fmt.Printf("Modules (%d):\n", len(result.Payload))
	for _, mod := range result.Payload {
		fmt.Printf("  - Type: %-20s | Flavor: %-15s | Clouds: %v | Created By: %s\n",
			mod.Type, mod.Flavor, mod.Clouds, mod.CreatedBy)
	}
}

// getModuleByID retrieves a specific module by its ID.
func getModuleByID(c *client.Facets, auth runtime.ClientAuthInfoWriter, moduleID string) {
	params := module_management.NewGetByIDParams()
	params.ID = moduleID

	result, err := c.ModuleManagement.GetByID(params, auth)
	if err != nil {
		log.Fatalf("Failed to get module: %v", err)
	}

	fmt.Printf("\nModule Details:\n")
	fmt.Printf("  %+v\n", result.Payload)
}

// getGroupedModules retrieves modules grouped by intent for a stack.
func getGroupedModules(c *client.Facets, auth runtime.ClientAuthInfoWriter, stackName string) {
	params := module_management.NewGetGroupedModulesForStackParams()
	params.StackName = stackName

	result, err := c.ModuleManagement.GetGroupedModulesForStack(params, auth)
	if err != nil {
		log.Fatalf("Failed to get grouped modules: %v", err)
	}

	resp := result.Payload
	fmt.Printf("\nGrouped Modules for %s:\n", stackName)
	for resourceType, flavors := range resp.Resources {
		fmt.Printf("  Resource Type: %s\n", resourceType)
		for flavor, modules := range flavors {
			fmt.Printf("    Flavor: %s → %+v\n", flavor, modules)
		}
	}
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
