# OrgSettingsPatchableDelegatesEnduser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Permissions** | Pointer to **[]string** | The permission that applies to this setting  | Permission | Description | |------|------| | &#x60;READ&#x60; | Allow end users to view their delegates | | &#x60;WRITE&#x60; | Allow end users to set their own delegates | | [optional] 

## Methods

### NewOrgSettingsPatchableDelegatesEnduser

`func NewOrgSettingsPatchableDelegatesEnduser() *OrgSettingsPatchableDelegatesEnduser`

NewOrgSettingsPatchableDelegatesEnduser instantiates a new OrgSettingsPatchableDelegatesEnduser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgSettingsPatchableDelegatesEnduserWithDefaults

`func NewOrgSettingsPatchableDelegatesEnduserWithDefaults() *OrgSettingsPatchableDelegatesEnduser`

NewOrgSettingsPatchableDelegatesEnduserWithDefaults instantiates a new OrgSettingsPatchableDelegatesEnduser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPermissions

`func (o *OrgSettingsPatchableDelegatesEnduser) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *OrgSettingsPatchableDelegatesEnduser) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *OrgSettingsPatchableDelegatesEnduser) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *OrgSettingsPatchableDelegatesEnduser) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


