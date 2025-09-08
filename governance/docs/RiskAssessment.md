# RiskAssessment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestSubmissionType** | [**RequestSubmissionType**](RequestSubmissionType.md) |  | 
**RiskRules** | [**[]RiskRule**](RiskRule.md) |  | 

## Methods

### NewRiskAssessment

`func NewRiskAssessment(requestSubmissionType RequestSubmissionType, riskRules []RiskRule, ) *RiskAssessment`

NewRiskAssessment instantiates a new RiskAssessment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAssessmentWithDefaults

`func NewRiskAssessmentWithDefaults() *RiskAssessment`

NewRiskAssessmentWithDefaults instantiates a new RiskAssessment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestSubmissionType

`func (o *RiskAssessment) GetRequestSubmissionType() RequestSubmissionType`

GetRequestSubmissionType returns the RequestSubmissionType field if non-nil, zero value otherwise.

### GetRequestSubmissionTypeOk

`func (o *RiskAssessment) GetRequestSubmissionTypeOk() (*RequestSubmissionType, bool)`

GetRequestSubmissionTypeOk returns a tuple with the RequestSubmissionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestSubmissionType

`func (o *RiskAssessment) SetRequestSubmissionType(v RequestSubmissionType)`

SetRequestSubmissionType sets RequestSubmissionType field to given value.


### GetRiskRules

`func (o *RiskAssessment) GetRiskRules() []RiskRule`

GetRiskRules returns the RiskRules field if non-nil, zero value otherwise.

### GetRiskRulesOk

`func (o *RiskAssessment) GetRiskRulesOk() (*[]RiskRule, bool)`

GetRiskRulesOk returns a tuple with the RiskRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskRules

`func (o *RiskAssessment) SetRiskRules(v []RiskRule)`

SetRiskRules sets RiskRules field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


