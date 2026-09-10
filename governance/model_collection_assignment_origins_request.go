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

// checks if the CollectionAssignmentOriginsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionAssignmentOriginsRequest{}

// CollectionAssignmentOriginsRequest Assignments to look up. Provide exactly one anchor: either a single principal (principalOrn + resourceOrns) or a single resource (resourceOrn + principalOrns). A request that sets both anchors or neither is rejected.
type CollectionAssignmentOriginsRequest struct {
	// Anchor principal. Provide together with resourceOrns.
	PrincipalOrn *string `json:"principalOrn,omitempty"`
	// Resource ORNs to check for the anchor principal
	ResourceOrns []string `json:"resourceOrns,omitempty"`
	// Anchor resource. Provide together with principalOrns.
	ResourceOrn *string `json:"resourceOrn,omitempty"`
	// Principal ORNs to check for the anchor resource
	PrincipalOrns        []string `json:"principalOrns,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CollectionAssignmentOriginsRequest CollectionAssignmentOriginsRequest

// NewCollectionAssignmentOriginsRequest instantiates a new CollectionAssignmentOriginsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionAssignmentOriginsRequest() *CollectionAssignmentOriginsRequest {
	this := CollectionAssignmentOriginsRequest{}
	return &this
}

// NewCollectionAssignmentOriginsRequestWithDefaults instantiates a new CollectionAssignmentOriginsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionAssignmentOriginsRequestWithDefaults() *CollectionAssignmentOriginsRequest {
	this := CollectionAssignmentOriginsRequest{}
	return &this
}

// GetPrincipalOrn returns the PrincipalOrn field value if set, zero value otherwise.
func (o *CollectionAssignmentOriginsRequest) GetPrincipalOrn() string {
	if o == nil || IsNil(o.PrincipalOrn) {
		var ret string
		return ret
	}
	return *o.PrincipalOrn
}

// GetPrincipalOrnOk returns a tuple with the PrincipalOrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionAssignmentOriginsRequest) GetPrincipalOrnOk() (*string, bool) {
	if o == nil || IsNil(o.PrincipalOrn) {
		return nil, false
	}
	return o.PrincipalOrn, true
}

// HasPrincipalOrn returns a boolean if a field has been set.
func (o *CollectionAssignmentOriginsRequest) HasPrincipalOrn() bool {
	if o != nil && !IsNil(o.PrincipalOrn) {
		return true
	}

	return false
}

// SetPrincipalOrn gets a reference to the given string and assigns it to the PrincipalOrn field.
func (o *CollectionAssignmentOriginsRequest) SetPrincipalOrn(v string) {
	o.PrincipalOrn = &v
}

// GetResourceOrns returns the ResourceOrns field value if set, zero value otherwise.
func (o *CollectionAssignmentOriginsRequest) GetResourceOrns() []string {
	if o == nil || IsNil(o.ResourceOrns) {
		var ret []string
		return ret
	}
	return o.ResourceOrns
}

// GetResourceOrnsOk returns a tuple with the ResourceOrns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionAssignmentOriginsRequest) GetResourceOrnsOk() ([]string, bool) {
	if o == nil || IsNil(o.ResourceOrns) {
		return nil, false
	}
	return o.ResourceOrns, true
}

// HasResourceOrns returns a boolean if a field has been set.
func (o *CollectionAssignmentOriginsRequest) HasResourceOrns() bool {
	if o != nil && !IsNil(o.ResourceOrns) {
		return true
	}

	return false
}

// SetResourceOrns gets a reference to the given []string and assigns it to the ResourceOrns field.
func (o *CollectionAssignmentOriginsRequest) SetResourceOrns(v []string) {
	o.ResourceOrns = v
}

// GetResourceOrn returns the ResourceOrn field value if set, zero value otherwise.
func (o *CollectionAssignmentOriginsRequest) GetResourceOrn() string {
	if o == nil || IsNil(o.ResourceOrn) {
		var ret string
		return ret
	}
	return *o.ResourceOrn
}

// GetResourceOrnOk returns a tuple with the ResourceOrn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionAssignmentOriginsRequest) GetResourceOrnOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceOrn) {
		return nil, false
	}
	return o.ResourceOrn, true
}

// HasResourceOrn returns a boolean if a field has been set.
func (o *CollectionAssignmentOriginsRequest) HasResourceOrn() bool {
	if o != nil && !IsNil(o.ResourceOrn) {
		return true
	}

	return false
}

// SetResourceOrn gets a reference to the given string and assigns it to the ResourceOrn field.
func (o *CollectionAssignmentOriginsRequest) SetResourceOrn(v string) {
	o.ResourceOrn = &v
}

// GetPrincipalOrns returns the PrincipalOrns field value if set, zero value otherwise.
func (o *CollectionAssignmentOriginsRequest) GetPrincipalOrns() []string {
	if o == nil || IsNil(o.PrincipalOrns) {
		var ret []string
		return ret
	}
	return o.PrincipalOrns
}

// GetPrincipalOrnsOk returns a tuple with the PrincipalOrns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionAssignmentOriginsRequest) GetPrincipalOrnsOk() ([]string, bool) {
	if o == nil || IsNil(o.PrincipalOrns) {
		return nil, false
	}
	return o.PrincipalOrns, true
}

// HasPrincipalOrns returns a boolean if a field has been set.
func (o *CollectionAssignmentOriginsRequest) HasPrincipalOrns() bool {
	if o != nil && !IsNil(o.PrincipalOrns) {
		return true
	}

	return false
}

// SetPrincipalOrns gets a reference to the given []string and assigns it to the PrincipalOrns field.
func (o *CollectionAssignmentOriginsRequest) SetPrincipalOrns(v []string) {
	o.PrincipalOrns = v
}

func (o CollectionAssignmentOriginsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionAssignmentOriginsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PrincipalOrn) {
		toSerialize["principalOrn"] = o.PrincipalOrn
	}
	if !IsNil(o.ResourceOrns) {
		toSerialize["resourceOrns"] = o.ResourceOrns
	}
	if !IsNil(o.ResourceOrn) {
		toSerialize["resourceOrn"] = o.ResourceOrn
	}
	if !IsNil(o.PrincipalOrns) {
		toSerialize["principalOrns"] = o.PrincipalOrns
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *CollectionAssignmentOriginsRequest) UnmarshalJSON(data []byte) (err error) {
	varCollectionAssignmentOriginsRequest := _CollectionAssignmentOriginsRequest{}

	err = json.Unmarshal(data, &varCollectionAssignmentOriginsRequest)

	if err != nil {
		return err
	}

	*o = CollectionAssignmentOriginsRequest(varCollectionAssignmentOriginsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "principalOrn")
		delete(additionalProperties, "resourceOrns")
		delete(additionalProperties, "resourceOrn")
		delete(additionalProperties, "principalOrns")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCollectionAssignmentOriginsRequest struct {
	value *CollectionAssignmentOriginsRequest
	isSet bool
}

func (v NullableCollectionAssignmentOriginsRequest) Get() *CollectionAssignmentOriginsRequest {
	return v.value
}

func (v *NullableCollectionAssignmentOriginsRequest) Set(val *CollectionAssignmentOriginsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionAssignmentOriginsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionAssignmentOriginsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionAssignmentOriginsRequest(val *CollectionAssignmentOriginsRequest) *NullableCollectionAssignmentOriginsRequest {
	return &NullableCollectionAssignmentOriginsRequest{value: val, isSet: true}
}

func (v NullableCollectionAssignmentOriginsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionAssignmentOriginsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
