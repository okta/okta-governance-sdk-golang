# ResourceOwnersPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceOrn** | **string** | The ID of the resource in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format. The resource can be an app, an entitlement value, an entitlement bundle, or a collection. See [supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | 
**Data** | [**[]ResourceOwnersPatchDataInner**](ResourceOwnersPatchDataInner.md) |  | 

## Methods

### NewResourceOwnersPatch

`func NewResourceOwnersPatch(resourceOrn string, data []ResourceOwnersPatchDataInner, ) *ResourceOwnersPatch`

NewResourceOwnersPatch instantiates a new ResourceOwnersPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceOwnersPatchWithDefaults

`func NewResourceOwnersPatchWithDefaults() *ResourceOwnersPatch`

NewResourceOwnersPatchWithDefaults instantiates a new ResourceOwnersPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceOrn

`func (o *ResourceOwnersPatch) GetResourceOrn() string`

GetResourceOrn returns the ResourceOrn field if non-nil, zero value otherwise.

### GetResourceOrnOk

`func (o *ResourceOwnersPatch) GetResourceOrnOk() (*string, bool)`

GetResourceOrnOk returns a tuple with the ResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrn

`func (o *ResourceOwnersPatch) SetResourceOrn(v string)`

SetResourceOrn sets ResourceOrn field to given value.


### GetData

`func (o *ResourceOwnersPatch) GetData() []ResourceOwnersPatchDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ResourceOwnersPatch) GetDataOk() (*[]ResourceOwnersPatchDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ResourceOwnersPatch) SetData(v []ResourceOwnersPatchDataInner)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


