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

// PrincipalEntitlementsHistory struct for PrincipalEntitlementsHistory
type PrincipalEntitlementsHistory struct {
	// The Okta app instance, in [ORN format](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn).  See the ORN format for a specific app in [Supported resouces](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	ResourceOrn *string         `json:"resourceOrn,omitempty"`
	Resource    *TargetResource `json:"resource,omitempty"`
	// The Okta user `id` in [ORN](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#okta-resource-name-orn) format.  See [Supported resources](https://developer.okta.com/docs/api/openapi/okta-management/guides/roles/#supported-resources).
	PrincipalOrn *string              `json:"principalOrn,omitempty"`
	Principal    *TargetPrincipalFull `json:"principal,omitempty"`
	// Principal entitlements history list
	EntitlementHistory   []EntitlementHistoryRecord `json:"entitlementHistory,omitempty"`
	Links                *ListLinks                 `json:"_links,omitempty"`
	Metadata             *ListMetadata              `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PrincipalEntitlementsHistory PrincipalEntitlementsHistory

// NewPrincipalEntitlementsHistory instantiates a new PrincipalEntitlementsHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPrincipalEntitlementsHistory() *PrincipalEntitlementsHistory {
	this := PrincipalEntitlementsHistory{}
	return &this
}

// NewPrincipalEntitlementsHistoryWithDefaults instantiates a new PrincipalEntitlementsHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPrincipalEntitlementsHistoryWithDefaults() *PrincipalEntitlementsHistory {
	this := PrincipalEntitlementsHistory{}
	return &this
}

// GetResourceOrn returns the ResourceOrn field value if set, zero value otherwise.
func (o *PrincipalEntitlementsHistory) GetResourceOrn() string {
	if o == nil || o.ResourceOrn == nil {
		var ret string
		return ret
	}
	return *o.ResourceOrn
}

// GetResourceOrnOk returns a tuple with the ResourceOrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementsHistory) GetResourceOrnOk() (*string, bool) {
	if o == nil || o.ResourceOrn == nil {
		return nil, false
	}
	return o.ResourceOrn, true
}

// HasResourceOrn returns a boolean if a field has been set.
func (o *PrincipalEntitlementsHistory) HasResourceOrn() bool {
	if o != nil && o.ResourceOrn != nil {
		return true
	}

	return false
}

// SetResourceOrn gets a reference to the given string and assigns it to the ResourceOrn field.
func (o *PrincipalEntitlementsHistory) SetResourceOrn(v string) {
	o.ResourceOrn = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *PrincipalEntitlementsHistory) GetResource() TargetResource {
	if o == nil || o.Resource == nil {
		var ret TargetResource
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementsHistory) GetResourceOk() (*TargetResource, bool) {
	if o == nil || o.Resource == nil {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *PrincipalEntitlementsHistory) HasResource() bool {
	if o != nil && o.Resource != nil {
		return true
	}

	return false
}

// SetResource gets a reference to the given TargetResource and assigns it to the Resource field.
func (o *PrincipalEntitlementsHistory) SetResource(v TargetResource) {
	o.Resource = &v
}

// GetPrincipalOrn returns the PrincipalOrn field value if set, zero value otherwise.
func (o *PrincipalEntitlementsHistory) GetPrincipalOrn() string {
	if o == nil || o.PrincipalOrn == nil {
		var ret string
		return ret
	}
	return *o.PrincipalOrn
}

// GetPrincipalOrnOk returns a tuple with the PrincipalOrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementsHistory) GetPrincipalOrnOk() (*string, bool) {
	if o == nil || o.PrincipalOrn == nil {
		return nil, false
	}
	return o.PrincipalOrn, true
}

// HasPrincipalOrn returns a boolean if a field has been set.
func (o *PrincipalEntitlementsHistory) HasPrincipalOrn() bool {
	if o != nil && o.PrincipalOrn != nil {
		return true
	}

	return false
}

// SetPrincipalOrn gets a reference to the given string and assigns it to the PrincipalOrn field.
func (o *PrincipalEntitlementsHistory) SetPrincipalOrn(v string) {
	o.PrincipalOrn = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *PrincipalEntitlementsHistory) GetPrincipal() TargetPrincipalFull {
	if o == nil || o.Principal == nil {
		var ret TargetPrincipalFull
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementsHistory) GetPrincipalOk() (*TargetPrincipalFull, bool) {
	if o == nil || o.Principal == nil {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *PrincipalEntitlementsHistory) HasPrincipal() bool {
	if o != nil && o.Principal != nil {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given TargetPrincipalFull and assigns it to the Principal field.
func (o *PrincipalEntitlementsHistory) SetPrincipal(v TargetPrincipalFull) {
	o.Principal = &v
}

// GetEntitlementHistory returns the EntitlementHistory field value if set, zero value otherwise.
func (o *PrincipalEntitlementsHistory) GetEntitlementHistory() []EntitlementHistoryRecord {
	if o == nil || o.EntitlementHistory == nil {
		var ret []EntitlementHistoryRecord
		return ret
	}
	return o.EntitlementHistory
}

// GetEntitlementHistoryOk returns a tuple with the EntitlementHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementsHistory) GetEntitlementHistoryOk() ([]EntitlementHistoryRecord, bool) {
	if o == nil || o.EntitlementHistory == nil {
		return nil, false
	}
	return o.EntitlementHistory, true
}

// HasEntitlementHistory returns a boolean if a field has been set.
func (o *PrincipalEntitlementsHistory) HasEntitlementHistory() bool {
	if o != nil && o.EntitlementHistory != nil {
		return true
	}

	return false
}

// SetEntitlementHistory gets a reference to the given []EntitlementHistoryRecord and assigns it to the EntitlementHistory field.
func (o *PrincipalEntitlementsHistory) SetEntitlementHistory(v []EntitlementHistoryRecord) {
	o.EntitlementHistory = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *PrincipalEntitlementsHistory) GetLinks() ListLinks {
	if o == nil || o.Links == nil {
		var ret ListLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementsHistory) GetLinksOk() (*ListLinks, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *PrincipalEntitlementsHistory) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ListLinks and assigns it to the Links field.
func (o *PrincipalEntitlementsHistory) SetLinks(v ListLinks) {
	o.Links = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *PrincipalEntitlementsHistory) GetMetadata() ListMetadata {
	if o == nil || o.Metadata == nil {
		var ret ListMetadata
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PrincipalEntitlementsHistory) GetMetadataOk() (*ListMetadata, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *PrincipalEntitlementsHistory) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ListMetadata and assigns it to the Metadata field.
func (o *PrincipalEntitlementsHistory) SetMetadata(v ListMetadata) {
	o.Metadata = &v
}

func (o PrincipalEntitlementsHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResourceOrn != nil {
		toSerialize["resourceOrn"] = o.ResourceOrn
	}
	if o.Resource != nil {
		toSerialize["resource"] = o.Resource
	}
	if o.PrincipalOrn != nil {
		toSerialize["principalOrn"] = o.PrincipalOrn
	}
	if o.Principal != nil {
		toSerialize["principal"] = o.Principal
	}
	if o.EntitlementHistory != nil {
		toSerialize["entitlementHistory"] = o.EntitlementHistory
	}
	if o.Links != nil {
		toSerialize["_links"] = o.Links
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PrincipalEntitlementsHistory) UnmarshalJSON(bytes []byte) (err error) {
	varPrincipalEntitlementsHistory := _PrincipalEntitlementsHistory{}

	err = json.Unmarshal(bytes, &varPrincipalEntitlementsHistory)
	if err == nil {
		*o = PrincipalEntitlementsHistory(varPrincipalEntitlementsHistory)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "resourceOrn")
		delete(additionalProperties, "resource")
		delete(additionalProperties, "principalOrn")
		delete(additionalProperties, "principal")
		delete(additionalProperties, "entitlementHistory")
		delete(additionalProperties, "_links")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullablePrincipalEntitlementsHistory struct {
	value *PrincipalEntitlementsHistory
	isSet bool
}

func (v NullablePrincipalEntitlementsHistory) Get() *PrincipalEntitlementsHistory {
	return v.value
}

func (v *NullablePrincipalEntitlementsHistory) Set(val *PrincipalEntitlementsHistory) {
	v.value = val
	v.isSet = true
}

func (v NullablePrincipalEntitlementsHistory) IsSet() bool {
	return v.isSet
}

func (v *NullablePrincipalEntitlementsHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrincipalEntitlementsHistory(val *PrincipalEntitlementsHistory) *NullablePrincipalEntitlementsHistory {
	return &NullablePrincipalEntitlementsHistory{value: val, isSet: true}
}

func (v NullablePrincipalEntitlementsHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrincipalEntitlementsHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
