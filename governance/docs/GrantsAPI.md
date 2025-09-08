# \GrantsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGrant**](GrantsAPI.md#CreateGrant) | **Post** /governance/api/v1/grants | Create a grant
[**GetGrant**](GrantsAPI.md#GetGrant) | **Get** /governance/api/v1/grants/{grantId} | Retrieve a grant
[**ListGrants**](GrantsAPI.md#ListGrants) | **Get** /governance/api/v1/grants | List all grants
[**ReplaceGrant**](GrantsAPI.md#ReplaceGrant) | **Put** /governance/api/v1/grants/{grantId} | Replace a grant
[**UpdateGrant**](GrantsAPI.md#UpdateGrant) | **Patch** /governance/api/v1/grants/{grantId} | Update a grant



## CreateGrant

> GrantFull CreateGrant(ctx).GrantCreatable(grantCreatable).Execute()

Create a grant



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
	grantCreatable := openapiclient.grant-creatable{GrantTypeBundleWriteable: openapiclient.NewGrantTypeBundleWriteable("GrantType_example", "EntitlementBundleId_example", *openapiclient.NewTargetPrincipal("00ub0oNGTSWTBKOLGLNR", openapiclient.principal-type("OKTA_USER")))} // GrantCreatable | The grant request parameters depend on the selected `grantType`

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GrantsAPI.CreateGrant(context.Background()).GrantCreatable(grantCreatable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GrantsAPI.CreateGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGrant`: GrantFull
	fmt.Fprintf(os.Stdout, "Response from `GrantsAPI.CreateGrant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **grantCreatable** | [**GrantCreatable**](GrantCreatable.md) | The grant request parameters depend on the selected &#x60;grantType&#x60; | 

### Return type

[**GrantFull**](GrantFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGrant

> GetGrant200Response GetGrant(ctx, grantId).Include(include).Execute()

Retrieve a grant



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
	grantId := "grantId_example" // string | The `id` of the grant
	include := []string{"Include_example"} // []string | The `include` parameter adds additional properties to the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GrantsAPI.GetGrant(context.Background(), grantId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GrantsAPI.GetGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGrant`: GetGrant200Response
	fmt.Fprintf(os.Stdout, "Response from `GrantsAPI.GetGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**grantId** | **string** | The &#x60;id&#x60; of the grant | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | The &#x60;include&#x60; parameter adds additional properties to the response. | 

### Return type

[**GetGrant200Response**](GetGrant200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGrants

> ListGrants200Response ListGrants(ctx).Filter(filter).After(after).Limit(limit).Include(include).Execute()

List all grants



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
	filter := "target.externalId eq "0oafxqCAJWWGELFTYASJ" AND target.type eq "APPLICATION" AND targetPrincipal.externalId eq "00ub0oNGTSWTBKOLGLNR" AND targetPrincipal.type eq "OKTA_USER"" // string | A [filter](https://developer.okta.com/docs/api/#filter) expression that returns entries based on the following properties: * `targetResourceOrn` (alternatively, you can use `target.externalId` and `target.type` for a specific resource ) * `targetPrincipalOrn` (alternatively, you can use `targetPrincipal.externalId` and `targetPrincipal.type` for a specific principal) * `entitlementBundleId` * `entitlements.id` * `entitlements.values.id` * `action`  The `eq` operator is supported for these properties. The `AND` and `OR` logical operators are supported for combining multiple expressions.  > **Note:** Query parameter percent encoding is required. See [Special characters](https://developer.okta.com/docs/api/#special-characters). 
	after := "after_example" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	include := []string{"Include_example"} // []string | The `include` parameter adds additional properties to the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GrantsAPI.ListGrants(context.Background()).Filter(filter).After(after).Limit(limit).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GrantsAPI.ListGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListGrants`: ListGrants200Response
	fmt.Fprintf(os.Stdout, "Response from `GrantsAPI.ListGrants`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A [filter](https://developer.okta.com/docs/api/#filter) expression that returns entries based on the following properties: * &#x60;targetResourceOrn&#x60; (alternatively, you can use &#x60;target.externalId&#x60; and &#x60;target.type&#x60; for a specific resource ) * &#x60;targetPrincipalOrn&#x60; (alternatively, you can use &#x60;targetPrincipal.externalId&#x60; and &#x60;targetPrincipal.type&#x60; for a specific principal) * &#x60;entitlementBundleId&#x60; * &#x60;entitlements.id&#x60; * &#x60;entitlements.values.id&#x60; * &#x60;action&#x60;  The &#x60;eq&#x60; operator is supported for these properties. The &#x60;AND&#x60; and &#x60;OR&#x60; logical operators are supported for combining multiple expressions.  &gt; **Note:** Query parameter percent encoding is required. See [Special characters](https://developer.okta.com/docs/api/#special-characters).  | 
 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **include** | **[]string** | The &#x60;include&#x60; parameter adds additional properties to the response. | 

### Return type

[**ListGrants200Response**](ListGrants200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceGrant

> GrantFull ReplaceGrant(ctx, grantId).GrantFull(grantFull).Execute()

Replace a grant



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/okta/okta-governance-sdk-golang"
)

func main() {
	grantId := "grantId_example" // string | The `id` of the grant
	grantFull := *openapiclient.NewGrantFull(openapiclient.grant-type("CUSTOM"), "orn:okta:directory:00o8rk36Bp5eZKOrw0g4:users:00u1ktfFMZ5HNoj7k0g4", *openapiclient.NewTargetPrincipalFull("00ub0oNGTSWTBKOLGLNR", openapiclient.principal-type("OKTA_USER")), openapiclient.grant-action("ALLOW"), openapiclient.grant-actor("API"), "TargetResourceOrn_example", *openapiclient.NewTargetResource("ExternalId_example", openapiclient.resource-type-2("APPLICATION")), *openapiclient.NewResourceGrantLinks(*openapiclient.NewLink("Href_example")), openapiclient.grant-status("ACTIVE"), "Id_example", "CreatedBy_example", time.Now(), time.Now(), "LastUpdatedBy_example") // GrantFull | The grant request parameters depend on the selected `grantType`

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GrantsAPI.ReplaceGrant(context.Background(), grantId).GrantFull(grantFull).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GrantsAPI.ReplaceGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceGrant`: GrantFull
	fmt.Fprintf(os.Stdout, "Response from `GrantsAPI.ReplaceGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**grantId** | **string** | The &#x60;id&#x60; of the grant | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **grantFull** | [**GrantFull**](GrantFull.md) | The grant request parameters depend on the selected &#x60;grantType&#x60; | 

### Return type

[**GrantFull**](GrantFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGrant

> GrantFull UpdateGrant(ctx, grantId).GrantPatch(grantPatch).Execute()

Update a grant



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
	grantId := "grantId_example" // string | The `id` of the grant
	grantPatch := *openapiclient.NewGrantPatch("Id_example", *openapiclient.NewScheduleSettingsWriteable()) // GrantPatch | Request parameters for a grant expiration date update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GrantsAPI.UpdateGrant(context.Background(), grantId).GrantPatch(grantPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GrantsAPI.UpdateGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGrant`: GrantFull
	fmt.Fprintf(os.Stdout, "Response from `GrantsAPI.UpdateGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**grantId** | **string** | The &#x60;id&#x60; of the grant | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **grantPatch** | [**GrantPatch**](GrantPatch.md) | Request parameters for a grant expiration date update | 

### Return type

[**GrantFull**](GrantFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

