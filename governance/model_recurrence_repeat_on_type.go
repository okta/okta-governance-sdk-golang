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

// RecurrenceRepeatOnType Applicable only, if `interval` is in terms of N months. Specifies the day-of-the-month on which, the campaign should be repeated. If the `interval` is in terms of days, weeks or years, this property can be ignored.  1) `SAME_DAY_AS_START_DATE` - Specifies the day-of-the-month in `startDate`. For eg: if `startDate` happens to be `5th` of some date, then campaign will repeat on 5th day of the matching month.  2) `SAME_WEEKDAY_AS_START_DATE` - Specifies the day-of-the-week in `startDate`. For eg if `startDate` happens to be `Thursday`, then campaign will repeat on `Thursday` of the matching week.  2) `LAST_WEEKDAY_AS_START_DATE` - Specifies the last day-of-the-week in `startDate`. For eg jf `startDate` happens to be `Thursday` and it is the last week of the month, then campaign will repeat on `Thursday` of the matching last week.  Refer few examples below:   1) Let's say the campaign start date is `June 1st 2023`, the `interval: P3M` (Every 3 months) and `repeatOnType: SAME_WEEKDAY_AS_START_DATE`- the `startDate` happens to be in the first week, and it is a `Thursday`. Hence the campaign will repeat every `3 months` and will start on `first week's Thursday`. In this example, some sample campaign occurrences are below -        | Campaign No |Date of recurrence        |       |-------------|--------------------------|       | 1           | Jun  1st, Thursday,  2023|       | 2           | Sep  7th, Thursday,  2023|       | 3           | Dec  7th, Thursday,  2023|       |...          |                          |       |...          |                          |    2) Let's say the campaign start date is `June 1st 2023`, the `interval: P1M` (Every 1 month) and `repeatOnType: SAME_DAY_AS_START_DATE` - the `startDate` happens to be `1st` day of the month. Hence the campaign will repeat every `1 month` and will start on `1st day of the month`. In this example, some sample campaign occurrences are below -        | Campaign No |Date of recurrence        |       |-------------|--------------------------|       | 1           | Jun  1st, Thursday,  2023|       | 2           | Jul  1st, Saturday,  2023|       | 3           | Aug  1st, Tuesday,   2023|       |...          |                          |       |...          |                          |    3) Let's say the campaign start date is `June 1st 2023`, the `interval: P1M` (Every 1 month) and `repeatOnType: SAME_WEEKDAY_AS_START_DATE` - the `startDate` happens to be in the first week, and it is a `Thursday`. Hence the campaign will repeat every `1 month` and will start on `first week's Thursday`. In this example, some sample campaign occurrences are below -        | Campaign No |Date of recurrence        |       |-------------|--------------------------|       | 1           | Jun  1st, Thursday,  2023|       | 2           | Jul  6th, Thursday,  2023|       | 3           | Aug  3rd, Thursday,  2023|       |...          |                          |       |...          |                          |    4) Let's say the campaign start date is `May 23rd 2023`, the `interval: P1M` (Every 1 month) and `repeatOnType: SAME_WEEKDAY_AS_START_DATE` - the `startDate` happens to be in the fourth week, and it is a `Tuesday`. Hence the campaign will repeat every `1 month` and will start on `fourth week's Tuesday`. In this example, some sample campaign occurrences are below -        | Campaign No |Date of recurrence        |       |-------------|--------------------------|       | 1           | May  23rd, Tuesday,  2023|       | 2           | Jun  27th, Tuesday,  2023|       | 3           | Jul  25th, Tuesday,  2023|       | 4           | Aug  22nd, Tuesday,  2023|       |...          |                          |    5) Let's say the campaign start date is `May 29rd 2023`, the `interval: P1M` (Every 1 month) and `repeatOnType: LAST_WEEKDAY_AS_START_DATE` - the `startDate` happens to be in the last week, and it is a `Monday`. Hence the campaign will repeat every `1 month` and will start on `last week's Monday`. In this example, some sample campaign occurrences are below -        | Campaign No |Date of recurrence        |       |-------------|--------------------------|       | 1           | May  29th, Monday,   2023|       | 2           | Jun  26th, Monday,   2023|       | 3           | Jul  31st, Monday,   2023|       | 4           | Aug  28th, Monday,   2023|       |...          |                          |    6) Let's say the campaign start date is `June 28th 2023`, the `interval: P1M` (Every 1 month) and `repeatOnType: SAME_DAY_AS_START_DATE` -  the `startDate` happens to be `28th` day of the month. Hence the campaign will repeat every `1 month` and will start on `28th day of the month`. In this example, some sample campaign occurrences are below -        | Campaign No |Date of recurrence        |       |-------------|--------------------------|       | 1           | Jun 28th, Wednesday, 2023|       | 2           | Jul 28th, Friday,    2023|       | 3           | Aug 28th, Monday,    2023|       |...          |                          |       |...          |                          |    7) Let's say the campaign start date is `June 28th 2023`, the `interval: P1M` (Every 1 month) and `repeatOnType: LAST_WEEKDAY_AS_START_DATE` -  the `startDate` happens to be the last week of the month, and it is a `Wednesday`. Hence the campaign will repeat every `1 month` and will start on `last week's Wednesday`. In this example, some sample campaign occurrences are below -        | Campaign No |Date of recurrence        |       |-------------|--------------------------|       | 1           | Jun 28th, Wednesday, 2023|       | 2           | Jul 26th, Wednesday, 2023|       | 3           | Aug 30th, Wednesday, 2023|       |...          |                          |       |...          |                          |    8) Let's say the campaign start date is `Jan 31st 2024`, the `interval: P1M` (Every 1 month) and `repeatOnType: LAST_WEEKDAY_AS_START_DATE` -  the `startDate` happens to be the last week of the month, and it is a `Wednesday`. Hence the campaign will repeat every `1 month` and will start on `last week's Wednesday`. In this example, some sample campaign occurrences are below -        | Campaign No |Date of recurrence        |       |-------------|--------------------------|       | 1           | Jan 31st, Wednesday, 2024|       | 2           | Feb 28th, Wednesday, 2024|       | 3           | Mar 27th, Wednesday, 2024|       |...          |                          |       |...          |                          |    9) Let's say the campaign start date is `Jan 31st 2024`, the `interval: P1M` (Every 1 month) and `repeatOnType: SAME_DAY_AS_START_DATE` -  the `startDate` happens to be `31st` day of the month. Hence the campaign will repeat every `1 month` and will start on `31st` or `30th` or `28/29th`(for February). In this example, some sample campaign occurrences are below -        | Campaign No |Date of recurrence        |       |-------------|--------------------------|       | 1           | Jan 31st, Wednesday, 2024|       | 2           | Feb 29th, Thursday,  2024|       | 3           | Mar 31st, Sunday,    2024|       |...          |                          |       |...          |                          |
type RecurrenceRepeatOnType string

