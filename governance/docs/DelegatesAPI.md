# \DelegatesAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListDelegateAppointments**](DelegatesAPI.md#ListDelegateAppointments) | **Get** /governance/api/v1/delegates | List all delegate appointments



## ListDelegateAppointments

> DelegateAppointmentList ListDelegateAppointments(ctx).Filter(filter).Limit(limit).After(after).Execute()

List all delegate appointments



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
	filter := "delegatorId eq "00ub0oNGTSWTBKOLGLNR" OR delegatorId eq "00ub0oNGTSWTBKOLGLNS"" // string | Filter expression for retrieving delegates based on specific criteria. Only the `delegatorId` property is supported with the `eq` operator.  (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	after := "00u68w6vzKLultXS97g6" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DelegatesAPI.ListDelegateAppointments(context.Background()).Filter(filter).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DelegatesAPI.ListDelegateAppointments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDelegateAppointments`: DelegateAppointmentList
	fmt.Fprintf(os.Stdout, "Response from `DelegatesAPI.ListDelegateAppointments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDelegateAppointmentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Filter expression for retrieving delegates based on specific criteria. Only the &#x60;delegatorId&#x60; property is supported with the &#x60;eq&#x60; operator.  | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. | 

### Return type

[**DelegateAppointmentList**](DelegateAppointmentList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

