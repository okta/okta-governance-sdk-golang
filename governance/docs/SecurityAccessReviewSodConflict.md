# SecurityAccessReviewSodConflict

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | ID of SOD conflict rule | 
**Name** | **string** | Name of SOD conflict rule | 
**Description** | Pointer to **string** | Description of SOD conflict rule | [optional] 
**Operator** | [**SecurityAccessReviewSodConflictOperator**](SecurityAccessReviewSodConflictOperator.md) |  | 
**ConflictedEntitlements** | [**[]SecurityAccessReviewSodConflictConflictedEntitlementInfo**](SecurityAccessReviewSodConflictConflictedEntitlementInfo.md) | List of conflicted entitlements that caused this SOD conflict | 

## Methods

### NewSecurityAccessReviewSodConflict

`func NewSecurityAccessReviewSodConflict(id string, name string, operator SecurityAccessReviewSodConflictOperator, conflictedEntitlements []SecurityAccessReviewSodConflictConflictedEntitlementInfo, ) *SecurityAccessReviewSodConflict`

NewSecurityAccessReviewSodConflict instantiates a new SecurityAccessReviewSodConflict object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAccessReviewSodConflictWithDefaults

`func NewSecurityAccessReviewSodConflictWithDefaults() *SecurityAccessReviewSodConflict`

NewSecurityAccessReviewSodConflictWithDefaults instantiates a new SecurityAccessReviewSodConflict object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityAccessReviewSodConflict) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityAccessReviewSodConflict) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityAccessReviewSodConflict) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *SecurityAccessReviewSodConflict) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityAccessReviewSodConflict) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityAccessReviewSodConflict) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SecurityAccessReviewSodConflict) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityAccessReviewSodConflict) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityAccessReviewSodConflict) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityAccessReviewSodConflict) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetOperator

`func (o *SecurityAccessReviewSodConflict) GetOperator() SecurityAccessReviewSodConflictOperator`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *SecurityAccessReviewSodConflict) GetOperatorOk() (*SecurityAccessReviewSodConflictOperator, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *SecurityAccessReviewSodConflict) SetOperator(v SecurityAccessReviewSodConflictOperator)`

SetOperator sets Operator field to given value.


### GetConflictedEntitlements

`func (o *SecurityAccessReviewSodConflict) GetConflictedEntitlements() []SecurityAccessReviewSodConflictConflictedEntitlementInfo`

GetConflictedEntitlements returns the ConflictedEntitlements field if non-nil, zero value otherwise.

### GetConflictedEntitlementsOk

`func (o *SecurityAccessReviewSodConflict) GetConflictedEntitlementsOk() (*[]SecurityAccessReviewSodConflictConflictedEntitlementInfo, bool)`

GetConflictedEntitlementsOk returns a tuple with the ConflictedEntitlements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictedEntitlements

`func (o *SecurityAccessReviewSodConflict) SetConflictedEntitlements(v []SecurityAccessReviewSodConflictConflictedEntitlementInfo)`

SetConflictedEntitlements sets ConflictedEntitlements field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


