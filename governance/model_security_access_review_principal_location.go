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

// SecurityAccessReviewPrincipalLocation struct for SecurityAccessReviewPrincipalLocation
type SecurityAccessReviewPrincipalLocation struct {
	// City of the location
	City *string `json:"city,omitempty"`
	// State of the location
	State *string `json:"state,omitempty"`
	// Country of the location
	Country              *string `json:"country,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SecurityAccessReviewPrincipalLocation SecurityAccessReviewPrincipalLocation

// NewSecurityAccessReviewPrincipalLocation instantiates a new SecurityAccessReviewPrincipalLocation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSecurityAccessReviewPrincipalLocation() *SecurityAccessReviewPrincipalLocation {
	this := SecurityAccessReviewPrincipalLocation{}
	return &this
}

// NewSecurityAccessReviewPrincipalLocationWithDefaults instantiates a new SecurityAccessReviewPrincipalLocation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSecurityAccessReviewPrincipalLocationWithDefaults() *SecurityAccessReviewPrincipalLocation {
	this := SecurityAccessReviewPrincipalLocation{}
	return &this
}

// GetCity returns the City field value if set, zero value otherwise.
func (o *SecurityAccessReviewPrincipalLocation) GetCity() string {
	if o == nil || o.City == nil {
		var ret string
		return ret
	}
	return *o.City
}

// GetCityOk returns a tuple with the City field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipalLocation) GetCityOk() (*string, bool) {
	if o == nil || o.City == nil {
		return nil, false
	}
	return o.City, true
}

// HasCity returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipalLocation) HasCity() bool {
	if o != nil && o.City != nil {
		return true
	}

	return false
}

// SetCity gets a reference to the given string and assigns it to the City field.
func (o *SecurityAccessReviewPrincipalLocation) SetCity(v string) {
	o.City = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SecurityAccessReviewPrincipalLocation) GetState() string {
	if o == nil || o.State == nil {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipalLocation) GetStateOk() (*string, bool) {
	if o == nil || o.State == nil {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipalLocation) HasState() bool {
	if o != nil && o.State != nil {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *SecurityAccessReviewPrincipalLocation) SetState(v string) {
	o.State = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *SecurityAccessReviewPrincipalLocation) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SecurityAccessReviewPrincipalLocation) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *SecurityAccessReviewPrincipalLocation) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *SecurityAccessReviewPrincipalLocation) SetCountry(v string) {
	o.Country = &v
}

func (o SecurityAccessReviewPrincipalLocation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.City != nil {
		toSerialize["city"] = o.City
	}
	if o.State != nil {
		toSerialize["state"] = o.State
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SecurityAccessReviewPrincipalLocation) UnmarshalJSON(bytes []byte) (err error) {
	varSecurityAccessReviewPrincipalLocation := _SecurityAccessReviewPrincipalLocation{}

	err = json.Unmarshal(bytes, &varSecurityAccessReviewPrincipalLocation)
	if err == nil {
		*o = SecurityAccessReviewPrincipalLocation(varSecurityAccessReviewPrincipalLocation)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "city")
		delete(additionalProperties, "state")
		delete(additionalProperties, "country")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableSecurityAccessReviewPrincipalLocation struct {
	value *SecurityAccessReviewPrincipalLocation
	isSet bool
}

func (v NullableSecurityAccessReviewPrincipalLocation) Get() *SecurityAccessReviewPrincipalLocation {
	return v.value
}

func (v *NullableSecurityAccessReviewPrincipalLocation) Set(val *SecurityAccessReviewPrincipalLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityAccessReviewPrincipalLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityAccessReviewPrincipalLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityAccessReviewPrincipalLocation(val *SecurityAccessReviewPrincipalLocation) *NullableSecurityAccessReviewPrincipalLocation {
	return &NullableSecurityAccessReviewPrincipalLocation{value: val, isSet: true}
}

func (v NullableSecurityAccessReviewPrincipalLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityAccessReviewPrincipalLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
