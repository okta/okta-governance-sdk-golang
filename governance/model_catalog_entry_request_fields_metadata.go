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

// CatalogEntryRequestFieldsMetadata struct for CatalogEntryRequestFieldsMetadata
type CatalogEntryRequestFieldsMetadata struct {
	RiskAssessment       *RiskAssessment `json:"riskAssessment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CatalogEntryRequestFieldsMetadata CatalogEntryRequestFieldsMetadata

// NewCatalogEntryRequestFieldsMetadata instantiates a new CatalogEntryRequestFieldsMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogEntryRequestFieldsMetadata() *CatalogEntryRequestFieldsMetadata {
	this := CatalogEntryRequestFieldsMetadata{}
	return &this
}

// NewCatalogEntryRequestFieldsMetadataWithDefaults instantiates a new CatalogEntryRequestFieldsMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogEntryRequestFieldsMetadataWithDefaults() *CatalogEntryRequestFieldsMetadata {
	this := CatalogEntryRequestFieldsMetadata{}
	return &this
}

// GetRiskAssessment returns the RiskAssessment field value if set, zero value otherwise.
func (o *CatalogEntryRequestFieldsMetadata) GetRiskAssessment() RiskAssessment {
	if o == nil || o.RiskAssessment == nil {
		var ret RiskAssessment
		return ret
	}
	return *o.RiskAssessment
}

// GetRiskAssessmentOk returns a tuple with the RiskAssessment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogEntryRequestFieldsMetadata) GetRiskAssessmentOk() (*RiskAssessment, bool) {
	if o == nil || o.RiskAssessment == nil {
		return nil, false
	}
	return o.RiskAssessment, true
}

// HasRiskAssessment returns a boolean if a field has been set.
func (o *CatalogEntryRequestFieldsMetadata) HasRiskAssessment() bool {
	if o != nil && o.RiskAssessment != nil {
		return true
	}

	return false
}

// SetRiskAssessment gets a reference to the given RiskAssessment and assigns it to the RiskAssessment field.
func (o *CatalogEntryRequestFieldsMetadata) SetRiskAssessment(v RiskAssessment) {
	o.RiskAssessment = &v
}

func (o CatalogEntryRequestFieldsMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RiskAssessment != nil {
		toSerialize["riskAssessment"] = o.RiskAssessment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CatalogEntryRequestFieldsMetadata) UnmarshalJSON(bytes []byte) (err error) {
	varCatalogEntryRequestFieldsMetadata := _CatalogEntryRequestFieldsMetadata{}

	err = json.Unmarshal(bytes, &varCatalogEntryRequestFieldsMetadata)
	if err == nil {
		*o = CatalogEntryRequestFieldsMetadata(varCatalogEntryRequestFieldsMetadata)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "riskAssessment")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableCatalogEntryRequestFieldsMetadata struct {
	value *CatalogEntryRequestFieldsMetadata
	isSet bool
}

func (v NullableCatalogEntryRequestFieldsMetadata) Get() *CatalogEntryRequestFieldsMetadata {
	return v.value
}

func (v *NullableCatalogEntryRequestFieldsMetadata) Set(val *CatalogEntryRequestFieldsMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogEntryRequestFieldsMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogEntryRequestFieldsMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogEntryRequestFieldsMetadata(val *CatalogEntryRequestFieldsMetadata) *NullableCatalogEntryRequestFieldsMetadata {
	return &NullableCatalogEntryRequestFieldsMetadata{value: val, isSet: true}
}

func (v NullableCatalogEntryRequestFieldsMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogEntryRequestFieldsMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
