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
)

// CampaignType Identifies if it is a resource campaign or a user campaign. By default it is RESOURCE.
type CampaignType string

// List of campaign-type
const (
	CAMPAIGNTYPE_RESOURCE CampaignType = "RESOURCE"
	CAMPAIGNTYPE_USER     CampaignType = "USER"
)

// All allowed values of CampaignType enum
var AllowedCampaignTypeEnumValues = []CampaignType{
	"RESOURCE",
	"USER",
}

func (v *CampaignType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CampaignType(value)
	for _, existing := range AllowedCampaignTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CampaignType", value)
}

// NewCampaignTypeFromValue returns a pointer to a valid CampaignType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignTypeFromValue(v string) (*CampaignType, error) {
	ev := CampaignType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignType: valid values are %v", v, AllowedCampaignTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignType) IsValid() bool {
	for _, existing := range AllowedCampaignTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign-type value
func (v CampaignType) Ptr() *CampaignType {
	return &v
}

type NullableCampaignType struct {
	value *CampaignType
	isSet bool
}

func (v NullableCampaignType) Get() *CampaignType {
	return v.value
}

func (v *NullableCampaignType) Set(val *CampaignType) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignType) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignType(val *CampaignType) *NullableCampaignType {
	return &NullableCampaignType{value: val, isSet: true}
}

func (v NullableCampaignType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
