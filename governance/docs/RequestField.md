# RequestField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A &#x60;read-only&#x60; field ID.  Useful for specifying requesterFieldValues when adding a request.  | 
**Type** | [**RequestFieldType**](RequestFieldType.md) |  | 
**Choices** | Pointer to **[]string** | Valid choices when type is SELECT or MULTISELECT. | [optional] 
**Required** | **bool** | Whether a value to this field is required to advance the request | [default to true]
**Label** | Pointer to **string** | label of the requester field | [optional] 
**Value** | Pointer to **string** | An admin configured value for this field. Only applies to DURATION fields. | [optional] 
**MaximumValue** | Pointer to **string** | The maximum value allowed for this field. Only applies to DURATION fields. | [optional] 
**ReadOnly** | Pointer to **bool** | Indicates this field is immutable. | [optional] 

## Methods

### NewRequestField

`func NewRequestField(id string, type_ RequestFieldType, required bool, ) *RequestField`

NewRequestField instantiates a new RequestField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestFieldWithDefaults

`func NewRequestFieldWithDefaults() *RequestField`

NewRequestFieldWithDefaults instantiates a new RequestField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestField) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestField) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestField) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *RequestField) GetType() RequestFieldType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestField) GetTypeOk() (*RequestFieldType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestField) SetType(v RequestFieldType)`

SetType sets Type field to given value.


### GetChoices

`func (o *RequestField) GetChoices() []string`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *RequestField) GetChoicesOk() (*[]string, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *RequestField) SetChoices(v []string)`

SetChoices sets Choices field to given value.

### HasChoices

`func (o *RequestField) HasChoices() bool`

HasChoices returns a boolean if a field has been set.

### GetRequired

`func (o *RequestField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *RequestField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *RequestField) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetLabel

`func (o *RequestField) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RequestField) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RequestField) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *RequestField) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetValue

`func (o *RequestField) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RequestField) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RequestField) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *RequestField) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetMaximumValue

`func (o *RequestField) GetMaximumValue() string`

GetMaximumValue returns the MaximumValue field if non-nil, zero value otherwise.

### GetMaximumValueOk

`func (o *RequestField) GetMaximumValueOk() (*string, bool)`

GetMaximumValueOk returns a tuple with the MaximumValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaximumValue

`func (o *RequestField) SetMaximumValue(v string)`

SetMaximumValue sets MaximumValue field to given value.

### HasMaximumValue

`func (o *RequestField) HasMaximumValue() bool`

HasMaximumValue returns a boolean if a field has been set.

### GetReadOnly

`func (o *RequestField) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *RequestField) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *RequestField) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *RequestField) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


