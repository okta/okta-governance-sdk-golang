# SecurityAccessReviewPrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The Okta user &#x60;id&#x60; | 
**Email** | Pointer to **string** | The Okta user&#39;s email address | [optional] 
**FirstName** | Pointer to **string** | The Okta user&#39;s first name | [optional] 
**LastName** | Pointer to **string** | The Okta user&#39;s last name | [optional] 
**Login** | Pointer to **string** | The Okta user&#39;s login | [optional] 
**Status** | [**PrincipalProfileStatus**](PrincipalProfileStatus.md) |  | 
**Type** | [**PrincipalProfileType**](PrincipalProfileType.md) |  | 
**Department** | Pointer to **string** | The department of the Okta user | [optional] 
**Manager** | Pointer to **string** | The manager of the Okta user | [optional] 
**Role** | Pointer to **string** | The Okta user role | [optional] 
**HomeLocation** | Pointer to [**SecurityAccessReviewPrincipalLocation**](SecurityAccessReviewPrincipalLocation.md) |  | [optional] 
**LastLoginInfo** | Pointer to [**SecurityAccessReviewPrincipalLastLoginInfo**](SecurityAccessReviewPrincipalLastLoginInfo.md) |  | [optional] 
**OktaAdminRoles** | Pointer to [**[]SecurityAccessReviewPrincipalOktaAdminRole**](SecurityAccessReviewPrincipalOktaAdminRole.md) | List of Okta admin roles assigned to the principal | [optional] 
**Links** | Pointer to [**map[string]Link**](Link.md) |  | [optional] 

## Methods

### NewSecurityAccessReviewPrincipal

`func NewSecurityAccessReviewPrincipal(id string, status PrincipalProfileStatus, type_ PrincipalProfileType, ) *SecurityAccessReviewPrincipal`

NewSecurityAccessReviewPrincipal instantiates a new SecurityAccessReviewPrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityAccessReviewPrincipalWithDefaults

`func NewSecurityAccessReviewPrincipalWithDefaults() *SecurityAccessReviewPrincipal`

NewSecurityAccessReviewPrincipalWithDefaults instantiates a new SecurityAccessReviewPrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SecurityAccessReviewPrincipal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityAccessReviewPrincipal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityAccessReviewPrincipal) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *SecurityAccessReviewPrincipal) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SecurityAccessReviewPrincipal) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SecurityAccessReviewPrincipal) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SecurityAccessReviewPrincipal) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *SecurityAccessReviewPrincipal) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *SecurityAccessReviewPrincipal) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *SecurityAccessReviewPrincipal) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *SecurityAccessReviewPrincipal) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *SecurityAccessReviewPrincipal) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *SecurityAccessReviewPrincipal) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *SecurityAccessReviewPrincipal) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *SecurityAccessReviewPrincipal) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetLogin

`func (o *SecurityAccessReviewPrincipal) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *SecurityAccessReviewPrincipal) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *SecurityAccessReviewPrincipal) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *SecurityAccessReviewPrincipal) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetStatus

`func (o *SecurityAccessReviewPrincipal) GetStatus() PrincipalProfileStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SecurityAccessReviewPrincipal) GetStatusOk() (*PrincipalProfileStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SecurityAccessReviewPrincipal) SetStatus(v PrincipalProfileStatus)`

SetStatus sets Status field to given value.


### GetType

`func (o *SecurityAccessReviewPrincipal) GetType() PrincipalProfileType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityAccessReviewPrincipal) GetTypeOk() (*PrincipalProfileType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityAccessReviewPrincipal) SetType(v PrincipalProfileType)`

SetType sets Type field to given value.


### GetDepartment

`func (o *SecurityAccessReviewPrincipal) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *SecurityAccessReviewPrincipal) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *SecurityAccessReviewPrincipal) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *SecurityAccessReviewPrincipal) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### GetManager

`func (o *SecurityAccessReviewPrincipal) GetManager() string`

GetManager returns the Manager field if non-nil, zero value otherwise.

### GetManagerOk

`func (o *SecurityAccessReviewPrincipal) GetManagerOk() (*string, bool)`

GetManagerOk returns a tuple with the Manager field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManager

`func (o *SecurityAccessReviewPrincipal) SetManager(v string)`

SetManager sets Manager field to given value.

### HasManager

`func (o *SecurityAccessReviewPrincipal) HasManager() bool`

HasManager returns a boolean if a field has been set.

### GetRole

`func (o *SecurityAccessReviewPrincipal) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *SecurityAccessReviewPrincipal) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *SecurityAccessReviewPrincipal) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *SecurityAccessReviewPrincipal) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetHomeLocation

`func (o *SecurityAccessReviewPrincipal) GetHomeLocation() SecurityAccessReviewPrincipalLocation`

GetHomeLocation returns the HomeLocation field if non-nil, zero value otherwise.

### GetHomeLocationOk

`func (o *SecurityAccessReviewPrincipal) GetHomeLocationOk() (*SecurityAccessReviewPrincipalLocation, bool)`

GetHomeLocationOk returns a tuple with the HomeLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomeLocation

`func (o *SecurityAccessReviewPrincipal) SetHomeLocation(v SecurityAccessReviewPrincipalLocation)`

SetHomeLocation sets HomeLocation field to given value.

### HasHomeLocation

`func (o *SecurityAccessReviewPrincipal) HasHomeLocation() bool`

HasHomeLocation returns a boolean if a field has been set.

### GetLastLoginInfo

`func (o *SecurityAccessReviewPrincipal) GetLastLoginInfo() SecurityAccessReviewPrincipalLastLoginInfo`

GetLastLoginInfo returns the LastLoginInfo field if non-nil, zero value otherwise.

### GetLastLoginInfoOk

`func (o *SecurityAccessReviewPrincipal) GetLastLoginInfoOk() (*SecurityAccessReviewPrincipalLastLoginInfo, bool)`

GetLastLoginInfoOk returns a tuple with the LastLoginInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastLoginInfo

`func (o *SecurityAccessReviewPrincipal) SetLastLoginInfo(v SecurityAccessReviewPrincipalLastLoginInfo)`

SetLastLoginInfo sets LastLoginInfo field to given value.

### HasLastLoginInfo

`func (o *SecurityAccessReviewPrincipal) HasLastLoginInfo() bool`

HasLastLoginInfo returns a boolean if a field has been set.

### GetOktaAdminRoles

`func (o *SecurityAccessReviewPrincipal) GetOktaAdminRoles() []SecurityAccessReviewPrincipalOktaAdminRole`

GetOktaAdminRoles returns the OktaAdminRoles field if non-nil, zero value otherwise.

### GetOktaAdminRolesOk

`func (o *SecurityAccessReviewPrincipal) GetOktaAdminRolesOk() (*[]SecurityAccessReviewPrincipalOktaAdminRole, bool)`

GetOktaAdminRolesOk returns a tuple with the OktaAdminRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOktaAdminRoles

`func (o *SecurityAccessReviewPrincipal) SetOktaAdminRoles(v []SecurityAccessReviewPrincipalOktaAdminRole)`

SetOktaAdminRoles sets OktaAdminRoles field to given value.

### HasOktaAdminRoles

`func (o *SecurityAccessReviewPrincipal) HasOktaAdminRoles() bool`

HasOktaAdminRoles returns a boolean if a field has been set.

### GetLinks

`func (o *SecurityAccessReviewPrincipal) GetLinks() map[string]Link`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *SecurityAccessReviewPrincipal) GetLinksOk() (*map[string]Link, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *SecurityAccessReviewPrincipal) SetLinks(v map[string]Link)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *SecurityAccessReviewPrincipal) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


