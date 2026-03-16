# ClientCredentialPrincipal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** | The principal ID | 
**Type** | **string** | The type of principal | 

## Methods

### NewClientCredentialPrincipal

`func NewClientCredentialPrincipal(externalId string, type_ string, ) *ClientCredentialPrincipal`

NewClientCredentialPrincipal instantiates a new ClientCredentialPrincipal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCredentialPrincipalWithDefaults

`func NewClientCredentialPrincipalWithDefaults() *ClientCredentialPrincipal`

NewClientCredentialPrincipalWithDefaults instantiates a new ClientCredentialPrincipal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *ClientCredentialPrincipal) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ClientCredentialPrincipal) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ClientCredentialPrincipal) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetType

`func (o *ClientCredentialPrincipal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientCredentialPrincipal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientCredentialPrincipal) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


