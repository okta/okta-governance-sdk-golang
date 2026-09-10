# EntitlementDriftFull

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the object | 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 
**Links** | [**EntitlementDriftLinks**](EntitlementDriftLinks.md) |  | 
**ResourceOrn** | **string** | The Okta resource in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn)  See the ORN format for [supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).  | 
**PrincipalOrn** | **string** | The Okta user in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format | 
**ImportJobId** | **string** | The &#x60;id&#x60; of the import job that surfaced this drift | 
**DriftType** | [**EntitlementDriftType**](EntitlementDriftType.md) |  | 
**Status** | [**EntitlementDriftStatus**](EntitlementDriftStatus.md) |  | 
**Resolution** | [**EntitlementDriftResolution**](EntitlementDriftResolution.md) |  | 
**Entitlement** | [**EntitlementDriftEntitlementReference**](EntitlementDriftEntitlementReference.md) |  | 
**AffectedGrantSource** | Pointer to [**EntitlementDriftAffectedGrantSource**](EntitlementDriftAffectedGrantSource.md) |  | [optional] 
**AffectedGrantId** | Pointer to **string** | The &#x60;id&#x60; of the grant that held the affected entitlement when the drift was detected.  | [optional] 
**ResolvedGrantId** | Pointer to **string** | The &#x60;id&#x60; of the grant that resolution created or changed.  | [optional] 
**RevertedAt** | Pointer to **time.Time** | When the drift was reverted. Present only on a reverted drift. | [optional] 
**RevertedBy** | Pointer to **string** | The principal who reverted the drift. Present only on a reverted drift. | [optional] 
**ErrorMessage** | Pointer to **string** | A message that explains why Okta couldn&#39;t resolve the drift. Present only when &#x60;status&#x60; is &#x60;ERROR&#x60;. | [optional] 

## Methods

### NewEntitlementDriftFull

`func NewEntitlementDriftFull(id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, links EntitlementDriftLinks, resourceOrn string, principalOrn string, importJobId string, driftType EntitlementDriftType, status EntitlementDriftStatus, resolution EntitlementDriftResolution, entitlement EntitlementDriftEntitlementReference, ) *EntitlementDriftFull`

NewEntitlementDriftFull instantiates a new EntitlementDriftFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementDriftFullWithDefaults

`func NewEntitlementDriftFullWithDefaults() *EntitlementDriftFull`

NewEntitlementDriftFullWithDefaults instantiates a new EntitlementDriftFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntitlementDriftFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementDriftFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementDriftFull) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedBy

