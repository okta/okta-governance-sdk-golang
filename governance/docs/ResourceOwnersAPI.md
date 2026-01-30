# \ResourceOwnersAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConfigureResourceOwners**](ResourceOwnersAPI.md#ConfigureResourceOwners) | **Post** /governance/api/v1/resource-owners | Configure the resource owners
[**ListResourceOwnerCatalogResources**](ResourceOwnersAPI.md#ListResourceOwnerCatalogResources) | **Get** /governance/api/v1/resource-owners/catalog/resources | List all resources without owners
[**ListResourceOwners**](ResourceOwnersAPI.md#ListResourceOwners) | **Get** /governance/api/v1/resource-owners | List all resources with owners
[**UpdateResourceOwners**](ResourceOwnersAPI.md#UpdateResourceOwners) | **Patch** /governance/api/v1/resource-owners | Update a resource owner



## ConfigureResourceOwners

> ResourceOwnersResponse ConfigureResourceOwners(ctx).ResourceOwnersUpdatable(resourceOwnersUpdatable).Execute()

Configure the resource owners



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
	resourceOwnersUpdatable := *openapiclient.NewResourceOwnersUpdatable([]string{"ResourceOrns_example"}) // ResourceOwnersUpdatable | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceOwnersAPI.ConfigureResourceOwners(context.Background()).ResourceOwnersUpdatable(resourceOwnersUpdatable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceOwnersAPI.ConfigureResourceOwners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConfigureResourceOwners`: ResourceOwnersResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourceOwnersAPI.ConfigureResourceOwners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiConfigureResourceOwnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceOwnersUpdatable** | [**ResourceOwnersUpdatable**](ResourceOwnersUpdatable.md) |  | 

### Return type

[**ResourceOwnersResponse**](ResourceOwnersResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResourceOwnerCatalogResources

> ResourceOwnersCatalogResourcesResponse ListResourceOwnerCatalogResources(ctx).Filter(filter).Limit(limit).After(after).Execute()

List all resources without owners



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
	filter := "parentResourceOrn eq "orn:okta:idp:00o11edPwGqbUrsDm0g4:apps:salesforce:0oafxqCAJWWGELFTYASJ" AND resource.type eq "entitlement-bundles"" // string | A [filter](https://developer.okta.com/docs/api/#filter) expression that returns entries based on the following properties and supported operators: * `parentResourceOrn`: supports `eq` (required) * `resource.type`: supports `eq` * `resource.profile.*`:  supports `sw` and `co` (both `parentResourceOrn` and `resource.type` filters are required for `resource.profile.*` filtering)  > **Note:** Query parameter percent encoding is required. See [Special characters](https://developer.okta.com/docs/api/#special-characters). 
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	after := "00u68w6vzKLultXS97g6" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceOwnersAPI.ListResourceOwnerCatalogResources(context.Background()).Filter(filter).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceOwnersAPI.ListResourceOwnerCatalogResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListResourceOwnerCatalogResources`: ResourceOwnersCatalogResourcesResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourceOwnersAPI.ListResourceOwnerCatalogResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListResourceOwnerCatalogResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A [filter](https://developer.okta.com/docs/api/#filter) expression that returns entries based on the following properties and supported operators: * &#x60;parentResourceOrn&#x60;: supports &#x60;eq&#x60; (required) * &#x60;resource.type&#x60;: supports &#x60;eq&#x60; * &#x60;resource.profile.*&#x60;:  supports &#x60;sw&#x60; and &#x60;co&#x60; (both &#x60;parentResourceOrn&#x60; and &#x60;resource.type&#x60; filters are required for &#x60;resource.profile.*&#x60; filtering)  &gt; **Note:** Query parameter percent encoding is required. See [Special characters](https://developer.okta.com/docs/api/#special-characters).  | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. | 

### Return type

[**ResourceOwnersCatalogResourcesResponse**](ResourceOwnersCatalogResourcesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListResourceOwners

> ResourceOwnersListResponse ListResourceOwners(ctx).Filter(filter).Limit(limit).After(after).Include(include).Execute()

List all resources with owners



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
	filter := "parentResourceOrn eq "orn:okta:idp:00o11edPwGqbUrsDm0g4:apps:salesforce:0oafxqCAJWWGELFTYASJ"" // string | A [filter](https://developer.okta.com/docs/api/#filter) expression that returns entries based on the following properties and supported operators: * `parentResourceOrn`: supports `eq` * `resource.orn`: supports `eq` * `resource.type`: supports `eq` * `resource.profile.name`:  supports `sw` and `co` (both `parentResourceOrn` and `resource.type` filters are required for `resource.profile.name` filtering)  > **Note:** Query parameter percent encoding is required. See [Special characters](https://developer.okta.com/docs/api/#special-characters). 
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	after := "00u68w6vzKLultXS97g6" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. (optional)
	include := []string{"Include_example"} // []string | Adds additional properties in the response (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ResourceOwnersAPI.ListResourceOwners(context.Background()).Filter(filter).Limit(limit).After(after).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceOwnersAPI.ListResourceOwners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListResourceOwners`: ResourceOwnersListResponse
	fmt.Fprintf(os.Stdout, "Response from `ResourceOwnersAPI.ListResourceOwners`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListResourceOwnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A [filter](https://developer.okta.com/docs/api/#filter) expression that returns entries based on the following properties and supported operators: * &#x60;parentResourceOrn&#x60;: supports &#x60;eq&#x60; * &#x60;resource.orn&#x60;: supports &#x60;eq&#x60; * &#x60;resource.type&#x60;: supports &#x60;eq&#x60; * &#x60;resource.profile.name&#x60;:  supports &#x60;sw&#x60; and &#x60;co&#x60; (both &#x60;parentResourceOrn&#x60; and &#x60;resource.type&#x60; filters are required for &#x60;resource.profile.name&#x60; filtering)  &gt; **Note:** Query parameter percent encoding is required. See [Special characters](https://developer.okta.com/docs/api/#special-characters).  | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. | 
 **include** | **[]string** | Adds additional properties in the response | 

### Return type

[**ResourceOwnersListResponse**](ResourceOwnersListResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceOwners

> UpdateResourceOwners(ctx).ResourceOwnersPatch(resourceOwnersPatch).Execute()

Update a resource owner



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
	resourceOwnersPatch := *openapiclient.NewResourceOwnersPatch("ResourceOrn_example", []openapiclient.ResourceOwnersPatchDataInner{*openapiclient.NewResourceOwnersPatchDataInner()}) // ResourceOwnersPatch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ResourceOwnersAPI.UpdateResourceOwners(context.Background()).ResourceOwnersPatch(resourceOwnersPatch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ResourceOwnersAPI.UpdateResourceOwners``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResourceOwnersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceOwnersPatch** | [**ResourceOwnersPatch**](ResourceOwnersPatch.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

