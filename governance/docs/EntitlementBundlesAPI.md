# \EntitlementBundlesAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEntitlementBundle**](EntitlementBundlesAPI.md#CreateEntitlementBundle) | **Post** /governance/api/v1/entitlement-bundles | Create an entitlement bundle
[**DeleteEntitlementBundle**](EntitlementBundlesAPI.md#DeleteEntitlementBundle) | **Delete** /governance/api/v1/entitlement-bundles/{entitlementBundleId} | Delete an entitlement bundle
[**GetEntitlementBundle**](EntitlementBundlesAPI.md#GetEntitlementBundle) | **Get** /governance/api/v1/entitlement-bundles/{entitlementBundleId} | Retrieve an entitlement bundle
[**ListEntitlementBundles**](EntitlementBundlesAPI.md#ListEntitlementBundles) | **Get** /governance/api/v1/entitlement-bundles | List all entitlement bundles
[**ReplaceEntitlementBundle**](EntitlementBundlesAPI.md#ReplaceEntitlementBundle) | **Put** /governance/api/v1/entitlement-bundles/{entitlementBundleId} | Replace an entitlement bundle



## CreateEntitlementBundle

> EntitlementBundleFull CreateEntitlementBundle(ctx).EntitlementBundleCreatable(entitlementBundleCreatable).Execute()

Create an entitlement bundle



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
	entitlementBundleCreatable := *openapiclient.NewEntitlementBundleCreatable(*openapiclient.NewTargetResource("ExternalId_example", openapiclient.resource-type-2("APPLICATION")), []openapiclient.EntitlementCreatable{*openapiclient.NewEntitlementCreatable()}, "Name_example") // EntitlementBundleCreatable | The writable attributes of an entitlement bundle

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementBundlesAPI.CreateEntitlementBundle(context.Background()).EntitlementBundleCreatable(entitlementBundleCreatable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementBundlesAPI.CreateEntitlementBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEntitlementBundle`: EntitlementBundleFull
	fmt.Fprintf(os.Stdout, "Response from `EntitlementBundlesAPI.CreateEntitlementBundle`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEntitlementBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entitlementBundleCreatable** | [**EntitlementBundleCreatable**](EntitlementBundleCreatable.md) | The writable attributes of an entitlement bundle | 

### Return type

[**EntitlementBundleFull**](EntitlementBundleFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEntitlementBundle

> DeleteEntitlementBundle(ctx, entitlementBundleId).Execute()

Delete an entitlement bundle



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
	entitlementBundleId := "entitlementBundleId_example" // string | The `id` of the entitlement bundle

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EntitlementBundlesAPI.DeleteEntitlementBundle(context.Background(), entitlementBundleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementBundlesAPI.DeleteEntitlementBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entitlementBundleId** | **string** | The &#x60;id&#x60; of the entitlement bundle | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEntitlementBundleRequest struct via the builder pattern


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


## GetEntitlementBundle

> EntitlementBundleFullWithEntitlements GetEntitlementBundle(ctx, entitlementBundleId).Include(include).Execute()

Retrieve an entitlement bundle



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
	entitlementBundleId := "entitlementBundleId_example" // string | The `id` of the entitlement bundle
	include := []string{"Include_example"} // []string | The `include` filter adds additional properties that are available in the retrieve an entitlement bundle operation, but are omitted from the list response normally. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementBundlesAPI.GetEntitlementBundle(context.Background(), entitlementBundleId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementBundlesAPI.GetEntitlementBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlementBundle`: EntitlementBundleFullWithEntitlements
	fmt.Fprintf(os.Stdout, "Response from `EntitlementBundlesAPI.GetEntitlementBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entitlementBundleId** | **string** | The &#x60;id&#x60; of the entitlement bundle | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | The &#x60;include&#x60; filter adds additional properties that are available in the retrieve an entitlement bundle operation, but are omitted from the list response normally. | 

### Return type

[**EntitlementBundleFullWithEntitlements**](EntitlementBundleFullWithEntitlements.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntitlementBundles

> EntitlementBundlesListWithEntitlements ListEntitlementBundles(ctx).Filter(filter).After(after).Limit(limit).OrderBy(orderBy).Include(include).Execute()

List all entitlement bundles



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
	filter := "lastUpdated gt "2022-05-24T14:15:22Z"" // string | A [filter](https://developer.okta.com/docs/api/#filter) expression that returns entries based on the following properties and supported operators: * `id`: supports `gt` * `lastUpdated`: supports `gt`, `ge`, `le`, and `lt` * `targetResourceOrn`: supports `eq` * `target.externalId`:  supports `eq` * `target.type`:  supports `eq` * `status`:  supports `eq` * `name`:  supports `eq` and `co`  > **Note:** Query parameter percent encoding is required. See [Special characters](https://developer.okta.com/docs/api/#special-characters).  (optional)
	after := "after_example" // string | The after cursor provided by a prior request. (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	orderBy := "created%20desc" // string | Apply an ordering of entitlement-bundles by specifying a supported entitlement bundle property name with `%20asc` or `%20desc` suffix.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  (optional)
	include := []string{"Include_example"} // []string | The `include` filter adds additional properties that are available in the retrieve an entitlement bundle operation, but are omitted from the list response normally. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementBundlesAPI.ListEntitlementBundles(context.Background()).Filter(filter).After(after).Limit(limit).OrderBy(orderBy).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementBundlesAPI.ListEntitlementBundles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEntitlementBundles`: EntitlementBundlesListWithEntitlements
	fmt.Fprintf(os.Stdout, "Response from `EntitlementBundlesAPI.ListEntitlementBundles`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEntitlementBundlesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A [filter](https://developer.okta.com/docs/api/#filter) expression that returns entries based on the following properties and supported operators: * &#x60;id&#x60;: supports &#x60;gt&#x60; * &#x60;lastUpdated&#x60;: supports &#x60;gt&#x60;, &#x60;ge&#x60;, &#x60;le&#x60;, and &#x60;lt&#x60; * &#x60;targetResourceOrn&#x60;: supports &#x60;eq&#x60; * &#x60;target.externalId&#x60;:  supports &#x60;eq&#x60; * &#x60;target.type&#x60;:  supports &#x60;eq&#x60; * &#x60;status&#x60;:  supports &#x60;eq&#x60; * &#x60;name&#x60;:  supports &#x60;eq&#x60; and &#x60;co&#x60;  &gt; **Note:** Query parameter percent encoding is required. See [Special characters](https://developer.okta.com/docs/api/#special-characters).  | 
 **after** | **string** | The after cursor provided by a prior request. | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **orderBy** | **string** | Apply an ordering of entitlement-bundles by specifying a supported entitlement bundle property name with &#x60;%20asc&#x60; or &#x60;%20desc&#x60; suffix.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  | 
 **include** | **[]string** | The &#x60;include&#x60; filter adds additional properties that are available in the retrieve an entitlement bundle operation, but are omitted from the list response normally. | 

### Return type

[**EntitlementBundlesListWithEntitlements**](EntitlementBundlesListWithEntitlements.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceEntitlementBundle

> EntitlementBundleFull ReplaceEntitlementBundle(ctx, entitlementBundleId).EntitlementBundleUpdatable(entitlementBundleUpdatable).Execute()

Replace an entitlement bundle



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
	entitlementBundleId := "entitlementBundleId_example" // string | The `id` of the entitlement bundle
	entitlementBundleUpdatable := *openapiclient.NewEntitlementBundleUpdatable("Id_example", "TargetResourceOrn_example", *openapiclient.NewTargetResource("ExternalId_example", openapiclient.resource-type-2("APPLICATION")), "Name_example") // EntitlementBundleUpdatable | The writable attributes of an entitlement bundle

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementBundlesAPI.ReplaceEntitlementBundle(context.Background(), entitlementBundleId).EntitlementBundleUpdatable(entitlementBundleUpdatable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementBundlesAPI.ReplaceEntitlementBundle``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceEntitlementBundle`: EntitlementBundleFull
	fmt.Fprintf(os.Stdout, "Response from `EntitlementBundlesAPI.ReplaceEntitlementBundle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entitlementBundleId** | **string** | The &#x60;id&#x60; of the entitlement bundle | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceEntitlementBundleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entitlementBundleUpdatable** | [**EntitlementBundleUpdatable**](EntitlementBundleUpdatable.md) | The writable attributes of an entitlement bundle | 

### Return type

[**EntitlementBundleFull**](EntitlementBundleFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

