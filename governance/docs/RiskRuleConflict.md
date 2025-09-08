# RiskRuleConflict

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique identifier for the object | 
**Description** | Pointer to **string** | Description for the risk rule | [optional] 
**Type** | Pointer to **string** | Risk rule type | [optional] 
**ConflictCriteria** | Pointer to [**ConflictCriteria**](ConflictCriteria.md) |  | [optional] 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 
**Name** | Pointer to **string** | Name of the resource risk rule | [optional] 
**Notes** | Pointer to **string** | Additional information about the risk rule | [optional] 
**Resources** | Pointer to [**[]RuleConflictResource**](RuleConflictResource.md) | Resources that the risk rule applies to | [optional] 
**Status** | Pointer to **string** | Status of the risk rule | [optional] 

## Methods

### NewRiskRuleConflict

`func NewRiskRuleConflict(id string, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, ) *RiskRuleConflict`

NewRiskRuleConflict instantiates a new RiskRuleConflict object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRuleConflictWithDefaults

`func NewRiskRuleConflictWithDefaults() *RiskRuleConflict`

NewRiskRuleConflictWithDefaults instantiates a new RiskRuleConflict object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RiskRuleConflict) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskRuleConflict) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskRuleConflict) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *RiskRuleConflict) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RiskRuleConflict) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RiskRuleConflict) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RiskRuleConflict) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *RiskRuleConflict) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskRuleConflict) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskRuleConflict) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RiskRuleConflict) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConflictCriteria

`func (o *RiskRuleConflict) GetConflictCriteria() ConflictCriteria`

GetConflictCriteria returns the ConflictCriteria field if non-nil, zero value otherwise.

### GetConflictCriteriaOk

`func (o *RiskRuleConflict) GetConflictCriteriaOk() (*ConflictCriteria, bool)`

GetConflictCriteriaOk returns a tuple with the ConflictCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictCriteria

`func (o *RiskRuleConflict) SetConflictCriteria(v ConflictCriteria)`

SetConflictCriteria sets ConflictCriteria field to given value.

### HasConflictCriteria

`func (o *RiskRuleConflict) HasConflictCriteria() bool`

HasConflictCriteria returns a boolean if a field has been set.

### GetCreatedBy

`func (o *RiskRuleConflict) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RiskRuleConflict) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RiskRuleConflict) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *RiskRuleConflict) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RiskRuleConflict) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RiskRuleConflict) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *RiskRuleConflict) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RiskRuleConflict) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RiskRuleConflict) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *RiskRuleConflict) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *RiskRuleConflict) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *RiskRuleConflict) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.


### GetLinks

`func (o *RiskRuleConflict) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RiskRuleConflict) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RiskRuleConflict) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *RiskRuleConflict) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetName

`func (o *RiskRuleConflict) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskRuleConflict) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskRuleConflict) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RiskRuleConflict) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNotes

`func (o *RiskRuleConflict) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *RiskRuleConflict) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *RiskRuleConflict) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *RiskRuleConflict) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetResources

`func (o *RiskRuleConflict) GetResources() []RuleConflictResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *RiskRuleConflict) GetResourcesOk() (*[]RuleConflictResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *RiskRuleConflict) SetResources(v []RuleConflictResource)`

SetResources sets Resources field to given value.

### HasResources

`func (o *RiskRuleConflict) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetStatus

`func (o *RiskRuleConflict) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RiskRuleConflict) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RiskRuleConflict) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *RiskRuleConflict) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


