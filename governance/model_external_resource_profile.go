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

// checks if the ExternalResourceProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalResourceProfile{}

// ExternalResourceProfile A limited set of properties from the resource's profile
type ExternalResourceProfile struct {
	// Okta resource `id`
	Id string `json:"id"`
	// The display name for the resource.
	Name string `json:"name"`
	// The description of the resource.
	Description *string `json:"description,omitempty"`
	// The label of the Okta resource
	Label *string `json:"label,omitempty"`
	// List of logo resources
	Logo                 []Link `json:"logo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExternalResourceProfile ExternalResourceProfile

// NewExternalResourceProfile instantiates a new ExternalResourceProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalResourceProfile(id string, name string) *ExternalResourceProfile {
	this := ExternalResourceProfile{}
	this.Id = id
	this.Name = name
	return &this
}

// NewExternalResourceProfileWithDefaults instantiates a new ExternalResourceProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalResourceProfileWithDefaults() *ExternalResourceProfile {
	this := ExternalResourceProfile{}
	return &this
}

// GetId returns the Id field value
func (o *ExternalResourceProfile) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExternalResourceProfile) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalResourceProfile) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ExternalResourceProfile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExternalResourceProfile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExternalResourceProfile) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ExternalResourceProfile) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalResourceProfile) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ExternalResourceProfile) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ExternalResourceProfile) SetDescription(v string) {
	o.Description = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *ExternalResourceProfile) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalResourceProfile) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *ExternalResourceProfile) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *ExternalResourceProfile) SetLabel(v string) {
	o.Label = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *ExternalResourceProfile) GetLogo() []Link {
	if o == nil || IsNil(o.Logo) {
		var ret []Link
		return ret
	}
	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalResourceProfile) GetLogoOk() ([]Link, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *ExternalResourceProfile) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given []Link and assigns it to the Logo field.
func (o *ExternalResourceProfile) SetLogo(v []Link) {
	o.Logo = v
}

func (o ExternalResourceProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalResourceProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
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

func (o *ExternalResourceProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varExternalResourceProfile := _ExternalResourceProfile{}

	err = json.Unmarshal(data, &varExternalResourceProfile)

	if err != nil {
		return err
	}

	*o = ExternalResourceProfile(varExternalResourceProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "label")
		delete(additionalProperties, "logo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExternalResourceProfile struct {
	value *ExternalResourceProfile
	isSet bool
}

func (v NullableExternalResourceProfile) Get() *ExternalResourceProfile {
	return v.value
}

func (v *NullableExternalResourceProfile) Set(val *ExternalResourceProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalResourceProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalResourceProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalResourceProfile(val *ExternalResourceProfile) *NullableExternalResourceProfile {
	return &NullableExternalResourceProfile{value: val, isSet: true}
}

func (v NullableExternalResourceProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalResourceProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
