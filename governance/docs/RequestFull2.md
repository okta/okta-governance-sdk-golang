# RequestFull2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the object | 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 
**Links** | [**RequestLinks2**](RequestLinks2.md) |  | 
**Status** | [**RequestStatus**](RequestStatus.md) |  | 
**Resolved** | Pointer to **NullableTime** | The date the request was resolved. The property may transition from having a value to null if the request is reopened. | [optional] 
**GrantStatus** | Pointer to [**RequestGrantStatus**](RequestGrantStatus.md) |  | [optional] 
**Granted** | Pointer to **NullableTime** | The date the approved access was granted. Only set if request.status is APPROVED. | [optional] 
**RevocationStatus** | Pointer to [**RequestRevocationStatus**](RequestRevocationStatus.md) |  | [optional] 
**Revoked** | Pointer to **NullableTime** | The date the granted access was revoked. Only set if request.grantStatus is GRANTED and request.revocationStatus is REVOKED. | [optional] 
**RequestedBy** | [**ClientCredentialPrincipal**](ClientCredentialPrincipal.md) |  | 
**RequestedFor** | [**TargetPrincipal**](TargetPrincipal.md) |  | 
**Requested** | [**Requested**](Requested.md) |  | 
**AccessDuration** | Pointer to **NullableString** | How long the requester retains access after their request is approved and fulfilled.  Specified in [ISO 8601 duration format](https://en.wikipedia.org/wiki/ISO_8601#Durations).  #### Known limitation  Only single time unit ISO 8601 duration formats (D, H, M) are supported, for units (days, hours, minutes).  ##### Supported  | Unit       | Example | | ---------- | ------- | | D, days    | P40D    | | H, hours   | PT65H   | | M, minutes | PT90M   |  &gt; **Note:** Mixes of units, as well as month/year/week designations, are not supported. For example, &#x60;P40DT65H&#x60;, &#x60;P40M&#x60;, &#x60;P1W&#x60; and &#x60;P1Y&#x60; are not supported. | [optional] 
**RevocationScheduled** | Pointer to **NullableTime** | The date the granted access is scheduled for recovation. Only set if request.accessDuration exists, and request.grantStatus is GRANTED. | [optional] 
**RequesterFieldValues** | Pointer to [**[]RequestFieldValue**](RequestFieldValue.md) | The requester input fields required by the approval system.  **Note:** The fields required are determined by the approval system.  For the Okta approval system, the required fields are defined in the approval sequence. Ensure that the requester input fields match up with this definition to avoid request approval flow failure.  For external approval systems, the requester input fields are for recording purposes only and do not affect the approval process.  | [optional] 
**RequestApproval** | Pointer to [**RequestApproval2**](RequestApproval2.md) |  | [optional] 
**RiskAssessment** | Pointer to [**RiskAssessment**](RiskAssessment.md) |  | [optional] 

## Methods

### NewRequestFull2

`func NewRequestFull2(id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, links RequestLinks2, status RequestStatus, requestedBy ClientCredentialPrincipal, requestedFor TargetPrincipal, requested Requested, ) *RequestFull2`

NewRequestFull2 instantiates a new RequestFull2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRequestFull2WithDefaults

`func NewRequestFull2WithDefaults() *RequestFull2`

NewRequestFull2WithDefaults instantiates a new RequestFull2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RequestFull2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RequestFull2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RequestFull2) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *RequestFull2) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RequestFull2) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RequestFull2) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *RequestFull2) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RequestFull2) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RequestFull2) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *RequestFull2) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RequestFull2) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RequestFull2) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *RequestFull2) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *RequestFull2) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *RequestFull2) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.


### GetLinks

