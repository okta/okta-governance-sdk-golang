# EntitlementValue2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**EntitlementLink**](EntitlementLink.md) |  | 
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

### NewEntitlementValue2

`func NewEntitlementValue2(links EntitlementLink, id string, name string, externalValue string, orn string, entitlementId string, parentResourceOrn string, parent TargetResource, ) *EntitlementValue2`

NewEntitlementValue2 instantiates a new EntitlementValue2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementValue2WithDefaults

`func NewEntitlementValue2WithDefaults() *EntitlementValue2`

NewEntitlementValue2WithDefaults instantiates a new EntitlementValue2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *EntitlementValue2) GetLinks() EntitlementLink`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EntitlementValue2) GetLinksOk() (*EntitlementLink, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EntitlementValue2) SetLinks(v EntitlementLink)`

SetLinks sets Links field to given value.


### GetId

`func (o *EntitlementValue2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementValue2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementValue2) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *EntitlementValue2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitlementValue2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitlementValue2) SetName(v string)`

SetName sets Name field to given value.


### GetExternalValue

`func (o *EntitlementValue2) GetExternalValue() string`

GetExternalValue returns the ExternalValue field if non-nil, zero value otherwise.

### GetExternalValueOk

`func (o *EntitlementValue2) GetExternalValueOk() (*string, bool)`

GetExternalValueOk returns a tuple with the ExternalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalValue

`func (o *EntitlementValue2) SetExternalValue(v string)`

SetExternalValue sets ExternalValue field to given value.


### GetExternalId

`func (o *EntitlementValue2) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *EntitlementValue2) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *EntitlementValue2) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *EntitlementValue2) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDescription

`func (o *EntitlementValue2) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitlementValue2) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitlementValue2) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntitlementValue2) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrn

`func (o *EntitlementValue2) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *EntitlementValue2) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *EntitlementValue2) SetOrn(v string)`

SetOrn sets Orn field to given value.


### GetEntitlementId

`func (o *EntitlementValue2) GetEntitlementId() string`

GetEntitlementId returns the EntitlementId field if non-nil, zero value otherwise.

### GetEntitlementIdOk

`func (o *EntitlementValue2) GetEntitlementIdOk() (*string, bool)`

GetEntitlementIdOk returns a tuple with the EntitlementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlementId

`func (o *EntitlementValue2) SetEntitlementId(v string)`

SetEntitlementId sets EntitlementId field to given value.


### GetParentResourceOrn

`func (o *EntitlementValue2) GetParentResourceOrn() string`

GetParentResourceOrn returns the ParentResourceOrn field if non-nil, zero value otherwise.

### GetParentResourceOrnOk

`func (o *EntitlementValue2) GetParentResourceOrnOk() (*string, bool)`

GetParentResourceOrnOk returns a tuple with the ParentResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentResourceOrn

`func (o *EntitlementValue2) SetParentResourceOrn(v string)`

SetParentResourceOrn sets ParentResourceOrn field to given value.


### GetParent

`func (o *EntitlementValue2) GetParent() TargetResource`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *EntitlementValue2) GetParentOk() (*TargetResource, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *EntitlementValue2) SetParent(v TargetResource)`

SetParent sets Parent field to given value.


### GetLabels

`func (o *EntitlementValue2) GetLabels() []Label`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *EntitlementValue2) GetLabelsOk() (*[]Label, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *EntitlementValue2) SetLabels(v []Label)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *EntitlementValue2) HasLabels() bool`

HasLabels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


