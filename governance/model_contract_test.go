//go:build contract

package governance

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// =============================================================================
// Label Model Contract Tests
// These tests verify the Label model structure matches the API contract.
// If the API changes (fields added/removed/renamed), these tests will fail.
// =============================================================================

func TestLabel_StructureAndSerialization(t *testing.T) {
	// Test that Label struct has expected fields and serializes correctly
	label := Label{
		LabelId: "lbc123456789",
		Name:    "Environment",
		Values: []LabelValue{
			{LabelValueId: "lbl111", Name: "production"},
			{LabelValueId: "lbl222", Name: "staging"},
		},
	}

	// Serialize to JSON
	data, err := json.Marshal(label)
	require.NoError(t, err)

	// Parse JSON to verify field names
	var parsed map[string]interface{}
	err = json.Unmarshal(data, &parsed)
	require.NoError(t, err)

	// Verify JSON field names match API contract
	assert.Contains(t, parsed, "labelId", "Label must have 'labelId' field")
	assert.Contains(t, parsed, "name", "Label must have 'name' field")
	assert.Contains(t, parsed, "values", "Label must have 'values' field")

	// Verify field values
	assert.Equal(t, "lbc123456789", parsed["labelId"])
	assert.Equal(t, "Environment", parsed["name"])

	// Verify nested values structure
	values, ok := parsed["values"].([]interface{})
	require.True(t, ok, "values must be an array")
	require.Len(t, values, 2)

	value1 := values[0].(map[string]interface{})
	assert.Contains(t, value1, "labelValueId")
	assert.Contains(t, value1, "name")
}

func TestLabel_ParseFromJSON(t *testing.T) {
	// Test that JSON from API can be parsed into Label struct
	jsonData := `{
		"labelId": "lbc999",
		"name": "Region",
		"values": [
			{"labelValueId": "lbl001", "name": "us-east"},
			{"labelValueId": "lbl002", "name": "eu-west"}
		],
		"_links": {
			"self": {"href": "https://example.com/labels/lbc999"}
		}
	}`

	var label Label
	err := json.Unmarshal([]byte(jsonData), &label)

	require.NoError(t, err)
	assert.Equal(t, "lbc999", label.LabelId)
	assert.Equal(t, "Region", label.Name)
	require.Len(t, label.Values, 2)
	assert.Equal(t, "lbl001", label.Values[0].LabelValueId)
	assert.Equal(t, "us-east", label.Values[0].Name)
}

// =============================================================================
// LabelCreate Request Contract Tests
// =============================================================================

func TestLabelCreate_StructureAndSerialization(t *testing.T) {
	req := LabelCreate{
		Name: "test-label",
		Values: []LabelValueCreate{
			{Name: "value1"},
			{Name: "value2"},
		},
	}

	data, err := json.Marshal(req)
	require.NoError(t, err)

	var parsed map[string]interface{}
	err = json.Unmarshal(data, &parsed)
	require.NoError(t, err)

	// Verify required fields are present
	assert.Contains(t, parsed, "name", "LabelCreate must have 'name' field")
	assert.Contains(t, parsed, "values", "LabelCreate must have 'values' field")

	// Verify structure
	assert.Equal(t, "test-label", parsed["name"])
	values := parsed["values"].([]interface{})
	assert.Len(t, values, 2)
}

// =============================================================================
// Bundle Model Contract Tests (EntitlementBundle)
// =============================================================================

func TestBundle_StructureAndSerialization(t *testing.T) {
	bundle := Bundle{
		Id:   "enb123456",
		Name: "Admin Bundle",
	}

	data, err := json.Marshal(bundle)
	require.NoError(t, err)

	var parsed map[string]interface{}
	err = json.Unmarshal(data, &parsed)
	require.NoError(t, err)

	// Verify JSON field names
	assert.Contains(t, parsed, "id", "Bundle must have 'id' field")
	assert.Contains(t, parsed, "name", "Bundle must have 'name' field")
}

