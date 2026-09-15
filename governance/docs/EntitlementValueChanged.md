# EntitlementValueChanged

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChangeType** | Pointer to **string** | Type of change that occurred to the entitlement | [optional] 
**Id** | Pointer to **string** | The &#x60;id&#x60; of the entitlement value | [optional] 
**Name** | Pointer to **string** | The display name for an entitlement value | [optional] 
**ExternalValue** | Pointer to **string** | The value of an entitlement property value | [optional] 
**ExternalId** | Pointer to **string** | The read-only ID of an entitlement property value in the downstream app | [optional] 
**Description** | Pointer to **string** | The description of an entitlement value | [optional] 
**Orn** | Pointer to **string** | The entitlement value resource, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) | [optional] 
**CreatedBy** | Pointer to **string** | The &#x60;id&#x60; of the Okta user who created the resource | [optional] [readonly] 
**Created** | Pointer to **time.Time** | The ISO 8601 formatted date and time when the resource was created | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [optional] [readonly] 
**LastUpdatedBy** | Pointer to **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [optional] [readonly] 

## Methods

### NewEntitlementValueChanged

`func NewEntitlementValueChanged() *EntitlementValueChanged`

NewEntitlementValueChanged instantiates a new EntitlementValueChanged object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementValueChangedWithDefaults

`func NewEntitlementValueChangedWithDefaults() *EntitlementValueChanged`

NewEntitlementValueChangedWithDefaults instantiates a new EntitlementValueChanged object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChangeType

`func (o *EntitlementValueChanged) GetChangeType() string`

GetChangeType returns the ChangeType field if non-nil, zero value otherwise.

### GetChangeTypeOk

`func (o *EntitlementValueChanged) GetChangeTypeOk() (*string, bool)`

GetChangeTypeOk returns a tuple with the ChangeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeType

`func (o *EntitlementValueChanged) SetChangeType(v string)`

SetChangeType sets ChangeType field to given value.

### HasChangeType

`func (o *EntitlementValueChanged) HasChangeType() bool`

HasChangeType returns a boolean if a field has been set.

### GetId

`func (o *EntitlementValueChanged) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementValueChanged) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementValueChanged) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntitlementValueChanged) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EntitlementValueChanged) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitlementValueChanged) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitlementValueChanged) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntitlementValueChanged) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalValue

`func (o *EntitlementValueChanged) GetExternalValue() string`

GetExternalValue returns the ExternalValue field if non-nil, zero value otherwise.

### GetExternalValueOk

`func (o *EntitlementValueChanged) GetExternalValueOk() (*string, bool)`

GetExternalValueOk returns a tuple with the ExternalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalValue

`func (o *EntitlementValueChanged) SetExternalValue(v string)`

SetExternalValue sets ExternalValue field to given value.

### HasExternalValue

`func (o *EntitlementValueChanged) HasExternalValue() bool`

HasExternalValue returns a boolean if a field has been set.

### GetExternalId

`func (o *EntitlementValueChanged) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *EntitlementValueChanged) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *EntitlementValueChanged) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *EntitlementValueChanged) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDescription

`func (o *EntitlementValueChanged) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitlementValueChanged) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitlementValueChanged) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntitlementValueChanged) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrn

`func (o *EntitlementValueChanged) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *EntitlementValueChanged) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *EntitlementValueChanged) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *EntitlementValueChanged) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetCreatedBy

`func (o *EntitlementValueChanged) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EntitlementValueChanged) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EntitlementValueChanged) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *EntitlementValueChanged) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreated

`func (o *EntitlementValueChanged) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EntitlementValueChanged) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EntitlementValueChanged) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *EntitlementValueChanged) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *EntitlementValueChanged) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *EntitlementValueChanged) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *EntitlementValueChanged) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *EntitlementValueChanged) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *EntitlementValueChanged) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *EntitlementValueChanged) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *EntitlementValueChanged) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *EntitlementValueChanged) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


