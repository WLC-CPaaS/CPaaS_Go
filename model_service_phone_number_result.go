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

// checks if the ServicePhoneNumberResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServicePhoneNumberResult{}

// ServicePhoneNumberResult struct for ServicePhoneNumberResult
type ServicePhoneNumberResult struct {
	Message *string `json:"message,omitempty"`
	Phonenumber *string `json:"phonenumber,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewServicePhoneNumberResult instantiates a new ServicePhoneNumberResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServicePhoneNumberResult() *ServicePhoneNumberResult {
	this := ServicePhoneNumberResult{}
	return &this
}

// NewServicePhoneNumberResultWithDefaults instantiates a new ServicePhoneNumberResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicePhoneNumberResultWithDefaults() *ServicePhoneNumberResult {
	this := ServicePhoneNumberResult{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ServicePhoneNumberResult) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePhoneNumberResult) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ServicePhoneNumberResult) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ServicePhoneNumberResult) SetMessage(v string) {
	o.Message = &v
}

// GetPhonenumber returns the Phonenumber field value if set, zero value otherwise.
func (o *ServicePhoneNumberResult) GetPhonenumber() string {
	if o == nil || IsNil(o.Phonenumber) {
		var ret string
		return ret
	}
	return *o.Phonenumber
}

// GetPhonenumberOk returns a tuple with the Phonenumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePhoneNumberResult) GetPhonenumberOk() (*string, bool) {
	if o == nil || IsNil(o.Phonenumber) {
		return nil, false
	}
	return o.Phonenumber, true
}

// HasPhonenumber returns a boolean if a field has been set.
func (o *ServicePhoneNumberResult) HasPhonenumber() bool {
	if o != nil && !IsNil(o.Phonenumber) {
		return true
	}

	return false
}

// SetPhonenumber gets a reference to the given string and assigns it to the Phonenumber field.
func (o *ServicePhoneNumberResult) SetPhonenumber(v string) {
	o.Phonenumber = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ServicePhoneNumberResult) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServicePhoneNumberResult) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ServicePhoneNumberResult) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ServicePhoneNumberResult) SetStatus(v string) {
	o.Status = &v
}

func (o ServicePhoneNumberResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServicePhoneNumberResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Phonenumber) {
		toSerialize["phonenumber"] = o.Phonenumber
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableServicePhoneNumberResult struct {
	value *ServicePhoneNumberResult
	isSet bool
}

func (v NullableServicePhoneNumberResult) Get() *ServicePhoneNumberResult {
	return v.value
}

func (v *NullableServicePhoneNumberResult) Set(val *ServicePhoneNumberResult) {
	v.value = val
	v.isSet = true
}

func (v NullableServicePhoneNumberResult) IsSet() bool {
	return v.isSet
}

func (v *NullableServicePhoneNumberResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServicePhoneNumberResult(val *ServicePhoneNumberResult) *NullableServicePhoneNumberResult {
	return &NullableServicePhoneNumberResult{value: val, isSet: true}
}

func (v NullableServicePhoneNumberResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServicePhoneNumberResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


