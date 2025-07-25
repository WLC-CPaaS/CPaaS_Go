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

// checks if the ServiceDocsE911RemoveLocationOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceDocsE911RemoveLocationOutput{}

// ServiceDocsE911RemoveLocationOutput struct for ServiceDocsE911RemoveLocationOutput
type ServiceDocsE911RemoveLocationOutput struct {
	Data *ServiceE911RemoveLocationOutput `json:"data,omitempty"`
	// Unique id for each request
	RequestId *string `json:"request_id,omitempty"`
	// HTTP response status code
	StatusCode *int32 `json:"status_code,omitempty"`
}

// NewServiceDocsE911RemoveLocationOutput instantiates a new ServiceDocsE911RemoveLocationOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDocsE911RemoveLocationOutput() *ServiceDocsE911RemoveLocationOutput {
	this := ServiceDocsE911RemoveLocationOutput{}
	return &this
}

// NewServiceDocsE911RemoveLocationOutputWithDefaults instantiates a new ServiceDocsE911RemoveLocationOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDocsE911RemoveLocationOutputWithDefaults() *ServiceDocsE911RemoveLocationOutput {
	this := ServiceDocsE911RemoveLocationOutput{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ServiceDocsE911RemoveLocationOutput) GetData() ServiceE911RemoveLocationOutput {
	if o == nil || IsNil(o.Data) {
		var ret ServiceE911RemoveLocationOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsE911RemoveLocationOutput) GetDataOk() (*ServiceE911RemoveLocationOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ServiceDocsE911RemoveLocationOutput) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ServiceE911RemoveLocationOutput and assigns it to the Data field.
func (o *ServiceDocsE911RemoveLocationOutput) SetData(v ServiceE911RemoveLocationOutput) {
	o.Data = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ServiceDocsE911RemoveLocationOutput) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsE911RemoveLocationOutput) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ServiceDocsE911RemoveLocationOutput) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ServiceDocsE911RemoveLocationOutput) SetRequestId(v string) {
	o.RequestId = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ServiceDocsE911RemoveLocationOutput) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsE911RemoveLocationOutput) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ServiceDocsE911RemoveLocationOutput) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *ServiceDocsE911RemoveLocationOutput) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o ServiceDocsE911RemoveLocationOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceDocsE911RemoveLocationOutput) ToMap() (map[string]interface{}, error) {
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

type NullableServiceDocsE911RemoveLocationOutput struct {
	value *ServiceDocsE911RemoveLocationOutput
	isSet bool
}

func (v NullableServiceDocsE911RemoveLocationOutput) Get() *ServiceDocsE911RemoveLocationOutput {
	return v.value
}

func (v *NullableServiceDocsE911RemoveLocationOutput) Set(val *ServiceDocsE911RemoveLocationOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDocsE911RemoveLocationOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDocsE911RemoveLocationOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDocsE911RemoveLocationOutput(val *ServiceDocsE911RemoveLocationOutput) *NullableServiceDocsE911RemoveLocationOutput {
	return &NullableServiceDocsE911RemoveLocationOutput{value: val, isSet: true}
}

func (v NullableServiceDocsE911RemoveLocationOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDocsE911RemoveLocationOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


