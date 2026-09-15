# ScopedEntitlementResourceAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The identifier of the resource asset | 
**Orn** | **string** | The Okta resource in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn)  See the ORN format for [supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).  | 
**Type** | [**ResourceAssetTypeSparse**](ResourceAssetTypeSparse.md) |  | 
**Name** | **string** | The display name for a resource asset | 

## Methods

### NewScopedEntitlementResourceAsset

`func NewScopedEntitlementResourceAsset(id string, orn string, type_ ResourceAssetTypeSparse, name string, ) *ScopedEntitlementResourceAsset`

NewScopedEntitlementResourceAsset instantiates a new ScopedEntitlementResourceAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopedEntitlementResourceAssetWithDefaults

`func NewScopedEntitlementResourceAssetWithDefaults() *ScopedEntitlementResourceAsset`

NewScopedEntitlementResourceAssetWithDefaults instantiates a new ScopedEntitlementResourceAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ScopedEntitlementResourceAsset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScopedEntitlementResourceAsset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScopedEntitlementResourceAsset) SetId(v string)`

SetId sets Id field to given value.


### GetOrn

`func (o *ScopedEntitlementResourceAsset) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *ScopedEntitlementResourceAsset) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *ScopedEntitlementResourceAsset) SetOrn(v string)`

SetOrn sets Orn field to given value.


### GetType

`func (o *ScopedEntitlementResourceAsset) GetType() ResourceAssetTypeSparse`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScopedEntitlementResourceAsset) GetTypeOk() (*ResourceAssetTypeSparse, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScopedEntitlementResourceAsset) SetType(v ResourceAssetTypeSparse)`

SetType sets Type field to given value.


### GetName

`func (o *ScopedEntitlementResourceAsset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ScopedEntitlementResourceAsset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ScopedEntitlementResourceAsset) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


