# RiskRuleConflicts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier of the risk rule | [optional] 
**Type** | Pointer to **string** | The type of the rule | [optional] 
**ConflictCriteria** | Pointer to [**RiskRuleConflictsConflictCriteria**](RiskRuleConflictsConflictCriteria.md) |  | [optional] 

## Methods

### NewRiskRuleConflicts

`func NewRiskRuleConflicts() *RiskRuleConflicts`

NewRiskRuleConflicts instantiates a new RiskRuleConflicts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRuleConflictsWithDefaults

`func NewRiskRuleConflictsWithDefaults() *RiskRuleConflicts`

NewRiskRuleConflictsWithDefaults instantiates a new RiskRuleConflicts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RiskRuleConflicts) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskRuleConflicts) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskRuleConflicts) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RiskRuleConflicts) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RiskRuleConflicts) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskRuleConflicts) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskRuleConflicts) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RiskRuleConflicts) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConflictCriteria

`func (o *RiskRuleConflicts) GetConflictCriteria() RiskRuleConflictsConflictCriteria`

GetConflictCriteria returns the ConflictCriteria field if non-nil, zero value otherwise.

### GetConflictCriteriaOk

`func (o *RiskRuleConflicts) GetConflictCriteriaOk() (*RiskRuleConflictsConflictCriteria, bool)`

GetConflictCriteriaOk returns a tuple with the ConflictCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictCriteria

`func (o *RiskRuleConflicts) SetConflictCriteria(v RiskRuleConflictsConflictCriteria)`

SetConflictCriteria sets ConflictCriteria field to given value.

### HasConflictCriteria

`func (o *RiskRuleConflicts) HasConflictCriteria() bool`

HasConflictCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


