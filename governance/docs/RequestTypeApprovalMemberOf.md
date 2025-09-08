# RequestTypeApprovalMemberOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApproverType** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ApproverMemberOf** | **[]string** | Okta groups the user persona must be a member of | 
**ApproverFields** | Pointer to [**[]Field**](Field.md) | A list of fields with which to gather input.  The order of field object controls the order with which the fields are presented to users.  #### Known limitation  Unlike requester fields, all approver fields are *required* and may not be set as optional.  | [optional] [default to []]

## Methods

### NewRequestTypeApprovalMemberOf

`func NewRequestTypeApprovalMemberOf(approverType string, approverMemberOf []string, ) *RequestTypeApprovalMemberOf`

NewRequestTypeApprovalMemberOf instantiates a new RequestTypeApprovalMemberOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTypeApprovalMemberOfWithDefaults

`func NewRequestTypeApprovalMemberOfWithDefaults() *RequestTypeApprovalMemberOf`

NewRequestTypeApprovalMemberOfWithDefaults instantiates a new RequestTypeApprovalMemberOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproverType

`func (o *RequestTypeApprovalMemberOf) GetApproverType() string`

GetApproverType returns the ApproverType field if non-nil, zero value otherwise.

### GetApproverTypeOk

`func (o *RequestTypeApprovalMemberOf) GetApproverTypeOk() (*string, bool)`

GetApproverTypeOk returns a tuple with the ApproverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverType

`func (o *RequestTypeApprovalMemberOf) SetApproverType(v string)`

SetApproverType sets ApproverType field to given value.


### GetDescription

`func (o *RequestTypeApprovalMemberOf) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestTypeApprovalMemberOf) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestTypeApprovalMemberOf) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequestTypeApprovalMemberOf) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetApproverMemberOf

`func (o *RequestTypeApprovalMemberOf) GetApproverMemberOf() []string`

GetApproverMemberOf returns the ApproverMemberOf field if non-nil, zero value otherwise.

### GetApproverMemberOfOk

`func (o *RequestTypeApprovalMemberOf) GetApproverMemberOfOk() (*[]string, bool)`

GetApproverMemberOfOk returns a tuple with the ApproverMemberOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverMemberOf

`func (o *RequestTypeApprovalMemberOf) SetApproverMemberOf(v []string)`

SetApproverMemberOf sets ApproverMemberOf field to given value.


### GetApproverFields

`func (o *RequestTypeApprovalMemberOf) GetApproverFields() []Field`

GetApproverFields returns the ApproverFields field if non-nil, zero value otherwise.

### GetApproverFieldsOk

`func (o *RequestTypeApprovalMemberOf) GetApproverFieldsOk() (*[]Field, bool)`

GetApproverFieldsOk returns a tuple with the ApproverFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverFields

`func (o *RequestTypeApprovalMemberOf) SetApproverFields(v []Field)`

SetApproverFields sets ApproverFields field to given value.

### HasApproverFields

`func (o *RequestTypeApprovalMemberOf) HasApproverFields() bool`

HasApproverFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


