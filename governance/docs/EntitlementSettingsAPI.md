# \EntitlementSettingsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEntitlementSettings**](EntitlementSettingsAPI.md#GetEntitlementSettings) | **Get** /governance/api/v2/resources/{resourceOrn}/entitlement-settings | Retrieve the entitlement settings for the resource
[**UpdateEntitlementSettings**](EntitlementSettingsAPI.md#UpdateEntitlementSettings) | **Patch** /governance/api/v2/resources/{resourceOrn}/entitlement-settings | Update the entitlement settings for the resource



## GetEntitlementSettings

> EntitlementSettingsFull GetEntitlementSettings(ctx, resourceOrn).Execute()

Retrieve the entitlement settings for the resource



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
	resourceOrn := "resourceOrn_example" // string | The resource ORN

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementSettingsAPI.GetEntitlementSettings(context.Background(), resourceOrn).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementSettingsAPI.GetEntitlementSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlementSettings`: EntitlementSettingsFull
	fmt.Fprintf(os.Stdout, "Response from `EntitlementSettingsAPI.GetEntitlementSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceOrn** | **string** | The resource ORN | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntitlementSettingsFull**](EntitlementSettingsFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntitlementSettings

> EntitlementSettingsFull UpdateEntitlementSettings(ctx, resourceOrn).EntitlementSettingsUpdatable(entitlementSettingsUpdatable).Execute()

Update the entitlement settings for the resource



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
	resourceOrn := "resourceOrn_example" // string | The resource ORN
	entitlementSettingsUpdatable := *openapiclient.NewEntitlementSettingsUpdatable(openapiclient.entitlement-settings-request-status("OPTED_IN")) // EntitlementSettingsUpdatable | The updatable attributes of the entitlement settings

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementSettingsAPI.UpdateEntitlementSettings(context.Background(), resourceOrn).EntitlementSettingsUpdatable(entitlementSettingsUpdatable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementSettingsAPI.UpdateEntitlementSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEntitlementSettings`: EntitlementSettingsFull
	fmt.Fprintf(os.Stdout, "Response from `EntitlementSettingsAPI.UpdateEntitlementSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceOrn** | **string** | The resource ORN | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntitlementSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entitlementSettingsUpdatable** | [**EntitlementSettingsUpdatable**](EntitlementSettingsUpdatable.md) | The updatable attributes of the entitlement settings | 

### Return type

[**EntitlementSettingsFull**](EntitlementSettingsFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

