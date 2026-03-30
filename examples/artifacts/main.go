package main

import (
	"fmt"
	"log"
	"os"

	runtime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Facets-cloud/facets-sdk-go/facets/client"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_artifact_ci_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_artifacts_controller"
)

func main() {
	facetsClient, auth := newClient()

	// List all artifact CI configurations
	listArtifactCIs(facetsClient, auth)

	// Get artifact CI by name
	getArtifactCIByName(facetsClient, auth, "my-service-ci")

	// Get artifact details for an application in a cluster
	getArtifactByApp(facetsClient, auth, "my-cluster-id", "my-app")

	// Register a new artifact
	registerArtifact(facetsClient, auth)
}

// listArtifactCIs retrieves all artifact CI configurations.
func listArtifactCIs(c *client.Facets, auth runtime.ClientAuthInfoWriter) {
	params := ui_artifact_ci_controller.NewGetAllArtifactsCIParams()
	result, err := c.UIArtifactCiController.GetAllArtifactsCI(params, auth)
	if err != nil {
		log.Fatalf("Failed to list artifact CIs: %v", err)
	}

	fmt.Printf("Artifact CIs (%d):\n", len(result.Payload))
	for _, ci := range result.Payload {
		fmt.Printf("  - CI Name: %-25s | Type: %-15s | Stack: %s\n",
			stringVal(ci.CiName), stringVal(ci.RegistrationType), ci.StackName)
	}
}

// getArtifactCIByName retrieves a specific artifact CI configuration.
func getArtifactCIByName(c *client.Facets, auth runtime.ClientAuthInfoWriter, ciName string) {
	params := ui_artifact_ci_controller.NewGetArtifactCiByNameParams()
	params.CiName = ciName

	result, err := c.UIArtifactCiController.GetArtifactCiByName(params, auth)
	if err != nil {
		log.Fatalf("Failed to get artifact CI %q: %v", ciName, err)
	}

	ci := result.Payload
	fmt.Printf("\nArtifact CI Details:\n")
	fmt.Printf("  ID:                %s\n", ci.ID)
	fmt.Printf("  CI Name:           %s\n", stringVal(ci.CiName))
	fmt.Printf("  Registration Type: %s\n", stringVal(ci.RegistrationType))
	fmt.Printf("  Stack:             %s\n", ci.StackName)
	fmt.Printf("  Created By:        %s\n", ci.CreatedBy)
}

// getArtifactByApp retrieves artifact details for an application in a cluster.
func getArtifactByApp(c *client.Facets, auth runtime.ClientAuthInfoWriter, clusterID, appName string) {
	params := ui_artifacts_controller.NewGetArtifactByApplicationNameParams()
	params.ClusterID = clusterID
	params.ApplicationName = appName

	result, err := c.UIArtifactsController.GetArtifactByApplicationName(params, auth)
	if err != nil {
		log.Fatalf("Failed to get artifact for app %q: %v", appName, err)
	}

	fmt.Printf("\nArtifact for %s:\n", appName)
	fmt.Printf("  %+v\n", result.Payload)
}

// registerArtifact registers a new artifact.
func registerArtifact(c *client.Facets, auth runtime.ClientAuthInfoWriter) {
	params := ui_artifacts_controller.NewRegisterArtifactParams()
	// Populate params.Body with your artifact details:
	// params.Body = &models.Artifact{...}

	result, err := c.UIArtifactsController.RegisterArtifact(params, auth)
	if err != nil {
		log.Fatalf("Failed to register artifact: %v", err)
	}

	fmt.Printf("\nRegistered Artifact:\n")
	fmt.Printf("  %+v\n", result.Payload)
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
