# \MyCatalogsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMyCatalogEntryRequestFieldsV2**](MyCatalogsAPI.md#GetMyCatalogEntryRequestFieldsV2) | **Get** /governance/api/v2/my/catalogs/default/entries/{entryId}/request-fields | Retrieve the request fields for my catalog entry
[**GetMyCatalogEntryUserRequestFieldsV2**](MyCatalogsAPI.md#GetMyCatalogEntryUserRequestFieldsV2) | **Get** /governance/api/v2/my/catalogs/default/entries/{entryId}/users/{userId}/request-fields | Retrieve the entry request fields for a user
[**GetMyEntryV2**](MyCatalogsAPI.md#GetMyEntryV2) | **Get** /governance/api/v2/my/catalogs/default/entries/{entryId} | Retrieve my catalog entry
[**ListMyDefaultEntriesV2**](MyCatalogsAPI.md#ListMyDefaultEntriesV2) | **Get** /governance/api/v2/my/catalogs/default/entries | List my entries for the default access request catalog
[**ListMyEntryUsersV2**](MyCatalogsAPI.md#ListMyEntryUsersV2) | **Get** /governance/api/v2/my/catalogs/default/entries/{entryId}/users | List my catalog entry users



## GetMyCatalogEntryRequestFieldsV2

> CatalogEntryRequestFields GetMyCatalogEntryRequestFieldsV2(ctx, entryId).Execute()

Retrieve the request fields for my catalog entry



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
	resp, r, err := apiClient.MyCatalogsAPI.GetMyCatalogEntryRequestFieldsV2(context.Background(), entryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyCatalogsAPI.GetMyCatalogEntryRequestFieldsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyCatalogEntryRequestFieldsV2`: CatalogEntryRequestFields
	fmt.Fprintf(os.Stdout, "Response from `MyCatalogsAPI.GetMyCatalogEntryRequestFieldsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryId** | **string** | The ID of the catalog entry | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyCatalogEntryRequestFieldsV2Request struct via the builder pattern


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


## GetMyCatalogEntryUserRequestFieldsV2

> CatalogEntryRequestFields GetMyCatalogEntryUserRequestFieldsV2(ctx, entryId, userId).Execute()

Retrieve the entry request fields for a user



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
	resp, r, err := apiClient.MyCatalogsAPI.GetMyCatalogEntryUserRequestFieldsV2(context.Background(), entryId, userId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyCatalogsAPI.GetMyCatalogEntryUserRequestFieldsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyCatalogEntryUserRequestFieldsV2`: CatalogEntryRequestFields
	fmt.Fprintf(os.Stdout, "Response from `MyCatalogsAPI.GetMyCatalogEntryUserRequestFieldsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryId** | **string** | The ID of the catalog entry | 
**userId** | **string** | The &#x60;id&#x60; of the user | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyCatalogEntryUserRequestFieldsV2Request struct via the builder pattern


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


## GetMyEntryV2

> RcarEntryGet GetMyEntryV2(ctx, entryId).Execute()

Retrieve my catalog entry



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
	resp, r, err := apiClient.MyCatalogsAPI.GetMyEntryV2(context.Background(), entryId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyCatalogsAPI.GetMyEntryV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyEntryV2`: RcarEntryGet
	fmt.Fprintf(os.Stdout, "Response from `MyCatalogsAPI.GetMyEntryV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryId** | **string** | The ID of the catalog entry | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyEntryV2Request struct via the builder pattern


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


## ListMyDefaultEntriesV2

> RcarEntriesListV2 ListMyDefaultEntriesV2(ctx).Filter(filter).After(after).Match(match).Limit(limit).Execute()

List my entries for the default access request catalog



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
	filter := "not(parent pr)" // string | A required [filter](https://developer.okta.com/docs/api/#filter) expression that returns entries based on the [`parent`](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Catalogs/#tag/Catalogs/operation/listAllDefaultEntriesV2!c=200&path=data/parent&t=response) property: * This filter expression only supports the `parent` property and the `eq` and `pr` [operators](https://developer.okta.com/docs/api/#operators). * If you want the query to return child entries, then you must specify the `parent` ID with the `eq` operator.  > **Notes:** > * If you don't use the `parent` property in the filter expression, undesireable results are returned. > * Query parameter percent encoding is required. See [Special characters]( https://developer.okta.com/docs/api/#special-characters ). 
	after := "after_example" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous response.  The maximum number of entries returned in a response is determined by the [`limit`](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Catalogs/#tag/Catalogs/operation/listAllDefaultEntriesV2!in=query&path=limit&t=request) query parameter. If there are more entries to return, the `_links.next.href` link contains the `after` cursor for the next page of results. (optional)
	match := "figma" // string | Return catalog entries that match a substring value in the [`name`](https://developer.okta.com/docs/api/iga/openapi/governance.requests.enduser.v2/tag/My-Catalogs/#tag/My-Catalogs/operation/listMyDefaultEntriesV2!c=200&path=data/name&t=response) or [`description`](https://developer.okta.com/docs/api/iga/openapi/governance.requests.enduser.v2/tag/My-Catalogs/#tag/My-Catalogs/operation/listMyDefaultEntriesV2!c=200&path=data/description&t=response) properties. At least three characters are required for fuzzy search. (optional)
	limit := int32(20) // int32 | The maximum number of records returned in a response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyCatalogsAPI.ListMyDefaultEntriesV2(context.Background()).Filter(filter).After(after).Match(match).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyCatalogsAPI.ListMyDefaultEntriesV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMyDefaultEntriesV2`: RcarEntriesListV2
	fmt.Fprintf(os.Stdout, "Response from `MyCatalogsAPI.ListMyDefaultEntriesV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMyDefaultEntriesV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A required [filter](https://developer.okta.com/docs/api/#filter) expression that returns entries based on the [&#x60;parent&#x60;](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Catalogs/#tag/Catalogs/operation/listAllDefaultEntriesV2!c&#x3D;200&amp;path&#x3D;data/parent&amp;t&#x3D;response) property: * This filter expression only supports the &#x60;parent&#x60; property and the &#x60;eq&#x60; and &#x60;pr&#x60; [operators](https://developer.okta.com/docs/api/#operators). * If you want the query to return child entries, then you must specify the &#x60;parent&#x60; ID with the &#x60;eq&#x60; operator.  &gt; **Notes:** &gt; * If you don&#39;t use the &#x60;parent&#x60; property in the filter expression, undesireable results are returned. &gt; * Query parameter percent encoding is required. See [Special characters]( https://developer.okta.com/docs/api/#special-characters ).  | 
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


## ListMyEntryUsersV2

> EntryUsersList ListMyEntryUsersV2(ctx, entryId).Filter(filter).After(after).Limit(limit).Execute()

List my catalog entry users



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
	filter := "firstName sw "John" OR lastName sw "John"" // string | A required filter expression that returns users based on the `firstName` or `lastName` properties. This [filter](https://developer.okta.com/docs/api/#filter) expression supports the `sw` [operator](https://developer.okta.com/docs/api/#operators).  **Note:** Query parameter percent encoding is required. See [Special characters]( https://developer.okta.com/docs/api/#special-characters ). 
	after := "00u68w6vzKLultXS97g6" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyCatalogsAPI.ListMyEntryUsersV2(context.Background(), entryId).Filter(filter).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyCatalogsAPI.ListMyEntryUsersV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMyEntryUsersV2`: EntryUsersList
	fmt.Fprintf(os.Stdout, "Response from `MyCatalogsAPI.ListMyEntryUsersV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryId** | **string** | The ID of the catalog entry | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMyEntryUsersV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | A required filter expression that returns users based on the &#x60;firstName&#x60; or &#x60;lastName&#x60; properties. This [filter](https://developer.okta.com/docs/api/#filter) expression supports the &#x60;sw&#x60; [operator](https://developer.okta.com/docs/api/#operators).  **Note:** Query parameter percent encoding is required. See [Special characters]( https://developer.okta.com/docs/api/#special-characters ).  | 
 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]

### Return type

[**EntryUsersList**](EntryUsersList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

