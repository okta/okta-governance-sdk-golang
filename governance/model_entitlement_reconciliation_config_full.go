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
	"time"
)

// checks if the EntitlementReconciliationConfigFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EntitlementReconciliationConfigFull{}

// EntitlementReconciliationConfigFull A resource's reconciliation configuration.  `additive` covers entitlements added directly in the app, the drifts reported with `driftType` of `ADD`. `subtractive` covers entitlements removed there, reported as `SUB`. Those two directions are exhaustive over `driftType`.  A response describes one of three states.  Unconfigured. The resource has no saved configuration, so reconciliation is off. `mode` is `DISABLED`, and only `resourceOrn` and `_links` come with it. There's no `id`, no audit fields, and no `additive` or `subtractive`.  Saved and disabled. A configuration exists but reconciliation is off. `mode` is `DISABLED`, and `id` and the audit fields are present. `additive` and `subtractive` are echoed back whenever the saved configuration has them, so an admin UI can repopulate the form without losing what was configured before it was turned off. Either one can be absent on its own, because an upsert while `mode` is `DISABLED` may send just one of them.  Saved and enabled. `mode` is `ENABLED`, and `additive` and `subtractive` are both present.  Presence of `id` is the only thing separating unconfigured from saved and disabled, because both report `mode` as `DISABLED`. Use it to tell them apart.  A successful `upsertEntitlementReconciliationConfig` always returns one of the two saved states.
type EntitlementReconciliationConfigFull struct {
	// The unique ID of the configuration. Present only after a configuration has been saved for the resource.
	Id *string `json:"id,omitempty"`
	// The `id` of the Okta user who created the resource
	CreatedBy *string `json:"createdBy,omitempty"`
	// The ISO 8601 formatted date and time when the resource was created
	Created *time.Time `json:"created,omitempty"`
	// The ISO 8601 formatted date and time when the object was last updated
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	// The `id` of the Okta user who last updated the object
	LastUpdatedBy *string                                 `json:"lastUpdatedBy,omitempty"`
	Mode          ReconciliationMode                      `json:"mode"`
	Additive      *EntitlementReconciliationDirectionFull `json:"additive,omitempty"`
	Subtractive   *EntitlementReconciliationDirectionFull `json:"subtractive,omitempty"`
	// The Okta resource in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn)  See the ORN format for [supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	ResourceOrn          string                               `json:"resourceOrn"`
	Links                EntitlementReconciliationConfigLinks `json:"_links"`
	AdditionalProperties map[string]interface{}
}

type _EntitlementReconciliationConfigFull EntitlementReconciliationConfigFull

// NewEntitlementReconciliationConfigFull instantiates a new EntitlementReconciliationConfigFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlementReconciliationConfigFull(mode ReconciliationMode, resourceOrn string, links EntitlementReconciliationConfigLinks) *EntitlementReconciliationConfigFull {
	this := EntitlementReconciliationConfigFull{}
	this.Mode = mode
	this.ResourceOrn = resourceOrn
	this.Links = links
	return &this
}

// NewEntitlementReconciliationConfigFullWithDefaults instantiates a new EntitlementReconciliationConfigFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementReconciliationConfigFullWithDefaults() *EntitlementReconciliationConfigFull {
	this := EntitlementReconciliationConfigFull{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EntitlementReconciliationConfigFull) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationConfigFull) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EntitlementReconciliationConfigFull) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EntitlementReconciliationConfigFull) SetId(v string) {
	o.Id = &v
}

// GetCreatedBy returns the CreatedBy field value if set, zero value otherwise.
func (o *EntitlementReconciliationConfigFull) GetCreatedBy() string {
	if o == nil || IsNil(o.CreatedBy) {
		var ret string
		return ret
	}
	return *o.CreatedBy
}

// GetCreatedByOk returns a tuple with the CreatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationConfigFull) GetCreatedByOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedBy) {
		return nil, false
	}
	return o.CreatedBy, true
}

// HasCreatedBy returns a boolean if a field has been set.
func (o *EntitlementReconciliationConfigFull) HasCreatedBy() bool {
	if o != nil && !IsNil(o.CreatedBy) {
		return true
	}

	return false
}

// SetCreatedBy gets a reference to the given string and assigns it to the CreatedBy field.
func (o *EntitlementReconciliationConfigFull) SetCreatedBy(v string) {
	o.CreatedBy = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *EntitlementReconciliationConfigFull) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationConfigFull) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *EntitlementReconciliationConfigFull) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *EntitlementReconciliationConfigFull) SetCreated(v time.Time) {
	o.Created = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *EntitlementReconciliationConfigFull) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationConfigFull) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *EntitlementReconciliationConfigFull) HasLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *EntitlementReconciliationConfigFull) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetLastUpdatedBy returns the LastUpdatedBy field value if set, zero value otherwise.
func (o *EntitlementReconciliationConfigFull) GetLastUpdatedBy() string {
	if o == nil || IsNil(o.LastUpdatedBy) {
		var ret string
		return ret
	}
	return *o.LastUpdatedBy
}

// GetLastUpdatedByOk returns a tuple with the LastUpdatedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationConfigFull) GetLastUpdatedByOk() (*string, bool) {
	if o == nil || IsNil(o.LastUpdatedBy) {
		return nil, false
	}
	return o.LastUpdatedBy, true
}

// HasLastUpdatedBy returns a boolean if a field has been set.
func (o *EntitlementReconciliationConfigFull) HasLastUpdatedBy() bool {
	if o != nil && !IsNil(o.LastUpdatedBy) {
		return true
	}

	return false
}

