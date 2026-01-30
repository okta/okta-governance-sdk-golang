# ResourceOwnerPrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | A unique identifier for the principal in Okta | 
**Type** | **string** | The principal type. This value is the &#x60;{objectType}&#x60; attribute from the [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) string.  Examples: &#x60;groups&#x60; or &#x60;users&#x60; | 
**Orn** | **string** | The Okta user or group &#x60;id&#x60; in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format. The resource can be an user id, or a group id. See [Supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | 
**Profile** | [**ExternalPrincipalProfile**](ExternalPrincipalProfile.md) |  | 

## Methods

### NewResourceOwnerPrincipal

`func NewResourceOwnerPrincipal(id string, type_ string, orn string, profile ExternalPrincipalProfile, ) *ResourceOwnerPrincipal`

NewResourceOwnerPrincipal instantiates a new ResourceOwnerPrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceOwnerPrincipalWithDefaults

`func NewResourceOwnerPrincipalWithDefaults() *ResourceOwnerPrincipal`

NewResourceOwnerPrincipalWithDefaults instantiates a new ResourceOwnerPrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResourceOwnerPrincipal) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceOwnerPrincipal) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceOwnerPrincipal) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *ResourceOwnerPrincipal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ResourceOwnerPrincipal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ResourceOwnerPrincipal) SetType(v string)`

SetType sets Type field to given value.


### GetOrn

`func (o *ResourceOwnerPrincipal) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *ResourceOwnerPrincipal) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *ResourceOwnerPrincipal) SetOrn(v string)`

SetOrn sets Orn field to given value.


### GetProfile

`func (o *ResourceOwnerPrincipal) GetProfile() ExternalPrincipalProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ResourceOwnerPrincipal) GetProfileOk() (*ExternalPrincipalProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ResourceOwnerPrincipal) SetProfile(v ExternalPrincipalProfile)`

SetProfile sets Profile field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


