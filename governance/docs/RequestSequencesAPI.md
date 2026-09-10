# \RequestSequencesAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRequestSequenceV2**](RequestSequencesAPI.md#DeleteRequestSequenceV2) | **Delete** /governance/api/v2/request-sequences/{sequenceId} | Delete a request sequence
[**GetResourceRequestSequenceV2**](RequestSequencesAPI.md#GetResourceRequestSequenceV2) | **Get** /governance/api/v2/resources/{resourceId}/request-sequences/{sequenceId} | Retrieve a resource request sequence
[**ListResourceRequestSequencesV2**](RequestSequencesAPI.md#ListResourceRequestSequencesV2) | **Get** /governance/api/v2/resources/{resourceId}/request-sequences | List all resource request sequences



## DeleteRequestSequenceV2

> DeleteRequestSequenceV2(ctx, sequenceId).Execute()

Delete a request sequence



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
	sequenceId := "sequenceId_example" // string | The `id` of the sequence

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RequestSequencesAPI.DeleteRequestSequenceV2(context.Background(), sequenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestSequencesAPI.DeleteRequestSequenceV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sequenceId** | **string** | The &#x60;id&#x60; of the sequence | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequestSequenceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceRequestSequenceV2

> RequestSequence GetResourceRequestSequenceV2(ctx, resourceId, sequenceId).Execute()

Retrieve a resource request sequence



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
	resourceId := "resourceId_example" // string | Unique identifier for the resource in Okta instance ID format or [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn)
	sequenceId := "sequenceId_example" // string | The `id` of the sequence

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestSequencesAPI.GetResourceRequestSequenceV2(context.Background(), resourceId, sequenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestSequencesAPI.GetResourceRequestSequenceV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResourceRequestSequenceV2`: RequestSequence
	fmt.Fprintf(os.Stdout, "Response from `RequestSequencesAPI.GetResourceRequestSequenceV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | Unique identifier for the resource in Okta instance ID format or [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) | 
**sequenceId** | **string** | The &#x60;id&#x60; of the sequence | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceRequestSequenceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RequestSequence**](RequestSequence.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResourceRequestSequencesV2

> RequestSequencesList ListResourceRequestSequencesV2(ctx, resourceId).Execute()

List all resource request sequences



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
	resourceId := "resourceId_example" // string | Unique identifier for the resource in Okta instance ID format or [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestSequencesAPI.ListResourceRequestSequencesV2(context.Background(), resourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestSequencesAPI.ListResourceRequestSequencesV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListResourceRequestSequencesV2`: RequestSequencesList
	fmt.Fprintf(os.Stdout, "Response from `RequestSequencesAPI.ListResourceRequestSequencesV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | Unique identifier for the resource in Okta instance ID format or [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) | 

### Other Parameters

Other parameters are passed through a pointer to a apiListResourceRequestSequencesV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestSequencesList**](RequestSequencesList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

