# \RequestConditionsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateResourceRequestConditionV2**](RequestConditionsAPI.md#ActivateResourceRequestConditionV2) | **Post** /governance/api/v2/resources/{resourceId}/request-conditions/{requestConditionId}/activate | Activate the request condition
[**CreateResourceRequestConditionV2**](RequestConditionsAPI.md#CreateResourceRequestConditionV2) | **Post** /governance/api/v2/resources/{resourceId}/request-conditions | Create a request condition
[**DeactivateResourceRequestConditionV2**](RequestConditionsAPI.md#DeactivateResourceRequestConditionV2) | **Post** /governance/api/v2/resources/{resourceId}/request-conditions/{requestConditionId}/deactivate | Deactivate the request condition
[**DeleteResourceRequestConditionV2**](RequestConditionsAPI.md#DeleteResourceRequestConditionV2) | **Delete** /governance/api/v2/resources/{resourceId}/request-conditions/{requestConditionId} | Delete a request condition
[**GetResourceRequestConditionV2**](RequestConditionsAPI.md#GetResourceRequestConditionV2) | **Get** /governance/api/v2/resources/{resourceId}/request-conditions/{requestConditionId} | Retrieve a resource request condition
[**ListResourceRequestConditionsV2**](RequestConditionsAPI.md#ListResourceRequestConditionsV2) | **Get** /governance/api/v2/resources/{resourceId}/request-conditions | List all resource request conditions
[**UpdateResourceRequestConditionV2**](RequestConditionsAPI.md#UpdateResourceRequestConditionV2) | **Patch** /governance/api/v2/resources/{resourceId}/request-conditions/{requestConditionId} | Update the request condition



## ActivateResourceRequestConditionV2

> RequestConditionFull ActivateResourceRequestConditionV2(ctx, resourceId, requestConditionId).Execute()

Activate the request condition



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
	resourceId := "resourceId_example" // string | The `id` of the resource in Okta ID format or ORN format
	requestConditionId := "requestConditionId_example" // string | The `id` of the request condition

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestConditionsAPI.ActivateResourceRequestConditionV2(context.Background(), resourceId, requestConditionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestConditionsAPI.ActivateResourceRequestConditionV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ActivateResourceRequestConditionV2`: RequestConditionFull
	fmt.Fprintf(os.Stdout, "Response from `RequestConditionsAPI.ActivateResourceRequestConditionV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The &#x60;id&#x60; of the resource in Okta ID format or ORN format | 
**requestConditionId** | **string** | The &#x60;id&#x60; of the request condition | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateResourceRequestConditionV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RequestConditionFull**](RequestConditionFull.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateResourceRequestConditionV2

> RequestConditionFull CreateResourceRequestConditionV2(ctx, resourceId).RequestConditionCreatable(requestConditionCreatable).Execute()

Create a request condition



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
	resourceId := "resourceId_example" // string | The `id` of the resource in Okta ID format or ORN format
	requestConditionCreatable := *openapiclient.NewRequestConditionCreatable(openapiclient.requester-settings-creatable_RequesterSettings{EveryoneRequesterSettings: openapiclient.NewEveryoneRequesterSettings("Type_example")}, openapiclient.access-scope-settings-creatable_AccessScopeSettings{AccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings: openapiclient.NewAccessScopeSettingsCreatableEntitlementBundleAccessScopeSettings("Type_example", []openapiclient.EntitlementBundlesArrayCreatableInner{*openapiclient.NewEntitlementBundlesArrayCreatableInner("Id_example")})}, "ApprovalSequenceId_example", "Name_example") // RequestConditionCreatable | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestConditionsAPI.CreateResourceRequestConditionV2(context.Background(), resourceId).RequestConditionCreatable(requestConditionCreatable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestConditionsAPI.CreateResourceRequestConditionV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateResourceRequestConditionV2`: RequestConditionFull
	fmt.Fprintf(os.Stdout, "Response from `RequestConditionsAPI.CreateResourceRequestConditionV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The &#x60;id&#x60; of the resource in Okta ID format or ORN format | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateResourceRequestConditionV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestConditionCreatable** | [**RequestConditionCreatable**](RequestConditionCreatable.md) |  | 

### Return type

[**RequestConditionFull**](RequestConditionFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateResourceRequestConditionV2

> RequestConditionFull DeactivateResourceRequestConditionV2(ctx, resourceId, requestConditionId).Execute()

Deactivate the request condition



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
	resourceId := "resourceId_example" // string | The `id` of the resource in Okta ID format or ORN format
	requestConditionId := "requestConditionId_example" // string | The `id` of the request condition

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestConditionsAPI.DeactivateResourceRequestConditionV2(context.Background(), resourceId, requestConditionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestConditionsAPI.DeactivateResourceRequestConditionV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeactivateResourceRequestConditionV2`: RequestConditionFull
	fmt.Fprintf(os.Stdout, "Response from `RequestConditionsAPI.DeactivateResourceRequestConditionV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The &#x60;id&#x60; of the resource in Okta ID format or ORN format | 
**requestConditionId** | **string** | The &#x60;id&#x60; of the request condition | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateResourceRequestConditionV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RequestConditionFull**](RequestConditionFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResourceRequestConditionV2

> DeleteResourceRequestConditionV2(ctx, resourceId, requestConditionId).Execute()

Delete a request condition



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
	resourceId := "resourceId_example" // string | The `id` of the resource in Okta ID format or ORN format
	requestConditionId := "requestConditionId_example" // string | The `id` of the request condition

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RequestConditionsAPI.DeleteResourceRequestConditionV2(context.Background(), resourceId, requestConditionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestConditionsAPI.DeleteResourceRequestConditionV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The &#x60;id&#x60; of the resource in Okta ID format or ORN format | 
**requestConditionId** | **string** | The &#x60;id&#x60; of the request condition | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceRequestConditionV2Request struct via the builder pattern


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


## GetResourceRequestConditionV2

> RequestConditionFull GetResourceRequestConditionV2(ctx, resourceId, requestConditionId).Execute()

Retrieve a resource request condition



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
	resourceId := "resourceId_example" // string | The `id` of the resource in Okta ID format or ORN format
	requestConditionId := "requestConditionId_example" // string | The `id` of the request condition

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestConditionsAPI.GetResourceRequestConditionV2(context.Background(), resourceId, requestConditionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestConditionsAPI.GetResourceRequestConditionV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResourceRequestConditionV2`: RequestConditionFull
	fmt.Fprintf(os.Stdout, "Response from `RequestConditionsAPI.GetResourceRequestConditionV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The &#x60;id&#x60; of the resource in Okta ID format or ORN format | 
**requestConditionId** | **string** | The &#x60;id&#x60; of the request condition | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceRequestConditionV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RequestConditionFull**](RequestConditionFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResourceRequestConditionsV2

> RequestConditionsList ListResourceRequestConditionsV2(ctx, resourceId).Execute()

List all resource request conditions



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
	resourceId := "resourceId_example" // string | The `id` of the resource in Okta ID format or ORN format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestConditionsAPI.ListResourceRequestConditionsV2(context.Background(), resourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestConditionsAPI.ListResourceRequestConditionsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListResourceRequestConditionsV2`: RequestConditionsList
	fmt.Fprintf(os.Stdout, "Response from `RequestConditionsAPI.ListResourceRequestConditionsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The &#x60;id&#x60; of the resource in Okta ID format or ORN format | 

### Other Parameters

Other parameters are passed through a pointer to a apiListResourceRequestConditionsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestConditionsList**](RequestConditionsList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceRequestConditionV2

> RequestConditionFull UpdateResourceRequestConditionV2(ctx, resourceId, requestConditionId).RequestConditionPatchable(requestConditionPatchable).Execute()

Update the request condition



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
	resourceId := "resourceId_example" // string | The `id` of the resource in Okta ID format or ORN format
	requestConditionId := "requestConditionId_example" // string | The `id` of the request condition
	requestConditionPatchable := *openapiclient.NewRequestConditionPatchable() // RequestConditionPatchable | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestConditionsAPI.UpdateResourceRequestConditionV2(context.Background(), resourceId, requestConditionId).RequestConditionPatchable(requestConditionPatchable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestConditionsAPI.UpdateResourceRequestConditionV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateResourceRequestConditionV2`: RequestConditionFull
	fmt.Fprintf(os.Stdout, "Response from `RequestConditionsAPI.UpdateResourceRequestConditionV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The &#x60;id&#x60; of the resource in Okta ID format or ORN format | 
**requestConditionId** | **string** | The &#x60;id&#x60; of the request condition | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResourceRequestConditionV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **requestConditionPatchable** | [**RequestConditionPatchable**](RequestConditionPatchable.md) |  | 

### Return type

[**RequestConditionFull**](RequestConditionFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

