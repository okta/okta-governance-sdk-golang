# \CollectionsV2API

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddResourcesToCollectionV2**](CollectionsV2API.md#AddResourcesToCollectionV2) | **Post** /governance/api/v2/collections/{collectionId}/resources | Add the resources to a collection
[**AssignCollectionV2**](CollectionsV2API.md#AssignCollectionV2) | **Post** /governance/api/v2/collections/{collectionId}/assignments | Assign a collection to principals
[**CreateCollectionV2**](CollectionsV2API.md#CreateCollectionV2) | **Post** /governance/api/v2/collections | Create a resource collection
[**DeleteCollectionResourceV2**](CollectionsV2API.md#DeleteCollectionResourceV2) | **Delete** /governance/api/v2/collections/{collectionId}/resources/{resourceId} | Delete a collection resource
[**DeleteCollectionV2**](CollectionsV2API.md#DeleteCollectionV2) | **Delete** /governance/api/v2/collections/{collectionId} | Delete a collection
[**DeletePrincipalAssignmentV2**](CollectionsV2API.md#DeletePrincipalAssignmentV2) | **Delete** /governance/api/v2/collections/{collectionId}/assignments/{assignmentId} | Delete a principal assignment
[**GetCollectionResourceV2**](CollectionsV2API.md#GetCollectionResourceV2) | **Get** /governance/api/v2/collections/{collectionId}/resources/{resourceId} | Retrieve a collection resource
[**GetCollectionV2**](CollectionsV2API.md#GetCollectionV2) | **Get** /governance/api/v2/collections/{collectionId} | Retrieve a resource collection
[**GetUnassignedUsersV2**](CollectionsV2API.md#GetUnassignedUsersV2) | **Get** /governance/api/v2/collections/{collectionId}/catalog/users | Retrieve the unassigned users
[**ListCollectionAssignmentsV2**](CollectionsV2API.md#ListCollectionAssignmentsV2) | **Get** /governance/api/v2/collections/{collectionId}/assignments | List all assignments for a collection
[**ListCollectionResourceEntitlementsV2**](CollectionsV2API.md#ListCollectionResourceEntitlementsV2) | **Get** /governance/api/v2/collections/{collectionId}/resources/{resourceId}/entitlements | List all entitlements for a collection resource
[**ListCollectionResourcesV2**](CollectionsV2API.md#ListCollectionResourcesV2) | **Get** /governance/api/v2/collections/{collectionId}/resources | List all collection resources
[**ListCollectionsAssignmentsV2**](CollectionsV2API.md#ListCollectionsAssignmentsV2) | **Get** /governance/api/v2/collections/assignments | List all assignments for all collections
[**ListCollectionsV2**](CollectionsV2API.md#ListCollectionsV2) | **Get** /governance/api/v2/collections | List all resource collections
[**ReplaceCollectionResourceV2**](CollectionsV2API.md#ReplaceCollectionResourceV2) | **Put** /governance/api/v2/collections/{collectionId}/resources/{resourceId} | Replace a collection resource
[**ReplaceCollectionV2**](CollectionsV2API.md#ReplaceCollectionV2) | **Put** /governance/api/v2/collections/{collectionId} | Replace a resource collection
[**RetrieveCollectionAssignmentOriginsV2**](CollectionsV2API.md#RetrieveCollectionAssignmentOriginsV2) | **Post** /governance/api/v2/collections/assignments/origins | Retrieve the collection origins for a batch of assignments
[**UpdatePrincipalAssignmentV2**](CollectionsV2API.md#UpdatePrincipalAssignmentV2) | **Patch** /governance/api/v2/collections/{collectionId}/assignments/{assignmentId} | Update a principal assignment



## AddResourcesToCollectionV2

> CollectionResourcesListV2 AddResourcesToCollectionV2(ctx, collectionId).CollectionResourceCreatableV2(collectionResourceCreatableV2).Execute()

Add the resources to a collection



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
	collectionId := "collectionId_example" // string | Unique identifier for the collection
	collectionResourceCreatableV2 := []openapiclient.CollectionResourceCreatableV2{*openapiclient.NewCollectionResourceCreatableV2("ResourceOrn_example")} // []CollectionResourceCreatableV2 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsV2API.AddResourcesToCollectionV2(context.Background(), collectionId).CollectionResourceCreatableV2(collectionResourceCreatableV2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsV2API.AddResourcesToCollectionV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddResourcesToCollectionV2`: CollectionResourcesListV2
	fmt.Fprintf(os.Stdout, "Response from `CollectionsV2API.AddResourcesToCollectionV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Unique identifier for the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddResourcesToCollectionV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **collectionResourceCreatableV2** | [**[]CollectionResourceCreatableV2**](CollectionResourceCreatableV2.md) |  | 

### Return type

[**CollectionResourcesListV2**](CollectionResourcesListV2.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignCollectionV2

> []AssignedPrincipalFull AssignCollectionV2(ctx, collectionId).AssignedPrincipal(assignedPrincipal).Execute()

Assign a collection to principals



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
	collectionId := "collectionId_example" // string | Unique identifier for the collection
	assignedPrincipal := []openapiclient.AssignedPrincipal{*openapiclient.NewAssignedPrincipal()} // []AssignedPrincipal | Attributes of collection assignment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsV2API.AssignCollectionV2(context.Background(), collectionId).AssignedPrincipal(assignedPrincipal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsV2API.AssignCollectionV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignCollectionV2`: []AssignedPrincipalFull
	fmt.Fprintf(os.Stdout, "Response from `CollectionsV2API.AssignCollectionV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Unique identifier for the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignCollectionV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **assignedPrincipal** | [**[]AssignedPrincipal**](AssignedPrincipal.md) | Attributes of collection assignment | 

### Return type

[**[]AssignedPrincipalFull**](AssignedPrincipalFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCollectionV2

> CollectionFullV2 CreateCollectionV2(ctx).CollectionCreatable(collectionCreatable).Execute()

Create a resource collection



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
	collectionCreatable := *openapiclient.NewCollectionCreatable("Name_example") // CollectionCreatable | The writeable attributes of a resource collection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsV2API.CreateCollectionV2(context.Background()).CollectionCreatable(collectionCreatable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsV2API.CreateCollectionV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCollectionV2`: CollectionFullV2
	fmt.Fprintf(os.Stdout, "Response from `CollectionsV2API.CreateCollectionV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCollectionV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collectionCreatable** | [**CollectionCreatable**](CollectionCreatable.md) | The writeable attributes of a resource collection | 

### Return type

[**CollectionFullV2**](CollectionFullV2.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCollectionResourceV2

> DeleteCollectionResourceV2(ctx, collectionId, resourceId).Execute()

Delete a collection resource



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
	collectionId := "collectionId_example" // string | Unique identifier for the collection
	resourceId := "resourceId_example" // string | Unique identifier for the resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionsV2API.DeleteCollectionResourceV2(context.Background(), collectionId, resourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsV2API.DeleteCollectionResourceV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Unique identifier for the collection | 
**resourceId** | **string** | Unique identifier for the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCollectionResourceV2Request struct via the builder pattern


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


## DeleteCollectionV2

> DeleteCollectionV2(ctx, collectionId).Execute()

Delete a collection



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
	collectionId := "collectionId_example" // string | Unique identifier for the collection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionsV2API.DeleteCollectionV2(context.Background(), collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsV2API.DeleteCollectionV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Unique identifier for the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCollectionV2Request struct via the builder pattern


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


## DeletePrincipalAssignmentV2

> DeletePrincipalAssignmentV2(ctx, collectionId, assignmentId).Execute()

Delete a principal assignment



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
	collectionId := "collectionId_example" // string | Unique identifier for the collection
	assignmentId := "assignmentId_example" // string | Unique identifier for the collection assignment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionsV2API.DeletePrincipalAssignmentV2(context.Background(), collectionId, assignmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsV2API.DeletePrincipalAssignmentV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Unique identifier for the collection | 
**assignmentId** | **string** | Unique identifier for the collection assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrincipalAssignmentV2Request struct via the builder pattern


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


## GetCollectionResourceV2

> CollectionResourceFullV2 GetCollectionResourceV2(ctx, collectionId, resourceId).Include(include).Execute()

Retrieve a collection resource



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
	collectionId := "collectionId_example" // string | Unique identifier for the collection
	resourceId := "resourceId_example" // string | Unique identifier for the resource
	include := []string{"Include_example"} // []string | The `include` query parameter returns additional properties in the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsV2API.GetCollectionResourceV2(context.Background(), collectionId, resourceId).Include(include).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsV2API.GetCollectionResourceV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCollectionResourceV2`: CollectionResourceFullV2
	fmt.Fprintf(os.Stdout, "Response from `CollectionsV2API.GetCollectionResourceV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Unique identifier for the collection | 
**resourceId** | **string** | Unique identifier for the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionResourceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **include** | **[]string** | The &#x60;include&#x60; query parameter returns additional properties in the response. | 

### Return type

[**CollectionResourceFullV2**](CollectionResourceFullV2.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionV2

> CollectionFullV2 GetCollectionV2(ctx, collectionId).Execute()

Retrieve a resource collection



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
	collectionId := "collectionId_example" // string | Unique identifier for the collection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsV2API.GetCollectionV2(context.Background(), collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsV2API.GetCollectionV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCollectionV2`: CollectionFullV2
	fmt.Fprintf(os.Stdout, "Response from `CollectionsV2API.GetCollectionV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Unique identifier for the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionFullV2**](CollectionFullV2.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnassignedUsersV2

> ListCatalogUsers GetUnassignedUsersV2(ctx, collectionId).Filter(filter).Execute()

Retrieve the unassigned users



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
	collectionId := "collectionId_example" // string | Unique identifier for the collection
	filter := "firstName sw "John"" // string | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters users based on `firstName`, `lastName`, or `email` properties. The `sw` [operator](https://developer.okta.com/docs/api/#operators) is supported for these properties.  > **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsV2API.GetUnassignedUsersV2(context.Background(), collectionId).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsV2API.GetUnassignedUsersV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUnassignedUsersV2`: ListCatalogUsers
	fmt.Fprintf(os.Stdout, "Response from `CollectionsV2API.GetUnassignedUsersV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Unique identifier for the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnassignedUsersV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters users based on &#x60;firstName&#x60;, &#x60;lastName&#x60;, or &#x60;email&#x60; properties. The &#x60;sw&#x60; [operator](https://developer.okta.com/docs/api/#operators) is supported for these properties.  &gt; **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  | 

### Return type

[**ListCatalogUsers**](ListCatalogUsers.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollectionAssignmentsV2

> ListAssignedPrincipals ListCollectionAssignmentsV2(ctx, collectionId).Filter(filter).Limit(limit).After(after).Execute()

List all assignments for a collection



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
	collectionId := "collectionId_example" // string | Unique identifier for the collection
	filter := "principal.externalId eq "00ub0oNGTSWTBKOLGLNR" AND principal.type eq "OKTA_USER"" // string | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters principals based on `principal` or `principalProfile` properties. The `sw` [operator](https://developer.okta.com/docs/api/#operators) is supported for these properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	after := "00u68w6vzKLultXS97g6" // string | Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the standard link headers. See [pagination](https://developer.okta.com/docs/api/#pagination). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsV2API.ListCollectionAssignmentsV2(context.Background(), collectionId).Filter(filter).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsV2API.ListCollectionAssignmentsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCollectionAssignmentsV2`: ListAssignedPrincipals
	fmt.Fprintf(os.Stdout, "Response from `CollectionsV2API.ListCollectionAssignmentsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Unique identifier for the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionAssignmentsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters principals based on &#x60;principal&#x60; or &#x60;principalProfile&#x60; properties. The &#x60;sw&#x60; [operator](https://developer.okta.com/docs/api/#operators) is supported for these properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **after** | **string** | Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the standard link headers. See [pagination](https://developer.okta.com/docs/api/#pagination). | 

### Return type

[**ListAssignedPrincipals**](ListAssignedPrincipals.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollectionResourceEntitlementsV2

> CollectionResourceEntitlementsList ListCollectionResourceEntitlementsV2(ctx, collectionId, resourceId).Limit(limit).After(after).Execute()

List all entitlements for a collection resource



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
	collectionId := "collectionId_example" // string | Unique identifier for the collection
	resourceId := "resourceId_example" // string | Unique identifier for the resource
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	after := "00u68w6vzKLultXS97g6" // string | Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the standard link headers. See [pagination](https://developer.okta.com/docs/api/#pagination). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsV2API.ListCollectionResourceEntitlementsV2(context.Background(), collectionId, resourceId).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsV2API.ListCollectionResourceEntitlementsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCollectionResourceEntitlementsV2`: CollectionResourceEntitlementsList
	fmt.Fprintf(os.Stdout, "Response from `CollectionsV2API.ListCollectionResourceEntitlementsV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Unique identifier for the collection | 
**resourceId** | **string** | Unique identifier for the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionResourceEntitlementsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **after** | **string** | Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the standard link headers. See [pagination](https://developer.okta.com/docs/api/#pagination). | 

### Return type

[**CollectionResourceEntitlementsList**](CollectionResourceEntitlementsList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollectionResourcesV2

> CollectionResourcesListV2 ListCollectionResourcesV2(ctx, collectionId).Include(include).Limit(limit).After(after).Execute()

List all collection resources



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
	collectionId := "collectionId_example" // string | Unique identifier for the collection
	include := []string{"Include_example"} // []string | The `include` query parameter returns additional properties in the response. (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	after := "00u68w6vzKLultXS97g6" // string | Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the standard link headers. See [pagination](https://developer.okta.com/docs/api/#pagination). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsV2API.ListCollectionResourcesV2(context.Background(), collectionId).Include(include).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsV2API.ListCollectionResourcesV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCollectionResourcesV2`: CollectionResourcesListV2
	fmt.Fprintf(os.Stdout, "Response from `CollectionsV2API.ListCollectionResourcesV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Unique identifier for the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionResourcesV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | The &#x60;include&#x60; query parameter returns additional properties in the response. | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **after** | **string** | Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the standard link headers. See [pagination](https://developer.okta.com/docs/api/#pagination). | 

### Return type

[**CollectionResourcesListV2**](CollectionResourcesListV2.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollectionsAssignmentsV2

> ListAssignedPrincipals ListCollectionsAssignmentsV2(ctx).Filter(filter).Limit(limit).After(after).Execute()

List all assignments for all collections



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
	filter := "principal.externalId eq "00ub0oNGTSWTBKOLGLNR" AND principal.type eq "OKTA_USER"" // string | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters principals based on `principal` or `principalProfile` properties. The `sw` [operator](https://developer.okta.com/docs/api/#operators) is supported for these properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	after := "00u68w6vzKLultXS97g6" // string | Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the standard link headers. See [pagination](https://developer.okta.com/docs/api/#pagination). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsV2API.ListCollectionsAssignmentsV2(context.Background()).Filter(filter).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsV2API.ListCollectionsAssignmentsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCollectionsAssignmentsV2`: ListAssignedPrincipals
	fmt.Fprintf(os.Stdout, "Response from `CollectionsV2API.ListCollectionsAssignmentsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionsAssignmentsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters principals based on &#x60;principal&#x60; or &#x60;principalProfile&#x60; properties. The &#x60;sw&#x60; [operator](https://developer.okta.com/docs/api/#operators) is supported for these properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **after** | **string** | Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the standard link headers. See [pagination](https://developer.okta.com/docs/api/#pagination). | 

### Return type

[**ListAssignedPrincipals**](ListAssignedPrincipals.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollectionsV2

> CollectionsListV2 ListCollectionsV2(ctx).Include(include).Limit(limit).After(after).Filter(filter).Execute()

List all resource collections



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
	include := []string{"Include_example"} // []string | The `include` filter adds additional properties in the response. (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	after := "00u68w6vzKLultXS97g6" // string | Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the standard link headers. See [pagination](https://developer.okta.com/docs/api/#pagination). (optional)
	filter := "name sw "Sales"" // string | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters data based on `id`, `name`, `resourceOrn`, or `principalOrn` resource collection properties. The `sw` and `co` [operators](https://developer.okta.com/docs/api/#operators) are supported for `name`, and the `eq` operator is supported for `id`, `resourceOrn`, and `principalOrn`.  When filtering by `resourceOrn`, Okta returns only collections that contain the specified resource. Each collection in the response includes a `resourceRelationship` object with details about when the resource was added.  When filtering by `principalOrn`, Okta returns only collections assigned to the specified principal. Each collection in the response includes a `principalAssignment` object with assignment details.  > **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsV2API.ListCollectionsV2(context.Background()).Include(include).Limit(limit).After(after).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsV2API.ListCollectionsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCollectionsV2`: CollectionsListV2
	fmt.Fprintf(os.Stdout, "Response from `CollectionsV2API.ListCollectionsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | The &#x60;include&#x60; filter adds additional properties in the response. | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **after** | **string** | Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the standard link headers. See [pagination](https://developer.okta.com/docs/api/#pagination). | 
 **filter** | **string** | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters data based on &#x60;id&#x60;, &#x60;name&#x60;, &#x60;resourceOrn&#x60;, or &#x60;principalOrn&#x60; resource collection properties. The &#x60;sw&#x60; and &#x60;co&#x60; [operators](https://developer.okta.com/docs/api/#operators) are supported for &#x60;name&#x60;, and the &#x60;eq&#x60; operator is supported for &#x60;id&#x60;, &#x60;resourceOrn&#x60;, and &#x60;principalOrn&#x60;.  When filtering by &#x60;resourceOrn&#x60;, Okta returns only collections that contain the specified resource. Each collection in the response includes a &#x60;resourceRelationship&#x60; object with details about when the resource was added.  When filtering by &#x60;principalOrn&#x60;, Okta returns only collections assigned to the specified principal. Each collection in the response includes a &#x60;principalAssignment&#x60; object with assignment details.  &gt; **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  | 

### Return type

[**CollectionsListV2**](CollectionsListV2.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceCollectionResourceV2

> CollectionResourceAppWithEntitlements ReplaceCollectionResourceV2(ctx, collectionId, resourceId).CollectionResourceUpdatable(collectionResourceUpdatable).Execute()

Replace a collection resource



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
	collectionId := "collectionId_example" // string | Unique identifier for the collection
	resourceId := "resourceId_example" // string | Unique identifier for the resource
	collectionResourceUpdatable := *openapiclient.NewCollectionResourceUpdatable() // CollectionResourceUpdatable | The updatable attributes of a collection resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsV2API.ReplaceCollectionResourceV2(context.Background(), collectionId, resourceId).CollectionResourceUpdatable(collectionResourceUpdatable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsV2API.ReplaceCollectionResourceV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceCollectionResourceV2`: CollectionResourceAppWithEntitlements
	fmt.Fprintf(os.Stdout, "Response from `CollectionsV2API.ReplaceCollectionResourceV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Unique identifier for the collection | 
**resourceId** | **string** | Unique identifier for the resource | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceCollectionResourceV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **collectionResourceUpdatable** | [**CollectionResourceUpdatable**](CollectionResourceUpdatable.md) | The updatable attributes of a collection resource | 

### Return type

[**CollectionResourceAppWithEntitlements**](CollectionResourceAppWithEntitlements.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceCollectionV2

> CollectionFullV2 ReplaceCollectionV2(ctx, collectionId).CollectionUpdatable(collectionUpdatable).Execute()

Replace a resource collection



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
	collectionId := "collectionId_example" // string | Unique identifier for the collection
	collectionUpdatable := *openapiclient.NewCollectionUpdatable("Name_example") // CollectionUpdatable | The writeable attributes of a resource collection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsV2API.ReplaceCollectionV2(context.Background(), collectionId).CollectionUpdatable(collectionUpdatable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsV2API.ReplaceCollectionV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceCollectionV2`: CollectionFullV2
	fmt.Fprintf(os.Stdout, "Response from `CollectionsV2API.ReplaceCollectionV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Unique identifier for the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceCollectionV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **collectionUpdatable** | [**CollectionUpdatable**](CollectionUpdatable.md) | The writeable attributes of a resource collection | 

### Return type

[**CollectionFullV2**](CollectionFullV2.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveCollectionAssignmentOriginsV2

> CollectionAssignmentOriginsResponse RetrieveCollectionAssignmentOriginsV2(ctx).CollectionAssignmentOriginsRequest(collectionAssignmentOriginsRequest).Execute()

Retrieve the collection origins for a batch of assignments



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
	collectionAssignmentOriginsRequest := *openapiclient.NewCollectionAssignmentOriginsRequest() // CollectionAssignmentOriginsRequest | The assignments to look up, anchored on a principal or a resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsV2API.RetrieveCollectionAssignmentOriginsV2(context.Background()).CollectionAssignmentOriginsRequest(collectionAssignmentOriginsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsV2API.RetrieveCollectionAssignmentOriginsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveCollectionAssignmentOriginsV2`: CollectionAssignmentOriginsResponse
	fmt.Fprintf(os.Stdout, "Response from `CollectionsV2API.RetrieveCollectionAssignmentOriginsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveCollectionAssignmentOriginsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collectionAssignmentOriginsRequest** | [**CollectionAssignmentOriginsRequest**](CollectionAssignmentOriginsRequest.md) | The assignments to look up, anchored on a principal or a resource | 

### Return type

[**CollectionAssignmentOriginsResponse**](CollectionAssignmentOriginsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrincipalAssignmentV2

> UpdatePrincipalAssignmentV2(ctx, collectionId, assignmentId).AssignmentPatchOperation(assignmentPatchOperation).Execute()

Update a principal assignment



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
	collectionId := "collectionId_example" // string | Unique identifier for the collection
	assignmentId := "assignmentId_example" // string | Unique identifier for the collection assignment
	assignmentPatchOperation := []openapiclient.AssignmentPatchOperation{*openapiclient.NewAssignmentPatchOperation(openapiclient.assignment-patch-op("ADD"), "Path_example")} // []AssignmentPatchOperation | The writable attributes of a collection assignment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionsV2API.UpdatePrincipalAssignmentV2(context.Background(), collectionId, assignmentId).AssignmentPatchOperation(assignmentPatchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsV2API.UpdatePrincipalAssignmentV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | Unique identifier for the collection | 
**assignmentId** | **string** | Unique identifier for the collection assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePrincipalAssignmentV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **assignmentPatchOperation** | [**[]AssignmentPatchOperation**](AssignmentPatchOperation.md) | The writable attributes of a collection assignment | 

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

