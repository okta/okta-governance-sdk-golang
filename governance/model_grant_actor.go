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
	"fmt"
)

// GrantActor The actor sending the grant request
type GrantActor string

// List of grant-actor
const (
	GRANTACTOR_API            GrantActor = "API"
	GRANTACTOR_ACCESS_REQUEST GrantActor = "ACCESS_REQUEST"
	GRANTACTOR_NONE           GrantActor = "NONE"
	GRANTACTOR_ADMIN          GrantActor = "ADMIN"
)

// All allowed values of GrantActor enum
var AllowedGrantActorEnumValues = []GrantActor{
	"API",
	"ACCESS_REQUEST",
	"NONE",
	"ADMIN",
}

func (v *GrantActor) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GrantActor(value)
	for _, existing := range AllowedGrantActorEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GrantActor", value)
}

// NewGrantActorFromValue returns a pointer to a valid GrantActor
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGrantActorFromValue(v string) (*GrantActor, error) {
	ev := GrantActor(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GrantActor: valid values are %v", v, AllowedGrantActorEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GrantActor) IsValid() bool {
	for _, existing := range AllowedGrantActorEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to grant-actor value
func (v GrantActor) Ptr() *GrantActor {
	return &v
}

type NullableGrantActor struct {
	value *GrantActor
	isSet bool
}

func (v NullableGrantActor) Get() *GrantActor {
	return v.value
}

func (v *NullableGrantActor) Set(val *GrantActor) {
	v.value = val
	v.isSet = true
}

func (v NullableGrantActor) IsSet() bool {
	return v.isSet
}

func (v *NullableGrantActor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrantActor(val *GrantActor) *NullableGrantActor {
	return &NullableGrantActor{value: val, isSet: true}
}

func (v NullableGrantActor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrantActor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
