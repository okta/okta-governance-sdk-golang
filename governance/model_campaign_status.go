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

// CampaignStatus the model 'CampaignStatus'
type CampaignStatus string

// List of campaign-status
const (
	CAMPAIGNSTATUS_SCHEDULED CampaignStatus = "SCHEDULED"
	CAMPAIGNSTATUS_LAUNCHING CampaignStatus = "LAUNCHING"
	CAMPAIGNSTATUS_ACTIVE    CampaignStatus = "ACTIVE"
	CAMPAIGNSTATUS_COMPLETED CampaignStatus = "COMPLETED"
	CAMPAIGNSTATUS_DELETED   CampaignStatus = "DELETED"
	CAMPAIGNSTATUS_ERROR     CampaignStatus = "ERROR"
)

// All allowed values of CampaignStatus enum
var AllowedCampaignStatusEnumValues = []CampaignStatus{
	"SCHEDULED",
	"LAUNCHING",
	"ACTIVE",
	"COMPLETED",
	"DELETED",
	"ERROR",
}

func (v *CampaignStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CampaignStatus(value)
	for _, existing := range AllowedCampaignStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CampaignStatus", value)
}

// NewCampaignStatusFromValue returns a pointer to a valid CampaignStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCampaignStatusFromValue(v string) (*CampaignStatus, error) {
	ev := CampaignStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CampaignStatus: valid values are %v", v, AllowedCampaignStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CampaignStatus) IsValid() bool {
	for _, existing := range AllowedCampaignStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to campaign-status value
func (v CampaignStatus) Ptr() *CampaignStatus {
	return &v
}

type NullableCampaignStatus struct {
	value *CampaignStatus
	isSet bool
}

func (v NullableCampaignStatus) Get() *CampaignStatus {
	return v.value
}

func (v *NullableCampaignStatus) Set(val *CampaignStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignStatus(val *CampaignStatus) *NullableCampaignStatus {
	return &NullableCampaignStatus{value: val, isSet: true}
}

func (v NullableCampaignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
