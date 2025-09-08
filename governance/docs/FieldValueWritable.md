# FieldValueWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The &#x60;id&#x60; of a &#x60;requesterField&#x60; in the related request type&#39;s &#x60;requestSettings.requesterFields&#x60;.  | 
**Value** | [**FieldValueWritableAllowedValues**](FieldValueWritableAllowedValues.md) |  | 

## Methods

### NewFieldValueWritable

`func NewFieldValueWritable(id string, value FieldValueWritableAllowedValues, ) *FieldValueWritable`

NewFieldValueWritable instantiates a new FieldValueWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldValueWritableWithDefaults

`func NewFieldValueWritableWithDefaults() *FieldValueWritable`

NewFieldValueWritableWithDefaults instantiates a new FieldValueWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FieldValueWritable) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FieldValueWritable) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FieldValueWritable) SetId(v string)`

SetId sets Id field to given value.


### GetValue

`func (o *FieldValueWritable) GetValue() FieldValueWritableAllowedValues`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FieldValueWritable) GetValueOk() (*FieldValueWritableAllowedValues, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FieldValueWritable) SetValue(v FieldValueWritableAllowedValues)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


