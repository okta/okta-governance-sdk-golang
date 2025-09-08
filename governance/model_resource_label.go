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
	if o == nil || o.Orn == nil {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceLabel) GetOrnOk() (*string, bool) {
	if o == nil || o.Orn == nil {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *ResourceLabel) HasOrn() bool {
	if o != nil && o.Orn != nil {
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
	if o == nil || o.Profile == nil {
		var ret ExternalResourceProfile
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceLabel) GetProfileOk() (*ExternalResourceProfile, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *ResourceLabel) HasProfile() bool {
	if o != nil && o.Profile != nil {
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
	if o == nil || o.Labels == nil {
		var ret []Label
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceLabel) GetLabelsOk() ([]Label, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *ResourceLabel) HasLabels() bool {
	if o != nil && o.Labels != nil {
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
	toSerialize := map[string]interface{}{}
	if o.Orn != nil {
		toSerialize["orn"] = o.Orn
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if true {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ResourceLabel) UnmarshalJSON(bytes []byte) (err error) {
	varResourceLabel := _ResourceLabel{}

	err = json.Unmarshal(bytes, &varResourceLabel)
	if err == nil {
		*o = ResourceLabel(varResourceLabel)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "orn")
		delete(additionalProperties, "profile")
		delete(additionalProperties, "labels")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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
