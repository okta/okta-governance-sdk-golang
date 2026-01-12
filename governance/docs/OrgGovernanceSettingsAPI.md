# \OrgGovernanceSettingsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetOrgSettings**](OrgGovernanceSettingsAPI.md#GetOrgSettings) | **Get** /governance/api/v1/settings | Retrieve the org settings
[**UpdateOrgSettings**](OrgGovernanceSettingsAPI.md#UpdateOrgSettings) | **Patch** /governance/api/v1/settings | Update the org settings



## GetOrgSettings

> OrgSettings GetOrgSettings(ctx).Execute()

Retrieve the org settings



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
	resp, r, err := apiClient.OrgGovernanceSettingsAPI.GetOrgSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgGovernanceSettingsAPI.GetOrgSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgSettings`: OrgSettings
	fmt.Fprintf(os.Stdout, "Response from `OrgGovernanceSettingsAPI.GetOrgSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgSettingsRequest struct via the builder pattern


### Return type

[**OrgSettings**](OrgSettings.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgSettings

> OrgSettingsPatchable UpdateOrgSettings(ctx).OrgSettingsPatchable(orgSettingsPatchable).Execute()

Update the org settings



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
	orgSettingsPatchable := *openapiclient.NewOrgSettingsPatchable() // OrgSettingsPatchable | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgGovernanceSettingsAPI.UpdateOrgSettings(context.Background()).OrgSettingsPatchable(orgSettingsPatchable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgGovernanceSettingsAPI.UpdateOrgSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgSettings`: OrgSettingsPatchable
	fmt.Fprintf(os.Stdout, "Response from `OrgGovernanceSettingsAPI.UpdateOrgSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgSettingsPatchable** | [**OrgSettingsPatchable**](OrgSettingsPatchable.md) |  | 

### Return type

[**OrgSettingsPatchable**](OrgSettingsPatchable.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

