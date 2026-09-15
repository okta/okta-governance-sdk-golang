# \MyTasksAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMyTaskV2**](MyTasksAPI.md#GetMyTaskV2) | **Get** /governance/api/v2/my/tasks/{taskId} | Retrieve my task
[**ListAllMyTasksV2**](MyTasksAPI.md#ListAllMyTasksV2) | **Get** /governance/api/v2/my/tasks | List all my tasks
[**ResolveMyTaskV2**](MyTasksAPI.md#ResolveMyTaskV2) | **Post** /governance/api/v2/my/tasks/{taskId}/resolve | Resolve my task



## GetMyTaskV2

> TaskFull GetMyTaskV2(ctx, taskId).Execute()

Retrieve my task



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
	taskId := "taskId_example" // string | The `id` of the task

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyTasksAPI.GetMyTaskV2(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyTasksAPI.GetMyTaskV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyTaskV2`: TaskFull
	fmt.Fprintf(os.Stdout, "Response from `MyTasksAPI.GetMyTaskV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | The &#x60;id&#x60; of the task | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMyTaskV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TaskFull**](TaskFull.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllMyTasksV2

> TaskList ListAllMyTasksV2(ctx).Filter(filter).Limit(limit).After(after).OrderBy(orderBy).Execute()

List all my tasks



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
	filter := "status eq "OPEN"" // string | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters a collection of tasks in the response. The filter expression supports the following properties and [operators](https://developer.okta.com/docs/api/#operators):  - `status`: The status of the task. Possible values: `OPEN`, `COMPLETED`, `ASSIGNING`, and `LOCKED`. Supports the `eq` operator. - `type`: The type of the task. Possible values: `QUESTION`, `APPROVAL`, and `TODO`. Supports the `eq` operator. - `requestId`: The ID of the request for the task. Supports the `eq` operator. - `createdAt`: The task created time. Supports the `gt`, `ge`, `le`, and `lt` operators. - `updatedAt`: The latest task updated time. Supports the `gt`, `ge`, `le`, and `lt` operators.  > **Note**: Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	after := "after_example" // string | Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the standard link headers. See [pagination](https://developer.okta.com/docs/api/#pagination).  (optional)
	orderBy := "createdAt desc" // string | Specifies a property to sort the results. The following properties are supported: - `createdAt` - `updatedAt`  Append `desc` or `asc` to indicate the sorting direction. > **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyTasksAPI.ListAllMyTasksV2(context.Background()).Filter(filter).Limit(limit).After(after).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyTasksAPI.ListAllMyTasksV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAllMyTasksV2`: TaskList
	fmt.Fprintf(os.Stdout, "Response from `MyTasksAPI.ListAllMyTasksV2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAllMyTasksV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters a collection of tasks in the response. The filter expression supports the following properties and [operators](https://developer.okta.com/docs/api/#operators):  - &#x60;status&#x60;: The status of the task. Possible values: &#x60;OPEN&#x60;, &#x60;COMPLETED&#x60;, &#x60;ASSIGNING&#x60;, and &#x60;LOCKED&#x60;. Supports the &#x60;eq&#x60; operator. - &#x60;type&#x60;: The type of the task. Possible values: &#x60;QUESTION&#x60;, &#x60;APPROVAL&#x60;, and &#x60;TODO&#x60;. Supports the &#x60;eq&#x60; operator. - &#x60;requestId&#x60;: The ID of the request for the task. Supports the &#x60;eq&#x60; operator. - &#x60;createdAt&#x60;: The task created time. Supports the &#x60;gt&#x60;, &#x60;ge&#x60;, &#x60;le&#x60;, and &#x60;lt&#x60; operators. - &#x60;updatedAt&#x60;: The latest task updated time. Supports the &#x60;gt&#x60;, &#x60;ge&#x60;, &#x60;le&#x60;, and &#x60;lt&#x60; operators.  &gt; **Note**: Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **after** | **string** | Specifies the pagination cursor for the next page of results. Treat this as an opaque value obtained through the standard link headers. See [pagination](https://developer.okta.com/docs/api/#pagination).  | 
 **orderBy** | **string** | Specifies a property to sort the results. The following properties are supported: - &#x60;createdAt&#x60; - &#x60;updatedAt&#x60;  Append &#x60;desc&#x60; or &#x60;asc&#x60; to indicate the sorting direction. &gt; **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  | 

### Return type

[**TaskList**](TaskList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResolveMyTaskV2

> TaskFull ResolveMyTaskV2(ctx, taskId).ResolveMyTaskV2Request(resolveMyTaskV2Request).Execute()

Resolve my task



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
	taskId := "taskId_example" // string | The `id` of the task
	resolveMyTaskV2Request := *openapiclient.NewResolveMyTaskV2Request() // ResolveMyTaskV2Request |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MyTasksAPI.ResolveMyTaskV2(context.Background(), taskId).ResolveMyTaskV2Request(resolveMyTaskV2Request).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MyTasksAPI.ResolveMyTaskV2``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ResolveMyTaskV2`: TaskFull
	fmt.Fprintf(os.Stdout, "Response from `MyTasksAPI.ResolveMyTaskV2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | The &#x60;id&#x60; of the task | 

### Other Parameters

Other parameters are passed through a pointer to a apiResolveMyTaskV2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resolveMyTaskV2Request** | [**ResolveMyTaskV2Request**](ResolveMyTaskV2Request.md) |  | 

### Return type

[**TaskFull**](TaskFull.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

