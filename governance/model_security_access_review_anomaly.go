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

// SecurityAccessReviewAnomaly struct for SecurityAccessReviewAnomaly
type SecurityAccessReviewAnomaly struct {
	Type     SecurityAccessReviewAnomalyType        `json:"type"`
	Severity SecurityAccessReviewAccessItemSeverity `json:"severity"`
	Subtext  ServerMessage                          `json:"subtext"`
	// A list of separation of duties (SOD) conflicts caused by a user's access to an entitlement
	SodConflicts         []SecurityAccessReviewSodConflict `json:"sodConflicts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewAnomaly SecurityAccessReviewAnomaly

// NewSecurityAccessReviewAnomaly instantiates a new SecurityAccessReviewAnomaly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewAnomaly(type_ SecurityAccessReviewAnomalyType, severity SecurityAccessReviewAccessItemSeverity, subtext ServerMessage) *SecurityAccessReviewAnomaly {
	this := SecurityAccessReviewAnomaly{}
	this.Type = type_
	this.Severity = severity
	this.Subtext = subtext
	return &this
}

// NewSecurityAccessReviewAnomalyWithDefaults instantiates a new SecurityAccessReviewAnomaly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewAnomalyWithDefaults() *SecurityAccessReviewAnomaly {
	this := SecurityAccessReviewAnomaly{}
	return &this
}

// GetType returns the Type field value
func (o *SecurityAccessReviewAnomaly) GetType() SecurityAccessReviewAnomalyType {
	if o == nil {
		var ret SecurityAccessReviewAnomalyType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewAnomaly) GetTypeOk() (*SecurityAccessReviewAnomalyType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SecurityAccessReviewAnomaly) SetType(v SecurityAccessReviewAnomalyType) {
	o.Type = v
}

// GetSeverity returns the Severity field value
func (o *SecurityAccessReviewAnomaly) GetSeverity() SecurityAccessReviewAccessItemSeverity {
	if o == nil {
		var ret SecurityAccessReviewAccessItemSeverity
		return ret
	}

	return o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewAnomaly) GetSeverityOk() (*SecurityAccessReviewAccessItemSeverity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Severity, true
}

// SetSeverity sets field value
func (o *SecurityAccessReviewAnomaly) SetSeverity(v SecurityAccessReviewAccessItemSeverity) {
	o.Severity = v
}

// GetSubtext returns the Subtext field value
func (o *SecurityAccessReviewAnomaly) GetSubtext() ServerMessage {
	if o == nil {
		var ret ServerMessage
		return ret
	}

	return o.Subtext
}

// GetSubtextOk returns a tuple with the Subtext field value
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewAnomaly) GetSubtextOk() (*ServerMessage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subtext, true
}

// SetSubtext sets field value
func (o *SecurityAccessReviewAnomaly) SetSubtext(v ServerMessage) {
	o.Subtext = v
}

// GetSodConflicts returns the SodConflicts field value if set, zero value otherwise.
func (o *SecurityAccessReviewAnomaly) GetSodConflicts() []SecurityAccessReviewSodConflict {
	if o == nil || o.SodConflicts == nil {
		var ret []SecurityAccessReviewSodConflict
		return ret
	}
	return o.SodConflicts
}

// GetSodConflictsOk returns a tuple with the SodConflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewAnomaly) GetSodConflictsOk() ([]SecurityAccessReviewSodConflict, bool) {
	if o == nil || o.SodConflicts == nil {
		return nil, false
	}
	return o.SodConflicts, true
}

// HasSodConflicts returns a boolean if a field has been set.
func (o *SecurityAccessReviewAnomaly) HasSodConflicts() bool {
	if o != nil && o.SodConflicts != nil {
		return true
	}

	return false
}

// SetSodConflicts gets a reference to the given []SecurityAccessReviewSodConflict and assigns it to the SodConflicts field.
func (o *SecurityAccessReviewAnomaly) SetSodConflicts(v []SecurityAccessReviewSodConflict) {
	o.SodConflicts = v
}

func (o SecurityAccessReviewAnomaly) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["severity"] = o.Severity
	}
	if true {
		toSerialize["subtext"] = o.Subtext
	}
	if o.SodConflicts != nil {
		toSerialize["sodConflicts"] = o.SodConflicts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityAccessReviewAnomaly) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityAccessReviewAnomaly := _SecurityAccessReviewAnomaly{}

	err = json.Unmarshal(bytes, &varSecurityAccessReviewAnomaly)
	if err == nil {
		*o = SecurityAccessReviewAnomaly(varSecurityAccessReviewAnomaly)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "severity")
		delete(additionalProperties, "subtext")
		delete(additionalProperties, "sodConflicts")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityAccessReviewAnomaly struct {
	value *SecurityAccessReviewAnomaly
	isSet bool
}

func (v NullableSecurityAccessReviewAnomaly) Get() *SecurityAccessReviewAnomaly {
	return v.value
}

func (v *NullableSecurityAccessReviewAnomaly) Set(val *SecurityAccessReviewAnomaly) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewAnomaly) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewAnomaly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewAnomaly(val *SecurityAccessReviewAnomaly) *NullableSecurityAccessReviewAnomaly {
	return &NullableSecurityAccessReviewAnomaly{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewAnomaly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewAnomaly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
