# \LabelsAPI

All URIs are relative to *https://subdomain.okta.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignResourceLabels**](LabelsAPI.md#AssignResourceLabels) | **Post** /governance/api/v1/resource-labels/assign | Assign the labels to resources
[**CreateLabel**](LabelsAPI.md#CreateLabel) | **Post** /governance/api/v1/labels | Create a label
[**DeleteLabel**](LabelsAPI.md#DeleteLabel) | **Delete** /governance/api/v1/labels/{labelId} | Delete a label
[**GetLabel**](LabelsAPI.md#GetLabel) | **Get** /governance/api/v1/labels/{labelId} | Retrieve a label
[**ListLabelResources**](LabelsAPI.md#ListLabelResources) | **Get** /governance/api/v1/resource-labels | List all labeled resources
[**ListLabels**](LabelsAPI.md#ListLabels) | **Get** /governance/api/v1/labels | List all labels
[**RemoveResourceLabels**](LabelsAPI.md#RemoveResourceLabels) | **Post** /governance/api/v1/resource-labels/unassign | Remove the labels from resources
[**UpdateLabel**](LabelsAPI.md#UpdateLabel) | **Patch** /governance/api/v1/labels/{labelId} | Update a label



## AssignResourceLabels

> AssignmentResourceLabelsResult AssignResourceLabels(ctx).AssignResourceLabels(assignResourceLabels).Execute()

Assign the labels to resources



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
	assignResourceLabels := *openapiclient.NewAssignResourceLabels([]string{"ResourceOrns_example"}, []string{"LabelValueIds_example"}) // AssignResourceLabels | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelsAPI.AssignResourceLabels(context.Background()).AssignResourceLabels(assignResourceLabels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.AssignResourceLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignResourceLabels`: AssignmentResourceLabelsResult
	fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.AssignResourceLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssignResourceLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assignResourceLabels** | [**AssignResourceLabels**](AssignResourceLabels.md) |  | 

### Return type

[**AssignmentResourceLabelsResult**](AssignmentResourceLabelsResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateLabel

> Label CreateLabel(ctx).LabelCreate(labelCreate).Execute()

Create a label



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
	labelCreate := *openapiclient.NewLabelCreate("Name_example", []openapiclient.LabelValueCreate{*openapiclient.NewLabelValueCreate("Name_example")}) // LabelCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelsAPI.CreateLabel(context.Background()).LabelCreate(labelCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.CreateLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateLabel`: Label
	fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.CreateLabel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **labelCreate** | [**LabelCreate**](LabelCreate.md) |  | 

### Return type

[**Label**](Label.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteLabel

> DeleteLabel(ctx, labelId).Execute()

Delete a label



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
	labelId := "labelId_example" // string | The ID of the label

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LabelsAPI.DeleteLabel(context.Background(), labelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.DeleteLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**labelId** | **string** | The ID of the label | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLabelRequest struct via the builder pattern


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


## GetLabel

> Label GetLabel(ctx, labelId).Execute()

Retrieve a label



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
	labelId := "labelId_example" // string | The ID of the label

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelsAPI.GetLabel(context.Background(), labelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.GetLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLabel`: Label
	fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.GetLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**labelId** | **string** | The ID of the label | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Label**](Label.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLabelResources

> ListResourceLabels ListLabelResources(ctx).Filter(filter).Limit(limit).After(after).Execute()

List all labeled resources



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
	filter := "orn eq "orn:okta:idp:00o11edPwGqbUrsDm0g4:apps:oidc:0oafxqCAJWWGELFTYASJ"" // string | A [filter](https://developer.okta.com/docs/api/#filter) expression that returns entries based on the following properties and supported operators: * `orn`:  supports `eq` * `labelValueId`: supports `eq` * `resourceType`: supports `eq` (The `resourceType` value is taken from the `{objectType}` attribute in the resource [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn). For example, `entitlement-values` for entitlement value resources.)  > **Note:** Query parameter percent encoding is required. See [Special characters](https://developer.okta.com/docs/api/#special-characters). 
	limit := int32(56) // int32 | The maximum number of records returned in a response (optional) (default to 20)
	after := "00u68w6vzKLultXS97g6" // string | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelsAPI.ListLabelResources(context.Background()).Filter(filter).Limit(limit).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.ListLabelResources``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLabelResources`: ListResourceLabels
	fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.ListLabelResources`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLabelResourcesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A [filter](https://developer.okta.com/docs/api/#filter) expression that returns entries based on the following properties and supported operators: * &#x60;orn&#x60;:  supports &#x60;eq&#x60; * &#x60;labelValueId&#x60;: supports &#x60;eq&#x60; * &#x60;resourceType&#x60;: supports &#x60;eq&#x60; (The &#x60;resourceType&#x60; value is taken from the &#x60;{objectType}&#x60; attribute in the resource [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn). For example, &#x60;entitlement-values&#x60; for entitlement value resources.)  &gt; **Note:** Query parameter percent encoding is required. See [Special characters](https://developer.okta.com/docs/api/#special-characters).  | 
 **limit** | **int32** | The maximum number of records returned in a response | [default to 20]
 **after** | **string** | The [pagination](https://developer.okta.com/docs/api/#pagination) cursor that points to the last record of the previous request. | 

### Return type

[**ListResourceLabels**](ListResourceLabels.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListLabels

> ListLabels ListLabels(ctx).Filter(filter).Execute()

List all labels



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
	filter := "values.name sw "SOD"" // string | A [filter](https://developer.okta.com/docs/api/#filter) expression that returns entries based on the following properties and supported operators: * `values.name`:  supports `sw` and `co`  > **Note:** Query parameter percent encoding is required. See [Special characters](https://developer.okta.com/docs/api/#special-characters).  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelsAPI.ListLabels(context.Background()).Filter(filter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.ListLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListLabels`: ListLabels
	fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.ListLabels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | A [filter](https://developer.okta.com/docs/api/#filter) expression that returns entries based on the following properties and supported operators: * &#x60;values.name&#x60;:  supports &#x60;sw&#x60; and &#x60;co&#x60;  &gt; **Note:** Query parameter percent encoding is required. See [Special characters](https://developer.okta.com/docs/api/#special-characters).  | 

### Return type

[**ListLabels**](ListLabels.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveResourceLabels

> RemoveResourceLabels(ctx).AssignResourceLabels(assignResourceLabels).Execute()

Remove the labels from resources



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
	assignResourceLabels := *openapiclient.NewAssignResourceLabels([]string{"ResourceOrns_example"}, []string{"LabelValueIds_example"}) // AssignResourceLabels | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.LabelsAPI.RemoveResourceLabels(context.Background()).AssignResourceLabels(assignResourceLabels).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.RemoveResourceLabels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveResourceLabelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assignResourceLabels** | [**AssignResourceLabels**](AssignResourceLabels.md) |  | 

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


## UpdateLabel

> Label UpdateLabel(ctx, labelId).PatchLabelsInner(patchLabelsInner).Execute()

Update a label



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
	labelId := "labelId_example" // string | The ID of the label
	patchLabelsInner := []openapiclient.PatchLabelsInner{openapiclient.patch_labels_inner{PatchLabelOperation: openapiclient.NewPatchLabelOperation(openapiclient.label-patch-op("REPLACE"), "Path_example", "RefType_example")}} // []PatchLabelsInner | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LabelsAPI.UpdateLabel(context.Background(), labelId).PatchLabelsInner(patchLabelsInner).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LabelsAPI.UpdateLabel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateLabel`: Label
	fmt.Fprintf(os.Stdout, "Response from `LabelsAPI.UpdateLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**labelId** | **string** | The ID of the label | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **patchLabelsInner** | [**[]PatchLabelsInner**](PatchLabelsInner.md) |  | 

### Return type

[**Label**](Label.md)

### Authorization

[OAuth2](../README.md#OAuth2), [ApiKey](../README.md#ApiKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

