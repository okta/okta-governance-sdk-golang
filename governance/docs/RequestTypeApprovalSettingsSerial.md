# RequestTypeApprovalSettingsSerial

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | When there are multiple approvals, an approval is not actionable until the previous approval has been approved. A denial will terminate the request and no subsequent approvals will be made actionable.  | 
**Approvals** | [**[]RequestTypeApproval**](RequestTypeApproval.md) | What approval(s) are required to grant access?  All specified approvals are considered required in order to fulfill an access request.  At least one approval must be specified for a request type. The maximum number of approvals is 5.  Approvals are serial.  When a request type has two approvals, and a user creates a request using that request type, then only the first approval will be immediately actionable by an approver.  After an approver has made a decision for that approval, then the second approval will be actionable by its approver.  The approval type of &#x60;RESOURCE_OWNER&#x60; is only supported for resource type of &#x60;GROUP&#x60;  | 

## Methods

### NewRequestTypeApprovalSettingsSerial

`func NewRequestTypeApprovalSettingsSerial(type_ string, approvals []RequestTypeApproval, ) *RequestTypeApprovalSettingsSerial`

NewRequestTypeApprovalSettingsSerial instantiates a new RequestTypeApprovalSettingsSerial object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestTypeApprovalSettingsSerialWithDefaults

`func NewRequestTypeApprovalSettingsSerialWithDefaults() *RequestTypeApprovalSettingsSerial`

NewRequestTypeApprovalSettingsSerialWithDefaults instantiates a new RequestTypeApprovalSettingsSerial object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *RequestTypeApprovalSettingsSerial) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestTypeApprovalSettingsSerial) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestTypeApprovalSettingsSerial) SetType(v string)`

SetType sets Type field to given value.


### GetApprovals

`func (o *RequestTypeApprovalSettingsSerial) GetApprovals() []RequestTypeApproval`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *RequestTypeApprovalSettingsSerial) GetApprovalsOk() (*[]RequestTypeApproval, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *RequestTypeApprovalSettingsSerial) SetApprovals(v []RequestTypeApproval)`

SetApprovals sets Approvals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


