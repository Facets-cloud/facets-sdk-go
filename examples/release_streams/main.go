package main

import (
	"fmt"
	"log"
	"os"

	runtime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Facets-cloud/facets-sdk-go/facets/client"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_release_stream_controller"
)

func main() {
	facetsClient, auth := newClient()

	// List all release streams
	listReleaseStreams(facetsClient, auth)
}

// listReleaseStreams retrieves all configured release streams.
func listReleaseStreams(c *client.Facets, auth runtime.ClientAuthInfoWriter) {
	params := ui_release_stream_controller.NewGetAll1Params()
	result, err := c.UIReleaseStreamController.GetAll1(params, auth)
	if err != nil {
		log.Fatalf("Failed to list release streams: %v", err)
	}

	fmt.Printf("Release Streams (%d):\n", len(result.Payload))
	for _, rs := range result.Payload {
		fmt.Printf("  - %+v\n", rs)
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
