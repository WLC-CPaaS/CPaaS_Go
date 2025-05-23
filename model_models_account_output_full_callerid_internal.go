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

// checks if the ModelsAccountOutputFullCalleridInternal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelsAccountOutputFullCalleridInternal{}

// ModelsAccountOutputFullCalleridInternal struct for ModelsAccountOutputFullCalleridInternal
type ModelsAccountOutputFullCalleridInternal struct {
	Number *string `json:"number,omitempty"`
}

// NewModelsAccountOutputFullCalleridInternal instantiates a new ModelsAccountOutputFullCalleridInternal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelsAccountOutputFullCalleridInternal() *ModelsAccountOutputFullCalleridInternal {
	this := ModelsAccountOutputFullCalleridInternal{}
	return &this
}

// NewModelsAccountOutputFullCalleridInternalWithDefaults instantiates a new ModelsAccountOutputFullCalleridInternal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelsAccountOutputFullCalleridInternalWithDefaults() *ModelsAccountOutputFullCalleridInternal {
	this := ModelsAccountOutputFullCalleridInternal{}
	return &this
}

// GetNumber returns the Number field value if set, zero value otherwise.
func (o *ModelsAccountOutputFullCalleridInternal) GetNumber() string {
	if o == nil || IsNil(o.Number) {
		var ret string
		return ret
	}
	return *o.Number
}

// GetNumberOk returns a tuple with the Number field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ModelsAccountOutputFullCalleridInternal) GetNumberOk() (*string, bool) {
	if o == nil || IsNil(o.Number) {
		return nil, false
	}
	return o.Number, true
}

// HasNumber returns a boolean if a field has been set.
func (o *ModelsAccountOutputFullCalleridInternal) HasNumber() bool {
	if o != nil && !IsNil(o.Number) {
		return true
	}

	return false
}

// SetNumber gets a reference to the given string and assigns it to the Number field.
func (o *ModelsAccountOutputFullCalleridInternal) SetNumber(v string) {
	o.Number = &v
}

func (o ModelsAccountOutputFullCalleridInternal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelsAccountOutputFullCalleridInternal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Number) {
		toSerialize["number"] = o.Number
	}
	return toSerialize, nil
}

type NullableModelsAccountOutputFullCalleridInternal struct {
	value *ModelsAccountOutputFullCalleridInternal
	isSet bool
}

func (v NullableModelsAccountOutputFullCalleridInternal) Get() *ModelsAccountOutputFullCalleridInternal {
	return v.value
}

func (v *NullableModelsAccountOutputFullCalleridInternal) Set(val *ModelsAccountOutputFullCalleridInternal) {
	v.value = val
	v.isSet = true
}

func (v NullableModelsAccountOutputFullCalleridInternal) IsSet() bool {
	return v.isSet
}

func (v *NullableModelsAccountOutputFullCalleridInternal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelsAccountOutputFullCalleridInternal(val *ModelsAccountOutputFullCalleridInternal) *NullableModelsAccountOutputFullCalleridInternal {
	return &NullableModelsAccountOutputFullCalleridInternal{value: val, isSet: true}
}

func (v NullableModelsAccountOutputFullCalleridInternal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelsAccountOutputFullCalleridInternal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


