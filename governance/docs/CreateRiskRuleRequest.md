# CreateRiskRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the resource risk rule | 
**Description** | Pointer to **string** | Description of the risk rule | [optional] 
**Notes** | Pointer to **string** | Additional information about the risk rule | [optional] 
**Type** | **string** | Risk rule type | 
**Resources** | [**[]RuleConflictResource**](RuleConflictResource.md) | Resources that the risk rule applies to | 
**ConflictCriteria** | [**ConflictCriteriaCreatable**](ConflictCriteriaCreatable.md) |  | 

## Methods

### NewCreateRiskRuleRequest

`func NewCreateRiskRuleRequest(name string, type_ string, resources []RuleConflictResource, conflictCriteria ConflictCriteriaCreatable, ) *CreateRiskRuleRequest`

NewCreateRiskRuleRequest instantiates a new CreateRiskRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRiskRuleRequestWithDefaults

`func NewCreateRiskRuleRequestWithDefaults() *CreateRiskRuleRequest`

NewCreateRiskRuleRequestWithDefaults instantiates a new CreateRiskRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateRiskRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateRiskRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateRiskRuleRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateRiskRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateRiskRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateRiskRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateRiskRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotes

`func (o *CreateRiskRuleRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreateRiskRuleRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreateRiskRuleRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreateRiskRuleRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetType

`func (o *CreateRiskRuleRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateRiskRuleRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateRiskRuleRequest) SetType(v string)`

SetType sets Type field to given value.


### GetResources

`func (o *CreateRiskRuleRequest) GetResources() []RuleConflictResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *CreateRiskRuleRequest) GetResourcesOk() (*[]RuleConflictResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *CreateRiskRuleRequest) SetResources(v []RuleConflictResource)`

SetResources sets Resources field to given value.


### GetConflictCriteria

`func (o *CreateRiskRuleRequest) GetConflictCriteria() ConflictCriteriaCreatable`

GetConflictCriteria returns the ConflictCriteria field if non-nil, zero value otherwise.

### GetConflictCriteriaOk

`func (o *CreateRiskRuleRequest) GetConflictCriteriaOk() (*ConflictCriteriaCreatable, bool)`

GetConflictCriteriaOk returns a tuple with the ConflictCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictCriteria

`func (o *CreateRiskRuleRequest) SetConflictCriteria(v ConflictCriteriaCreatable)`

SetConflictCriteria sets ConflictCriteria field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


