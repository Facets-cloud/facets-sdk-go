package main

import (
	"fmt"
	"log"
	"os"

	runtime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Facets-cloud/facets-sdk-go/facets/client"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_user_controller"
)

func main() {
	facetsClient, auth := newClient()

	// Get current authenticated user
	getCurrentUser(facetsClient, auth)

	// List all users
	listAllUsers(facetsClient, auth)
}

// getCurrentUser retrieves the currently authenticated user's details.
func getCurrentUser(c *client.Facets, auth runtime.ClientAuthInfoWriter) {
	params := ui_user_controller.NewGetCurrentUserParams()
	result, err := c.UIUserController.GetCurrentUser(params, auth)
	if err != nil {
		log.Fatalf("Failed to get current user: %v", err)
	}

	user := result.Payload
	fmt.Printf("Current User:\n")
	fmt.Printf("  ID:       %s\n", user.ID)
	fmt.Printf("  Username: %s\n", user.UserName)
	fmt.Printf("  Roles:    %v\n", user.Roles)
}

// listAllUsers retrieves all users in the organization.
func listAllUsers(c *client.Facets, auth runtime.ClientAuthInfoWriter) {
	params := ui_user_controller.NewGetAllUsersParams()
	result, err := c.UIUserController.GetAllUsers(params, auth)
	if err != nil {
		log.Fatalf("Failed to list users: %v", err)
	}

	fmt.Printf("\nAll Users (%d):\n", len(result.Payload))
	for _, user := range result.Payload {
		fmt.Printf("  - %s (ID: %s) Roles: %v\n",
			user.UserName, user.ID, user.Roles)
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