func TestBundle_ParseFromJSON(t *testing.T) {
	jsonData := `{
		"id": "enb789",
		"name": "Super Admin"
	}`

	var bundle Bundle
	err := json.Unmarshal([]byte(jsonData), &bundle)

	require.NoError(t, err)
	assert.Equal(t, "enb789", bundle.Id)
	assert.Equal(t, "Super Admin", bundle.Name)
}

// =============================================================================
// Team Model Contract Tests
// =============================================================================

func TestTeam_StructureAndSerialization(t *testing.T) {
	now := time.Now()
	team := Team{
		Id:          "team123",
		Name:        "IT Team",
		Created:     now,
		LastUpdated: now,
	}

	data, err := json.Marshal(team)
	require.NoError(t, err)

	var parsed map[string]interface{}
	err = json.Unmarshal(data, &parsed)
	require.NoError(t, err)

	assert.Contains(t, parsed, "id", "Team must have 'id' field")
	assert.Contains(t, parsed, "name", "Team must have 'name' field")
	assert.Contains(t, parsed, "created", "Team must have 'created' field")
	assert.Contains(t, parsed, "lastUpdated", "Team must have 'lastUpdated' field")
}

func TestTeam_ParseFromJSON(t *testing.T) {
	jsonData := `{
		"id": "68d4d26d899ab659c87c5317",
		"name": "IT",
		"created": "2025-09-25T05:26:05.307Z",
		"lastUpdated": "2025-09-25T05:26:05.307Z"
	}`

	var team Team
	err := json.Unmarshal([]byte(jsonData), &team)

	require.NoError(t, err)
	assert.Equal(t, "68d4d26d899ab659c87c5317", team.Id)
	assert.Equal(t, "IT", team.Name)
	assert.False(t, team.Created.IsZero(), "Created timestamp should be parsed")
}

// =============================================================================
// CampaignDetailsReadOnly Model Contract Tests
// Note: This model has many required fields - simplified test focuses on key fields
// =============================================================================

func TestCampaignDetailsReadOnly_HasRequiredFields(t *testing.T) {
	// Test that the model struct has the expected key fields
	campaign := CampaignDetailsReadOnly{
		Name: "Q1 Access Review",
	}

	// Verify the struct can be created with Name field
	assert.Equal(t, "Q1 Access Review", campaign.Name)

	// Verify Description pointer field exists
	desc := "Test description"
	campaign.Description = &desc
	assert.Equal(t, "Test description", *campaign.Description)
}

// =============================================================================
// OrgSettings Model Contract Tests
// =============================================================================

func TestOrgSettings_ParseFromJSON(t *testing.T) {
	jsonData := `{
		"delegates": {
			"endUser": {
				"enabled": true
			}
		}
	}`

	var settings OrgSettings
	err := json.Unmarshal([]byte(jsonData), &settings)

	require.NoError(t, err)
	assert.NotNil(t, settings.Delegates)
}

// =============================================================================
// DelegateAppointment Model Contract Tests
// =============================================================================

func TestDelegateAppointment_HasRequiredFields(t *testing.T) {
	// Test that DelegateAppointment struct has expected fields
	now := time.Now()
	appointment := DelegateAppointment{
		Id:        "del123",
		CreatedBy: "admin001",
		Created:   now,
	}

	assert.Equal(t, "del123", appointment.Id)
	assert.Equal(t, "admin001", appointment.CreatedBy)
	assert.False(t, appointment.Created.IsZero())
}

// =============================================================================
// OrgRequestSettings Model Contract Tests
// =============================================================================

func TestOrgRequestSettings_HasRequiredFields(t *testing.T) {
	settings := OrgRequestSettings{
		SubprocessorsAcknowledged: true,
		LongTimePastProvisioned:   false,
	}

	assert.True(t, settings.SubprocessorsAcknowledged)
	assert.False(t, settings.LongTimePastProvisioned)
}
