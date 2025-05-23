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

// checks if the ServiceUserOutputFullCallerid type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceUserOutputFullCallerid{}

// ServiceUserOutputFullCallerid struct for ServiceUserOutputFullCallerid
type ServiceUserOutputFullCallerid struct {
	Emergency *ServiceUserOutputFullCalleridEmergency `json:"emergency,omitempty"`
	External *ServiceUserOutputFullCalleridExternal `json:"external,omitempty"`
	Internal *ServiceUserOutputFullCalleridInternal `json:"internal,omitempty"`
}

// NewServiceUserOutputFullCallerid instantiates a new ServiceUserOutputFullCallerid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceUserOutputFullCallerid() *ServiceUserOutputFullCallerid {
	this := ServiceUserOutputFullCallerid{}
	return &this
}

// NewServiceUserOutputFullCalleridWithDefaults instantiates a new ServiceUserOutputFullCallerid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceUserOutputFullCalleridWithDefaults() *ServiceUserOutputFullCallerid {
	this := ServiceUserOutputFullCallerid{}
	return &this
}

// GetEmergency returns the Emergency field value if set, zero value otherwise.
func (o *ServiceUserOutputFullCallerid) GetEmergency() ServiceUserOutputFullCalleridEmergency {
	if o == nil || IsNil(o.Emergency) {
		var ret ServiceUserOutputFullCalleridEmergency
		return ret
	}
	return *o.Emergency
}

// GetEmergencyOk returns a tuple with the Emergency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUserOutputFullCallerid) GetEmergencyOk() (*ServiceUserOutputFullCalleridEmergency, bool) {
	if o == nil || IsNil(o.Emergency) {
		return nil, false
	}
	return o.Emergency, true
}

// HasEmergency returns a boolean if a field has been set.
func (o *ServiceUserOutputFullCallerid) HasEmergency() bool {
	if o != nil && !IsNil(o.Emergency) {
		return true
	}

	return false
}

// SetEmergency gets a reference to the given ServiceUserOutputFullCalleridEmergency and assigns it to the Emergency field.
func (o *ServiceUserOutputFullCallerid) SetEmergency(v ServiceUserOutputFullCalleridEmergency) {
	o.Emergency = &v
}

// GetExternal returns the External field value if set, zero value otherwise.
func (o *ServiceUserOutputFullCallerid) GetExternal() ServiceUserOutputFullCalleridExternal {
	if o == nil || IsNil(o.External) {
		var ret ServiceUserOutputFullCalleridExternal
		return ret
	}
	return *o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUserOutputFullCallerid) GetExternalOk() (*ServiceUserOutputFullCalleridExternal, bool) {
	if o == nil || IsNil(o.External) {
		return nil, false
	}
	return o.External, true
}

// HasExternal returns a boolean if a field has been set.
func (o *ServiceUserOutputFullCallerid) HasExternal() bool {
	if o != nil && !IsNil(o.External) {
		return true
	}

	return false
}

// SetExternal gets a reference to the given ServiceUserOutputFullCalleridExternal and assigns it to the External field.
func (o *ServiceUserOutputFullCallerid) SetExternal(v ServiceUserOutputFullCalleridExternal) {
	o.External = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *ServiceUserOutputFullCallerid) GetInternal() ServiceUserOutputFullCalleridInternal {
	if o == nil || IsNil(o.Internal) {
		var ret ServiceUserOutputFullCalleridInternal
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceUserOutputFullCallerid) GetInternalOk() (*ServiceUserOutputFullCalleridInternal, bool) {
	if o == nil || IsNil(o.Internal) {
		return nil, false
	}
	return o.Internal, true
}

// HasInternal returns a boolean if a field has been set.
func (o *ServiceUserOutputFullCallerid) HasInternal() bool {
	if o != nil && !IsNil(o.Internal) {
		return true
	}

	return false
}

// SetInternal gets a reference to the given ServiceUserOutputFullCalleridInternal and assigns it to the Internal field.
func (o *ServiceUserOutputFullCallerid) SetInternal(v ServiceUserOutputFullCalleridInternal) {
	o.Internal = &v
}

func (o ServiceUserOutputFullCallerid) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceUserOutputFullCallerid) ToMap() (map[string]interface{}, error) {
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

type NullableServiceUserOutputFullCallerid struct {
	value *ServiceUserOutputFullCallerid
	isSet bool
}

func (v NullableServiceUserOutputFullCallerid) Get() *ServiceUserOutputFullCallerid {
	return v.value
}

func (v *NullableServiceUserOutputFullCallerid) Set(val *ServiceUserOutputFullCallerid) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceUserOutputFullCallerid) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceUserOutputFullCallerid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceUserOutputFullCallerid(val *ServiceUserOutputFullCallerid) *NullableServiceUserOutputFullCallerid {
	return &NullableServiceUserOutputFullCallerid{value: val, isSet: true}
}

func (v NullableServiceUserOutputFullCallerid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceUserOutputFullCallerid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


