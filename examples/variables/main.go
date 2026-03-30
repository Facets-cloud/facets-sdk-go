package main

import (
	"fmt"
	"log"
	"os"

	runtime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Facets-cloud/facets-sdk-go/facets/client"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/variable_management"
	"github.com/Facets-cloud/facets-sdk-go/facets/models"
)

func main() {
	facetsClient, auth := newClient()

	stackName := "my-stack"

	// Add a variable to a stack
	addVariable(facetsClient, auth, stackName)

	// Get a variable's values across all environments
	getVariableAcrossEnvs(facetsClient, auth, stackName, "MY_VAR")
}

// addVariable adds a new variable to a stack.
func addVariable(c *client.Facets, auth runtime.ClientAuthInfoWriter, stackName string) {
	params := variable_management.NewAddVariableParams()
	params.StackName = stackName
	params.Body = &models.VariableRequest{
		// Populate with your variable details
	}

	_, err := c.VariableManagement.AddVariable(params, auth)
	if err != nil {
		log.Fatalf("Failed to add variable: %v", err)
	}

	fmt.Printf("Variable added successfully\n")
}

// getVariableAcrossEnvs retrieves a variable's values across all environments.
func getVariableAcrossEnvs(c *client.Facets, auth runtime.ClientAuthInfoWriter, stackName, varName string) {
	params := variable_management.NewGetVariableAcrossEnvironmentsParams()
	params.StackName = stackName
	params.VariableName = varName

	result, err := c.VariableManagement.GetVariableAcrossEnvironments(params, auth)
	if err != nil {
		log.Fatalf("Failed to get variable: %v", err)
	}

	resp := result.Payload
	fmt.Printf("\nVariable %q across environments:\n", varName)
	fmt.Printf("  Description:   %s\n", resp.Description)
	fmt.Printf("  Stack Default: %s\n", resp.StackDefault)
	fmt.Printf("  Is Secret:     %v\n", resp.IsSecret)
	for _, envVal := range resp.EnvironmentValues {
		fmt.Printf("  - %+v\n", envVal)
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
