# CollectionResourceCreatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entitlements** | Pointer to [**[]EntitlementCreatable**](EntitlementCreatable.md) | Collection of entitlements and associated value identifiers | [optional] 
**ResourceOrn** | **string** | The ORN identifier for a specific app. Other resource types aren&#39;t supported.  See the [supported-resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources) endpoint for reference.  | 

## Methods

### NewCollectionResourceCreatable

`func NewCollectionResourceCreatable(resourceOrn string, ) *CollectionResourceCreatable`

NewCollectionResourceCreatable instantiates a new CollectionResourceCreatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCollectionResourceCreatableWithDefaults

`func NewCollectionResourceCreatableWithDefaults() *CollectionResourceCreatable`

NewCollectionResourceCreatableWithDefaults instantiates a new CollectionResourceCreatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntitlements

`func (o *CollectionResourceCreatable) GetEntitlements() []EntitlementCreatable`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *CollectionResourceCreatable) GetEntitlementsOk() (*[]EntitlementCreatable, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *CollectionResourceCreatable) SetEntitlements(v []EntitlementCreatable)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *CollectionResourceCreatable) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetResourceOrn

`func (o *CollectionResourceCreatable) GetResourceOrn() string`

GetResourceOrn returns the ResourceOrn field if non-nil, zero value otherwise.

### GetResourceOrnOk

`func (o *CollectionResourceCreatable) GetResourceOrnOk() (*string, bool)`

GetResourceOrnOk returns a tuple with the ResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrn

`func (o *CollectionResourceCreatable) SetResourceOrn(v string)`

SetResourceOrn sets ResourceOrn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


