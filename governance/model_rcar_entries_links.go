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

// RcarEntriesLinks Links available for the list of catalog entries. * `atspoke`: returned only for top level entries and if request types is supported * `match`: returned only if match is supported
type RcarEntriesLinks struct {
	Self                 Link  `json:"self"`
	Next                 *Link `json:"next,omitempty"`
	Atspoke              *Link `json:"atspoke,omitempty"`
	Match                *Link `json:"match,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RcarEntriesLinks RcarEntriesLinks

// NewRcarEntriesLinks instantiates a new RcarEntriesLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRcarEntriesLinks(self Link) *RcarEntriesLinks {
	this := RcarEntriesLinks{}
	this.Self = self
	return &this
}

// NewRcarEntriesLinksWithDefaults instantiates a new RcarEntriesLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRcarEntriesLinksWithDefaults() *RcarEntriesLinks {
	this := RcarEntriesLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *RcarEntriesLinks) GetSelf() Link {
	if o == nil {
		var ret Link
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *RcarEntriesLinks) GetSelfOk() (*Link, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *RcarEntriesLinks) SetSelf(v Link) {
	o.Self = v
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *RcarEntriesLinks) GetNext() Link {
	if o == nil || o.Next == nil {
		var ret Link
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcarEntriesLinks) GetNextOk() (*Link, bool) {
	if o == nil || o.Next == nil {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *RcarEntriesLinks) HasNext() bool {
	if o != nil && o.Next != nil {
		return true
	}

	return false
}

// SetNext gets a reference to the given Link and assigns it to the Next field.
func (o *RcarEntriesLinks) SetNext(v Link) {
	o.Next = &v
}

// GetAtspoke returns the Atspoke field value if set, zero value otherwise.
func (o *RcarEntriesLinks) GetAtspoke() Link {
	if o == nil || o.Atspoke == nil {
		var ret Link
		return ret
	}
	return *o.Atspoke
}

// GetAtspokeOk returns a tuple with the Atspoke field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcarEntriesLinks) GetAtspokeOk() (*Link, bool) {
	if o == nil || o.Atspoke == nil {
		return nil, false
	}
	return o.Atspoke, true
}

// HasAtspoke returns a boolean if a field has been set.
func (o *RcarEntriesLinks) HasAtspoke() bool {
	if o != nil && o.Atspoke != nil {
		return true
	}

	return false
}

// SetAtspoke gets a reference to the given Link and assigns it to the Atspoke field.
func (o *RcarEntriesLinks) SetAtspoke(v Link) {
	o.Atspoke = &v
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *RcarEntriesLinks) GetMatch() Link {
	if o == nil || o.Match == nil {
		var ret Link
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RcarEntriesLinks) GetMatchOk() (*Link, bool) {
	if o == nil || o.Match == nil {
		return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *RcarEntriesLinks) HasMatch() bool {
	if o != nil && o.Match != nil {
		return true
	}

	return false
}

// SetMatch gets a reference to the given Link and assigns it to the Match field.
func (o *RcarEntriesLinks) SetMatch(v Link) {
	o.Match = &v
}

func (o RcarEntriesLinks) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["self"] = o.Self
	}
	if o.Next != nil {
		toSerialize["next"] = o.Next
	}
	if o.Atspoke != nil {
		toSerialize["atspoke"] = o.Atspoke
	}
	if o.Match != nil {
		toSerialize["match"] = o.Match
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RcarEntriesLinks) UnmarshalJSON(bytes []byte) (err error) {
	varRcarEntriesLinks := _RcarEntriesLinks{}

	err = json.Unmarshal(bytes, &varRcarEntriesLinks)
	if err == nil {
		*o = RcarEntriesLinks(varRcarEntriesLinks)
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "self")
		delete(additionalProperties, "next")
		delete(additionalProperties, "atspoke")
		delete(additionalProperties, "match")
		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableRcarEntriesLinks struct {
	value *RcarEntriesLinks
	isSet bool
}

func (v NullableRcarEntriesLinks) Get() *RcarEntriesLinks {
	return v.value
}

func (v *NullableRcarEntriesLinks) Set(val *RcarEntriesLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableRcarEntriesLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableRcarEntriesLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRcarEntriesLinks(val *RcarEntriesLinks) *NullableRcarEntriesLinks {
	return &NullableRcarEntriesLinks{value: val, isSet: true}
}

func (v NullableRcarEntriesLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRcarEntriesLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
