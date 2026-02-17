package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/resourcegraph/armresourcegraph"
)

func main() {
	ctx := context.Background()

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatal(err)
	}

	client, err := armresourcegraph.NewClient(cred, nil)
	if err != nil {
		log.Fatal(err)
	}

	query := `
Resources
| where type in (
  "microsoft.compute/virtualmachines",
  "microsoft.containerservice/managedclusters",
  "microsoft.web/sites",
  "microsoft.app/containerapps"
)
| project name, type, location, resourceGroup, subscriptionId, id
`

	subscriptions := []string{
		"<SUBSCRIPTION_ID_1>",
		"<SUBSCRIPTION_ID_2>",
	}

	req := armresourcegraph.QueryRequest{
		Subscriptions: subscriptions,
		Query:         &query,
	}

	resp, err := client.Resources(ctx, req, nil)
	if err != nil {
		log.Fatal(err)
	}

	// resp.Data is basically JSON
	fmt.Printf("%v\n", resp.Data)
}
