# EntitlementCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parent** | Pointer to [**TargetResource**](TargetResource.md) |  | [optional] 
**Values** | Pointer to [**[]EntitlementValueWritableProperties**](EntitlementValueWritableProperties.md) |  | [optional] 
**Name** | **string** | The display name for an entitlement property | 
**ExternalValue** | **string** | The value of an entitlement property | 
**Description** | Pointer to **string** | The description of an entitlement property | [optional] 
**MultiValue** | **bool** | The property that determines if the entitlement property can hold multiple values. If this is set to true, the data type is replaced with an array. | 
**DataType** | [**EntitlementPropertyDatatype**](EntitlementPropertyDatatype.md) |  | 

## Methods

### NewEntitlementCreate

`func NewEntitlementCreate(name string, externalValue string, multiValue bool, dataType EntitlementPropertyDatatype, ) *EntitlementCreate`

NewEntitlementCreate instantiates a new EntitlementCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementCreateWithDefaults

`func NewEntitlementCreateWithDefaults() *EntitlementCreate`

NewEntitlementCreateWithDefaults instantiates a new EntitlementCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParent

`func (o *EntitlementCreate) GetParent() TargetResource`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *EntitlementCreate) GetParentOk() (*TargetResource, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *EntitlementCreate) SetParent(v TargetResource)`

SetParent sets Parent field to given value.

### HasParent

`func (o *EntitlementCreate) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetValues

`func (o *EntitlementCreate) GetValues() []EntitlementValueWritableProperties`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *EntitlementCreate) GetValuesOk() (*[]EntitlementValueWritableProperties, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *EntitlementCreate) SetValues(v []EntitlementValueWritableProperties)`

SetValues sets Values field to given value.

### HasValues

`func (o *EntitlementCreate) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetName

`func (o *EntitlementCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EntitlementCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EntitlementCreate) SetName(v string)`

SetName sets Name field to given value.


### GetExternalValue

`func (o *EntitlementCreate) GetExternalValue() string`

GetExternalValue returns the ExternalValue field if non-nil, zero value otherwise.

### GetExternalValueOk

`func (o *EntitlementCreate) GetExternalValueOk() (*string, bool)`

GetExternalValueOk returns a tuple with the ExternalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalValue

`func (o *EntitlementCreate) SetExternalValue(v string)`

SetExternalValue sets ExternalValue field to given value.


### GetDescription

`func (o *EntitlementCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EntitlementCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EntitlementCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EntitlementCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMultiValue

`func (o *EntitlementCreate) GetMultiValue() bool`

GetMultiValue returns the MultiValue field if non-nil, zero value otherwise.

### GetMultiValueOk

`func (o *EntitlementCreate) GetMultiValueOk() (*bool, bool)`

GetMultiValueOk returns a tuple with the MultiValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValue

`func (o *EntitlementCreate) SetMultiValue(v bool)`

SetMultiValue sets MultiValue field to given value.


### GetDataType

`func (o *EntitlementCreate) GetDataType() EntitlementPropertyDatatype`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *EntitlementCreate) GetDataTypeOk() (*EntitlementPropertyDatatype, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *EntitlementCreate) SetDataType(v EntitlementPropertyDatatype)`

SetDataType sets DataType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


