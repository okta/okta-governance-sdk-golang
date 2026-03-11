//go:build unit

package governance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCatalogsAPI_ListAllDefaultEntriesV2(t *testing.T) {
	cassetteName := CassetteName("catalogs", "list_all_default_entries")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// List parent entries (top-level)
	filter := "not(parent pr)"
	resp, httpRes, err := client.CatalogsAPI.ListAllDefaultEntriesV2(ctx).Filter(filter).Limit(10).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func TestCatalogsAPI_ListAllDefaultEntriesV2_WithMatch(t *testing.T) {
	cassetteName := CassetteName("catalogs", "list_all_default_entries_with_match")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// List parent entries with fuzzy match
	filter := "not(parent pr)"
	resp, httpRes, err := client.CatalogsAPI.ListAllDefaultEntriesV2(ctx).
		Filter(filter).
		Match("test").
		Limit(10).
		Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func TestCatalogsAPI_GetCatalogEntryV2(t *testing.T) {
	cassetteName := CassetteName("catalogs", "get_catalog_entry")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First list entries to get a valid entry ID
	filter := "not(parent pr)"
	listResp, _, err := client.CatalogsAPI.ListAllDefaultEntriesV2(ctx).Filter(filter).Limit(1).Execute()
	require.NoError(t, err)

	if len(listResp.Data) == 0 {
		t.Skip("No catalog entries available to test GetCatalogEntryV2")
	}

	entryId := listResp.Data[0].Id

	// Get the specific catalog entry
	resp, httpRes, err := client.CatalogsAPI.GetCatalogEntryV2(ctx, entryId).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
	assert.Equal(t, entryId, resp.Id)
}

func TestCatalogsAPI_ListAllDefaultUserEntriesV2(t *testing.T) {
	cassetteName := CassetteName("catalogs", "list_all_default_user_entries")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// List parent entries for a specific user
	filter := "not(parent pr)"
	resp, httpRes, err := client.CatalogsAPI.ListAllDefaultUserEntriesV2(ctx, TestPrincipalId).
		Filter(filter).
		Limit(10).
		Execute()

	// May return 404 if user doesn't exist
	if err != nil && httpRes != nil && httpRes.StatusCode == 404 {
		t.Skipf("User %s not found - use a valid user ID during recording", TestPrincipalId)
	}

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}
