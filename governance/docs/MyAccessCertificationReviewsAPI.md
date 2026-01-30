# \MyAccessCertificationReviewsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListMyManagedConnections**](MyAccessCertificationReviewsAPI.md#ListMyManagedConnections) | **Get** /governance/api/v1/my/campaigns/{campaignId}/reviews/{reviewId}/agent-managed-connections/{agentId} | List all managed connections for my review



## ListMyManagedConnections

> MyManagedConnections ListMyManagedConnections(ctx, campaignId, reviewId, agentId).After(after).Limit(limit).Execute()

List all managed connections for my review



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
	reviewId := "reviewId_example" // string | The `id` of the review
	agentId := "wlpkhjm0jnDp8RrUu0g4" // string | ID of the agent
	after := "00u68w6vzKLultXS97g6" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyAccessCertificationReviewsAPI.ListMyManagedConnections(context.Background(), campaignId, reviewId, agentId).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyAccessCertificationReviewsAPI.ListMyManagedConnections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMyManagedConnections`: MyManagedConnections
	fmt.Fprintf(os.Stdout, "Response from `MyAccessCertificationReviewsAPI.ListMyManagedConnections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The &#x60;id&#x60; of the campaign | 
**reviewId** | **string** | The &#x60;id&#x60; of the review | 
**agentId** | **string** | ID of the agent | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMyManagedConnectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]

### Return type

[**MyManagedConnections**](MyManagedConnections.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

