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

// checks if the AutoRemediationSettingsIncludeOnlyInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AutoRemediationSettingsIncludeOnlyInner{}

// AutoRemediationSettingsIncludeOnlyInner Represents a list of resources that will be automatically remediated
type AutoRemediationSettingsIncludeOnlyInner struct {
	ResourceType *AutoRemediationResourceType `json:"resourceType,omitempty"`
	// The resource ID of the target resource When `type = GROUP`, it will point to the group ID.
	ResourceId           *string `json:"resourceId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AutoRemediationSettingsIncludeOnlyInner AutoRemediationSettingsIncludeOnlyInner

// NewAutoRemediationSettingsIncludeOnlyInner instantiates a new AutoRemediationSettingsIncludeOnlyInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAutoRemediationSettingsIncludeOnlyInner() *AutoRemediationSettingsIncludeOnlyInner {
	this := AutoRemediationSettingsIncludeOnlyInner{}
	return &this
}

// NewAutoRemediationSettingsIncludeOnlyInnerWithDefaults instantiates a new AutoRemediationSettingsIncludeOnlyInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAutoRemediationSettingsIncludeOnlyInnerWithDefaults() *AutoRemediationSettingsIncludeOnlyInner {
	this := AutoRemediationSettingsIncludeOnlyInner{}
	return &this
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *AutoRemediationSettingsIncludeOnlyInner) GetResourceType() AutoRemediationResourceType {
	if o == nil || IsNil(o.ResourceType) {
		var ret AutoRemediationResourceType
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoRemediationSettingsIncludeOnlyInner) GetResourceTypeOk() (*AutoRemediationResourceType, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *AutoRemediationSettingsIncludeOnlyInner) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given AutoRemediationResourceType and assigns it to the ResourceType field.
func (o *AutoRemediationSettingsIncludeOnlyInner) SetResourceType(v AutoRemediationResourceType) {
	o.ResourceType = &v
}

// GetResourceId returns the ResourceId field value if set, zero value otherwise.
func (o *AutoRemediationSettingsIncludeOnlyInner) GetResourceId() string {
	if o == nil || IsNil(o.ResourceId) {
		var ret string
		return ret
	}
	return *o.ResourceId
}

// GetResourceIdOk returns a tuple with the ResourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AutoRemediationSettingsIncludeOnlyInner) GetResourceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceId) {
		return nil, false
	}
	return o.ResourceId, true
}

// HasResourceId returns a boolean if a field has been set.
func (o *AutoRemediationSettingsIncludeOnlyInner) HasResourceId() bool {
	if o != nil && !IsNil(o.ResourceId) {
		return true
	}

	return false
}

// SetResourceId gets a reference to the given string and assigns it to the ResourceId field.
func (o *AutoRemediationSettingsIncludeOnlyInner) SetResourceId(v string) {
	o.ResourceId = &v
}

func (o AutoRemediationSettingsIncludeOnlyInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AutoRemediationSettingsIncludeOnlyInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}
	if !IsNil(o.ResourceId) {
		toSerialize["resourceId"] = o.ResourceId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AutoRemediationSettingsIncludeOnlyInner) UnmarshalJSON(data []byte) (err error) {
	varAutoRemediationSettingsIncludeOnlyInner := _AutoRemediationSettingsIncludeOnlyInner{}

	err = json.Unmarshal(data, &varAutoRemediationSettingsIncludeOnlyInner)

	if err != nil {
		return err
	}

	*o = AutoRemediationSettingsIncludeOnlyInner(varAutoRemediationSettingsIncludeOnlyInner)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "resourceType")
		delete(additionalProperties, "resourceId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAutoRemediationSettingsIncludeOnlyInner struct {
	value *AutoRemediationSettingsIncludeOnlyInner
	isSet bool
}

func (v NullableAutoRemediationSettingsIncludeOnlyInner) Get() *AutoRemediationSettingsIncludeOnlyInner {
	return v.value
}

func (v *NullableAutoRemediationSettingsIncludeOnlyInner) Set(val *AutoRemediationSettingsIncludeOnlyInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAutoRemediationSettingsIncludeOnlyInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAutoRemediationSettingsIncludeOnlyInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAutoRemediationSettingsIncludeOnlyInner(val *AutoRemediationSettingsIncludeOnlyInner) *NullableAutoRemediationSettingsIncludeOnlyInner {
	return &NullableAutoRemediationSettingsIncludeOnlyInner{value: val, isSet: true}
}

func (v NullableAutoRemediationSettingsIncludeOnlyInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAutoRemediationSettingsIncludeOnlyInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
