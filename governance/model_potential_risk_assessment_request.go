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

// checks if the PotentialRiskAssessmentRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PotentialRiskAssessmentRequest{}

// PotentialRiskAssessmentRequest struct for PotentialRiskAssessmentRequest
type PotentialRiskAssessmentRequest struct {
	// The Okta user `id` in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format.  See [Supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	PrincipalOrn string `json:"principalOrn"`
	// The `id` of the resource in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format. The resource can be an app, a collection, or a bundle.
	ResourceOrn          string `json:"resourceOrn"`
	AdditionalProperties map[string]interface{}
}

type _PotentialRiskAssessmentRequest PotentialRiskAssessmentRequest

// NewPotentialRiskAssessmentRequest instantiates a new PotentialRiskAssessmentRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPotentialRiskAssessmentRequest(principalOrn string, resourceOrn string) *PotentialRiskAssessmentRequest {
	this := PotentialRiskAssessmentRequest{}
	this.PrincipalOrn = principalOrn
	this.ResourceOrn = resourceOrn
	return &this
}

// NewPotentialRiskAssessmentRequestWithDefaults instantiates a new PotentialRiskAssessmentRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPotentialRiskAssessmentRequestWithDefaults() *PotentialRiskAssessmentRequest {
	this := PotentialRiskAssessmentRequest{}
	return &this
}

// GetPrincipalOrn returns the PrincipalOrn field value
func (o *PotentialRiskAssessmentRequest) GetPrincipalOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrincipalOrn
}

// GetPrincipalOrnOk returns a tuple with the PrincipalOrn field value
// and a boolean to check if the value has been set.
func (o *PotentialRiskAssessmentRequest) GetPrincipalOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalOrn, true
}

// SetPrincipalOrn sets field value
func (o *PotentialRiskAssessmentRequest) SetPrincipalOrn(v string) {
	o.PrincipalOrn = v
}

// GetResourceOrn returns the ResourceOrn field value
func (o *PotentialRiskAssessmentRequest) GetResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceOrn
}

// GetResourceOrnOk returns a tuple with the ResourceOrn field value
// and a boolean to check if the value has been set.
func (o *PotentialRiskAssessmentRequest) GetResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceOrn, true
}

// SetResourceOrn sets field value
func (o *PotentialRiskAssessmentRequest) SetResourceOrn(v string) {
	o.ResourceOrn = v
}

func (o PotentialRiskAssessmentRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PotentialRiskAssessmentRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["principalOrn"] = o.PrincipalOrn
	toSerialize["resourceOrn"] = o.ResourceOrn

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PotentialRiskAssessmentRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"principalOrn",
		"resourceOrn",
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

	varPotentialRiskAssessmentRequest := _PotentialRiskAssessmentRequest{}

	err = json.Unmarshal(data, &varPotentialRiskAssessmentRequest)

	if err != nil {
		return err
	}

	*o = PotentialRiskAssessmentRequest(varPotentialRiskAssessmentRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "principalOrn")
		delete(additionalProperties, "resourceOrn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePotentialRiskAssessmentRequest struct {
	value *PotentialRiskAssessmentRequest
	isSet bool
}

func (v NullablePotentialRiskAssessmentRequest) Get() *PotentialRiskAssessmentRequest {
	return v.value
}

func (v *NullablePotentialRiskAssessmentRequest) Set(val *PotentialRiskAssessmentRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePotentialRiskAssessmentRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePotentialRiskAssessmentRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePotentialRiskAssessmentRequest(val *PotentialRiskAssessmentRequest) *NullablePotentialRiskAssessmentRequest {
	return &NullablePotentialRiskAssessmentRequest{value: val, isSet: true}
}

func (v NullablePotentialRiskAssessmentRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePotentialRiskAssessmentRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
