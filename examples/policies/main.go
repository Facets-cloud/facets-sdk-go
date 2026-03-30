package main

import (
	"fmt"
	"log"
	"os"

	runtime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Facets-cloud/facets-sdk-go/facets/client"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_opa_controller"
)

func main() {
	facetsClient, auth := newClient()

	// List all OPA policies
	listPolicies(facetsClient, auth)

	// Get a specific policy
	getPolicy(facetsClient, auth, "my-policy-name")

	// List all policy templates
	listPolicyTemplates(facetsClient, auth)
}

// listPolicies retrieves all OPA policies.
func listPolicies(c *client.Facets, auth runtime.ClientAuthInfoWriter) {
	params := ui_opa_controller.NewGetAllPoliciesParams()
	result, err := c.UIOpaController.GetAllPolicies(params, auth)
	if err != nil {
		log.Fatalf("Failed to list policies: %v", err)
	}

	fmt.Printf("OPA Policies (%d):\n", len(result.Payload))
	for _, policy := range result.Payload {
		fmt.Printf("  - %+v\n", policy)
	}
}

// getPolicy retrieves a specific OPA policy by name.
func getPolicy(c *client.Facets, auth runtime.ClientAuthInfoWriter, name string) {
	params := ui_opa_controller.NewGetPolicyParams()
	params.PolicyName = name

	result, err := c.UIOpaController.GetPolicy(params, auth)
	if err != nil {
		log.Fatalf("Failed to get policy: %v", err)
	}

	fmt.Printf("\nPolicy Details:\n")
	fmt.Printf("  %+v\n", result.Payload)
}

// listPolicyTemplates retrieves all available policy templates.
func listPolicyTemplates(c *client.Facets, auth runtime.ClientAuthInfoWriter) {
	params := ui_opa_controller.NewGetAllPolicyTemplatesParams()
	result, err := c.UIOpaController.GetAllPolicyTemplates(params, auth)
	if err != nil {
		log.Fatalf("Failed to list policy templates: %v", err)
	}

	fmt.Printf("\nPolicy Templates (%d):\n", len(result.Payload))
	for _, tmpl := range result.Payload {
		fmt.Printf("  - %+v\n", tmpl)
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
