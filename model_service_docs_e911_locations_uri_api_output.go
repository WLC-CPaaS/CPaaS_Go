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

// checks if the ServiceDocsE911LocationsURIApiOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceDocsE911LocationsURIApiOutput{}

// ServiceDocsE911LocationsURIApiOutput struct for ServiceDocsE911LocationsURIApiOutput
type ServiceDocsE911LocationsURIApiOutput struct {
	Data []ServiceE911LocationURI `json:"data,omitempty"`
	// Unique id for each request
	RequestId *string `json:"request_id,omitempty"`
	// HTTP response status code
	StatusCode *int32 `json:"status_code,omitempty"`
}

// NewServiceDocsE911LocationsURIApiOutput instantiates a new ServiceDocsE911LocationsURIApiOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDocsE911LocationsURIApiOutput() *ServiceDocsE911LocationsURIApiOutput {
	this := ServiceDocsE911LocationsURIApiOutput{}
	return &this
}

// NewServiceDocsE911LocationsURIApiOutputWithDefaults instantiates a new ServiceDocsE911LocationsURIApiOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDocsE911LocationsURIApiOutputWithDefaults() *ServiceDocsE911LocationsURIApiOutput {
	this := ServiceDocsE911LocationsURIApiOutput{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ServiceDocsE911LocationsURIApiOutput) GetData() []ServiceE911LocationURI {
	if o == nil || IsNil(o.Data) {
		var ret []ServiceE911LocationURI
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsE911LocationsURIApiOutput) GetDataOk() ([]ServiceE911LocationURI, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ServiceDocsE911LocationsURIApiOutput) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ServiceE911LocationURI and assigns it to the Data field.
func (o *ServiceDocsE911LocationsURIApiOutput) SetData(v []ServiceE911LocationURI) {
	o.Data = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ServiceDocsE911LocationsURIApiOutput) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsE911LocationsURIApiOutput) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ServiceDocsE911LocationsURIApiOutput) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ServiceDocsE911LocationsURIApiOutput) SetRequestId(v string) {
	o.RequestId = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ServiceDocsE911LocationsURIApiOutput) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsE911LocationsURIApiOutput) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ServiceDocsE911LocationsURIApiOutput) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *ServiceDocsE911LocationsURIApiOutput) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o ServiceDocsE911LocationsURIApiOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceDocsE911LocationsURIApiOutput) ToMap() (map[string]interface{}, error) {
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

type NullableServiceDocsE911LocationsURIApiOutput struct {
	value *ServiceDocsE911LocationsURIApiOutput
	isSet bool
}

func (v NullableServiceDocsE911LocationsURIApiOutput) Get() *ServiceDocsE911LocationsURIApiOutput {
	return v.value
}

func (v *NullableServiceDocsE911LocationsURIApiOutput) Set(val *ServiceDocsE911LocationsURIApiOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDocsE911LocationsURIApiOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDocsE911LocationsURIApiOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDocsE911LocationsURIApiOutput(val *ServiceDocsE911LocationsURIApiOutput) *NullableServiceDocsE911LocationsURIApiOutput {
	return &NullableServiceDocsE911LocationsURIApiOutput{value: val, isSet: true}
}

func (v NullableServiceDocsE911LocationsURIApiOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDocsE911LocationsURIApiOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


