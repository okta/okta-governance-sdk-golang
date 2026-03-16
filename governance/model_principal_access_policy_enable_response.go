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

// checks if the PrincipalAccessPolicyEnableResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrincipalAccessPolicyEnableResponse{}

// PrincipalAccessPolicyEnableResponse Principal access policy enable
type PrincipalAccessPolicyEnableResponse struct {
	// The Okta user, in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format.
	PrincipalOrn string `json:"principalOrn"`
	// The Okta resource, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for [supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	ResourceOrn    string                `json:"resourceOrn"`
	AccessDuration *AssignmentProperties `json:"accessDuration,omitempty"`
	// Collection of entitlements with associated values
	EffectiveEntitlements []EntitlementFull `json:"effectiveEntitlements"`
	// Entitlement access details
	EntitlementAccessDetails []EntitlementAccessDetailsObject `json:"entitlementAccessDetails,omitempty"`
	AdditionalProperties     map[string]interface{}
}

type _PrincipalAccessPolicyEnableResponse PrincipalAccessPolicyEnableResponse

// NewPrincipalAccessPolicyEnableResponse instantiates a new PrincipalAccessPolicyEnableResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalAccessPolicyEnableResponse(principalOrn string, resourceOrn string, effectiveEntitlements []EntitlementFull) *PrincipalAccessPolicyEnableResponse {
	this := PrincipalAccessPolicyEnableResponse{}
	this.PrincipalOrn = principalOrn
	this.ResourceOrn = resourceOrn
	this.EffectiveEntitlements = effectiveEntitlements
	return &this
}

// NewPrincipalAccessPolicyEnableResponseWithDefaults instantiates a new PrincipalAccessPolicyEnableResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalAccessPolicyEnableResponseWithDefaults() *PrincipalAccessPolicyEnableResponse {
	this := PrincipalAccessPolicyEnableResponse{}
	return &this
}

// GetPrincipalOrn returns the PrincipalOrn field value
func (o *PrincipalAccessPolicyEnableResponse) GetPrincipalOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrincipalOrn
}

// GetPrincipalOrnOk returns a tuple with the PrincipalOrn field value
// and a boolean to check if the value has been set.
func (o *PrincipalAccessPolicyEnableResponse) GetPrincipalOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalOrn, true
}

// SetPrincipalOrn sets field value
func (o *PrincipalAccessPolicyEnableResponse) SetPrincipalOrn(v string) {
	o.PrincipalOrn = v
}

// GetResourceOrn returns the ResourceOrn field value
func (o *PrincipalAccessPolicyEnableResponse) GetResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceOrn
}

// GetResourceOrnOk returns a tuple with the ResourceOrn field value
// and a boolean to check if the value has been set.
func (o *PrincipalAccessPolicyEnableResponse) GetResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceOrn, true
}

// SetResourceOrn sets field value
func (o *PrincipalAccessPolicyEnableResponse) SetResourceOrn(v string) {
	o.ResourceOrn = v
}

// GetAccessDuration returns the AccessDuration field value if set, zero value otherwise.
func (o *PrincipalAccessPolicyEnableResponse) GetAccessDuration() AssignmentProperties {
	if o == nil || IsNil(o.AccessDuration) {
		var ret AssignmentProperties
		return ret
	}
	return *o.AccessDuration
}

// GetAccessDurationOk returns a tuple with the AccessDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalAccessPolicyEnableResponse) GetAccessDurationOk() (*AssignmentProperties, bool) {
	if o == nil || IsNil(o.AccessDuration) {
		return nil, false
	}
	return o.AccessDuration, true
}

// HasAccessDuration returns a boolean if a field has been set.
func (o *PrincipalAccessPolicyEnableResponse) HasAccessDuration() bool {
	if o != nil && !IsNil(o.AccessDuration) {
		return true
	}

	return false
}

// SetAccessDuration gets a reference to the given AssignmentProperties and assigns it to the AccessDuration field.
func (o *PrincipalAccessPolicyEnableResponse) SetAccessDuration(v AssignmentProperties) {
	o.AccessDuration = &v
}

