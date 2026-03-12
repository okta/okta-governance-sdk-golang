//go:build vcr

package governance

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLabelsAPI_ListLabels(t *testing.T) {
	cassetteName := CassetteName("labels", "list_labels")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()
	resp, httpRes, err := client.LabelsAPI.ListLabels(ctx).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func TestLabelsAPI_GetLabel(t *testing.T) {
	cassetteName := CassetteName("labels", "get_label")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// First list labels to get a valid label ID
	listResp, _, err := client.LabelsAPI.ListLabels(ctx).Execute()
	require.NoError(t, err)

	if len(listResp.Data) == 0 {
		t.Skip("No labels available to test GetLabel")
	}

	labelId := listResp.Data[0].LabelId

	// Get the specific label
	resp, httpRes, err := client.LabelsAPI.GetLabel(ctx, labelId).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
	assert.Equal(t, labelId, resp.LabelId)
}

func TestLabelsAPI_ListLabelResources(t *testing.T) {
	cassetteName := CassetteName("labels", "list_label_resources")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()
	// Filter is required for this API - resourceType uses ORN object type (e.g., "apps")
	filter := `resourceType eq "apps"`
	resp, httpRes, err := client.LabelsAPI.ListLabelResources(ctx).Filter(filter).Limit(10).Execute()

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, 200, httpRes.StatusCode)
}

func TestLabelsAPI_CreateAndDeleteLabel(t *testing.T) {
	cassetteName := CassetteName("labels", "create_delete_label")
	skipIfNotRecording(t, cassetteName)

	client := NewRecorderClient(t, cassetteName)
	defer func() {
		require.NoError(t, client.Stop())
	}()

	ctx := context.Background()

	// Create a new label with specific test data
	labelName := "test-sdk-label"
	labelValues := []LabelValueCreate{
		{Name: "value1"},
		{Name: "value2"},
	}
	createReq := LabelCreate{
		Name:   labelName,
		Values: labelValues,
	}

	createResp, httpRes, err := client.LabelsAPI.CreateLabel(ctx).LabelCreate(createReq).Execute()

	// Verify the response contains the correct data
	require.NoError(t, err)
	require.NotNil(t, createResp)
	assert.Equal(t, 201, httpRes.StatusCode)

	// Verify the label was created with the correct name
	assert.Equal(t, labelName, createResp.Name)
	assert.NotEmpty(t, createResp.LabelId, "Label should have an ID")

	// Verify the label values were created
	require.Len(t, createResp.Values, 2, "Label should have 2 values")
	assert.Equal(t, "value1", createResp.Values[0].Name)
	assert.Equal(t, "value2", createResp.Values[1].Name)
	assert.NotEmpty(t, createResp.Values[0].LabelValueId, "Label value should have an ID")
	assert.NotEmpty(t, createResp.Values[1].LabelValueId, "Label value should have an ID")

	// Delete the label
	_, err = client.LabelsAPI.DeleteLabel(ctx, createResp.LabelId).Execute()
	require.NoError(t, err)
}
