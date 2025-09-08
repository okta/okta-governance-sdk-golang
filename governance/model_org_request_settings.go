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

// OrgRequestSettings Request settings for the org
type OrgRequestSettings struct {
	// Whether a customer has acknowledged Access Requests subprocessors
	SubprocessorsAcknowledged bool               `json:"subprocessorsAcknowledged"`
	ProvisioningStatus        ProvisioningStatus `json:"provisioningStatus"`
	// Which request experiences this org supports
	RequestExperiences []RequestExperience `json:"requestExperiences"`
	// Whether it has been a long time since the Access Requests org has been provisioned
	LongTimePastProvisioned bool `json:"longTimePastProvisioned"`
	AdditionalProperties    map[string]interface{}
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

func (o OrgRequestSettings) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subprocessorsAcknowledged"] = o.SubprocessorsAcknowledged
	}
	if true {
		toSerialize["provisioningStatus"] = o.ProvisioningStatus
	}
	if true {
		toSerialize["requestExperiences"] = o.RequestExperiences
	}
	if true {
		toSerialize["longTimePastProvisioned"] = o.LongTimePastProvisioned
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *OrgRequestSettings) UnmarshalJSON(bytes []byte) (err error) {
	varOrgRequestSettings := _OrgRequestSettings{}

	err = json.Unmarshal(bytes, &varOrgRequestSettings)
	if err == nil {
		*o = OrgRequestSettings(varOrgRequestSettings)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "subprocessorsAcknowledged")
		delete(additionalProperties, "provisioningStatus")
		delete(additionalProperties, "requestExperiences")
		delete(additionalProperties, "longTimePastProvisioned")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
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
