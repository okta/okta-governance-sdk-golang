# ResourceConnectionVaultedSecret

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orn** | **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the vaulted secret | 
**Name** | **string** | Display name of the secret | 

## Methods

### NewResourceConnectionVaultedSecret

`func NewResourceConnectionVaultedSecret(orn string, name string, ) *ResourceConnectionVaultedSecret`

NewResourceConnectionVaultedSecret instantiates a new ResourceConnectionVaultedSecret object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceConnectionVaultedSecretWithDefaults

`func NewResourceConnectionVaultedSecretWithDefaults() *ResourceConnectionVaultedSecret`

NewResourceConnectionVaultedSecretWithDefaults instantiates a new ResourceConnectionVaultedSecret object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrn

`func (o *ResourceConnectionVaultedSecret) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *ResourceConnectionVaultedSecret) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *ResourceConnectionVaultedSecret) SetOrn(v string)`

SetOrn sets Orn field to given value.


### GetName

`func (o *ResourceConnectionVaultedSecret) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceConnectionVaultedSecret) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceConnectionVaultedSecret) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


