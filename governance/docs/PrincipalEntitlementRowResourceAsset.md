# PrincipalEntitlementRowResourceAsset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The identifier of the resource asset | 
**Orn** | **string** | The Okta resource in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn)  See the ORN format for [supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).  | 
**Type** | [**ResourceAssetTypeSparse**](ResourceAssetTypeSparse.md) |  | 
**Name** | **string** | The display name for a resource asset | 
**Description** | Pointer to **string** | The description of a resource asset | [optional] 
**HasChildren** | **bool** | Whether this asset has any children in the resource hierarchy | 

## Methods

### NewPrincipalEntitlementRowResourceAsset

`func NewPrincipalEntitlementRowResourceAsset(id string, orn string, type_ ResourceAssetTypeSparse, name string, hasChildren bool, ) *PrincipalEntitlementRowResourceAsset`

NewPrincipalEntitlementRowResourceAsset instantiates a new PrincipalEntitlementRowResourceAsset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalEntitlementRowResourceAssetWithDefaults

`func NewPrincipalEntitlementRowResourceAssetWithDefaults() *PrincipalEntitlementRowResourceAsset`

NewPrincipalEntitlementRowResourceAssetWithDefaults instantiates a new PrincipalEntitlementRowResourceAsset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrincipalEntitlementRowResourceAsset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrincipalEntitlementRowResourceAsset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrincipalEntitlementRowResourceAsset) SetId(v string)`

SetId sets Id field to given value.


### GetOrn

`func (o *PrincipalEntitlementRowResourceAsset) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *PrincipalEntitlementRowResourceAsset) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *PrincipalEntitlementRowResourceAsset) SetOrn(v string)`

SetOrn sets Orn field to given value.


### GetType

`func (o *PrincipalEntitlementRowResourceAsset) GetType() ResourceAssetTypeSparse`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrincipalEntitlementRowResourceAsset) GetTypeOk() (*ResourceAssetTypeSparse, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrincipalEntitlementRowResourceAsset) SetType(v ResourceAssetTypeSparse)`

SetType sets Type field to given value.


### GetName

`func (o *PrincipalEntitlementRowResourceAsset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrincipalEntitlementRowResourceAsset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrincipalEntitlementRowResourceAsset) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *PrincipalEntitlementRowResourceAsset) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrincipalEntitlementRowResourceAsset) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrincipalEntitlementRowResourceAsset) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrincipalEntitlementRowResourceAsset) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetHasChildren

`func (o *PrincipalEntitlementRowResourceAsset) GetHasChildren() bool`

GetHasChildren returns the HasChildren field if non-nil, zero value otherwise.

### GetHasChildrenOk

`func (o *PrincipalEntitlementRowResourceAsset) GetHasChildrenOk() (*bool, bool)`

GetHasChildrenOk returns a tuple with the HasChildren field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasChildren

`func (o *PrincipalEntitlementRowResourceAsset) SetHasChildren(v bool)`

SetHasChildren sets HasChildren field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


