# RequestTypeRequesterMemberOfWritable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**RequesterMemberOf** | **[]string** | Okta groups the user persona must be a member of | 
**RequesterFields** | Pointer to [**[]FieldWritable**](FieldWritable.md) | A list of fields with which to gather input. The order of the field object controls the order with which the fields are presented to users. | [optional] [default to []]

## Methods

### NewRequestTypeRequesterMemberOfWritable

`func NewRequestTypeRequesterMemberOfWritable(type_ string, requesterMemberOf []string, ) *RequestTypeRequesterMemberOfWritable`

NewRequestTypeRequesterMemberOfWritable instantiates a new RequestTypeRequesterMemberOfWritable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTypeRequesterMemberOfWritableWithDefaults

`func NewRequestTypeRequesterMemberOfWritableWithDefaults() *RequestTypeRequesterMemberOfWritable`

NewRequestTypeRequesterMemberOfWritableWithDefaults instantiates a new RequestTypeRequesterMemberOfWritable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RequestTypeRequesterMemberOfWritable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestTypeRequesterMemberOfWritable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestTypeRequesterMemberOfWritable) SetType(v string)`

SetType sets Type field to given value.


### GetRequesterMemberOf

`func (o *RequestTypeRequesterMemberOfWritable) GetRequesterMemberOf() []string`

GetRequesterMemberOf returns the RequesterMemberOf field if non-nil, zero value otherwise.

### GetRequesterMemberOfOk

`func (o *RequestTypeRequesterMemberOfWritable) GetRequesterMemberOfOk() (*[]string, bool)`

GetRequesterMemberOfOk returns a tuple with the RequesterMemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterMemberOf

`func (o *RequestTypeRequesterMemberOfWritable) SetRequesterMemberOf(v []string)`

SetRequesterMemberOf sets RequesterMemberOf field to given value.


### GetRequesterFields

`func (o *RequestTypeRequesterMemberOfWritable) GetRequesterFields() []FieldWritable`

GetRequesterFields returns the RequesterFields field if non-nil, zero value otherwise.

### GetRequesterFieldsOk

`func (o *RequestTypeRequesterMemberOfWritable) GetRequesterFieldsOk() (*[]FieldWritable, bool)`

GetRequesterFieldsOk returns a tuple with the RequesterFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterFields

`func (o *RequestTypeRequesterMemberOfWritable) SetRequesterFields(v []FieldWritable)`

SetRequesterFields sets RequesterFields field to given value.

### HasRequesterFields

`func (o *RequestTypeRequesterMemberOfWritable) HasRequesterFields() bool`

HasRequesterFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


