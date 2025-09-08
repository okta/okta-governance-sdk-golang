# SecurityAccessReviewAnomaly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**SecurityAccessReviewAnomalyType**](SecurityAccessReviewAnomalyType.md) |  | 
**Severity** | [**SecurityAccessReviewAccessItemSeverity**](SecurityAccessReviewAccessItemSeverity.md) |  | 
**Subtext** | [**ServerMessage**](ServerMessage.md) |  | 
**SodConflicts** | Pointer to [**[]SecurityAccessReviewSodConflict**](SecurityAccessReviewSodConflict.md) | A list of separation of duties (SOD) conflicts caused by a user&#39;s access to an entitlement | [optional] 

## Methods

### NewSecurityAccessReviewAnomaly

`func NewSecurityAccessReviewAnomaly(type_ SecurityAccessReviewAnomalyType, severity SecurityAccessReviewAccessItemSeverity, subtext ServerMessage, ) *SecurityAccessReviewAnomaly`

NewSecurityAccessReviewAnomaly instantiates a new SecurityAccessReviewAnomaly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAccessReviewAnomalyWithDefaults

`func NewSecurityAccessReviewAnomalyWithDefaults() *SecurityAccessReviewAnomaly`

NewSecurityAccessReviewAnomalyWithDefaults instantiates a new SecurityAccessReviewAnomaly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SecurityAccessReviewAnomaly) GetType() SecurityAccessReviewAnomalyType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityAccessReviewAnomaly) GetTypeOk() (*SecurityAccessReviewAnomalyType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityAccessReviewAnomaly) SetType(v SecurityAccessReviewAnomalyType)`

SetType sets Type field to given value.


### GetSeverity

`func (o *SecurityAccessReviewAnomaly) GetSeverity() SecurityAccessReviewAccessItemSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *SecurityAccessReviewAnomaly) GetSeverityOk() (*SecurityAccessReviewAccessItemSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *SecurityAccessReviewAnomaly) SetSeverity(v SecurityAccessReviewAccessItemSeverity)`

SetSeverity sets Severity field to given value.


### GetSubtext

`func (o *SecurityAccessReviewAnomaly) GetSubtext() ServerMessage`

GetSubtext returns the Subtext field if non-nil, zero value otherwise.

### GetSubtextOk

`func (o *SecurityAccessReviewAnomaly) GetSubtextOk() (*ServerMessage, bool)`

GetSubtextOk returns a tuple with the Subtext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubtext

`func (o *SecurityAccessReviewAnomaly) SetSubtext(v ServerMessage)`

SetSubtext sets Subtext field to given value.


### GetSodConflicts

`func (o *SecurityAccessReviewAnomaly) GetSodConflicts() []SecurityAccessReviewSodConflict`

GetSodConflicts returns the SodConflicts field if non-nil, zero value otherwise.

### GetSodConflictsOk

`func (o *SecurityAccessReviewAnomaly) GetSodConflictsOk() (*[]SecurityAccessReviewSodConflict, bool)`

GetSodConflictsOk returns a tuple with the SodConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSodConflicts

`func (o *SecurityAccessReviewAnomaly) SetSodConflicts(v []SecurityAccessReviewSodConflict)`

SetSodConflicts sets SodConflicts field to given value.

### HasSodConflicts

`func (o *SecurityAccessReviewAnomaly) HasSodConflicts() bool`

HasSodConflicts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


