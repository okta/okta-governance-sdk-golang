# RiskRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Links** | [**LinkSelf**](LinkSelf.md) |  | 
**Id** | **string** | Unique identifier for the object | 
**Description** | Pointer to **string** | Description for the risk rule | [optional] 
**Type** | **string** | Risk rule type | 
**ConflictCriteria** | [**ConflictCriteria**](ConflictCriteria.md) |  | 
**CreatedBy** | **string** | The &#x60;id&#x60; of the Okta user who created the resource | [readonly] 
**Created** | **time.Time** | The ISO 8601 formatted date and time when the resource was created | [readonly] 
**LastUpdated** | **time.Time** | The ISO 8601 formatted date and time when the object was last updated | [readonly] 
**LastUpdatedBy** | **string** | The &#x60;id&#x60; of the Okta user who last updated the object | [readonly] 
**Name** | **string** | Name of the resource risk rule | 
**Notes** | Pointer to **string** | Additional information about the risk rule | [optional] 
**Resources** | [**[]RuleConflictResource**](RuleConflictResource.md) | Resources that the risk rule applies to | 
**Status** | **string** | Status of the risk rule | 

## Methods

### NewRiskRuleResponse

`func NewRiskRuleResponse(links LinkSelf, id string, type_ string, conflictCriteria ConflictCriteria, createdBy string, created time.Time, lastUpdated time.Time, lastUpdatedBy string, name string, resources []RuleConflictResource, status string, ) *RiskRuleResponse`

NewRiskRuleResponse instantiates a new RiskRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskRuleResponseWithDefaults

`func NewRiskRuleResponseWithDefaults() *RiskRuleResponse`

NewRiskRuleResponseWithDefaults instantiates a new RiskRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinks

`func (o *RiskRuleResponse) GetLinks() LinkSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RiskRuleResponse) GetLinksOk() (*LinkSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RiskRuleResponse) SetLinks(v LinkSelf)`

SetLinks sets Links field to given value.


### GetId

`func (o *RiskRuleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RiskRuleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RiskRuleResponse) SetId(v string)`

SetId sets Id field to given value.


### GetDescription

`func (o *RiskRuleResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RiskRuleResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RiskRuleResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RiskRuleResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *RiskRuleResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RiskRuleResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RiskRuleResponse) SetType(v string)`

SetType sets Type field to given value.


### GetConflictCriteria

`func (o *RiskRuleResponse) GetConflictCriteria() ConflictCriteria`

GetConflictCriteria returns the ConflictCriteria field if non-nil, zero value otherwise.

### GetConflictCriteriaOk

`func (o *RiskRuleResponse) GetConflictCriteriaOk() (*ConflictCriteria, bool)`

GetConflictCriteriaOk returns a tuple with the ConflictCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflictCriteria

`func (o *RiskRuleResponse) SetConflictCriteria(v ConflictCriteria)`

SetConflictCriteria sets ConflictCriteria field to given value.


### GetCreatedBy

`func (o *RiskRuleResponse) GetCreatedBy() string`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *RiskRuleResponse) GetCreatedByOk() (*string, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *RiskRuleResponse) SetCreatedBy(v string)`

SetCreatedBy sets CreatedBy field to given value.


### GetCreated

`func (o *RiskRuleResponse) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *RiskRuleResponse) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *RiskRuleResponse) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetLastUpdated

`func (o *RiskRuleResponse) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *RiskRuleResponse) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *RiskRuleResponse) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.


### GetLastUpdatedBy

`func (o *RiskRuleResponse) GetLastUpdatedBy() string`

GetLastUpdatedBy returns the LastUpdatedBy field if non-nil, zero value otherwise.

### GetLastUpdatedByOk

`func (o *RiskRuleResponse) GetLastUpdatedByOk() (*string, bool)`

GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdatedBy

`func (o *RiskRuleResponse) SetLastUpdatedBy(v string)`

SetLastUpdatedBy sets LastUpdatedBy field to given value.


### GetName

`func (o *RiskRuleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RiskRuleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RiskRuleResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNotes

`func (o *RiskRuleResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *RiskRuleResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *RiskRuleResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *RiskRuleResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### GetResources

`func (o *RiskRuleResponse) GetResources() []RuleConflictResource`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *RiskRuleResponse) GetResourcesOk() (*[]RuleConflictResource, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *RiskRuleResponse) SetResources(v []RuleConflictResource)`

SetResources sets Resources field to given value.


### GetStatus

`func (o *RiskRuleResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *RiskRuleResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *RiskRuleResponse) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


