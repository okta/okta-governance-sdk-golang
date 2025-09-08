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

// ValidRequesterSetting struct for ValidRequesterSetting
type ValidRequesterSetting struct {
	Type                 *RequesterSettingsType `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ValidRequesterSetting ValidRequesterSetting

// NewValidRequesterSetting instantiates a new ValidRequesterSetting object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidRequesterSetting() *ValidRequesterSetting {
	this := ValidRequesterSetting{}
	return &this
}

// NewValidRequesterSettingWithDefaults instantiates a new ValidRequesterSetting object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidRequesterSettingWithDefaults() *ValidRequesterSetting {
	this := ValidRequesterSetting{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ValidRequesterSetting) GetType() RequesterSettingsType {
	if o == nil || o.Type == nil {
		var ret RequesterSettingsType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ValidRequesterSetting) GetTypeOk() (*RequesterSettingsType, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ValidRequesterSetting) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given RequesterSettingsType and assigns it to the Type field.
func (o *ValidRequesterSetting) SetType(v RequesterSettingsType) {
	o.Type = &v
}

func (o ValidRequesterSetting) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ValidRequesterSetting) UnmarshalJSON(bytes []byte) (err error) {
	varValidRequesterSetting := _ValidRequesterSetting{}

	err = json.Unmarshal(bytes, &varValidRequesterSetting)
	if err == nil {
		*o = ValidRequesterSetting(varValidRequesterSetting)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableValidRequesterSetting struct {
	value *ValidRequesterSetting
	isSet bool
}

func (v NullableValidRequesterSetting) Get() *ValidRequesterSetting {
	return v.value
}

func (v *NullableValidRequesterSetting) Set(val *ValidRequesterSetting) {
	v.value = val
	v.isSet = true
}

func (v NullableValidRequesterSetting) IsSet() bool {
	return v.isSet
}

func (v *NullableValidRequesterSetting) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidRequesterSetting(val *ValidRequesterSetting) *NullableValidRequesterSetting {
	return &NullableValidRequesterSetting{value: val, isSet: true}
}

func (v NullableValidRequesterSetting) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidRequesterSetting) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
