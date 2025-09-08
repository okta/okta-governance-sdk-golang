# \RiskRulesAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRiskRule**](RiskRulesAPI.md#CreateRiskRule) | **Post** /governance/api/v1/risk-rules | Create a risk rule
[**DeleteRiskRule**](RiskRulesAPI.md#DeleteRiskRule) | **Delete** /governance/api/v1/risk-rules/{ruleId} | Delete a risk rule
[**GeneratePotentialRiskAssessments**](RiskRulesAPI.md#GeneratePotentialRiskAssessments) | **Post** /governance/api/v1/risk-rule-assessments | Generate a risk assessment
[**GetRiskRule**](RiskRulesAPI.md#GetRiskRule) | **Get** /governance/api/v1/risk-rules/{ruleId} | Retrieve a risk rule
[**ListRiskRules**](RiskRulesAPI.md#ListRiskRules) | **Get** /governance/api/v1/risk-rules | List all risk rules
[**ReplaceRiskRule**](RiskRulesAPI.md#ReplaceRiskRule) | **Put** /governance/api/v1/risk-rules/{ruleId} | Replace a risk rule



## CreateRiskRule

> RiskRuleResponse CreateRiskRule(ctx).CreateRiskRuleRequest(createRiskRuleRequest).Execute()

Create a risk rule



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
	createRiskRuleRequest := *openapiclient.NewCreateRiskRuleRequest("Name_example", "Type_example", []openapiclient.RuleConflictResource{*openapiclient.NewRuleConflictResource()}, *openapiclient.NewConflictCriteriaCreatable()) // CreateRiskRuleRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskRulesAPI.CreateRiskRule(context.Background()).CreateRiskRuleRequest(createRiskRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskRulesAPI.CreateRiskRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRiskRule`: RiskRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `RiskRulesAPI.CreateRiskRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateRiskRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createRiskRuleRequest** | [**CreateRiskRuleRequest**](CreateRiskRuleRequest.md) |  | 

### Return type

[**RiskRuleResponse**](RiskRuleResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRiskRule

> DeleteRiskRule(ctx, ruleId).Execute()

Delete a risk rule



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
	ruleId := "ruleId_example" // string | The `id` of the risk rule

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RiskRulesAPI.DeleteRiskRule(context.Background(), ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskRulesAPI.DeleteRiskRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | The &#x60;id&#x60; of the risk rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRiskRuleRequest struct via the builder pattern


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


## GeneratePotentialRiskAssessments

> RiskAssessmentResponse GeneratePotentialRiskAssessments(ctx).PotentialRiskAssessmentRequest(potentialRiskAssessmentRequest).Execute()

Generate a risk assessment



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
	potentialRiskAssessmentRequest := *openapiclient.NewPotentialRiskAssessmentRequest("orn:okta:directory:00o8rk36Bp5eZKOrw0g4:users:00u1ktfFMZ5HNoj7k0g4", "ResourceOrn_example") // PotentialRiskAssessmentRequest | Description of the requested access resource by user

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskRulesAPI.GeneratePotentialRiskAssessments(context.Background()).PotentialRiskAssessmentRequest(potentialRiskAssessmentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskRulesAPI.GeneratePotentialRiskAssessments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GeneratePotentialRiskAssessments`: RiskAssessmentResponse
	fmt.Fprintf(os.Stdout, "Response from `RiskRulesAPI.GeneratePotentialRiskAssessments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGeneratePotentialRiskAssessmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **potentialRiskAssessmentRequest** | [**PotentialRiskAssessmentRequest**](PotentialRiskAssessmentRequest.md) | Description of the requested access resource by user | 

### Return type

[**RiskAssessmentResponse**](RiskAssessmentResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRiskRule

> RiskRuleResponse GetRiskRule(ctx, ruleId).Execute()

Retrieve a risk rule



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
	ruleId := "ruleId_example" // string | The `id` of the risk rule

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskRulesAPI.GetRiskRule(context.Background(), ruleId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskRulesAPI.GetRiskRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRiskRule`: RiskRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `RiskRulesAPI.GetRiskRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | The &#x60;id&#x60; of the risk rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRiskRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RiskRuleResponse**](RiskRuleResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListRiskRules

> ListRiskRuleResponse ListRiskRules(ctx).Limit(limit).After(after).Filter(filter).Execute()

List all risk rules



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
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	after := "00u68w6vzKLultXS97g6" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. (optional)
	filter := "name sw "Process"" // string | A filter expression that returns entries based on the `resourceOrn` and `name` properties and supports the following operators: *  `eq` operator for the `resourceOrn` property *  `sw` and `co` operator for the `name` property  > **Note:** Query parameter percent encoding is required. See [Special characters](https://developer.okta.com/docs/api/#special-characters).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskRulesAPI.ListRiskRules(context.Background()).Limit(limit).After(after).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskRulesAPI.ListRiskRules``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListRiskRules`: ListRiskRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `RiskRulesAPI.ListRiskRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListRiskRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. | 
 **filter** | **string** | A filter expression that returns entries based on the &#x60;resourceOrn&#x60; and &#x60;name&#x60; properties and supports the following operators: *  &#x60;eq&#x60; operator for the &#x60;resourceOrn&#x60; property *  &#x60;sw&#x60; and &#x60;co&#x60; operator for the &#x60;name&#x60; property  &gt; **Note:** Query parameter percent encoding is required. See [Special characters](https://developer.okta.com/docs/api/#special-characters).  | 

### Return type

[**ListRiskRuleResponse**](ListRiskRuleResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceRiskRule

> RiskRuleResponse ReplaceRiskRule(ctx, ruleId).UpdateRiskRuleRequest(updateRiskRuleRequest).Execute()

Replace a risk rule



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
	ruleId := "ruleId_example" // string | The `id` of the risk rule
	updateRiskRuleRequest := *openapiclient.NewUpdateRiskRuleRequest("Id_example") // UpdateRiskRuleRequest | The updatable attributes of a risk rule

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RiskRulesAPI.ReplaceRiskRule(context.Background(), ruleId).UpdateRiskRuleRequest(updateRiskRuleRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RiskRulesAPI.ReplaceRiskRule``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceRiskRule`: RiskRuleResponse
	fmt.Fprintf(os.Stdout, "Response from `RiskRulesAPI.ReplaceRiskRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** | The &#x60;id&#x60; of the risk rule | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceRiskRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateRiskRuleRequest** | [**UpdateRiskRuleRequest**](UpdateRiskRuleRequest.md) | The updatable attributes of a risk rule | 

### Return type

[**RiskRuleResponse**](RiskRuleResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

