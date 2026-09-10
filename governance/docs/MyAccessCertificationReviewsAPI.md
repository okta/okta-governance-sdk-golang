# \MyAccessCertificationReviewsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCampaignBulkDecisionsJobStatus**](MyAccessCertificationReviewsAPI.md#GetCampaignBulkDecisionsJobStatus) | **Get** /governance/api/v1/my/campaigns/{campaignId}/reviews/bulk-decisions/jobs/{jobId} | Retrieve the status of a bulk-review submission
[**ListCampaignReviews**](MyAccessCertificationReviewsAPI.md#ListCampaignReviews) | **Get** /governance/api/v1/my/campaigns/{campaignId}/reviews | List my reviews for a campaign
[**ListMyResourceConnections**](MyAccessCertificationReviewsAPI.md#ListMyResourceConnections) | **Get** /governance/api/v1/my/campaigns/{campaignId}/reviews/{reviewId}/agent-resource-connections/{agentId} | List all resource connections for my review
[**SubmitCampaignBulkDecisions**](MyAccessCertificationReviewsAPI.md#SubmitCampaignBulkDecisions) | **Post** /governance/api/v1/my/campaigns/{campaignId}/reviews/bulk-decisions | Submit a bulk-review decision
[**SubmitMyCampaignReviewActions**](MyAccessCertificationReviewsAPI.md#SubmitMyCampaignReviewActions) | **Post** /governance/api/v1/my/campaigns/{campaignId}/reviews/actions | Submit my review actions



## GetCampaignBulkDecisionsJobStatus

> BulkReviewSubmissionJobDetails GetCampaignBulkDecisionsJobStatus(ctx, campaignId, jobId).Execute()

Retrieve the status of a bulk-review submission



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-governance-sdk-golang"
)

func main() {
	campaignId := "campaignId_example" // string | Unique identifier for the campaign
	jobId := "jobId_example" // string | Unique identifier for the job

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyAccessCertificationReviewsAPI.GetCampaignBulkDecisionsJobStatus(context.Background(), campaignId, jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccessCertificationReviewsAPI.GetCampaignBulkDecisionsJobStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignBulkDecisionsJobStatus`: BulkReviewSubmissionJobDetails
	fmt.Fprintf(os.Stdout, "Response from `MyAccessCertificationReviewsAPI.GetCampaignBulkDecisionsJobStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Unique identifier for the campaign | 
**jobId** | **string** | Unique identifier for the job | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignBulkDecisionsJobStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**BulkReviewSubmissionJobDetails**](BulkReviewSubmissionJobDetails.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCampaignReviews

> ReviewList ListCampaignReviews(ctx, campaignId).Filter(filter).After(after).Limit(limit).OrderBy(orderBy).Execute()

List my reviews for a campaign



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-governance-sdk-golang"
)

func main() {
	campaignId := "campaignId_example" // string | Unique identifier for the campaign
	filter := "decision eq "UNREVIEWED"" // string | Apply various filters by using supported review filtering properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  Supported filters are:  - `id`: string - `principalId`: string - `reviewerId`: string - `decision`: string (APPROVE, REVOKE, UNREVIEWED) - `resourceId`: string (`GroupId` or `AppId`) - `reviewerType`: string (`USER` or `GROUP` or `RESOURCE_OWNER`) - `reviewerLevel`: string (`FIRST` or `SECOND`) - `entitlementValueId`: string - `entitlementBundleId`: string  (optional)
	after := "00u68w6vzKLultXS97g6" // string | Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the standard link headers. See [pagination](https://developer.okta.com/docs/api/#pagination). (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	orderBy := []string{"Inner_example"} // []string | Specifies a property to sort the results. The following properties are supported: - `decided` - `decision` - `remediationStatus` - `created`  Append `desc` or `asc` to indicate the sorting direction.  > **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyAccessCertificationReviewsAPI.ListCampaignReviews(context.Background(), campaignId).Filter(filter).After(after).Limit(limit).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccessCertificationReviewsAPI.ListCampaignReviews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCampaignReviews`: ReviewList
	fmt.Fprintf(os.Stdout, "Response from `MyAccessCertificationReviewsAPI.ListCampaignReviews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Unique identifier for the campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCampaignReviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | Apply various filters by using supported review filtering properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  Supported filters are:  - &#x60;id&#x60;: string - &#x60;principalId&#x60;: string - &#x60;reviewerId&#x60;: string - &#x60;decision&#x60;: string (APPROVE, REVOKE, UNREVIEWED) - &#x60;resourceId&#x60;: string (&#x60;GroupId&#x60; or &#x60;AppId&#x60;) - &#x60;reviewerType&#x60;: string (&#x60;USER&#x60; or &#x60;GROUP&#x60; or &#x60;RESOURCE_OWNER&#x60;) - &#x60;reviewerLevel&#x60;: string (&#x60;FIRST&#x60; or &#x60;SECOND&#x60;) - &#x60;entitlementValueId&#x60;: string - &#x60;entitlementBundleId&#x60;: string  | 
 **after** | **string** | Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the standard link headers. See [pagination](https://developer.okta.com/docs/api/#pagination). | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **orderBy** | **[]string** | Specifies a property to sort the results. The following properties are supported: - &#x60;decided&#x60; - &#x60;decision&#x60; - &#x60;remediationStatus&#x60; - &#x60;created&#x60;  Append &#x60;desc&#x60; or &#x60;asc&#x60; to indicate the sorting direction.  &gt; **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  | 

### Return type

[**ReviewList**](ReviewList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMyResourceConnections

> MyResourceConnections ListMyResourceConnections(ctx, campaignId, reviewId, agentId).After(after).Limit(limit).Execute()

List all resource connections for my review



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-governance-sdk-golang"
)

func main() {
	campaignId := "campaignId_example" // string | Unique identifier for the campaign
	reviewId := "reviewId_example" // string | Unique identifier for the review
	agentId := "wlpkhjm0jnDp8RrUu0g4" // string | ID of the agent
	after := "00u68w6vzKLultXS97g6" // string | Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the standard link headers. See [pagination](https://developer.okta.com/docs/api/#pagination). (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyAccessCertificationReviewsAPI.ListMyResourceConnections(context.Background(), campaignId, reviewId, agentId).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccessCertificationReviewsAPI.ListMyResourceConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMyResourceConnections`: MyResourceConnections
	fmt.Fprintf(os.Stdout, "Response from `MyAccessCertificationReviewsAPI.ListMyResourceConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Unique identifier for the campaign | 
**reviewId** | **string** | Unique identifier for the review | 
**agentId** | **string** | ID of the agent | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMyResourceConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **after** | **string** | Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the standard link headers. See [pagination](https://developer.okta.com/docs/api/#pagination). | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]

### Return type

[**MyResourceConnections**](MyResourceConnections.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitCampaignBulkDecisions

> BulkReviewSubmissionDetails SubmitCampaignBulkDecisions(ctx, campaignId).BulkReviewSubmissionMutable(bulkReviewSubmissionMutable).Execute()

Submit a bulk-review decision



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-governance-sdk-golang"
)

func main() {
	campaignId := "campaignId_example" // string | Unique identifier for the campaign
	bulkReviewSubmissionMutable := *openapiclient.NewBulkReviewSubmissionMutable(*openapiclient.NewBulkReviewSubmissionCriteriaMutable(openapiclient.bulk-review-criteria-type("GOVERNANCE_ANALYZER")), openapiclient.bulk-review-decision("APPROVE")) // BulkReviewSubmissionMutable | Request body for submitting a decision for bulk reviews 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyAccessCertificationReviewsAPI.SubmitCampaignBulkDecisions(context.Background(), campaignId).BulkReviewSubmissionMutable(bulkReviewSubmissionMutable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccessCertificationReviewsAPI.SubmitCampaignBulkDecisions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitCampaignBulkDecisions`: BulkReviewSubmissionDetails
	fmt.Fprintf(os.Stdout, "Response from `MyAccessCertificationReviewsAPI.SubmitCampaignBulkDecisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Unique identifier for the campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitCampaignBulkDecisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkReviewSubmissionMutable** | [**BulkReviewSubmissionMutable**](BulkReviewSubmissionMutable.md) | Request body for submitting a decision for bulk reviews  | 

### Return type

[**BulkReviewSubmissionDetails**](BulkReviewSubmissionDetails.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitMyCampaignReviewActions

> ReviewList SubmitMyCampaignReviewActions(ctx, campaignId).ReviewItemsActionMutable(reviewItemsActionMutable).Execute()

Submit my review actions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/okta/okta-governance-sdk-golang"
)

func main() {
	campaignId := "campaignId_example" // string | Unique identifier for the campaign
	reviewItemsActionMutable := *openapiclient.NewReviewItemsActionMutable([]openapiclient.ReviewActionItem{*openapiclient.NewReviewActionItem("ReviewId_example", openapiclient.review-action("APPROVE"))}) // ReviewItemsActionMutable | Submit my review actions request body

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyAccessCertificationReviewsAPI.SubmitMyCampaignReviewActions(context.Background(), campaignId).ReviewItemsActionMutable(reviewItemsActionMutable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccessCertificationReviewsAPI.SubmitMyCampaignReviewActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitMyCampaignReviewActions`: ReviewList
	fmt.Fprintf(os.Stdout, "Response from `MyAccessCertificationReviewsAPI.SubmitMyCampaignReviewActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Unique identifier for the campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitMyCampaignReviewActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reviewItemsActionMutable** | [**ReviewItemsActionMutable**](ReviewItemsActionMutable.md) | Submit my review actions request body | 

### Return type

[**ReviewList**](ReviewList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