`func (o *RequestFull2) GetLinks() RequestLinks2`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RequestFull2) GetLinksOk() (*RequestLinks2, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RequestFull2) SetLinks(v RequestLinks2)`

SetLinks sets Links field to given value.


### GetStatus

`func (o *RequestFull2) GetStatus() RequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RequestFull2) GetStatusOk() (*RequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RequestFull2) SetStatus(v RequestStatus)`

SetStatus sets Status field to given value.


### GetResolved

`func (o *RequestFull2) GetResolved() time.Time`

GetResolved returns the Resolved field if non-nil, zero value otherwise.

### GetResolvedOk

`func (o *RequestFull2) GetResolvedOk() (*time.Time, bool)`

GetResolvedOk returns a tuple with the Resolved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolved

`func (o *RequestFull2) SetResolved(v time.Time)`

SetResolved sets Resolved field to given value.

### HasResolved

`func (o *RequestFull2) HasResolved() bool`

HasResolved returns a boolean if a field has been set.

### SetResolvedNil

`func (o *RequestFull2) SetResolvedNil(b bool)`

 SetResolvedNil sets the value for Resolved to be an explicit nil

### UnsetResolved
`func (o *RequestFull2) UnsetResolved()`

UnsetResolved ensures that no value is present for Resolved, not even an explicit nil
### GetGrantStatus

`func (o *RequestFull2) GetGrantStatus() RequestGrantStatus`

GetGrantStatus returns the GrantStatus field if non-nil, zero value otherwise.

### GetGrantStatusOk

`func (o *RequestFull2) GetGrantStatusOk() (*RequestGrantStatus, bool)`

GetGrantStatusOk returns a tuple with the GrantStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantStatus

`func (o *RequestFull2) SetGrantStatus(v RequestGrantStatus)`

SetGrantStatus sets GrantStatus field to given value.

### HasGrantStatus

`func (o *RequestFull2) HasGrantStatus() bool`

HasGrantStatus returns a boolean if a field has been set.

### GetGranted

`func (o *RequestFull2) GetGranted() time.Time`

GetGranted returns the Granted field if non-nil, zero value otherwise.

### GetGrantedOk

`func (o *RequestFull2) GetGrantedOk() (*time.Time, bool)`

GetGrantedOk returns a tuple with the Granted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranted

`func (o *RequestFull2) SetGranted(v time.Time)`

SetGranted sets Granted field to given value.

### HasGranted

`func (o *RequestFull2) HasGranted() bool`

HasGranted returns a boolean if a field has been set.

### SetGrantedNil

`func (o *RequestFull2) SetGrantedNil(b bool)`

 SetGrantedNil sets the value for Granted to be an explicit nil

### UnsetGranted
`func (o *RequestFull2) UnsetGranted()`

UnsetGranted ensures that no value is present for Granted, not even an explicit nil
### GetRevocationStatus

`func (o *RequestFull2) GetRevocationStatus() RequestRevocationStatus`

GetRevocationStatus returns the RevocationStatus field if non-nil, zero value otherwise.

### GetRevocationStatusOk

`func (o *RequestFull2) GetRevocationStatusOk() (*RequestRevocationStatus, bool)`

GetRevocationStatusOk returns a tuple with the RevocationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationStatus

`func (o *RequestFull2) SetRevocationStatus(v RequestRevocationStatus)`

SetRevocationStatus sets RevocationStatus field to given value.

### HasRevocationStatus

`func (o *RequestFull2) HasRevocationStatus() bool`

HasRevocationStatus returns a boolean if a field has been set.

### GetRevoked

`func (o *RequestFull2) GetRevoked() time.Time`

GetRevoked returns the Revoked field if non-nil, zero value otherwise.

### GetRevokedOk

`func (o *RequestFull2) GetRevokedOk() (*time.Time, bool)`

GetRevokedOk returns a tuple with the Revoked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevoked

`func (o *RequestFull2) SetRevoked(v time.Time)`

SetRevoked sets Revoked field to given value.

### HasRevoked

`func (o *RequestFull2) HasRevoked() bool`

HasRevoked returns a boolean if a field has been set.

### SetRevokedNil

`func (o *RequestFull2) SetRevokedNil(b bool)`

 SetRevokedNil sets the value for Revoked to be an explicit nil

### UnsetRevoked
`func (o *RequestFull2) UnsetRevoked()`

UnsetRevoked ensures that no value is present for Revoked, not even an explicit nil
### GetRequestedBy

`func (o *RequestFull2) GetRequestedBy() ClientCredentialPrincipal`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *RequestFull2) GetRequestedByOk() (*ClientCredentialPrincipal, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *RequestFull2) SetRequestedBy(v ClientCredentialPrincipal)`

SetRequestedBy sets RequestedBy field to given value.


### GetRequestedFor

`func (o *RequestFull2) GetRequestedFor() TargetPrincipal`

GetRequestedFor returns the RequestedFor field if non-nil, zero value otherwise.

### GetRequestedForOk

`func (o *RequestFull2) GetRequestedForOk() (*TargetPrincipal, bool)`

GetRequestedForOk returns a tuple with the RequestedFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedFor

`func (o *RequestFull2) SetRequestedFor(v TargetPrincipal)`

SetRequestedFor sets RequestedFor field to given value.


