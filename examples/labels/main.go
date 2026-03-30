package main

import (
	"fmt"
	"log"
	"os"

	runtime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Facets-cloud/facets-sdk-go/facets/client"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/label_management"
	"github.com/Facets-cloud/facets-sdk-go/facets/models"
)

func main() {
	facetsClient, auth := newClient()

	// List all labels
	listLabels(facetsClient, auth)

	// Create a new label
	createLabel(facetsClient, auth, "production", "#FF0000")

	// Get a label by ID
	getLabel(facetsClient, auth, "label-id-here")
}

// listLabels retrieves all labels in the organization.
func listLabels(c *client.Facets, auth runtime.ClientAuthInfoWriter) {
	params := label_management.NewGetAllLabelsParams()
	result, err := c.LabelManagement.GetAllLabels(params, auth)
	if err != nil {
		log.Fatalf("Failed to list labels: %v", err)
	}

	fmt.Printf("Labels (%d):\n", len(result.Payload))
	for _, label := range result.Payload {
		fmt.Printf("  - Name: %-20s | Color: %s | ID: %s\n",
			label.Name, label.ColorCode, label.ID)
	}
}

// createLabel creates a new label with the given name and color.
func createLabel(c *client.Facets, auth runtime.ClientAuthInfoWriter, name, color string) {
	params := label_management.NewCreateLabelParams()
	params.Body = &models.LabelRequest{
		Name:      &name,
		ColorCode: color,
	}

	result, err := c.LabelManagement.CreateLabel(params, auth)
	if err != nil {
		log.Fatalf("Failed to create label: %v", err)
	}

	fmt.Printf("\nCreated Label:\n")
	fmt.Printf("  Name:  %s\n", result.Payload.Name)
	fmt.Printf("  Color: %s\n", result.Payload.ColorCode)
	fmt.Printf("  ID:    %s\n", result.Payload.ID)
}

// getLabel retrieves a label by its ID.
func getLabel(c *client.Facets, auth runtime.ClientAuthInfoWriter, labelID string) {
	params := label_management.NewGetLabelByIDParams()
	params.LabelID = labelID

	result, err := c.LabelManagement.GetLabelByID(params, auth)
	if err != nil {
		log.Fatalf("Failed to get label: %v", err)
	}

	fmt.Printf("\nLabel Details:\n")
	fmt.Printf("  Name:       %s\n", result.Payload.Name)
	fmt.Printf("  Color:      %s\n", result.Payload.ColorCode)
	fmt.Printf("  Created By: %s\n", result.Payload.CreatedBy)
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
