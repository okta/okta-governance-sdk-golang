# BaseField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Prompt** | **string** | Text to prompt the user with | 
**Required** | Pointer to **bool** | Whether a value to this field is required to advance the request | [optional] [default to true]
**Id** | **string** | A &#x60;read-only&#x60; field id.  Useful for specifying requesterFieldValues when adding a request.  The generated value is a UUID v4 from [RFC4122](https://www.ietf.org/rfc/rfc4122.txt).  | [readonly] 

## Methods

### NewBaseField

`func NewBaseField(prompt string, id string, ) *BaseField`

NewBaseField instantiates a new BaseField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseFieldWithDefaults

`func NewBaseFieldWithDefaults() *BaseField`

NewBaseFieldWithDefaults instantiates a new BaseField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrompt

`func (o *BaseField) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *BaseField) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *BaseField) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetRequired

`func (o *BaseField) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *BaseField) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *BaseField) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *BaseField) HasRequired() bool`

HasRequired returns a boolean if a field has been set.

### GetId

`func (o *BaseField) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseField) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseField) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


