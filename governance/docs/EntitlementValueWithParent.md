# EntitlementValueWithParent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The &#x60;id&#x60; of the entitlement value | 
**Name** | **string** | The display name for an entitlement value | 
**ExternalValue** | **string** | The value of an entitlement property value | 
**ExternalId** | Pointer to **string** | The read-only ID of an entitlement property value in the downstream app | [optional] 
**Description** | Pointer to **string** | The description of an entitlement value | [optional] 
**Orn** | **string** | The entitlement value resource, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) | 
**EntitlementId** | **string** | The &#x60;id&#x60; property of an entitlement | 
**ParentResourceOrn** | **string** | The Okta resource, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for [supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | 
**Parent** | [**TargetResource**](TargetResource.md) |  | 
**Labels** | Pointer to [**[]Label**](Label.md) | List of assigned labels | [optional] 

## Methods

### NewEntitlementValueWithParent

`func NewEntitlementValueWithParent(id string, name string, externalValue string, orn string, entitlementId string, parentResourceOrn string, parent TargetResource, ) *EntitlementValueWithParent`

NewEntitlementValueWithParent instantiates a new EntitlementValueWithParent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementValueWithParentWithDefaults

`func NewEntitlementValueWithParentWithDefaults() *EntitlementValueWithParent`

NewEntitlementValueWithParentWithDefaults instantiates a new EntitlementValueWithParent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntitlementValueWithParent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementValueWithParent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementValueWithParent) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EntitlementValueWithParent) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitlementValueWithParent) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitlementValueWithParent) SetName(v string)`

SetName sets Name field to given value.


### GetExternalValue

`func (o *EntitlementValueWithParent) GetExternalValue() string`

GetExternalValue returns the ExternalValue field if non-nil, zero value otherwise.

### GetExternalValueOk

`func (o *EntitlementValueWithParent) GetExternalValueOk() (*string, bool)`

GetExternalValueOk returns a tuple with the ExternalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalValue

`func (o *EntitlementValueWithParent) SetExternalValue(v string)`

SetExternalValue sets ExternalValue field to given value.


### GetExternalId

`func (o *EntitlementValueWithParent) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *EntitlementValueWithParent) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *EntitlementValueWithParent) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *EntitlementValueWithParent) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDescription

`func (o *EntitlementValueWithParent) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitlementValueWithParent) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitlementValueWithParent) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntitlementValueWithParent) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrn

`func (o *EntitlementValueWithParent) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *EntitlementValueWithParent) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *EntitlementValueWithParent) SetOrn(v string)`

SetOrn sets Orn field to given value.


### GetEntitlementId

`func (o *EntitlementValueWithParent) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *EntitlementValueWithParent) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *EntitlementValueWithParent) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.


### GetParentResourceOrn

`func (o *EntitlementValueWithParent) GetParentResourceOrn() string`

GetParentResourceOrn returns the ParentResourceOrn field if non-nil, zero value otherwise.

### GetParentResourceOrnOk

`func (o *EntitlementValueWithParent) GetParentResourceOrnOk() (*string, bool)`

GetParentResourceOrnOk returns a tuple with the ParentResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentResourceOrn

`func (o *EntitlementValueWithParent) SetParentResourceOrn(v string)`

SetParentResourceOrn sets ParentResourceOrn field to given value.


### GetParent

`func (o *EntitlementValueWithParent) GetParent() TargetResource`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *EntitlementValueWithParent) GetParentOk() (*TargetResource, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *EntitlementValueWithParent) SetParent(v TargetResource)`

SetParent sets Parent field to given value.


### GetLabels

`func (o *EntitlementValueWithParent) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EntitlementValueWithParent) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EntitlementValueWithParent) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *EntitlementValueWithParent) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


