package main

import (
	"fmt"
	"log"
	"os"
	"time"

	runtime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Facets-cloud/facets-sdk-go/facets/client"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_audit_logs_controller"
)

func main() {
	facetsClient, auth := newClient()

	// Get audit logs from the last 7 days
	getRecentAuditLogs(facetsClient, auth)

	// Get audit logs filtered by stack
	getAuditLogsByStack(facetsClient, auth, "my-stack")
}

// getRecentAuditLogs retrieves audit logs from the past 7 days.
func getRecentAuditLogs(c *client.Facets, auth runtime.ClientAuthInfoWriter) {
	params := ui_audit_logs_controller.NewGetAuditLogsParams()

	// Start is required — set to 7 days ago
	start := strfmt.DateTime(time.Now().AddDate(0, 0, -7))
	params.Start = start

	// Optional: limit page size
	size := int32(20)
	params.Size = &size

	result, err := c.UIAuditLogsController.GetAuditLogs(params, auth)
	if err != nil {
		log.Fatalf("Failed to get audit logs: %v", err)
	}

	fmt.Printf("Recent Audit Logs:\n")
	fmt.Printf("  %+v\n", result.Payload)
}

// getAuditLogsByStack retrieves audit logs filtered by stack name.
func getAuditLogsByStack(c *client.Facets, auth runtime.ClientAuthInfoWriter, stackName string) {
	params := ui_audit_logs_controller.NewGetAuditLogsParams()

	start := strfmt.DateTime(time.Now().AddDate(0, 0, -30))
	params.Start = start
	params.StackName = &stackName

	size := int32(50)
	params.Size = &size

	result, err := c.UIAuditLogsController.GetAuditLogs(params, auth)
	if err != nil {
		log.Fatalf("Failed to get audit logs for stack %q: %v", stackName, err)
	}

	fmt.Printf("\nAudit Logs for stack %q:\n", stackName)
	fmt.Printf("  %+v\n", result.Payload)
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
