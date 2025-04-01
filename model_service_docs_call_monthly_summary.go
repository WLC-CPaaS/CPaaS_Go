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

// checks if the ServiceDocsCallMonthlySummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceDocsCallMonthlySummary{}

// ServiceDocsCallMonthlySummary struct for ServiceDocsCallMonthlySummary
type ServiceDocsCallMonthlySummary struct {
	Data []ModelCallMonthlySummary `json:"data,omitempty"`
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

// NewServiceDocsCallMonthlySummary instantiates a new ServiceDocsCallMonthlySummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDocsCallMonthlySummary() *ServiceDocsCallMonthlySummary {
	this := ServiceDocsCallMonthlySummary{}
	return &this
}

// NewServiceDocsCallMonthlySummaryWithDefaults instantiates a new ServiceDocsCallMonthlySummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDocsCallMonthlySummaryWithDefaults() *ServiceDocsCallMonthlySummary {
	this := ServiceDocsCallMonthlySummary{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *ServiceDocsCallMonthlySummary) GetData() []ModelCallMonthlySummary {
	if o == nil || IsNil(o.Data) {
		var ret []ModelCallMonthlySummary
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsCallMonthlySummary) GetDataOk() ([]ModelCallMonthlySummary, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *ServiceDocsCallMonthlySummary) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []ModelCallMonthlySummary and assigns it to the Data field.
func (o *ServiceDocsCallMonthlySummary) SetData(v []ModelCallMonthlySummary) {
	o.Data = v
}

// GetNextStartKey returns the NextStartKey field value if set, zero value otherwise.
func (o *ServiceDocsCallMonthlySummary) GetNextStartKey() string {
	if o == nil || IsNil(o.NextStartKey) {
		var ret string
		return ret
	}
	return *o.NextStartKey
}

// GetNextStartKeyOk returns a tuple with the NextStartKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsCallMonthlySummary) GetNextStartKeyOk() (*string, bool) {
	if o == nil || IsNil(o.NextStartKey) {
		return nil, false
	}
	return o.NextStartKey, true
}

// HasNextStartKey returns a boolean if a field has been set.
func (o *ServiceDocsCallMonthlySummary) HasNextStartKey() bool {
	if o != nil && !IsNil(o.NextStartKey) {
		return true
	}

	return false
}

// SetNextStartKey gets a reference to the given string and assigns it to the NextStartKey field.
func (o *ServiceDocsCallMonthlySummary) SetNextStartKey(v string) {
	o.NextStartKey = &v
}

// GetPageSize returns the PageSize field value if set, zero value otherwise.
func (o *ServiceDocsCallMonthlySummary) GetPageSize() int32 {
	if o == nil || IsNil(o.PageSize) {
		var ret int32
		return ret
	}
	return *o.PageSize
}

// GetPageSizeOk returns a tuple with the PageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsCallMonthlySummary) GetPageSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.PageSize) {
		return nil, false
	}
	return o.PageSize, true
}

// HasPageSize returns a boolean if a field has been set.
func (o *ServiceDocsCallMonthlySummary) HasPageSize() bool {
	if o != nil && !IsNil(o.PageSize) {
		return true
	}

	return false
}

// SetPageSize gets a reference to the given int32 and assigns it to the PageSize field.
func (o *ServiceDocsCallMonthlySummary) SetPageSize(v int32) {
	o.PageSize = &v
}

// GetRequestId returns the RequestId field value if set, zero value otherwise.
func (o *ServiceDocsCallMonthlySummary) GetRequestId() string {
	if o == nil || IsNil(o.RequestId) {
		var ret string
		return ret
	}
	return *o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsCallMonthlySummary) GetRequestIdOk() (*string, bool) {
	if o == nil || IsNil(o.RequestId) {
		return nil, false
	}
	return o.RequestId, true
}

// HasRequestId returns a boolean if a field has been set.
func (o *ServiceDocsCallMonthlySummary) HasRequestId() bool {
	if o != nil && !IsNil(o.RequestId) {
		return true
	}

	return false
}

// SetRequestId gets a reference to the given string and assigns it to the RequestId field.
func (o *ServiceDocsCallMonthlySummary) SetRequestId(v string) {
	o.RequestId = &v
}

// GetStartKey returns the StartKey field value if set, zero value otherwise.
func (o *ServiceDocsCallMonthlySummary) GetStartKey() string {
	if o == nil || IsNil(o.StartKey) {
		var ret string
		return ret
	}
	return *o.StartKey
}

// GetStartKeyOk returns a tuple with the StartKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsCallMonthlySummary) GetStartKeyOk() (*string, bool) {
	if o == nil || IsNil(o.StartKey) {
		return nil, false
	}
	return o.StartKey, true
}

// HasStartKey returns a boolean if a field has been set.
func (o *ServiceDocsCallMonthlySummary) HasStartKey() bool {
	if o != nil && !IsNil(o.StartKey) {
		return true
	}

	return false
}

// SetStartKey gets a reference to the given string and assigns it to the StartKey field.
func (o *ServiceDocsCallMonthlySummary) SetStartKey(v string) {
	o.StartKey = &v
}

// GetStatusCode returns the StatusCode field value if set, zero value otherwise.
func (o *ServiceDocsCallMonthlySummary) GetStatusCode() int32 {
	if o == nil || IsNil(o.StatusCode) {
		var ret int32
		return ret
	}
	return *o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceDocsCallMonthlySummary) GetStatusCodeOk() (*int32, bool) {
	if o == nil || IsNil(o.StatusCode) {
		return nil, false
	}
	return o.StatusCode, true
}

// HasStatusCode returns a boolean if a field has been set.
func (o *ServiceDocsCallMonthlySummary) HasStatusCode() bool {
	if o != nil && !IsNil(o.StatusCode) {
		return true
	}

	return false
}

// SetStatusCode gets a reference to the given int32 and assigns it to the StatusCode field.
func (o *ServiceDocsCallMonthlySummary) SetStatusCode(v int32) {
	o.StatusCode = &v
}

func (o ServiceDocsCallMonthlySummary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceDocsCallMonthlySummary) ToMap() (map[string]interface{}, error) {
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

type NullableServiceDocsCallMonthlySummary struct {
	value *ServiceDocsCallMonthlySummary
	isSet bool
}

func (v NullableServiceDocsCallMonthlySummary) Get() *ServiceDocsCallMonthlySummary {
	return v.value
}

func (v *NullableServiceDocsCallMonthlySummary) Set(val *ServiceDocsCallMonthlySummary) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDocsCallMonthlySummary) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDocsCallMonthlySummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDocsCallMonthlySummary(val *ServiceDocsCallMonthlySummary) *NullableServiceDocsCallMonthlySummary {
	return &NullableServiceDocsCallMonthlySummary{value: val, isSet: true}
}

func (v NullableServiceDocsCallMonthlySummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDocsCallMonthlySummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


