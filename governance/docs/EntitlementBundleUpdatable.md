# EntitlementBundleUpdatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the object | 
**TargetResourceOrn** | **string** | The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | 
**Target** | [**TargetResource**](TargetResource.md) |  | 
**Status** | Pointer to [**EntitlementBundleStatus**](EntitlementBundleStatus.md) |  | [optional] 
**Entitlements** | Pointer to [**[]EntitlementCreatable**](EntitlementCreatable.md) | Collection of entitlements and associated value identifiers | [optional] 
**Name** | **string** | The unique name of the entitlement bundle | 
**Description** | Pointer to **string** | The human-readable description | [optional] 

## Methods

### NewEntitlementBundleUpdatable

`func NewEntitlementBundleUpdatable(id string, targetResourceOrn string, target TargetResource, name string, ) *EntitlementBundleUpdatable`

NewEntitlementBundleUpdatable instantiates a new EntitlementBundleUpdatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementBundleUpdatableWithDefaults

`func NewEntitlementBundleUpdatableWithDefaults() *EntitlementBundleUpdatable`

NewEntitlementBundleUpdatableWithDefaults instantiates a new EntitlementBundleUpdatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntitlementBundleUpdatable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementBundleUpdatable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementBundleUpdatable) SetId(v string)`

SetId sets Id field to given value.


### GetTargetResourceOrn

`func (o *EntitlementBundleUpdatable) GetTargetResourceOrn() string`

GetTargetResourceOrn returns the TargetResourceOrn field if non-nil, zero value otherwise.

### GetTargetResourceOrnOk

`func (o *EntitlementBundleUpdatable) GetTargetResourceOrnOk() (*string, bool)`

GetTargetResourceOrnOk returns a tuple with the TargetResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetResourceOrn

`func (o *EntitlementBundleUpdatable) SetTargetResourceOrn(v string)`

SetTargetResourceOrn sets TargetResourceOrn field to given value.


### GetTarget

`func (o *EntitlementBundleUpdatable) GetTarget() TargetResource`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *EntitlementBundleUpdatable) GetTargetOk() (*TargetResource, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *EntitlementBundleUpdatable) SetTarget(v TargetResource)`

SetTarget sets Target field to given value.


### GetStatus

`func (o *EntitlementBundleUpdatable) GetStatus() EntitlementBundleStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EntitlementBundleUpdatable) GetStatusOk() (*EntitlementBundleStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EntitlementBundleUpdatable) SetStatus(v EntitlementBundleStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EntitlementBundleUpdatable) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEntitlements

`func (o *EntitlementBundleUpdatable) GetEntitlements() []EntitlementCreatable`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *EntitlementBundleUpdatable) GetEntitlementsOk() (*[]EntitlementCreatable, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *EntitlementBundleUpdatable) SetEntitlements(v []EntitlementCreatable)`

SetEntitlements sets Entitlements field to given value.

### HasEntitlements

`func (o *EntitlementBundleUpdatable) HasEntitlements() bool`

HasEntitlements returns a boolean if a field has been set.

### GetName

`func (o *EntitlementBundleUpdatable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitlementBundleUpdatable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitlementBundleUpdatable) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *EntitlementBundleUpdatable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitlementBundleUpdatable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitlementBundleUpdatable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntitlementBundleUpdatable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


