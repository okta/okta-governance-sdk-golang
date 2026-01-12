# ResourceOwnersUpdatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalOrns** | Pointer to **[]string** | Owners of the resource. If no owners are provided (empty list), then all current owners are removed. | [optional] 
**ResourceOrns** | **[]string** | The resources to assign owners | 

## Methods

### NewResourceOwnersUpdatable

`func NewResourceOwnersUpdatable(resourceOrns []string, ) *ResourceOwnersUpdatable`

NewResourceOwnersUpdatable instantiates a new ResourceOwnersUpdatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceOwnersUpdatableWithDefaults

`func NewResourceOwnersUpdatableWithDefaults() *ResourceOwnersUpdatable`

NewResourceOwnersUpdatableWithDefaults instantiates a new ResourceOwnersUpdatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalOrns

`func (o *ResourceOwnersUpdatable) GetPrincipalOrns() []string`

GetPrincipalOrns returns the PrincipalOrns field if non-nil, zero value otherwise.

### GetPrincipalOrnsOk

`func (o *ResourceOwnersUpdatable) GetPrincipalOrnsOk() (*[]string, bool)`

GetPrincipalOrnsOk returns a tuple with the PrincipalOrns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalOrns

`func (o *ResourceOwnersUpdatable) SetPrincipalOrns(v []string)`

SetPrincipalOrns sets PrincipalOrns field to given value.

### HasPrincipalOrns

`func (o *ResourceOwnersUpdatable) HasPrincipalOrns() bool`

HasPrincipalOrns returns a boolean if a field has been set.

### GetResourceOrns

`func (o *ResourceOwnersUpdatable) GetResourceOrns() []string`

GetResourceOrns returns the ResourceOrns field if non-nil, zero value otherwise.

### GetResourceOrnsOk

`func (o *ResourceOwnersUpdatable) GetResourceOrnsOk() (*[]string, bool)`

GetResourceOrnsOk returns a tuple with the ResourceOrns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrns

`func (o *ResourceOwnersUpdatable) SetResourceOrns(v []string)`

SetResourceOrns sets ResourceOrns field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


