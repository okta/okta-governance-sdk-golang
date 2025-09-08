# \MyRequestsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMyRequestV2**](MyRequestsAPI.md#CreateMyRequestV2) | **Post** /governance/api/v2/my/catalogs/default/entries/{entryId}/requests | Create a request
[**GetMyRequestV2**](MyRequestsAPI.md#GetMyRequestV2) | **Get** /governance/api/v2/my/catalogs/default/entries/{entryId}/requests/{requestId} | Retrieve my request



## CreateMyRequestV2

> RequestFull2 CreateMyRequestV2(ctx, entryId).MyRequestCreatable(myRequestCreatable).Execute()

Create a request



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
	myRequestCreatable := *openapiclient.NewMyRequestCreatable() // MyRequestCreatable | Creates a resource access request for a given user  You can use this endpoint to create access requests managed by access request conditions.  If `requestedBy` and `requestedFor` are not the same, you must also enable the `requestOnBehalfOfSettings` property on the Access request settings. See [Request Settings](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Request-Settings/#tag/Request-Settings/operation/updateResourceRequestSettingsV2!path=requestOnBehalfOfSettings&t=request).  As part of the payload for the Create a request endpoint, include the following information:  - The Okta user ID for the user who requires access. Add the user ID in the `requestedFor.externalId` parameter. - The Catalog entry ID of the resource required by the user. Add the catalog ID in the `requested.entryId` parameter. - If the request conditions include requester input fields, add the field and information for the field to the `requesterFieldValues` array. See [Retrieve an entry's request fields](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Catalogs/#tag/Catalogs/operation/getCatalogEntryRequestFieldsV2). - Optional: The user ID of the person submitting the request. By default, this value is the admin user ID calling the endpoint and doesn't need to be provided. However, to add a different Okta user ID for the request, include the `requestedBy.externalId` parameter in the request body. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyRequestsAPI.CreateMyRequestV2(context.Background(), entryId).MyRequestCreatable(myRequestCreatable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyRequestsAPI.CreateMyRequestV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMyRequestV2`: RequestFull2
	fmt.Fprintf(os.Stdout, "Response from `MyRequestsAPI.CreateMyRequestV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryId** | **string** | The ID of the catalog entry | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMyRequestV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **myRequestCreatable** | [**MyRequestCreatable**](MyRequestCreatable.md) | Creates a resource access request for a given user  You can use this endpoint to create access requests managed by access request conditions.  If &#x60;requestedBy&#x60; and &#x60;requestedFor&#x60; are not the same, you must also enable the &#x60;requestOnBehalfOfSettings&#x60; property on the Access request settings. See [Request Settings](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Request-Settings/#tag/Request-Settings/operation/updateResourceRequestSettingsV2!path&#x3D;requestOnBehalfOfSettings&amp;t&#x3D;request).  As part of the payload for the Create a request endpoint, include the following information:  - The Okta user ID for the user who requires access. Add the user ID in the &#x60;requestedFor.externalId&#x60; parameter. - The Catalog entry ID of the resource required by the user. Add the catalog ID in the &#x60;requested.entryId&#x60; parameter. - If the request conditions include requester input fields, add the field and information for the field to the &#x60;requesterFieldValues&#x60; array. See [Retrieve an entry&#39;s request fields](https://developer.okta.com/docs/api/iga/openapi/governance.requests.admin.v2/tag/Catalogs/#tag/Catalogs/operation/getCatalogEntryRequestFieldsV2). - Optional: The user ID of the person submitting the request. By default, this value is the admin user ID calling the endpoint and doesn&#39;t need to be provided. However, to add a different Okta user ID for the request, include the &#x60;requestedBy.externalId&#x60; parameter in the request body.  | 

### Return type

[**RequestFull2**](RequestFull2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyRequestV2

> RequestFull2 GetMyRequestV2(ctx, entryId, requestId).Execute()

Retrieve my request



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
	requestId := "requestId_example" // string | The `id` of the request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyRequestsAPI.GetMyRequestV2(context.Background(), entryId, requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyRequestsAPI.GetMyRequestV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyRequestV2`: RequestFull2
	fmt.Fprintf(os.Stdout, "Response from `MyRequestsAPI.GetMyRequestV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entryId** | **string** | The ID of the catalog entry | 
**requestId** | **string** | The &#x60;id&#x60; of the request | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyRequestV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RequestFull2**](RequestFull2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

