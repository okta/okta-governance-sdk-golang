/*
Okta Governance API

Allows customers to easily access the Okta API

Copyright 2025 - Present Okta, Inc.

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
	"fmt"
)

// checks if the RiskAssessmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RiskAssessmentResponse{}

// RiskAssessmentResponse struct for RiskAssessmentResponse
type RiskAssessmentResponse struct {
	// List of conflicts for the requested resource
	Data                 []RuleConflict      `json:"data"`
	Links                RiskAssessmentLinks `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _RiskAssessmentResponse RiskAssessmentResponse

// NewRiskAssessmentResponse instantiates a new RiskAssessmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRiskAssessmentResponse(data []RuleConflict, links RiskAssessmentLinks) *RiskAssessmentResponse {
	this := RiskAssessmentResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewRiskAssessmentResponseWithDefaults instantiates a new RiskAssessmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRiskAssessmentResponseWithDefaults() *RiskAssessmentResponse {
	this := RiskAssessmentResponse{}
	return &this
}

// GetData returns the Data field value
func (o *RiskAssessmentResponse) GetData() []RuleConflict {
	if o == nil {
		var ret []RuleConflict
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *RiskAssessmentResponse) GetDataOk() ([]RuleConflict, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *RiskAssessmentResponse) SetData(v []RuleConflict) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *RiskAssessmentResponse) GetLinks() RiskAssessmentLinks {
	if o == nil {
		var ret RiskAssessmentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *RiskAssessmentResponse) GetLinksOk() (*RiskAssessmentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *RiskAssessmentResponse) SetLinks(v RiskAssessmentLinks) {
	o.Links = v
}

func (o RiskAssessmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RiskAssessmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RiskAssessmentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
		"_links",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varRiskAssessmentResponse := _RiskAssessmentResponse{}

	err = json.Unmarshal(data, &varRiskAssessmentResponse)

	if err != nil {
		return err
	}

	*o = RiskAssessmentResponse(varRiskAssessmentResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "data")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRiskAssessmentResponse struct {
	value *RiskAssessmentResponse
	isSet bool
}

func (v NullableRiskAssessmentResponse) Get() *RiskAssessmentResponse {
	return v.value
}

func (v *NullableRiskAssessmentResponse) Set(val *RiskAssessmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRiskAssessmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRiskAssessmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRiskAssessmentResponse(val *RiskAssessmentResponse) *NullableRiskAssessmentResponse {
	return &NullableRiskAssessmentResponse{value: val, isSet: true}
}

func (v NullableRiskAssessmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRiskAssessmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
