# RequestCreatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestTypeId** | **string** | The request type &#x60;id&#x60;.  | 
**Subject** | **string** | The subject of the request | 
**RequesterUserIds** | Pointer to **[]string** | A list of requester Okta user &#x60;id&#x60;s. | [optional] 
**RequesterFieldValues** | Pointer to [**[]FieldValueWritable**](FieldValueWritable.md) | Field values provided when adding the request.  If a request type has required requesterFields, they must be provided when the request is created.  Non-required fields may be omitted when creating the request.  | [optional] 

## Methods

### NewRequestCreatable

`func NewRequestCreatable(requestTypeId string, subject string, ) *RequestCreatable`

NewRequestCreatable instantiates a new RequestCreatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestCreatableWithDefaults

`func NewRequestCreatableWithDefaults() *RequestCreatable`

NewRequestCreatableWithDefaults instantiates a new RequestCreatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestTypeId

`func (o *RequestCreatable) GetRequestTypeId() string`

GetRequestTypeId returns the RequestTypeId field if non-nil, zero value otherwise.

### GetRequestTypeIdOk

`func (o *RequestCreatable) GetRequestTypeIdOk() (*string, bool)`

GetRequestTypeIdOk returns a tuple with the RequestTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTypeId

`func (o *RequestCreatable) SetRequestTypeId(v string)`

SetRequestTypeId sets RequestTypeId field to given value.


### GetSubject

`func (o *RequestCreatable) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RequestCreatable) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RequestCreatable) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetRequesterUserIds

`func (o *RequestCreatable) GetRequesterUserIds() []string`

GetRequesterUserIds returns the RequesterUserIds field if non-nil, zero value otherwise.

### GetRequesterUserIdsOk

`func (o *RequestCreatable) GetRequesterUserIdsOk() (*[]string, bool)`

GetRequesterUserIdsOk returns a tuple with the RequesterUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterUserIds

`func (o *RequestCreatable) SetRequesterUserIds(v []string)`

SetRequesterUserIds sets RequesterUserIds field to given value.

### HasRequesterUserIds

`func (o *RequestCreatable) HasRequesterUserIds() bool`

HasRequesterUserIds returns a boolean if a field has been set.

### GetRequesterFieldValues

`func (o *RequestCreatable) GetRequesterFieldValues() []FieldValueWritable`

GetRequesterFieldValues returns the RequesterFieldValues field if non-nil, zero value otherwise.

### GetRequesterFieldValuesOk

`func (o *RequestCreatable) GetRequesterFieldValuesOk() (*[]FieldValueWritable, bool)`

GetRequesterFieldValuesOk returns a tuple with the RequesterFieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterFieldValues

`func (o *RequestCreatable) SetRequesterFieldValues(v []FieldValueWritable)`

SetRequesterFieldValues sets RequesterFieldValues field to given value.

### HasRequesterFieldValues

`func (o *RequestCreatable) HasRequesterFieldValues() bool`

HasRequesterFieldValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


