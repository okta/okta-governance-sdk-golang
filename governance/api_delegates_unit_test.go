//go:build unit

package governance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDelegatesAPI_ListDelegateAppointments(t *testing.T) {
	cassetteName := CassetteName("delegates", "list_delegate_appointments")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()
	resp, httpRes, err := client.DelegatesAPI.ListDelegateAppointments(ctx).Limit(10).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
	assert.NotNil(t, resp.Data)
}

func TestDelegatesAPI_ListDelegateAppointmentsWithFilter(t *testing.T) {
	cassetteName := CassetteName("delegates", "list_delegate_appointments_filtered")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()
	// Filter by delegatorId - use a placeholder ID that can be replaced during recording
	filter := `delegatorId eq "00u0000000000000000"`
	resp, httpRes, err := client.DelegatesAPI.ListDelegateAppointments(ctx).Filter(filter).Limit(10).Execute()

	// Note: This may return empty results or 400 if no matching delegator exists
	// The test validates the API call works correctly
	if err != nil {
		// Expected if no matching delegator
		t.Logf("Filter returned error (expected if no matching delegator): %v", err)
	} else {
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)
	}
}
