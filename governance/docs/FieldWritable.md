# FieldWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FieldDateTimeType**](FieldDateTimeType.md) |  | 
**Options** | [**[]FieldOption**](FieldOption.md) | The options available for the select input. | 
**Prompt** | **string** | Text to prompt the user with | 
**Required** | Pointer to **bool** | Whether a value to this field is required to advance the request | [optional] [default to true]

## Methods

### NewFieldWritable

`func NewFieldWritable(type_ FieldDateTimeType, options []FieldOption, prompt string, ) *FieldWritable`

NewFieldWritable instantiates a new FieldWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldWritableWithDefaults

`func NewFieldWritableWithDefaults() *FieldWritable`

NewFieldWritableWithDefaults instantiates a new FieldWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FieldWritable) GetType() FieldDateTimeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FieldWritable) GetTypeOk() (*FieldDateTimeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FieldWritable) SetType(v FieldDateTimeType)`

SetType sets Type field to given value.


### GetOptions

`func (o *FieldWritable) GetOptions() []FieldOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FieldWritable) GetOptionsOk() (*[]FieldOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FieldWritable) SetOptions(v []FieldOption)`

SetOptions sets Options field to given value.


### GetPrompt

`func (o *FieldWritable) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *FieldWritable) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *FieldWritable) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetRequired

`func (o *FieldWritable) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FieldWritable) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FieldWritable) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FieldWritable) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


