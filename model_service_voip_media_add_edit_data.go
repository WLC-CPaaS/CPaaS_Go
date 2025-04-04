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
	"bytes"
	"fmt"
)

// checks if the ServiceVOIPMediaAddEditData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceVOIPMediaAddEditData{}

// ServiceVOIPMediaAddEditData struct for ServiceVOIPMediaAddEditData
type ServiceVOIPMediaAddEditData struct {
	Description *string `json:"description,omitempty"`
	MediaSource *string `json:"media_source,omitempty"`
	Name string `json:"name"`
	Tts *ServiceTTS `json:"tts,omitempty"`
}

type _ServiceVOIPMediaAddEditData ServiceVOIPMediaAddEditData

// NewServiceVOIPMediaAddEditData instantiates a new ServiceVOIPMediaAddEditData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceVOIPMediaAddEditData(name string) *ServiceVOIPMediaAddEditData {
	this := ServiceVOIPMediaAddEditData{}
	this.Name = name
	return &this
}

// NewServiceVOIPMediaAddEditDataWithDefaults instantiates a new ServiceVOIPMediaAddEditData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceVOIPMediaAddEditDataWithDefaults() *ServiceVOIPMediaAddEditData {
	this := ServiceVOIPMediaAddEditData{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ServiceVOIPMediaAddEditData) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPMediaAddEditData) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ServiceVOIPMediaAddEditData) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ServiceVOIPMediaAddEditData) SetDescription(v string) {
	o.Description = &v
}

// GetMediaSource returns the MediaSource field value if set, zero value otherwise.
func (o *ServiceVOIPMediaAddEditData) GetMediaSource() string {
	if o == nil || IsNil(o.MediaSource) {
		var ret string
		return ret
	}
	return *o.MediaSource
}

// GetMediaSourceOk returns a tuple with the MediaSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPMediaAddEditData) GetMediaSourceOk() (*string, bool) {
	if o == nil || IsNil(o.MediaSource) {
		return nil, false
	}
	return o.MediaSource, true
}

// HasMediaSource returns a boolean if a field has been set.
func (o *ServiceVOIPMediaAddEditData) HasMediaSource() bool {
	if o != nil && !IsNil(o.MediaSource) {
		return true
	}

	return false
}

// SetMediaSource gets a reference to the given string and assigns it to the MediaSource field.
func (o *ServiceVOIPMediaAddEditData) SetMediaSource(v string) {
	o.MediaSource = &v
}

// GetName returns the Name field value
func (o *ServiceVOIPMediaAddEditData) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ServiceVOIPMediaAddEditData) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ServiceVOIPMediaAddEditData) SetName(v string) {
	o.Name = v
}

// GetTts returns the Tts field value if set, zero value otherwise.
func (o *ServiceVOIPMediaAddEditData) GetTts() ServiceTTS {
	if o == nil || IsNil(o.Tts) {
		var ret ServiceTTS
		return ret
	}
	return *o.Tts
}

// GetTtsOk returns a tuple with the Tts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPMediaAddEditData) GetTtsOk() (*ServiceTTS, bool) {
	if o == nil || IsNil(o.Tts) {
		return nil, false
	}
	return o.Tts, true
}

// HasTts returns a boolean if a field has been set.
func (o *ServiceVOIPMediaAddEditData) HasTts() bool {
	if o != nil && !IsNil(o.Tts) {
		return true
	}

	return false
}

// SetTts gets a reference to the given ServiceTTS and assigns it to the Tts field.
func (o *ServiceVOIPMediaAddEditData) SetTts(v ServiceTTS) {
	o.Tts = &v
}

func (o ServiceVOIPMediaAddEditData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceVOIPMediaAddEditData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MediaSource) {
		toSerialize["media_source"] = o.MediaSource
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Tts) {
		toSerialize["tts"] = o.Tts
	}
	return toSerialize, nil
}

func (o *ServiceVOIPMediaAddEditData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varServiceVOIPMediaAddEditData := _ServiceVOIPMediaAddEditData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceVOIPMediaAddEditData)

	if err != nil {
		return err
	}

	*o = ServiceVOIPMediaAddEditData(varServiceVOIPMediaAddEditData)

	return err
}

type NullableServiceVOIPMediaAddEditData struct {
	value *ServiceVOIPMediaAddEditData
	isSet bool
}

func (v NullableServiceVOIPMediaAddEditData) Get() *ServiceVOIPMediaAddEditData {
	return v.value
}

func (v *NullableServiceVOIPMediaAddEditData) Set(val *ServiceVOIPMediaAddEditData) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceVOIPMediaAddEditData) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceVOIPMediaAddEditData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceVOIPMediaAddEditData(val *ServiceVOIPMediaAddEditData) *NullableServiceVOIPMediaAddEditData {
	return &NullableServiceVOIPMediaAddEditData{value: val, isSet: true}
}

func (v NullableServiceVOIPMediaAddEditData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceVOIPMediaAddEditData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


