/*
Okta Governance API

Allows customers to easily access the Okta API

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 3.2.0
Contact: devex-public@okta.com
*/

package governance

import (
	"encoding/json"
)

// RiskAssessment A risk assessment indicates whether request submission is allowed or restricted and contains the risk rules that lead to possible conflicts for the requested resource.
type RiskAssessment struct {
	RequestSubmissionType RequestSubmissionType `json:"requestSubmissionType"`
	RiskRules             []RiskRule            `json:"riskRules"`
	AdditionalProperties  map[string]interface{}
}

type _RiskAssessment RiskAssessment

// NewRiskAssessment instantiates a new RiskAssessment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskAssessment(requestSubmissionType RequestSubmissionType, riskRules []RiskRule) *RiskAssessment {
	this := RiskAssessment{}
	this.RequestSubmissionType = requestSubmissionType
	this.RiskRules = riskRules
	return &this
}

// NewRiskAssessmentWithDefaults instantiates a new RiskAssessment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskAssessmentWithDefaults() *RiskAssessment {
	this := RiskAssessment{}
	return &this
}

// GetRequestSubmissionType returns the RequestSubmissionType field value
func (o *RiskAssessment) GetRequestSubmissionType() RequestSubmissionType {
	if o == nil {
		var ret RequestSubmissionType
		return ret
	}

	return o.RequestSubmissionType
}

// GetRequestSubmissionTypeOk returns a tuple with the RequestSubmissionType field value
// and a boolean to check if the value has been set.
func (o *RiskAssessment) GetRequestSubmissionTypeOk() (*RequestSubmissionType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestSubmissionType, true
}

// SetRequestSubmissionType sets field value
func (o *RiskAssessment) SetRequestSubmissionType(v RequestSubmissionType) {
	o.RequestSubmissionType = v
}

// GetRiskRules returns the RiskRules field value
func (o *RiskAssessment) GetRiskRules() []RiskRule {
	if o == nil {
		var ret []RiskRule
		return ret
	}

	return o.RiskRules
}

// GetRiskRulesOk returns a tuple with the RiskRules field value
// and a boolean to check if the value has been set.
func (o *RiskAssessment) GetRiskRulesOk() ([]RiskRule, bool) {
	if o == nil {
		return nil, false
	}
	return o.RiskRules, true
}

// SetRiskRules sets field value
func (o *RiskAssessment) SetRiskRules(v []RiskRule) {
	o.RiskRules = v
}

func (o RiskAssessment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["requestSubmissionType"] = o.RequestSubmissionType
	}
	if true {
		toSerialize["riskRules"] = o.RiskRules
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RiskAssessment) UnmarshalJSON(bytes []byte) (err error) {
	varRiskAssessment := _RiskAssessment{}

	err = json.Unmarshal(bytes, &varRiskAssessment)
	if err == nil {
		*o = RiskAssessment(varRiskAssessment)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "requestSubmissionType")
		delete(additionalProperties, "riskRules")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRiskAssessment struct {
	value *RiskAssessment
	isSet bool
}

func (v NullableRiskAssessment) Get() *RiskAssessment {
	return v.value
}

func (v *NullableRiskAssessment) Set(val *RiskAssessment) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskAssessment) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskAssessment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskAssessment(val *RiskAssessment) *NullableRiskAssessment {
	return &NullableRiskAssessment{value: val, isSet: true}
}

func (v NullableRiskAssessment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskAssessment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
