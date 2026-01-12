# GrantedEntitlementValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The &#x60;id&#x60; of the entitlement value | 
**Name** | **string** | The display name for an entitlement value | 
**ExternalValue** | **string** | The value of an entitlement property value | 
**Orn** | **string** | The entitlement value resource, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) | 
**Effective** | Pointer to **bool** | A granted entitlement may not be effective if the same entitlement is granted by a higher priority additional grant | [optional] 

## Methods

### NewGrantedEntitlementValue

`func NewGrantedEntitlementValue(id string, name string, externalValue string, orn string, ) *GrantedEntitlementValue`

NewGrantedEntitlementValue instantiates a new GrantedEntitlementValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantedEntitlementValueWithDefaults

`func NewGrantedEntitlementValueWithDefaults() *GrantedEntitlementValue`

NewGrantedEntitlementValueWithDefaults instantiates a new GrantedEntitlementValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GrantedEntitlementValue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GrantedEntitlementValue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GrantedEntitlementValue) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GrantedEntitlementValue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GrantedEntitlementValue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GrantedEntitlementValue) SetName(v string)`

SetName sets Name field to given value.


### GetExternalValue

`func (o *GrantedEntitlementValue) GetExternalValue() string`

GetExternalValue returns the ExternalValue field if non-nil, zero value otherwise.

### GetExternalValueOk

`func (o *GrantedEntitlementValue) GetExternalValueOk() (*string, bool)`

GetExternalValueOk returns a tuple with the ExternalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalValue

`func (o *GrantedEntitlementValue) SetExternalValue(v string)`

SetExternalValue sets ExternalValue field to given value.


### GetOrn

`func (o *GrantedEntitlementValue) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *GrantedEntitlementValue) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *GrantedEntitlementValue) SetOrn(v string)`

SetOrn sets Orn field to given value.


### GetEffective

`func (o *GrantedEntitlementValue) GetEffective() bool`

GetEffective returns the Effective field if non-nil, zero value otherwise.

### GetEffectiveOk

`func (o *GrantedEntitlementValue) GetEffectiveOk() (*bool, bool)`

GetEffectiveOk returns a tuple with the Effective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffective

`func (o *GrantedEntitlementValue) SetEffective(v bool)`

SetEffective sets Effective field to given value.

### HasEffective

`func (o *GrantedEntitlementValue) HasEffective() bool`

HasEffective returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


