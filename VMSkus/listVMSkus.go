// To use this code:
// go run listVMSkus your_subscription_id
// The code lists out all available virtual machine SKUs in an Azure subscription

package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2020-06-01/compute"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func main() {
	subID := os.Args[1]

	AzureVMSize(subID)
}

func AzureVMSize(subID string) *compute.ResourceSkusClient {
	auth1, err := auth.NewAuthorizerFromCLI()

	if err != nil {
		fmt.Println(err)
	}

	vmClient := compute.NewResourceSkusClient(subID)
	vmClient.Authorizer = auth1

	sizes, err := vmClient.List(context.Background(), "centralus")

	if err != nil {
		fmt.Println(err)
	}

	for _, list := range sizes.Values() {
		fmt.Printf(*list.Name)
	}

	return &vmClient
}
