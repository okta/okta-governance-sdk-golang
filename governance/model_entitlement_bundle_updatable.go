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

// EntitlementBundleUpdatable struct for EntitlementBundleUpdatable
type EntitlementBundleUpdatable struct {
	// Unique identifier for the object
	Id string `json:"id"`
	// The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	TargetResourceOrn string                   `json:"targetResourceOrn"`
	Target            TargetResource           `json:"target"`
	Status            *EntitlementBundleStatus `json:"status,omitempty"`
	// Collection of entitlements and associated value identifiers
	Entitlements []EntitlementCreatable `json:"entitlements,omitempty"`
	// The unique name of the entitlement bundle
	Name string `json:"name"`
	// The human-readable description
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementBundleUpdatable EntitlementBundleUpdatable

// NewEntitlementBundleUpdatable instantiates a new EntitlementBundleUpdatable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementBundleUpdatable(id string, targetResourceOrn string, target TargetResource, name string) *EntitlementBundleUpdatable {
	this := EntitlementBundleUpdatable{}
	this.Name = name
	return &this
}

// NewEntitlementBundleUpdatableWithDefaults instantiates a new EntitlementBundleUpdatable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementBundleUpdatableWithDefaults() *EntitlementBundleUpdatable {
	this := EntitlementBundleUpdatable{}
	return &this
}

// GetId returns the Id field value
func (o *EntitlementBundleUpdatable) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EntitlementBundleUpdatable) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EntitlementBundleUpdatable) SetId(v string) {
	o.Id = v
}

// GetTargetResourceOrn returns the TargetResourceOrn field value
func (o *EntitlementBundleUpdatable) GetTargetResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetResourceOrn
}

// GetTargetResourceOrnOk returns a tuple with the TargetResourceOrn field value
// and a boolean to check if the value has been set.
func (o *EntitlementBundleUpdatable) GetTargetResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetResourceOrn, true
}

// SetTargetResourceOrn sets field value
func (o *EntitlementBundleUpdatable) SetTargetResourceOrn(v string) {
	o.TargetResourceOrn = v
}

// GetTarget returns the Target field value
func (o *EntitlementBundleUpdatable) GetTarget() TargetResource {
	if o == nil {
		var ret TargetResource
		return ret
	}

	return o.Target
}

// GetTargetOk returns a tuple with the Target field value
// and a boolean to check if the value has been set.
func (o *EntitlementBundleUpdatable) GetTargetOk() (*TargetResource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Target, true
}

// SetTarget sets field value
func (o *EntitlementBundleUpdatable) SetTarget(v TargetResource) {
	o.Target = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *EntitlementBundleUpdatable) GetStatus() EntitlementBundleStatus {
	if o == nil || o.Status == nil {
		var ret EntitlementBundleStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementBundleUpdatable) GetStatusOk() (*EntitlementBundleStatus, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *EntitlementBundleUpdatable) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given EntitlementBundleStatus and assigns it to the Status field.
func (o *EntitlementBundleUpdatable) SetStatus(v EntitlementBundleStatus) {
	o.Status = &v
}

// GetEntitlements returns the Entitlements field value if set, zero value otherwise.
func (o *EntitlementBundleUpdatable) GetEntitlements() []EntitlementCreatable {
	if o == nil || o.Entitlements == nil {
		var ret []EntitlementCreatable
		return ret
	}
	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementBundleUpdatable) GetEntitlementsOk() ([]EntitlementCreatable, bool) {
	if o == nil || o.Entitlements == nil {
		return nil, false
	}
	return o.Entitlements, true
}

// HasEntitlements returns a boolean if a field has been set.
func (o *EntitlementBundleUpdatable) HasEntitlements() bool {
	if o != nil && o.Entitlements != nil {
		return true
	}

	return false
}

// SetEntitlements gets a reference to the given []EntitlementCreatable and assigns it to the Entitlements field.
func (o *EntitlementBundleUpdatable) SetEntitlements(v []EntitlementCreatable) {
	o.Entitlements = v
}

// GetName returns the Name field value
func (o *EntitlementBundleUpdatable) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EntitlementBundleUpdatable) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EntitlementBundleUpdatable) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *EntitlementBundleUpdatable) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementBundleUpdatable) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *EntitlementBundleUpdatable) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *EntitlementBundleUpdatable) SetDescription(v string) {
	o.Description = &v
}

func (o EntitlementBundleUpdatable) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["targetResourceOrn"] = o.TargetResourceOrn
	}
	if true {
		toSerialize["target"] = o.Target
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Entitlements != nil {
		toSerialize["entitlements"] = o.Entitlements
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EntitlementBundleUpdatable) UnmarshalJSON(bytes []byte) (err error) {
	varEntitlementBundleUpdatable := _EntitlementBundleUpdatable{}

	err = json.Unmarshal(bytes, &varEntitlementBundleUpdatable)
	if err == nil {
		*o = EntitlementBundleUpdatable(varEntitlementBundleUpdatable)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "targetResourceOrn")
		delete(additionalProperties, "target")
		delete(additionalProperties, "status")
		delete(additionalProperties, "entitlements")
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableEntitlementBundleUpdatable struct {
	value *EntitlementBundleUpdatable
	isSet bool
}

func (v NullableEntitlementBundleUpdatable) Get() *EntitlementBundleUpdatable {
	return v.value
}

func (v *NullableEntitlementBundleUpdatable) Set(val *EntitlementBundleUpdatable) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementBundleUpdatable) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementBundleUpdatable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementBundleUpdatable(val *EntitlementBundleUpdatable) *NullableEntitlementBundleUpdatable {
	return &NullableEntitlementBundleUpdatable{value: val, isSet: true}
}

func (v NullableEntitlementBundleUpdatable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementBundleUpdatable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
