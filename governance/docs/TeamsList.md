# TeamsList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]Team**](Team.md) | All teams on the current page | [optional] 
**Links** | Pointer to [**TeamsListLinks**](TeamsListLinks.md) |  | [optional] 

## Methods

### NewTeamsList

`func NewTeamsList() *TeamsList`

NewTeamsList instantiates a new TeamsList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTeamsListWithDefaults

`func NewTeamsListWithDefaults() *TeamsList`

NewTeamsListWithDefaults instantiates a new TeamsList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TeamsList) GetData() []Team`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TeamsList) GetDataOk() (*[]Team, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TeamsList) SetData(v []Team)`

SetData sets Data field to given value.

### HasData

`func (o *TeamsList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetLinks

`func (o *TeamsList) GetLinks() TeamsListLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TeamsList) GetLinksOk() (*TeamsListLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TeamsList) SetLinks(v TeamsListLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *TeamsList) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


