# \OrgGovernanceSettingsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrgIntegration**](OrgGovernanceSettingsAPI.md#CreateOrgIntegration) | **Post** /governance/api/v1/settings/integrations | Create an org integration
[**DeleteOrgIntegration**](OrgGovernanceSettingsAPI.md#DeleteOrgIntegration) | **Delete** /governance/api/v1/settings/integrations/{integrationId} | Delete an org integration
[**GetOrgCertificationSettings**](OrgGovernanceSettingsAPI.md#GetOrgCertificationSettings) | **Get** /governance/api/v1/settings/certification | Retrieve the org certification settings
[**GetOrgSettings**](OrgGovernanceSettingsAPI.md#GetOrgSettings) | **Get** /governance/api/v1/settings | Retrieve the org settings
[**ListOrgIntegrations**](OrgGovernanceSettingsAPI.md#ListOrgIntegrations) | **Get** /governance/api/v1/settings/integrations | List all org integrations
[**UpdateOrgCertificationSettings**](OrgGovernanceSettingsAPI.md#UpdateOrgCertificationSettings) | **Patch** /governance/api/v1/settings/certification | Update the org certification settings
[**UpdateOrgSettings**](OrgGovernanceSettingsAPI.md#UpdateOrgSettings) | **Patch** /governance/api/v1/settings | Update the org settings



## CreateOrgIntegration

> IntegrationFull CreateOrgIntegration(ctx).IntegrationCreatable(integrationCreatable).Execute()

Create an org integration



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
	integrationCreatable := *openapiclient.NewIntegrationCreatable() // IntegrationCreatable | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgGovernanceSettingsAPI.CreateOrgIntegration(context.Background()).IntegrationCreatable(integrationCreatable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgGovernanceSettingsAPI.CreateOrgIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrgIntegration`: IntegrationFull
	fmt.Fprintf(os.Stdout, "Response from `OrgGovernanceSettingsAPI.CreateOrgIntegration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrgIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **integrationCreatable** | [**IntegrationCreatable**](IntegrationCreatable.md) |  | 

### Return type

[**IntegrationFull**](IntegrationFull.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrgIntegration

> DeleteOrgIntegration(ctx, integrationId).Execute()

Delete an org integration



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
	integrationId := "integrationId_example" // string | Unique identifier for the integration

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrgGovernanceSettingsAPI.DeleteOrgIntegration(context.Background(), integrationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgGovernanceSettingsAPI.DeleteOrgIntegration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**integrationId** | **string** | Unique identifier for the integration | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrgIntegrationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrgCertificationSettings

> OrgCertificationSettings GetOrgCertificationSettings(ctx).Execute()

Retrieve the org certification settings



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
	resp, r, err := apiClient.OrgGovernanceSettingsAPI.GetOrgCertificationSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgGovernanceSettingsAPI.GetOrgCertificationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrgCertificationSettings`: OrgCertificationSettings
	fmt.Fprintf(os.Stdout, "Response from `OrgGovernanceSettingsAPI.GetOrgCertificationSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrgCertificationSettingsRequest struct via the builder pattern


### Return type

[**OrgCertificationSettings**](OrgCertificationSettings.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


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


## ListOrgIntegrations

> IntegrationsReadable ListOrgIntegrations(ctx).Execute()

List all org integrations



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
	resp, r, err := apiClient.OrgGovernanceSettingsAPI.ListOrgIntegrations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgGovernanceSettingsAPI.ListOrgIntegrations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrgIntegrations`: IntegrationsReadable
	fmt.Fprintf(os.Stdout, "Response from `OrgGovernanceSettingsAPI.ListOrgIntegrations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOrgIntegrationsRequest struct via the builder pattern


### Return type

[**IntegrationsReadable**](IntegrationsReadable.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateOrgCertificationSettings

> OrgCertificationSettingsPatchable UpdateOrgCertificationSettings(ctx).OrgCertificationSettingsPatchable(orgCertificationSettingsPatchable).Execute()

Update the org certification settings



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
	orgCertificationSettingsPatchable := *openapiclient.NewOrgCertificationSettingsPatchable() // OrgCertificationSettingsPatchable | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrgGovernanceSettingsAPI.UpdateOrgCertificationSettings(context.Background()).OrgCertificationSettingsPatchable(orgCertificationSettingsPatchable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrgGovernanceSettingsAPI.UpdateOrgCertificationSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOrgCertificationSettings`: OrgCertificationSettingsPatchable
	fmt.Fprintf(os.Stdout, "Response from `OrgGovernanceSettingsAPI.UpdateOrgCertificationSettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrgCertificationSettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orgCertificationSettingsPatchable** | [**OrgCertificationSettingsPatchable**](OrgCertificationSettingsPatchable.md) |  | 

### Return type

[**OrgCertificationSettingsPatchable**](OrgCertificationSettingsPatchable.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
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

