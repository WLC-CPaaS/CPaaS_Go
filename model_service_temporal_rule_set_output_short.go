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

// checks if the ServiceTemporalRuleSetOutputShort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceTemporalRuleSetOutputShort{}

// ServiceTemporalRuleSetOutputShort struct for ServiceTemporalRuleSetOutputShort
type ServiceTemporalRuleSetOutputShort struct {
	Features []string `json:"features,omitempty"`
	Flags []string `json:"flags,omitempty"`
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Rules []string `json:"rules,omitempty"`
}

// NewServiceTemporalRuleSetOutputShort instantiates a new ServiceTemporalRuleSetOutputShort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceTemporalRuleSetOutputShort() *ServiceTemporalRuleSetOutputShort {
	this := ServiceTemporalRuleSetOutputShort{}
	return &this
}

// NewServiceTemporalRuleSetOutputShortWithDefaults instantiates a new ServiceTemporalRuleSetOutputShort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceTemporalRuleSetOutputShortWithDefaults() *ServiceTemporalRuleSetOutputShort {
	this := ServiceTemporalRuleSetOutputShort{}
	return &this
}

// GetFeatures returns the Features field value if set, zero value otherwise.
func (o *ServiceTemporalRuleSetOutputShort) GetFeatures() []string {
	if o == nil || IsNil(o.Features) {
		var ret []string
		return ret
	}
	return o.Features
}

// GetFeaturesOk returns a tuple with the Features field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTemporalRuleSetOutputShort) GetFeaturesOk() ([]string, bool) {
	if o == nil || IsNil(o.Features) {
		return nil, false
	}
	return o.Features, true
}

// HasFeatures returns a boolean if a field has been set.
func (o *ServiceTemporalRuleSetOutputShort) HasFeatures() bool {
	if o != nil && !IsNil(o.Features) {
		return true
	}

	return false
}

// SetFeatures gets a reference to the given []string and assigns it to the Features field.
func (o *ServiceTemporalRuleSetOutputShort) SetFeatures(v []string) {
	o.Features = v
}

// GetFlags returns the Flags field value if set, zero value otherwise.
func (o *ServiceTemporalRuleSetOutputShort) GetFlags() []string {
	if o == nil || IsNil(o.Flags) {
		var ret []string
		return ret
	}
	return o.Flags
}

// GetFlagsOk returns a tuple with the Flags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTemporalRuleSetOutputShort) GetFlagsOk() ([]string, bool) {
	if o == nil || IsNil(o.Flags) {
		return nil, false
	}
	return o.Flags, true
}

// HasFlags returns a boolean if a field has been set.
func (o *ServiceTemporalRuleSetOutputShort) HasFlags() bool {
	if o != nil && !IsNil(o.Flags) {
		return true
	}

	return false
}

// SetFlags gets a reference to the given []string and assigns it to the Flags field.
func (o *ServiceTemporalRuleSetOutputShort) SetFlags(v []string) {
	o.Flags = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceTemporalRuleSetOutputShort) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTemporalRuleSetOutputShort) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceTemporalRuleSetOutputShort) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceTemporalRuleSetOutputShort) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceTemporalRuleSetOutputShort) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTemporalRuleSetOutputShort) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceTemporalRuleSetOutputShort) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceTemporalRuleSetOutputShort) SetName(v string) {
	o.Name = &v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *ServiceTemporalRuleSetOutputShort) GetRules() []string {
	if o == nil || IsNil(o.Rules) {
		var ret []string
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceTemporalRuleSetOutputShort) GetRulesOk() ([]string, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *ServiceTemporalRuleSetOutputShort) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []string and assigns it to the Rules field.
func (o *ServiceTemporalRuleSetOutputShort) SetRules(v []string) {
	o.Rules = v
}

func (o ServiceTemporalRuleSetOutputShort) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceTemporalRuleSetOutputShort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Features) {
		toSerialize["features"] = o.Features
	}
	if !IsNil(o.Flags) {
		toSerialize["flags"] = o.Flags
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableServiceTemporalRuleSetOutputShort struct {
	value *ServiceTemporalRuleSetOutputShort
	isSet bool
}

func (v NullableServiceTemporalRuleSetOutputShort) Get() *ServiceTemporalRuleSetOutputShort {
	return v.value
}

func (v *NullableServiceTemporalRuleSetOutputShort) Set(val *ServiceTemporalRuleSetOutputShort) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTemporalRuleSetOutputShort) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTemporalRuleSetOutputShort) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTemporalRuleSetOutputShort(val *ServiceTemporalRuleSetOutputShort) *NullableServiceTemporalRuleSetOutputShort {
	return &NullableServiceTemporalRuleSetOutputShort{value: val, isSet: true}
}

func (v NullableServiceTemporalRuleSetOutputShort) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTemporalRuleSetOutputShort) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


