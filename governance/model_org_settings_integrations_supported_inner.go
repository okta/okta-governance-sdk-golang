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
)

// checks if the OrgSettingsIntegrationsSupportedInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgSettingsIntegrationsSupportedInner{}

// OrgSettingsIntegrationsSupportedInner struct for OrgSettingsIntegrationsSupportedInner
type OrgSettingsIntegrationsSupportedInner struct {
	Type                 *IntegrationType `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _OrgSettingsIntegrationsSupportedInner OrgSettingsIntegrationsSupportedInner

// NewOrgSettingsIntegrationsSupportedInner instantiates a new OrgSettingsIntegrationsSupportedInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgSettingsIntegrationsSupportedInner() *OrgSettingsIntegrationsSupportedInner {
	this := OrgSettingsIntegrationsSupportedInner{}
	return &this
}

// NewOrgSettingsIntegrationsSupportedInnerWithDefaults instantiates a new OrgSettingsIntegrationsSupportedInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgSettingsIntegrationsSupportedInnerWithDefaults() *OrgSettingsIntegrationsSupportedInner {
	this := OrgSettingsIntegrationsSupportedInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *OrgSettingsIntegrationsSupportedInner) GetType() IntegrationType {
	if o == nil || IsNil(o.Type) {
		var ret IntegrationType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgSettingsIntegrationsSupportedInner) GetTypeOk() (*IntegrationType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *OrgSettingsIntegrationsSupportedInner) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given IntegrationType and assigns it to the Type field.
func (o *OrgSettingsIntegrationsSupportedInner) SetType(v IntegrationType) {
	o.Type = &v
}

func (o OrgSettingsIntegrationsSupportedInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgSettingsIntegrationsSupportedInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgSettingsIntegrationsSupportedInner) UnmarshalJSON(data []byte) (err error) {
	varOrgSettingsIntegrationsSupportedInner := _OrgSettingsIntegrationsSupportedInner{}

	err = json.Unmarshal(data, &varOrgSettingsIntegrationsSupportedInner)

	if err != nil {
		return err
	}

	*o = OrgSettingsIntegrationsSupportedInner(varOrgSettingsIntegrationsSupportedInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgSettingsIntegrationsSupportedInner struct {
	value *OrgSettingsIntegrationsSupportedInner
	isSet bool
}

func (v NullableOrgSettingsIntegrationsSupportedInner) Get() *OrgSettingsIntegrationsSupportedInner {
	return v.value
}

func (v *NullableOrgSettingsIntegrationsSupportedInner) Set(val *OrgSettingsIntegrationsSupportedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgSettingsIntegrationsSupportedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgSettingsIntegrationsSupportedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgSettingsIntegrationsSupportedInner(val *OrgSettingsIntegrationsSupportedInner) *NullableOrgSettingsIntegrationsSupportedInner {
	return &NullableOrgSettingsIntegrationsSupportedInner{value: val, isSet: true}
}

func (v NullableOrgSettingsIntegrationsSupportedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgSettingsIntegrationsSupportedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
