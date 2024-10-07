package sdk_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	sdk "github.com/registry-tools/rt-sdk"
	"github.com/registry-tools/rt-sdk/generated/api"
	"github.com/registry-tools/rt-sdk/generated/models"
)

func testClient(t *testing.T) sdk.SDK {
	t.Helper()

	client, err := sdk.NewSDK(os.Getenv("REGISTRY_TOOLS_HOST"), os.Getenv("REGISTRY_TOOLS_CLIENT_ID"), os.Getenv("REGISTRY_TOOLS_CLIENT_SECRET"))
	if err != nil {
		t.Fatalf("Failed to create API Client: %v", err)
	}

	if client == nil {
		t.Fatal("API Client was nil")
	}

	return client
}

func TestAccServiceDiscovery(t *testing.T) {
	client := testClient(t)

	ctx := context.Background()

	response, err := client.WellKnown().TerraformJson().GetAsTerraformGetResponse(ctx, nil)
	if err != nil {
		t.Fatalf("Failed to fetch terraform service discovery: %v", err)
	}

	data := response.GetAdditionalData()

	services := []string{
		"modules.v1",
		"tfe.v2",
		"modules-rt.v1",
	}

	for _, svc := range services {
		path, ok := data[svc]
		if !ok {
			t.Errorf("Service discovery did not contain a %s service", svc)
		}
		if *path.(*string) == "" {
			t.Errorf("Service discovery service \"%s\" did not contain a service path", svc)
		}
	}
}

func TestAccNamespacesList(t *testing.T) {
	client := testClient(t)

	ctx := context.Background()

	response, err := client.Api().Namespaces().GetAsNamespacesGetResponse(ctx, nil)
	if err != nil {
		if modelError, ok := err.(*models.Errors); ok {
			fatalError(t, modelError)
		} else {
			t.Fatalf("Got unexpected error %T: %v", err, err)
		}
	}

	for _, ns := range response.GetData() {
		t.Logf("Got namespace %s / %s", *ns.GetId(), *ns.GetName())
	}
}

func TestAccNamespacesCreate(t *testing.T) {
	client := testClient(t)

	ctx := context.Background()

	newNamespace := models.NewNamespace()
	name := fmt.Sprintf("test-namespace-%d", time.Now().Unix())
	description := "Test namespace description"
	newNamespace.SetName(&name)
	newNamespace.SetDescription(&description)

	newNamespaceBody := api.NewNamespacesPostRequestBody()
	newNamespaceBody.SetNamespace(newNamespace)

	response, err := client.Api().Namespaces().GetAsNamespacesGetResponse(ctx, nil)
	if err != nil {
		if modelError, ok := err.(*models.Errors); ok {
			fatalError(t, modelError)
		} else {
			t.Fatalf("Got unexpected error %T: %v", err, err)
		}
	}

	for _, ns := range response.GetData() {
		t.Logf("Got namespace %s / %s", *ns.GetId(), *ns.GetName())
	}

	namespace, err := client.Api().Namespaces().PostAsNamespacesPostResponse(ctx, newNamespaceBody, nil)
	if err != nil {
		if modelError, ok := err.(*models.Errors); ok {
			fatalError(t, modelError)
		} else {
			t.Fatalf("Got unexpected error %T: %v", err, err)
		}
	}

	t.Logf("Created namespace %s / %s", *namespace.GetData().GetId(), *namespace.GetData().GetName())
}

func fatalError(t *testing.T, err *models.Errors) {
	t.Helper()

	t.Logf("Got error status %d", err.ResponseStatusCode)
	for _, e := range err.GetErrors() {
		t.Logf("  - %s: %s", *e.GetTitle(), *e.GetDetail())
	}
	t.Fatal()
}
