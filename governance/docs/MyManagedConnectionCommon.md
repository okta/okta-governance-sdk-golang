# MyManagedConnectionCommon

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the managed connection | [optional] 
**Orn** | Pointer to **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the managed connection | [optional] 
**Status** | Pointer to **string** | The status of the connection | [optional] 

## Methods

### NewMyManagedConnectionCommon

`func NewMyManagedConnectionCommon() *MyManagedConnectionCommon`

NewMyManagedConnectionCommon instantiates a new MyManagedConnectionCommon object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMyManagedConnectionCommonWithDefaults

`func NewMyManagedConnectionCommonWithDefaults() *MyManagedConnectionCommon`

NewMyManagedConnectionCommonWithDefaults instantiates a new MyManagedConnectionCommon object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MyManagedConnectionCommon) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MyManagedConnectionCommon) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MyManagedConnectionCommon) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MyManagedConnectionCommon) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrn

`func (o *MyManagedConnectionCommon) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *MyManagedConnectionCommon) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *MyManagedConnectionCommon) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *MyManagedConnectionCommon) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetStatus

`func (o *MyManagedConnectionCommon) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MyManagedConnectionCommon) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MyManagedConnectionCommon) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MyManagedConnectionCommon) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


