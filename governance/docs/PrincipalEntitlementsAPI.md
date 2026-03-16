# \PrincipalEntitlementsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPrincipalEntitlements**](PrincipalEntitlementsAPI.md#GetPrincipalEntitlements) | **Get** /governance/api/v1/principal-entitlements | Retrieve the principal&#39;s effective entitlements for a resource
[**GetPrincipalEntitlementsChanges**](PrincipalEntitlementsAPI.md#GetPrincipalEntitlementsChanges) | **Get** /governance/api/v1/principal-entitlements-changes/{principalEntitlementsChangeId} | Retrieve the principal entitlement changes
[**GetPrincipalEntitlementsHistory**](PrincipalEntitlementsAPI.md#GetPrincipalEntitlementsHistory) | **Get** /governance/api/v1/principal-entitlements/history | Retrieve an entitlement history



## GetPrincipalEntitlements

> PrincipalEntitlementsList GetPrincipalEntitlements(ctx).Filter(filter).Execute()

Retrieve the principal's effective entitlements for a resource



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
	filter := "parent.externalId%20eq%20%220oafxqCAJWWGELFTYASJ%22%20AND%20parent.type%20eq%20%22APPLICATION%22%20AND%20targetPrincipal.externalId%20eq%20%2200ub0oNGTSWTBKOLGLNR%22%20AND%20targetPrincipal.type%20eq%20%22OKTA_USER%22" // string | Apply various filters by using supported principal entitlements filtering properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding) 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrincipalEntitlementsAPI.GetPrincipalEntitlements(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrincipalEntitlementsAPI.GetPrincipalEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrincipalEntitlements`: PrincipalEntitlementsList
	fmt.Fprintf(os.Stdout, "Response from `PrincipalEntitlementsAPI.GetPrincipalEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrincipalEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Apply various filters by using supported principal entitlements filtering properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  | 

### Return type

[**PrincipalEntitlementsList**](PrincipalEntitlementsList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrincipalEntitlementsChanges

> PrincipalEntitlementsChange GetPrincipalEntitlementsChanges(ctx, principalEntitlementsChangeId).Execute()

Retrieve the principal entitlement changes



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
	principalEntitlementsChangeId := "espo3v6yijysPW3ep3q9:espo3v6xlwdtEX2il1d6" // string | The entitlement change ID, in the format of `{old_effective_grant_id}:{new_effective_grant_id}`.  Only the `{new_effective_grant_id}` value is required if the `{old_effective_grant_id}` value isn't available.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrincipalEntitlementsAPI.GetPrincipalEntitlementsChanges(context.Background(), principalEntitlementsChangeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrincipalEntitlementsAPI.GetPrincipalEntitlementsChanges``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrincipalEntitlementsChanges`: PrincipalEntitlementsChange
	fmt.Fprintf(os.Stdout, "Response from `PrincipalEntitlementsAPI.GetPrincipalEntitlementsChanges`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**principalEntitlementsChangeId** | **string** | The entitlement change ID, in the format of &#x60;{old_effective_grant_id}:{new_effective_grant_id}&#x60;.  Only the &#x60;{new_effective_grant_id}&#x60; value is required if the &#x60;{old_effective_grant_id}&#x60; value isn&#39;t available. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrincipalEntitlementsChangesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PrincipalEntitlementsChange**](PrincipalEntitlementsChange.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrincipalEntitlementsHistory

> PrincipalEntitlementsHistory GetPrincipalEntitlementsHistory(ctx).Filter(filter).Limit(limit).After(after).Include(include).Execute()

Retrieve an entitlement history



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
	filter := "resource.externalId eq "0oafxqCAJWWGELFTYASJ" AND resource.type eq "APPLICATION" AND principal.externalId eq "00ub0oNGTSWTBKOLGLNR" AND principal.type eq "OKTA_USER"" // string | This [filter](https://developer.okta.com/docs/api/#filter) expression supports the `eq` [operator](https://developer.okta.com/docs/api/#operators) and the following required and optional properties.  **Required**: You must specify both principal and resource references with one of these sets of properties: * `principal.externalId`, `principal.type`, `resource.externalId`, and `resource.type` * `principalOrn` and `resourceOrn` * `principalId` and `resourceId`  **Optional**: You can optionally filter by a date range with the following properties: * `startDate`: Start of the date range (inclusive) in ISO 8601 UTC format. If omitted, data is retrieved from the earliest available records. * `endDate`: End of the date range (inclusive) in ISO 8601 UTC format. If omitted, data is retrieved up to the current time.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding). 
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	after := "00u68w6vzKLultXS97g6" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. (optional)
	include := []string{"counts"} // []string | An optional parameter that adds additional properties in the `metadata` response object (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrincipalEntitlementsAPI.GetPrincipalEntitlementsHistory(context.Background()).Filter(filter).Limit(limit).After(after).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrincipalEntitlementsAPI.GetPrincipalEntitlementsHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrincipalEntitlementsHistory`: PrincipalEntitlementsHistory
	fmt.Fprintf(os.Stdout, "Response from `PrincipalEntitlementsAPI.GetPrincipalEntitlementsHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrincipalEntitlementsHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | This [filter](https://developer.okta.com/docs/api/#filter) expression supports the &#x60;eq&#x60; [operator](https://developer.okta.com/docs/api/#operators) and the following required and optional properties.  **Required**: You must specify both principal and resource references with one of these sets of properties: * &#x60;principal.externalId&#x60;, &#x60;principal.type&#x60;, &#x60;resource.externalId&#x60;, and &#x60;resource.type&#x60; * &#x60;principalOrn&#x60; and &#x60;resourceOrn&#x60; * &#x60;principalId&#x60; and &#x60;resourceId&#x60;  **Optional**: You can optionally filter by a date range with the following properties: * &#x60;startDate&#x60;: Start of the date range (inclusive) in ISO 8601 UTC format. If omitted, data is retrieved from the earliest available records. * &#x60;endDate&#x60;: End of the date range (inclusive) in ISO 8601 UTC format. If omitted, data is retrieved up to the current time.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. | 
 **include** | **[]string** | An optional parameter that adds additional properties in the &#x60;metadata&#x60; response object | 

### Return type

[**PrincipalEntitlementsHistory**](PrincipalEntitlementsHistory.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

