package main

import (
	"fmt"
	"log"
	"os"

	runtime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Facets-cloud/facets-sdk-go/facets/client"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_deployment_controller"
)

func main() {
	facetsClient, auth := newClient()

	clusterID := "my-cluster-id"

	// List deployments for a cluster
	listDeployments(facetsClient, auth, clusterID)

	// Get the latest release for a cluster
	getLatestRelease(facetsClient, auth, clusterID)

	// Trigger a release (commented out to prevent accidental triggers)
	// triggerRelease(facetsClient, auth, clusterID)
}

// listDeployments retrieves all deployments for a cluster.
func listDeployments(c *client.Facets, auth runtime.ClientAuthInfoWriter, clusterID string) {
	params := ui_deployment_controller.NewGetDeploymentsParams()
	params.ClusterID = clusterID

	result, err := c.UIDeploymentController.GetDeployments(params, auth)
	if err != nil {
		log.Fatalf("Failed to list deployments: %v", err)
	}

	wrapper := result.Payload
	fmt.Printf("Deployments for cluster %s:\n", wrapper.ClusterID)
	for _, dep := range wrapper.Deployments {
		fmt.Printf("  - ID: %s | Status: %s | Type: %s | Created: %s\n",
			dep.ID, dep.Status, dep.ReleaseType, dep.CreatedOn)
	}
}

// getLatestRelease retrieves the most recent release for a cluster.
func getLatestRelease(c *client.Facets, auth runtime.ClientAuthInfoWriter, clusterID string) {
	params := ui_deployment_controller.NewGetLatestReleaseParams()
	params.ClusterID = clusterID

	result, err := c.UIDeploymentController.GetLatestRelease(params, auth)
	if err != nil {
		log.Fatalf("Failed to get latest release: %v", err)
	}

	release := result.Payload
	fmt.Printf("\nLatest Release:\n")
	fmt.Printf("  ID:           %s\n", release.ID)
	fmt.Printf("  Status:       %s\n", release.Status)
	fmt.Printf("  Release Type: %s\n", release.ReleaseType)
	fmt.Printf("  Created On:   %s\n", release.CreatedOn)
}

// triggerRelease triggers a new release on a cluster.
func triggerRelease(c *client.Facets, auth runtime.ClientAuthInfoWriter, clusterID string) {
	params := ui_deployment_controller.NewReleaseParams()
	params.ClusterID = clusterID

	// Optional parameters
	allowDestroy := false
	params.AllowDestroy = &allowDestroy

	withRefresh := true
	params.WithRefresh = &withRefresh

	comment := "Release triggered via SDK"
	params.Comment = &comment

	result, err := c.UIDeploymentController.Release(params, auth)
	if err != nil {
		log.Fatalf("Failed to trigger release: %v", err)
	}

	fmt.Printf("\nRelease triggered:\n")
	fmt.Printf("  ID:     %s\n", result.Payload.ID)
	fmt.Printf("  Status: %s\n", result.Payload.Status)
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
