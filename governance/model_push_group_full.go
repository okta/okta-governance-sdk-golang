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

// checks if the PushGroupFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PushGroupFull{}

// PushGroupFull Full representation of a push group
type PushGroupFull struct {
	// Unique Okta Group ID for the push group
	Id *string `json:"id,omitempty"`
	// The name of the push group
	Name *string `json:"name,omitempty"`
	// The description of the push group
	Description *string `json:"description,omitempty"`
	// List of push group logo resources
	Logo                 []Link `json:"logo,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PushGroupFull PushGroupFull

// NewPushGroupFull instantiates a new PushGroupFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPushGroupFull() *PushGroupFull {
	this := PushGroupFull{}
	return &this
}

// NewPushGroupFullWithDefaults instantiates a new PushGroupFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPushGroupFullWithDefaults() *PushGroupFull {
	this := PushGroupFull{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PushGroupFull) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushGroupFull) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PushGroupFull) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PushGroupFull) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PushGroupFull) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushGroupFull) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PushGroupFull) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PushGroupFull) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PushGroupFull) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushGroupFull) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PushGroupFull) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PushGroupFull) SetDescription(v string) {
	o.Description = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *PushGroupFull) GetLogo() []Link {
	if o == nil || IsNil(o.Logo) {
		var ret []Link
		return ret
	}
	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PushGroupFull) GetLogoOk() ([]Link, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *PushGroupFull) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given []Link and assigns it to the Logo field.
func (o *PushGroupFull) SetLogo(v []Link) {
	o.Logo = v
}

func (o PushGroupFull) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PushGroupFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PushGroupFull) UnmarshalJSON(data []byte) (err error) {
	varPushGroupFull := _PushGroupFull{}

	err = json.Unmarshal(data, &varPushGroupFull)

	if err != nil {
		return err
	}

	*o = PushGroupFull(varPushGroupFull)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "logo")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePushGroupFull struct {
	value *PushGroupFull
	isSet bool
}

func (v NullablePushGroupFull) Get() *PushGroupFull {
	return v.value
}

func (v *NullablePushGroupFull) Set(val *PushGroupFull) {
	v.value = val
	v.isSet = true
}

func (v NullablePushGroupFull) IsSet() bool {
	return v.isSet
}

func (v *NullablePushGroupFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePushGroupFull(val *PushGroupFull) *NullablePushGroupFull {
	return &NullablePushGroupFull{value: val, isSet: true}
}

func (v NullablePushGroupFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePushGroupFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
