# \PrincipalSettingsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdatePrincipalSettings**](PrincipalSettingsAPI.md#UpdatePrincipalSettings) | **Patch** /governance/api/v1/principal-settings/{targetPrincipalId} | Update the principal settings



## UpdatePrincipalSettings

> PrincipalSettings UpdatePrincipalSettings(ctx, targetPrincipalId).PrincipalSettingsPatchable(principalSettingsPatchable).Execute()

Update the principal settings



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
	targetPrincipalId := "targetPrincipalId_example" // string | The `id` of the resource in Okta ID format or [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format
	principalSettingsPatchable := *openapiclient.NewPrincipalSettingsPatchable() // PrincipalSettingsPatchable | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrincipalSettingsAPI.UpdatePrincipalSettings(context.Background(), targetPrincipalId).PrincipalSettingsPatchable(principalSettingsPatchable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrincipalSettingsAPI.UpdatePrincipalSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePrincipalSettings`: PrincipalSettings
	fmt.Fprintf(os.Stdout, "Response from `PrincipalSettingsAPI.UpdatePrincipalSettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**targetPrincipalId** | **string** | The &#x60;id&#x60; of the resource in Okta ID format or [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePrincipalSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **principalSettingsPatchable** | [**PrincipalSettingsPatchable**](PrincipalSettingsPatchable.md) |  | 

### Return type

[**PrincipalSettings**](PrincipalSettings.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

