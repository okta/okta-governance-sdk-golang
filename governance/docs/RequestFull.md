# RequestFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approvals** | [**[]RequestApproval**](RequestApproval.md) |  | 
**Actions** | [**[]RequestAction**](RequestAction.md) | A list of actions. Currently only supports one action per request. | 
**PermalinkId** | **int32** | The immutable, persistent identifier that always resolves to the request | 
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

### NewRequestFull

`func NewRequestFull(approvals []RequestApproval, actions []RequestAction, permalinkId int32, type_ string, id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, links RequestLinks, requestTypeId string, subject string, requesterUserIds []string, requestStatus RequestRequestStatus, resolved NullableTime, requesterFieldValues []FieldValue, ) *RequestFull`

NewRequestFull instantiates a new RequestFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestFullWithDefaults

`func NewRequestFullWithDefaults() *RequestFull`

NewRequestFullWithDefaults instantiates a new RequestFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovals

`func (o *RequestFull) GetApprovals() []RequestApproval`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *RequestFull) GetApprovalsOk() (*[]RequestApproval, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *RequestFull) SetApprovals(v []RequestApproval)`

SetApprovals sets Approvals field to given value.


### GetActions

`func (o *RequestFull) GetActions() []RequestAction`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *RequestFull) GetActionsOk() (*[]RequestAction, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *RequestFull) SetActions(v []RequestAction)`

SetActions sets Actions field to given value.


### SetActionsNil

`func (o *RequestFull) SetActionsNil(b bool)`

 SetActionsNil sets the value for Actions to be an explicit nil

### UnsetActions
`func (o *RequestFull) UnsetActions()`

UnsetActions ensures that no value is present for Actions, not even an explicit nil
### GetPermalinkId

`func (o *RequestFull) GetPermalinkId() int32`

GetPermalinkId returns the PermalinkId field if non-nil, zero value otherwise.

### GetPermalinkIdOk

`func (o *RequestFull) GetPermalinkIdOk() (*int32, bool)`

GetPermalinkIdOk returns a tuple with the PermalinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermalinkId

`func (o *RequestFull) SetPermalinkId(v int32)`

SetPermalinkId sets PermalinkId field to given value.


### GetType

`func (o *RequestFull) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RequestFull) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RequestFull) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *RequestFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestFull) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *RequestFull) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RequestFull) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RequestFull) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *RequestFull) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RequestFull) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RequestFull) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *RequestFull) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RequestFull) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RequestFull) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *RequestFull) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *RequestFull) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *RequestFull) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.


### GetLinks

`func (o *RequestFull) GetLinks() RequestLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RequestFull) GetLinksOk() (*RequestLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RequestFull) SetLinks(v RequestLinks)`

SetLinks sets Links field to given value.


### GetRequestTypeId

`func (o *RequestFull) GetRequestTypeId() string`

GetRequestTypeId returns the RequestTypeId field if non-nil, zero value otherwise.

### GetRequestTypeIdOk

`func (o *RequestFull) GetRequestTypeIdOk() (*string, bool)`

GetRequestTypeIdOk returns a tuple with the RequestTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTypeId

`func (o *RequestFull) SetRequestTypeId(v string)`

SetRequestTypeId sets RequestTypeId field to given value.


### GetSubject

`func (o *RequestFull) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *RequestFull) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *RequestFull) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetRequesterUserIds

`func (o *RequestFull) GetRequesterUserIds() []string`

GetRequesterUserIds returns the RequesterUserIds field if non-nil, zero value otherwise.

### GetRequesterUserIdsOk

`func (o *RequestFull) GetRequesterUserIdsOk() (*[]string, bool)`

GetRequesterUserIdsOk returns a tuple with the RequesterUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterUserIds

`func (o *RequestFull) SetRequesterUserIds(v []string)`

SetRequesterUserIds sets RequesterUserIds field to given value.


### GetRequestStatus

`func (o *RequestFull) GetRequestStatus() RequestRequestStatus`

GetRequestStatus returns the RequestStatus field if non-nil, zero value otherwise.

### GetRequestStatusOk

`func (o *RequestFull) GetRequestStatusOk() (*RequestRequestStatus, bool)`

GetRequestStatusOk returns a tuple with the RequestStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestStatus

`func (o *RequestFull) SetRequestStatus(v RequestRequestStatus)`

SetRequestStatus sets RequestStatus field to given value.


### GetResolved

`func (o *RequestFull) GetResolved() time.Time`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *RequestFull) GetResolvedOk() (*time.Time, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *RequestFull) SetResolved(v time.Time)`

SetResolved sets Resolved field to given value.


### SetResolvedNil

`func (o *RequestFull) SetResolvedNil(b bool)`

 SetResolvedNil sets the value for Resolved to be an explicit nil

### UnsetResolved
`func (o *RequestFull) UnsetResolved()`

UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
### GetRequesterFieldValues

`func (o *RequestFull) GetRequesterFieldValues() []FieldValue`

GetRequesterFieldValues returns the RequesterFieldValues field if non-nil, zero value otherwise.

### GetRequesterFieldValuesOk

`func (o *RequestFull) GetRequesterFieldValuesOk() (*[]FieldValue, bool)`

GetRequesterFieldValuesOk returns a tuple with the RequesterFieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterFieldValues

`func (o *RequestFull) SetRequesterFieldValues(v []FieldValue)`

SetRequesterFieldValues sets RequesterFieldValues field to given value.


### SetRequesterFieldValuesNil

`func (o *RequestFull) SetRequesterFieldValuesNil(b bool)`

 SetRequesterFieldValuesNil sets the value for RequesterFieldValues to be an explicit nil

### UnsetRequesterFieldValues
`func (o *RequestFull) UnsetRequesterFieldValues()`

UnsetRequesterFieldValues ensures that no value is present for RequesterFieldValues, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


