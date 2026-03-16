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

// checks if the PrincipalEntitlementsChange type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PrincipalEntitlementsChange{}

// PrincipalEntitlementsChange struct for PrincipalEntitlementsChange
type PrincipalEntitlementsChange struct {
	// List of changed entitlements
	EntitlementsChanged []EntitlementChangedFull `json:"entitlementsChanged,omitempty"`
	// The Okta resource, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for [supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	ResourceOrn *string         `json:"resourceOrn,omitempty"`
	Resource    *TargetResource `json:"resource,omitempty"`
	// The Okta user, in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format.
	PrincipalOrn         *string                           `json:"principalOrn,omitempty"`
	Principal            *TargetPrincipalFull              `json:"principal,omitempty"`
	Links                *PrincipalEntitlementsChangeLinks `json:"_links,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrincipalEntitlementsChange PrincipalEntitlementsChange

// NewPrincipalEntitlementsChange instantiates a new PrincipalEntitlementsChange object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalEntitlementsChange() *PrincipalEntitlementsChange {
	this := PrincipalEntitlementsChange{}
	return &this
}

// NewPrincipalEntitlementsChangeWithDefaults instantiates a new PrincipalEntitlementsChange object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalEntitlementsChangeWithDefaults() *PrincipalEntitlementsChange {
	this := PrincipalEntitlementsChange{}
	return &this
}

// GetEntitlementsChanged returns the EntitlementsChanged field value if set, zero value otherwise.
func (o *PrincipalEntitlementsChange) GetEntitlementsChanged() []EntitlementChangedFull {
	if o == nil || IsNil(o.EntitlementsChanged) {
		var ret []EntitlementChangedFull
		return ret
	}
	return o.EntitlementsChanged
}

// GetEntitlementsChangedOk returns a tuple with the EntitlementsChanged field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementsChange) GetEntitlementsChangedOk() ([]EntitlementChangedFull, bool) {
	if o == nil || IsNil(o.EntitlementsChanged) {
		return nil, false
	}
	return o.EntitlementsChanged, true
}

// HasEntitlementsChanged returns a boolean if a field has been set.
func (o *PrincipalEntitlementsChange) HasEntitlementsChanged() bool {
	if o != nil && !IsNil(o.EntitlementsChanged) {
		return true
	}

	return false
}

// SetEntitlementsChanged gets a reference to the given []EntitlementChangedFull and assigns it to the EntitlementsChanged field.
func (o *PrincipalEntitlementsChange) SetEntitlementsChanged(v []EntitlementChangedFull) {
	o.EntitlementsChanged = v
}

// GetResourceOrn returns the ResourceOrn field value if set, zero value otherwise.
func (o *PrincipalEntitlementsChange) GetResourceOrn() string {
	if o == nil || IsNil(o.ResourceOrn) {
		var ret string
		return ret
	}
	return *o.ResourceOrn
}

// GetResourceOrnOk returns a tuple with the ResourceOrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementsChange) GetResourceOrnOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceOrn) {
		return nil, false
	}
	return o.ResourceOrn, true
}

// HasResourceOrn returns a boolean if a field has been set.
func (o *PrincipalEntitlementsChange) HasResourceOrn() bool {
	if o != nil && !IsNil(o.ResourceOrn) {
		return true
	}

	return false
}

// SetResourceOrn gets a reference to the given string and assigns it to the ResourceOrn field.
func (o *PrincipalEntitlementsChange) SetResourceOrn(v string) {
	o.ResourceOrn = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *PrincipalEntitlementsChange) GetResource() TargetResource {
	if o == nil || IsNil(o.Resource) {
		var ret TargetResource
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementsChange) GetResourceOk() (*TargetResource, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *PrincipalEntitlementsChange) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given TargetResource and assigns it to the Resource field.
func (o *PrincipalEntitlementsChange) SetResource(v TargetResource) {
	o.Resource = &v
}

// GetPrincipalOrn returns the PrincipalOrn field value if set, zero value otherwise.
func (o *PrincipalEntitlementsChange) GetPrincipalOrn() string {
	if o == nil || IsNil(o.PrincipalOrn) {
		var ret string
		return ret
	}
	return *o.PrincipalOrn
}

// GetPrincipalOrnOk returns a tuple with the PrincipalOrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementsChange) GetPrincipalOrnOk() (*string, bool) {
	if o == nil || IsNil(o.PrincipalOrn) {
		return nil, false
	}
	return o.PrincipalOrn, true
}

// HasPrincipalOrn returns a boolean if a field has been set.
func (o *PrincipalEntitlementsChange) HasPrincipalOrn() bool {
	if o != nil && !IsNil(o.PrincipalOrn) {
		return true
	}

	return false
}

// SetPrincipalOrn gets a reference to the given string and assigns it to the PrincipalOrn field.
func (o *PrincipalEntitlementsChange) SetPrincipalOrn(v string) {
	o.PrincipalOrn = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *PrincipalEntitlementsChange) GetPrincipal() TargetPrincipalFull {
	if o == nil || IsNil(o.Principal) {
		var ret TargetPrincipalFull
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementsChange) GetPrincipalOk() (*TargetPrincipalFull, bool) {
	if o == nil || IsNil(o.Principal) {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *PrincipalEntitlementsChange) HasPrincipal() bool {
	if o != nil && !IsNil(o.Principal) {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given TargetPrincipalFull and assigns it to the Principal field.
func (o *PrincipalEntitlementsChange) SetPrincipal(v TargetPrincipalFull) {
	o.Principal = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PrincipalEntitlementsChange) GetLinks() PrincipalEntitlementsChangeLinks {
	if o == nil || IsNil(o.Links) {
		var ret PrincipalEntitlementsChangeLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementsChange) GetLinksOk() (*PrincipalEntitlementsChangeLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PrincipalEntitlementsChange) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given PrincipalEntitlementsChangeLinks and assigns it to the Links field.
func (o *PrincipalEntitlementsChange) SetLinks(v PrincipalEntitlementsChangeLinks) {
	o.Links = &v
}

func (o PrincipalEntitlementsChange) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PrincipalEntitlementsChange) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EntitlementsChanged) {
		toSerialize["entitlementsChanged"] = o.EntitlementsChanged
	}
	if !IsNil(o.ResourceOrn) {
		toSerialize["resourceOrn"] = o.ResourceOrn
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if !IsNil(o.PrincipalOrn) {
		toSerialize["principalOrn"] = o.PrincipalOrn
	}
	if !IsNil(o.Principal) {
		toSerialize["principal"] = o.Principal
	}
	if !IsNil(o.Links) {
		toSerialize["_links"] = o.Links
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *PrincipalEntitlementsChange) UnmarshalJSON(data []byte) (err error) {
	varPrincipalEntitlementsChange := _PrincipalEntitlementsChange{}

	err = json.Unmarshal(data, &varPrincipalEntitlementsChange)

	if err != nil {
		return err
	}

	*o = PrincipalEntitlementsChange(varPrincipalEntitlementsChange)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "entitlementsChanged")
		delete(additionalProperties, "resourceOrn")
		delete(additionalProperties, "resource")
		delete(additionalProperties, "principalOrn")
		delete(additionalProperties, "principal")
		delete(additionalProperties, "_links")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePrincipalEntitlementsChange struct {
	value *PrincipalEntitlementsChange
	isSet bool
}

func (v NullablePrincipalEntitlementsChange) Get() *PrincipalEntitlementsChange {
	return v.value
}

func (v *NullablePrincipalEntitlementsChange) Set(val *PrincipalEntitlementsChange) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalEntitlementsChange) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalEntitlementsChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalEntitlementsChange(val *PrincipalEntitlementsChange) *NullablePrincipalEntitlementsChange {
	return &NullablePrincipalEntitlementsChange{value: val, isSet: true}
}

func (v NullablePrincipalEntitlementsChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalEntitlementsChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
