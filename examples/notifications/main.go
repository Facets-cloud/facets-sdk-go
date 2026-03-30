package main

import (
	"fmt"
	"log"
	"os"

	runtime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Facets-cloud/facets-sdk-go/facets/client"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_notification_controller"
)

func main() {
	facetsClient, auth := newClient()

	// List all notification channels
	listChannels(facetsClient, auth)

	// List all notification subscriptions
	listSubscriptions(facetsClient, auth)

	// List available channel types
	listChannelTypes(facetsClient, auth)
}

// listChannels retrieves all configured notification channels.
func listChannels(c *client.Facets, auth runtime.ClientAuthInfoWriter) {
	params := ui_notification_controller.NewGetAllChannelsParams()
	result, err := c.UINotificationController.GetAllChannels(params, auth)
	if err != nil {
		log.Fatalf("Failed to list channels: %v", err)
	}

	fmt.Printf("Notification Channels (%d):\n", len(result.Payload))
	for _, ch := range result.Payload {
		fmt.Printf("  - %+v\n", ch)
	}
}

// listSubscriptions retrieves all notification subscriptions.
func listSubscriptions(c *client.Facets, auth runtime.ClientAuthInfoWriter) {
	params := ui_notification_controller.NewGetAllSubscriptions1Params()
	result, err := c.UINotificationController.GetAllSubscriptions1(params, auth)
	if err != nil {
		log.Fatalf("Failed to list subscriptions: %v", err)
	}

	fmt.Printf("\nNotification Subscriptions (%d):\n", len(result.Payload))
	for _, sub := range result.Payload {
		fmt.Printf("  - %+v\n", sub)
	}
}

// listChannelTypes retrieves all available notification channel types.
func listChannelTypes(c *client.Facets, auth runtime.ClientAuthInfoWriter) {
	params := ui_notification_controller.NewGetAllChannelTypesParams()
	result, err := c.UINotificationController.GetAllChannelTypes(params, auth)
	if err != nil {
		log.Fatalf("Failed to list channel types: %v", err)
	}

	fmt.Printf("\nAvailable Channel Types (%d):\n", len(result.Payload))
	for _, ct := range result.Payload {
		fmt.Printf("  - %+v\n", ct)
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
