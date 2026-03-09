# EntitlementDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | Pointer to [**[]EntitlementValueFull**](EntitlementValueFull.md) | Collection of entitlement values. | [optional] 
**Id** | Pointer to **string** | The &#x60;id&#x60; property of an entitlement | [optional] 
**Name** | Pointer to **string** | The display name for an entitlement property | [optional] 
**ExternalValue** | Pointer to **string** | The value of an entitlement property | [optional] 
**Description** | Pointer to **string** | The description of an entitlement property | [optional] 
**MultiValue** | Pointer to **bool** | Indicate if the entitlement property can hold multiple values. If this property is &#x60;true&#x60;, then the &#x60;dataType&#x60; property is set to  &#x60;array&#x60;. | [optional] 
**Required** | Pointer to **bool** | The property that determines if the entitlement property is a required attribute | [optional] 
**DataType** | Pointer to [**EntitlementPropertyDatatype**](EntitlementPropertyDatatype.md) |  | [optional] 

## Methods

### NewEntitlementDetail

`func NewEntitlementDetail() *EntitlementDetail`

NewEntitlementDetail instantiates a new EntitlementDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementDetailWithDefaults

`func NewEntitlementDetailWithDefaults() *EntitlementDetail`

NewEntitlementDetailWithDefaults instantiates a new EntitlementDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *EntitlementDetail) GetValues() []EntitlementValueFull`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *EntitlementDetail) GetValuesOk() (*[]EntitlementValueFull, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *EntitlementDetail) SetValues(v []EntitlementValueFull)`

SetValues sets Values field to given value.

### HasValues

`func (o *EntitlementDetail) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetId

`func (o *EntitlementDetail) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementDetail) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementDetail) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntitlementDetail) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EntitlementDetail) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitlementDetail) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitlementDetail) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EntitlementDetail) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalValue

`func (o *EntitlementDetail) GetExternalValue() string`

GetExternalValue returns the ExternalValue field if non-nil, zero value otherwise.

### GetExternalValueOk

`func (o *EntitlementDetail) GetExternalValueOk() (*string, bool)`

GetExternalValueOk returns a tuple with the ExternalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalValue

`func (o *EntitlementDetail) SetExternalValue(v string)`

SetExternalValue sets ExternalValue field to given value.

### HasExternalValue

`func (o *EntitlementDetail) HasExternalValue() bool`

HasExternalValue returns a boolean if a field has been set.

### GetDescription

`func (o *EntitlementDetail) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitlementDetail) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitlementDetail) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntitlementDetail) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMultiValue

`func (o *EntitlementDetail) GetMultiValue() bool`

GetMultiValue returns the MultiValue field if non-nil, zero value otherwise.

### GetMultiValueOk

`func (o *EntitlementDetail) GetMultiValueOk() (*bool, bool)`

GetMultiValueOk returns a tuple with the MultiValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValue

`func (o *EntitlementDetail) SetMultiValue(v bool)`

SetMultiValue sets MultiValue field to given value.

### HasMultiValue

`func (o *EntitlementDetail) HasMultiValue() bool`

HasMultiValue returns a boolean if a field has been set.

### GetRequired

`func (o *EntitlementDetail) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *EntitlementDetail) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *EntitlementDetail) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *EntitlementDetail) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetDataType

`func (o *EntitlementDetail) GetDataType() EntitlementPropertyDatatype`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *EntitlementDetail) GetDataTypeOk() (*EntitlementPropertyDatatype, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *EntitlementDetail) SetDataType(v EntitlementPropertyDatatype)`

SetDataType sets DataType field to given value.

### HasDataType

`func (o *EntitlementDetail) HasDataType() bool`

HasDataType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


