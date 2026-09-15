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

// checks if the EntitlementDriftEntitlementReference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementDriftEntitlementReference{}

// EntitlementDriftEntitlementReference Identifies the entitlement the drift is about. `entitlementId` is always present. At least one of `id`, `value`, and `externalId` is also present.
type EntitlementDriftEntitlementReference struct {
	// The `id` of the entitlement the drift applies to
	EntitlementId string `json:"entitlementId"`
	// The `id` of the entitlement value, after Okta has registered it
	Id *string `json:"id,omitempty"`
	// The entitlement value as seen in the import, or as Okta governed it before this import
	Value *string `json:"value,omitempty"`
	// The entitlement value's ID in the source app
	ExternalId           *string `json:"externalId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementDriftEntitlementReference EntitlementDriftEntitlementReference

// NewEntitlementDriftEntitlementReference instantiates a new EntitlementDriftEntitlementReference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementDriftEntitlementReference(entitlementId string) *EntitlementDriftEntitlementReference {
	this := EntitlementDriftEntitlementReference{}
	this.EntitlementId = entitlementId
	return &this
}

// NewEntitlementDriftEntitlementReferenceWithDefaults instantiates a new EntitlementDriftEntitlementReference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementDriftEntitlementReferenceWithDefaults() *EntitlementDriftEntitlementReference {
	this := EntitlementDriftEntitlementReference{}
	return &this
}

// GetEntitlementId returns the EntitlementId field value
func (o *EntitlementDriftEntitlementReference) GetEntitlementId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EntitlementId
}

// GetEntitlementIdOk returns a tuple with the EntitlementId field value
// and a boolean to check if the value has been set.
func (o *EntitlementDriftEntitlementReference) GetEntitlementIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EntitlementId, true
}

// SetEntitlementId sets field value
func (o *EntitlementDriftEntitlementReference) SetEntitlementId(v string) {
	o.EntitlementId = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntitlementDriftEntitlementReference) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDriftEntitlementReference) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntitlementDriftEntitlementReference) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EntitlementDriftEntitlementReference) SetId(v string) {
	o.Id = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *EntitlementDriftEntitlementReference) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDriftEntitlementReference) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *EntitlementDriftEntitlementReference) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *EntitlementDriftEntitlementReference) SetValue(v string) {
	o.Value = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *EntitlementDriftEntitlementReference) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementDriftEntitlementReference) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *EntitlementDriftEntitlementReference) HasExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *EntitlementDriftEntitlementReference) SetExternalId(v string) {
	o.ExternalId = &v
}

func (o EntitlementDriftEntitlementReference) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementDriftEntitlementReference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["entitlementId"] = o.EntitlementId
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementDriftEntitlementReference) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"entitlementId",
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

	varEntitlementDriftEntitlementReference := _EntitlementDriftEntitlementReference{}

	err = json.Unmarshal(data, &varEntitlementDriftEntitlementReference)

	if err != nil {
		return err
	}

	*o = EntitlementDriftEntitlementReference(varEntitlementDriftEntitlementReference)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entitlementId")
		delete(additionalProperties, "id")
		delete(additionalProperties, "value")
		delete(additionalProperties, "externalId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementDriftEntitlementReference struct {
	value *EntitlementDriftEntitlementReference
	isSet bool
}

func (v NullableEntitlementDriftEntitlementReference) Get() *EntitlementDriftEntitlementReference {
	return v.value
}

func (v *NullableEntitlementDriftEntitlementReference) Set(val *EntitlementDriftEntitlementReference) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementDriftEntitlementReference) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementDriftEntitlementReference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementDriftEntitlementReference(val *EntitlementDriftEntitlementReference) *NullableEntitlementDriftEntitlementReference {
	return &NullableEntitlementDriftEntitlementReference{value: val, isSet: true}
}

func (v NullableEntitlementDriftEntitlementReference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementDriftEntitlementReference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
