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

// checks if the ServiceDocsCampaignImportOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceDocsCampaignImportOutput{}

// ServiceDocsCampaignImportOutput struct for ServiceDocsCampaignImportOutput
type ServiceDocsCampaignImportOutput struct {
	Data *ServiceCampaignImportOutput `json:"data,omitempty"`
	RequestId *string `json:"request_id,omitempty"`
	StatusCode *int32 `json:"status_code,omitempty"`
}

// NewServiceDocsCampaignImportOutput instantiates a new ServiceDocsCampaignImportOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDocsCampaignImportOutput() *ServiceDocsCampaignImportOutput {
	this := ServiceDocsCampaignImportOutput{}
	return &this
}

// NewServiceDocsCampaignImportOutputWithDefaults instantiates a new ServiceDocsCampaignImportOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDocsCampaignImportOutputWithDefaults() *ServiceDocsCampaignImportOutput {
	this := ServiceDocsCampaignImportOutput{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ServiceDocsCampaignImportOutput) GetData() ServiceCampaignImportOutput {
	if o == nil || IsNil(o.Data) {
		var ret ServiceCampaignImportOutput
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsCampaignImportOutput) GetDataOk() (*ServiceCampaignImportOutput, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ServiceDocsCampaignImportOutput) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given ServiceCampaignImportOutput and assigns it to the Data field.
func (o *ServiceDocsCampaignImportOutput) SetData(v ServiceCampaignImportOutput) {
	o.Data = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ServiceDocsCampaignImportOutput) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsCampaignImportOutput) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ServiceDocsCampaignImportOutput) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ServiceDocsCampaignImportOutput) SetRequestId(v string) {
	o.RequestId = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ServiceDocsCampaignImportOutput) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsCampaignImportOutput) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ServiceDocsCampaignImportOutput) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *ServiceDocsCampaignImportOutput) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o ServiceDocsCampaignImportOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceDocsCampaignImportOutput) ToMap() (map[string]interface{}, error) {
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

type NullableServiceDocsCampaignImportOutput struct {
	value *ServiceDocsCampaignImportOutput
	isSet bool
}

func (v NullableServiceDocsCampaignImportOutput) Get() *ServiceDocsCampaignImportOutput {
	return v.value
}

func (v *NullableServiceDocsCampaignImportOutput) Set(val *ServiceDocsCampaignImportOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDocsCampaignImportOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDocsCampaignImportOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDocsCampaignImportOutput(val *ServiceDocsCampaignImportOutput) *NullableServiceDocsCampaignImportOutput {
	return &NullableServiceDocsCampaignImportOutput{value: val, isSet: true}
}

func (v NullableServiceDocsCampaignImportOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDocsCampaignImportOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