`func (o *EntitlementDriftFull) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *EntitlementDriftFull) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *EntitlementDriftFull) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *EntitlementDriftFull) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *EntitlementDriftFull) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *EntitlementDriftFull) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *EntitlementDriftFull) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *EntitlementDriftFull) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *EntitlementDriftFull) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *EntitlementDriftFull) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *EntitlementDriftFull) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *EntitlementDriftFull) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.


### GetLinks

`func (o *EntitlementDriftFull) GetLinks() EntitlementDriftLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *EntitlementDriftFull) GetLinksOk() (*EntitlementDriftLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *EntitlementDriftFull) SetLinks(v EntitlementDriftLinks)`

SetLinks sets Links field to given value.


### GetResourceOrn

`func (o *EntitlementDriftFull) GetResourceOrn() string`

GetResourceOrn returns the ResourceOrn field if non-nil, zero value otherwise.

### GetResourceOrnOk

`func (o *EntitlementDriftFull) GetResourceOrnOk() (*string, bool)`

GetResourceOrnOk returns a tuple with the ResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrn

`func (o *EntitlementDriftFull) SetResourceOrn(v string)`

SetResourceOrn sets ResourceOrn field to given value.


### GetPrincipalOrn

`func (o *EntitlementDriftFull) GetPrincipalOrn() string`

GetPrincipalOrn returns the PrincipalOrn field if non-nil, zero value otherwise.

### GetPrincipalOrnOk

`func (o *EntitlementDriftFull) GetPrincipalOrnOk() (*string, bool)`

GetPrincipalOrnOk returns a tuple with the PrincipalOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalOrn

`func (o *EntitlementDriftFull) SetPrincipalOrn(v string)`

SetPrincipalOrn sets PrincipalOrn field to given value.


### GetImportJobId

`func (o *EntitlementDriftFull) GetImportJobId() string`

GetImportJobId returns the ImportJobId field if non-nil, zero value otherwise.

### GetImportJobIdOk

`func (o *EntitlementDriftFull) GetImportJobIdOk() (*string, bool)`

GetImportJobIdOk returns a tuple with the ImportJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportJobId

`func (o *EntitlementDriftFull) SetImportJobId(v string)`

SetImportJobId sets ImportJobId field to given value.


### GetDriftType

`func (o *EntitlementDriftFull) GetDriftType() EntitlementDriftType`

GetDriftType returns the DriftType field if non-nil, zero value otherwise.

### GetDriftTypeOk

`func (o *EntitlementDriftFull) GetDriftTypeOk() (*EntitlementDriftType, bool)`

GetDriftTypeOk returns a tuple with the DriftType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriftType

`func (o *EntitlementDriftFull) SetDriftType(v EntitlementDriftType)`

SetDriftType sets DriftType field to given value.


### GetStatus

`func (o *EntitlementDriftFull) GetStatus() EntitlementDriftStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EntitlementDriftFull) GetStatusOk() (*EntitlementDriftStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EntitlementDriftFull) SetStatus(v EntitlementDriftStatus)`

SetStatus sets Status field to given value.


### GetResolution

`func (o *EntitlementDriftFull) GetResolution() EntitlementDriftResolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *EntitlementDriftFull) GetResolutionOk() (*EntitlementDriftResolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *EntitlementDriftFull) SetResolution(v EntitlementDriftResolution)`

SetResolution sets Resolution field to given value.


### GetEntitlement

`func (o *EntitlementDriftFull) GetEntitlement() EntitlementDriftEntitlementReference`

GetEntitlement returns the Entitlement field if non-nil, zero value otherwise.

### GetEntitlementOk

`func (o *EntitlementDriftFull) GetEntitlementOk() (*EntitlementDriftEntitlementReference, bool)`

GetEntitlementOk returns a tuple with the Entitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlement

`func (o *EntitlementDriftFull) SetEntitlement(v EntitlementDriftEntitlementReference)`

SetEntitlement sets Entitlement field to given value.


### GetAffectedGrantSource

`func (o *EntitlementDriftFull) GetAffectedGrantSource() EntitlementDriftAffectedGrantSource`

GetAffectedGrantSource returns the AffectedGrantSource field if non-nil, zero value otherwise.

### GetAffectedGrantSourceOk

`func (o *EntitlementDriftFull) GetAffectedGrantSourceOk() (*EntitlementDriftAffectedGrantSource, bool)`

GetAffectedGrantSourceOk returns a tuple with the AffectedGrantSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedGrantSource

`func (o *EntitlementDriftFull) SetAffectedGrantSource(v EntitlementDriftAffectedGrantSource)`

SetAffectedGrantSource sets AffectedGrantSource field to given value.

### HasAffectedGrantSource

`func (o *EntitlementDriftFull) HasAffectedGrantSource() bool`

HasAffectedGrantSource returns a boolean if a field has been set.

### GetAffectedGrantId

`func (o *EntitlementDriftFull) GetAffectedGrantId() string`

GetAffectedGrantId returns the AffectedGrantId field if non-nil, zero value otherwise.

### GetAffectedGrantIdOk

`func (o *EntitlementDriftFull) GetAffectedGrantIdOk() (*string, bool)`

GetAffectedGrantIdOk returns a tuple with the AffectedGrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffectedGrantId

`func (o *EntitlementDriftFull) SetAffectedGrantId(v string)`

SetAffectedGrantId sets AffectedGrantId field to given value.

### HasAffectedGrantId

`func (o *EntitlementDriftFull) HasAffectedGrantId() bool`

HasAffectedGrantId returns a boolean if a field has been set.

### GetResolvedGrantId

`func (o *EntitlementDriftFull) GetResolvedGrantId() string`

GetResolvedGrantId returns the ResolvedGrantId field if non-nil, zero value otherwise.

### GetResolvedGrantIdOk

`func (o *EntitlementDriftFull) GetResolvedGrantIdOk() (*string, bool)`

GetResolvedGrantIdOk returns a tuple with the ResolvedGrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedGrantId

`func (o *EntitlementDriftFull) SetResolvedGrantId(v string)`

SetResolvedGrantId sets ResolvedGrantId field to given value.

### HasResolvedGrantId

`func (o *EntitlementDriftFull) HasResolvedGrantId() bool`

HasResolvedGrantId returns a boolean if a field has been set.

### GetRevertedAt

`func (o *EntitlementDriftFull) GetRevertedAt() time.Time`

GetRevertedAt returns the RevertedAt field if non-nil, zero value otherwise.

### GetRevertedAtOk

`func (o *EntitlementDriftFull) GetRevertedAtOk() (*time.Time, bool)`

GetRevertedAtOk returns a tuple with the RevertedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevertedAt

`func (o *EntitlementDriftFull) SetRevertedAt(v time.Time)`

SetRevertedAt sets RevertedAt field to given value.

### HasRevertedAt

`func (o *EntitlementDriftFull) HasRevertedAt() bool`

HasRevertedAt returns a boolean if a field has been set.

### GetRevertedBy

`func (o *EntitlementDriftFull) GetRevertedBy() string`

GetRevertedBy returns the RevertedBy field if non-nil, zero value otherwise.

### GetRevertedByOk

`func (o *EntitlementDriftFull) GetRevertedByOk() (*string, bool)`

GetRevertedByOk returns a tuple with the RevertedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevertedBy

`func (o *EntitlementDriftFull) SetRevertedBy(v string)`

SetRevertedBy sets RevertedBy field to given value.

### HasRevertedBy

`func (o *EntitlementDriftFull) HasRevertedBy() bool`

HasRevertedBy returns a boolean if a field has been set.

### GetErrorMessage

`func (o *EntitlementDriftFull) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *EntitlementDriftFull) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *EntitlementDriftFull) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *EntitlementDriftFull) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


