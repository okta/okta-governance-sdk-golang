# \MySecurityAccessReviewsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMySecurityAccessReviewComment**](MySecurityAccessReviewsAPI.md#AddMySecurityAccessReviewComment) | **Post** /governance/api/v2/my/security-access-reviews/{securityAccessReviewId}/comment | Add a comment for a security access review
[**ExecuteMySecurityAccessReviewAccessesAction**](MySecurityAccessReviewsAPI.md#ExecuteMySecurityAccessReviewAccessesAction) | **Post** /governance/api/v2/my/security-access-reviews/{securityAccessReviewId}/accesses/{securityAccessReviewTargetId}/actions | Execute an action on an access item
[**ExecuteMySecurityAccessReviewAction**](MySecurityAccessReviewsAPI.md#ExecuteMySecurityAccessReviewAction) | **Post** /governance/api/v2/my/security-access-reviews/{securityAccessReviewId}/actions | Execute an action on a security access review
[**GenerateMySecurityAccessReviewAccessesSummary**](MySecurityAccessReviewsAPI.md#GenerateMySecurityAccessReviewAccessesSummary) | **Post** /governance/api/v2/my/security-access-reviews/{securityAccessReviewId}/accesses/{securityAccessReviewTargetId}/summary | Generate a summary for an access item
[**GenerateMySecurityAccessReviewSummary**](MySecurityAccessReviewsAPI.md#GenerateMySecurityAccessReviewSummary) | **Post** /governance/api/v2/my/security-access-reviews/{securityAccessReviewId}/summary | Generate a summary for a security access review
[**GetMySecurityAccessReview**](MySecurityAccessReviewsAPI.md#GetMySecurityAccessReview) | **Get** /governance/api/v2/my/security-access-reviews/{securityAccessReviewId} | Retrieve a security access review
[**GetMySecurityAccessReviewPrincipalDetails**](MySecurityAccessReviewsAPI.md#GetMySecurityAccessReviewPrincipalDetails) | **Get** /governance/api/v2/my/security-access-reviews/{securityAccessReviewId}/principal | Retrieve the principal for a security access review
[**GetMySecurityAccessReviewsStats**](MySecurityAccessReviewsAPI.md#GetMySecurityAccessReviewsStats) | **Get** /governance/api/v2/my/security-access-reviews/stats | Retrieve the statistics for security access reviews
[**ListMySecurityAccessReviewAccesses**](MySecurityAccessReviewsAPI.md#ListMySecurityAccessReviewAccesses) | **Get** /governance/api/v2/my/security-access-reviews/{securityAccessReviewId}/accesses | List the access items for a security access review
[**ListMySecurityAccessReviewAccessesAnomalies**](MySecurityAccessReviewsAPI.md#ListMySecurityAccessReviewAccessesAnomalies) | **Get** /governance/api/v2/my/security-access-reviews/{securityAccessReviewId}/accesses/{securityAccessReviewTargetId}/anomalies | List the anomalies for an access item
[**ListMySecurityAccessReviewActions**](MySecurityAccessReviewsAPI.md#ListMySecurityAccessReviewActions) | **Get** /governance/api/v2/my/security-access-reviews/{securityAccessReviewId}/actions | List all actions for a security access review
[**ListMySecurityAccessReviewHistory**](MySecurityAccessReviewsAPI.md#ListMySecurityAccessReviewHistory) | **Get** /governance/api/v2/my/security-access-reviews/{securityAccessReviewId}/history | List the history of a security access review
[**ListMySecurityAccessReviewSubAccesses**](MySecurityAccessReviewsAPI.md#ListMySecurityAccessReviewSubAccesses) | **Get** /governance/api/v2/my/security-access-reviews/{securityAccessReviewId}/accesses/{securityAccessReviewAccessId}/sub-accesses | List the sub-access items for an access item
[**ListMySecurityAccessReviews**](MySecurityAccessReviewsAPI.md#ListMySecurityAccessReviews) | **Get** /governance/api/v2/my/security-access-reviews | List the security access reviews



## AddMySecurityAccessReviewComment

> AddMySecurityAccessReviewComment(ctx, securityAccessReviewId).SecurityAccessReviewCommentRequest(securityAccessReviewCommentRequest).Execute()

Add a comment for a security access review



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
	securityAccessReviewId := "securityAccessReviewId_example" // string | The ID of the security access review
	securityAccessReviewCommentRequest := *openapiclient.NewSecurityAccessReviewCommentRequest("Comment_example") // SecurityAccessReviewCommentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MySecurityAccessReviewsAPI.AddMySecurityAccessReviewComment(context.Background(), securityAccessReviewId).SecurityAccessReviewCommentRequest(securityAccessReviewCommentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MySecurityAccessReviewsAPI.AddMySecurityAccessReviewComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityAccessReviewId** | **string** | The ID of the security access review | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMySecurityAccessReviewCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **securityAccessReviewCommentRequest** | [**SecurityAccessReviewCommentRequest**](SecurityAccessReviewCommentRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteMySecurityAccessReviewAccessesAction

> ExecuteMySecurityAccessReviewAccessesAction(ctx, securityAccessReviewId, securityAccessReviewTargetId).SecurityAccessReviewAccessesActionRequest(securityAccessReviewAccessesActionRequest).Execute()

Execute an action on an access item



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
	securityAccessReviewId := "securityAccessReviewId_example" // string | The ID of the security access review
	securityAccessReviewTargetId := "securityAccessReviewTargetId_example" // string | The ID of the access or sub-access item in a security access review
	securityAccessReviewAccessesActionRequest := *openapiclient.NewSecurityAccessReviewAccessesActionRequest(openapiclient.security-access-review-access-item-supported-action("REVOKE_ACCESS")) // SecurityAccessReviewAccessesActionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MySecurityAccessReviewsAPI.ExecuteMySecurityAccessReviewAccessesAction(context.Background(), securityAccessReviewId, securityAccessReviewTargetId).SecurityAccessReviewAccessesActionRequest(securityAccessReviewAccessesActionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MySecurityAccessReviewsAPI.ExecuteMySecurityAccessReviewAccessesAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityAccessReviewId** | **string** | The ID of the security access review | 
**securityAccessReviewTargetId** | **string** | The ID of the access or sub-access item in a security access review | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteMySecurityAccessReviewAccessesActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **securityAccessReviewAccessesActionRequest** | [**SecurityAccessReviewAccessesActionRequest**](SecurityAccessReviewAccessesActionRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteMySecurityAccessReviewAction

> ExecuteMySecurityAccessReviewAction(ctx, securityAccessReviewId).SecurityAccessReviewActionRequest(securityAccessReviewActionRequest).Execute()

Execute an action on a security access review



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
	securityAccessReviewId := "securityAccessReviewId_example" // string | The ID of the security access review
	securityAccessReviewActionRequest := *openapiclient.NewSecurityAccessReviewActionRequest(openapiclient.security-access-review-action-type("CLOSE_REVIEW")) // SecurityAccessReviewActionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MySecurityAccessReviewsAPI.ExecuteMySecurityAccessReviewAction(context.Background(), securityAccessReviewId).SecurityAccessReviewActionRequest(securityAccessReviewActionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MySecurityAccessReviewsAPI.ExecuteMySecurityAccessReviewAction``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityAccessReviewId** | **string** | The ID of the security access review | 

### Other Parameters

Other parameters are passed through a pointer to a apiExecuteMySecurityAccessReviewActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **securityAccessReviewActionRequest** | [**SecurityAccessReviewActionRequest**](SecurityAccessReviewActionRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateMySecurityAccessReviewAccessesSummary

> AiMessage GenerateMySecurityAccessReviewAccessesSummary(ctx, securityAccessReviewId, securityAccessReviewTargetId).Execute()

Generate a summary for an access item



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
	securityAccessReviewId := "securityAccessReviewId_example" // string | The ID of the security access review
	securityAccessReviewTargetId := "securityAccessReviewTargetId_example" // string | The ID of the access or sub-access item in a security access review

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MySecurityAccessReviewsAPI.GenerateMySecurityAccessReviewAccessesSummary(context.Background(), securityAccessReviewId, securityAccessReviewTargetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MySecurityAccessReviewsAPI.GenerateMySecurityAccessReviewAccessesSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateMySecurityAccessReviewAccessesSummary`: AiMessage
	fmt.Fprintf(os.Stdout, "Response from `MySecurityAccessReviewsAPI.GenerateMySecurityAccessReviewAccessesSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityAccessReviewId** | **string** | The ID of the security access review | 
**securityAccessReviewTargetId** | **string** | The ID of the access or sub-access item in a security access review | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateMySecurityAccessReviewAccessesSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AiMessage**](AiMessage.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GenerateMySecurityAccessReviewSummary

> AiMessage GenerateMySecurityAccessReviewSummary(ctx, securityAccessReviewId).Execute()

Generate a summary for a security access review



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
	securityAccessReviewId := "securityAccessReviewId_example" // string | The ID of the security access review

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MySecurityAccessReviewsAPI.GenerateMySecurityAccessReviewSummary(context.Background(), securityAccessReviewId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MySecurityAccessReviewsAPI.GenerateMySecurityAccessReviewSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateMySecurityAccessReviewSummary`: AiMessage
	fmt.Fprintf(os.Stdout, "Response from `MySecurityAccessReviewsAPI.GenerateMySecurityAccessReviewSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityAccessReviewId** | **string** | The ID of the security access review | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateMySecurityAccessReviewSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AiMessage**](AiMessage.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMySecurityAccessReview

> SecurityAccessReview GetMySecurityAccessReview(ctx, securityAccessReviewId).Execute()

Retrieve a security access review



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
	securityAccessReviewId := "securityAccessReviewId_example" // string | The ID of the security access review

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MySecurityAccessReviewsAPI.GetMySecurityAccessReview(context.Background(), securityAccessReviewId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MySecurityAccessReviewsAPI.GetMySecurityAccessReview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMySecurityAccessReview`: SecurityAccessReview
	fmt.Fprintf(os.Stdout, "Response from `MySecurityAccessReviewsAPI.GetMySecurityAccessReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityAccessReviewId** | **string** | The ID of the security access review | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMySecurityAccessReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecurityAccessReview**](SecurityAccessReview.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMySecurityAccessReviewPrincipalDetails

> SecurityAccessReviewPrincipal GetMySecurityAccessReviewPrincipalDetails(ctx, securityAccessReviewId).Execute()

Retrieve the principal for a security access review



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
	securityAccessReviewId := "securityAccessReviewId_example" // string | The ID of the security access review

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MySecurityAccessReviewsAPI.GetMySecurityAccessReviewPrincipalDetails(context.Background(), securityAccessReviewId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MySecurityAccessReviewsAPI.GetMySecurityAccessReviewPrincipalDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMySecurityAccessReviewPrincipalDetails`: SecurityAccessReviewPrincipal
	fmt.Fprintf(os.Stdout, "Response from `MySecurityAccessReviewsAPI.GetMySecurityAccessReviewPrincipalDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityAccessReviewId** | **string** | The ID of the security access review | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMySecurityAccessReviewPrincipalDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecurityAccessReviewPrincipal**](SecurityAccessReviewPrincipal.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMySecurityAccessReviewsStats

> SecurityAccessReviewsStats GetMySecurityAccessReviewsStats(ctx).Execute()

Retrieve the statistics for security access reviews



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MySecurityAccessReviewsAPI.GetMySecurityAccessReviewsStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MySecurityAccessReviewsAPI.GetMySecurityAccessReviewsStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMySecurityAccessReviewsStats`: SecurityAccessReviewsStats
	fmt.Fprintf(os.Stdout, "Response from `MySecurityAccessReviewsAPI.GetMySecurityAccessReviewsStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMySecurityAccessReviewsStatsRequest struct via the builder pattern


### Return type

[**SecurityAccessReviewsStats**](SecurityAccessReviewsStats.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMySecurityAccessReviewAccesses

> SecurityAccessReviewAccessItems ListMySecurityAccessReviewAccesses(ctx, securityAccessReviewId).Filter(filter).OrderBy(orderBy).After(after).Limit(limit).Execute()

List the access items for a security access review



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
	securityAccessReviewId := "securityAccessReviewId_example" // string | The ID of the security access review
	filter := "name%20co%20%22Git%22" // string | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters access items.  > **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding). (optional)
	orderBy := []string{"Inner_example"} // []string | The field to sort the results, in ascending (asc) or descending (desc) order. Sorting is applied to only one field.  > **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding) (optional) (default to ["priority desc"])
	after := "00u68w6vzKLultXS97g6" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MySecurityAccessReviewsAPI.ListMySecurityAccessReviewAccesses(context.Background(), securityAccessReviewId).Filter(filter).OrderBy(orderBy).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MySecurityAccessReviewsAPI.ListMySecurityAccessReviewAccesses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMySecurityAccessReviewAccesses`: SecurityAccessReviewAccessItems
	fmt.Fprintf(os.Stdout, "Response from `MySecurityAccessReviewsAPI.ListMySecurityAccessReviewAccesses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityAccessReviewId** | **string** | The ID of the security access review | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMySecurityAccessReviewAccessesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | **string** | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters access items.  &gt; **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding). | 
 **orderBy** | **[]string** | The field to sort the results, in ascending (asc) or descending (desc) order. Sorting is applied to only one field.  &gt; **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding) | [default to [&quot;priority desc&quot;]]
 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]

### Return type

[**SecurityAccessReviewAccessItems**](SecurityAccessReviewAccessItems.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMySecurityAccessReviewAccessesAnomalies

> SecurityAccessReviewAnomalies ListMySecurityAccessReviewAccessesAnomalies(ctx, securityAccessReviewId, securityAccessReviewTargetId).Execute()

List the anomalies for an access item



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
	securityAccessReviewId := "securityAccessReviewId_example" // string | The ID of the security access review
	securityAccessReviewTargetId := "securityAccessReviewTargetId_example" // string | The ID of the access or sub-access item in a security access review

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MySecurityAccessReviewsAPI.ListMySecurityAccessReviewAccessesAnomalies(context.Background(), securityAccessReviewId, securityAccessReviewTargetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MySecurityAccessReviewsAPI.ListMySecurityAccessReviewAccessesAnomalies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMySecurityAccessReviewAccessesAnomalies`: SecurityAccessReviewAnomalies
	fmt.Fprintf(os.Stdout, "Response from `MySecurityAccessReviewsAPI.ListMySecurityAccessReviewAccessesAnomalies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityAccessReviewId** | **string** | The ID of the security access review | 
**securityAccessReviewTargetId** | **string** | The ID of the access or sub-access item in a security access review | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMySecurityAccessReviewAccessesAnomaliesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SecurityAccessReviewAnomalies**](SecurityAccessReviewAnomalies.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMySecurityAccessReviewActions

> SecurityAccessReviewActions ListMySecurityAccessReviewActions(ctx, securityAccessReviewId).Execute()

List all actions for a security access review



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
	securityAccessReviewId := "securityAccessReviewId_example" // string | The ID of the security access review

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MySecurityAccessReviewsAPI.ListMySecurityAccessReviewActions(context.Background(), securityAccessReviewId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MySecurityAccessReviewsAPI.ListMySecurityAccessReviewActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMySecurityAccessReviewActions`: SecurityAccessReviewActions
	fmt.Fprintf(os.Stdout, "Response from `MySecurityAccessReviewsAPI.ListMySecurityAccessReviewActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityAccessReviewId** | **string** | The ID of the security access review | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMySecurityAccessReviewActionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecurityAccessReviewActions**](SecurityAccessReviewActions.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMySecurityAccessReviewHistory

> SecurityAccessReviewHistoryItems ListMySecurityAccessReviewHistory(ctx, securityAccessReviewId).After(after).Limit(limit).Execute()

List the history of a security access review



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
	securityAccessReviewId := "securityAccessReviewId_example" // string | The ID of the security access review
	after := "00u68w6vzKLultXS97g6" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MySecurityAccessReviewsAPI.ListMySecurityAccessReviewHistory(context.Background(), securityAccessReviewId).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MySecurityAccessReviewsAPI.ListMySecurityAccessReviewHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMySecurityAccessReviewHistory`: SecurityAccessReviewHistoryItems
	fmt.Fprintf(os.Stdout, "Response from `MySecurityAccessReviewsAPI.ListMySecurityAccessReviewHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityAccessReviewId** | **string** | The ID of the security access review | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMySecurityAccessReviewHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]

### Return type

[**SecurityAccessReviewHistoryItems**](SecurityAccessReviewHistoryItems.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMySecurityAccessReviewSubAccesses

> SecurityAccessReviewSubAccessItems ListMySecurityAccessReviewSubAccesses(ctx, securityAccessReviewId, securityAccessReviewAccessId).Filter(filter).OrderBy(orderBy).After(after).Limit(limit).Execute()

List the sub-access items for an access item



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
	securityAccessReviewId := "securityAccessReviewId_example" // string | The ID of the security access review
	securityAccessReviewAccessId := "securityAccessReviewAccessId_example" // string | The ID of the access item in a security access review
	filter := "name%20co%20%22Git%22" // string | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters sub-access items.  > **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  (optional)
	orderBy := []string{"Inner_example"} // []string | A field by which results can be sorted. For now, sorting by a single field is supported.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  (optional) (default to ["priority desc"])
	after := "00u68w6vzKLultXS97g6" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MySecurityAccessReviewsAPI.ListMySecurityAccessReviewSubAccesses(context.Background(), securityAccessReviewId, securityAccessReviewAccessId).Filter(filter).OrderBy(orderBy).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MySecurityAccessReviewsAPI.ListMySecurityAccessReviewSubAccesses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMySecurityAccessReviewSubAccesses`: SecurityAccessReviewSubAccessItems
	fmt.Fprintf(os.Stdout, "Response from `MySecurityAccessReviewsAPI.ListMySecurityAccessReviewSubAccesses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityAccessReviewId** | **string** | The ID of the security access review | 
**securityAccessReviewAccessId** | **string** | The ID of the access item in a security access review | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMySecurityAccessReviewSubAccessesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters sub-access items.  &gt; **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding).  | 
 **orderBy** | **[]string** | A field by which results can be sorted. For now, sorting by a single field is supported.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  | [default to [&quot;priority desc&quot;]]
 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]

### Return type

[**SecurityAccessReviewSubAccessItems**](SecurityAccessReviewSubAccessItems.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMySecurityAccessReviews

> SecurityAccessReviews ListMySecurityAccessReviews(ctx).Filter(filter).OrderBy(orderBy).After(after).Limit(limit).Execute()

List the security access reviews



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
	filter := "name%20co%20%22Git%22" // string | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters security access reviews. The `eq` and `co` [operators](https://developer.okta.com/docs/api/#operators) are supported for string properties. The `gt` and `lt` operators are supported for date properties.  > **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding). (optional)
	orderBy := []string{"Inner_example"} // []string | The field to sort the results, in ascending (`asc`) or descending (`desc`) order. Sorting is applied to only one field.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  (optional) (default to ["created asc"])
	after := "00u68w6vzKLultXS97g6" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MySecurityAccessReviewsAPI.ListMySecurityAccessReviews(context.Background()).Filter(filter).OrderBy(orderBy).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MySecurityAccessReviewsAPI.ListMySecurityAccessReviews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMySecurityAccessReviews`: SecurityAccessReviews
	fmt.Fprintf(os.Stdout, "Response from `MySecurityAccessReviewsAPI.ListMySecurityAccessReviews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMySecurityAccessReviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A [filter](https://developer.okta.com/docs/api/#filter) expression that filters security access reviews. The &#x60;eq&#x60; and &#x60;co&#x60; [operators](https://developer.okta.com/docs/api/#operators) are supported for string properties. The &#x60;gt&#x60; and &#x60;lt&#x60; operators are supported for date properties.  &gt; **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding). | 
 **orderBy** | **[]string** | The field to sort the results, in ascending (&#x60;asc&#x60;) or descending (&#x60;desc&#x60;) order. Sorting is applied to only one field.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  | [default to [&quot;created asc&quot;]]
 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]

### Return type

[**SecurityAccessReviews**](SecurityAccessReviews.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

