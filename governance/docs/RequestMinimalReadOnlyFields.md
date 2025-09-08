# RequestMinimalReadOnlyFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestStatus** | Pointer to [**RequestRequestStatus**](RequestRequestStatus.md) |  | [optional] 
**RequestTypeId** | Pointer to **string** | The Request Type enabling this Request. | [optional] 
**Resolved** | Pointer to **NullableTime** | The date the request was resolved. The property may transition from having a value to null if the request is reopened. | [optional] 
**RequesterFieldValues** | Pointer to [**[]FieldValue**](FieldValue.md) | Field values provided when adding the request.  If a request type has required requesterFields, they must be provided when the request is created.  Non-required fields may be omitted when creating the request.  | [optional] 
**Links** | Pointer to [**RequestLinks**](RequestLinks.md) |  | [optional] 

## Methods

### NewRequestMinimalReadOnlyFields

`func NewRequestMinimalReadOnlyFields() *RequestMinimalReadOnlyFields`

NewRequestMinimalReadOnlyFields instantiates a new RequestMinimalReadOnlyFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestMinimalReadOnlyFieldsWithDefaults

`func NewRequestMinimalReadOnlyFieldsWithDefaults() *RequestMinimalReadOnlyFields`

NewRequestMinimalReadOnlyFieldsWithDefaults instantiates a new RequestMinimalReadOnlyFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestStatus

`func (o *RequestMinimalReadOnlyFields) GetRequestStatus() RequestRequestStatus`

GetRequestStatus returns the RequestStatus field if non-nil, zero value otherwise.

### GetRequestStatusOk

`func (o *RequestMinimalReadOnlyFields) GetRequestStatusOk() (*RequestRequestStatus, bool)`

GetRequestStatusOk returns a tuple with the RequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestStatus

`func (o *RequestMinimalReadOnlyFields) SetRequestStatus(v RequestRequestStatus)`

SetRequestStatus sets RequestStatus field to given value.

### HasRequestStatus

`func (o *RequestMinimalReadOnlyFields) HasRequestStatus() bool`

HasRequestStatus returns a boolean if a field has been set.

### GetRequestTypeId

`func (o *RequestMinimalReadOnlyFields) GetRequestTypeId() string`

GetRequestTypeId returns the RequestTypeId field if non-nil, zero value otherwise.

### GetRequestTypeIdOk

`func (o *RequestMinimalReadOnlyFields) GetRequestTypeIdOk() (*string, bool)`

GetRequestTypeIdOk returns a tuple with the RequestTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTypeId

`func (o *RequestMinimalReadOnlyFields) SetRequestTypeId(v string)`

SetRequestTypeId sets RequestTypeId field to given value.

### HasRequestTypeId

`func (o *RequestMinimalReadOnlyFields) HasRequestTypeId() bool`

HasRequestTypeId returns a boolean if a field has been set.

### GetResolved

`func (o *RequestMinimalReadOnlyFields) GetResolved() time.Time`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *RequestMinimalReadOnlyFields) GetResolvedOk() (*time.Time, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *RequestMinimalReadOnlyFields) SetResolved(v time.Time)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *RequestMinimalReadOnlyFields) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### SetResolvedNil

`func (o *RequestMinimalReadOnlyFields) SetResolvedNil(b bool)`

 SetResolvedNil sets the value for Resolved to be an explicit nil

### UnsetResolved
`func (o *RequestMinimalReadOnlyFields) UnsetResolved()`

UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
### GetRequesterFieldValues

`func (o *RequestMinimalReadOnlyFields) GetRequesterFieldValues() []FieldValue`

GetRequesterFieldValues returns the RequesterFieldValues field if non-nil, zero value otherwise.

### GetRequesterFieldValuesOk

`func (o *RequestMinimalReadOnlyFields) GetRequesterFieldValuesOk() (*[]FieldValue, bool)`

GetRequesterFieldValuesOk returns a tuple with the RequesterFieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterFieldValues

`func (o *RequestMinimalReadOnlyFields) SetRequesterFieldValues(v []FieldValue)`

SetRequesterFieldValues sets RequesterFieldValues field to given value.

### HasRequesterFieldValues

`func (o *RequestMinimalReadOnlyFields) HasRequesterFieldValues() bool`

HasRequesterFieldValues returns a boolean if a field has been set.

### SetRequesterFieldValuesNil

`func (o *RequestMinimalReadOnlyFields) SetRequesterFieldValuesNil(b bool)`

 SetRequesterFieldValuesNil sets the value for RequesterFieldValues to be an explicit nil

### UnsetRequesterFieldValues
`func (o *RequestMinimalReadOnlyFields) UnsetRequesterFieldValues()`

UnsetRequesterFieldValues ensures that no value is present for RequesterFieldValues, not even an explicit nil
### GetLinks

`func (o *RequestMinimalReadOnlyFields) GetLinks() RequestLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RequestMinimalReadOnlyFields) GetLinksOk() (*RequestLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RequestMinimalReadOnlyFields) SetLinks(v RequestLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RequestMinimalReadOnlyFields) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


