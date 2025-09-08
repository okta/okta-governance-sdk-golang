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

// AutoRemediationSettings When a group is selected to be automatically remediated, if a user is assigned to an app through this group,  the user will be automatically removed from the group if their access to the app is revoked in the Access Certification Campaign. Group is currently the only supported type of resources that can be automatically remediated.
type AutoRemediationSettings struct {
	// An array of resources to be automatically remediated
	IncludeOnly []AutoRemediationSettingsIncludeOnlyInner `json:"includeOnly,omitempty"`
	// If `includeAllIndirectAssignments` is set to `true`, the user's access to all groups that can assign the user to the application(s) will be removed during remediation. Only app assignments through groups can be automatically remediated. **Note:** You can only specify either  `includeAllIndirectAssignments` or `includeOnly`.
	IncludeAllIndirectAssignments *bool `json:"includeAllIndirectAssignments,omitempty"`
	AdditionalProperties          map[string]interface{}
}

type _AutoRemediationSettings AutoRemediationSettings

// NewAutoRemediationSettings instantiates a new AutoRemediationSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoRemediationSettings() *AutoRemediationSettings {
	this := AutoRemediationSettings{}
	return &this
}

// NewAutoRemediationSettingsWithDefaults instantiates a new AutoRemediationSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoRemediationSettingsWithDefaults() *AutoRemediationSettings {
	this := AutoRemediationSettings{}
	return &this
}

// GetIncludeOnly returns the IncludeOnly field value if set, zero value otherwise.
func (o *AutoRemediationSettings) GetIncludeOnly() []AutoRemediationSettingsIncludeOnlyInner {
	if o == nil || o.IncludeOnly == nil {
		var ret []AutoRemediationSettingsIncludeOnlyInner
		return ret
	}
	return o.IncludeOnly
}

// GetIncludeOnlyOk returns a tuple with the IncludeOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoRemediationSettings) GetIncludeOnlyOk() ([]AutoRemediationSettingsIncludeOnlyInner, bool) {
	if o == nil || o.IncludeOnly == nil {
		return nil, false
	}
	return o.IncludeOnly, true
}

// HasIncludeOnly returns a boolean if a field has been set.
func (o *AutoRemediationSettings) HasIncludeOnly() bool {
	if o != nil && o.IncludeOnly != nil {
		return true
	}

	return false
}

// SetIncludeOnly gets a reference to the given []AutoRemediationSettingsIncludeOnlyInner and assigns it to the IncludeOnly field.
func (o *AutoRemediationSettings) SetIncludeOnly(v []AutoRemediationSettingsIncludeOnlyInner) {
	o.IncludeOnly = v
}

// GetIncludeAllIndirectAssignments returns the IncludeAllIndirectAssignments field value if set, zero value otherwise.
func (o *AutoRemediationSettings) GetIncludeAllIndirectAssignments() bool {
	if o == nil || o.IncludeAllIndirectAssignments == nil {
		var ret bool
		return ret
	}
	return *o.IncludeAllIndirectAssignments
}

// GetIncludeAllIndirectAssignmentsOk returns a tuple with the IncludeAllIndirectAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoRemediationSettings) GetIncludeAllIndirectAssignmentsOk() (*bool, bool) {
	if o == nil || o.IncludeAllIndirectAssignments == nil {
		return nil, false
	}
	return o.IncludeAllIndirectAssignments, true
}

// HasIncludeAllIndirectAssignments returns a boolean if a field has been set.
func (o *AutoRemediationSettings) HasIncludeAllIndirectAssignments() bool {
	if o != nil && o.IncludeAllIndirectAssignments != nil {
		return true
	}

	return false
}

// SetIncludeAllIndirectAssignments gets a reference to the given bool and assigns it to the IncludeAllIndirectAssignments field.
func (o *AutoRemediationSettings) SetIncludeAllIndirectAssignments(v bool) {
	o.IncludeAllIndirectAssignments = &v
}

func (o AutoRemediationSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IncludeOnly != nil {
		toSerialize["includeOnly"] = o.IncludeOnly
	}
	if o.IncludeAllIndirectAssignments != nil {
		toSerialize["includeAllIndirectAssignments"] = o.IncludeAllIndirectAssignments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AutoRemediationSettings) UnmarshalJSON(bytes []byte) (err error) {
	varAutoRemediationSettings := _AutoRemediationSettings{}

	err = json.Unmarshal(bytes, &varAutoRemediationSettings)
	if err == nil {
		*o = AutoRemediationSettings(varAutoRemediationSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "includeOnly")
		delete(additionalProperties, "includeAllIndirectAssignments")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableAutoRemediationSettings struct {
	value *AutoRemediationSettings
	isSet bool
}

func (v NullableAutoRemediationSettings) Get() *AutoRemediationSettings {
	return v.value
}

func (v *NullableAutoRemediationSettings) Set(val *AutoRemediationSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoRemediationSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoRemediationSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoRemediationSettings(val *AutoRemediationSettings) *NullableAutoRemediationSettings {
	return &NullableAutoRemediationSettings{value: val, isSet: true}
}

func (v NullableAutoRemediationSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoRemediationSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
