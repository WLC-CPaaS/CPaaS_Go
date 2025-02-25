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

// checks if the ServiceStoragePlan type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceStoragePlan{}

// ServiceStoragePlan struct for ServiceStoragePlan
type ServiceStoragePlan struct {
	Account *ServiceStoragePlanDatabase `json:"account,omitempty"`
	Aggregate *ServiceStoragePlanDatabase `json:"aggregate,omitempty"`
	Modb *ServiceStoragePlanDatabase `json:"modb,omitempty"`
	System *ServiceStoragePlanDatabase `json:"system,omitempty"`
}

// NewServiceStoragePlan instantiates a new ServiceStoragePlan object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceStoragePlan() *ServiceStoragePlan {
	this := ServiceStoragePlan{}
	return &this
}

// NewServiceStoragePlanWithDefaults instantiates a new ServiceStoragePlan object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceStoragePlanWithDefaults() *ServiceStoragePlan {
	this := ServiceStoragePlan{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *ServiceStoragePlan) GetAccount() ServiceStoragePlanDatabase {
	if o == nil || IsNil(o.Account) {
		var ret ServiceStoragePlanDatabase
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceStoragePlan) GetAccountOk() (*ServiceStoragePlanDatabase, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *ServiceStoragePlan) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given ServiceStoragePlanDatabase and assigns it to the Account field.
func (o *ServiceStoragePlan) SetAccount(v ServiceStoragePlanDatabase) {
	o.Account = &v
}

// GetAggregate returns the Aggregate field value if set, zero value otherwise.
func (o *ServiceStoragePlan) GetAggregate() ServiceStoragePlanDatabase {
	if o == nil || IsNil(o.Aggregate) {
		var ret ServiceStoragePlanDatabase
		return ret
	}
	return *o.Aggregate
}

// GetAggregateOk returns a tuple with the Aggregate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceStoragePlan) GetAggregateOk() (*ServiceStoragePlanDatabase, bool) {
	if o == nil || IsNil(o.Aggregate) {
		return nil, false
	}
	return o.Aggregate, true
}

// HasAggregate returns a boolean if a field has been set.
func (o *ServiceStoragePlan) HasAggregate() bool {
	if o != nil && !IsNil(o.Aggregate) {
		return true
	}

	return false
}

// SetAggregate gets a reference to the given ServiceStoragePlanDatabase and assigns it to the Aggregate field.
func (o *ServiceStoragePlan) SetAggregate(v ServiceStoragePlanDatabase) {
	o.Aggregate = &v
}

// GetModb returns the Modb field value if set, zero value otherwise.
func (o *ServiceStoragePlan) GetModb() ServiceStoragePlanDatabase {
	if o == nil || IsNil(o.Modb) {
		var ret ServiceStoragePlanDatabase
		return ret
	}
	return *o.Modb
}

// GetModbOk returns a tuple with the Modb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceStoragePlan) GetModbOk() (*ServiceStoragePlanDatabase, bool) {
	if o == nil || IsNil(o.Modb) {
		return nil, false
	}
	return o.Modb, true
}

// HasModb returns a boolean if a field has been set.
func (o *ServiceStoragePlan) HasModb() bool {
	if o != nil && !IsNil(o.Modb) {
		return true
	}

	return false
}

// SetModb gets a reference to the given ServiceStoragePlanDatabase and assigns it to the Modb field.
func (o *ServiceStoragePlan) SetModb(v ServiceStoragePlanDatabase) {
	o.Modb = &v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *ServiceStoragePlan) GetSystem() ServiceStoragePlanDatabase {
	if o == nil || IsNil(o.System) {
		var ret ServiceStoragePlanDatabase
		return ret
	}
	return *o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceStoragePlan) GetSystemOk() (*ServiceStoragePlanDatabase, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *ServiceStoragePlan) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given ServiceStoragePlanDatabase and assigns it to the System field.
func (o *ServiceStoragePlan) SetSystem(v ServiceStoragePlanDatabase) {
	o.System = &v
}

func (o ServiceStoragePlan) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceStoragePlan) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Aggregate) {
		toSerialize["aggregate"] = o.Aggregate
	}
	if !IsNil(o.Modb) {
		toSerialize["modb"] = o.Modb
	}
	if !IsNil(o.System) {
		toSerialize["system"] = o.System
	}
	return toSerialize, nil
}

type NullableServiceStoragePlan struct {
	value *ServiceStoragePlan
	isSet bool
}

func (v NullableServiceStoragePlan) Get() *ServiceStoragePlan {
	return v.value
}

func (v *NullableServiceStoragePlan) Set(val *ServiceStoragePlan) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceStoragePlan) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceStoragePlan) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceStoragePlan(val *ServiceStoragePlan) *NullableServiceStoragePlan {
	return &NullableServiceStoragePlan{value: val, isSet: true}
}

func (v NullableServiceStoragePlan) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceStoragePlan) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


