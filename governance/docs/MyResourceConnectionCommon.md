# MyResourceConnectionCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the resource connection | [optional] 
**Orn** | Pointer to **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the resource connection | [optional] 
**Status** | Pointer to **string** | The status of the connection | [optional] 

## Methods

### NewMyResourceConnectionCommon

`func NewMyResourceConnectionCommon() *MyResourceConnectionCommon`

NewMyResourceConnectionCommon instantiates a new MyResourceConnectionCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyResourceConnectionCommonWithDefaults

`func NewMyResourceConnectionCommonWithDefaults() *MyResourceConnectionCommon`

NewMyResourceConnectionCommonWithDefaults instantiates a new MyResourceConnectionCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MyResourceConnectionCommon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MyResourceConnectionCommon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MyResourceConnectionCommon) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MyResourceConnectionCommon) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrn

`func (o *MyResourceConnectionCommon) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *MyResourceConnectionCommon) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *MyResourceConnectionCommon) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *MyResourceConnectionCommon) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetStatus

`func (o *MyResourceConnectionCommon) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MyResourceConnectionCommon) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MyResourceConnectionCommon) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MyResourceConnectionCommon) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