// List of recurrence-repeat-on-type
const (
	RECURRENCEREPEATONTYPE_SAME_DAY_AS_START_DATE     RecurrenceRepeatOnType = "SAME_DAY_AS_START_DATE"
	RECURRENCEREPEATONTYPE_SAME_WEEKDAY_AS_START_DATE RecurrenceRepeatOnType = "SAME_WEEKDAY_AS_START_DATE"
	RECURRENCEREPEATONTYPE_LAST_WEEKDAY_AS_START_DATE RecurrenceRepeatOnType = "LAST_WEEKDAY_AS_START_DATE"
)

// All allowed values of RecurrenceRepeatOnType enum
var AllowedRecurrenceRepeatOnTypeEnumValues = []RecurrenceRepeatOnType{
	"SAME_DAY_AS_START_DATE",
	"SAME_WEEKDAY_AS_START_DATE",
	"LAST_WEEKDAY_AS_START_DATE",
}

func (v *RecurrenceRepeatOnType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RecurrenceRepeatOnType(value)
	for _, existing := range AllowedRecurrenceRepeatOnTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RecurrenceRepeatOnType", value)
}

// NewRecurrenceRepeatOnTypeFromValue returns a pointer to a valid RecurrenceRepeatOnType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRecurrenceRepeatOnTypeFromValue(v string) (*RecurrenceRepeatOnType, error) {
	ev := RecurrenceRepeatOnType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RecurrenceRepeatOnType: valid values are %v", v, AllowedRecurrenceRepeatOnTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RecurrenceRepeatOnType) IsValid() bool {
	for _, existing := range AllowedRecurrenceRepeatOnTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to recurrence-repeat-on-type value
func (v RecurrenceRepeatOnType) Ptr() *RecurrenceRepeatOnType {
	return &v
}

type NullableRecurrenceRepeatOnType struct {
	value *RecurrenceRepeatOnType
	isSet bool
}

func (v NullableRecurrenceRepeatOnType) Get() *RecurrenceRepeatOnType {
	return v.value
}

func (v *NullableRecurrenceRepeatOnType) Set(val *RecurrenceRepeatOnType) {
	v.value = val
	v.isSet = true
}

func (v NullableRecurrenceRepeatOnType) IsSet() bool {
	return v.isSet
}

func (v *NullableRecurrenceRepeatOnType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecurrenceRepeatOnType(val *RecurrenceRepeatOnType) *NullableRecurrenceRepeatOnType {
	return &NullableRecurrenceRepeatOnType{value: val, isSet: true}
}

func (v NullableRecurrenceRepeatOnType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecurrenceRepeatOnType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
