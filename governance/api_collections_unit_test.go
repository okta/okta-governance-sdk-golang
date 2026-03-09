//go:build unit

package governance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCollectionsAPI_ListCollections(t *testing.T) {
	cassetteName := CassetteName("collections", "list_collections")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()
	resp, httpRes, err := client.CollectionsAPI.ListCollections(ctx).Limit(10).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
	assert.NotNil(t, resp.Data)
}

func TestCollectionsAPI_CreateAndDeleteCollection(t *testing.T) {
	cassetteName := CassetteName("collections", "create_delete_collection")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// Create collection
	collection := BuildTestCollection()
	createdCollection, httpRes, err := client.CollectionsAPI.CreateCollection(ctx).CollectionCreatable(collection).Execute()

	require.NoError(t, err)
	require.NotNil(t, createdCollection)
	assert.Equal(t, 201, httpRes.StatusCode)
	assert.NotEmpty(t, createdCollection.Id)

	// Clean up - delete the collection
	t.Cleanup(func() {
		if createdCollection.Id != "" {
			_, _ = client.CollectionsAPI.DeleteCollection(ctx, createdCollection.Id).Execute()
		}
	})
}

func TestCollectionsAPI_GetCollection(t *testing.T) {
	cassetteName := CassetteName("collections", "get_collection")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First create a collection
	collection := BuildTestCollectionWithName("GetCollection Test")
	createdCollection, _, err := client.CollectionsAPI.CreateCollection(ctx).CollectionCreatable(collection).Execute()
	require.NoError(t, err)
	require.NotNil(t, createdCollection)

	t.Cleanup(func() {
		if createdCollection.Id != "" {
			_, _ = client.CollectionsAPI.DeleteCollection(ctx, createdCollection.Id).Execute()
		}
	})

	// Get the collection
	gotCollection, httpRes, err := client.CollectionsAPI.GetCollection(ctx, createdCollection.Id).Execute()

	require.NoError(t, err)
	require.NotNil(t, gotCollection)
	assert.Equal(t, 200, httpRes.StatusCode)
	assert.Equal(t, createdCollection.Id, gotCollection.Id)
}

func TestCollectionsAPI_UpdateCollection(t *testing.T) {
	cassetteName := CassetteName("collections", "update_collection")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First create a collection
	collection := BuildTestCollectionWithName("UpdateCollection Test")
	createdCollection, _, err := client.CollectionsAPI.CreateCollection(ctx).CollectionCreatable(collection).Execute()
	require.NoError(t, err)
	require.NotNil(t, createdCollection)

	t.Cleanup(func() {
		if createdCollection.Id != "" {
			_, _ = client.CollectionsAPI.DeleteCollection(ctx, createdCollection.Id).Execute()
		}
	})

	// Replace the collection
	updatedName := "Updated Collection Name"
	updatedDesc := "Updated description"
	updateData := CollectionUpdatable{
		Name:        updatedName,
		Description: &updatedDesc,
	}

	updatedCollection, httpRes, err := client.CollectionsAPI.ReplaceCollection(ctx, createdCollection.Id).CollectionUpdatable(updateData).Execute()

	require.NoError(t, err)
	require.NotNil(t, updatedCollection)
	assert.Equal(t, 200, httpRes.StatusCode)
	assert.Equal(t, updatedName, updatedCollection.Name)
}

func TestCollectionsAPI_ListCollectionResources(t *testing.T) {
	cassetteName := CassetteName("collections", "list_collection_resources")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First list collections to get a valid collection ID
	listResp, _, err := client.CollectionsAPI.ListCollections(ctx).Limit(1).Execute()
	require.NoError(t, err)

	if len(listResp.Data) == 0 {
		t.Skip("No collections available to test ListCollectionResources")
	}

	collectionId := listResp.Data[0].Id

	// List collection resources
	resp, httpRes, err := client.CollectionsAPI.ListCollectionResources(ctx, collectionId).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}
