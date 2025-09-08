# AutoRemediationSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeOnly** | Pointer to [**[]AutoRemediationSettingsIncludeOnlyInner**](AutoRemediationSettingsIncludeOnlyInner.md) | An array of resources to be automatically remediated | [optional] 
**IncludeAllIndirectAssignments** | Pointer to **bool** | If &#x60;includeAllIndirectAssignments&#x60; is set to &#x60;true&#x60;, the user&#39;s access to all groups that can assign the user to the application(s) will be removed during remediation. Only app assignments through groups can be automatically remediated. **Note:** You can only specify either  &#x60;includeAllIndirectAssignments&#x60; or &#x60;includeOnly&#x60;.  | [optional] 

## Methods

### NewAutoRemediationSettings

`func NewAutoRemediationSettings() *AutoRemediationSettings`

NewAutoRemediationSettings instantiates a new AutoRemediationSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoRemediationSettingsWithDefaults

`func NewAutoRemediationSettingsWithDefaults() *AutoRemediationSettings`

NewAutoRemediationSettingsWithDefaults instantiates a new AutoRemediationSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeOnly

`func (o *AutoRemediationSettings) GetIncludeOnly() []AutoRemediationSettingsIncludeOnlyInner`

GetIncludeOnly returns the IncludeOnly field if non-nil, zero value otherwise.

### GetIncludeOnlyOk

`func (o *AutoRemediationSettings) GetIncludeOnlyOk() (*[]AutoRemediationSettingsIncludeOnlyInner, bool)`

GetIncludeOnlyOk returns a tuple with the IncludeOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOnly

`func (o *AutoRemediationSettings) SetIncludeOnly(v []AutoRemediationSettingsIncludeOnlyInner)`

SetIncludeOnly sets IncludeOnly field to given value.

### HasIncludeOnly

`func (o *AutoRemediationSettings) HasIncludeOnly() bool`

HasIncludeOnly returns a boolean if a field has been set.

### GetIncludeAllIndirectAssignments

`func (o *AutoRemediationSettings) GetIncludeAllIndirectAssignments() bool`

GetIncludeAllIndirectAssignments returns the IncludeAllIndirectAssignments field if non-nil, zero value otherwise.

### GetIncludeAllIndirectAssignmentsOk

`func (o *AutoRemediationSettings) GetIncludeAllIndirectAssignmentsOk() (*bool, bool)`

GetIncludeAllIndirectAssignmentsOk returns a tuple with the IncludeAllIndirectAssignments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAllIndirectAssignments

`func (o *AutoRemediationSettings) SetIncludeAllIndirectAssignments(v bool)`

SetIncludeAllIndirectAssignments sets IncludeAllIndirectAssignments field to given value.

### HasIncludeAllIndirectAssignments

`func (o *AutoRemediationSettings) HasIncludeAllIndirectAssignments() bool`

HasIncludeAllIndirectAssignments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


