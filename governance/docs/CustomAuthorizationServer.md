# CustomAuthorizationServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Orn** | **string** | The [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) of the authorization server | 
**Name** | **string** | Display name of the authorization server | 
**IssuerUrl** | **string** | Issuer URL for the authorization server | 

## Methods

### NewCustomAuthorizationServer

`func NewCustomAuthorizationServer(orn string, name string, issuerUrl string, ) *CustomAuthorizationServer`

NewCustomAuthorizationServer instantiates a new CustomAuthorizationServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomAuthorizationServerWithDefaults

`func NewCustomAuthorizationServerWithDefaults() *CustomAuthorizationServer`

NewCustomAuthorizationServerWithDefaults instantiates a new CustomAuthorizationServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrn

`func (o *CustomAuthorizationServer) GetOrn() string`

GetOrn returns the Orn field if non-nil, zero value otherwise.

### GetOrnOk

`func (o *CustomAuthorizationServer) GetOrnOk() (*string, bool)`

GetOrnOk returns a tuple with the Orn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrn

`func (o *CustomAuthorizationServer) SetOrn(v string)`

SetOrn sets Orn field to given value.


### GetName

`func (o *CustomAuthorizationServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomAuthorizationServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomAuthorizationServer) SetName(v string)`

SetName sets Name field to given value.


### GetIssuerUrl

`func (o *CustomAuthorizationServer) GetIssuerUrl() string`

GetIssuerUrl returns the IssuerUrl field if non-nil, zero value otherwise.

### GetIssuerUrlOk

`func (o *CustomAuthorizationServer) GetIssuerUrlOk() (*string, bool)`

GetIssuerUrlOk returns a tuple with the IssuerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuerUrl

`func (o *CustomAuthorizationServer) SetIssuerUrl(v string)`

SetIssuerUrl sets IssuerUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


