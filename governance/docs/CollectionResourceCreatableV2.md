# CollectionResourceCreatableV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entitlements** | Pointer to [**[]EntitlementCreatable**](EntitlementCreatable.md) | Collection of entitlements and associated value identifiers | [optional] 
**ResourceOrn** | **string** | The ORN identifier for a collection resource (app, group, or push group).  See the [supported-resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources) endpoint.  | 

## Methods

### NewCollectionResourceCreatableV2

`func NewCollectionResourceCreatableV2(resourceOrn string, ) *CollectionResourceCreatableV2`

NewCollectionResourceCreatableV2 instantiates a new CollectionResourceCreatableV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResourceCreatableV2WithDefaults

`func NewCollectionResourceCreatableV2WithDefaults() *CollectionResourceCreatableV2`

NewCollectionResourceCreatableV2WithDefaults instantiates a new CollectionResourceCreatableV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlements

`func (o *CollectionResourceCreatableV2) GetEntitlements() []EntitlementCreatable`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *CollectionResourceCreatableV2) GetEntitlementsOk() (*[]EntitlementCreatable, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *CollectionResourceCreatableV2) SetEntitlements(v []EntitlementCreatable)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *CollectionResourceCreatableV2) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetResourceOrn

`func (o *CollectionResourceCreatableV2) GetResourceOrn() string`

GetResourceOrn returns the ResourceOrn field if non-nil, zero value otherwise.

### GetResourceOrnOk

`func (o *CollectionResourceCreatableV2) GetResourceOrnOk() (*string, bool)`

GetResourceOrnOk returns a tuple with the ResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrn

`func (o *CollectionResourceCreatableV2) SetResourceOrn(v string)`

SetResourceOrn sets ResourceOrn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


