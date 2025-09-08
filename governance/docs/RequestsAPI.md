# \RequestsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRequest**](RequestsAPI.md#CreateRequest) | **Post** /governance/api/v1/requests | Create a request
[**CreateRequestMessage**](RequestsAPI.md#CreateRequestMessage) | **Post** /governance/api/v1/requests/{requestId}/messages | Create a Message for a Request
[**CreateRequestMessageV2**](RequestsAPI.md#CreateRequestMessageV2) | **Post** /governance/api/v2/requests/{requestId}/messages | Create a request message
[**CreateRequestV2**](RequestsAPI.md#CreateRequestV2) | **Post** /governance/api/v2/requests | Create a request
[**GetRequest**](RequestsAPI.md#GetRequest) | **Get** /governance/api/v1/requests/{requestId} | Retrieve a request
[**GetRequestV2**](RequestsAPI.md#GetRequestV2) | **Get** /governance/api/v2/requests/{requestId} | Retrieve a request
[**ListAllRequests**](RequestsAPI.md#ListAllRequests) | **Get** /governance/api/v1/requests | List all requests
[**ListAllRequestsV2**](RequestsAPI.md#ListAllRequestsV2) | **Get** /governance/api/v2/requests | List all requests



## CreateRequest

> RequestFull CreateRequest(ctx).RequestCreatable(requestCreatable).Execute()

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
	requestCreatable := *openapiclient.NewRequestCreatable("RequestTypeId_example", "Subject_example") // RequestCreatable | The writable attributes of a request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.CreateRequest(context.Background()).RequestCreatable(requestCreatable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.CreateRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRequest`: RequestFull
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.CreateRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestCreatable** | [**RequestCreatable**](RequestCreatable.md) | The writable attributes of a request | 

### Return type

[**RequestFull**](RequestFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateRequestMessage

> CreateRequestMessage(ctx, requestId).RequestMessageCreatable(requestMessageCreatable).Execute()

Create a Message for a Request



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
	requestId := "requestId_example" // string | The `id` of the request
	requestMessageCreatable := *openapiclient.NewRequestMessageCreatable("Message_example") // RequestMessageCreatable | The writable attributes of a request message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RequestsAPI.CreateRequestMessage(context.Background(), requestId).RequestMessageCreatable(requestMessageCreatable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.CreateRequestMessage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | The &#x60;id&#x60; of the request | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequestMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestMessageCreatable** | [**RequestMessageCreatable**](RequestMessageCreatable.md) | The writable attributes of a request message | 

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


## CreateRequestMessageV2

> CreateRequestMessageV2(ctx, requestId).RequestMessageCreatable(requestMessageCreatable).Execute()

Create a request message



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
	requestId := "requestId_example" // string | The `id` of the request
	requestMessageCreatable := *openapiclient.NewRequestMessageCreatable("Message_example") // RequestMessageCreatable | The writable attributes of a request message

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RequestsAPI.CreateRequestMessageV2(context.Background(), requestId).RequestMessageCreatable(requestMessageCreatable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.CreateRequestMessageV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | The &#x60;id&#x60; of the request | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequestMessageV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestMessageCreatable** | [**RequestMessageCreatable**](RequestMessageCreatable.md) | The writable attributes of a request message | 

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


## CreateRequestV2

> RequestSubmissionFull CreateRequestV2(ctx).RequestCreatable2(requestCreatable2).Execute()

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
	requestCreatable2 := *openapiclient.NewRequestCreatable2(openapiclient.request-resource-creatable{RequestResourceCatalogEntryCreatable: openapiclient.NewRequestResourceCatalogEntryCreatable("Type_example", "cenp2rjyxK1Js2Fc41d5")}, *openapiclient.NewTargetPrincipal("00ub0oNGTSWTBKOLGLNR", openapiclient.principal-type("OKTA_USER"))) // RequestCreatable2 | The writable attributes of a request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.CreateRequestV2(context.Background()).RequestCreatable2(requestCreatable2).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.CreateRequestV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRequestV2`: RequestSubmissionFull
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.CreateRequestV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequestV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestCreatable2** | [**RequestCreatable2**](RequestCreatable2.md) | The writable attributes of a request | 

### Return type

[**RequestSubmissionFull**](RequestSubmissionFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequest

> RequestFull GetRequest(ctx, requestId).Execute()

Retrieve a request



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
	requestId := "requestId_example" // string | The `id` of the request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.GetRequest(context.Background(), requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.GetRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequest`: RequestFull
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.GetRequest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | The &#x60;id&#x60; of the request | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestFull**](RequestFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequestV2

> RequestFull2 GetRequestV2(ctx, requestId).Execute()

Retrieve a request



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
	requestId := "requestId_example" // string | The `id` of the request

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.GetRequestV2(context.Background(), requestId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.GetRequestV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequestV2`: RequestFull2
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.GetRequestV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**requestId** | **string** | The &#x60;id&#x60; of the request | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequestV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RequestFull2**](RequestFull2.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllRequests

> RequestList ListAllRequests(ctx).Filter(filter).After(after).Limit(limit).OrderBy(orderBy).Execute()

List all requests



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
	filter := "requestStatus%20eq%20%22RESOLVED%22" // string | Apply various filters by using supported request filtering properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  (optional)
	after := "after_example" // string | The after cursor provided by a prior request. (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	orderBy := "created%20desc" // string | Apply an ordering of requests by specifying a supported request property name with `%20asc` or `%20desc` suffix.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.ListAllRequests(context.Background()).Filter(filter).After(after).Limit(limit).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.ListAllRequests``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllRequests`: RequestList
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.ListAllRequests`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllRequestsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Apply various filters by using supported request filtering properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  | 
 **after** | **string** | The after cursor provided by a prior request. | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **orderBy** | **string** | Apply an ordering of requests by specifying a supported request property name with &#x60;%20asc&#x60; or &#x60;%20desc&#x60; suffix.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  | 

### Return type

[**RequestList**](RequestList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllRequestsV2

> RequestList2 ListAllRequestsV2(ctx).Filter(filter).After(after).Limit(limit).OrderBy(orderBy).Execute()

List all requests



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
	filter := "status%20eq%20%22APPROVED%22" // string | Apply various filters by using supported request filtering properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  (optional)
	after := "after_example" // string | The after cursor provided by a prior request. (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	orderBy := "created%20desc" // string | Apply an ordering of requests by specifying a supported request property name with `%20asc` or `%20desc` suffix.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RequestsAPI.ListAllRequestsV2(context.Background()).Filter(filter).After(after).Limit(limit).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RequestsAPI.ListAllRequestsV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllRequestsV2`: RequestList2
	fmt.Fprintf(os.Stdout, "Response from `RequestsAPI.ListAllRequestsV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllRequestsV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Apply various filters by using supported request filtering properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  | 
 **after** | **string** | The after cursor provided by a prior request. | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **orderBy** | **string** | Apply an ordering of requests by specifying a supported request property name with &#x60;%20asc&#x60; or &#x60;%20desc&#x60; suffix.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  | 

### Return type

[**RequestList2**](RequestList2.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

