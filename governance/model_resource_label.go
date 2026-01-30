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

// checks if the ResourceLabel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceLabel{}

// ResourceLabel struct for ResourceLabel
type ResourceLabel struct {
	// The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	Orn     *string                  `json:"orn,omitempty"`
	Profile *ExternalResourceProfile `json:"profile,omitempty"`
	// List of assigned labels.
	Labels               []Label  `json:"labels,omitempty"`
	Links                LinkSelf `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _ResourceLabel ResourceLabel

// NewResourceLabel instantiates a new ResourceLabel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceLabel(links LinkSelf) *ResourceLabel {
	this := ResourceLabel{}
	this.Links = links
	return &this
}

// NewResourceLabelWithDefaults instantiates a new ResourceLabel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceLabelWithDefaults() *ResourceLabel {
	this := ResourceLabel{}
	return &this
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *ResourceLabel) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceLabel) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *ResourceLabel) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *ResourceLabel) SetOrn(v string) {
	o.Orn = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *ResourceLabel) GetProfile() ExternalResourceProfile {
	if o == nil || IsNil(o.Profile) {
		var ret ExternalResourceProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceLabel) GetProfileOk() (*ExternalResourceProfile, bool) {
	if o == nil || IsNil(o.Profile) {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *ResourceLabel) HasProfile() bool {
	if o != nil && !IsNil(o.Profile) {
		return true
	}

	return false
}

// SetProfile gets a reference to the given ExternalResourceProfile and assigns it to the Profile field.
func (o *ResourceLabel) SetProfile(v ExternalResourceProfile) {
	o.Profile = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *ResourceLabel) GetLabels() []Label {
	if o == nil || IsNil(o.Labels) {
		var ret []Label
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceLabel) GetLabelsOk() ([]Label, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ResourceLabel) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []Label and assigns it to the Labels field.
func (o *ResourceLabel) SetLabels(v []Label) {
	o.Labels = v
}

// GetLinks returns the Links field value
func (o *ResourceLabel) GetLinks() LinkSelf {
	if o == nil {
		var ret LinkSelf
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ResourceLabel) GetLinksOk() (*LinkSelf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ResourceLabel) SetLinks(v LinkSelf) {
	o.Links = v
}

func (o ResourceLabel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceLabel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Orn) {
		toSerialize["orn"] = o.Orn
	}
	if !IsNil(o.Profile) {
		toSerialize["profile"] = o.Profile
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceLabel) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varResourceLabel := _ResourceLabel{}

	err = json.Unmarshal(data, &varResourceLabel)

	if err != nil {
		return err
	}

	*o = ResourceLabel(varResourceLabel)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "orn")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "labels")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceLabel struct {
	value *ResourceLabel
	isSet bool
}

func (v NullableResourceLabel) Get() *ResourceLabel {
	return v.value
}

func (v *NullableResourceLabel) Set(val *ResourceLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceLabel(val *ResourceLabel) *NullableResourceLabel {
	return &NullableResourceLabel{value: val, isSet: true}
}

func (v NullableResourceLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
