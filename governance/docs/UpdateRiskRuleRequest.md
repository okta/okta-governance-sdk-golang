# UpdateRiskRuleRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the object | 
**Name** | Pointer to **NullableString** | Name of the resource risk rule | [optional] 
**Notes** | Pointer to **NullableString** | Additional information about the rule | [optional] 
**Description** | Pointer to **NullableString** | Description of the risk rule | [optional] 
**ConflictCriteria** | Pointer to [**NullableConflictCriteriaUpdatable**](ConflictCriteriaUpdatable.md) |  | [optional] 

## Methods

### NewUpdateRiskRuleRequest

`func NewUpdateRiskRuleRequest(id string, ) *UpdateRiskRuleRequest`

NewUpdateRiskRuleRequest instantiates a new UpdateRiskRuleRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRiskRuleRequestWithDefaults

`func NewUpdateRiskRuleRequestWithDefaults() *UpdateRiskRuleRequest`

NewUpdateRiskRuleRequestWithDefaults instantiates a new UpdateRiskRuleRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateRiskRuleRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateRiskRuleRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateRiskRuleRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateRiskRuleRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateRiskRuleRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateRiskRuleRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateRiskRuleRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *UpdateRiskRuleRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *UpdateRiskRuleRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetNotes

`func (o *UpdateRiskRuleRequest) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UpdateRiskRuleRequest) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UpdateRiskRuleRequest) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UpdateRiskRuleRequest) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *UpdateRiskRuleRequest) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *UpdateRiskRuleRequest) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetDescription

`func (o *UpdateRiskRuleRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateRiskRuleRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateRiskRuleRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateRiskRuleRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateRiskRuleRequest) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateRiskRuleRequest) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetConflictCriteria

`func (o *UpdateRiskRuleRequest) GetConflictCriteria() ConflictCriteriaUpdatable`

GetConflictCriteria returns the ConflictCriteria field if non-nil, zero value otherwise.

### GetConflictCriteriaOk

`func (o *UpdateRiskRuleRequest) GetConflictCriteriaOk() (*ConflictCriteriaUpdatable, bool)`

GetConflictCriteriaOk returns a tuple with the ConflictCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictCriteria

`func (o *UpdateRiskRuleRequest) SetConflictCriteria(v ConflictCriteriaUpdatable)`

SetConflictCriteria sets ConflictCriteria field to given value.

### HasConflictCriteria

`func (o *UpdateRiskRuleRequest) HasConflictCriteria() bool`

HasConflictCriteria returns a boolean if a field has been set.

### SetConflictCriteriaNil

`func (o *UpdateRiskRuleRequest) SetConflictCriteriaNil(b bool)`

 SetConflictCriteriaNil sets the value for ConflictCriteria to be an explicit nil

### UnsetConflictCriteria
`func (o *UpdateRiskRuleRequest) UnsetConflictCriteria()`

UnsetConflictCriteria ensures that no value is present for ConflictCriteria, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


