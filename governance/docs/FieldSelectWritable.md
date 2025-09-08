# FieldSelectWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FieldSelectType**](FieldSelectType.md) |  | 
**Options** | [**[]FieldOption**](FieldOption.md) | The options available for the select input. | 
**Prompt** | **string** | Text to prompt the user with | 
**Required** | Pointer to **bool** | Whether a value to this field is required to advance the request | [optional] [default to true]

## Methods

### NewFieldSelectWritable

`func NewFieldSelectWritable(type_ FieldSelectType, options []FieldOption, prompt string, ) *FieldSelectWritable`

NewFieldSelectWritable instantiates a new FieldSelectWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldSelectWritableWithDefaults

`func NewFieldSelectWritableWithDefaults() *FieldSelectWritable`

NewFieldSelectWritableWithDefaults instantiates a new FieldSelectWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FieldSelectWritable) GetType() FieldSelectType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FieldSelectWritable) GetTypeOk() (*FieldSelectType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FieldSelectWritable) SetType(v FieldSelectType)`

SetType sets Type field to given value.


### GetOptions

`func (o *FieldSelectWritable) GetOptions() []FieldOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FieldSelectWritable) GetOptionsOk() (*[]FieldOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FieldSelectWritable) SetOptions(v []FieldOption)`

SetOptions sets Options field to given value.


### GetPrompt

`func (o *FieldSelectWritable) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *FieldSelectWritable) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *FieldSelectWritable) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetRequired

`func (o *FieldSelectWritable) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FieldSelectWritable) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FieldSelectWritable) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FieldSelectWritable) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


