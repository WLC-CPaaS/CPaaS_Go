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

// checks if the ServiceStoragePlanDatabaseAttachment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceStoragePlanDatabaseAttachment{}

// ServiceStoragePlanDatabaseAttachment struct for ServiceStoragePlanDatabaseAttachment
type ServiceStoragePlanDatabaseAttachment struct {
	Handler *string `json:"handler,omitempty"`
	Params map[string]interface{} `json:"params,omitempty"`
	Stub *bool `json:"stub,omitempty"`
}

// NewServiceStoragePlanDatabaseAttachment instantiates a new ServiceStoragePlanDatabaseAttachment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceStoragePlanDatabaseAttachment() *ServiceStoragePlanDatabaseAttachment {
	this := ServiceStoragePlanDatabaseAttachment{}
	return &this
}

// NewServiceStoragePlanDatabaseAttachmentWithDefaults instantiates a new ServiceStoragePlanDatabaseAttachment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceStoragePlanDatabaseAttachmentWithDefaults() *ServiceStoragePlanDatabaseAttachment {
	this := ServiceStoragePlanDatabaseAttachment{}
	return &this
}

// GetHandler returns the Handler field value if set, zero value otherwise.
func (o *ServiceStoragePlanDatabaseAttachment) GetHandler() string {
	if o == nil || IsNil(o.Handler) {
		var ret string
		return ret
	}
	return *o.Handler
}

// GetHandlerOk returns a tuple with the Handler field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceStoragePlanDatabaseAttachment) GetHandlerOk() (*string, bool) {
	if o == nil || IsNil(o.Handler) {
		return nil, false
	}
	return o.Handler, true
}

// HasHandler returns a boolean if a field has been set.
func (o *ServiceStoragePlanDatabaseAttachment) HasHandler() bool {
	if o != nil && !IsNil(o.Handler) {
		return true
	}

	return false
}

// SetHandler gets a reference to the given string and assigns it to the Handler field.
func (o *ServiceStoragePlanDatabaseAttachment) SetHandler(v string) {
	o.Handler = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *ServiceStoragePlanDatabaseAttachment) GetParams() map[string]interface{} {
	if o == nil || IsNil(o.Params) {
		var ret map[string]interface{}
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceStoragePlanDatabaseAttachment) GetParamsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Params) {
		return map[string]interface{}{}, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *ServiceStoragePlanDatabaseAttachment) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given map[string]interface{} and assigns it to the Params field.
func (o *ServiceStoragePlanDatabaseAttachment) SetParams(v map[string]interface{}) {
	o.Params = v
}

// GetStub returns the Stub field value if set, zero value otherwise.
func (o *ServiceStoragePlanDatabaseAttachment) GetStub() bool {
	if o == nil || IsNil(o.Stub) {
		var ret bool
		return ret
	}
	return *o.Stub
}

// GetStubOk returns a tuple with the Stub field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceStoragePlanDatabaseAttachment) GetStubOk() (*bool, bool) {
	if o == nil || IsNil(o.Stub) {
		return nil, false
	}
	return o.Stub, true
}

// HasStub returns a boolean if a field has been set.
func (o *ServiceStoragePlanDatabaseAttachment) HasStub() bool {
	if o != nil && !IsNil(o.Stub) {
		return true
	}

	return false
}

// SetStub gets a reference to the given bool and assigns it to the Stub field.
func (o *ServiceStoragePlanDatabaseAttachment) SetStub(v bool) {
	o.Stub = &v
}

func (o ServiceStoragePlanDatabaseAttachment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceStoragePlanDatabaseAttachment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Handler) {
		toSerialize["handler"] = o.Handler
	}
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	if !IsNil(o.Stub) {
		toSerialize["stub"] = o.Stub
	}
	return toSerialize, nil
}

type NullableServiceStoragePlanDatabaseAttachment struct {
	value *ServiceStoragePlanDatabaseAttachment
	isSet bool
}

func (v NullableServiceStoragePlanDatabaseAttachment) Get() *ServiceStoragePlanDatabaseAttachment {
	return v.value
}

func (v *NullableServiceStoragePlanDatabaseAttachment) Set(val *ServiceStoragePlanDatabaseAttachment) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceStoragePlanDatabaseAttachment) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceStoragePlanDatabaseAttachment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceStoragePlanDatabaseAttachment(val *ServiceStoragePlanDatabaseAttachment) *NullableServiceStoragePlanDatabaseAttachment {
	return &NullableServiceStoragePlanDatabaseAttachment{value: val, isSet: true}
}

func (v NullableServiceStoragePlanDatabaseAttachment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceStoragePlanDatabaseAttachment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


