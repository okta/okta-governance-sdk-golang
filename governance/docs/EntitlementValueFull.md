# EntitlementValueFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewEntitlementValueFull

`func NewEntitlementValueFull() *EntitlementValueFull`

NewEntitlementValueFull instantiates a new EntitlementValueFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementValueFullWithDefaults

`func NewEntitlementValueFullWithDefaults() *EntitlementValueFull`

NewEntitlementValueFullWithDefaults instantiates a new EntitlementValueFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntitlementValueFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementValueFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementValueFull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntitlementValueFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EntitlementValueFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitlementValueFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitlementValueFull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntitlementValueFull) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalValue

`func (o *EntitlementValueFull) GetExternalValue() string`

GetExternalValue returns the ExternalValue field if non-nil, zero value otherwise.

### GetExternalValueOk

`func (o *EntitlementValueFull) GetExternalValueOk() (*string, bool)`

GetExternalValueOk returns a tuple with the ExternalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalValue

`func (o *EntitlementValueFull) SetExternalValue(v string)`

SetExternalValue sets ExternalValue field to given value.

### HasExternalValue

`func (o *EntitlementValueFull) HasExternalValue() bool`

HasExternalValue returns a boolean if a field has been set.

### GetExternalId

`func (o *EntitlementValueFull) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *EntitlementValueFull) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *EntitlementValueFull) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *EntitlementValueFull) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDescription

`func (o *EntitlementValueFull) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitlementValueFull) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitlementValueFull) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntitlementValueFull) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrn

`func (o *EntitlementValueFull) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *EntitlementValueFull) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *EntitlementValueFull) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *EntitlementValueFull) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetCreatedBy

`func (o *EntitlementValueFull) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EntitlementValueFull) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EntitlementValueFull) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *EntitlementValueFull) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreated

`func (o *EntitlementValueFull) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EntitlementValueFull) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EntitlementValueFull) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *EntitlementValueFull) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *EntitlementValueFull) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *EntitlementValueFull) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *EntitlementValueFull) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *EntitlementValueFull) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *EntitlementValueFull) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *EntitlementValueFull) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *EntitlementValueFull) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *EntitlementValueFull) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


