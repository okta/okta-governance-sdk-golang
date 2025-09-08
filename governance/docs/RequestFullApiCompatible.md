# RequestFullApiCompatible

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approvals** | [**[]RequestApproval**](RequestApproval.md) |  | 
**Actions** | [**[]RequestAction**](RequestAction.md) | A list of actions. Currently only supports one action per request. | 
**Type** | **string** | This request is associated with a request type with no &#x60;CUSTOM&#x60; settings. | 
**Id** | **string** | Unique identifier for the object | 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 
**Links** | [**RequestLinks**](RequestLinks.md) |  | 
**RequestTypeId** | **string** | The Request Type enabling this Request. | 
**Subject** | **string** | The subject of the request | 
**RequesterUserIds** | **[]string** | A list of requester Okta user &#x60;id&#x60;s. | 
**RequestStatus** | [**RequestRequestStatus**](RequestRequestStatus.md) |  | 
**Resolved** | **NullableTime** | The date the request was resolved. The property may transition from having a value to null if the request is reopened. | 
**RequesterFieldValues** | [**[]FieldValue**](FieldValue.md) | Field values provided when adding the request.  If a request type has required requesterFields, they must be provided when the request is created.  Non-required fields may be omitted when creating the request.  | 

## Methods

### NewRequestFullApiCompatible

`func NewRequestFullApiCompatible(approvals []RequestApproval, actions []RequestAction, type_ string, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, links RequestLinks, requestTypeId string, subject string, requesterUserIds []string, requestStatus RequestRequestStatus, resolved NullableTime, requesterFieldValues []FieldValue, ) *RequestFullApiCompatible`

NewRequestFullApiCompatible instantiates a new RequestFullApiCompatible object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestFullApiCompatibleWithDefaults

`func NewRequestFullApiCompatibleWithDefaults() *RequestFullApiCompatible`

NewRequestFullApiCompatibleWithDefaults instantiates a new RequestFullApiCompatible object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovals

`func (o *RequestFullApiCompatible) GetApprovals() []RequestApproval`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *RequestFullApiCompatible) GetApprovalsOk() (*[]RequestApproval, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *RequestFullApiCompatible) SetApprovals(v []RequestApproval)`

SetApprovals sets Approvals field to given value.


### GetActions

`func (o *RequestFullApiCompatible) GetActions() []RequestAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *RequestFullApiCompatible) GetActionsOk() (*[]RequestAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *RequestFullApiCompatible) SetActions(v []RequestAction)`

SetActions sets Actions field to given value.


### SetActionsNil

`func (o *RequestFullApiCompatible) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *RequestFullApiCompatible) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetType

`func (o *RequestFullApiCompatible) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestFullApiCompatible) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestFullApiCompatible) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *RequestFullApiCompatible) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestFullApiCompatible) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestFullApiCompatible) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *RequestFullApiCompatible) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RequestFullApiCompatible) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RequestFullApiCompatible) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *RequestFullApiCompatible) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RequestFullApiCompatible) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RequestFullApiCompatible) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *RequestFullApiCompatible) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RequestFullApiCompatible) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RequestFullApiCompatible) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *RequestFullApiCompatible) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *RequestFullApiCompatible) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *RequestFullApiCompatible) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.


### GetLinks

`func (o *RequestFullApiCompatible) GetLinks() RequestLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RequestFullApiCompatible) GetLinksOk() (*RequestLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RequestFullApiCompatible) SetLinks(v RequestLinks)`

SetLinks sets Links field to given value.


### GetRequestTypeId

`func (o *RequestFullApiCompatible) GetRequestTypeId() string`

GetRequestTypeId returns the RequestTypeId field if non-nil, zero value otherwise.

### GetRequestTypeIdOk

`func (o *RequestFullApiCompatible) GetRequestTypeIdOk() (*string, bool)`

GetRequestTypeIdOk returns a tuple with the RequestTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTypeId

`func (o *RequestFullApiCompatible) SetRequestTypeId(v string)`

SetRequestTypeId sets RequestTypeId field to given value.


### GetSubject

`func (o *RequestFullApiCompatible) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RequestFullApiCompatible) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RequestFullApiCompatible) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetRequesterUserIds

`func (o *RequestFullApiCompatible) GetRequesterUserIds() []string`

GetRequesterUserIds returns the RequesterUserIds field if non-nil, zero value otherwise.

### GetRequesterUserIdsOk

`func (o *RequestFullApiCompatible) GetRequesterUserIdsOk() (*[]string, bool)`

GetRequesterUserIdsOk returns a tuple with the RequesterUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterUserIds

`func (o *RequestFullApiCompatible) SetRequesterUserIds(v []string)`

SetRequesterUserIds sets RequesterUserIds field to given value.


### GetRequestStatus

`func (o *RequestFullApiCompatible) GetRequestStatus() RequestRequestStatus`

GetRequestStatus returns the RequestStatus field if non-nil, zero value otherwise.

### GetRequestStatusOk

`func (o *RequestFullApiCompatible) GetRequestStatusOk() (*RequestRequestStatus, bool)`

GetRequestStatusOk returns a tuple with the RequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestStatus

`func (o *RequestFullApiCompatible) SetRequestStatus(v RequestRequestStatus)`

SetRequestStatus sets RequestStatus field to given value.


### GetResolved

`func (o *RequestFullApiCompatible) GetResolved() time.Time`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *RequestFullApiCompatible) GetResolvedOk() (*time.Time, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *RequestFullApiCompatible) SetResolved(v time.Time)`

SetResolved sets Resolved field to given value.


### SetResolvedNil

`func (o *RequestFullApiCompatible) SetResolvedNil(b bool)`

 SetResolvedNil sets the value for Resolved to be an explicit nil

### UnsetResolved
`func (o *RequestFullApiCompatible) UnsetResolved()`

UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
### GetRequesterFieldValues

`func (o *RequestFullApiCompatible) GetRequesterFieldValues() []FieldValue`

GetRequesterFieldValues returns the RequesterFieldValues field if non-nil, zero value otherwise.

### GetRequesterFieldValuesOk

`func (o *RequestFullApiCompatible) GetRequesterFieldValuesOk() (*[]FieldValue, bool)`

GetRequesterFieldValuesOk returns a tuple with the RequesterFieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterFieldValues

`func (o *RequestFullApiCompatible) SetRequesterFieldValues(v []FieldValue)`

SetRequesterFieldValues sets RequesterFieldValues field to given value.


### SetRequesterFieldValuesNil

`func (o *RequestFullApiCompatible) SetRequesterFieldValuesNil(b bool)`

 SetRequesterFieldValuesNil sets the value for RequesterFieldValues to be an explicit nil

### UnsetRequesterFieldValues
`func (o *RequestFullApiCompatible) UnsetRequesterFieldValues()`

UnsetRequesterFieldValues ensures that no value is present for RequesterFieldValues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


