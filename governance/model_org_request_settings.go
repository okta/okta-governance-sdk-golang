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

// checks if the OrgRequestSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrgRequestSettings{}

// OrgRequestSettings Request settings for the org
type OrgRequestSettings struct {
	// Indicates that Access Requests provisioning was triggered by the customer (such as in [Govern Okta admin roles](https://help.okta.com/okta_help.htm?type=oie&id=csh-governance-admin-roles))
	SubprocessorsAcknowledged bool               `json:"subprocessorsAcknowledged"`
	ProvisioningStatus        ProvisioningStatus `json:"provisioningStatus"`
	// The supported request experiences in the org
	RequestExperiences []RequestExperience `json:"requestExperiences"`
	// Indicates that Access Request provisioning was triggered more than 10 minutes ago
	LongTimePastProvisioned   bool                       `json:"longTimePastProvisioned"`
	ResourceCatalogVisibility *ResourceCatalogVisibility `json:"resourceCatalogVisibility,omitempty"`
	Integrations              *RequestIntegrations       `json:"integrations,omitempty"`
	AdditionalProperties      map[string]interface{}
}

type _OrgRequestSettings OrgRequestSettings

// NewOrgRequestSettings instantiates a new OrgRequestSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrgRequestSettings(subprocessorsAcknowledged bool, provisioningStatus ProvisioningStatus, requestExperiences []RequestExperience, longTimePastProvisioned bool) *OrgRequestSettings {
	this := OrgRequestSettings{}
	this.SubprocessorsAcknowledged = subprocessorsAcknowledged
	this.ProvisioningStatus = provisioningStatus
	this.RequestExperiences = requestExperiences
	this.LongTimePastProvisioned = longTimePastProvisioned
	return &this
}

// NewOrgRequestSettingsWithDefaults instantiates a new OrgRequestSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrgRequestSettingsWithDefaults() *OrgRequestSettings {
	this := OrgRequestSettings{}
	return &this
}

// GetSubprocessorsAcknowledged returns the SubprocessorsAcknowledged field value
func (o *OrgRequestSettings) GetSubprocessorsAcknowledged() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SubprocessorsAcknowledged
}

// GetSubprocessorsAcknowledgedOk returns a tuple with the SubprocessorsAcknowledged field value
// and a boolean to check if the value has been set.
func (o *OrgRequestSettings) GetSubprocessorsAcknowledgedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SubprocessorsAcknowledged, true
}

// SetSubprocessorsAcknowledged sets field value
func (o *OrgRequestSettings) SetSubprocessorsAcknowledged(v bool) {
	o.SubprocessorsAcknowledged = v
}

// GetProvisioningStatus returns the ProvisioningStatus field value
func (o *OrgRequestSettings) GetProvisioningStatus() ProvisioningStatus {
	if o == nil {
		var ret ProvisioningStatus
		return ret
	}

	return o.ProvisioningStatus
}

// GetProvisioningStatusOk returns a tuple with the ProvisioningStatus field value
// and a boolean to check if the value has been set.
func (o *OrgRequestSettings) GetProvisioningStatusOk() (*ProvisioningStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProvisioningStatus, true
}

// SetProvisioningStatus sets field value
func (o *OrgRequestSettings) SetProvisioningStatus(v ProvisioningStatus) {
	o.ProvisioningStatus = v
}

// GetRequestExperiences returns the RequestExperiences field value
func (o *OrgRequestSettings) GetRequestExperiences() []RequestExperience {
	if o == nil {
		var ret []RequestExperience
		return ret
	}

	return o.RequestExperiences
}

// GetRequestExperiencesOk returns a tuple with the RequestExperiences field value
// and a boolean to check if the value has been set.
func (o *OrgRequestSettings) GetRequestExperiencesOk() ([]RequestExperience, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestExperiences, true
}

// SetRequestExperiences sets field value
func (o *OrgRequestSettings) SetRequestExperiences(v []RequestExperience) {
	o.RequestExperiences = v
}

// GetLongTimePastProvisioned returns the LongTimePastProvisioned field value
func (o *OrgRequestSettings) GetLongTimePastProvisioned() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.LongTimePastProvisioned
}

