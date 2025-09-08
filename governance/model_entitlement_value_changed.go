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

// EntitlementValueChanged Attributes related to an entitlement value
type EntitlementValueChanged struct {
	// Type of change that occurred to the entitlement
	ChangeType *string `json:"changeType,omitempty"`
	// The `id` of an entitlement value
	Id *string `json:"id,omitempty"`
	// The display name for an entitlement value
	Name *string `json:"name,omitempty"`
	// The value of an entitlement property value
	ExternalValue *string `json:"externalValue,omitempty"`
	// The read-only `id` of an entitlement property value in the downstream application.
	ExternalId *string `json:"externalId,omitempty"`
	// The description of an entitlement value
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementValueChanged EntitlementValueChanged

// NewEntitlementValueChanged instantiates a new EntitlementValueChanged object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementValueChanged() *EntitlementValueChanged {
	this := EntitlementValueChanged{}
	return &this
}

// NewEntitlementValueChangedWithDefaults instantiates a new EntitlementValueChanged object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementValueChangedWithDefaults() *EntitlementValueChanged {
	this := EntitlementValueChanged{}
	return &this
}

// GetChangeType returns the ChangeType field value if set, zero value otherwise.
func (o *EntitlementValueChanged) GetChangeType() string {
	if o == nil || o.ChangeType == nil {
		var ret string
		return ret
	}
	return *o.ChangeType
}

// GetChangeTypeOk returns a tuple with the ChangeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValueChanged) GetChangeTypeOk() (*string, bool) {
	if o == nil || o.ChangeType == nil {
		return nil, false
	}
	return o.ChangeType, true
}

// HasChangeType returns a boolean if a field has been set.
func (o *EntitlementValueChanged) HasChangeType() bool {
	if o != nil && o.ChangeType != nil {
		return true
	}

	return false
}

// SetChangeType gets a reference to the given string and assigns it to the ChangeType field.
func (o *EntitlementValueChanged) SetChangeType(v string) {
	o.ChangeType = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntitlementValueChanged) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValueChanged) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntitlementValueChanged) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EntitlementValueChanged) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EntitlementValueChanged) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValueChanged) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EntitlementValueChanged) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EntitlementValueChanged) SetName(v string) {
	o.Name = &v
}

// GetExternalValue returns the ExternalValue field value if set, zero value otherwise.
func (o *EntitlementValueChanged) GetExternalValue() string {
	if o == nil || o.ExternalValue == nil {
		var ret string
		return ret
	}
	return *o.ExternalValue
}

// GetExternalValueOk returns a tuple with the ExternalValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValueChanged) GetExternalValueOk() (*string, bool) {
	if o == nil || o.ExternalValue == nil {
		return nil, false
	}
	return o.ExternalValue, true
}

// HasExternalValue returns a boolean if a field has been set.
func (o *EntitlementValueChanged) HasExternalValue() bool {
	if o != nil && o.ExternalValue != nil {
		return true
	}

	return false
}

// SetExternalValue gets a reference to the given string and assigns it to the ExternalValue field.
func (o *EntitlementValueChanged) SetExternalValue(v string) {
	o.ExternalValue = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *EntitlementValueChanged) GetExternalId() string {
	if o == nil || o.ExternalId == nil {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValueChanged) GetExternalIdOk() (*string, bool) {
	if o == nil || o.ExternalId == nil {
		return nil, false
	}
	return o.ExternalId, true
}

// HasExternalId returns a boolean if a field has been set.
func (o *EntitlementValueChanged) HasExternalId() bool {
	if o != nil && o.ExternalId != nil {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *EntitlementValueChanged) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitlementValueChanged) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementValueChanged) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementValueChanged) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitlementValueChanged) SetDescription(v string) {
	o.Description = &v
}

func (o EntitlementValueChanged) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ChangeType != nil {
		toSerialize["changeType"] = o.ChangeType
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.ExternalValue != nil {
		toSerialize["externalValue"] = o.ExternalValue
	}
	if o.ExternalId != nil {
		toSerialize["externalId"] = o.ExternalId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntitlementValueChanged) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementValueChanged := _EntitlementValueChanged{}

	err = json.Unmarshal(bytes, &varEntitlementValueChanged)
	if err == nil {
		*o = EntitlementValueChanged(varEntitlementValueChanged)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "changeType")
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "externalValue")
		delete(additionalProperties, "externalId")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEntitlementValueChanged struct {
	value *EntitlementValueChanged
	isSet bool
}

func (v NullableEntitlementValueChanged) Get() *EntitlementValueChanged {
	return v.value
}

func (v *NullableEntitlementValueChanged) Set(val *EntitlementValueChanged) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementValueChanged) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementValueChanged) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementValueChanged(val *EntitlementValueChanged) *NullableEntitlementValueChanged {
	return &NullableEntitlementValueChanged{value: val, isSet: true}
}

func (v NullableEntitlementValueChanged) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementValueChanged) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
