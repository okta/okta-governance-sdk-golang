# ResourceOwnersPatchDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | Pointer to **string** | The operation to be performed for update. | [optional] 
**Path** | Pointer to **string** | The path of the property being updated. Only &#x60;/principalOrn&#x60; can be updated. | [optional] 
**Value** | Pointer to **string** | The Okta user or group &#x60;id&#x60; in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format. The resource can be an user id, or a group id. See [Supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | [optional] 

## Methods

### NewResourceOwnersPatchDataInner

`func NewResourceOwnersPatchDataInner() *ResourceOwnersPatchDataInner`

NewResourceOwnersPatchDataInner instantiates a new ResourceOwnersPatchDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceOwnersPatchDataInnerWithDefaults

`func NewResourceOwnersPatchDataInnerWithDefaults() *ResourceOwnersPatchDataInner`

NewResourceOwnersPatchDataInnerWithDefaults instantiates a new ResourceOwnersPatchDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *ResourceOwnersPatchDataInner) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *ResourceOwnersPatchDataInner) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *ResourceOwnersPatchDataInner) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *ResourceOwnersPatchDataInner) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetPath

`func (o *ResourceOwnersPatchDataInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ResourceOwnersPatchDataInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ResourceOwnersPatchDataInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ResourceOwnersPatchDataInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetValue

`func (o *ResourceOwnersPatchDataInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ResourceOwnersPatchDataInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ResourceOwnersPatchDataInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ResourceOwnersPatchDataInner) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


