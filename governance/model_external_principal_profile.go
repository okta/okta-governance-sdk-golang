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

// checks if the ExternalPrincipalProfile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalPrincipalProfile{}

// ExternalPrincipalProfile A limited set of properties from the principal's profile
type ExternalPrincipalProfile struct {
	// Okta user `id` or Okta group `id`
	Id string `json:"id"`
	// User name or Group Name
	Name string `json:"name"`
	// Email of the resource owner, if applicable.
	Email *string `json:"email,omitempty"`
	// List of logo resources
	Logo                 []Link   `json:"logo,omitempty"`
	Links                LinkSelf `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _ExternalPrincipalProfile ExternalPrincipalProfile

// NewExternalPrincipalProfile instantiates a new ExternalPrincipalProfile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalPrincipalProfile(id string, name string, links LinkSelf) *ExternalPrincipalProfile {
	this := ExternalPrincipalProfile{}
	this.Id = id
	this.Name = name
	this.Links = links
	return &this
}

// NewExternalPrincipalProfileWithDefaults instantiates a new ExternalPrincipalProfile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalPrincipalProfileWithDefaults() *ExternalPrincipalProfile {
	this := ExternalPrincipalProfile{}
	return &this
}

// GetId returns the Id field value
func (o *ExternalPrincipalProfile) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ExternalPrincipalProfile) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ExternalPrincipalProfile) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ExternalPrincipalProfile) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ExternalPrincipalProfile) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ExternalPrincipalProfile) SetName(v string) {
	o.Name = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ExternalPrincipalProfile) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPrincipalProfile) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ExternalPrincipalProfile) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ExternalPrincipalProfile) SetEmail(v string) {
	o.Email = &v
}

// GetLogo returns the Logo field value if set, zero value otherwise.
func (o *ExternalPrincipalProfile) GetLogo() []Link {
	if o == nil || IsNil(o.Logo) {
		var ret []Link
		return ret
	}
	return o.Logo
}

// GetLogoOk returns a tuple with the Logo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalPrincipalProfile) GetLogoOk() ([]Link, bool) {
	if o == nil || IsNil(o.Logo) {
		return nil, false
	}
	return o.Logo, true
}

// HasLogo returns a boolean if a field has been set.
func (o *ExternalPrincipalProfile) HasLogo() bool {
	if o != nil && !IsNil(o.Logo) {
		return true
	}

	return false
}

// SetLogo gets a reference to the given []Link and assigns it to the Logo field.
func (o *ExternalPrincipalProfile) SetLogo(v []Link) {
	o.Logo = v
}

// GetLinks returns the Links field value
func (o *ExternalPrincipalProfile) GetLinks() LinkSelf {
	if o == nil {
		var ret LinkSelf
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ExternalPrincipalProfile) GetLinksOk() (*LinkSelf, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ExternalPrincipalProfile) SetLinks(v LinkSelf) {
	o.Links = v
}

func (o ExternalPrincipalProfile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalPrincipalProfile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["name"] = o.Name
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Logo) {
		toSerialize["logo"] = o.Logo
	}
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ExternalPrincipalProfile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"name",
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

	varExternalPrincipalProfile := _ExternalPrincipalProfile{}

	err = json.Unmarshal(data, &varExternalPrincipalProfile)

	if err != nil {
		return err
	}

	*o = ExternalPrincipalProfile(varExternalPrincipalProfile)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "email")
		delete(additionalProperties, "logo")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExternalPrincipalProfile struct {
	value *ExternalPrincipalProfile
	isSet bool
}

func (v NullableExternalPrincipalProfile) Get() *ExternalPrincipalProfile {
	return v.value
}

func (v *NullableExternalPrincipalProfile) Set(val *ExternalPrincipalProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalPrincipalProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalPrincipalProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalPrincipalProfile(val *ExternalPrincipalProfile) *NullableExternalPrincipalProfile {
	return &NullableExternalPrincipalProfile{value: val, isSet: true}
}

func (v NullableExternalPrincipalProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalPrincipalProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
