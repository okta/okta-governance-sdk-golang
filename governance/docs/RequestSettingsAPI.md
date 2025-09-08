# \RequestSettingsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgRequestSettingsV2**](RequestSettingsAPI.md#GetOrgRequestSettingsV2) | **Get** /governance/api/v2/request-settings | Retrieve the request settings for the organization
[**GetRequestSettingsV2**](RequestSettingsAPI.md#GetRequestSettingsV2) | **Get** /governance/api/v2/resources/{resourceId}/request-settings | Retrieve the request settings for a resource
[**UpdateOrgRequestSettingsV2**](RequestSettingsAPI.md#UpdateOrgRequestSettingsV2) | **Patch** /governance/api/v2/request-settings | Update the request settings for the organization
[**UpdateResourceRequestSettingsV2**](RequestSettingsAPI.md#UpdateResourceRequestSettingsV2) | **Patch** /governance/api/v2/resources/{resourceId}/request-settings | Update the resource request settings



## GetOrgRequestSettingsV2

> OrgRequestSettings GetOrgRequestSettingsV2(ctx).Execute()

Retrieve the request settings for the organization



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestSettingsAPI.GetOrgRequestSettingsV2(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestSettingsAPI.GetOrgRequestSettingsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgRequestSettingsV2`: OrgRequestSettings
	fmt.Fprintf(os.Stdout, "Response from `RequestSettingsAPI.GetOrgRequestSettingsV2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgRequestSettingsV2Request struct via the builder pattern


### Return type

[**OrgRequestSettings**](OrgRequestSettings.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequestSettingsV2

> RequestSettings GetRequestSettingsV2(ctx, resourceId).Execute()

Retrieve the request settings for a resource



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
	resp, r, err := apiClient.RequestSettingsAPI.GetRequestSettingsV2(context.Background(), resourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestSettingsAPI.GetRequestSettingsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequestSettingsV2`: RequestSettings
	fmt.Fprintf(os.Stdout, "Response from `RequestSettingsAPI.GetRequestSettingsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The &#x60;id&#x60; of the resource in Okta ID format or ORN format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestSettingsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestSettings**](RequestSettings.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgRequestSettingsV2

> OrgRequestSettingsPatchable UpdateOrgRequestSettingsV2(ctx).OrgRequestSettingsPatchable(orgRequestSettingsPatchable).Execute()

Update the request settings for the organization



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
	orgRequestSettingsPatchable := *openapiclient.NewOrgRequestSettingsPatchable(false) // OrgRequestSettingsPatchable | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestSettingsAPI.UpdateOrgRequestSettingsV2(context.Background()).OrgRequestSettingsPatchable(orgRequestSettingsPatchable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestSettingsAPI.UpdateOrgRequestSettingsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgRequestSettingsV2`: OrgRequestSettingsPatchable
	fmt.Fprintf(os.Stdout, "Response from `RequestSettingsAPI.UpdateOrgRequestSettingsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgRequestSettingsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgRequestSettingsPatchable** | [**OrgRequestSettingsPatchable**](OrgRequestSettingsPatchable.md) |  | 

### Return type

[**OrgRequestSettingsPatchable**](OrgRequestSettingsPatchable.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceRequestSettingsV2

> RequestSettings UpdateResourceRequestSettingsV2(ctx, resourceId).ResourceRequestSettingsPatchable(resourceRequestSettingsPatchable).Execute()

Update the resource request settings



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
	resourceRequestSettingsPatchable := *openapiclient.NewResourceRequestSettingsPatchable() // ResourceRequestSettingsPatchable | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestSettingsAPI.UpdateResourceRequestSettingsV2(context.Background(), resourceId).ResourceRequestSettingsPatchable(resourceRequestSettingsPatchable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestSettingsAPI.UpdateResourceRequestSettingsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateResourceRequestSettingsV2`: RequestSettings
	fmt.Fprintf(os.Stdout, "Response from `RequestSettingsAPI.UpdateResourceRequestSettingsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | The &#x60;id&#x60; of the resource in Okta ID format or ORN format | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResourceRequestSettingsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceRequestSettingsPatchable** | [**ResourceRequestSettingsPatchable**](ResourceRequestSettingsPatchable.md) |  | 

### Return type

[**RequestSettings**](RequestSettings.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