### GetRequested

`func (o *RequestFull2) GetRequested() Requested`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *RequestFull2) GetRequestedOk() (*Requested, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *RequestFull2) SetRequested(v Requested)`

SetRequested sets Requested field to given value.


### GetAccessDuration

`func (o *RequestFull2) GetAccessDuration() string`

GetAccessDuration returns the AccessDuration field if non-nil, zero value otherwise.

### GetAccessDurationOk

`func (o *RequestFull2) GetAccessDurationOk() (*string, bool)`

GetAccessDurationOk returns a tuple with the AccessDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessDuration

`func (o *RequestFull2) SetAccessDuration(v string)`

SetAccessDuration sets AccessDuration field to given value.

### HasAccessDuration

`func (o *RequestFull2) HasAccessDuration() bool`

HasAccessDuration returns a boolean if a field has been set.

### SetAccessDurationNil

`func (o *RequestFull2) SetAccessDurationNil(b bool)`

 SetAccessDurationNil sets the value for AccessDuration to be an explicit nil

### UnsetAccessDuration
`func (o *RequestFull2) UnsetAccessDuration()`

UnsetAccessDuration ensures that no value is present for AccessDuration, not even an explicit nil
### GetRevocationScheduled

`func (o *RequestFull2) GetRevocationScheduled() time.Time`

GetRevocationScheduled returns the RevocationScheduled field if non-nil, zero value otherwise.

### GetRevocationScheduledOk

`func (o *RequestFull2) GetRevocationScheduledOk() (*time.Time, bool)`

GetRevocationScheduledOk returns a tuple with the RevocationScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevocationScheduled

`func (o *RequestFull2) SetRevocationScheduled(v time.Time)`

SetRevocationScheduled sets RevocationScheduled field to given value.

### HasRevocationScheduled

`func (o *RequestFull2) HasRevocationScheduled() bool`

HasRevocationScheduled returns a boolean if a field has been set.

### SetRevocationScheduledNil

`func (o *RequestFull2) SetRevocationScheduledNil(b bool)`

 SetRevocationScheduledNil sets the value for RevocationScheduled to be an explicit nil

### UnsetRevocationScheduled
`func (o *RequestFull2) UnsetRevocationScheduled()`

UnsetRevocationScheduled ensures that no value is present for RevocationScheduled, not even an explicit nil
### GetRequesterFieldValues

`func (o *RequestFull2) GetRequesterFieldValues() []RequestFieldValue`

GetRequesterFieldValues returns the RequesterFieldValues field if non-nil, zero value otherwise.

### GetRequesterFieldValuesOk

`func (o *RequestFull2) GetRequesterFieldValuesOk() (*[]RequestFieldValue, bool)`

GetRequesterFieldValuesOk returns a tuple with the RequesterFieldValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterFieldValues

`func (o *RequestFull2) SetRequesterFieldValues(v []RequestFieldValue)`

SetRequesterFieldValues sets RequesterFieldValues field to given value.

### HasRequesterFieldValues

`func (o *RequestFull2) HasRequesterFieldValues() bool`

HasRequesterFieldValues returns a boolean if a field has been set.

### GetRequestApproval

`func (o *RequestFull2) GetRequestApproval() RequestApproval2`

GetRequestApproval returns the RequestApproval field if non-nil, zero value otherwise.

### GetRequestApprovalOk

`func (o *RequestFull2) GetRequestApprovalOk() (*RequestApproval2, bool)`

GetRequestApprovalOk returns a tuple with the RequestApproval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestApproval

`func (o *RequestFull2) SetRequestApproval(v RequestApproval2)`

SetRequestApproval sets RequestApproval field to given value.

### HasRequestApproval

`func (o *RequestFull2) HasRequestApproval() bool`

HasRequestApproval returns a boolean if a field has been set.

### GetRiskAssessment

`func (o *RequestFull2) GetRiskAssessment() RiskAssessment`

GetRiskAssessment returns the RiskAssessment field if non-nil, zero value otherwise.

### GetRiskAssessmentOk

`func (o *RequestFull2) GetRiskAssessmentOk() (*RiskAssessment, bool)`

GetRiskAssessmentOk returns a tuple with the RiskAssessment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskAssessment

`func (o *RequestFull2) SetRiskAssessment(v RiskAssessment)`

SetRiskAssessment sets RiskAssessment field to given value.

### HasRiskAssessment

`func (o *RequestFull2) HasRiskAssessment() bool`

HasRiskAssessment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


