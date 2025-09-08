# BaseFieldWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prompt** | **string** | Text to prompt the user with | 
**Required** | Pointer to **bool** | Whether a value to this field is required to advance the request | [optional] [default to true]

## Methods

### NewBaseFieldWritable

`func NewBaseFieldWritable(prompt string, ) *BaseFieldWritable`

NewBaseFieldWritable instantiates a new BaseFieldWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseFieldWritableWithDefaults

`func NewBaseFieldWritableWithDefaults() *BaseFieldWritable`

NewBaseFieldWritableWithDefaults instantiates a new BaseFieldWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrompt

`func (o *BaseFieldWritable) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *BaseFieldWritable) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *BaseFieldWritable) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetRequired

`func (o *BaseFieldWritable) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *BaseFieldWritable) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *BaseFieldWritable) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *BaseFieldWritable) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


