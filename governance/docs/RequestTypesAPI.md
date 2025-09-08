# \RequestTypesAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRequestType**](RequestTypesAPI.md#CreateRequestType) | **Post** /governance/api/v1/request-types | Create a request type
[**DeleteRequestType**](RequestTypesAPI.md#DeleteRequestType) | **Delete** /governance/api/v1/request-types/{requestTypeId} | Delete a request type
[**GetRequestType**](RequestTypesAPI.md#GetRequestType) | **Get** /governance/api/v1/request-types/{requestTypeId} | Retrieve a request type
[**ListAllRequestTeams**](RequestTypesAPI.md#ListAllRequestTeams) | **Get** /governance/api/v1/teams | List all teams
[**ListAllRequestTypes**](RequestTypesAPI.md#ListAllRequestTypes) | **Get** /governance/api/v1/request-types | List all request types
[**PublishRequestType**](RequestTypesAPI.md#PublishRequestType) | **Post** /governance/api/v1/request-types/{requestTypeId}/publish | Publish a request type
[**UnpublishRequestType**](RequestTypesAPI.md#UnpublishRequestType) | **Post** /governance/api/v1/request-types/{requestTypeId}/un-publish | Unpublish a request type



## CreateRequestType

> RequestTypeFull CreateRequestType(ctx).RequestTypeCreatable(requestTypeCreatable).Execute()

Create a request type



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
	requestTypeCreatable := *openapiclient.NewRequestTypeCreatable("OwnerId_example", openapiclient.request-type-resource-settings-mutable{RequestTypeResourceSettingsApps: openapiclient.NewRequestTypeResourceSettingsApps("Type_example", []openapiclient.OktaApplicationResource{*openapiclient.NewOktaApplicationResource("ResourceId_example")})}, openapiclient.request-type-approval-settings-mutable{RequestTypeApprovalSettingsNone: openapiclient.NewRequestTypeApprovalSettingsNone("Type_example")}, "Name_example") // RequestTypeCreatable | The writable attributes of a request type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestTypesAPI.CreateRequestType(context.Background()).RequestTypeCreatable(requestTypeCreatable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestTypesAPI.CreateRequestType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRequestType`: RequestTypeFull
	fmt.Fprintf(os.Stdout, "Response from `RequestTypesAPI.CreateRequestType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequestTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestTypeCreatable** | [**RequestTypeCreatable**](RequestTypeCreatable.md) | The writable attributes of a request type | 

### Return type

[**RequestTypeFull**](RequestTypeFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRequestType

> DeleteRequestType(ctx, requestTypeId).Execute()

Delete a request type



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
	requestTypeId := "requestTypeId_example" // string | The `id` of the request type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RequestTypesAPI.DeleteRequestType(context.Background(), requestTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestTypesAPI.DeleteRequestType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestTypeId** | **string** | The &#x60;id&#x60; of the request type | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRequestTypeRequest struct via the builder pattern


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


## GetRequestType

> RequestTypeFull GetRequestType(ctx, requestTypeId).Execute()

Retrieve a request type



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
	requestTypeId := "requestTypeId_example" // string | The `id` of the request type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestTypesAPI.GetRequestType(context.Background(), requestTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestTypesAPI.GetRequestType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequestType`: RequestTypeFull
	fmt.Fprintf(os.Stdout, "Response from `RequestTypesAPI.GetRequestType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestTypeId** | **string** | The &#x60;id&#x60; of the request type | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestTypeFull**](RequestTypeFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllRequestTeams

> TeamsList ListAllRequestTeams(ctx).Filter(filter).After(after).Limit(limit).Execute()

List all teams



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
	filter := "name%20eq%20%22Salesforce%20admins%22" // string | Apply various filters by using supported team filtering properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  (optional)
	after := "after_example" // string | The after cursor provided by a prior request. (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestTypesAPI.ListAllRequestTeams(context.Background()).Filter(filter).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestTypesAPI.ListAllRequestTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllRequestTeams`: TeamsList
	fmt.Fprintf(os.Stdout, "Response from `RequestTypesAPI.ListAllRequestTeams`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllRequestTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Apply various filters by using supported team filtering properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  | 
 **after** | **string** | The after cursor provided by a prior request. | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]

### Return type

[**TeamsList**](TeamsList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllRequestTypes

> RequestTypesList ListAllRequestTypes(ctx).Filter(filter).After(after).Limit(limit).OrderBy(orderBy).Execute()

List all request types



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
	filter := "status%20eq%20%22ACTIVE%22" // string | Apply various filters by using supported request types filtering properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  (optional)
	after := "after_example" // string | The after cursor provided by a prior request. (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	orderBy := "created%20desc" // string | Apply an ordering of request types by specifying a supported request type property name with `%20asc` or `%20desc` suffix.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestTypesAPI.ListAllRequestTypes(context.Background()).Filter(filter).After(after).Limit(limit).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestTypesAPI.ListAllRequestTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllRequestTypes`: RequestTypesList
	fmt.Fprintf(os.Stdout, "Response from `RequestTypesAPI.ListAllRequestTypes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllRequestTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Apply various filters by using supported request types filtering properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  | 
 **after** | **string** | The after cursor provided by a prior request. | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **orderBy** | **string** | Apply an ordering of request types by specifying a supported request type property name with &#x60;%20asc&#x60; or &#x60;%20desc&#x60; suffix.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  | 

### Return type

[**RequestTypesList**](RequestTypesList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PublishRequestType

> RequestTypeFull PublishRequestType(ctx, requestTypeId).Execute()

Publish a request type



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
	requestTypeId := "requestTypeId_example" // string | The `id` of the request type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestTypesAPI.PublishRequestType(context.Background(), requestTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestTypesAPI.PublishRequestType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PublishRequestType`: RequestTypeFull
	fmt.Fprintf(os.Stdout, "Response from `RequestTypesAPI.PublishRequestType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestTypeId** | **string** | The &#x60;id&#x60; of the request type | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishRequestTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestTypeFull**](RequestTypeFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnpublishRequestType

> RequestTypeFull UnpublishRequestType(ctx, requestTypeId).Execute()

Unpublish a request type



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
	requestTypeId := "requestTypeId_example" // string | The `id` of the request type

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestTypesAPI.UnpublishRequestType(context.Background(), requestTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestTypesAPI.UnpublishRequestType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnpublishRequestType`: RequestTypeFull
	fmt.Fprintf(os.Stdout, "Response from `RequestTypesAPI.UnpublishRequestType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestTypeId** | **string** | The &#x60;id&#x60; of the request type | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnpublishRequestTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestTypeFull**](RequestTypeFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

