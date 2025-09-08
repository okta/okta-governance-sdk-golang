# PrincipalScopeSettingsMutable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | [**PrincipalScopeType**](PrincipalScopeType.md) |  | 
**UserScopeExpression** | Pointer to **string** | The Okta expression language user expression on the &#x60;resourceSettings&#x60; to include users in the campaign. | [optional] 
**ExcludedUserIds** | Pointer to **[]string** | An array of Okta user IDs excluded from access certification or the campaign. This field is optional. A maximum of 50 users can be specified in the array. | [optional] 
**UserIds** | Pointer to **[]string** | An array of Okta user IDs included from access certification or the campaign. &#x60;userIds&#x60;, &#x60;groupIds&#x60; or &#x60;userScopeExpression&#x60; is required if campaign type is &#x60;USER&#x60;. A maximum of 100 users can be specified in the array. | [optional] 
**GroupIds** | Pointer to **[]string** | An array of Okta group IDs included from access certification or the campaign. &#x60;userIds&#x60;, &#x60;groupIds&#x60; or &#x60;userScopeExpression&#x60; is required if campaign type is &#x60;USER&#x60;. A maximum of 5 groups can be specified in the array. | [optional] 
**IncludeOnlyActiveUsers** | Pointer to **bool** | If set to &#x60;true&#x60;, only active Okta users are included in the campaign | [optional] 
**PredefinedInactiveUsersScope** | Pointer to [**PredefinedInactiveUsersScopeSettings**](PredefinedInactiveUsersScopeSettings.md) |  | [optional] 
**OnlyIncludeUsersWithSODConflicts** | Pointer to **bool** | If set to &#x60;true&#x60;, only includes users that have at least one SOD conflict that was caused due to entitlement(s) within Campaign scope | [optional] 

## Methods

### NewPrincipalScopeSettingsMutable

`func NewPrincipalScopeSettingsMutable(type_ PrincipalScopeType, ) *PrincipalScopeSettingsMutable`

NewPrincipalScopeSettingsMutable instantiates a new PrincipalScopeSettingsMutable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalScopeSettingsMutableWithDefaults

`func NewPrincipalScopeSettingsMutableWithDefaults() *PrincipalScopeSettingsMutable`

