# \MySettingsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMySettings**](MySettingsAPI.md#GetMySettings) | **Get** /governance/api/v1/my/settings | Retrieve my settings
[**ListMyDelegateUsers**](MySettingsAPI.md#ListMyDelegateUsers) | **Get** /governance/api/v1/my/settings/delegate/users | List my eligible delegates
[**UpdateMySettings**](MySettingsAPI.md#UpdateMySettings) | **Patch** /governance/api/v1/my/settings | Update my settings



## GetMySettings

> MySettingsGet GetMySettings(ctx).Execute()

Retrieve my settings



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
	resp, r, err := apiClient.MySettingsAPI.GetMySettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MySettingsAPI.GetMySettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMySettings`: MySettingsGet
	fmt.Fprintf(os.Stdout, "Response from `MySettingsAPI.GetMySettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMySettingsRequest struct via the builder pattern


### Return type

[**MySettingsGet**](MySettingsGet.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMyDelegateUsers

> DelegateUsersList ListMyDelegateUsers(ctx).Filter(filter).After(after).Limit(limit).Execute()

List my eligible delegates



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
	filter := "firstName sw "John" OR lastName sw "John"" // string | A required filter expression that returns users based on the `firstName` or `lastName` properties. This [filter](https://developer.okta.com/docs/api/#filter) expression supports the `sw` [operator](https://developer.okta.com/docs/api/#operators).  **Note:** Query parameter percent encoding is required. See [Special characters]( https://developer.okta.com/docs/api/#special-characters ). 
	after := "00u68w6vzKLultXS97g6" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. (optional)
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MySettingsAPI.ListMyDelegateUsers(context.Background()).Filter(filter).After(after).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MySettingsAPI.ListMyDelegateUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListMyDelegateUsers`: DelegateUsersList
	fmt.Fprintf(os.Stdout, "Response from `MySettingsAPI.ListMyDelegateUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMyDelegateUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A required filter expression that returns users based on the &#x60;firstName&#x60; or &#x60;lastName&#x60; properties. This [filter](https://developer.okta.com/docs/api/#filter) expression supports the &#x60;sw&#x60; [operator](https://developer.okta.com/docs/api/#operators).  **Note:** Query parameter percent encoding is required. See [Special characters]( https://developer.okta.com/docs/api/#special-characters ).  | 
 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]

### Return type

[**DelegateUsersList**](DelegateUsersList.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMySettings

> MySettings UpdateMySettings(ctx).MySettingsPatchable(mySettingsPatchable).Execute()

Update my settings



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
	mySettingsPatchable := *openapiclient.NewMySettingsPatchable() // MySettingsPatchable | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MySettingsAPI.UpdateMySettings(context.Background()).MySettingsPatchable(mySettingsPatchable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MySettingsAPI.UpdateMySettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMySettings`: MySettings
	fmt.Fprintf(os.Stdout, "Response from `MySettingsAPI.UpdateMySettings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMySettingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **mySettingsPatchable** | [**MySettingsPatchable**](MySettingsPatchable.md) |  | 

### Return type

[**MySettings**](MySettings.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

