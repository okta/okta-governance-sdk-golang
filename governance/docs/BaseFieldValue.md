# BaseFieldValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A &#x60;read-only&#x60; field id.  Useful for specifying requesterFieldValues when adding a request.  The generated value is a UUID v4 from [RFC4122](https://www.ietf.org/rfc/rfc4122.txt).  | [readonly] 
**Prompt** | **string** | Text to prompt the user with | 
**Required** | **bool** | Whether a value to this field is required to advance the request | [default to true]

## Methods

### NewBaseFieldValue

`func NewBaseFieldValue(id string, prompt string, required bool, ) *BaseFieldValue`

NewBaseFieldValue instantiates a new BaseFieldValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseFieldValueWithDefaults

`func NewBaseFieldValueWithDefaults() *BaseFieldValue`

NewBaseFieldValueWithDefaults instantiates a new BaseFieldValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BaseFieldValue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseFieldValue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseFieldValue) SetId(v string)`

SetId sets Id field to given value.


### GetPrompt

`func (o *BaseFieldValue) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *BaseFieldValue) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *BaseFieldValue) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetRequired

`func (o *BaseFieldValue) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *BaseFieldValue) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *BaseFieldValue) SetRequired(v bool)`

SetRequired sets Required field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


