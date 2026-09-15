# PrincipalEntitlementValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The &#x60;id&#x60; of the entitlement value | 
**Name** | **string** | The display name for an entitlement value | 
**ExternalValue** | **string** | The value of an entitlement property value | 
**ExternalId** | Pointer to **string** | The read-only ID of an entitlement property value in the downstream app | [optional] 
**Description** | Pointer to **string** | The description of an entitlement value | [optional] 
**Orn** | Pointer to **string** | The entitlement value resource, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) | [optional] 
**CreatedBy** | Pointer to **string** | The &#x60;id&#x60; of the Okta user who created the resource | [optional] [readonly] 
**Created** | Pointer to **time.Time** | The ISO 8601 formatted date and time when the resource was created | [optional] [readonly] 
**LastUpdated** | Pointer to **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [optional] [readonly] 
**LastUpdatedBy** | Pointer to **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [optional] [readonly] 

## Methods

### NewPrincipalEntitlementValue

`func NewPrincipalEntitlementValue(id string, name string, externalValue string, ) *PrincipalEntitlementValue`

NewPrincipalEntitlementValue instantiates a new PrincipalEntitlementValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalEntitlementValueWithDefaults

`func NewPrincipalEntitlementValueWithDefaults() *PrincipalEntitlementValue`

NewPrincipalEntitlementValueWithDefaults instantiates a new PrincipalEntitlementValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrincipalEntitlementValue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrincipalEntitlementValue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrincipalEntitlementValue) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *PrincipalEntitlementValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrincipalEntitlementValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrincipalEntitlementValue) SetName(v string)`

SetName sets Name field to given value.


### GetExternalValue

`func (o *PrincipalEntitlementValue) GetExternalValue() string`

GetExternalValue returns the ExternalValue field if non-nil, zero value otherwise.

### GetExternalValueOk

`func (o *PrincipalEntitlementValue) GetExternalValueOk() (*string, bool)`

GetExternalValueOk returns a tuple with the ExternalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalValue

`func (o *PrincipalEntitlementValue) SetExternalValue(v string)`

SetExternalValue sets ExternalValue field to given value.


### GetExternalId

`func (o *PrincipalEntitlementValue) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *PrincipalEntitlementValue) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *PrincipalEntitlementValue) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *PrincipalEntitlementValue) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetDescription

`func (o *PrincipalEntitlementValue) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrincipalEntitlementValue) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrincipalEntitlementValue) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrincipalEntitlementValue) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOrn

`func (o *PrincipalEntitlementValue) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *PrincipalEntitlementValue) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *PrincipalEntitlementValue) SetOrn(v string)`

SetOrn sets Orn field to given value.

### HasOrn

`func (o *PrincipalEntitlementValue) HasOrn() bool`

HasOrn returns a boolean if a field has been set.

### GetCreatedBy

`func (o *PrincipalEntitlementValue) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *PrincipalEntitlementValue) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *PrincipalEntitlementValue) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *PrincipalEntitlementValue) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetCreated

`func (o *PrincipalEntitlementValue) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *PrincipalEntitlementValue) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *PrincipalEntitlementValue) SetCreated(v time.Time)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *PrincipalEntitlementValue) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *PrincipalEntitlementValue) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *PrincipalEntitlementValue) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *PrincipalEntitlementValue) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *PrincipalEntitlementValue) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastUpdatedBy

`func (o *PrincipalEntitlementValue) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *PrincipalEntitlementValue) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *PrincipalEntitlementValue) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.

### HasLastUpdatedBy

`func (o *PrincipalEntitlementValue) HasLastUpdatedBy() bool`

HasLastUpdatedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


