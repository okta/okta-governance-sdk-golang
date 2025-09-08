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

// RequestSequencesListLinks struct for RequestSequencesListLinks
type RequestSequencesListLinks struct {
	Self                 Link  `json:"self"`
	NewSequence          *Link `json:"newSequence,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestSequencesListLinks RequestSequencesListLinks

// NewRequestSequencesListLinks instantiates a new RequestSequencesListLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestSequencesListLinks(self Link) *RequestSequencesListLinks {
	this := RequestSequencesListLinks{}
	this.Self = self
	return &this
}

// NewRequestSequencesListLinksWithDefaults instantiates a new RequestSequencesListLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestSequencesListLinksWithDefaults() *RequestSequencesListLinks {
	this := RequestSequencesListLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *RequestSequencesListLinks) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *RequestSequencesListLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *RequestSequencesListLinks) SetSelf(v Link) {
	o.Self = v
}

// GetNewSequence returns the NewSequence field value if set, zero value otherwise.
func (o *RequestSequencesListLinks) GetNewSequence() Link {
	if o == nil || o.NewSequence == nil {
		var ret Link
		return ret
	}
	return *o.NewSequence
}

// GetNewSequenceOk returns a tuple with the NewSequence field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestSequencesListLinks) GetNewSequenceOk() (*Link, bool) {
	if o == nil || o.NewSequence == nil {
		return nil, false
	}
	return o.NewSequence, true
}

// HasNewSequence returns a boolean if a field has been set.
func (o *RequestSequencesListLinks) HasNewSequence() bool {
	if o != nil && o.NewSequence != nil {
		return true
	}

	return false
}

// SetNewSequence gets a reference to the given Link and assigns it to the NewSequence field.
func (o *RequestSequencesListLinks) SetNewSequence(v Link) {
	o.NewSequence = &v
}

func (o RequestSequencesListLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["self"] = o.Self
	}
	if o.NewSequence != nil {
		toSerialize["newSequence"] = o.NewSequence
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestSequencesListLinks) UnmarshalJSON(bytes []byte) (err error) {
	varRequestSequencesListLinks := _RequestSequencesListLinks{}

	err = json.Unmarshal(bytes, &varRequestSequencesListLinks)
	if err == nil {
		*o = RequestSequencesListLinks(varRequestSequencesListLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "newSequence")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRequestSequencesListLinks struct {
	value *RequestSequencesListLinks
	isSet bool
}

func (v NullableRequestSequencesListLinks) Get() *RequestSequencesListLinks {
	return v.value
}

func (v *NullableRequestSequencesListLinks) Set(val *RequestSequencesListLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestSequencesListLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestSequencesListLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestSequencesListLinks(val *RequestSequencesListLinks) *NullableRequestSequencesListLinks {
	return &NullableRequestSequencesListLinks{value: val, isSet: true}
}

func (v NullableRequestSequencesListLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestSequencesListLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
