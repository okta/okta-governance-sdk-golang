# EntitlementPropertyFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The &#x60;id&#x60; property of an entitlement | [optional] 
**Name** | Pointer to **string** | The display name for an entitlement property | [optional] 
**ExternalValue** | Pointer to **string** | The value of an entitlement property | [optional] 
**Description** | Pointer to **string** | The description of an entitlement property | [optional] 
**MultiValue** | Pointer to **bool** | The property that determines if the entitlement property can hold multiple values. If this is set to true, the data type is replaced with an array. | [optional] 
**Required** | Pointer to **bool** | The property that determines if the entitlement property is a required attribute | [optional] 
**DataType** | Pointer to [**EntitlementPropertyDatatype**](EntitlementPropertyDatatype.md) |  | [optional] 

## Methods

### NewEntitlementPropertyFull

`func NewEntitlementPropertyFull() *EntitlementPropertyFull`

NewEntitlementPropertyFull instantiates a new EntitlementPropertyFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementPropertyFullWithDefaults

`func NewEntitlementPropertyFullWithDefaults() *EntitlementPropertyFull`

NewEntitlementPropertyFullWithDefaults instantiates a new EntitlementPropertyFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntitlementPropertyFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementPropertyFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementPropertyFull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntitlementPropertyFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EntitlementPropertyFull) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitlementPropertyFull) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitlementPropertyFull) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntitlementPropertyFull) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalValue

`func (o *EntitlementPropertyFull) GetExternalValue() string`

GetExternalValue returns the ExternalValue field if non-nil, zero value otherwise.

### GetExternalValueOk

`func (o *EntitlementPropertyFull) GetExternalValueOk() (*string, bool)`

GetExternalValueOk returns a tuple with the ExternalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalValue

`func (o *EntitlementPropertyFull) SetExternalValue(v string)`

SetExternalValue sets ExternalValue field to given value.

### HasExternalValue

`func (o *EntitlementPropertyFull) HasExternalValue() bool`

HasExternalValue returns a boolean if a field has been set.

### GetDescription

`func (o *EntitlementPropertyFull) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitlementPropertyFull) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitlementPropertyFull) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntitlementPropertyFull) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMultiValue

`func (o *EntitlementPropertyFull) GetMultiValue() bool`

GetMultiValue returns the MultiValue field if non-nil, zero value otherwise.

### GetMultiValueOk

`func (o *EntitlementPropertyFull) GetMultiValueOk() (*bool, bool)`

GetMultiValueOk returns a tuple with the MultiValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValue

`func (o *EntitlementPropertyFull) SetMultiValue(v bool)`

SetMultiValue sets MultiValue field to given value.

### HasMultiValue

`func (o *EntitlementPropertyFull) HasMultiValue() bool`

HasMultiValue returns a boolean if a field has been set.

### GetRequired

`func (o *EntitlementPropertyFull) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *EntitlementPropertyFull) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *EntitlementPropertyFull) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *EntitlementPropertyFull) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetDataType

`func (o *EntitlementPropertyFull) GetDataType() EntitlementPropertyDatatype`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *EntitlementPropertyFull) GetDataTypeOk() (*EntitlementPropertyDatatype, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *EntitlementPropertyFull) SetDataType(v EntitlementPropertyDatatype)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *EntitlementPropertyFull) HasDataType() bool`

HasDataType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


