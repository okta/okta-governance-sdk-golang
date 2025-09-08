# EntitlementBundleCreatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Target** | [**TargetResource**](TargetResource.md) |  | 
**Entitlements** | [**[]EntitlementCreatable**](EntitlementCreatable.md) | Collection of entitlements and associated value identifiers | 
**Name** | **string** | The unique name of the entitlement bundle | 
**Description** | Pointer to **string** | The human-readable description | [optional] 

## Methods

### NewEntitlementBundleCreatable

`func NewEntitlementBundleCreatable(target TargetResource, entitlements []EntitlementCreatable, name string, ) *EntitlementBundleCreatable`

NewEntitlementBundleCreatable instantiates a new EntitlementBundleCreatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementBundleCreatableWithDefaults

`func NewEntitlementBundleCreatableWithDefaults() *EntitlementBundleCreatable`

NewEntitlementBundleCreatableWithDefaults instantiates a new EntitlementBundleCreatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTarget

`func (o *EntitlementBundleCreatable) GetTarget() TargetResource`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *EntitlementBundleCreatable) GetTargetOk() (*TargetResource, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *EntitlementBundleCreatable) SetTarget(v TargetResource)`

SetTarget sets Target field to given value.


### GetEntitlements

`func (o *EntitlementBundleCreatable) GetEntitlements() []EntitlementCreatable`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *EntitlementBundleCreatable) GetEntitlementsOk() (*[]EntitlementCreatable, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *EntitlementBundleCreatable) SetEntitlements(v []EntitlementCreatable)`

SetEntitlements sets Entitlements field to given value.


### GetName

`func (o *EntitlementBundleCreatable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitlementBundleCreatable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitlementBundleCreatable) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EntitlementBundleCreatable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitlementBundleCreatable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitlementBundleCreatable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntitlementBundleCreatable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