NewPrincipalScopeSettingsMutableWithDefaults instantiates a new PrincipalScopeSettingsMutable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PrincipalScopeSettingsMutable) GetType() PrincipalScopeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PrincipalScopeSettingsMutable) GetTypeOk() (*PrincipalScopeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PrincipalScopeSettingsMutable) SetType(v PrincipalScopeType)`

SetType sets Type field to given value.


### GetUserScopeExpression

`func (o *PrincipalScopeSettingsMutable) GetUserScopeExpression() string`

GetUserScopeExpression returns the UserScopeExpression field if non-nil, zero value otherwise.

### GetUserScopeExpressionOk

`func (o *PrincipalScopeSettingsMutable) GetUserScopeExpressionOk() (*string, bool)`

GetUserScopeExpressionOk returns a tuple with the UserScopeExpression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserScopeExpression

`func (o *PrincipalScopeSettingsMutable) SetUserScopeExpression(v string)`

SetUserScopeExpression sets UserScopeExpression field to given value.

### HasUserScopeExpression

`func (o *PrincipalScopeSettingsMutable) HasUserScopeExpression() bool`

HasUserScopeExpression returns a boolean if a field has been set.

### GetExcludedUserIds

`func (o *PrincipalScopeSettingsMutable) GetExcludedUserIds() []string`

GetExcludedUserIds returns the ExcludedUserIds field if non-nil, zero value otherwise.

### GetExcludedUserIdsOk

`func (o *PrincipalScopeSettingsMutable) GetExcludedUserIdsOk() (*[]string, bool)`

GetExcludedUserIdsOk returns a tuple with the ExcludedUserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedUserIds

`func (o *PrincipalScopeSettingsMutable) SetExcludedUserIds(v []string)`

SetExcludedUserIds sets ExcludedUserIds field to given value.

### HasExcludedUserIds

`func (o *PrincipalScopeSettingsMutable) HasExcludedUserIds() bool`

HasExcludedUserIds returns a boolean if a field has been set.

### GetUserIds

`func (o *PrincipalScopeSettingsMutable) GetUserIds() []string`

GetUserIds returns the UserIds field if non-nil, zero value otherwise.

### GetUserIdsOk

`func (o *PrincipalScopeSettingsMutable) GetUserIdsOk() (*[]string, bool)`

GetUserIdsOk returns a tuple with the UserIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserIds

`func (o *PrincipalScopeSettingsMutable) SetUserIds(v []string)`

SetUserIds sets UserIds field to given value.

### HasUserIds

`func (o *PrincipalScopeSettingsMutable) HasUserIds() bool`

HasUserIds returns a boolean if a field has been set.

### SetUserIdsNil

`func (o *PrincipalScopeSettingsMutable) SetUserIdsNil(b bool)`

 SetUserIdsNil sets the value for UserIds to be an explicit nil

### UnsetUserIds
`func (o *PrincipalScopeSettingsMutable) UnsetUserIds()`

UnsetUserIds ensures that no value is present for UserIds, not even an explicit nil
### GetGroupIds

`func (o *PrincipalScopeSettingsMutable) GetGroupIds() []string`

GetGroupIds returns the GroupIds field if non-nil, zero value otherwise.

### GetGroupIdsOk

`func (o *PrincipalScopeSettingsMutable) GetGroupIdsOk() (*[]string, bool)`

GetGroupIdsOk returns a tuple with the GroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupIds

`func (o *PrincipalScopeSettingsMutable) SetGroupIds(v []string)`

SetGroupIds sets GroupIds field to given value.

### HasGroupIds

`func (o *PrincipalScopeSettingsMutable) HasGroupIds() bool`

HasGroupIds returns a boolean if a field has been set.

### SetGroupIdsNil

`func (o *PrincipalScopeSettingsMutable) SetGroupIdsNil(b bool)`

 SetGroupIdsNil sets the value for GroupIds to be an explicit nil

### UnsetGroupIds
`func (o *PrincipalScopeSettingsMutable) UnsetGroupIds()`

UnsetGroupIds ensures that no value is present for GroupIds, not even an explicit nil
### GetIncludeOnlyActiveUsers

`func (o *PrincipalScopeSettingsMutable) GetIncludeOnlyActiveUsers() bool`

GetIncludeOnlyActiveUsers returns the IncludeOnlyActiveUsers field if non-nil, zero value otherwise.

### GetIncludeOnlyActiveUsersOk

`func (o *PrincipalScopeSettingsMutable) GetIncludeOnlyActiveUsersOk() (*bool, bool)`

GetIncludeOnlyActiveUsersOk returns a tuple with the IncludeOnlyActiveUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOnlyActiveUsers

`func (o *PrincipalScopeSettingsMutable) SetIncludeOnlyActiveUsers(v bool)`

SetIncludeOnlyActiveUsers sets IncludeOnlyActiveUsers field to given value.

### HasIncludeOnlyActiveUsers

`func (o *PrincipalScopeSettingsMutable) HasIncludeOnlyActiveUsers() bool`

HasIncludeOnlyActiveUsers returns a boolean if a field has been set.

### GetPredefinedInactiveUsersScope

`func (o *PrincipalScopeSettingsMutable) GetPredefinedInactiveUsersScope() PredefinedInactiveUsersScopeSettings`

GetPredefinedInactiveUsersScope returns the PredefinedInactiveUsersScope field if non-nil, zero value otherwise.

### GetPredefinedInactiveUsersScopeOk

`func (o *PrincipalScopeSettingsMutable) GetPredefinedInactiveUsersScopeOk() (*PredefinedInactiveUsersScopeSettings, bool)`

GetPredefinedInactiveUsersScopeOk returns a tuple with the PredefinedInactiveUsersScope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredefinedInactiveUsersScope

`func (o *PrincipalScopeSettingsMutable) SetPredefinedInactiveUsersScope(v PredefinedInactiveUsersScopeSettings)`

SetPredefinedInactiveUsersScope sets PredefinedInactiveUsersScope field to given value.

### HasPredefinedInactiveUsersScope

`func (o *PrincipalScopeSettingsMutable) HasPredefinedInactiveUsersScope() bool`

HasPredefinedInactiveUsersScope returns a boolean if a field has been set.

### GetOnlyIncludeUsersWithSODConflicts

`func (o *PrincipalScopeSettingsMutable) GetOnlyIncludeUsersWithSODConflicts() bool`

GetOnlyIncludeUsersWithSODConflicts returns the OnlyIncludeUsersWithSODConflicts field if non-nil, zero value otherwise.

### GetOnlyIncludeUsersWithSODConflictsOk

`func (o *PrincipalScopeSettingsMutable) GetOnlyIncludeUsersWithSODConflictsOk() (*bool, bool)`

GetOnlyIncludeUsersWithSODConflictsOk returns a tuple with the OnlyIncludeUsersWithSODConflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyIncludeUsersWithSODConflicts

`func (o *PrincipalScopeSettingsMutable) SetOnlyIncludeUsersWithSODConflicts(v bool)`

SetOnlyIncludeUsersWithSODConflicts sets OnlyIncludeUsersWithSODConflicts field to given value.

### HasOnlyIncludeUsersWithSODConflicts

`func (o *PrincipalScopeSettingsMutable) HasOnlyIncludeUsersWithSODConflicts() bool`

HasOnlyIncludeUsersWithSODConflicts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


