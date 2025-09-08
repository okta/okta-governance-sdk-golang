# RequestTypeApprovalManager

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApproverType** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**ApproverFields** | Pointer to [**[]Field**](Field.md) | A list of fields with which to gather input.  The order of field object controls the order with which the fields are presented to users.  #### Known limitation  Unlike requester fields, all approver fields are *required* and may not be set as optional.  | [optional] [default to []]

## Methods

### NewRequestTypeApprovalManager

`func NewRequestTypeApprovalManager(approverType string, ) *RequestTypeApprovalManager`

NewRequestTypeApprovalManager instantiates a new RequestTypeApprovalManager object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTypeApprovalManagerWithDefaults

`func NewRequestTypeApprovalManagerWithDefaults() *RequestTypeApprovalManager`

NewRequestTypeApprovalManagerWithDefaults instantiates a new RequestTypeApprovalManager object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApproverType

`func (o *RequestTypeApprovalManager) GetApproverType() string`

GetApproverType returns the ApproverType field if non-nil, zero value otherwise.

### GetApproverTypeOk

`func (o *RequestTypeApprovalManager) GetApproverTypeOk() (*string, bool)`

GetApproverTypeOk returns a tuple with the ApproverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverType

`func (o *RequestTypeApprovalManager) SetApproverType(v string)`

SetApproverType sets ApproverType field to given value.


### GetDescription

`func (o *RequestTypeApprovalManager) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RequestTypeApprovalManager) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RequestTypeApprovalManager) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RequestTypeApprovalManager) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetApproverFields

`func (o *RequestTypeApprovalManager) GetApproverFields() []Field`

GetApproverFields returns the ApproverFields field if non-nil, zero value otherwise.

### GetApproverFieldsOk

`func (o *RequestTypeApprovalManager) GetApproverFieldsOk() (*[]Field, bool)`

GetApproverFieldsOk returns a tuple with the ApproverFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproverFields

`func (o *RequestTypeApprovalManager) SetApproverFields(v []Field)`

SetApproverFields sets ApproverFields field to given value.

### HasApproverFields

`func (o *RequestTypeApprovalManager) HasApproverFields() bool`

HasApproverFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


