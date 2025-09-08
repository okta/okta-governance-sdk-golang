# FieldValueDate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**FieldDateTimeType**](FieldDateTimeType.md) |  | 
**Value** | **NullableTime** | Value provided by a user in ISO 8601 string format | 
**Id** | **string** | A &#x60;read-only&#x60; field id.  Useful for specifying requesterFieldValues when adding a request.  The generated value is a UUID v4 from [RFC4122](https://www.ietf.org/rfc/rfc4122.txt).  | [readonly] 
**Prompt** | **string** | Text to prompt the user with | 
**Required** | **bool** | Whether a value to this field is required to advance the request | [default to true]

## Methods

### NewFieldValueDate

`func NewFieldValueDate(type_ FieldDateTimeType, value NullableTime, id string, prompt string, required bool, ) *FieldValueDate`

NewFieldValueDate instantiates a new FieldValueDate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldValueDateWithDefaults

`func NewFieldValueDateWithDefaults() *FieldValueDate`

NewFieldValueDateWithDefaults instantiates a new FieldValueDate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *FieldValueDate) GetType() FieldDateTimeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FieldValueDate) GetTypeOk() (*FieldDateTimeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FieldValueDate) SetType(v FieldDateTimeType)`

SetType sets Type field to given value.


### GetValue

`func (o *FieldValueDate) GetValue() time.Time`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FieldValueDate) GetValueOk() (*time.Time, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FieldValueDate) SetValue(v time.Time)`

SetValue sets Value field to given value.


### SetValueNil

`func (o *FieldValueDate) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *FieldValueDate) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil
### GetId

`func (o *FieldValueDate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FieldValueDate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FieldValueDate) SetId(v string)`

SetId sets Id field to given value.


### GetPrompt

`func (o *FieldValueDate) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *FieldValueDate) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *FieldValueDate) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetRequired

`func (o *FieldValueDate) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FieldValueDate) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FieldValueDate) SetRequired(v bool)`

SetRequired sets Required field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


