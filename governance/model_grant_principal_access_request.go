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

// checks if the GrantPrincipalAccessRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GrantPrincipalAccessRequest{}

// GrantPrincipalAccessRequest Request body for granting a principal access to a resource. At minimum, `principalOrn` and `resourceOrn` are required to grant app-level access. Optionally, specify `entitlementData` to grant specific entitlements or bundles, and `accessDuration` to set an expiration on the app assignment.
type GrantPrincipalAccessRequest struct {
	// The Okta user in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format
	PrincipalOrn string `json:"principalOrn"`
	// The ORN of the app resource to grant access to
	ResourceOrn          string                         `json:"resourceOrn"`
	Actor                *GrantActor                    `json:"actor,omitempty"`
	AccessDuration       *AccessDuration                `json:"accessDuration,omitempty"`
	EntitlementData      *GrantPrincipalEntitlementData `json:"entitlementData,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _GrantPrincipalAccessRequest GrantPrincipalAccessRequest

// NewGrantPrincipalAccessRequest instantiates a new GrantPrincipalAccessRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrantPrincipalAccessRequest(principalOrn string, resourceOrn string) *GrantPrincipalAccessRequest {
	this := GrantPrincipalAccessRequest{}
	this.PrincipalOrn = principalOrn
	this.ResourceOrn = resourceOrn
	var actor GrantActor = GRANTACTOR_API
	this.Actor = &actor
	return &this
}

// NewGrantPrincipalAccessRequestWithDefaults instantiates a new GrantPrincipalAccessRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGrantPrincipalAccessRequestWithDefaults() *GrantPrincipalAccessRequest {
	this := GrantPrincipalAccessRequest{}
	var actor GrantActor = GRANTACTOR_API
	this.Actor = &actor
	return &this
}

// GetPrincipalOrn returns the PrincipalOrn field value
func (o *GrantPrincipalAccessRequest) GetPrincipalOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrincipalOrn
}

// GetPrincipalOrnOk returns a tuple with the PrincipalOrn field value
// and a boolean to check if the value has been set.
func (o *GrantPrincipalAccessRequest) GetPrincipalOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrincipalOrn, true
}

// SetPrincipalOrn sets field value
func (o *GrantPrincipalAccessRequest) SetPrincipalOrn(v string) {
	o.PrincipalOrn = v
}

// GetResourceOrn returns the ResourceOrn field value
func (o *GrantPrincipalAccessRequest) GetResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceOrn
}

// GetResourceOrnOk returns a tuple with the ResourceOrn field value
// and a boolean to check if the value has been set.
func (o *GrantPrincipalAccessRequest) GetResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceOrn, true
}

// SetResourceOrn sets field value
func (o *GrantPrincipalAccessRequest) SetResourceOrn(v string) {
	o.ResourceOrn = v
}

// GetActor returns the Actor field value if set, zero value otherwise.
func (o *GrantPrincipalAccessRequest) GetActor() GrantActor {
	if o == nil || IsNil(o.Actor) {
		var ret GrantActor
		return ret
	}
	return *o.Actor
}

// GetActorOk returns a tuple with the Actor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantPrincipalAccessRequest) GetActorOk() (*GrantActor, bool) {
	if o == nil || IsNil(o.Actor) {
		return nil, false
	}
	return o.Actor, true
}

// HasActor returns a boolean if a field has been set.
func (o *GrantPrincipalAccessRequest) HasActor() bool {
	if o != nil && !IsNil(o.Actor) {
		return true
	}

	return false
}

// SetActor gets a reference to the given GrantActor and assigns it to the Actor field.
func (o *GrantPrincipalAccessRequest) SetActor(v GrantActor) {
	o.Actor = &v
}

// GetAccessDuration returns the AccessDuration field value if set, zero value otherwise.
func (o *GrantPrincipalAccessRequest) GetAccessDuration() AccessDuration {
	if o == nil || IsNil(o.AccessDuration) {
		var ret AccessDuration
		return ret
	}
	return *o.AccessDuration
}

// GetAccessDurationOk returns a tuple with the AccessDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantPrincipalAccessRequest) GetAccessDurationOk() (*AccessDuration, bool) {
	if o == nil || IsNil(o.AccessDuration) {
		return nil, false
	}
	return o.AccessDuration, true
}

// HasAccessDuration returns a boolean if a field has been set.
func (o *GrantPrincipalAccessRequest) HasAccessDuration() bool {
	if o != nil && !IsNil(o.AccessDuration) {
		return true
	}

	return false
}

// SetAccessDuration gets a reference to the given AccessDuration and assigns it to the AccessDuration field.
func (o *GrantPrincipalAccessRequest) SetAccessDuration(v AccessDuration) {
	o.AccessDuration = &v
}

// GetEntitlementData returns the EntitlementData field value if set, zero value otherwise.
func (o *GrantPrincipalAccessRequest) GetEntitlementData() GrantPrincipalEntitlementData {
	if o == nil || IsNil(o.EntitlementData) {
		var ret GrantPrincipalEntitlementData
		return ret
	}
	return *o.EntitlementData
}

// GetEntitlementDataOk returns a tuple with the EntitlementData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GrantPrincipalAccessRequest) GetEntitlementDataOk() (*GrantPrincipalEntitlementData, bool) {
	if o == nil || IsNil(o.EntitlementData) {
		return nil, false
	}
	return o.EntitlementData, true
}

// HasEntitlementData returns a boolean if a field has been set.
func (o *GrantPrincipalAccessRequest) HasEntitlementData() bool {
	if o != nil && !IsNil(o.EntitlementData) {
		return true
	}

	return false
}

// SetEntitlementData gets a reference to the given GrantPrincipalEntitlementData and assigns it to the EntitlementData field.
func (o *GrantPrincipalAccessRequest) SetEntitlementData(v GrantPrincipalEntitlementData) {
	o.EntitlementData = &v
}

func (o GrantPrincipalAccessRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GrantPrincipalAccessRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["principalOrn"] = o.PrincipalOrn
	toSerialize["resourceOrn"] = o.ResourceOrn
	if !IsNil(o.Actor) {
		toSerialize["actor"] = o.Actor
	}
	if !IsNil(o.AccessDuration) {
		toSerialize["accessDuration"] = o.AccessDuration
	}
	if !IsNil(o.EntitlementData) {
		toSerialize["entitlementData"] = o.EntitlementData
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *GrantPrincipalAccessRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"principalOrn",
		"resourceOrn",
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

	varGrantPrincipalAccessRequest := _GrantPrincipalAccessRequest{}

	err = json.Unmarshal(data, &varGrantPrincipalAccessRequest)

	if err != nil {
		return err
	}

	*o = GrantPrincipalAccessRequest(varGrantPrincipalAccessRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "principalOrn")
		delete(additionalProperties, "resourceOrn")
		delete(additionalProperties, "actor")
		delete(additionalProperties, "accessDuration")
		delete(additionalProperties, "entitlementData")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableGrantPrincipalAccessRequest struct {
	value *GrantPrincipalAccessRequest
	isSet bool
}

func (v NullableGrantPrincipalAccessRequest) Get() *GrantPrincipalAccessRequest {
	return v.value
}

func (v *NullableGrantPrincipalAccessRequest) Set(val *GrantPrincipalAccessRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantPrincipalAccessRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantPrincipalAccessRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantPrincipalAccessRequest(val *GrantPrincipalAccessRequest) *NullableGrantPrincipalAccessRequest {
	return &NullableGrantPrincipalAccessRequest{value: val, isSet: true}
}

func (v NullableGrantPrincipalAccessRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantPrincipalAccessRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
