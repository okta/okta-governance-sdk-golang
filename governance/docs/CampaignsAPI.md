# \CampaignsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCampaign**](CampaignsAPI.md#CreateCampaign) | **Post** /governance/api/v1/campaigns | Create a campaign
[**DeleteCampaign**](CampaignsAPI.md#DeleteCampaign) | **Delete** /governance/api/v1/campaigns/{campaignId} | Delete a campaign
[**EndCampaign**](CampaignsAPI.md#EndCampaign) | **Post** /governance/api/v1/campaigns/{campaignId}/end | End a campaign
[**GetCampaign**](CampaignsAPI.md#GetCampaign) | **Get** /governance/api/v1/campaigns/{campaignId} | Retrieve a campaign
[**LaunchCampaign**](CampaignsAPI.md#LaunchCampaign) | **Post** /governance/api/v1/campaigns/{campaignId}/launch | Launch a campaign
[**ListCampaigns**](CampaignsAPI.md#ListCampaigns) | **Get** /governance/api/v1/campaigns | List all campaigns



## CreateCampaign

> CampaignFull CreateCampaign(ctx).CampaignMutable(campaignMutable).Execute()

Create a campaign



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/okta/okta-governance-sdk-golang"
)

func main() {
	campaignMutable := *openapiclient.NewCampaignMutable("Name_example", *openapiclient.NewScheduleSettingsMutable(openapiclient.schedule-type("ONE_OFF"), time.Now(), float32(123), "TimeZone_example"), *openapiclient.NewResourceSettingsMutable(openapiclient.campaign-resource-type("GROUP")), *openapiclient.NewReviewerSettingsMutable(openapiclient.campaign-reviewer-type("USER")), *openapiclient.NewRemediationSettings(openapiclient.approved-remediation-action("NO_ACTION"), openapiclient.revoked-remediation-action("NO_ACTION"), openapiclient.no-response-remediation-action("NO_ACTION"))) // CampaignMutable | Specifies the characteristics of a single campaign

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.CreateCampaign(context.Background()).CampaignMutable(campaignMutable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.CreateCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCampaign`: CampaignFull
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.CreateCampaign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignMutable** | [**CampaignMutable**](CampaignMutable.md) | Specifies the characteristics of a single campaign | 

### Return type

[**CampaignFull**](CampaignFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCampaign

> DeleteCampaign(ctx, campaignId).Execute()

Delete a campaign



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
	campaignId := "campaignId_example" // string | The `id` of the campaign

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CampaignsAPI.DeleteCampaign(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.DeleteCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The &#x60;id&#x60; of the campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignRequest struct via the builder pattern


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


## EndCampaign

> EndCampaign(ctx, campaignId).CampaignEndSkipRemediation(campaignEndSkipRemediation).Execute()

End a campaign



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
	campaignId := "campaignId_example" // string | The `id` of the campaign
	campaignEndSkipRemediation := *openapiclient.NewCampaignEndSkipRemediation() // CampaignEndSkipRemediation | Ends a campaign with the option to skip remediation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CampaignsAPI.EndCampaign(context.Background(), campaignId).CampaignEndSkipRemediation(campaignEndSkipRemediation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.EndCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The &#x60;id&#x60; of the campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiEndCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **campaignEndSkipRemediation** | [**CampaignEndSkipRemediation**](CampaignEndSkipRemediation.md) | Ends a campaign with the option to skip remediation. | 

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


## GetCampaign

> CampaignFull GetCampaign(ctx, campaignId).Execute()

Retrieve a campaign



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
	campaignId := "campaignId_example" // string | The `id` of the campaign

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.GetCampaign(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.GetCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaign`: CampaignFull
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.GetCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The &#x60;id&#x60; of the campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CampaignFull**](CampaignFull.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LaunchCampaign

> LaunchCampaign(ctx, campaignId).Execute()

Launch a campaign



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
	campaignId := "campaignId_example" // string | The `id` of the campaign

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CampaignsAPI.LaunchCampaign(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.LaunchCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | The &#x60;id&#x60; of the campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiLaunchCampaignRequest struct via the builder pattern


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


## ListCampaigns

> CampaignsList ListCampaigns(ctx).Filter(filter).After(after).Limit(limit).OrderBy(orderBy).Execute()

List all campaigns



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
	filter := "name%20eq%20%22Sales%20Review%22" // string | Apply various filters by using supported campaign filtering properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  (optional)
	after := "00u68w6vzKLultXS97g6" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	orderBy := []string{"Inner_example"} // []string | Apply an ordering of campaigns by specifying a supported campaign property name with `%20asc` or `%20desc` suffix.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  (optional) (default to ["created asc"])

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignsAPI.ListCampaigns(context.Background()).Filter(filter).After(after).Limit(limit).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignsAPI.ListCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListCampaigns`: CampaignsList
	fmt.Fprintf(os.Stdout, "Response from `CampaignsAPI.ListCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Apply various filters by using supported campaign filtering properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  | 
 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **orderBy** | **[]string** | Apply an ordering of campaigns by specifying a supported campaign property name with &#x60;%20asc&#x60; or &#x60;%20desc&#x60; suffix.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  | [default to [&quot;created asc&quot;]]

### Return type

[**CampaignsList**](CampaignsList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