// GetLongTimePastProvisionedOk returns a tuple with the LongTimePastProvisioned field value
// and a boolean to check if the value has been set.
func (o *OrgRequestSettings) GetLongTimePastProvisionedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LongTimePastProvisioned, true
}

// SetLongTimePastProvisioned sets field value
func (o *OrgRequestSettings) SetLongTimePastProvisioned(v bool) {
	o.LongTimePastProvisioned = v
}

// GetResourceCatalogVisibility returns the ResourceCatalogVisibility field value if set, zero value otherwise.
func (o *OrgRequestSettings) GetResourceCatalogVisibility() ResourceCatalogVisibility {
	if o == nil || IsNil(o.ResourceCatalogVisibility) {
		var ret ResourceCatalogVisibility
		return ret
	}
	return *o.ResourceCatalogVisibility
}

// GetResourceCatalogVisibilityOk returns a tuple with the ResourceCatalogVisibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgRequestSettings) GetResourceCatalogVisibilityOk() (*ResourceCatalogVisibility, bool) {
	if o == nil || IsNil(o.ResourceCatalogVisibility) {
		return nil, false
	}
	return o.ResourceCatalogVisibility, true
}

// HasResourceCatalogVisibility returns a boolean if a field has been set.
func (o *OrgRequestSettings) HasResourceCatalogVisibility() bool {
	if o != nil && !IsNil(o.ResourceCatalogVisibility) {
		return true
	}

	return false
}

// SetResourceCatalogVisibility gets a reference to the given ResourceCatalogVisibility and assigns it to the ResourceCatalogVisibility field.
func (o *OrgRequestSettings) SetResourceCatalogVisibility(v ResourceCatalogVisibility) {
	o.ResourceCatalogVisibility = &v
}

// GetIntegrations returns the Integrations field value if set, zero value otherwise.
func (o *OrgRequestSettings) GetIntegrations() RequestIntegrations {
	if o == nil || IsNil(o.Integrations) {
		var ret RequestIntegrations
		return ret
	}
	return *o.Integrations
}

// GetIntegrationsOk returns a tuple with the Integrations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrgRequestSettings) GetIntegrationsOk() (*RequestIntegrations, bool) {
	if o == nil || IsNil(o.Integrations) {
		return nil, false
	}
	return o.Integrations, true
}

// HasIntegrations returns a boolean if a field has been set.
func (o *OrgRequestSettings) HasIntegrations() bool {
	if o != nil && !IsNil(o.Integrations) {
		return true
	}

	return false
}

// SetIntegrations gets a reference to the given RequestIntegrations and assigns it to the Integrations field.
func (o *OrgRequestSettings) SetIntegrations(v RequestIntegrations) {
	o.Integrations = &v
}

func (o OrgRequestSettings) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrgRequestSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subprocessorsAcknowledged"] = o.SubprocessorsAcknowledged
	toSerialize["provisioningStatus"] = o.ProvisioningStatus
	toSerialize["requestExperiences"] = o.RequestExperiences
	toSerialize["longTimePastProvisioned"] = o.LongTimePastProvisioned
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

func (o *OrgRequestSettings) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subprocessorsAcknowledged",
		"provisioningStatus",
		"requestExperiences",
		"longTimePastProvisioned",
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

	varOrgRequestSettings := _OrgRequestSettings{}

	err = json.Unmarshal(data, &varOrgRequestSettings)

	if err != nil {
		return err
	}

	*o = OrgRequestSettings(varOrgRequestSettings)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "subprocessorsAcknowledged")
		delete(additionalProperties, "provisioningStatus")
		delete(additionalProperties, "requestExperiences")
		delete(additionalProperties, "longTimePastProvisioned")
		delete(additionalProperties, "resourceCatalogVisibility")
		delete(additionalProperties, "integrations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableOrgRequestSettings struct {
	value *OrgRequestSettings
	isSet bool
}

func (v NullableOrgRequestSettings) Get() *OrgRequestSettings {
	return v.value
}

func (v *NullableOrgRequestSettings) Set(val *OrgRequestSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableOrgRequestSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableOrgRequestSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrgRequestSettings(val *OrgRequestSettings) *NullableOrgRequestSettings {
	return &NullableOrgRequestSettings{value: val, isSet: true}
}

func (v NullableOrgRequestSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrgRequestSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
