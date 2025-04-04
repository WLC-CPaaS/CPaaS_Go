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

// checks if the ServiceCampaignPhoneNumberOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceCampaignPhoneNumberOutput{}

// ServiceCampaignPhoneNumberOutput struct for ServiceCampaignPhoneNumberOutput
type ServiceCampaignPhoneNumberOutput struct {
	TelephoneNumbers []string `json:"telephone_numbers,omitempty"`
	TotalCount *int32 `json:"total_count,omitempty"`
}

// NewServiceCampaignPhoneNumberOutput instantiates a new ServiceCampaignPhoneNumberOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceCampaignPhoneNumberOutput() *ServiceCampaignPhoneNumberOutput {
	this := ServiceCampaignPhoneNumberOutput{}
	return &this
}

// NewServiceCampaignPhoneNumberOutputWithDefaults instantiates a new ServiceCampaignPhoneNumberOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceCampaignPhoneNumberOutputWithDefaults() *ServiceCampaignPhoneNumberOutput {
	this := ServiceCampaignPhoneNumberOutput{}
	return &this
}

// GetTelephoneNumbers returns the TelephoneNumbers field value if set, zero value otherwise.
func (o *ServiceCampaignPhoneNumberOutput) GetTelephoneNumbers() []string {
	if o == nil || IsNil(o.TelephoneNumbers) {
		var ret []string
		return ret
	}
	return o.TelephoneNumbers
}

// GetTelephoneNumbersOk returns a tuple with the TelephoneNumbers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCampaignPhoneNumberOutput) GetTelephoneNumbersOk() ([]string, bool) {
	if o == nil || IsNil(o.TelephoneNumbers) {
		return nil, false
	}
	return o.TelephoneNumbers, true
}

// HasTelephoneNumbers returns a boolean if a field has been set.
func (o *ServiceCampaignPhoneNumberOutput) HasTelephoneNumbers() bool {
	if o != nil && !IsNil(o.TelephoneNumbers) {
		return true
	}

	return false
}

// SetTelephoneNumbers gets a reference to the given []string and assigns it to the TelephoneNumbers field.
func (o *ServiceCampaignPhoneNumberOutput) SetTelephoneNumbers(v []string) {
	o.TelephoneNumbers = v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *ServiceCampaignPhoneNumberOutput) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceCampaignPhoneNumberOutput) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *ServiceCampaignPhoneNumberOutput) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *ServiceCampaignPhoneNumberOutput) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o ServiceCampaignPhoneNumberOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceCampaignPhoneNumberOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TelephoneNumbers) {
		toSerialize["telephone_numbers"] = o.TelephoneNumbers
	}
	if !IsNil(o.TotalCount) {
		toSerialize["total_count"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableServiceCampaignPhoneNumberOutput struct {
	value *ServiceCampaignPhoneNumberOutput
	isSet bool
}

func (v NullableServiceCampaignPhoneNumberOutput) Get() *ServiceCampaignPhoneNumberOutput {
	return v.value
}

func (v *NullableServiceCampaignPhoneNumberOutput) Set(val *ServiceCampaignPhoneNumberOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceCampaignPhoneNumberOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceCampaignPhoneNumberOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceCampaignPhoneNumberOutput(val *ServiceCampaignPhoneNumberOutput) *NullableServiceCampaignPhoneNumberOutput {
	return &NullableServiceCampaignPhoneNumberOutput{value: val, isSet: true}
}

func (v NullableServiceCampaignPhoneNumberOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceCampaignPhoneNumberOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


