# PrincipalEntitlement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentResourceOrn** | Pointer to **string** | The Okta resource, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for [supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | [optional] 
**Parent** | Pointer to [**TargetResource**](TargetResource.md) |  | [optional] 
**Values** | Pointer to [**[]EntitlementValueFull**](EntitlementValueFull.md) | Collection of entitlement values. | [optional] 
**TargetPrincipalOrn** | Pointer to **string** | The Okta user, in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format. | [optional] 
**TargetPrincipal** | Pointer to [**TargetPrincipalFull**](TargetPrincipalFull.md) |  | [optional] 
**Id** | Pointer to **string** | The &#x60;id&#x60; property of an entitlement | [optional] 
**Name** | Pointer to **string** | The display name for an entitlement property | [optional] 
**ExternalValue** | Pointer to **string** | The value of an entitlement property | [optional] 
**Description** | Pointer to **string** | The description of an entitlement property | [optional] 
**MultiValue** | Pointer to **bool** | Indicate if the entitlement property can hold multiple values. If this property is &#x60;true&#x60;, then the &#x60;dataType&#x60; property is set to  &#x60;array&#x60;. | [optional] 
**Required** | Pointer to **bool** | The property that determines if the entitlement property is a required attribute | [optional] 
**DataType** | Pointer to [**EntitlementPropertyDatatype**](EntitlementPropertyDatatype.md) |  | [optional] 

## Methods

### NewPrincipalEntitlement

`func NewPrincipalEntitlement() *PrincipalEntitlement`

NewPrincipalEntitlement instantiates a new PrincipalEntitlement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalEntitlementWithDefaults

`func NewPrincipalEntitlementWithDefaults() *PrincipalEntitlement`

NewPrincipalEntitlementWithDefaults instantiates a new PrincipalEntitlement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentResourceOrn

`func (o *PrincipalEntitlement) GetParentResourceOrn() string`

GetParentResourceOrn returns the ParentResourceOrn field if non-nil, zero value otherwise.

### GetParentResourceOrnOk

`func (o *PrincipalEntitlement) GetParentResourceOrnOk() (*string, bool)`

GetParentResourceOrnOk returns a tuple with the ParentResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentResourceOrn

`func (o *PrincipalEntitlement) SetParentResourceOrn(v string)`

SetParentResourceOrn sets ParentResourceOrn field to given value.

### HasParentResourceOrn

`func (o *PrincipalEntitlement) HasParentResourceOrn() bool`

HasParentResourceOrn returns a boolean if a field has been set.

### GetParent

`func (o *PrincipalEntitlement) GetParent() TargetResource`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *PrincipalEntitlement) GetParentOk() (*TargetResource, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *PrincipalEntitlement) SetParent(v TargetResource)`

SetParent sets Parent field to given value.

### HasParent

`func (o *PrincipalEntitlement) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetValues

`func (o *PrincipalEntitlement) GetValues() []EntitlementValueFull`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PrincipalEntitlement) GetValuesOk() (*[]EntitlementValueFull, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PrincipalEntitlement) SetValues(v []EntitlementValueFull)`

SetValues sets Values field to given value.

### HasValues

`func (o *PrincipalEntitlement) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetTargetPrincipalOrn

`func (o *PrincipalEntitlement) GetTargetPrincipalOrn() string`

GetTargetPrincipalOrn returns the TargetPrincipalOrn field if non-nil, zero value otherwise.

### GetTargetPrincipalOrnOk

`func (o *PrincipalEntitlement) GetTargetPrincipalOrnOk() (*string, bool)`

GetTargetPrincipalOrnOk returns a tuple with the TargetPrincipalOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPrincipalOrn

`func (o *PrincipalEntitlement) SetTargetPrincipalOrn(v string)`

SetTargetPrincipalOrn sets TargetPrincipalOrn field to given value.

### HasTargetPrincipalOrn

`func (o *PrincipalEntitlement) HasTargetPrincipalOrn() bool`

HasTargetPrincipalOrn returns a boolean if a field has been set.

### GetTargetPrincipal

`func (o *PrincipalEntitlement) GetTargetPrincipal() TargetPrincipalFull`

GetTargetPrincipal returns the TargetPrincipal field if non-nil, zero value otherwise.

### GetTargetPrincipalOk

`func (o *PrincipalEntitlement) GetTargetPrincipalOk() (*TargetPrincipalFull, bool)`

GetTargetPrincipalOk returns a tuple with the TargetPrincipal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPrincipal

`func (o *PrincipalEntitlement) SetTargetPrincipal(v TargetPrincipalFull)`

SetTargetPrincipal sets TargetPrincipal field to given value.

### HasTargetPrincipal

`func (o *PrincipalEntitlement) HasTargetPrincipal() bool`

HasTargetPrincipal returns a boolean if a field has been set.

### GetId

`func (o *PrincipalEntitlement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrincipalEntitlement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrincipalEntitlement) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrincipalEntitlement) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PrincipalEntitlement) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrincipalEntitlement) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrincipalEntitlement) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrincipalEntitlement) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalValue

`func (o *PrincipalEntitlement) GetExternalValue() string`

GetExternalValue returns the ExternalValue field if non-nil, zero value otherwise.

### GetExternalValueOk

`func (o *PrincipalEntitlement) GetExternalValueOk() (*string, bool)`

GetExternalValueOk returns a tuple with the ExternalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalValue

`func (o *PrincipalEntitlement) SetExternalValue(v string)`

SetExternalValue sets ExternalValue field to given value.

### HasExternalValue

`func (o *PrincipalEntitlement) HasExternalValue() bool`

HasExternalValue returns a boolean if a field has been set.

### GetDescription

`func (o *PrincipalEntitlement) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PrincipalEntitlement) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PrincipalEntitlement) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PrincipalEntitlement) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMultiValue

`func (o *PrincipalEntitlement) GetMultiValue() bool`

GetMultiValue returns the MultiValue field if non-nil, zero value otherwise.

### GetMultiValueOk

`func (o *PrincipalEntitlement) GetMultiValueOk() (*bool, bool)`

GetMultiValueOk returns a tuple with the MultiValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValue

`func (o *PrincipalEntitlement) SetMultiValue(v bool)`

SetMultiValue sets MultiValue field to given value.

### HasMultiValue

`func (o *PrincipalEntitlement) HasMultiValue() bool`

HasMultiValue returns a boolean if a field has been set.

### GetRequired

`func (o *PrincipalEntitlement) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *PrincipalEntitlement) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *PrincipalEntitlement) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *PrincipalEntitlement) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetDataType

`func (o *PrincipalEntitlement) GetDataType() EntitlementPropertyDatatype`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *PrincipalEntitlement) GetDataTypeOk() (*EntitlementPropertyDatatype, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *PrincipalEntitlement) SetDataType(v EntitlementPropertyDatatype)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *PrincipalEntitlement) HasDataType() bool`

HasDataType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


