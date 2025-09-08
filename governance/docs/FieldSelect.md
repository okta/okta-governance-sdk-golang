# FieldSelect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FieldSelectType**](FieldSelectType.md) |  | 
**Options** | [**[]FieldOption**](FieldOption.md) | The options available for the select input. | 
**Prompt** | **string** | Text to prompt the user with | 
**Required** | Pointer to **bool** | Whether a value to this field is required to advance the request | [optional] [default to true]
**Id** | **string** | A &#x60;read-only&#x60; field id.  Useful for specifying requesterFieldValues when adding a request.  The generated value is a UUID v4 from [RFC4122](https://www.ietf.org/rfc/rfc4122.txt).  | [readonly] 

## Methods

### NewFieldSelect

`func NewFieldSelect(type_ FieldSelectType, options []FieldOption, prompt string, id string, ) *FieldSelect`

NewFieldSelect instantiates a new FieldSelect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldSelectWithDefaults

`func NewFieldSelectWithDefaults() *FieldSelect`

NewFieldSelectWithDefaults instantiates a new FieldSelect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FieldSelect) GetType() FieldSelectType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FieldSelect) GetTypeOk() (*FieldSelectType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FieldSelect) SetType(v FieldSelectType)`

SetType sets Type field to given value.


### GetOptions

`func (o *FieldSelect) GetOptions() []FieldOption`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *FieldSelect) GetOptionsOk() (*[]FieldOption, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *FieldSelect) SetOptions(v []FieldOption)`

SetOptions sets Options field to given value.


### GetPrompt

`func (o *FieldSelect) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *FieldSelect) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *FieldSelect) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetRequired

`func (o *FieldSelect) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FieldSelect) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FieldSelect) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FieldSelect) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetId

`func (o *FieldSelect) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FieldSelect) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FieldSelect) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


