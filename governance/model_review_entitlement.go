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

// ReviewEntitlement Entitlement and it's entitlement values.
type ReviewEntitlement struct {
	// The entitlement id
	Id string `json:"id"`
	// The entitlement name
	Name string `json:"name"`
	// Collection of entitlement values.
	Values               []ReviewerEntitlementValue `json:"values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewEntitlement ReviewEntitlement

// NewReviewEntitlement instantiates a new ReviewEntitlement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewEntitlement(id string, name string) *ReviewEntitlement {
	this := ReviewEntitlement{}
	this.Id = id
	this.Name = name
	return &this
}

// NewReviewEntitlementWithDefaults instantiates a new ReviewEntitlement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewEntitlementWithDefaults() *ReviewEntitlement {
	this := ReviewEntitlement{}
	return &this
}

// GetId returns the Id field value
func (o *ReviewEntitlement) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReviewEntitlement) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReviewEntitlement) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *ReviewEntitlement) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ReviewEntitlement) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ReviewEntitlement) SetName(v string) {
	o.Name = v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *ReviewEntitlement) GetValues() []ReviewerEntitlementValue {
	if o == nil || o.Values == nil {
		var ret []ReviewerEntitlementValue
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewEntitlement) GetValuesOk() ([]ReviewerEntitlementValue, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *ReviewEntitlement) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []ReviewerEntitlementValue and assigns it to the Values field.
func (o *ReviewEntitlement) SetValues(v []ReviewerEntitlementValue) {
	o.Values = v
}

func (o ReviewEntitlement) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Values != nil {
		toSerialize["values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ReviewEntitlement) UnmarshalJSON(bytes []byte) (err error) {
	varReviewEntitlement := _ReviewEntitlement{}

	err = json.Unmarshal(bytes, &varReviewEntitlement)
	if err == nil {
		*o = ReviewEntitlement(varReviewEntitlement)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "values")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableReviewEntitlement struct {
	value *ReviewEntitlement
	isSet bool
}

func (v NullableReviewEntitlement) Get() *ReviewEntitlement {
	return v.value
}

func (v *NullableReviewEntitlement) Set(val *ReviewEntitlement) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewEntitlement) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewEntitlement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewEntitlement(val *ReviewEntitlement) *NullableReviewEntitlement {
	return &NullableReviewEntitlement{value: val, isSet: true}
}

func (v NullableReviewEntitlement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewEntitlement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
