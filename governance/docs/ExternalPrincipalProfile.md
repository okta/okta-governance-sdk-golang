# ExternalPrincipalProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Okta user &#x60;id&#x60; or Okta group &#x60;id&#x60; | 
**Name** | **string** | User name or Group Name | 
**Email** | Pointer to **string** | Email of the resource owner, if applicable. | [optional] 
**Logo** | Pointer to [**[]Link**](Link.md) | List of logo resources | [optional] 
**Links** | [**LinkSelf**](LinkSelf.md) |  | 

## Methods

### NewExternalPrincipalProfile

`func NewExternalPrincipalProfile(id string, name string, links LinkSelf, ) *ExternalPrincipalProfile`

NewExternalPrincipalProfile instantiates a new ExternalPrincipalProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalPrincipalProfileWithDefaults

`func NewExternalPrincipalProfileWithDefaults() *ExternalPrincipalProfile`

NewExternalPrincipalProfileWithDefaults instantiates a new ExternalPrincipalProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalPrincipalProfile) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalPrincipalProfile) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalPrincipalProfile) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ExternalPrincipalProfile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExternalPrincipalProfile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExternalPrincipalProfile) SetName(v string)`

SetName sets Name field to given value.


### GetEmail

`func (o *ExternalPrincipalProfile) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ExternalPrincipalProfile) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ExternalPrincipalProfile) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ExternalPrincipalProfile) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLogo

`func (o *ExternalPrincipalProfile) GetLogo() []Link`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *ExternalPrincipalProfile) GetLogoOk() (*[]Link, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *ExternalPrincipalProfile) SetLogo(v []Link)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *ExternalPrincipalProfile) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetLinks

`func (o *ExternalPrincipalProfile) GetLinks() LinkSelf`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ExternalPrincipalProfile) GetLinksOk() (*LinkSelf, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ExternalPrincipalProfile) SetLinks(v LinkSelf)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