// SetLastUpdatedBy gets a reference to the given string and assigns it to the LastUpdatedBy field.
func (o *EntitlementReconciliationConfigFull) SetLastUpdatedBy(v string) {
	o.LastUpdatedBy = &v
}

// GetMode returns the Mode field value
func (o *EntitlementReconciliationConfigFull) GetMode() ReconciliationMode {
	if o == nil {
		var ret ReconciliationMode
		return ret
	}

	return o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationConfigFull) GetModeOk() (*ReconciliationMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mode, true
}

// SetMode sets field value
func (o *EntitlementReconciliationConfigFull) SetMode(v ReconciliationMode) {
	o.Mode = v
}

// GetAdditive returns the Additive field value if set, zero value otherwise.
func (o *EntitlementReconciliationConfigFull) GetAdditive() EntitlementReconciliationDirectionFull {
	if o == nil || IsNil(o.Additive) {
		var ret EntitlementReconciliationDirectionFull
		return ret
	}
	return *o.Additive
}

// GetAdditiveOk returns a tuple with the Additive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationConfigFull) GetAdditiveOk() (*EntitlementReconciliationDirectionFull, bool) {
	if o == nil || IsNil(o.Additive) {
		return nil, false
	}
	return o.Additive, true
}

// HasAdditive returns a boolean if a field has been set.
func (o *EntitlementReconciliationConfigFull) HasAdditive() bool {
	if o != nil && !IsNil(o.Additive) {
		return true
	}

	return false
}

// SetAdditive gets a reference to the given EntitlementReconciliationDirectionFull and assigns it to the Additive field.
func (o *EntitlementReconciliationConfigFull) SetAdditive(v EntitlementReconciliationDirectionFull) {
	o.Additive = &v
}

// GetSubtractive returns the Subtractive field value if set, zero value otherwise.
func (o *EntitlementReconciliationConfigFull) GetSubtractive() EntitlementReconciliationDirectionFull {
	if o == nil || IsNil(o.Subtractive) {
		var ret EntitlementReconciliationDirectionFull
		return ret
	}
	return *o.Subtractive
}

// GetSubtractiveOk returns a tuple with the Subtractive field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationConfigFull) GetSubtractiveOk() (*EntitlementReconciliationDirectionFull, bool) {
	if o == nil || IsNil(o.Subtractive) {
		return nil, false
	}
	return o.Subtractive, true
}

// HasSubtractive returns a boolean if a field has been set.
func (o *EntitlementReconciliationConfigFull) HasSubtractive() bool {
	if o != nil && !IsNil(o.Subtractive) {
		return true
	}

	return false
}

// SetSubtractive gets a reference to the given EntitlementReconciliationDirectionFull and assigns it to the Subtractive field.
func (o *EntitlementReconciliationConfigFull) SetSubtractive(v EntitlementReconciliationDirectionFull) {
	o.Subtractive = &v
}

// GetResourceOrn returns the ResourceOrn field value
func (o *EntitlementReconciliationConfigFull) GetResourceOrn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResourceOrn
}

// GetResourceOrnOk returns a tuple with the ResourceOrn field value
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationConfigFull) GetResourceOrnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResourceOrn, true
}

// SetResourceOrn sets field value
func (o *EntitlementReconciliationConfigFull) SetResourceOrn(v string) {
	o.ResourceOrn = v
}

// GetLinks returns the Links field value
func (o *EntitlementReconciliationConfigFull) GetLinks() EntitlementReconciliationConfigLinks {
	if o == nil {
		var ret EntitlementReconciliationConfigLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *EntitlementReconciliationConfigFull) GetLinksOk() (*EntitlementReconciliationConfigLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *EntitlementReconciliationConfigFull) SetLinks(v EntitlementReconciliationConfigLinks) {
	o.Links = v
}

func (o EntitlementReconciliationConfigFull) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EntitlementReconciliationConfigFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CreatedBy) {
		toSerialize["createdBy"] = o.CreatedBy
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.LastUpdatedBy) {
		toSerialize["lastUpdatedBy"] = o.LastUpdatedBy
	}
	toSerialize["mode"] = o.Mode
	if !IsNil(o.Additive) {
		toSerialize["additive"] = o.Additive
	}
	if !IsNil(o.Subtractive) {
		toSerialize["subtractive"] = o.Subtractive
	}
	toSerialize["resourceOrn"] = o.ResourceOrn
	toSerialize["_links"] = o.Links

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *EntitlementReconciliationConfigFull) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mode",
		"resourceOrn",
		"_links",
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

	varEntitlementReconciliationConfigFull := _EntitlementReconciliationConfigFull{}

	err = json.Unmarshal(data, &varEntitlementReconciliationConfigFull)

	if err != nil {
		return err
	}

	*o = EntitlementReconciliationConfigFull(varEntitlementReconciliationConfigFull)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "createdBy")
		delete(additionalProperties, "created")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "lastUpdatedBy")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "additive")
		delete(additionalProperties, "subtractive")
		delete(additionalProperties, "resourceOrn")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEntitlementReconciliationConfigFull struct {
	value *EntitlementReconciliationConfigFull
	isSet bool
}

func (v NullableEntitlementReconciliationConfigFull) Get() *EntitlementReconciliationConfigFull {
	return v.value
}

func (v *NullableEntitlementReconciliationConfigFull) Set(val *EntitlementReconciliationConfigFull) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlementReconciliationConfigFull) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlementReconciliationConfigFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlementReconciliationConfigFull(val *EntitlementReconciliationConfigFull) *NullableEntitlementReconciliationConfigFull {
	return &NullableEntitlementReconciliationConfigFull{value: val, isSet: true}
}

func (v NullableEntitlementReconciliationConfigFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlementReconciliationConfigFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
