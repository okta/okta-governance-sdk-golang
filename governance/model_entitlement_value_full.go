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

// checks if the EntitlementValueFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementValueFull{}

// EntitlementValueFull Attributes related to Entitlement value
type EntitlementValueFull struct {
	// The `id` of the entitlement value
	Id *string `json:"id,omitempty"`
	// The display name for an entitlement value
	Name *string `json:"name,omitempty"`
	// The value of an entitlement property value
	ExternalValue *string `json:"externalValue,omitempty"`
	// The read-only ID of an entitlement property value in the downstream app
	ExternalId *string `json:"externalId,omitempty"`
	// The description of an entitlement value
	Description *string `json:"description,omitempty"`
	// The entitlement value resource, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn)
	Orn                  *string `json:"orn,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementValueFull EntitlementValueFull

// NewEntitlementValueFull instantiates a new EntitlementValueFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementValueFull() *EntitlementValueFull {
	this := EntitlementValueFull{}
	return &this
}

// NewEntitlementValueFullWithDefaults instantiates a new EntitlementValueFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementValueFullWithDefaults() *EntitlementValueFull {
	this := EntitlementValueFull{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntitlementValueFull) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValueFull) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntitlementValueFull) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EntitlementValueFull) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EntitlementValueFull) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValueFull) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EntitlementValueFull) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EntitlementValueFull) SetName(v string) {
	o.Name = &v
}

// GetExternalValue returns the ExternalValue field value if set, zero value otherwise.
func (o *EntitlementValueFull) GetExternalValue() string {
	if o == nil || IsNil(o.ExternalValue) {
		var ret string
		return ret
	}
	return *o.ExternalValue
}

// GetExternalValueOk returns a tuple with the ExternalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValueFull) GetExternalValueOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalValue) {
		return nil, false
	}
	return o.ExternalValue, true
}

// HasExternalValue returns a boolean if a field has been set.
func (o *EntitlementValueFull) HasExternalValue() bool {
	if o != nil && !IsNil(o.ExternalValue) {
		return true
	}

	return false
}

// SetExternalValue gets a reference to the given string and assigns it to the ExternalValue field.
func (o *EntitlementValueFull) SetExternalValue(v string) {
	o.ExternalValue = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *EntitlementValueFull) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValueFull) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *EntitlementValueFull) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *EntitlementValueFull) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitlementValueFull) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValueFull) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementValueFull) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitlementValueFull) SetDescription(v string) {
	o.Description = &v
}

// GetOrn returns the Orn field value if set, zero value otherwise.
func (o *EntitlementValueFull) GetOrn() string {
	if o == nil || IsNil(o.Orn) {
		var ret string
		return ret
	}
	return *o.Orn
}

// GetOrnOk returns a tuple with the Orn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValueFull) GetOrnOk() (*string, bool) {
	if o == nil || IsNil(o.Orn) {
		return nil, false
	}
	return o.Orn, true
}

// HasOrn returns a boolean if a field has been set.
func (o *EntitlementValueFull) HasOrn() bool {
	if o != nil && !IsNil(o.Orn) {
		return true
	}

	return false
}

// SetOrn gets a reference to the given string and assigns it to the Orn field.
func (o *EntitlementValueFull) SetOrn(v string) {
	o.Orn = &v
}

func (o EntitlementValueFull) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementValueFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ExternalValue) {
		toSerialize["externalValue"] = o.ExternalValue
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Orn) {
		toSerialize["orn"] = o.Orn
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementValueFull) UnmarshalJSON(data []byte) (err error) {
	varEntitlementValueFull := _EntitlementValueFull{}

	err = json.Unmarshal(data, &varEntitlementValueFull)

	if err != nil {
		return err
	}

	*o = EntitlementValueFull(varEntitlementValueFull)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "externalValue")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "description")
		delete(additionalProperties, "orn")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementValueFull struct {
	value *EntitlementValueFull
	isSet bool
}

func (v NullableEntitlementValueFull) Get() *EntitlementValueFull {
	return v.value
}

func (v *NullableEntitlementValueFull) Set(val *EntitlementValueFull) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementValueFull) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementValueFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementValueFull(val *EntitlementValueFull) *NullableEntitlementValueFull {
	return &NullableEntitlementValueFull{value: val, isSet: true}
}

func (v NullableEntitlementValueFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementValueFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
