package main

import (
	"fmt"
	"log"
	"os"

	runtime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Facets-cloud/facets-sdk-go/facets/client"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_common_cluster_controller"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_stack_controller"
)

func main() {
	facetsClient, auth := newClient()

	// List all clusters across all stacks
	listAllClusters(facetsClient, auth)

	// Get cluster details by cluster ID
	getClusterDetails(facetsClient, auth, "my-cluster-id")

	// Get cluster info by stack name and cluster name
	getClusterInfoByName(facetsClient, auth, "my-stack", "my-cluster")

	// Get cluster variables
	getClusterVariables(facetsClient, auth, "my-cluster-id")
}

// listAllClusters retrieves all clusters across all stacks (paginated).
func listAllClusters(c *client.Facets, auth runtime.ClientAuthInfoWriter) {
	params := ui_stack_controller.NewGetAllClustersParams()
	result, err := c.UIStackController.GetAllClusters(params, auth)
	if err != nil {
		log.Fatalf("Failed to list clusters: %v", err)
	}

	page := result.Payload
	fmt.Printf("All Clusters (page %d, %d total):\n", page.Number, page.TotalElements)
	for _, cluster := range page.Content {
		fmt.Printf("  - ID: %s | Name: %s | Stack: %s | Cloud: %s\n",
			cluster.ID, stringVal(cluster.Name), stringVal(cluster.StackName), cluster.Cloud)
	}
}

// getClusterDetails retrieves detailed information about a specific cluster.
func getClusterDetails(c *client.Facets, auth runtime.ClientAuthInfoWriter, clusterID string) {
	params := ui_common_cluster_controller.NewGetClusterCommonParams()
	params.ClusterID = clusterID

	result, err := c.UICommonClusterController.GetClusterCommon(params, auth)
	if err != nil {
		log.Fatalf("Failed to get cluster details: %v", err)
	}

	fmt.Printf("\nCluster Details:\n")
	fmt.Printf("  %+v\n", result.Payload)
}

// getClusterInfoByName looks up a cluster by stack name and cluster name.
func getClusterInfoByName(c *client.Facets, auth runtime.ClientAuthInfoWriter, stackName, clusterName string) {
	params := ui_common_cluster_controller.NewGetClusterInfoByNameParams()
	params.StackName = stackName
	params.ClusterName = clusterName

	result, err := c.UICommonClusterController.GetClusterInfoByName(params, auth)
	if err != nil {
		log.Fatalf("Failed to get cluster info: %v", err)
	}

	info := result.Payload
	fmt.Printf("\nCluster Info:\n")
	fmt.Printf("  ID:    %s\n", info.ID)
	fmt.Printf("  Name:  %s\n", info.Name)
	fmt.Printf("  Cloud: %s\n", info.Cloud)
	fmt.Printf("  State: %s\n", info.ClusterState)
}

// getClusterVariables retrieves variables configured for a cluster.
func getClusterVariables(c *client.Facets, auth runtime.ClientAuthInfoWriter, clusterID string) {
	params := ui_common_cluster_controller.NewGetVarsParams()
	params.ClusterID = clusterID

	result, err := c.UICommonClusterController.GetVars(params, auth)
	if err != nil {
		log.Fatalf("Failed to get cluster variables: %v", err)
	}

	cluster := result.Payload
	fmt.Printf("\nCluster Variables for %s:\n", stringVal(cluster.Name))
	for key, val := range cluster.Variables {
		fmt.Printf("  %s = %+v\n", key, val)
	}
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
