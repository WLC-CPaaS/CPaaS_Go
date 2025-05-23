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

// checks if the ServiceVOIPDeviceAddEdit3c type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceVOIPDeviceAddEdit3c{}

// ServiceVOIPDeviceAddEdit3c struct for ServiceVOIPDeviceAddEdit3c
type ServiceVOIPDeviceAddEdit3c struct {
	Emergency *ServiceVOIPDeviceAddEdit4 `json:"emergency,omitempty"`
	External *ServiceVOIPDeviceAddEdit4 `json:"external,omitempty"`
	Internal *ServiceVOIPDeviceAddEdit4 `json:"internal,omitempty"`
}

// NewServiceVOIPDeviceAddEdit3c instantiates a new ServiceVOIPDeviceAddEdit3c object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceVOIPDeviceAddEdit3c() *ServiceVOIPDeviceAddEdit3c {
	this := ServiceVOIPDeviceAddEdit3c{}
	return &this
}

// NewServiceVOIPDeviceAddEdit3cWithDefaults instantiates a new ServiceVOIPDeviceAddEdit3c object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceVOIPDeviceAddEdit3cWithDefaults() *ServiceVOIPDeviceAddEdit3c {
	this := ServiceVOIPDeviceAddEdit3c{}
	return &this
}

// GetEmergency returns the Emergency field value if set, zero value otherwise.
func (o *ServiceVOIPDeviceAddEdit3c) GetEmergency() ServiceVOIPDeviceAddEdit4 {
	if o == nil || IsNil(o.Emergency) {
		var ret ServiceVOIPDeviceAddEdit4
		return ret
	}
	return *o.Emergency
}

// GetEmergencyOk returns a tuple with the Emergency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPDeviceAddEdit3c) GetEmergencyOk() (*ServiceVOIPDeviceAddEdit4, bool) {
	if o == nil || IsNil(o.Emergency) {
		return nil, false
	}
	return o.Emergency, true
}

// HasEmergency returns a boolean if a field has been set.
func (o *ServiceVOIPDeviceAddEdit3c) HasEmergency() bool {
	if o != nil && !IsNil(o.Emergency) {
		return true
	}

	return false
}

// SetEmergency gets a reference to the given ServiceVOIPDeviceAddEdit4 and assigns it to the Emergency field.
func (o *ServiceVOIPDeviceAddEdit3c) SetEmergency(v ServiceVOIPDeviceAddEdit4) {
	o.Emergency = &v
}

// GetExternal returns the External field value if set, zero value otherwise.
func (o *ServiceVOIPDeviceAddEdit3c) GetExternal() ServiceVOIPDeviceAddEdit4 {
	if o == nil || IsNil(o.External) {
		var ret ServiceVOIPDeviceAddEdit4
		return ret
	}
	return *o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPDeviceAddEdit3c) GetExternalOk() (*ServiceVOIPDeviceAddEdit4, bool) {
	if o == nil || IsNil(o.External) {
		return nil, false
	}
	return o.External, true
}

// HasExternal returns a boolean if a field has been set.
func (o *ServiceVOIPDeviceAddEdit3c) HasExternal() bool {
	if o != nil && !IsNil(o.External) {
		return true
	}

	return false
}

// SetExternal gets a reference to the given ServiceVOIPDeviceAddEdit4 and assigns it to the External field.
func (o *ServiceVOIPDeviceAddEdit3c) SetExternal(v ServiceVOIPDeviceAddEdit4) {
	o.External = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *ServiceVOIPDeviceAddEdit3c) GetInternal() ServiceVOIPDeviceAddEdit4 {
	if o == nil || IsNil(o.Internal) {
		var ret ServiceVOIPDeviceAddEdit4
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPDeviceAddEdit3c) GetInternalOk() (*ServiceVOIPDeviceAddEdit4, bool) {
	if o == nil || IsNil(o.Internal) {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *ServiceVOIPDeviceAddEdit3c) HasInternal() bool {
	if o != nil && !IsNil(o.Internal) {
		return true
	}

	return false
}

// SetInternal gets a reference to the given ServiceVOIPDeviceAddEdit4 and assigns it to the Internal field.
func (o *ServiceVOIPDeviceAddEdit3c) SetInternal(v ServiceVOIPDeviceAddEdit4) {
	o.Internal = &v
}

func (o ServiceVOIPDeviceAddEdit3c) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceVOIPDeviceAddEdit3c) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Emergency) {
		toSerialize["emergency"] = o.Emergency
	}
	if !IsNil(o.External) {
		toSerialize["external"] = o.External
	}
	if !IsNil(o.Internal) {
		toSerialize["internal"] = o.Internal
	}
	return toSerialize, nil
}

type NullableServiceVOIPDeviceAddEdit3c struct {
	value *ServiceVOIPDeviceAddEdit3c
	isSet bool
}

func (v NullableServiceVOIPDeviceAddEdit3c) Get() *ServiceVOIPDeviceAddEdit3c {
	return v.value
}

func (v *NullableServiceVOIPDeviceAddEdit3c) Set(val *ServiceVOIPDeviceAddEdit3c) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceVOIPDeviceAddEdit3c) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceVOIPDeviceAddEdit3c) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceVOIPDeviceAddEdit3c(val *ServiceVOIPDeviceAddEdit3c) *NullableServiceVOIPDeviceAddEdit3c {
	return &NullableServiceVOIPDeviceAddEdit3c{value: val, isSet: true}
}

func (v NullableServiceVOIPDeviceAddEdit3c) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceVOIPDeviceAddEdit3c) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


