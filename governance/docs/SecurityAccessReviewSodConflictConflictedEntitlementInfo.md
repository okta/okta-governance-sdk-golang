# SecurityAccessReviewSodConflictConflictedEntitlementInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListName** | **string** | The entitlement list name | 
**Scope** | [**SecurityAccessReviewConflictedEntitlementListScope**](SecurityAccessReviewConflictedEntitlementListScope.md) |  | 
**Entitlements** | [**[]SecurityAccessReviewConflictedEntitlement**](SecurityAccessReviewConflictedEntitlement.md) | List of entitlements | 

## Methods

### NewSecurityAccessReviewSodConflictConflictedEntitlementInfo

`func NewSecurityAccessReviewSodConflictConflictedEntitlementInfo(listName string, scope SecurityAccessReviewConflictedEntitlementListScope, entitlements []SecurityAccessReviewConflictedEntitlement, ) *SecurityAccessReviewSodConflictConflictedEntitlementInfo`

NewSecurityAccessReviewSodConflictConflictedEntitlementInfo instantiates a new SecurityAccessReviewSodConflictConflictedEntitlementInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAccessReviewSodConflictConflictedEntitlementInfoWithDefaults

`func NewSecurityAccessReviewSodConflictConflictedEntitlementInfoWithDefaults() *SecurityAccessReviewSodConflictConflictedEntitlementInfo`

NewSecurityAccessReviewSodConflictConflictedEntitlementInfoWithDefaults instantiates a new SecurityAccessReviewSodConflictConflictedEntitlementInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListName

`func (o *SecurityAccessReviewSodConflictConflictedEntitlementInfo) GetListName() string`

GetListName returns the ListName field if non-nil, zero value otherwise.

### GetListNameOk

`func (o *SecurityAccessReviewSodConflictConflictedEntitlementInfo) GetListNameOk() (*string, bool)`

GetListNameOk returns a tuple with the ListName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListName

`func (o *SecurityAccessReviewSodConflictConflictedEntitlementInfo) SetListName(v string)`

SetListName sets ListName field to given value.


### GetScope

`func (o *SecurityAccessReviewSodConflictConflictedEntitlementInfo) GetScope() SecurityAccessReviewConflictedEntitlementListScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *SecurityAccessReviewSodConflictConflictedEntitlementInfo) GetScopeOk() (*SecurityAccessReviewConflictedEntitlementListScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *SecurityAccessReviewSodConflictConflictedEntitlementInfo) SetScope(v SecurityAccessReviewConflictedEntitlementListScope)`

SetScope sets Scope field to given value.


### GetEntitlements

`func (o *SecurityAccessReviewSodConflictConflictedEntitlementInfo) GetEntitlements() []SecurityAccessReviewConflictedEntitlement`

GetEntitlements returns the Entitlements field if non-nil, zero value otherwise.

### GetEntitlementsOk

`func (o *SecurityAccessReviewSodConflictConflictedEntitlementInfo) GetEntitlementsOk() (*[]SecurityAccessReviewConflictedEntitlement, bool)`

GetEntitlementsOk returns a tuple with the Entitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlements

`func (o *SecurityAccessReviewSodConflictConflictedEntitlementInfo) SetEntitlements(v []SecurityAccessReviewConflictedEntitlement)`

SetEntitlements sets Entitlements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


