# RequestTypeApproval

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApproverType** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ApproverMemberOf** | **[]string** | Okta groups the user persona must be a member of | 
**ApproverFields** | Pointer to [**[]Field**](Field.md) | A list of fields with which to gather input.  The order of field object controls the order with which the fields are presented to users.  #### Known limitation  Unlike requester fields, all approver fields are *required* and may not be set as optional.  | [optional] [default to []]
**ApproverUserId** | **string** | The Okta user &#x60;id&#x60; | 

## Methods

### NewRequestTypeApproval

`func NewRequestTypeApproval(approverType string, approverMemberOf []string, approverUserId string, ) *RequestTypeApproval`

NewRequestTypeApproval instantiates a new RequestTypeApproval object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTypeApprovalWithDefaults

`func NewRequestTypeApprovalWithDefaults() *RequestTypeApproval`

NewRequestTypeApprovalWithDefaults instantiates a new RequestTypeApproval object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproverType

`func (o *RequestTypeApproval) GetApproverType() string`

GetApproverType returns the ApproverType field if non-nil, zero value otherwise.

### GetApproverTypeOk

`func (o *RequestTypeApproval) GetApproverTypeOk() (*string, bool)`

GetApproverTypeOk returns a tuple with the ApproverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverType

`func (o *RequestTypeApproval) SetApproverType(v string)`

SetApproverType sets ApproverType field to given value.


### GetDescription

`func (o *RequestTypeApproval) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestTypeApproval) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestTypeApproval) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequestTypeApproval) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetApproverMemberOf

`func (o *RequestTypeApproval) GetApproverMemberOf() []string`

GetApproverMemberOf returns the ApproverMemberOf field if non-nil, zero value otherwise.

### GetApproverMemberOfOk

`func (o *RequestTypeApproval) GetApproverMemberOfOk() (*[]string, bool)`

GetApproverMemberOfOk returns a tuple with the ApproverMemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverMemberOf

`func (o *RequestTypeApproval) SetApproverMemberOf(v []string)`

SetApproverMemberOf sets ApproverMemberOf field to given value.


### GetApproverFields

`func (o *RequestTypeApproval) GetApproverFields() []Field`

GetApproverFields returns the ApproverFields field if non-nil, zero value otherwise.

### GetApproverFieldsOk

`func (o *RequestTypeApproval) GetApproverFieldsOk() (*[]Field, bool)`

GetApproverFieldsOk returns a tuple with the ApproverFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverFields

`func (o *RequestTypeApproval) SetApproverFields(v []Field)`

SetApproverFields sets ApproverFields field to given value.

### HasApproverFields

`func (o *RequestTypeApproval) HasApproverFields() bool`

HasApproverFields returns a boolean if a field has been set.

### GetApproverUserId

`func (o *RequestTypeApproval) GetApproverUserId() string`

GetApproverUserId returns the ApproverUserId field if non-nil, zero value otherwise.

### GetApproverUserIdOk

`func (o *RequestTypeApproval) GetApproverUserIdOk() (*string, bool)`

GetApproverUserIdOk returns a tuple with the ApproverUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverUserId

`func (o *RequestTypeApproval) SetApproverUserId(v string)`

SetApproverUserId sets ApproverUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


