# FieldValueText

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FieldTextType**](FieldTextType.md) |  | 
**Value** | **NullableString** | Value provided by a user. | 
**Id** | **string** | A &#x60;read-only&#x60; field id.  Useful for specifying requesterFieldValues when adding a request.  The generated value is a UUID v4 from [RFC4122](https://www.ietf.org/rfc/rfc4122.txt).  | [readonly] 
**Prompt** | **string** | Text to prompt the user with | 
**Required** | **bool** | Whether a value to this field is required to advance the request | [default to true]

## Methods

### NewFieldValueText

`func NewFieldValueText(type_ FieldTextType, value NullableString, id string, prompt string, required bool, ) *FieldValueText`

NewFieldValueText instantiates a new FieldValueText object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldValueTextWithDefaults

`func NewFieldValueTextWithDefaults() *FieldValueText`

NewFieldValueTextWithDefaults instantiates a new FieldValueText object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FieldValueText) GetType() FieldTextType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FieldValueText) GetTypeOk() (*FieldTextType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FieldValueText) SetType(v FieldTextType)`

SetType sets Type field to given value.


### GetValue

`func (o *FieldValueText) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FieldValueText) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FieldValueText) SetValue(v string)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *FieldValueText) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *FieldValueText) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetId

`func (o *FieldValueText) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FieldValueText) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FieldValueText) SetId(v string)`

SetId sets Id field to given value.


### GetPrompt

`func (o *FieldValueText) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *FieldValueText) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *FieldValueText) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetRequired

`func (o *FieldValueText) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FieldValueText) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FieldValueText) SetRequired(v bool)`

SetRequired sets Required field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


