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

// checks if the ServiceDocsOrderPhonenumber type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceDocsOrderPhonenumber{}

// ServiceDocsOrderPhonenumber struct for ServiceDocsOrderPhonenumber
type ServiceDocsOrderPhonenumber struct {
	Data []ServicePhoneNumberResult `json:"data,omitempty"`
	// Unique id for each request
	RequestId *string `json:"request_id,omitempty"`
	// HTTP response status code
	StatusCode *int32 `json:"status_code,omitempty"`
}

// NewServiceDocsOrderPhonenumber instantiates a new ServiceDocsOrderPhonenumber object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDocsOrderPhonenumber() *ServiceDocsOrderPhonenumber {
	this := ServiceDocsOrderPhonenumber{}
	return &this
}

// NewServiceDocsOrderPhonenumberWithDefaults instantiates a new ServiceDocsOrderPhonenumber object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDocsOrderPhonenumberWithDefaults() *ServiceDocsOrderPhonenumber {
	this := ServiceDocsOrderPhonenumber{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ServiceDocsOrderPhonenumber) GetData() []ServicePhoneNumberResult {
	if o == nil || IsNil(o.Data) {
		var ret []ServicePhoneNumberResult
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsOrderPhonenumber) GetDataOk() ([]ServicePhoneNumberResult, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ServiceDocsOrderPhonenumber) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ServicePhoneNumberResult and assigns it to the Data field.
func (o *ServiceDocsOrderPhonenumber) SetData(v []ServicePhoneNumberResult) {
	o.Data = v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ServiceDocsOrderPhonenumber) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsOrderPhonenumber) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ServiceDocsOrderPhonenumber) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ServiceDocsOrderPhonenumber) SetRequestId(v string) {
	o.RequestId = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ServiceDocsOrderPhonenumber) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsOrderPhonenumber) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ServiceDocsOrderPhonenumber) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *ServiceDocsOrderPhonenumber) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o ServiceDocsOrderPhonenumber) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceDocsOrderPhonenumber) ToMap() (map[string]interface{}, error) {
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

type NullableServiceDocsOrderPhonenumber struct {
	value *ServiceDocsOrderPhonenumber
	isSet bool
}

func (v NullableServiceDocsOrderPhonenumber) Get() *ServiceDocsOrderPhonenumber {
	return v.value
}

func (v *NullableServiceDocsOrderPhonenumber) Set(val *ServiceDocsOrderPhonenumber) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDocsOrderPhonenumber) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDocsOrderPhonenumber) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDocsOrderPhonenumber(val *ServiceDocsOrderPhonenumber) *NullableServiceDocsOrderPhonenumber {
	return &NullableServiceDocsOrderPhonenumber{value: val, isSet: true}
}

func (v NullableServiceDocsOrderPhonenumber) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDocsOrderPhonenumber) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


