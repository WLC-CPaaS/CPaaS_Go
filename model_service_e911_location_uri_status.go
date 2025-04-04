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

// checks if the ServiceE911LocationURIStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceE911LocationURIStatus{}

// ServiceE911LocationURIStatus struct for ServiceE911LocationURIStatus
type ServiceE911LocationURIStatus struct {
	Code *string `json:"code,omitempty"`
	Description *string `json:"description,omitempty"`
}

// NewServiceE911LocationURIStatus instantiates a new ServiceE911LocationURIStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceE911LocationURIStatus() *ServiceE911LocationURIStatus {
	this := ServiceE911LocationURIStatus{}
	return &this
}

// NewServiceE911LocationURIStatusWithDefaults instantiates a new ServiceE911LocationURIStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceE911LocationURIStatusWithDefaults() *ServiceE911LocationURIStatus {
	this := ServiceE911LocationURIStatus{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ServiceE911LocationURIStatus) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationURIStatus) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ServiceE911LocationURIStatus) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *ServiceE911LocationURIStatus) SetCode(v string) {
	o.Code = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceE911LocationURIStatus) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceE911LocationURIStatus) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceE911LocationURIStatus) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceE911LocationURIStatus) SetDescription(v string) {
	o.Description = &v
}

func (o ServiceE911LocationURIStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceE911LocationURIStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	return toSerialize, nil
}

type NullableServiceE911LocationURIStatus struct {
	value *ServiceE911LocationURIStatus
	isSet bool
}

func (v NullableServiceE911LocationURIStatus) Get() *ServiceE911LocationURIStatus {
	return v.value
}

func (v *NullableServiceE911LocationURIStatus) Set(val *ServiceE911LocationURIStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceE911LocationURIStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceE911LocationURIStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceE911LocationURIStatus(val *ServiceE911LocationURIStatus) *NullableServiceE911LocationURIStatus {
	return &NullableServiceE911LocationURIStatus{value: val, isSet: true}
}

func (v NullableServiceE911LocationURIStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceE911LocationURIStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


