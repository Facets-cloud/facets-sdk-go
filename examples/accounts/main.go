package main

import (
	"fmt"
	"log"
	"os"

	runtime "github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/Facets-cloud/facets-sdk-go/facets/client"
	"github.com/Facets-cloud/facets-sdk-go/facets/client/ui_accounts_controller"
)

func main() {
	facetsClient, auth := newClient()

	// List all accounts (cloud, VCS, etc.)
	listAllAccounts(facetsClient, auth)

	// List accounts filtered by type (e.g., "AWS", "GCP", "AZURE", "GITHUB", "BITBUCKET", "GITLAB", "KUBERNETES")
	listAccountsByType(facetsClient, auth, "AWS")
}

// listAllAccounts retrieves all configured accounts.
func listAllAccounts(c *client.Facets, auth runtime.ClientAuthInfoWriter) {
	params := ui_accounts_controller.NewGetAllAccountsParams()
	result, err := c.UIAccountsController.GetAllAccounts(params, auth)
	if err != nil {
		log.Fatalf("Failed to list accounts: %v", err)
	}

	fmt.Printf("All Accounts (%d):\n", len(result.Payload))
	for _, acct := range result.Payload {
		fmt.Printf("  - Name: %-20s | Provider: %-12s | Type: %-16s | In Use: %v\n",
			acct.Name, acct.Provider, acct.AccountType, acct.InUse)
	}
}

// listAccountsByType retrieves accounts filtered by provider type.
func listAccountsByType(c *client.Facets, auth runtime.ClientAuthInfoWriter, accountType string) {
	params := ui_accounts_controller.NewGetAccountsByTypeParams()
	params.Type = accountType

	result, err := c.UIAccountsController.GetAccountsByType(params, auth)
	if err != nil {
		log.Fatalf("Failed to list %s accounts: %v", accountType, err)
	}

	fmt.Printf("\n%s Accounts (%d):\n", accountType, len(result.Payload))
	for _, acct := range result.Payload {
		fmt.Printf("  - Name: %s | ID: %s | Created By: %s\n",
			acct.Name, acct.ID, acct.CreatedBy)
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
