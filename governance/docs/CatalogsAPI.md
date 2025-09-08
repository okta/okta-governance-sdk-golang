# \CatalogsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCatalogEntryRequestFieldsV2**](CatalogsAPI.md#GetCatalogEntryRequestFieldsV2) | **Get** /governance/api/v2/catalogs/default/entries/{entryId}/users/{userId}/request-fields | Retrieve an entry&#39;s request fields
[**GetCatalogEntryV2**](CatalogsAPI.md#GetCatalogEntryV2) | **Get** /governance/api/v2/catalogs/default/entries/{entryId} | Retrieve a catalog entry
[**ListAllDefaultEntriesV2**](CatalogsAPI.md#ListAllDefaultEntriesV2) | **Get** /governance/api/v2/catalogs/default/entries | List all entries for the default access request catalog
[**ListAllDefaultUserEntriesV2**](CatalogsAPI.md#ListAllDefaultUserEntriesV2) | **Get** /governance/api/v2/catalogs/default/user/{userId}/entries | List all access request catalog entries for a user



## GetCatalogEntryRequestFieldsV2

> CatalogEntryRequestFields GetCatalogEntryRequestFieldsV2(ctx, entryId, userId).Execute()

Retrieve an entry's request fields



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
	entryId := "entryId_example" // string | The ID of the catalog entry
	userId := "00ucvnr9rbONeZdRp1d7" // string | The `id` of the user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetCatalogEntryRequestFieldsV2(context.Background(), entryId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetCatalogEntryRequestFieldsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalogEntryRequestFieldsV2`: CatalogEntryRequestFields
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetCatalogEntryRequestFieldsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryId** | **string** | The ID of the catalog entry | 
**userId** | **string** | The &#x60;id&#x60; of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogEntryRequestFieldsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CatalogEntryRequestFields**](CatalogEntryRequestFields.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCatalogEntryV2

> RcarEntryGet GetCatalogEntryV2(ctx, entryId).Execute()

Retrieve a catalog entry



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
	entryId := "entryId_example" // string | The ID of the catalog entry

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.GetCatalogEntryV2(context.Background(), entryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.GetCatalogEntryV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCatalogEntryV2`: RcarEntryGet
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.GetCatalogEntryV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryId** | **string** | The ID of the catalog entry | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogEntryV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RcarEntryGet**](RcarEntryGet.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllDefaultEntriesV2

> RcarEntriesListV2 ListAllDefaultEntriesV2(ctx).Filter(filter).After(after).Match(match).Limit(limit).Execute()

List all entries for the default access request catalog



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
	filter := "not(parent pr)" // string | A required filter expression that returns entries based on the [`parent`](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Catalogs/#tag/Catalogs/operation/listAllDefaultEntriesV2!c=200&path=data/parent&t=response) property. This [filter](https://developer.okta.com/docs/api/#filter) expression supports the `eq` and `pr` [operators](https://developer.okta.com/docs/api/#operators).  **Note:** Query parameter percent encoding is required. See [Special characters]( https://developer.okta.com/docs/api/#special-characters ). 
	after := "after_example" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous response.  The maximum number of entries returned in a response is determined by the [`limit`](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Catalogs/#tag/Catalogs/operation/listAllDefaultEntriesV2!in=query&path=limit&t=request) query parameter. If there are more entries to return, the `_links.next.href` link contains the `after` cursor for the next page of results. (optional)
	match := "figma" // string | Return catalog entries that match a substring value in the [`name`](https://developer.okta.com/docs/api/iga/openapi/governance.requests.enduser.v2/tag/My-Catalogs/#tag/My-Catalogs/operation/listMyDefaultEntriesV2!c=200&path=data/name&t=response) or [`description`](https://developer.okta.com/docs/api/iga/openapi/governance.requests.enduser.v2/tag/My-Catalogs/#tag/My-Catalogs/operation/listMyDefaultEntriesV2!c=200&path=data/description&t=response) properties. At least three characters are required for fuzzy search. (optional)
	limit := int32(20) // int32 | The maximum number of records returned in a response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.ListAllDefaultEntriesV2(context.Background()).Filter(filter).After(after).Match(match).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.ListAllDefaultEntriesV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllDefaultEntriesV2`: RcarEntriesListV2
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.ListAllDefaultEntriesV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllDefaultEntriesV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A required filter expression that returns entries based on the [&#x60;parent&#x60;](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Catalogs/#tag/Catalogs/operation/listAllDefaultEntriesV2!c&#x3D;200&amp;path&#x3D;data/parent&amp;t&#x3D;response) property. This [filter](https://developer.okta.com/docs/api/#filter) expression supports the &#x60;eq&#x60; and &#x60;pr&#x60; [operators](https://developer.okta.com/docs/api/#operators).  **Note:** Query parameter percent encoding is required. See [Special characters]( https://developer.okta.com/docs/api/#special-characters ).  | 
 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous response.  The maximum number of entries returned in a response is determined by the [&#x60;limit&#x60;](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Catalogs/#tag/Catalogs/operation/listAllDefaultEntriesV2!in&#x3D;query&amp;path&#x3D;limit&amp;t&#x3D;request) query parameter. If there are more entries to return, the &#x60;_links.next.href&#x60; link contains the &#x60;after&#x60; cursor for the next page of results. | 
 **match** | **string** | Return catalog entries that match a substring value in the [&#x60;name&#x60;](https://developer.okta.com/docs/api/iga/openapi/governance.requests.enduser.v2/tag/My-Catalogs/#tag/My-Catalogs/operation/listMyDefaultEntriesV2!c&#x3D;200&amp;path&#x3D;data/name&amp;t&#x3D;response) or [&#x60;description&#x60;](https://developer.okta.com/docs/api/iga/openapi/governance.requests.enduser.v2/tag/My-Catalogs/#tag/My-Catalogs/operation/listMyDefaultEntriesV2!c&#x3D;200&amp;path&#x3D;data/description&amp;t&#x3D;response) properties. At least three characters are required for fuzzy search. | 
 **limit** | **int32** | The maximum number of records returned in a response | 

### Return type

[**RcarEntriesListV2**](RcarEntriesListV2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllDefaultUserEntriesV2

> RcarEntriesListV2 ListAllDefaultUserEntriesV2(ctx, userId).Filter(filter).After(after).Match(match).Limit(limit).Execute()

List all access request catalog entries for a user



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
	userId := "00ucvnr9rbONeZdRp1d7" // string | The `id` of the user
	filter := "not(parent pr)" // string | A required filter expression that returns entries based on the [`parent`](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Catalogs/#tag/Catalogs/operation/listAllDefaultEntriesV2!c=200&path=data/parent&t=response) property. This [filter](https://developer.okta.com/docs/api/#filter) expression supports the `eq` and `pr` [operators](https://developer.okta.com/docs/api/#operators).  **Note:** Query parameter percent encoding is required. See [Special characters]( https://developer.okta.com/docs/api/#special-characters ). 
	after := "after_example" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous response.  The maximum number of entries returned in a response is determined by the [`limit`](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Catalogs/#tag/Catalogs/operation/listAllDefaultEntriesV2!in=query&path=limit&t=request) query parameter. If there are more entries to return, the `_links.next.href` link contains the `after` cursor for the next page of results. (optional)
	match := "figma" // string | Return catalog entries that match a substring value in the [`name`](https://developer.okta.com/docs/api/iga/openapi/governance.requests.enduser.v2/tag/My-Catalogs/#tag/My-Catalogs/operation/listMyDefaultEntriesV2!c=200&path=data/name&t=response) or [`description`](https://developer.okta.com/docs/api/iga/openapi/governance.requests.enduser.v2/tag/My-Catalogs/#tag/My-Catalogs/operation/listMyDefaultEntriesV2!c=200&path=data/description&t=response) properties. At least three characters are required for fuzzy search. (optional)
	limit := int32(20) // int32 | The maximum number of records returned in a response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CatalogsAPI.ListAllDefaultUserEntriesV2(context.Background(), userId).Filter(filter).After(after).Match(match).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CatalogsAPI.ListAllDefaultUserEntriesV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllDefaultUserEntriesV2`: RcarEntriesListV2
	fmt.Fprintf(os.Stdout, "Response from `CatalogsAPI.ListAllDefaultUserEntriesV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | The &#x60;id&#x60; of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllDefaultUserEntriesV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | A required filter expression that returns entries based on the [&#x60;parent&#x60;](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Catalogs/#tag/Catalogs/operation/listAllDefaultEntriesV2!c&#x3D;200&amp;path&#x3D;data/parent&amp;t&#x3D;response) property. This [filter](https://developer.okta.com/docs/api/#filter) expression supports the &#x60;eq&#x60; and &#x60;pr&#x60; [operators](https://developer.okta.com/docs/api/#operators).  **Note:** Query parameter percent encoding is required. See [Special characters]( https://developer.okta.com/docs/api/#special-characters ).  | 
 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous response.  The maximum number of entries returned in a response is determined by the [&#x60;limit&#x60;](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Catalogs/#tag/Catalogs/operation/listAllDefaultEntriesV2!in&#x3D;query&amp;path&#x3D;limit&amp;t&#x3D;request) query parameter. If there are more entries to return, the &#x60;_links.next.href&#x60; link contains the &#x60;after&#x60; cursor for the next page of results. | 
 **match** | **string** | Return catalog entries that match a substring value in the [&#x60;name&#x60;](https://developer.okta.com/docs/api/iga/openapi/governance.requests.enduser.v2/tag/My-Catalogs/#tag/My-Catalogs/operation/listMyDefaultEntriesV2!c&#x3D;200&amp;path&#x3D;data/name&amp;t&#x3D;response) or [&#x60;description&#x60;](https://developer.okta.com/docs/api/iga/openapi/governance.requests.enduser.v2/tag/My-Catalogs/#tag/My-Catalogs/operation/listMyDefaultEntriesV2!c&#x3D;200&amp;path&#x3D;data/description&amp;t&#x3D;response) properties. At least three characters are required for fuzzy search. | 
 **limit** | **int32** | The maximum number of records returned in a response | 

### Return type

[**RcarEntriesListV2**](RcarEntriesListV2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

