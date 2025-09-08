# \EntitlementsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEntitlement**](EntitlementsAPI.md#CreateEntitlement) | **Post** /governance/api/v1/entitlements | Create an entitlement
[**DeleteEntitlement**](EntitlementsAPI.md#DeleteEntitlement) | **Delete** /governance/api/v1/entitlements/{entitlementId} | Delete an entitlement
[**GetEntitlement**](EntitlementsAPI.md#GetEntitlement) | **Get** /governance/api/v1/entitlements/{entitlementId} | Retrieve an entitlement
[**GetEntitlementValue**](EntitlementsAPI.md#GetEntitlementValue) | **Get** /governance/api/v1/entitlements/{entitlementId}/values/{valueId} | Retrieve an entitlement value
[**ListAllEntitlementValues**](EntitlementsAPI.md#ListAllEntitlementValues) | **Get** /governance/api/v1/entitlements/values | List all entitlement values
[**ListEntitlementValues**](EntitlementsAPI.md#ListEntitlementValues) | **Get** /governance/api/v1/entitlements/{entitlementId}/values | List all values for an entitlement
[**ListEntitlements**](EntitlementsAPI.md#ListEntitlements) | **Get** /governance/api/v1/entitlements | List all entitlements
[**ReplaceEntitlement**](EntitlementsAPI.md#ReplaceEntitlement) | **Put** /governance/api/v1/entitlements/{entitlementId} | Replace an entitlement
[**UpdateEntitlement**](EntitlementsAPI.md#UpdateEntitlement) | **Patch** /governance/api/v1/entitlements/{entitlementId} | Update the entitlement



## CreateEntitlement

> EntitlementsFullWithParent CreateEntitlement(ctx).EntitlementCreate(entitlementCreate).Execute()

Create an entitlement



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
	entitlementCreate := *openapiclient.NewEntitlementCreate("Name_example", "ExternalValue_example", false, openapiclient.entitlement-property-datatype("string")) // EntitlementCreate | The writable attributes of an entitlement

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.CreateEntitlement(context.Background()).EntitlementCreate(entitlementCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.CreateEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEntitlement`: EntitlementsFullWithParent
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.CreateEntitlement`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **entitlementCreate** | [**EntitlementCreate**](EntitlementCreate.md) | The writable attributes of an entitlement | 

### Return type

[**EntitlementsFullWithParent**](EntitlementsFullWithParent.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteEntitlement

> DeleteEntitlement(ctx, entitlementId).Execute()

Delete an entitlement



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
	entitlementId := "entitlementId_example" // string | The `id` of the entitlement

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EntitlementsAPI.DeleteEntitlement(context.Background(), entitlementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.DeleteEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entitlementId** | **string** | The &#x60;id&#x60; of the entitlement | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEntitlementRequest struct via the builder pattern


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


## GetEntitlement

> EntitlementsFullWithParent GetEntitlement(ctx, entitlementId).Execute()

Retrieve an entitlement



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
	entitlementId := "entitlementId_example" // string | The `id` of the entitlement

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.GetEntitlement(context.Background(), entitlementId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.GetEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlement`: EntitlementsFullWithParent
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.GetEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entitlementId** | **string** | The &#x60;id&#x60; of the entitlement | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EntitlementsFullWithParent**](EntitlementsFullWithParent.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEntitlementValue

> EntitlementValue2 GetEntitlementValue(ctx, entitlementId, valueId).Execute()

Retrieve an entitlement value



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
	entitlementId := "entitlementId_example" // string | The `id` of the entitlement
	valueId := "valueId_example" // string | The `id` of the entitlement value

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.GetEntitlementValue(context.Background(), entitlementId, valueId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.GetEntitlementValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEntitlementValue`: EntitlementValue2
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.GetEntitlementValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entitlementId** | **string** | The &#x60;id&#x60; of the entitlement | 
**valueId** | **string** | The &#x60;id&#x60; of the entitlement value | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEntitlementValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EntitlementValue2**](EntitlementValue2.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllEntitlementValues

> EntitlementValuesList ListAllEntitlementValues(ctx).Filter(filter).Limit(limit).After(after).OrderBy(orderBy).Execute()

List all entitlement values



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
	filter := "parent.externalId eq "0oafxqCAJWWGELFTYASJ" AND parent.type eq "APPLICATION"" // string | A [filter](https://developer.okta.com/docs/api/#filter) expression that returns entries based on the following properties and supported operators: * `parent.externalId`: supports `eq` * `parent.type`: supports `eq` * `parentResourceOrn`: supports `eq` * `name`:  supports `sw` and `co` * `externalValue`: supports `eq`  > **Note:** Query parameter percent encoding is required. See [Special characters](https://developer.okta.com/docs/api/#special-characters). 
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 200)
	after := "00u68w6vzKLultXS97g6" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. (optional)
	orderBy := []string{"Inner_example"} // []string | A field by which results can be sorted. For now, sorting by a single field is supported.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  (optional) (default to ["id asc"])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.ListAllEntitlementValues(context.Background()).Filter(filter).Limit(limit).After(after).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.ListAllEntitlementValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllEntitlementValues`: EntitlementValuesList
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.ListAllEntitlementValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllEntitlementValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A [filter](https://developer.okta.com/docs/api/#filter) expression that returns entries based on the following properties and supported operators: * &#x60;parent.externalId&#x60;: supports &#x60;eq&#x60; * &#x60;parent.type&#x60;: supports &#x60;eq&#x60; * &#x60;parentResourceOrn&#x60;: supports &#x60;eq&#x60; * &#x60;name&#x60;:  supports &#x60;sw&#x60; and &#x60;co&#x60; * &#x60;externalValue&#x60;: supports &#x60;eq&#x60;  &gt; **Note:** Query parameter percent encoding is required. See [Special characters](https://developer.okta.com/docs/api/#special-characters).  | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 200]
 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. | 
 **orderBy** | **[]string** | A field by which results can be sorted. For now, sorting by a single field is supported.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  | [default to [&quot;id asc&quot;]]

### Return type

[**EntitlementValuesList**](EntitlementValuesList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntitlementValues

> EntitlementValuesList ListEntitlementValues(ctx, entitlementId).Limit(limit).After(after).Filter(filter).OrderBy(orderBy).Execute()

List all values for an entitlement



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
	entitlementId := "entitlementId_example" // string | The `id` of the entitlement
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	after := "00u68w6vzKLultXS97g6" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. (optional)
	filter := "name sw "License"" // string | A [filter](https://developer.okta.com/docs/api/#filter) expression that returns entries based on the following properties and supported operators: * `name`:supports the `sw` and `co` operators. * `labelValueId`: supports `eq`  > **Note:** Query parameter percent encoding is required. See [Special characters](https://developer.okta.com/docs/api/#special-characters).  (optional)
	orderBy := []string{"Inner_example"} // []string | A field by which results can be sorted. For now, sorting by a single field is supported.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  (optional) (default to ["name asc"])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.ListEntitlementValues(context.Background(), entitlementId).Limit(limit).After(after).Filter(filter).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.ListEntitlementValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEntitlementValues`: EntitlementValuesList
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.ListEntitlementValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entitlementId** | **string** | The &#x60;id&#x60; of the entitlement | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEntitlementValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. | 
 **filter** | **string** | A [filter](https://developer.okta.com/docs/api/#filter) expression that returns entries based on the following properties and supported operators: * &#x60;name&#x60;:supports the &#x60;sw&#x60; and &#x60;co&#x60; operators. * &#x60;labelValueId&#x60;: supports &#x60;eq&#x60;  &gt; **Note:** Query parameter percent encoding is required. See [Special characters](https://developer.okta.com/docs/api/#special-characters).  | 
 **orderBy** | **[]string** | A field by which results can be sorted. For now, sorting by a single field is supported.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  | [default to [&quot;name asc&quot;]]

### Return type

[**EntitlementValuesList**](EntitlementValuesList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListEntitlements

> EntitlementsList ListEntitlements(ctx).Filter(filter).Limit(limit).After(after).OrderBy(orderBy).Execute()

List all entitlements



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
	filter := "parent.externalId eq "0oafxqCAJWWGELFTYASJ" AND parent.type eq "APPLICATION"" // string | A [filter](https://developer.okta.com/docs/api/#filter) expression that returns entries based on the following properties and supported operators: * `parent.externalId`: supports `eq` * `parent.type`: supports `eq` * `parentResourceOrn`: supports `eq` * `name`:  supports `sw` and `co`  > **Note:** Query parameter percent encoding is required. See [Special characters](https://developer.okta.com/docs/api/#special-characters). 
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	after := "00u68w6vzKLultXS97g6" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. (optional)
	orderBy := []string{"Inner_example"} // []string | A field by which results can be sorted. For now, sorting by a single field is supported.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  (optional) (default to ["name asc"])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.ListEntitlements(context.Background()).Filter(filter).Limit(limit).After(after).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.ListEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEntitlements`: EntitlementsList
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.ListEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A [filter](https://developer.okta.com/docs/api/#filter) expression that returns entries based on the following properties and supported operators: * &#x60;parent.externalId&#x60;: supports &#x60;eq&#x60; * &#x60;parent.type&#x60;: supports &#x60;eq&#x60; * &#x60;parentResourceOrn&#x60;: supports &#x60;eq&#x60; * &#x60;name&#x60;:  supports &#x60;sw&#x60; and &#x60;co&#x60;  &gt; **Note:** Query parameter percent encoding is required. See [Special characters](https://developer.okta.com/docs/api/#special-characters).  | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. | 
 **orderBy** | **[]string** | A field by which results can be sorted. For now, sorting by a single field is supported.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  | [default to [&quot;name asc&quot;]]

### Return type

[**EntitlementsList**](EntitlementsList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceEntitlement

> EntitlementsFullWithParent ReplaceEntitlement(ctx, entitlementId).EntitlementsFullWithParent(entitlementsFullWithParent).Execute()

Replace an entitlement



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
	entitlementId := "entitlementId_example" // string | The `id` of the entitlement
	entitlementsFullWithParent := *openapiclient.NewEntitlementsFullWithParent("ParentResourceOrn_example", *openapiclient.NewTargetResource("ExternalId_example", openapiclient.resource-type-2("APPLICATION")), []openapiclient.EntitlementValueFull{*openapiclient.NewEntitlementValueFull()}, "Id_example", "Name_example", "ExternalValue_example", false, openapiclient.entitlement-property-datatype("string")) // EntitlementsFullWithParent | The writable attributes of an entitlement

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.ReplaceEntitlement(context.Background(), entitlementId).EntitlementsFullWithParent(entitlementsFullWithParent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.ReplaceEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceEntitlement`: EntitlementsFullWithParent
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.ReplaceEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entitlementId** | **string** | The &#x60;id&#x60; of the entitlement | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entitlementsFullWithParent** | [**EntitlementsFullWithParent**](EntitlementsFullWithParent.md) | The writable attributes of an entitlement | 

### Return type

[**EntitlementsFullWithParent**](EntitlementsFullWithParent.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEntitlement

> EntitlementsFullWithParent UpdateEntitlement(ctx, entitlementId).EntitlementPatchInner(entitlementPatchInner).Execute()

Update the entitlement



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
	entitlementId := "entitlementId_example" // string | The `id` of the entitlement
	entitlementPatchInner := []openapiclient.EntitlementPatchInner{openapiclient.entitlement_patch_inner{EntitlementPatchOperation: openapiclient.NewEntitlementPatchOperation(openapiclient.entitlement-patch-op("ADD"), "Path_example", "RefType_example")}} // []EntitlementPatchInner | The writable attributes of an entitlement.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EntitlementsAPI.UpdateEntitlement(context.Background(), entitlementId).EntitlementPatchInner(entitlementPatchInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EntitlementsAPI.UpdateEntitlement``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEntitlement`: EntitlementsFullWithParent
	fmt.Fprintf(os.Stdout, "Response from `EntitlementsAPI.UpdateEntitlement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entitlementId** | **string** | The &#x60;id&#x60; of the entitlement | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEntitlementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **entitlementPatchInner** | [**[]EntitlementPatchInner**](EntitlementPatchInner.md) | The writable attributes of an entitlement. | 

### Return type

[**EntitlementsFullWithParent**](EntitlementsFullWithParent.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

