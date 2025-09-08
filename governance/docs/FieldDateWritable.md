# FieldDateWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FieldDateTimeType**](FieldDateTimeType.md) |  | 
**Prompt** | **string** | Text to prompt the user with | 
**Required** | Pointer to **bool** | Whether a value to this field is required to advance the request | [optional] [default to true]

## Methods

### NewFieldDateWritable

`func NewFieldDateWritable(type_ FieldDateTimeType, prompt string, ) *FieldDateWritable`

NewFieldDateWritable instantiates a new FieldDateWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldDateWritableWithDefaults

`func NewFieldDateWritableWithDefaults() *FieldDateWritable`

NewFieldDateWritableWithDefaults instantiates a new FieldDateWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FieldDateWritable) GetType() FieldDateTimeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FieldDateWritable) GetTypeOk() (*FieldDateTimeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FieldDateWritable) SetType(v FieldDateTimeType)`

SetType sets Type field to given value.


### GetPrompt

`func (o *FieldDateWritable) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *FieldDateWritable) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *FieldDateWritable) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetRequired

`func (o *FieldDateWritable) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FieldDateWritable) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FieldDateWritable) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FieldDateWritable) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


