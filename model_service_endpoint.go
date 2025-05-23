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

// checks if the ServiceEndpoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceEndpoint{}

// ServiceEndpoint struct for ServiceEndpoint
type ServiceEndpoint struct {
	Type *string `json:"type,omitempty"`
}

// NewServiceEndpoint instantiates a new ServiceEndpoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceEndpoint() *ServiceEndpoint {
	this := ServiceEndpoint{}
	return &this
}

// NewServiceEndpointWithDefaults instantiates a new ServiceEndpoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceEndpointWithDefaults() *ServiceEndpoint {
	this := ServiceEndpoint{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *ServiceEndpoint) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceEndpoint) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ServiceEndpoint) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *ServiceEndpoint) SetType(v string) {
	o.Type = &v
}

func (o ServiceEndpoint) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceEndpoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableServiceEndpoint struct {
	value *ServiceEndpoint
	isSet bool
}

func (v NullableServiceEndpoint) Get() *ServiceEndpoint {
	return v.value
}

func (v *NullableServiceEndpoint) Set(val *ServiceEndpoint) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceEndpoint) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceEndpoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceEndpoint(val *ServiceEndpoint) *NullableServiceEndpoint {
	return &NullableServiceEndpoint{value: val, isSet: true}
}

func (v NullableServiceEndpoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceEndpoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


