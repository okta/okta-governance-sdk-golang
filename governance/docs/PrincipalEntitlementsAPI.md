# \PrincipalEntitlementsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPrincipalEntitlements**](PrincipalEntitlementsAPI.md#GetPrincipalEntitlements) | **Get** /governance/api/v1/principal-entitlements | Retrieve the principal&#39;s effective entitlements for a resource



## GetPrincipalEntitlements

> PrincipalEntitlementsList GetPrincipalEntitlements(ctx).Filter(filter).Execute()

Retrieve the principal's effective entitlements for a resource



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
	filter := "parent.externalId%20eq%20%220oafxqCAJWWGELFTYASJ%22%20AND%20parent.type%20eq%20%22APPLICATION%22%20AND%20targetPrincipal.externalId%20eq%20%2200ub0oNGTSWTBKOLGLNR%22%20AND%20targetPrincipal.type%20eq%20%22OKTA_USER%22" // string | Apply various filters by using supported principal entitlements filtering properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding) 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrincipalEntitlementsAPI.GetPrincipalEntitlements(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrincipalEntitlementsAPI.GetPrincipalEntitlements``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrincipalEntitlements`: PrincipalEntitlementsList
	fmt.Fprintf(os.Stdout, "Response from `PrincipalEntitlementsAPI.GetPrincipalEntitlements`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrincipalEntitlementsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | Apply various filters by using supported principal entitlements filtering properties.  **Note:** Query parameter percent encoding is required. See [Percent-encoding](https://developer.mozilla.org/en-US/docs/Glossary/Percent-encoding)  | 

### Return type

[**PrincipalEntitlementsList**](PrincipalEntitlementsList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

