/*
White Label Communications CPaas API Documentation

A CPaaS platform API

API version: 1.1
Contact: support@whitelabelcomm.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ModelCallDailySummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelCallDailySummary{}

// ModelCallDailySummary struct for ModelCallDailySummary
type ModelCallDailySummary struct {
	AccountId *string `json:"account_id,omitempty"`
	CallDate *string `json:"call_date,omitempty"`
	CallType *string `json:"call_type,omitempty"`
	CreatedAt *string `json:"created_at,omitempty"`
	TotalCallDuration *string `json:"total_call_duration,omitempty"`
}

// NewModelCallDailySummary instantiates a new ModelCallDailySummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelCallDailySummary() *ModelCallDailySummary {
	this := ModelCallDailySummary{}
	return &this
}

// NewModelCallDailySummaryWithDefaults instantiates a new ModelCallDailySummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelCallDailySummaryWithDefaults() *ModelCallDailySummary {
	this := ModelCallDailySummary{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ModelCallDailySummary) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCallDailySummary) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ModelCallDailySummary) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ModelCallDailySummary) SetAccountId(v string) {
	o.AccountId = &v
}

// GetCallDate returns the CallDate field value if set, zero value otherwise.
func (o *ModelCallDailySummary) GetCallDate() string {
	if o == nil || IsNil(o.CallDate) {
		var ret string
		return ret
	}
	return *o.CallDate
}

// GetCallDateOk returns a tuple with the CallDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCallDailySummary) GetCallDateOk() (*string, bool) {
	if o == nil || IsNil(o.CallDate) {
		return nil, false
	}
	return o.CallDate, true
}

// HasCallDate returns a boolean if a field has been set.
func (o *ModelCallDailySummary) HasCallDate() bool {
	if o != nil && !IsNil(o.CallDate) {
		return true
	}

	return false
}

// SetCallDate gets a reference to the given string and assigns it to the CallDate field.
func (o *ModelCallDailySummary) SetCallDate(v string) {
	o.CallDate = &v
}

// GetCallType returns the CallType field value if set, zero value otherwise.
func (o *ModelCallDailySummary) GetCallType() string {
	if o == nil || IsNil(o.CallType) {
		var ret string
		return ret
	}
	return *o.CallType
}

// GetCallTypeOk returns a tuple with the CallType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCallDailySummary) GetCallTypeOk() (*string, bool) {
	if o == nil || IsNil(o.CallType) {
		return nil, false
	}
	return o.CallType, true
}

// HasCallType returns a boolean if a field has been set.
func (o *ModelCallDailySummary) HasCallType() bool {
	if o != nil && !IsNil(o.CallType) {
		return true
	}

	return false
}

// SetCallType gets a reference to the given string and assigns it to the CallType field.
func (o *ModelCallDailySummary) SetCallType(v string) {
	o.CallType = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ModelCallDailySummary) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCallDailySummary) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ModelCallDailySummary) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *ModelCallDailySummary) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetTotalCallDuration returns the TotalCallDuration field value if set, zero value otherwise.
func (o *ModelCallDailySummary) GetTotalCallDuration() string {
	if o == nil || IsNil(o.TotalCallDuration) {
		var ret string
		return ret
	}
	return *o.TotalCallDuration
}

// GetTotalCallDurationOk returns a tuple with the TotalCallDuration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelCallDailySummary) GetTotalCallDurationOk() (*string, bool) {
	if o == nil || IsNil(o.TotalCallDuration) {
		return nil, false
	}
	return o.TotalCallDuration, true
}

// HasTotalCallDuration returns a boolean if a field has been set.
func (o *ModelCallDailySummary) HasTotalCallDuration() bool {
	if o != nil && !IsNil(o.TotalCallDuration) {
		return true
	}

	return false
}

// SetTotalCallDuration gets a reference to the given string and assigns it to the TotalCallDuration field.
func (o *ModelCallDailySummary) SetTotalCallDuration(v string) {
	o.TotalCallDuration = &v
}

func (o ModelCallDailySummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelCallDailySummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.CallDate) {
		toSerialize["call_date"] = o.CallDate
	}
	if !IsNil(o.CallType) {
		toSerialize["call_type"] = o.CallType
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.TotalCallDuration) {
		toSerialize["total_call_duration"] = o.TotalCallDuration
	}
	return toSerialize, nil
}

type NullableModelCallDailySummary struct {
	value *ModelCallDailySummary
	isSet bool
}

func (v NullableModelCallDailySummary) Get() *ModelCallDailySummary {
	return v.value
}

func (v *NullableModelCallDailySummary) Set(val *ModelCallDailySummary) {
	v.value = val
	v.isSet = true
}

func (v NullableModelCallDailySummary) IsSet() bool {
	return v.isSet
}

func (v *NullableModelCallDailySummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelCallDailySummary(val *ModelCallDailySummary) *NullableModelCallDailySummary {
	return &NullableModelCallDailySummary{value: val, isSet: true}
}

func (v NullableModelCallDailySummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelCallDailySummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


