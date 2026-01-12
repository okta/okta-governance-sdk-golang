# ManagedConnectionAppInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orn** | **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the app instance | 
**Name** | **string** | Display name of the app | 

## Methods

### NewManagedConnectionAppInstance

`func NewManagedConnectionAppInstance(orn string, name string, ) *ManagedConnectionAppInstance`

NewManagedConnectionAppInstance instantiates a new ManagedConnectionAppInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedConnectionAppInstanceWithDefaults

`func NewManagedConnectionAppInstanceWithDefaults() *ManagedConnectionAppInstance`

NewManagedConnectionAppInstanceWithDefaults instantiates a new ManagedConnectionAppInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrn

`func (o *ManagedConnectionAppInstance) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *ManagedConnectionAppInstance) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *ManagedConnectionAppInstance) SetOrn(v string)`

SetOrn sets Orn field to given value.


### GetName

`func (o *ManagedConnectionAppInstance) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ManagedConnectionAppInstance) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ManagedConnectionAppInstance) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


