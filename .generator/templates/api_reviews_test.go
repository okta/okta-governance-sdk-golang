package governance

import (
	apiClient "github.com/okta/okta-governance-sdk-golang"
	"testing"
)

func TestReviewsAPIService_ListReviewsExecute(t *testing.T) {
	filter := `campaignId eq "icizc47q85Kdw1wzRxQ1d6"`
	execute, a, err := apiClient.ReviewsAPI.ListReviews(apiClient.cfg.Context).Filter(filter).Limit(5).Execute()
	if err != nil {
		t.Errorf("Error getting review: %v", err)
		return
	}

	if a.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", a.StatusCode)
		return
	}

	if len(execute.Data) == 0 {
		t.Errorf("Response should have contained a review")
	}

	for _, review := range execute.Data {
		if review.Id == "" {
			t.Errorf("Review ID should not be empty")
		}
		if review.CampaignId == "" {
			t.Errorf("Campaign id should not be empty")
		}
	}
}

func TestReviewsAPIService_GetReviewExecute(t *testing.T) {
	campaignId := "icizc47q85dqwK1wzRxQ1d6"
	filter := `campaignId eq "icizc47qdqd85K1wzRxQ1d6"`
	execute, _, _ := apiClient.ReviewsAPI.ListReviews(apiClient.cfg.Context).Filter(filter).Limit(5).Execute()
	reviews, a, err := apiClient.ReviewsAPI.GetReview(apiClient.cfg.Context, execute.Data[0].Id).Execute()
	if err != nil {
		t.Errorf("Error getting review: %v", err)
		return
	}

	if a.StatusCode != 200 {
		t.Errorf("Expected status code 200, got %d", a.StatusCode)
		return
	}

	if reviews.CampaignId != campaignId {
		t.Errorf("Expected campaignId %s, got %s", campaignId, reviews.CampaignId)
		return
	}
}

func TestReviewsAPIService_ReassignReviewsExecute(t *testing.T) {
	campaignId := "icizc47q85K1wzRxQ1d6"
	filter := `campaignId eq "icizc47q85K1wzRxQ1d6"`
	reviews, _, _ := apiClient.ReviewsAPI.ListReviews(apiClient.cfg.Context).Filter(filter).Execute()

	var reviewID string
	for _, review := range reviews.Data {
		if review.Decision == DECISION_UNREVIEWED {
			reviewID = review.Id
			break
		}
	}

	r := ReviewsReassign{ReviewerId: "00unli90kor62oF5Z1d7", ReviewIds: []string{reviewID}, Note: "testing"}

	execute, _, err := apiClient.ReviewsAPI.ReassignReviews(apiClient.cfg.Context, campaignId).ReviewsReassign(r).Execute()
	if err != nil {
		t.Errorf("Error getting review: %v", err)
		return
	}

	if len(execute.Data) == 0 {
		t.Errorf("Expected at least one review to be reassigned")
		return
	}
}