// GetEffectiveEntitlements returns the EffectiveEntitlements field value
func (o *PrincipalAccessPolicyEnableResponse) GetEffectiveEntitlements() []EntitlementFull {
	if o == nil {
		var ret []EntitlementFull
		return ret
	}

	return o.EffectiveEntitlements
}

// GetEffectiveEntitlementsOk returns a tuple with the EffectiveEntitlements field value
// and a boolean to check if the value has been set.
func (o *PrincipalAccessPolicyEnableResponse) GetEffectiveEntitlementsOk() ([]EntitlementFull, bool) {
	if o == nil {
		return nil, false
	}
	return o.EffectiveEntitlements, true
}

// SetEffectiveEntitlements sets field value
func (o *PrincipalAccessPolicyEnableResponse) SetEffectiveEntitlements(v []EntitlementFull) {
	o.EffectiveEntitlements = v
}

// GetEntitlementAccessDetails returns the EntitlementAccessDetails field value if set, zero value otherwise.
func (o *PrincipalAccessPolicyEnableResponse) GetEntitlementAccessDetails() []EntitlementAccessDetailsObject {
	if o == nil || IsNil(o.EntitlementAccessDetails) {
		var ret []EntitlementAccessDetailsObject
		return ret
	}
	return o.EntitlementAccessDetails
}

// GetEntitlementAccessDetailsOk returns a tuple with the EntitlementAccessDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalAccessPolicyEnableResponse) GetEntitlementAccessDetailsOk() ([]EntitlementAccessDetailsObject, bool) {
	if o == nil || IsNil(o.EntitlementAccessDetails) {
		return nil, false
	}
	return o.EntitlementAccessDetails, true
}

// HasEntitlementAccessDetails returns a boolean if a field has been set.
func (o *PrincipalAccessPolicyEnableResponse) HasEntitlementAccessDetails() bool {
	if o != nil && !IsNil(o.EntitlementAccessDetails) {
		return true
	}

	return false
}

// SetEntitlementAccessDetails gets a reference to the given []EntitlementAccessDetailsObject and assigns it to the EntitlementAccessDetails field.
func (o *PrincipalAccessPolicyEnableResponse) SetEntitlementAccessDetails(v []EntitlementAccessDetailsObject) {
	o.EntitlementAccessDetails = v
}

func (o PrincipalAccessPolicyEnableResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrincipalAccessPolicyEnableResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["principalOrn"] = o.PrincipalOrn
	toSerialize["resourceOrn"] = o.ResourceOrn
	if !IsNil(o.AccessDuration) {
		toSerialize["accessDuration"] = o.AccessDuration
	}
	toSerialize["effectiveEntitlements"] = o.EffectiveEntitlements
	if !IsNil(o.EntitlementAccessDetails) {
		toSerialize["entitlementAccessDetails"] = o.EntitlementAccessDetails
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrincipalAccessPolicyEnableResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"principalOrn",
		"resourceOrn",
		"effectiveEntitlements",
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

	varPrincipalAccessPolicyEnableResponse := _PrincipalAccessPolicyEnableResponse{}

	err = json.Unmarshal(data, &varPrincipalAccessPolicyEnableResponse)

	if err != nil {
		return err
	}

	*o = PrincipalAccessPolicyEnableResponse(varPrincipalAccessPolicyEnableResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "principalOrn")
		delete(additionalProperties, "resourceOrn")
		delete(additionalProperties, "accessDuration")
		delete(additionalProperties, "effectiveEntitlements")
		delete(additionalProperties, "entitlementAccessDetails")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrincipalAccessPolicyEnableResponse struct {
	value *PrincipalAccessPolicyEnableResponse
	isSet bool
}

func (v NullablePrincipalAccessPolicyEnableResponse) Get() *PrincipalAccessPolicyEnableResponse {
	return v.value
}

func (v *NullablePrincipalAccessPolicyEnableResponse) Set(val *PrincipalAccessPolicyEnableResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalAccessPolicyEnableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalAccessPolicyEnableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalAccessPolicyEnableResponse(val *PrincipalAccessPolicyEnableResponse) *NullablePrincipalAccessPolicyEnableResponse {
	return &NullablePrincipalAccessPolicyEnableResponse{value: val, isSet: true}
}

func (v NullablePrincipalAccessPolicyEnableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalAccessPolicyEnableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
