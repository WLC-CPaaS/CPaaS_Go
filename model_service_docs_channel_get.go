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

// checks if the ServiceDocsChannelGet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceDocsChannelGet{}

// ServiceDocsChannelGet struct for ServiceDocsChannelGet
type ServiceDocsChannelGet struct {
	Data []ServiceChannelOutput `json:"data,omitempty"`
	// Unique id for each request
	RequestId *string `json:"request_id,omitempty"`
	// HTTP response status code
	StatusCode *int32 `json:"status_code,omitempty"`
}

// NewServiceDocsChannelGet instantiates a new ServiceDocsChannelGet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDocsChannelGet() *ServiceDocsChannelGet {
	this := ServiceDocsChannelGet{}
	return &this
}

// NewServiceDocsChannelGetWithDefaults instantiates a new ServiceDocsChannelGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDocsChannelGetWithDefaults() *ServiceDocsChannelGet {
	this := ServiceDocsChannelGet{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ServiceDocsChannelGet) GetData() []ServiceChannelOutput {
	if o == nil || IsNil(o.Data) {
		var ret []ServiceChannelOutput
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsChannelGet) GetDataOk() ([]ServiceChannelOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ServiceDocsChannelGet) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ServiceChannelOutput and assigns it to the Data field.
func (o *ServiceDocsChannelGet) SetData(v []ServiceChannelOutput) {
	o.Data = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ServiceDocsChannelGet) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsChannelGet) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ServiceDocsChannelGet) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ServiceDocsChannelGet) SetRequestId(v string) {
	o.RequestId = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ServiceDocsChannelGet) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsChannelGet) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ServiceDocsChannelGet) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *ServiceDocsChannelGet) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o ServiceDocsChannelGet) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceDocsChannelGet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	if !IsNil(o.StatusCode) {
		toSerialize["status_code"] = o.StatusCode
	}
	return toSerialize, nil
}

type NullableServiceDocsChannelGet struct {
	value *ServiceDocsChannelGet
	isSet bool
}

func (v NullableServiceDocsChannelGet) Get() *ServiceDocsChannelGet {
	return v.value
}

func (v *NullableServiceDocsChannelGet) Set(val *ServiceDocsChannelGet) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDocsChannelGet) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDocsChannelGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDocsChannelGet(val *ServiceDocsChannelGet) *NullableServiceDocsChannelGet {
	return &NullableServiceDocsChannelGet{value: val, isSet: true}
}

func (v NullableServiceDocsChannelGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDocsChannelGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


