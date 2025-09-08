# RiskAssessmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]RuleConflict**](RuleConflict.md) | List of conflicts for the requested resource | 
**Links** | [**RiskAssessmentLinks**](RiskAssessmentLinks.md) |  | 

## Methods

### NewRiskAssessmentResponse

`func NewRiskAssessmentResponse(data []RuleConflict, links RiskAssessmentLinks, ) *RiskAssessmentResponse`

NewRiskAssessmentResponse instantiates a new RiskAssessmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRiskAssessmentResponseWithDefaults

`func NewRiskAssessmentResponseWithDefaults() *RiskAssessmentResponse`

NewRiskAssessmentResponseWithDefaults instantiates a new RiskAssessmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *RiskAssessmentResponse) GetData() []RuleConflict`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *RiskAssessmentResponse) GetDataOk() (*[]RuleConflict, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *RiskAssessmentResponse) SetData(v []RuleConflict)`

SetData sets Data field to given value.


### GetLinks

`func (o *RiskAssessmentResponse) GetLinks() RiskAssessmentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *RiskAssessmentResponse) GetLinksOk() (*RiskAssessmentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *RiskAssessmentResponse) SetLinks(v RiskAssessmentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


