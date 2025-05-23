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

// checks if the ServiceChannelRunActionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceChannelRunActionData{}

// ServiceChannelRunActionData struct for ServiceChannelRunActionData
type ServiceChannelRunActionData struct {
	Action string `json:"action"`
	Moh *string `json:"moh,omitempty"`
	TakebackDtmf *string `json:"takeback_dtmf,omitempty"`
	Target *string `json:"target,omitempty"`
}

type _ServiceChannelRunActionData ServiceChannelRunActionData

// NewServiceChannelRunActionData instantiates a new ServiceChannelRunActionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceChannelRunActionData(action string) *ServiceChannelRunActionData {
	this := ServiceChannelRunActionData{}
	this.Action = action
	return &this
}

// NewServiceChannelRunActionDataWithDefaults instantiates a new ServiceChannelRunActionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceChannelRunActionDataWithDefaults() *ServiceChannelRunActionData {
	this := ServiceChannelRunActionData{}
	return &this
}

// GetAction returns the Action field value
func (o *ServiceChannelRunActionData) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ServiceChannelRunActionData) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ServiceChannelRunActionData) SetAction(v string) {
	o.Action = v
}

// GetMoh returns the Moh field value if set, zero value otherwise.
func (o *ServiceChannelRunActionData) GetMoh() string {
	if o == nil || IsNil(o.Moh) {
		var ret string
		return ret
	}
	return *o.Moh
}

// GetMohOk returns a tuple with the Moh field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceChannelRunActionData) GetMohOk() (*string, bool) {
	if o == nil || IsNil(o.Moh) {
		return nil, false
	}
	return o.Moh, true
}

// HasMoh returns a boolean if a field has been set.
func (o *ServiceChannelRunActionData) HasMoh() bool {
	if o != nil && !IsNil(o.Moh) {
		return true
	}

	return false
}

// SetMoh gets a reference to the given string and assigns it to the Moh field.
func (o *ServiceChannelRunActionData) SetMoh(v string) {
	o.Moh = &v
}

// GetTakebackDtmf returns the TakebackDtmf field value if set, zero value otherwise.
func (o *ServiceChannelRunActionData) GetTakebackDtmf() string {
	if o == nil || IsNil(o.TakebackDtmf) {
		var ret string
		return ret
	}
	return *o.TakebackDtmf
}

// GetTakebackDtmfOk returns a tuple with the TakebackDtmf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceChannelRunActionData) GetTakebackDtmfOk() (*string, bool) {
	if o == nil || IsNil(o.TakebackDtmf) {
		return nil, false
	}
	return o.TakebackDtmf, true
}

// HasTakebackDtmf returns a boolean if a field has been set.
func (o *ServiceChannelRunActionData) HasTakebackDtmf() bool {
	if o != nil && !IsNil(o.TakebackDtmf) {
		return true
	}

	return false
}

// SetTakebackDtmf gets a reference to the given string and assigns it to the TakebackDtmf field.
func (o *ServiceChannelRunActionData) SetTakebackDtmf(v string) {
	o.TakebackDtmf = &v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ServiceChannelRunActionData) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceChannelRunActionData) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ServiceChannelRunActionData) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *ServiceChannelRunActionData) SetTarget(v string) {
	o.Target = &v
}

func (o ServiceChannelRunActionData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceChannelRunActionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	if !IsNil(o.Moh) {
		toSerialize["moh"] = o.Moh
	}
	if !IsNil(o.TakebackDtmf) {
		toSerialize["takeback_dtmf"] = o.TakebackDtmf
	}
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	return toSerialize, nil
}

func (o *ServiceChannelRunActionData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"action",
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

	varServiceChannelRunActionData := _ServiceChannelRunActionData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varServiceChannelRunActionData)

	if err != nil {
		return err
	}

	*o = ServiceChannelRunActionData(varServiceChannelRunActionData)

	return err
}

type NullableServiceChannelRunActionData struct {
	value *ServiceChannelRunActionData
	isSet bool
}

func (v NullableServiceChannelRunActionData) Get() *ServiceChannelRunActionData {
	return v.value
}

func (v *NullableServiceChannelRunActionData) Set(val *ServiceChannelRunActionData) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceChannelRunActionData) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceChannelRunActionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceChannelRunActionData(val *ServiceChannelRunActionData) *NullableServiceChannelRunActionData {
	return &NullableServiceChannelRunActionData{value: val, isSet: true}
}

func (v NullableServiceChannelRunActionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceChannelRunActionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


