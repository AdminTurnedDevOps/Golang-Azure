package AzureAuth

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func TestAzureAuth(t *testing.T) {
	auth1, err := auth.NewAuthorizerFromCLI()

	assert.NotNil(t, auth1)

	if err != nil {
		log.Println(err)
	}
}
