package governance

import (
	"testing"
	"time"
)

func TestApiCreateCampaignRequest_Execute(t *testing.T) {
	createdCampaign, resp, err := apiClient.CampaignsAPI.CreateCampaign(apiClient.cfg.Context).CampaignMutable(buildCampaign()).Execute()
	if err != nil {
		t.Errorf("Error creating campaign: %v", err)
		return
	}
	if resp.StatusCode != 201 {
		t.Errorf("Expected status code 201, got %d", resp.StatusCode)
		return
	}

	campaigns, resp, err := apiClient.CampaignsAPI.ListCampaigns(apiClient.cfg.Context).Limit(10).Execute()
	if err != nil {
		t.Errorf("Error listing campaigns: %v", err)
		return
	}
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
	if len(campaigns.Data) == 0 {
		t.Errorf("Expected at least one campaign")
	}

	t.Cleanup(func() {
		if createdCampaign.Id != "" {
			execute, err := apiClient.CampaignsAPI.DeleteCampaign(apiClient.cfg.Context, createdCampaign.Id).Execute()
			if execute.StatusCode != 201 || err != nil {
				return
			} // nolint
		}
	})
}

func TestApiGetCampaignRequest_Execute(t *testing.T) {
	campaign, _, err := apiClient.CampaignsAPI.CreateCampaign(apiClient.cfg.Context).
		CampaignMutable(buildCampaign()).
		Execute()
	if err != nil {
		t.Fatalf("Error creating campaign: %v", err)
	}

	gotCampaign, resp, err := apiClient.CampaignsAPI.GetCampaign(apiClient.cfg.Context, campaign.Id).Execute()
	if err != nil {
		t.Errorf("Error getting campaign: %v", err)
		return
	}
	if resp.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", resp.StatusCode)
	}
	if gotCampaign.Id != campaign.Id {
		t.Errorf("Expected campaign ID %s, got %s", campaign.Id, gotCampaign.Id)
	}

	t.Cleanup(func() {
		if campaign.Id != "" {
			execute, err := apiClient.CampaignsAPI.DeleteCampaign(apiClient.cfg.Context, campaign.Id).Execute()
			if execute.StatusCode != 201 || err != nil {
				return
			} // nolint
		}
	})
}

func TestApiEndCampaignRequest_Execute(t *testing.T) {
	campaign, _, err := apiClient.CampaignsAPI.CreateCampaign(apiClient.cfg.Context).
		CampaignMutable(buildCampaign()).
		Execute()

	if campaign != nil {
		execute, err := apiClient.CampaignsAPI.LaunchCampaign(apiClient.cfg.Context, campaign.Id).Execute()
		if err != nil {
			return
		}

		if execute.StatusCode != 202 {
			t.Errorf("Expected status code 202, got %d", execute.StatusCode)
			return
		}
	}

	_, err = apiClient.CampaignsAPI.EndCampaign(apiClient.cfg.Context, campaign.Id).Execute()
	if err != nil {
		t.Errorf("Error ending campaign: %v", err)
	}
}

func TestApiDeleteCampaignRequest_Execute(t *testing.T) {
	campaign, _, err := apiClient.CampaignsAPI.CreateCampaign(apiClient.cfg.Context).
		CampaignMutable(buildCampaign()).
		Execute()
	if err != nil {
		t.Fatalf("Error creating campaign: %v", err)
	}

	_, err = apiClient.CampaignsAPI.DeleteCampaign(apiClient.cfg.Context, campaign.Id).Execute()
	if err != nil {
		t.Errorf("Error deleting campaign: %v", err)
	}
}

func buildCampaign() CampaignMutable {
	reviewerGroupID := "00gnkw1sdqL30MdGk1d7"
	reviewerId := "00unkw1sfbTw08c0g1d7"
	reviewerScopeExpression := "user.profile.reviewerId"
	str := "2025-10-04T13:43:40.000Z"
	rt := "GROUP"
	parsedTime, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return CampaignMutable{}
	}

	return CampaignMutable{
		Name: "Group Campaign | User Reviewer Expression_Test2",
		RemediationSettings: RemediationSettings{
			AccessApproved: "NO_ACTION",
			AccessRevoked:  "DENY",
			NoResponse:     "NO_ACTION",
		},
		ResourceSettings: ResourceSettingsMutable{
			Type: "GROUP",
			TargetResources: []TargetResourcesRequestInner{
				{
					ResourceId:   "00gnkw1sdqL30MdGk1d7",
					ResourceType: (*ResourceType)(&rt),
				},
			},
		},
		ReviewerSettings: ReviewerSettingsMutable{
			Type:                    "USER",
			ReviewerGroupId:         &reviewerGroupID,
			ReviewerId:              &reviewerId,
			ReviewerScopeExpression: &reviewerScopeExpression,
		},
		ScheduleSettings: ScheduleSettingsMutable{
			StartDate:      parsedTime,
			DurationInDays: 30,
			TimeZone:       "America/New_York",
			Type:           "ONE_OFF",
		},
	}
}
