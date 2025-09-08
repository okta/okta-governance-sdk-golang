# Criteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the criteria | [optional] 
**Attribute** | Pointer to **string** | Attribute that the criteria applies to | [optional] 
**Operation** | Pointer to **string** | Operation performed on the criteria value | [optional] 
**Value** | Pointer to [**CriteriaValue**](CriteriaValue.md) |  | [optional] 

## Methods

### NewCriteria

`func NewCriteria() *Criteria`

NewCriteria instantiates a new Criteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCriteriaWithDefaults

`func NewCriteriaWithDefaults() *Criteria`

NewCriteriaWithDefaults instantiates a new Criteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Criteria) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Criteria) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Criteria) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Criteria) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAttribute

`func (o *Criteria) GetAttribute() string`

GetAttribute returns the Attribute field if non-nil, zero value otherwise.

### GetAttributeOk

`func (o *Criteria) GetAttributeOk() (*string, bool)`

GetAttributeOk returns a tuple with the Attribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttribute

`func (o *Criteria) SetAttribute(v string)`

SetAttribute sets Attribute field to given value.

### HasAttribute

`func (o *Criteria) HasAttribute() bool`

HasAttribute returns a boolean if a field has been set.

### GetOperation

`func (o *Criteria) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *Criteria) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *Criteria) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *Criteria) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetValue

`func (o *Criteria) GetValue() CriteriaValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Criteria) GetValueOk() (*CriteriaValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Criteria) SetValue(v CriteriaValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *Criteria) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


