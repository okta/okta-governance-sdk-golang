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

// checks if the ResourceProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceProfile{}

// ResourceProfile A limited set of properties from the resource's profile
type ResourceProfile struct {
	// Okta application instance `id`
	Id *string `json:"id,omitempty"`
	// The name of the Okta application
	Name *string `json:"name,omitempty"`
	// The label of the Okta application
	Label *string `json:"label,omitempty"`
	// List of app logo resources
	Logo                 []Link `json:"logo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ResourceProfile ResourceProfile

// NewResourceProfile instantiates a new ResourceProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceProfile() *ResourceProfile {
	this := ResourceProfile{}
	return &this
}

// NewResourceProfileWithDefaults instantiates a new ResourceProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceProfileWithDefaults() *ResourceProfile {
	this := ResourceProfile{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ResourceProfile) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceProfile) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ResourceProfile) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ResourceProfile) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResourceProfile) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceProfile) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResourceProfile) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResourceProfile) SetName(v string) {
	o.Name = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ResourceProfile) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceProfile) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ResourceProfile) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ResourceProfile) SetLabel(v string) {
	o.Label = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *ResourceProfile) GetLogo() []Link {
	if o == nil || IsNil(o.Logo) {
		var ret []Link
		return ret
	}
	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceProfile) GetLogoOk() ([]Link, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *ResourceProfile) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given []Link and assigns it to the Logo field.
func (o *ResourceProfile) SetLogo(v []Link) {
	o.Logo = v
}

func (o ResourceProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ResourceProfile) UnmarshalJSON(data []byte) (err error) {
	varResourceProfile := _ResourceProfile{}

	err = json.Unmarshal(data, &varResourceProfile)

	if err != nil {
		return err
	}

	*o = ResourceProfile(varResourceProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "logo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableResourceProfile struct {
	value *ResourceProfile
	isSet bool
}

func (v NullableResourceProfile) Get() *ResourceProfile {
	return v.value
}

func (v *NullableResourceProfile) Set(val *ResourceProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceProfile(val *ResourceProfile) *NullableResourceProfile {
	return &NullableResourceProfile{value: val, isSet: true}
}

func (v NullableResourceProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
