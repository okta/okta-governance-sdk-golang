# FieldTextWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FieldTextType**](FieldTextType.md) |  | 
**Prompt** | **string** | Text to prompt the user with | 
**Required** | Pointer to **bool** | Whether a value to this field is required to advance the request | [optional] [default to true]

## Methods

### NewFieldTextWritable

`func NewFieldTextWritable(type_ FieldTextType, prompt string, ) *FieldTextWritable`

NewFieldTextWritable instantiates a new FieldTextWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldTextWritableWithDefaults

`func NewFieldTextWritableWithDefaults() *FieldTextWritable`

NewFieldTextWritableWithDefaults instantiates a new FieldTextWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FieldTextWritable) GetType() FieldTextType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FieldTextWritable) GetTypeOk() (*FieldTextType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FieldTextWritable) SetType(v FieldTextType)`

SetType sets Type field to given value.


### GetPrompt

`func (o *FieldTextWritable) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *FieldTextWritable) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *FieldTextWritable) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetRequired

`func (o *FieldTextWritable) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FieldTextWritable) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FieldTextWritable) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FieldTextWritable) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


