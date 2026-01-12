# GrantedEntitlements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | [**[]GrantedEntitlementValue**](GrantedEntitlementValue.md) |  | 
**Id** | **string** | The &#x60;id&#x60; property of an entitlement | 
**Name** | **string** | The display name for an entitlement property | 
**ExternalValue** | **string** | The value of an entitlement property | 
**Description** | Pointer to **string** | The description of an entitlement property | [optional] 
**MultiValue** | **bool** | Indicate if the entitlement property can hold multiple values. If this property is &#x60;true&#x60;, then the &#x60;dataType&#x60; property is set to  &#x60;array&#x60;. | 
**Required** | **bool** | The property that determines if the entitlement property is a required attribute | 
**DataType** | [**EntitlementPropertyDatatype**](EntitlementPropertyDatatype.md) |  | 

## Methods

### NewGrantedEntitlements

`func NewGrantedEntitlements(values []GrantedEntitlementValue, id string, name string, externalValue string, multiValue bool, required bool, dataType EntitlementPropertyDatatype, ) *GrantedEntitlements`

NewGrantedEntitlements instantiates a new GrantedEntitlements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGrantedEntitlementsWithDefaults

`func NewGrantedEntitlementsWithDefaults() *GrantedEntitlements`

NewGrantedEntitlementsWithDefaults instantiates a new GrantedEntitlements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *GrantedEntitlements) GetValues() []GrantedEntitlementValue`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GrantedEntitlements) GetValuesOk() (*[]GrantedEntitlementValue, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GrantedEntitlements) SetValues(v []GrantedEntitlementValue)`

SetValues sets Values field to given value.


### GetId

`func (o *GrantedEntitlements) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GrantedEntitlements) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GrantedEntitlements) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GrantedEntitlements) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GrantedEntitlements) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GrantedEntitlements) SetName(v string)`

SetName sets Name field to given value.


### GetExternalValue

`func (o *GrantedEntitlements) GetExternalValue() string`

GetExternalValue returns the ExternalValue field if non-nil, zero value otherwise.

### GetExternalValueOk

`func (o *GrantedEntitlements) GetExternalValueOk() (*string, bool)`

GetExternalValueOk returns a tuple with the ExternalValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalValue

`func (o *GrantedEntitlements) SetExternalValue(v string)`

SetExternalValue sets ExternalValue field to given value.


### GetDescription

`func (o *GrantedEntitlements) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GrantedEntitlements) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GrantedEntitlements) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GrantedEntitlements) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMultiValue

`func (o *GrantedEntitlements) GetMultiValue() bool`

GetMultiValue returns the MultiValue field if non-nil, zero value otherwise.

### GetMultiValueOk

`func (o *GrantedEntitlements) GetMultiValueOk() (*bool, bool)`

GetMultiValueOk returns a tuple with the MultiValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiValue

`func (o *GrantedEntitlements) SetMultiValue(v bool)`

SetMultiValue sets MultiValue field to given value.


### GetRequired

`func (o *GrantedEntitlements) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *GrantedEntitlements) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *GrantedEntitlements) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetDataType

`func (o *GrantedEntitlements) GetDataType() EntitlementPropertyDatatype`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *GrantedEntitlements) GetDataTypeOk() (*EntitlementPropertyDatatype, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *GrantedEntitlements) SetDataType(v EntitlementPropertyDatatype)`

SetDataType sets DataType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


