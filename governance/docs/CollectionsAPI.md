# \CollectionsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddResourcesToCollection**](CollectionsAPI.md#AddResourcesToCollection) | **Post** /governance/api/v1/collections/{collectionId}/resources | Add the resources to a collection
[**AssignCollection**](CollectionsAPI.md#AssignCollection) | **Post** /governance/api/v1/collections/{collectionId}/assignments | Assign a collection to principals
[**CreateCollection**](CollectionsAPI.md#CreateCollection) | **Post** /governance/api/v1/collections | Create a resource collection
[**DeleteCollection**](CollectionsAPI.md#DeleteCollection) | **Delete** /governance/api/v1/collections/{collectionId} | Delete a collection
[**DeleteCollectionResource**](CollectionsAPI.md#DeleteCollectionResource) | **Delete** /governance/api/v1/collections/{collectionId}/resources/{resourceId} | Delete a collection resource
[**DeletePrincipalAssignment**](CollectionsAPI.md#DeletePrincipalAssignment) | **Delete** /governance/api/v1/collections/{collectionId}/assignments/{assignmentId} | Delete a principal assignment
[**GetCollection**](CollectionsAPI.md#GetCollection) | **Get** /governance/api/v1/collections/{collectionId} | Retrieve a resource collection
[**GetCollectionResource**](CollectionsAPI.md#GetCollectionResource) | **Get** /governance/api/v1/collections/{collectionId}/resources/{resourceId} | Retrieve a collection resource
[**GetUnassignedUsers**](CollectionsAPI.md#GetUnassignedUsers) | **Get** /governance/api/v1/collections/{collectionId}/catalog/users | Retrieve the unassigned users
[**ListCollectionAssignments**](CollectionsAPI.md#ListCollectionAssignments) | **Get** /governance/api/v1/collections/{collectionId}/assignments | List all assignments for a collection
[**ListCollectionResources**](CollectionsAPI.md#ListCollectionResources) | **Get** /governance/api/v1/collections/{collectionId}/resources | List all collection resources
[**ListCollections**](CollectionsAPI.md#ListCollections) | **Get** /governance/api/v1/collections | List all resource collections
[**ListCollectionsAssignments**](CollectionsAPI.md#ListCollectionsAssignments) | **Get** /governance/api/v1/collections/assignments | List all assignments for all collections
[**ReplaceCollection**](CollectionsAPI.md#ReplaceCollection) | **Put** /governance/api/v1/collections/{collectionId} | Replace a resource collection
[**ReplaceCollectionResource**](CollectionsAPI.md#ReplaceCollectionResource) | **Put** /governance/api/v1/collections/{collectionId}/resources/{resourceId} | Replace a collection resource
[**UpdatePrincipalAssignment**](CollectionsAPI.md#UpdatePrincipalAssignment) | **Patch** /governance/api/v1/collections/{collectionId}/assignments/{assignmentId} | Update a principal assignment



## AddResourcesToCollection

> CollectionResourcesList AddResourcesToCollection(ctx, collectionId).CollectionResourceCreatable(collectionResourceCreatable).Execute()

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
	collectionId := "collectionId_example" // string | The `id` of the collection
	collectionResourceCreatable := []openapiclient.CollectionResourceCreatable{*openapiclient.NewCollectionResourceCreatable("ResourceOrn_example")} // []CollectionResourceCreatable |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.AddResourcesToCollection(context.Background(), collectionId).CollectionResourceCreatable(collectionResourceCreatable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.AddResourcesToCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddResourcesToCollection`: CollectionResourcesList
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.AddResourcesToCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The &#x60;id&#x60; of the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddResourcesToCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **collectionResourceCreatable** | [**[]CollectionResourceCreatable**](CollectionResourceCreatable.md) |  | 

### Return type

[**CollectionResourcesList**](CollectionResourcesList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignCollection

> []AssignedPrincipalFull AssignCollection(ctx, collectionId).AssignedPrincipal(assignedPrincipal).Execute()

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
	collectionId := "collectionId_example" // string | The `id` of the collection
	assignedPrincipal := []openapiclient.AssignedPrincipal{*openapiclient.NewAssignedPrincipal()} // []AssignedPrincipal | Attributes of collection assignment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.AssignCollection(context.Background(), collectionId).AssignedPrincipal(assignedPrincipal).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.AssignCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignCollection`: []AssignedPrincipalFull
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.AssignCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The &#x60;id&#x60; of the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignCollectionRequest struct via the builder pattern


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


## CreateCollection

> CollectionFull CreateCollection(ctx).CollectionCreatable(collectionCreatable).Execute()

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
	resp, r, err := apiClient.CollectionsAPI.CreateCollection(context.Background()).CollectionCreatable(collectionCreatable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.CreateCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCollection`: CollectionFull
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.CreateCollection`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collectionCreatable** | [**CollectionCreatable**](CollectionCreatable.md) | The writeable attributes of a resource collection | 

### Return type

[**CollectionFull**](CollectionFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCollection

> DeleteCollection(ctx, collectionId).Execute()

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
	collectionId := "collectionId_example" // string | The `id` of the collection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionsAPI.DeleteCollection(context.Background(), collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.DeleteCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The &#x60;id&#x60; of the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCollectionRequest struct via the builder pattern


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


## DeleteCollectionResource

> DeleteCollectionResource(ctx, collectionId, resourceId).Execute()

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
	collectionId := "collectionId_example" // string | The `id` of the collection
	resourceId := "resourceId_example" // string | The `id` of the resource in ORN format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionsAPI.DeleteCollectionResource(context.Background(), collectionId, resourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.DeleteCollectionResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The &#x60;id&#x60; of the collection | 
**resourceId** | **string** | The &#x60;id&#x60; of the resource in ORN format | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCollectionResourceRequest struct via the builder pattern


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


## DeletePrincipalAssignment

> DeletePrincipalAssignment(ctx, collectionId, assignmentId).Execute()

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
	collectionId := "collectionId_example" // string | The `id` of the collection
	assignmentId := "assignmentId_example" // string | The `id` of the collection assignment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionsAPI.DeletePrincipalAssignment(context.Background(), collectionId, assignmentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.DeletePrincipalAssignment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The &#x60;id&#x60; of the collection | 
**assignmentId** | **string** | The &#x60;id&#x60; of the collection assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrincipalAssignmentRequest struct via the builder pattern


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


## GetCollection

> CollectionFull GetCollection(ctx, collectionId).Execute()

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
	collectionId := "collectionId_example" // string | The `id` of the collection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.GetCollection(context.Background(), collectionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.GetCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCollection`: CollectionFull
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.GetCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The &#x60;id&#x60; of the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CollectionFull**](CollectionFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCollectionResource

> CollectionResourceFull GetCollectionResource(ctx, collectionId, resourceId).Execute()

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
	collectionId := "collectionId_example" // string | The `id` of the collection
	resourceId := "resourceId_example" // string | The `id` of the resource in ORN format

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.GetCollectionResource(context.Background(), collectionId, resourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.GetCollectionResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCollectionResource`: CollectionResourceFull
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.GetCollectionResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The &#x60;id&#x60; of the collection | 
**resourceId** | **string** | The &#x60;id&#x60; of the resource in ORN format | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCollectionResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CollectionResourceFull**](CollectionResourceFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUnassignedUsers

> ListCatalogUsers GetUnassignedUsers(ctx, collectionId).Filter(filter).Execute()

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
	collectionId := "collectionId_example" // string | The `id` of the collection
	filter := "firstName%20sw%20%22John%22" // string | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters users based on `firstName`, `lastName`, or `email` properties. The `sw` [operator](https://developer.okta.com/docs/api/#operators) is supported for these properties.  > **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.GetUnassignedUsers(context.Background(), collectionId).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.GetUnassignedUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUnassignedUsers`: ListCatalogUsers
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.GetUnassignedUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The &#x60;id&#x60; of the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUnassignedUsersRequest struct via the builder pattern


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


## ListCollectionAssignments

> ListAssignedPrincipals ListCollectionAssignments(ctx, collectionId).Filter(filter).Limit(limit).After(after).Execute()

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
	collectionId := "collectionId_example" // string | The `id` of the collection
	filter := "principal.externalId%20eq%20%2200ub0oNGTSWTBKOLGLNR%22%20AND%20principal.type%20eq%20%22OKTA_USER%22" // string | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters principals based on `principal` or `principalProfile` properties. The `sw` [operator](https://developer.okta.com/docs/api/#operators) is supported for these properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	after := "00u68w6vzKLultXS97g6" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.ListCollectionAssignments(context.Background(), collectionId).Filter(filter).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.ListCollectionAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCollectionAssignments`: ListAssignedPrincipals
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.ListCollectionAssignments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The &#x60;id&#x60; of the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters principals based on &#x60;principal&#x60; or &#x60;principalProfile&#x60; properties. The &#x60;sw&#x60; [operator](https://developer.okta.com/docs/api/#operators) is supported for these properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. | 

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


## ListCollectionResources

> CollectionResourcesList ListCollectionResources(ctx, collectionId).Include(include).Limit(limit).After(after).Execute()

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
	collectionId := "collectionId_example" // string | The `id` of the collection
	include := []string{"Include_example"} // []string | The `include` query parameter returns additional properties in the response. (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	after := "00u68w6vzKLultXS97g6" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.ListCollectionResources(context.Background(), collectionId).Include(include).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.ListCollectionResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCollectionResources`: CollectionResourcesList
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.ListCollectionResources`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The &#x60;id&#x60; of the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **include** | **[]string** | The &#x60;include&#x60; query parameter returns additional properties in the response. | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. | 

### Return type

[**CollectionResourcesList**](CollectionResourcesList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollections

> CollectionsList ListCollections(ctx).Include(include).Limit(limit).After(after).Filter(filter).Execute()

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
	after := "00u68w6vzKLultXS97g6" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. (optional)
	filter := "name sw "Sales"" // string | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters data based on `id` or `name` resource collection properties. The `sw` and `co` [operators](https://developer.okta.com/docs/api/#operators) are supported for `name`, and the `eq` operator is supported for `id`.  > **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.ListCollections(context.Background()).Include(include).Limit(limit).After(after).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.ListCollections``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCollections`: CollectionsList
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.ListCollections`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **include** | **[]string** | The &#x60;include&#x60; filter adds additional properties in the response. | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. | 
 **filter** | **string** | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters data based on &#x60;id&#x60; or &#x60;name&#x60; resource collection properties. The &#x60;sw&#x60; and &#x60;co&#x60; [operators](https://developer.okta.com/docs/api/#operators) are supported for &#x60;name&#x60;, and the &#x60;eq&#x60; operator is supported for &#x60;id&#x60;.  &gt; **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  | 

### Return type

[**CollectionsList**](CollectionsList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCollectionsAssignments

> ListAssignedPrincipals ListCollectionsAssignments(ctx).Filter(filter).Limit(limit).After(after).Execute()

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
	filter := "principal.externalId%20eq%20%2200ub0oNGTSWTBKOLGLNR%22%20AND%20principal.type%20eq%20%22OKTA_USER%22" // string | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters principals based on `principal` or `principalProfile` properties. The `sw` [operator](https://developer.okta.com/docs/api/#operators) is supported for these properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	after := "00u68w6vzKLultXS97g6" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.ListCollectionsAssignments(context.Background()).Filter(filter).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.ListCollectionsAssignments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCollectionsAssignments`: ListAssignedPrincipals
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.ListCollectionsAssignments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCollectionsAssignmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters principals based on &#x60;principal&#x60; or &#x60;principalProfile&#x60; properties. The &#x60;sw&#x60; [operator](https://developer.okta.com/docs/api/#operators) is supported for these properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. | 

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


## ReplaceCollection

> CollectionFull ReplaceCollection(ctx, collectionId).CollectionUpdatable(collectionUpdatable).Execute()

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
	collectionId := "collectionId_example" // string | The `id` of the collection
	collectionUpdatable := *openapiclient.NewCollectionUpdatable("Name_example") // CollectionUpdatable | The writeable attributes of a resource collection

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.ReplaceCollection(context.Background(), collectionId).CollectionUpdatable(collectionUpdatable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.ReplaceCollection``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceCollection`: CollectionFull
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.ReplaceCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The &#x60;id&#x60; of the collection | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **collectionUpdatable** | [**CollectionUpdatable**](CollectionUpdatable.md) | The writeable attributes of a resource collection | 

### Return type

[**CollectionFull**](CollectionFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceCollectionResource

> CollectionResourceFull ReplaceCollectionResource(ctx, collectionId, resourceId).CollectionResourceUpdatable(collectionResourceUpdatable).Execute()

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
	collectionId := "collectionId_example" // string | The `id` of the collection
	resourceId := "resourceId_example" // string | The `id` of the resource in ORN format
	collectionResourceUpdatable := *openapiclient.NewCollectionResourceUpdatable() // CollectionResourceUpdatable | The updatable attributes of a collection resource

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CollectionsAPI.ReplaceCollectionResource(context.Background(), collectionId, resourceId).CollectionResourceUpdatable(collectionResourceUpdatable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.ReplaceCollectionResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceCollectionResource`: CollectionResourceFull
	fmt.Fprintf(os.Stdout, "Response from `CollectionsAPI.ReplaceCollectionResource`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The &#x60;id&#x60; of the collection | 
**resourceId** | **string** | The &#x60;id&#x60; of the resource in ORN format | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceCollectionResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **collectionResourceUpdatable** | [**CollectionResourceUpdatable**](CollectionResourceUpdatable.md) | The updatable attributes of a collection resource | 

### Return type

[**CollectionResourceFull**](CollectionResourceFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrincipalAssignment

> UpdatePrincipalAssignment(ctx, collectionId, assignmentId).AssignmentPatchOperation(assignmentPatchOperation).Execute()

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
	collectionId := "collectionId_example" // string | The `id` of the collection
	assignmentId := "assignmentId_example" // string | The `id` of the collection assignment
	assignmentPatchOperation := []openapiclient.AssignmentPatchOperation{*openapiclient.NewAssignmentPatchOperation(openapiclient.assignment-patch-op("ADD"), "Path_example")} // []AssignmentPatchOperation | The writable attributes of a collection assignment

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CollectionsAPI.UpdatePrincipalAssignment(context.Background(), collectionId, assignmentId).AssignmentPatchOperation(assignmentPatchOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CollectionsAPI.UpdatePrincipalAssignment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** | The &#x60;id&#x60; of the collection | 
**assignmentId** | **string** | The &#x60;id&#x60; of the collection assignment | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePrincipalAssignmentRequest struct via the builder pattern


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

