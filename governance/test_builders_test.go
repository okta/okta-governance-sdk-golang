package governance

import (
	"time"
)

// Test data builder functions for creating test fixtures
// These builders create realistic test data for API testing

// BuildTestCampaign creates a campaign mutable object for testing
func BuildTestCampaign() CampaignMutable {
	reviewerGroupID := TestGroupId
	reviewerId := TestUserId
	reviewerScopeExpression := "user.profile.reviewerId"
	str := "2025-10-04T13:43:40.000Z"
	rt := "GROUP"
	parsedTime, err := time.Parse(time.RFC3339, str)
	if err != nil {
		return CampaignMutable{}
	}

	return CampaignMutable{
		Name: "Test Campaign | SDK Unit Test",
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



// BuildTestEntitlement creates an entitlement create object for testing
func BuildTestEntitlement() EntitlementCreate {
	return EntitlementCreate{
		Name:          "Test Entitlement",
		ExternalValue: "test-external-value",
		Description:   nil,
		DataType:      "string",
		Parent:        &TargetResource{Type: "APPLICATION", ExternalId: "0oao1otjwc8OEd6z61d7"},
		Values: []EntitlementValueWritableProperties{
			{
				Name:          "Test Value",
				ExternalValue: "test-value",
				Description:   nil,
			},
		},
		MultiValue: true,
	}
}

// BuildTestEntitlementWithParent creates an entitlement with a specific parent resource
func BuildTestEntitlementWithParent(appId string) EntitlementCreate {
	return EntitlementCreate{
		Name:          "Test Entitlement",
		ExternalValue: "test-external-value",
		Description:   nil,
		DataType:      "string",
		Parent:        &TargetResource{Type: "APPLICATION", ExternalId: appId},
		Values: []EntitlementValueWritableProperties{
			{
				Name:          "Test Value",
				ExternalValue: "test-value",
				Description:   nil,
			},
		},
		MultiValue: true,
	}
}

// BuildTestEntitlementFull creates an entitlement full object for update operations
func BuildTestEntitlementFull(entitlementId, appId string) EntitlementsFullWithParent {
	str := "Test Value Full"
	return EntitlementsFullWithParent{
		Name:              "Test Entitlement Full",
		ExternalValue:     "test-external-value-full",
		Id:                entitlementId,
		Description:       nil,
		DataType:          "string",
		Parent:            TargetResource{Type: "APPLICATION", ExternalId: appId},
		ParentResourceOrn: "orn:oktapreview:idp:00onkw1sbuAh3Q06I1d7:apps:saasure:" + appId,
		Values: []EntitlementValueFull{
			{
				Name:          &str,
				ExternalValue: &str,
			},
		},
	}
}



// BuildTestEntitlementBundle creates an entitlement bundle for testing
func BuildTestEntitlementBundle() EntitlementBundleCreatable {
	desc := "Test entitlement bundle created by SDK unit tests"
	return EntitlementBundleCreatable{
		Name:        "Test Entitlement Bundle",
		Description: &desc,
	}
}

// BuildTestEntitlementBundleWithName creates an entitlement bundle with a custom name
func BuildTestEntitlementBundleWithName(name string) EntitlementBundleCreatable {
	desc := "Test entitlement bundle: " + name
	return EntitlementBundleCreatable{
		Name:        name,
		Description: &desc,
	}
}

// BuildTestResourceOwnerPatch creates a resource owner patch for testing
func BuildTestResourceOwnerPatch(userOrns []string) ResourceOwnersPatch {
	data := make([]ResourceOwnersPatchDataInner, len(userOrns))
	op := "add"
	path := "/principalOrn"
	for i, userOrn := range userOrns {
		data[i] = ResourceOwnersPatchDataInner{
			Op:    &op,
			Path:  &path,
			Value: &userOrn,
		}
	}
	return ResourceOwnersPatch{
		Data: data,
	}
}

// Test IDs - Constants for test fixtures
// These can be overridden in tests that need specific IDs
// Note: These IDs are for dcp-sep25.oktapreview.com test org
const (
	TestApplicationId       = "0oaqd21lm797hdT4q1d7" // Okta Identity Governance app
	TestUserId              = "00uqd1g3jmuDjGhpE1d7"
	TestGroupId             = "00gqd1g3h7drGOpDL1d7" // Everyone group
	TestEntitlementBundleId = "enb18v0ta5boKOyxA1d7" // Group Administrator bundle
	TestPrincipalId         = "00uqd1g3jmuDjGhpE1d7" // Same as TestUserId - principal can be a user
)
