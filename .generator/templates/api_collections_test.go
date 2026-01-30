package governance

import (
	"testing"
)

func TestCollectionsAPIService_CreateCollectionExecute(t *testing.T) {
	execute, a, err := apiClient.CollectionsAPI.CreateCollection(apiClient.cfg.Context).CollectionCreatable(buildCollection()).Execute()
	if err != nil {
		t.Errorf("Error getting collection: %v", err)
		return
	}

	if a.StatusCode != 201 {
		t.Errorf("Expected status code 200, got %d", a.StatusCode)
		return
	}

	if execute.Id == "" {
		t.Errorf("Response should have contained a collection")
	}
}

func buildCollection() CollectionCreatable {
	desc := "This is a test collection"
	return CollectionCreatable{
		Name:        "Test Collection",
		Description: &desc,
	}
}
