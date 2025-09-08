# PotentialRiskAssessmentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrincipalOrn** | **string** | The Okta user &#x60;id&#x60; in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format.  See [Supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources). | 
**ResourceOrn** | **string** | The &#x60;id&#x60; of the resource in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format. The resource can be an app, a collection, or a bundle. | 

## Methods

### NewPotentialRiskAssessmentRequest

`func NewPotentialRiskAssessmentRequest(principalOrn string, resourceOrn string, ) *PotentialRiskAssessmentRequest`

NewPotentialRiskAssessmentRequest instantiates a new PotentialRiskAssessmentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPotentialRiskAssessmentRequestWithDefaults

`func NewPotentialRiskAssessmentRequestWithDefaults() *PotentialRiskAssessmentRequest`

NewPotentialRiskAssessmentRequestWithDefaults instantiates a new PotentialRiskAssessmentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipalOrn

`func (o *PotentialRiskAssessmentRequest) GetPrincipalOrn() string`

GetPrincipalOrn returns the PrincipalOrn field if non-nil, zero value otherwise.

### GetPrincipalOrnOk

`func (o *PotentialRiskAssessmentRequest) GetPrincipalOrnOk() (*string, bool)`

GetPrincipalOrnOk returns a tuple with the PrincipalOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipalOrn

`func (o *PotentialRiskAssessmentRequest) SetPrincipalOrn(v string)`

SetPrincipalOrn sets PrincipalOrn field to given value.


### GetResourceOrn

`func (o *PotentialRiskAssessmentRequest) GetResourceOrn() string`

GetResourceOrn returns the ResourceOrn field if non-nil, zero value otherwise.

### GetResourceOrnOk

`func (o *PotentialRiskAssessmentRequest) GetResourceOrnOk() (*string, bool)`

GetResourceOrnOk returns a tuple with the ResourceOrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceOrn

`func (o *PotentialRiskAssessmentRequest) SetResourceOrn(v string)`

SetResourceOrn sets ResourceOrn field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


