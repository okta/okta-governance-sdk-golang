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

// RequestLinks2 Links available on a single request representation
type RequestLinks2 struct {
	CatalogEntry         *Link `json:"catalogEntry,omitempty"`
	ApprovalCondition    *Link `json:"approvalCondition,omitempty"`
	ApprovalSequence     *Link `json:"approvalSequence,omitempty"`
	Self                 Link  `json:"self"`
	AdditionalProperties map[string]interface{}
}

type _RequestLinks2 RequestLinks2

// NewRequestLinks2 instantiates a new RequestLinks2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestLinks2(self Link) *RequestLinks2 {
	this := RequestLinks2{}
	this.Self = self
	return &this
}

// NewRequestLinks2WithDefaults instantiates a new RequestLinks2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestLinks2WithDefaults() *RequestLinks2 {
	this := RequestLinks2{}
	return &this
}

// GetCatalogEntry returns the CatalogEntry field value if set, zero value otherwise.
func (o *RequestLinks2) GetCatalogEntry() Link {
	if o == nil || o.CatalogEntry == nil {
		var ret Link
		return ret
	}
	return *o.CatalogEntry
}

// GetCatalogEntryOk returns a tuple with the CatalogEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestLinks2) GetCatalogEntryOk() (*Link, bool) {
	if o == nil || o.CatalogEntry == nil {
		return nil, false
	}
	return o.CatalogEntry, true
}

// HasCatalogEntry returns a boolean if a field has been set.
func (o *RequestLinks2) HasCatalogEntry() bool {
	if o != nil && o.CatalogEntry != nil {
		return true
	}

	return false
}

// SetCatalogEntry gets a reference to the given Link and assigns it to the CatalogEntry field.
func (o *RequestLinks2) SetCatalogEntry(v Link) {
	o.CatalogEntry = &v
}

// GetApprovalCondition returns the ApprovalCondition field value if set, zero value otherwise.
func (o *RequestLinks2) GetApprovalCondition() Link {
	if o == nil || o.ApprovalCondition == nil {
		var ret Link
		return ret
	}
	return *o.ApprovalCondition
}

// GetApprovalConditionOk returns a tuple with the ApprovalCondition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestLinks2) GetApprovalConditionOk() (*Link, bool) {
	if o == nil || o.ApprovalCondition == nil {
		return nil, false
	}
	return o.ApprovalCondition, true
}

// HasApprovalCondition returns a boolean if a field has been set.
func (o *RequestLinks2) HasApprovalCondition() bool {
	if o != nil && o.ApprovalCondition != nil {
		return true
	}

	return false
}

// SetApprovalCondition gets a reference to the given Link and assigns it to the ApprovalCondition field.
func (o *RequestLinks2) SetApprovalCondition(v Link) {
	o.ApprovalCondition = &v
}

// GetApprovalSequence returns the ApprovalSequence field value if set, zero value otherwise.
func (o *RequestLinks2) GetApprovalSequence() Link {
	if o == nil || o.ApprovalSequence == nil {
		var ret Link
		return ret
	}
	return *o.ApprovalSequence
}

// GetApprovalSequenceOk returns a tuple with the ApprovalSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestLinks2) GetApprovalSequenceOk() (*Link, bool) {
	if o == nil || o.ApprovalSequence == nil {
		return nil, false
	}
	return o.ApprovalSequence, true
}

// HasApprovalSequence returns a boolean if a field has been set.
func (o *RequestLinks2) HasApprovalSequence() bool {
	if o != nil && o.ApprovalSequence != nil {
		return true
	}

	return false
}

// SetApprovalSequence gets a reference to the given Link and assigns it to the ApprovalSequence field.
func (o *RequestLinks2) SetApprovalSequence(v Link) {
	o.ApprovalSequence = &v
}

// GetSelf returns the Self field value
func (o *RequestLinks2) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *RequestLinks2) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *RequestLinks2) SetSelf(v Link) {
	o.Self = v
}

func (o RequestLinks2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CatalogEntry != nil {
		toSerialize["catalogEntry"] = o.CatalogEntry
	}
	if o.ApprovalCondition != nil {
		toSerialize["approvalCondition"] = o.ApprovalCondition
	}
	if o.ApprovalSequence != nil {
		toSerialize["approvalSequence"] = o.ApprovalSequence
	}
	if true {
		toSerialize["self"] = o.Self
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestLinks2) UnmarshalJSON(bytes []byte) (err error) {
	varRequestLinks2 := _RequestLinks2{}

	err = json.Unmarshal(bytes, &varRequestLinks2)
	if err == nil {
		*o = RequestLinks2(varRequestLinks2)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "catalogEntry")
		delete(additionalProperties, "approvalCondition")
		delete(additionalProperties, "approvalSequence")
		delete(additionalProperties, "self")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestLinks2 struct {
	value *RequestLinks2
	isSet bool
}

func (v NullableRequestLinks2) Get() *RequestLinks2 {
	return v.value
}

func (v *NullableRequestLinks2) Set(val *RequestLinks2) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestLinks2) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestLinks2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestLinks2(val *RequestLinks2) *NullableRequestLinks2 {
	return &NullableRequestLinks2{value: val, isSet: true}
}

func (v NullableRequestLinks2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestLinks2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
