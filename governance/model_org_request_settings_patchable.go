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

// checks if the OrgRequestSettingsPatchable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgRequestSettingsPatchable{}

// OrgRequestSettingsPatchable Request settings for the org for patch
type OrgRequestSettingsPatchable struct {
	// Indicates that Access Requests provisioning was triggered by the customer (such as in [Govern Okta admin roles](https://help.okta.com/okta_help.htm?type=oie&id=csh-governance-admin-roles))
	SubprocessorsAcknowledged *bool                      `json:"subprocessorsAcknowledged,omitempty"`
	ResourceCatalogVisibility *ResourceCatalogVisibility `json:"resourceCatalogVisibility,omitempty"`
	Integrations              *RequestIntegrations       `json:"integrations,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _OrgRequestSettingsPatchable OrgRequestSettingsPatchable

// NewOrgRequestSettingsPatchable instantiates a new OrgRequestSettingsPatchable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgRequestSettingsPatchable() *OrgRequestSettingsPatchable {
	this := OrgRequestSettingsPatchable{}
	return &this
}

// NewOrgRequestSettingsPatchableWithDefaults instantiates a new OrgRequestSettingsPatchable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgRequestSettingsPatchableWithDefaults() *OrgRequestSettingsPatchable {
	this := OrgRequestSettingsPatchable{}
	return &this
}

// GetSubprocessorsAcknowledged returns the SubprocessorsAcknowledged field value if set, zero value otherwise.
func (o *OrgRequestSettingsPatchable) GetSubprocessorsAcknowledged() bool {
	if o == nil || IsNil(o.SubprocessorsAcknowledged) {
		var ret bool
		return ret
	}
	return *o.SubprocessorsAcknowledged
}

// GetSubprocessorsAcknowledgedOk returns a tuple with the SubprocessorsAcknowledged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgRequestSettingsPatchable) GetSubprocessorsAcknowledgedOk() (*bool, bool) {
	if o == nil || IsNil(o.SubprocessorsAcknowledged) {
		return nil, false
	}
	return o.SubprocessorsAcknowledged, true
}

// HasSubprocessorsAcknowledged returns a boolean if a field has been set.
func (o *OrgRequestSettingsPatchable) HasSubprocessorsAcknowledged() bool {
	if o != nil && !IsNil(o.SubprocessorsAcknowledged) {
		return true
	}

	return false
}

// SetSubprocessorsAcknowledged gets a reference to the given bool and assigns it to the SubprocessorsAcknowledged field.
func (o *OrgRequestSettingsPatchable) SetSubprocessorsAcknowledged(v bool) {
	o.SubprocessorsAcknowledged = &v
}

// GetResourceCatalogVisibility returns the ResourceCatalogVisibility field value if set, zero value otherwise.
func (o *OrgRequestSettingsPatchable) GetResourceCatalogVisibility() ResourceCatalogVisibility {
	if o == nil || IsNil(o.ResourceCatalogVisibility) {
		var ret ResourceCatalogVisibility
		return ret
	}
	return *o.ResourceCatalogVisibility
}

// GetResourceCatalogVisibilityOk returns a tuple with the ResourceCatalogVisibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgRequestSettingsPatchable) GetResourceCatalogVisibilityOk() (*ResourceCatalogVisibility, bool) {
	if o == nil || IsNil(o.ResourceCatalogVisibility) {
		return nil, false
	}
	return o.ResourceCatalogVisibility, true
}

// HasResourceCatalogVisibility returns a boolean if a field has been set.
func (o *OrgRequestSettingsPatchable) HasResourceCatalogVisibility() bool {
	if o != nil && !IsNil(o.ResourceCatalogVisibility) {
		return true
	}

	return false
}

// SetResourceCatalogVisibility gets a reference to the given ResourceCatalogVisibility and assigns it to the ResourceCatalogVisibility field.
func (o *OrgRequestSettingsPatchable) SetResourceCatalogVisibility(v ResourceCatalogVisibility) {
	o.ResourceCatalogVisibility = &v
}

// GetIntegrations returns the Integrations field value if set, zero value otherwise.
func (o *OrgRequestSettingsPatchable) GetIntegrations() RequestIntegrations {
	if o == nil || IsNil(o.Integrations) {
		var ret RequestIntegrations
		return ret
	}
	return *o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgRequestSettingsPatchable) GetIntegrationsOk() (*RequestIntegrations, bool) {
	if o == nil || IsNil(o.Integrations) {
		return nil, false
	}
	return o.Integrations, true
}

// HasIntegrations returns a boolean if a field has been set.
func (o *OrgRequestSettingsPatchable) HasIntegrations() bool {
	if o != nil && !IsNil(o.Integrations) {
		return true
	}

	return false
}

// SetIntegrations gets a reference to the given RequestIntegrations and assigns it to the Integrations field.
func (o *OrgRequestSettingsPatchable) SetIntegrations(v RequestIntegrations) {
	o.Integrations = &v
}

func (o OrgRequestSettingsPatchable) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgRequestSettingsPatchable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SubprocessorsAcknowledged) {
		toSerialize["subprocessorsAcknowledged"] = o.SubprocessorsAcknowledged
	}
	if !IsNil(o.ResourceCatalogVisibility) {
		toSerialize["resourceCatalogVisibility"] = o.ResourceCatalogVisibility
	}
	if !IsNil(o.Integrations) {
		toSerialize["integrations"] = o.Integrations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *OrgRequestSettingsPatchable) UnmarshalJSON(data []byte) (err error) {
	varOrgRequestSettingsPatchable := _OrgRequestSettingsPatchable{}

	err = json.Unmarshal(data, &varOrgRequestSettingsPatchable)

	if err != nil {
		return err
	}

	*o = OrgRequestSettingsPatchable(varOrgRequestSettingsPatchable)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "subprocessorsAcknowledged")
		delete(additionalProperties, "resourceCatalogVisibility")
		delete(additionalProperties, "integrations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgRequestSettingsPatchable struct {
	value *OrgRequestSettingsPatchable
	isSet bool
}

func (v NullableOrgRequestSettingsPatchable) Get() *OrgRequestSettingsPatchable {
	return v.value
}

func (v *NullableOrgRequestSettingsPatchable) Set(val *OrgRequestSettingsPatchable) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgRequestSettingsPatchable) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgRequestSettingsPatchable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgRequestSettingsPatchable(val *OrgRequestSettingsPatchable) *NullableOrgRequestSettingsPatchable {
	return &NullableOrgRequestSettingsPatchable{value: val, isSet: true}
}

func (v NullableOrgRequestSettingsPatchable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgRequestSettingsPatchable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
