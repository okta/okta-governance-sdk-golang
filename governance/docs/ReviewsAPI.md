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
	reviewId := "reviewId_example" // string | The `id` of the review

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
**reviewId** | **string** | The &#x60;id&#x60; of the review | 

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
	filter := "campaignId%20eq%20%22icitdyhndQ6qstyvR8g5%22" // string | Apply various filters by using supported review filtering properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  (optional)
	after := "00u68w6vzKLultXS97g6" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	orderBy := []string{"Inner_example"} // []string | A field by which results can be sorted. For now, sorting by a single field is supported.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  (optional) (default to ["created asc"])

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
 **filter** | **string** | Apply various filters by using supported review filtering properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  | 
 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **orderBy** | **[]string** | A field by which results can be sorted. For now, sorting by a single field is supported.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  | [default to [&quot;created asc&quot;]]

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
	campaignId := "campaignId_example" // string | The `id` of the campaign
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
**campaignId** | **string** | The &#x60;id&#x60; of the campaign | 

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

