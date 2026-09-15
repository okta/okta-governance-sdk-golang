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

// checks if the PrincipalEntitlementRow type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrincipalEntitlementRow{}

// PrincipalEntitlementRow One node of the principal's effective-entitlement view. It's either the resource (app) itself or one of its resource assets.  * `targetOrn` is the stable identifier that paginates `data[]`.  * When the row is a resource asset, `resourceAsset` carries its sparse representation. It's omitted when the row is the resource itself, which happens for resources that don't support assets.  * `entitlements[]` is the principal's complete effective set on the target, calculated on the server side. Clients must not recompute it. `entitlements[]` isn't paginated. Only each entitlement's `values` are paged.
type PrincipalEntitlementRow struct {
	// The Okta resource in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn)  See the ORN format for [supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	TargetOrn     string                                `json:"targetOrn"`
	ResourceAsset *PrincipalEntitlementRowResourceAsset `json:"resourceAsset,omitempty"`
	// The principal's complete effective entitlements on this target, calculated on the server side from all grant sources: direct grants, group-inherited grants, and hierarchy-inherited grants. Each item is one entitlement property with only the first page of its effective values returned.
	Entitlements         []EntitlementWithValues `json:"entitlements"`
	AdditionalProperties map[string]interface{}
}

type _PrincipalEntitlementRow PrincipalEntitlementRow

// NewPrincipalEntitlementRow instantiates a new PrincipalEntitlementRow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalEntitlementRow(targetOrn string, entitlements []EntitlementWithValues) *PrincipalEntitlementRow {
	this := PrincipalEntitlementRow{}
	this.TargetOrn = targetOrn
	this.Entitlements = entitlements
	return &this
}

// NewPrincipalEntitlementRowWithDefaults instantiates a new PrincipalEntitlementRow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalEntitlementRowWithDefaults() *PrincipalEntitlementRow {
	this := PrincipalEntitlementRow{}
	return &this
}

// GetTargetOrn returns the TargetOrn field value
func (o *PrincipalEntitlementRow) GetTargetOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetOrn
}

// GetTargetOrnOk returns a tuple with the TargetOrn field value
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementRow) GetTargetOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetOrn, true
}

// SetTargetOrn sets field value
func (o *PrincipalEntitlementRow) SetTargetOrn(v string) {
	o.TargetOrn = v
}

// GetResourceAsset returns the ResourceAsset field value if set, zero value otherwise.
func (o *PrincipalEntitlementRow) GetResourceAsset() PrincipalEntitlementRowResourceAsset {
	if o == nil || IsNil(o.ResourceAsset) {
		var ret PrincipalEntitlementRowResourceAsset
		return ret
	}
	return *o.ResourceAsset
}

// GetResourceAssetOk returns a tuple with the ResourceAsset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementRow) GetResourceAssetOk() (*PrincipalEntitlementRowResourceAsset, bool) {
	if o == nil || IsNil(o.ResourceAsset) {
		return nil, false
	}
	return o.ResourceAsset, true
}

// HasResourceAsset returns a boolean if a field has been set.
func (o *PrincipalEntitlementRow) HasResourceAsset() bool {
	if o != nil && !IsNil(o.ResourceAsset) {
		return true
	}

	return false
}

// SetResourceAsset gets a reference to the given PrincipalEntitlementRowResourceAsset and assigns it to the ResourceAsset field.
func (o *PrincipalEntitlementRow) SetResourceAsset(v PrincipalEntitlementRowResourceAsset) {
	o.ResourceAsset = &v
}

// GetEntitlements returns the Entitlements field value
func (o *PrincipalEntitlementRow) GetEntitlements() []EntitlementWithValues {
	if o == nil {
		var ret []EntitlementWithValues
		return ret
	}

	return o.Entitlements
}

// GetEntitlementsOk returns a tuple with the Entitlements field value
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementRow) GetEntitlementsOk() ([]EntitlementWithValues, bool) {
	if o == nil {
		return nil, false
	}
	return o.Entitlements, true
}

// SetEntitlements sets field value
func (o *PrincipalEntitlementRow) SetEntitlements(v []EntitlementWithValues) {
	o.Entitlements = v
}

func (o PrincipalEntitlementRow) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrincipalEntitlementRow) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["targetOrn"] = o.TargetOrn
	if !IsNil(o.ResourceAsset) {
		toSerialize["resourceAsset"] = o.ResourceAsset
	}
	toSerialize["entitlements"] = o.Entitlements

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrincipalEntitlementRow) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"targetOrn",
		"entitlements",
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

	varPrincipalEntitlementRow := _PrincipalEntitlementRow{}

	err = json.Unmarshal(data, &varPrincipalEntitlementRow)

	if err != nil {
		return err
	}

	*o = PrincipalEntitlementRow(varPrincipalEntitlementRow)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "targetOrn")
		delete(additionalProperties, "resourceAsset")
		delete(additionalProperties, "entitlements")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrincipalEntitlementRow struct {
	value *PrincipalEntitlementRow
	isSet bool
}

func (v NullablePrincipalEntitlementRow) Get() *PrincipalEntitlementRow {
	return v.value
}

func (v *NullablePrincipalEntitlementRow) Set(val *PrincipalEntitlementRow) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalEntitlementRow) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalEntitlementRow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalEntitlementRow(val *PrincipalEntitlementRow) *NullablePrincipalEntitlementRow {
	return &NullablePrincipalEntitlementRow{value: val, isSet: true}
}

func (v NullablePrincipalEntitlementRow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalEntitlementRow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
