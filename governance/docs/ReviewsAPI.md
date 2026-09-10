# \ReviewsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetReview**](ReviewsAPI.md#GetReview) | **Get** /governance/api/v1/reviews/{reviewId} | Retrieve a review
[**ListReviews**](ReviewsAPI.md#ListReviews) | **Get** /governance/api/v1/reviews | List all reviews
[**ReassignReviews**](ReviewsAPI.md#ReassignReviews) | **Post** /governance/api/v1/campaigns/{campaignId}/reviews/reassign | Reassign the reviews



## GetReview

> ReviewFull GetReview(ctx, reviewId).Execute()

Retrieve a review



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
	reviewId := "reviewId_example" // string | Unique identifier for the review

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReviewsAPI.GetReview(context.Background(), reviewId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReviewsAPI.GetReview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReview`: ReviewFull
	fmt.Fprintf(os.Stdout, "Response from `ReviewsAPI.GetReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reviewId** | **string** | Unique identifier for the review | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ReviewFull**](ReviewFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListReviews

> ReviewList ListReviews(ctx).Filter(filter).After(after).Limit(limit).OrderBy(orderBy).Execute()

List all reviews



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
	filter := "campaignId eq "icitdyhndQ6qstyvR8g5"" // string | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters a collection of reviews in the response. The filter expression supports the `eq` [operator](https://developer.okta.com/docs/api/#operators) and the following properties: * `campaignId` * `principalId` (corresponds to `principalProfile.id`) * `reviewerId` (corresponds to `reviewerProfile.id`) * `decision` * `resourceId` * `reviewerType` * `reviewerLevel` (corresponds to `currentReviewerLevel`) * `entitlementValueId` (corresponds to `entitlementValue.id`) * `entitlementBundleId` (corresponds to `entitlementBundle.id`)  > **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  (optional)
	after := "00u68w6vzKLultXS97g6" // string | Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the standard link headers. See [pagination](https://developer.okta.com/docs/api/#pagination). (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	orderBy := []string{"Inner_example"} // []string | Specifies a property to sort the results. The following properties are supported: - `decided` - `decision` - `remediationStatus` - `created`  Append `desc` or `asc` to indicate the sorting direction.  By default, the results are sorted by `created` in ascending order (`created asc`).  > **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  (optional) (default to ["created asc"])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReviewsAPI.ListReviews(context.Background()).Filter(filter).After(after).Limit(limit).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReviewsAPI.ListReviews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListReviews`: ReviewList
	fmt.Fprintf(os.Stdout, "Response from `ReviewsAPI.ListReviews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListReviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters a collection of reviews in the response. The filter expression supports the &#x60;eq&#x60; [operator](https://developer.okta.com/docs/api/#operators) and the following properties: * &#x60;campaignId&#x60; * &#x60;principalId&#x60; (corresponds to &#x60;principalProfile.id&#x60;) * &#x60;reviewerId&#x60; (corresponds to &#x60;reviewerProfile.id&#x60;) * &#x60;decision&#x60; * &#x60;resourceId&#x60; * &#x60;reviewerType&#x60; * &#x60;reviewerLevel&#x60; (corresponds to &#x60;currentReviewerLevel&#x60;) * &#x60;entitlementValueId&#x60; (corresponds to &#x60;entitlementValue.id&#x60;) * &#x60;entitlementBundleId&#x60; (corresponds to &#x60;entitlementBundle.id&#x60;)  &gt; **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  | 
 **after** | **string** | Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the standard link headers. See [pagination](https://developer.okta.com/docs/api/#pagination). | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **orderBy** | **[]string** | Specifies a property to sort the results. The following properties are supported: - &#x60;decided&#x60; - &#x60;decision&#x60; - &#x60;remediationStatus&#x60; - &#x60;created&#x60;  Append &#x60;desc&#x60; or &#x60;asc&#x60; to indicate the sorting direction.  By default, the results are sorted by &#x60;created&#x60; in ascending order (&#x60;created asc&#x60;).  &gt; **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  | [default to [&quot;created asc&quot;]]

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


## ReassignReviews

> ReviewReassignList ReassignReviews(ctx, campaignId).ReviewsReassign(reviewsReassign).Execute()

Reassign the reviews



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
	reviewsReassign := *openapiclient.NewReviewsReassign("ReviewerId_example", []string{"ReviewIds_example"}, "Note_example") // ReviewsReassign | The operation payload for reviews reassignment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ReviewsAPI.ReassignReviews(context.Background(), campaignId).ReviewsReassign(reviewsReassign).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ReviewsAPI.ReassignReviews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReassignReviews`: ReviewReassignList
	fmt.Fprintf(os.Stdout, "Response from `ReviewsAPI.ReassignReviews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | Unique identifier for the campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiReassignReviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reviewsReassign** | [**ReviewsReassign**](ReviewsReassign.md) | The operation payload for reviews reassignment | 

### Return type

[**ReviewReassignList**](ReviewReassignList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

