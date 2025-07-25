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

// checks if the ProvisioningDocsDocsModelOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningDocsDocsModelOutput{}

// ProvisioningDocsDocsModelOutput struct for ProvisioningDocsDocsModelOutput
type ProvisioningDocsDocsModelOutput struct {
	Data []ModelsModel `json:"data,omitempty"`
	// List Pagination: Used to get the next page of results. Will not exist if this is the last page.
	NextStartKey *string `json:"next_start_key,omitempty"`
	// List Pagination: The number of results returned in this page
	PageSize *int32 `json:"page_size,omitempty"`
	// Unique id for each request
	RequestId *string `json:"request_id,omitempty"`
	// List Pagination: Code for paged results
	StartKey *string `json:"start_key,omitempty"`
	// HTTP response status code
	StatusCode *int32 `json:"status_code,omitempty"`
}

// NewProvisioningDocsDocsModelOutput instantiates a new ProvisioningDocsDocsModelOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningDocsDocsModelOutput() *ProvisioningDocsDocsModelOutput {
	this := ProvisioningDocsDocsModelOutput{}
	return &this
}

// NewProvisioningDocsDocsModelOutputWithDefaults instantiates a new ProvisioningDocsDocsModelOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningDocsDocsModelOutputWithDefaults() *ProvisioningDocsDocsModelOutput {
	this := ProvisioningDocsDocsModelOutput{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ProvisioningDocsDocsModelOutput) GetData() []ModelsModel {
	if o == nil || IsNil(o.Data) {
		var ret []ModelsModel
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningDocsDocsModelOutput) GetDataOk() ([]ModelsModel, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ProvisioningDocsDocsModelOutput) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ModelsModel and assigns it to the Data field.
func (o *ProvisioningDocsDocsModelOutput) SetData(v []ModelsModel) {
	o.Data = v
}

// GetNextStartKey returns the NextStartKey field value if set, zero value otherwise.
func (o *ProvisioningDocsDocsModelOutput) GetNextStartKey() string {
	if o == nil || IsNil(o.NextStartKey) {
		var ret string
		return ret
	}
	return *o.NextStartKey
}

// GetNextStartKeyOk returns a tuple with the NextStartKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningDocsDocsModelOutput) GetNextStartKeyOk() (*string, bool) {
	if o == nil || IsNil(o.NextStartKey) {
		return nil, false
	}
	return o.NextStartKey, true
}

// HasNextStartKey returns a boolean if a field has been set.
func (o *ProvisioningDocsDocsModelOutput) HasNextStartKey() bool {
	if o != nil && !IsNil(o.NextStartKey) {
		return true
	}

	return false
}

// SetNextStartKey gets a reference to the given string and assigns it to the NextStartKey field.
func (o *ProvisioningDocsDocsModelOutput) SetNextStartKey(v string) {
	o.NextStartKey = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ProvisioningDocsDocsModelOutput) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningDocsDocsModelOutput) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ProvisioningDocsDocsModelOutput) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ProvisioningDocsDocsModelOutput) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ProvisioningDocsDocsModelOutput) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningDocsDocsModelOutput) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ProvisioningDocsDocsModelOutput) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ProvisioningDocsDocsModelOutput) SetRequestId(v string) {
	o.RequestId = &v
}

// GetStartKey returns the StartKey field value if set, zero value otherwise.
func (o *ProvisioningDocsDocsModelOutput) GetStartKey() string {
	if o == nil || IsNil(o.StartKey) {
		var ret string
		return ret
	}
	return *o.StartKey
}

// GetStartKeyOk returns a tuple with the StartKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningDocsDocsModelOutput) GetStartKeyOk() (*string, bool) {
	if o == nil || IsNil(o.StartKey) {
		return nil, false
	}
	return o.StartKey, true
}

// HasStartKey returns a boolean if a field has been set.
func (o *ProvisioningDocsDocsModelOutput) HasStartKey() bool {
	if o != nil && !IsNil(o.StartKey) {
		return true
	}

	return false
}

// SetStartKey gets a reference to the given string and assigns it to the StartKey field.
func (o *ProvisioningDocsDocsModelOutput) SetStartKey(v string) {
	o.StartKey = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ProvisioningDocsDocsModelOutput) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningDocsDocsModelOutput) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ProvisioningDocsDocsModelOutput) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *ProvisioningDocsDocsModelOutput) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o ProvisioningDocsDocsModelOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningDocsDocsModelOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	if !IsNil(o.NextStartKey) {
		toSerialize["next_start_key"] = o.NextStartKey
	}
	if !IsNil(o.PageSize) {
		toSerialize["page_size"] = o.PageSize
	}
	if !IsNil(o.RequestId) {
		toSerialize["request_id"] = o.RequestId
	}
	if !IsNil(o.StartKey) {
		toSerialize["start_key"] = o.StartKey
	}
	if !IsNil(o.StatusCode) {
		toSerialize["status_code"] = o.StatusCode
	}
	return toSerialize, nil
}

type NullableProvisioningDocsDocsModelOutput struct {
	value *ProvisioningDocsDocsModelOutput
	isSet bool
}

func (v NullableProvisioningDocsDocsModelOutput) Get() *ProvisioningDocsDocsModelOutput {
	return v.value
}

func (v *NullableProvisioningDocsDocsModelOutput) Set(val *ProvisioningDocsDocsModelOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningDocsDocsModelOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningDocsDocsModelOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningDocsDocsModelOutput(val *ProvisioningDocsDocsModelOutput) *NullableProvisioningDocsDocsModelOutput {
	return &NullableProvisioningDocsDocsModelOutput{value: val, isSet: true}
}

func (v NullableProvisioningDocsDocsModelOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningDocsDocsModelOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


