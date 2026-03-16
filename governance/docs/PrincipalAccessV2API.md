# \PrincipalAccessV2API

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RevokePrincipalAccess**](PrincipalAccessV2API.md#RevokePrincipalAccess) | **Post** /governance/api/v2/revoke-principal-access | Revoke a principal&#39;s access



## RevokePrincipalAccess

> RevokePrincipalAccessResponse RevokePrincipalAccess(ctx).RevokePrincipalAccessCreatable(revokePrincipalAccessCreatable).Execute()

Revoke a principal's access



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
	revokePrincipalAccessCreatable := *openapiclient.NewRevokePrincipalAccessCreatable("orn:okta:directory:00o8rk36Bp5eZKOrw0g4:users:00u1ktfFMZ5HNoj7k0g4", []string{"orn:okta:governance:00o11rndFqmZ5rNfs0g4:entitlement-values:ent65J1scEgrRZbjT0g2"}) // RevokePrincipalAccessCreatable | The revocation request parameters

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrincipalAccessV2API.RevokePrincipalAccess(context.Background()).RevokePrincipalAccessCreatable(revokePrincipalAccessCreatable).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrincipalAccessV2API.RevokePrincipalAccess``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevokePrincipalAccess`: RevokePrincipalAccessResponse
	fmt.Fprintf(os.Stdout, "Response from `PrincipalAccessV2API.RevokePrincipalAccess`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRevokePrincipalAccessRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **revokePrincipalAccessCreatable** | [**RevokePrincipalAccessCreatable**](RevokePrincipalAccessCreatable.md) | The revocation request parameters | 

### Return type

[**RevokePrincipalAccessResponse**](RevokePrincipalAccessResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

