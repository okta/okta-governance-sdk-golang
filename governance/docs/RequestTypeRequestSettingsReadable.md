# RequestTypeRequestSettingsReadable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**RequesterMemberOf** | **[]string** | Okta groups the user persona must be a member of | 
**RequesterFields** | Pointer to [**[]Field**](Field.md) | A list of fields with which to gather input. The order of the field object controls the order with which the fields are presented to users. | [optional] [default to []]

## Methods

### NewRequestTypeRequestSettingsReadable

`func NewRequestTypeRequestSettingsReadable(type_ string, requesterMemberOf []string, ) *RequestTypeRequestSettingsReadable`

NewRequestTypeRequestSettingsReadable instantiates a new RequestTypeRequestSettingsReadable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTypeRequestSettingsReadableWithDefaults

`func NewRequestTypeRequestSettingsReadableWithDefaults() *RequestTypeRequestSettingsReadable`

NewRequestTypeRequestSettingsReadableWithDefaults instantiates a new RequestTypeRequestSettingsReadable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RequestTypeRequestSettingsReadable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestTypeRequestSettingsReadable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestTypeRequestSettingsReadable) SetType(v string)`

SetType sets Type field to given value.


### GetRequesterMemberOf

`func (o *RequestTypeRequestSettingsReadable) GetRequesterMemberOf() []string`

GetRequesterMemberOf returns the RequesterMemberOf field if non-nil, zero value otherwise.

### GetRequesterMemberOfOk

`func (o *RequestTypeRequestSettingsReadable) GetRequesterMemberOfOk() (*[]string, bool)`

GetRequesterMemberOfOk returns a tuple with the RequesterMemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterMemberOf

`func (o *RequestTypeRequestSettingsReadable) SetRequesterMemberOf(v []string)`

SetRequesterMemberOf sets RequesterMemberOf field to given value.


### GetRequesterFields

`func (o *RequestTypeRequestSettingsReadable) GetRequesterFields() []Field`

GetRequesterFields returns the RequesterFields field if non-nil, zero value otherwise.

### GetRequesterFieldsOk

`func (o *RequestTypeRequestSettingsReadable) GetRequesterFieldsOk() (*[]Field, bool)`

GetRequesterFieldsOk returns a tuple with the RequesterFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterFields

`func (o *RequestTypeRequestSettingsReadable) SetRequesterFields(v []Field)`

SetRequesterFields sets RequesterFields field to given value.

### HasRequesterFields

`func (o *RequestTypeRequestSettingsReadable) HasRequesterFields() bool`

HasRequesterFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


