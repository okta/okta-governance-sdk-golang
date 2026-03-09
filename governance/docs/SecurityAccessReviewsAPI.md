# \SecurityAccessReviewsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSecurityAccessReviewComment**](SecurityAccessReviewsAPI.md#AddSecurityAccessReviewComment) | **Post** /governance/api/v2/security-access-reviews/{securityAccessReviewId}/comment | Add a comment for a security access review
[**CreateSecurityAccessReview**](SecurityAccessReviewsAPI.md#CreateSecurityAccessReview) | **Post** /governance/api/v2/security-access-reviews | Create a security access review
[**ExecuteSecurityAccessReviewAccessesAction**](SecurityAccessReviewsAPI.md#ExecuteSecurityAccessReviewAccessesAction) | **Post** /governance/api/v2/security-access-reviews/{securityAccessReviewId}/accesses/{securityAccessReviewTargetId}/actions | Execute an action on an access item
[**ExecuteSecurityAccessReviewAction**](SecurityAccessReviewsAPI.md#ExecuteSecurityAccessReviewAction) | **Post** /governance/api/v2/security-access-reviews/{securityAccessReviewId}/actions | Execute an action on a security access review
[**GenerateSecurityAccessReviewAccessesSummary**](SecurityAccessReviewsAPI.md#GenerateSecurityAccessReviewAccessesSummary) | **Post** /governance/api/v2/security-access-reviews/{securityAccessReviewId}/accesses/{securityAccessReviewTargetId}/summary | Generate a summary for an access item
[**GenerateSecurityAccessReviewSummary**](SecurityAccessReviewsAPI.md#GenerateSecurityAccessReviewSummary) | **Post** /governance/api/v2/security-access-reviews/{securityAccessReviewId}/summary | Generate a summary for a security access review
[**GetSecurityAccessReview**](SecurityAccessReviewsAPI.md#GetSecurityAccessReview) | **Get** /governance/api/v2/security-access-reviews/{securityAccessReviewId} | Retrieve a security access review
[**GetSecurityAccessReviewPrincipalDetails**](SecurityAccessReviewsAPI.md#GetSecurityAccessReviewPrincipalDetails) | **Get** /governance/api/v2/security-access-reviews/{securityAccessReviewId}/principal | Retrieve the principal for a security access review
[**GetSecurityAccessReviewsStats**](SecurityAccessReviewsAPI.md#GetSecurityAccessReviewsStats) | **Get** /governance/api/v2/security-access-reviews/stats | Retrieve the statistics for security access reviews
[**ListSecurityAccessReviewAccesses**](SecurityAccessReviewsAPI.md#ListSecurityAccessReviewAccesses) | **Get** /governance/api/v2/security-access-reviews/{securityAccessReviewId}/accesses | List the access items for a security access review
[**ListSecurityAccessReviewAccessesAnomalies**](SecurityAccessReviewsAPI.md#ListSecurityAccessReviewAccessesAnomalies) | **Get** /governance/api/v2/security-access-reviews/{securityAccessReviewId}/accesses/{securityAccessReviewTargetId}/anomalies | List the anomalies for an access item
[**ListSecurityAccessReviewActions**](SecurityAccessReviewsAPI.md#ListSecurityAccessReviewActions) | **Get** /governance/api/v2/security-access-reviews/{securityAccessReviewId}/actions | List the actions for a security access review
[**ListSecurityAccessReviewHistory**](SecurityAccessReviewsAPI.md#ListSecurityAccessReviewHistory) | **Get** /governance/api/v2/security-access-reviews/{securityAccessReviewId}/history | List the history of a security access review
[**ListSecurityAccessReviewSubAccesses**](SecurityAccessReviewsAPI.md#ListSecurityAccessReviewSubAccesses) | **Get** /governance/api/v2/security-access-reviews/{securityAccessReviewId}/accesses/{securityAccessReviewAccessId}/sub-accesses | List the sub-access items for an access item
[**ListSecurityAccessReviews**](SecurityAccessReviewsAPI.md#ListSecurityAccessReviews) | **Get** /governance/api/v2/security-access-reviews | List all security access reviews
[**UpdateSecurityAccessReview**](SecurityAccessReviewsAPI.md#UpdateSecurityAccessReview) | **Patch** /governance/api/v2/security-access-reviews/{securityAccessReviewId} | Update a security access review



## AddSecurityAccessReviewComment

> AddSecurityAccessReviewComment(ctx, securityAccessReviewId).SecurityAccessReviewCommentRequest(securityAccessReviewCommentRequest).Execute()

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
	r, err := apiClient.SecurityAccessReviewsAPI.AddSecurityAccessReviewComment(context.Background(), securityAccessReviewId).SecurityAccessReviewCommentRequest(securityAccessReviewCommentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAccessReviewsAPI.AddSecurityAccessReviewComment``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAddSecurityAccessReviewCommentRequest struct via the builder pattern


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


## CreateSecurityAccessReview

> SecurityAccessReview CreateSecurityAccessReview(ctx).SecurityAccessReviewRequest(securityAccessReviewRequest).Execute()

Create a security access review



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
	securityAccessReviewRequest := *openapiclient.NewSecurityAccessReviewRequest("PrincipalId_example", "Name_example", *openapiclient.NewSecurityAccessReviewReviewerSettings(openapiclient.security-access-review-reviewer-type("USER"))) // SecurityAccessReviewRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAccessReviewsAPI.CreateSecurityAccessReview(context.Background()).SecurityAccessReviewRequest(securityAccessReviewRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAccessReviewsAPI.CreateSecurityAccessReview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSecurityAccessReview`: SecurityAccessReview
	fmt.Fprintf(os.Stdout, "Response from `SecurityAccessReviewsAPI.CreateSecurityAccessReview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSecurityAccessReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **securityAccessReviewRequest** | [**SecurityAccessReviewRequest**](SecurityAccessReviewRequest.md) |  | 

### Return type

[**SecurityAccessReview**](SecurityAccessReview.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExecuteSecurityAccessReviewAccessesAction

> ExecuteSecurityAccessReviewAccessesAction(ctx, securityAccessReviewId, securityAccessReviewTargetId).SecurityAccessReviewAccessesActionRequest(securityAccessReviewAccessesActionRequest).Execute()

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
	r, err := apiClient.SecurityAccessReviewsAPI.ExecuteSecurityAccessReviewAccessesAction(context.Background(), securityAccessReviewId, securityAccessReviewTargetId).SecurityAccessReviewAccessesActionRequest(securityAccessReviewAccessesActionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAccessReviewsAPI.ExecuteSecurityAccessReviewAccessesAction``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExecuteSecurityAccessReviewAccessesActionRequest struct via the builder pattern


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


## ExecuteSecurityAccessReviewAction

> ExecuteSecurityAccessReviewAction(ctx, securityAccessReviewId).SecurityAccessReviewActionRequest(securityAccessReviewActionRequest).Execute()

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
	r, err := apiClient.SecurityAccessReviewsAPI.ExecuteSecurityAccessReviewAction(context.Background(), securityAccessReviewId).SecurityAccessReviewActionRequest(securityAccessReviewActionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAccessReviewsAPI.ExecuteSecurityAccessReviewAction``: %v\n", err)
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

Other parameters are passed through a pointer to a apiExecuteSecurityAccessReviewActionRequest struct via the builder pattern


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


## GenerateSecurityAccessReviewAccessesSummary

> AiMessage GenerateSecurityAccessReviewAccessesSummary(ctx, securityAccessReviewId, securityAccessReviewTargetId).Execute()

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
	resp, r, err := apiClient.SecurityAccessReviewsAPI.GenerateSecurityAccessReviewAccessesSummary(context.Background(), securityAccessReviewId, securityAccessReviewTargetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAccessReviewsAPI.GenerateSecurityAccessReviewAccessesSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateSecurityAccessReviewAccessesSummary`: AiMessage
	fmt.Fprintf(os.Stdout, "Response from `SecurityAccessReviewsAPI.GenerateSecurityAccessReviewAccessesSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityAccessReviewId** | **string** | The ID of the security access review | 
**securityAccessReviewTargetId** | **string** | The ID of the access or sub-access item in a security access review | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateSecurityAccessReviewAccessesSummaryRequest struct via the builder pattern


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


## GenerateSecurityAccessReviewSummary

> AiMessage GenerateSecurityAccessReviewSummary(ctx, securityAccessReviewId).Execute()

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
	resp, r, err := apiClient.SecurityAccessReviewsAPI.GenerateSecurityAccessReviewSummary(context.Background(), securityAccessReviewId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAccessReviewsAPI.GenerateSecurityAccessReviewSummary``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GenerateSecurityAccessReviewSummary`: AiMessage
	fmt.Fprintf(os.Stdout, "Response from `SecurityAccessReviewsAPI.GenerateSecurityAccessReviewSummary`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityAccessReviewId** | **string** | The ID of the security access review | 

### Other Parameters

Other parameters are passed through a pointer to a apiGenerateSecurityAccessReviewSummaryRequest struct via the builder pattern


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


## GetSecurityAccessReview

> SecurityAccessReview GetSecurityAccessReview(ctx, securityAccessReviewId).Execute()

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
	resp, r, err := apiClient.SecurityAccessReviewsAPI.GetSecurityAccessReview(context.Background(), securityAccessReviewId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAccessReviewsAPI.GetSecurityAccessReview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityAccessReview`: SecurityAccessReview
	fmt.Fprintf(os.Stdout, "Response from `SecurityAccessReviewsAPI.GetSecurityAccessReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityAccessReviewId** | **string** | The ID of the security access review | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityAccessReviewRequest struct via the builder pattern


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


## GetSecurityAccessReviewPrincipalDetails

> SecurityAccessReviewPrincipal GetSecurityAccessReviewPrincipalDetails(ctx, securityAccessReviewId).Execute()

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
	resp, r, err := apiClient.SecurityAccessReviewsAPI.GetSecurityAccessReviewPrincipalDetails(context.Background(), securityAccessReviewId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAccessReviewsAPI.GetSecurityAccessReviewPrincipalDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityAccessReviewPrincipalDetails`: SecurityAccessReviewPrincipal
	fmt.Fprintf(os.Stdout, "Response from `SecurityAccessReviewsAPI.GetSecurityAccessReviewPrincipalDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityAccessReviewId** | **string** | The ID of the security access review | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityAccessReviewPrincipalDetailsRequest struct via the builder pattern


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


## GetSecurityAccessReviewsStats

> SecurityAccessReviewsStats GetSecurityAccessReviewsStats(ctx).Execute()

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
	resp, r, err := apiClient.SecurityAccessReviewsAPI.GetSecurityAccessReviewsStats(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAccessReviewsAPI.GetSecurityAccessReviewsStats``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityAccessReviewsStats`: SecurityAccessReviewsStats
	fmt.Fprintf(os.Stdout, "Response from `SecurityAccessReviewsAPI.GetSecurityAccessReviewsStats`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityAccessReviewsStatsRequest struct via the builder pattern


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


## ListSecurityAccessReviewAccesses

> SecurityAccessReviewAccessItems ListSecurityAccessReviewAccesses(ctx, securityAccessReviewId).Filter(filter).OrderBy(orderBy).After(after).Limit(limit).Execute()

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
	resp, r, err := apiClient.SecurityAccessReviewsAPI.ListSecurityAccessReviewAccesses(context.Background(), securityAccessReviewId).Filter(filter).OrderBy(orderBy).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAccessReviewsAPI.ListSecurityAccessReviewAccesses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecurityAccessReviewAccesses`: SecurityAccessReviewAccessItems
	fmt.Fprintf(os.Stdout, "Response from `SecurityAccessReviewsAPI.ListSecurityAccessReviewAccesses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityAccessReviewId** | **string** | The ID of the security access review | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityAccessReviewAccessesRequest struct via the builder pattern


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


## ListSecurityAccessReviewAccessesAnomalies

> SecurityAccessReviewAnomalies ListSecurityAccessReviewAccessesAnomalies(ctx, securityAccessReviewId, securityAccessReviewTargetId).Execute()

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
	resp, r, err := apiClient.SecurityAccessReviewsAPI.ListSecurityAccessReviewAccessesAnomalies(context.Background(), securityAccessReviewId, securityAccessReviewTargetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAccessReviewsAPI.ListSecurityAccessReviewAccessesAnomalies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecurityAccessReviewAccessesAnomalies`: SecurityAccessReviewAnomalies
	fmt.Fprintf(os.Stdout, "Response from `SecurityAccessReviewsAPI.ListSecurityAccessReviewAccessesAnomalies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityAccessReviewId** | **string** | The ID of the security access review | 
**securityAccessReviewTargetId** | **string** | The ID of the access or sub-access item in a security access review | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityAccessReviewAccessesAnomaliesRequest struct via the builder pattern


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


## ListSecurityAccessReviewActions

> SecurityAccessReviewActions ListSecurityAccessReviewActions(ctx, securityAccessReviewId).Execute()

List the actions for a security access review



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
	resp, r, err := apiClient.SecurityAccessReviewsAPI.ListSecurityAccessReviewActions(context.Background(), securityAccessReviewId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAccessReviewsAPI.ListSecurityAccessReviewActions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecurityAccessReviewActions`: SecurityAccessReviewActions
	fmt.Fprintf(os.Stdout, "Response from `SecurityAccessReviewsAPI.ListSecurityAccessReviewActions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityAccessReviewId** | **string** | The ID of the security access review | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityAccessReviewActionsRequest struct via the builder pattern


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


## ListSecurityAccessReviewHistory

> SecurityAccessReviewHistoryItems ListSecurityAccessReviewHistory(ctx, securityAccessReviewId).After(after).Limit(limit).Execute()

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
	resp, r, err := apiClient.SecurityAccessReviewsAPI.ListSecurityAccessReviewHistory(context.Background(), securityAccessReviewId).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAccessReviewsAPI.ListSecurityAccessReviewHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecurityAccessReviewHistory`: SecurityAccessReviewHistoryItems
	fmt.Fprintf(os.Stdout, "Response from `SecurityAccessReviewsAPI.ListSecurityAccessReviewHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityAccessReviewId** | **string** | The ID of the security access review | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityAccessReviewHistoryRequest struct via the builder pattern


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


## ListSecurityAccessReviewSubAccesses

> SecurityAccessReviewSubAccessItems ListSecurityAccessReviewSubAccesses(ctx, securityAccessReviewId, securityAccessReviewAccessId).Filter(filter).OrderBy(orderBy).After(after).Limit(limit).Execute()

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
	resp, r, err := apiClient.SecurityAccessReviewsAPI.ListSecurityAccessReviewSubAccesses(context.Background(), securityAccessReviewId, securityAccessReviewAccessId).Filter(filter).OrderBy(orderBy).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAccessReviewsAPI.ListSecurityAccessReviewSubAccesses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecurityAccessReviewSubAccesses`: SecurityAccessReviewSubAccessItems
	fmt.Fprintf(os.Stdout, "Response from `SecurityAccessReviewsAPI.ListSecurityAccessReviewSubAccesses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityAccessReviewId** | **string** | The ID of the security access review | 
**securityAccessReviewAccessId** | **string** | The ID of the access item in a security access review | 

### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityAccessReviewSubAccessesRequest struct via the builder pattern


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


## ListSecurityAccessReviews

> SecurityAccessReviews ListSecurityAccessReviews(ctx).Filter(filter).OrderBy(orderBy).After(after).Limit(limit).Execute()

List all security access reviews



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
	resp, r, err := apiClient.SecurityAccessReviewsAPI.ListSecurityAccessReviews(context.Background()).Filter(filter).OrderBy(orderBy).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAccessReviewsAPI.ListSecurityAccessReviews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListSecurityAccessReviews`: SecurityAccessReviews
	fmt.Fprintf(os.Stdout, "Response from `SecurityAccessReviewsAPI.ListSecurityAccessReviews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSecurityAccessReviewsRequest struct via the builder pattern


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


## UpdateSecurityAccessReview

> SecurityAccessReview UpdateSecurityAccessReview(ctx, securityAccessReviewId).SecurityAccessReviewUpdateRequest(securityAccessReviewUpdateRequest).Execute()

Update a security access review



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
	securityAccessReviewUpdateRequest := *openapiclient.NewSecurityAccessReviewUpdateRequest() // SecurityAccessReviewUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityAccessReviewsAPI.UpdateSecurityAccessReview(context.Background(), securityAccessReviewId).SecurityAccessReviewUpdateRequest(securityAccessReviewUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityAccessReviewsAPI.UpdateSecurityAccessReview``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSecurityAccessReview`: SecurityAccessReview
	fmt.Fprintf(os.Stdout, "Response from `SecurityAccessReviewsAPI.UpdateSecurityAccessReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**securityAccessReviewId** | **string** | The ID of the security access review | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityAccessReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **securityAccessReviewUpdateRequest** | [**SecurityAccessReviewUpdateRequest**](SecurityAccessReviewUpdateRequest.md) |  | 

### Return type

[**SecurityAccessReview**](SecurityAccessReview.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

