# RuleConflictCreatable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name of the resource risk rule | [optional] 
**Description** | Pointer to **string** | Description of the risk rule | [optional] 
**Notes** | Pointer to **string** | Additional information about the risk rule | [optional] 
**Type** | Pointer to **string** | Risk rule type | [optional] 
**Resources** | Pointer to [**[]RuleConflictResource**](RuleConflictResource.md) | Resources that the risk rule applies to | [optional] 
**ConflictCriteria** | Pointer to [**ConflictCriteriaCreatable**](ConflictCriteriaCreatable.md) |  | [optional] 

## Methods

### NewRuleConflictCreatable

`func NewRuleConflictCreatable() *RuleConflictCreatable`

NewRuleConflictCreatable instantiates a new RuleConflictCreatable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleConflictCreatableWithDefaults

`func NewRuleConflictCreatableWithDefaults() *RuleConflictCreatable`

NewRuleConflictCreatableWithDefaults instantiates a new RuleConflictCreatable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *RuleConflictCreatable) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RuleConflictCreatable) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RuleConflictCreatable) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RuleConflictCreatable) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *RuleConflictCreatable) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RuleConflictCreatable) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RuleConflictCreatable) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RuleConflictCreatable) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetNotes

`func (o *RuleConflictCreatable) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *RuleConflictCreatable) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *RuleConflictCreatable) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *RuleConflictCreatable) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetType

`func (o *RuleConflictCreatable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RuleConflictCreatable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RuleConflictCreatable) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RuleConflictCreatable) HasType() bool`

HasType returns a boolean if a field has been set.

### GetResources

`func (o *RuleConflictCreatable) GetResources() []RuleConflictResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *RuleConflictCreatable) GetResourcesOk() (*[]RuleConflictResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *RuleConflictCreatable) SetResources(v []RuleConflictResource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *RuleConflictCreatable) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetConflictCriteria

`func (o *RuleConflictCreatable) GetConflictCriteria() ConflictCriteriaCreatable`

GetConflictCriteria returns the ConflictCriteria field if non-nil, zero value otherwise.

### GetConflictCriteriaOk

`func (o *RuleConflictCreatable) GetConflictCriteriaOk() (*ConflictCriteriaCreatable, bool)`

GetConflictCriteriaOk returns a tuple with the ConflictCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictCriteria

`func (o *RuleConflictCreatable) SetConflictCriteria(v ConflictCriteriaCreatable)`

SetConflictCriteria sets ConflictCriteria field to given value.

### HasConflictCriteria

`func (o *RuleConflictCreatable) HasConflictCriteria() bool`

HasConflictCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


