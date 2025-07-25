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

// checks if the ServiceVOIPDeviceAddEditProvision type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceVOIPDeviceAddEditProvision{}

// ServiceVOIPDeviceAddEditProvision struct for ServiceVOIPDeviceAddEditProvision
type ServiceVOIPDeviceAddEditProvision struct {
	EndpointBrand *string `json:"endpoint_brand,omitempty"`
	EndpointFamily *string `json:"endpoint_family,omitempty"`
	EndpointModel *string `json:"endpoint_model,omitempty"`
	Id *string `json:"id,omitempty"`
	LineKeys []ServiceVOIPDeviceAddEditLineKey `json:"line_keys,omitempty"`
}

// NewServiceVOIPDeviceAddEditProvision instantiates a new ServiceVOIPDeviceAddEditProvision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceVOIPDeviceAddEditProvision() *ServiceVOIPDeviceAddEditProvision {
	this := ServiceVOIPDeviceAddEditProvision{}
	return &this
}

// NewServiceVOIPDeviceAddEditProvisionWithDefaults instantiates a new ServiceVOIPDeviceAddEditProvision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceVOIPDeviceAddEditProvisionWithDefaults() *ServiceVOIPDeviceAddEditProvision {
	this := ServiceVOIPDeviceAddEditProvision{}
	return &this
}

// GetEndpointBrand returns the EndpointBrand field value if set, zero value otherwise.
func (o *ServiceVOIPDeviceAddEditProvision) GetEndpointBrand() string {
	if o == nil || IsNil(o.EndpointBrand) {
		var ret string
		return ret
	}
	return *o.EndpointBrand
}

// GetEndpointBrandOk returns a tuple with the EndpointBrand field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPDeviceAddEditProvision) GetEndpointBrandOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointBrand) {
		return nil, false
	}
	return o.EndpointBrand, true
}

// HasEndpointBrand returns a boolean if a field has been set.
func (o *ServiceVOIPDeviceAddEditProvision) HasEndpointBrand() bool {
	if o != nil && !IsNil(o.EndpointBrand) {
		return true
	}

	return false
}

// SetEndpointBrand gets a reference to the given string and assigns it to the EndpointBrand field.
func (o *ServiceVOIPDeviceAddEditProvision) SetEndpointBrand(v string) {
	o.EndpointBrand = &v
}

// GetEndpointFamily returns the EndpointFamily field value if set, zero value otherwise.
func (o *ServiceVOIPDeviceAddEditProvision) GetEndpointFamily() string {
	if o == nil || IsNil(o.EndpointFamily) {
		var ret string
		return ret
	}
	return *o.EndpointFamily
}

// GetEndpointFamilyOk returns a tuple with the EndpointFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPDeviceAddEditProvision) GetEndpointFamilyOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointFamily) {
		return nil, false
	}
	return o.EndpointFamily, true
}

// HasEndpointFamily returns a boolean if a field has been set.
func (o *ServiceVOIPDeviceAddEditProvision) HasEndpointFamily() bool {
	if o != nil && !IsNil(o.EndpointFamily) {
		return true
	}

	return false
}

// SetEndpointFamily gets a reference to the given string and assigns it to the EndpointFamily field.
func (o *ServiceVOIPDeviceAddEditProvision) SetEndpointFamily(v string) {
	o.EndpointFamily = &v
}

// GetEndpointModel returns the EndpointModel field value if set, zero value otherwise.
func (o *ServiceVOIPDeviceAddEditProvision) GetEndpointModel() string {
	if o == nil || IsNil(o.EndpointModel) {
		var ret string
		return ret
	}
	return *o.EndpointModel
}

// GetEndpointModelOk returns a tuple with the EndpointModel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPDeviceAddEditProvision) GetEndpointModelOk() (*string, bool) {
	if o == nil || IsNil(o.EndpointModel) {
		return nil, false
	}
	return o.EndpointModel, true
}

// HasEndpointModel returns a boolean if a field has been set.
func (o *ServiceVOIPDeviceAddEditProvision) HasEndpointModel() bool {
	if o != nil && !IsNil(o.EndpointModel) {
		return true
	}

	return false
}

// SetEndpointModel gets a reference to the given string and assigns it to the EndpointModel field.
func (o *ServiceVOIPDeviceAddEditProvision) SetEndpointModel(v string) {
	o.EndpointModel = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceVOIPDeviceAddEditProvision) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPDeviceAddEditProvision) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceVOIPDeviceAddEditProvision) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceVOIPDeviceAddEditProvision) SetId(v string) {
	o.Id = &v
}

// GetLineKeys returns the LineKeys field value if set, zero value otherwise.
func (o *ServiceVOIPDeviceAddEditProvision) GetLineKeys() []ServiceVOIPDeviceAddEditLineKey {
	if o == nil || IsNil(o.LineKeys) {
		var ret []ServiceVOIPDeviceAddEditLineKey
		return ret
	}
	return o.LineKeys
}

// GetLineKeysOk returns a tuple with the LineKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVOIPDeviceAddEditProvision) GetLineKeysOk() ([]ServiceVOIPDeviceAddEditLineKey, bool) {
	if o == nil || IsNil(o.LineKeys) {
		return nil, false
	}
	return o.LineKeys, true
}

// HasLineKeys returns a boolean if a field has been set.
func (o *ServiceVOIPDeviceAddEditProvision) HasLineKeys() bool {
	if o != nil && !IsNil(o.LineKeys) {
		return true
	}

	return false
}

// SetLineKeys gets a reference to the given []ServiceVOIPDeviceAddEditLineKey and assigns it to the LineKeys field.
func (o *ServiceVOIPDeviceAddEditProvision) SetLineKeys(v []ServiceVOIPDeviceAddEditLineKey) {
	o.LineKeys = v
}

func (o ServiceVOIPDeviceAddEditProvision) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceVOIPDeviceAddEditProvision) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EndpointBrand) {
		toSerialize["endpoint_brand"] = o.EndpointBrand
	}
	if !IsNil(o.EndpointFamily) {
		toSerialize["endpoint_family"] = o.EndpointFamily
	}
	if !IsNil(o.EndpointModel) {
		toSerialize["endpoint_model"] = o.EndpointModel
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.LineKeys) {
		toSerialize["line_keys"] = o.LineKeys
	}
	return toSerialize, nil
}

type NullableServiceVOIPDeviceAddEditProvision struct {
	value *ServiceVOIPDeviceAddEditProvision
	isSet bool
}

func (v NullableServiceVOIPDeviceAddEditProvision) Get() *ServiceVOIPDeviceAddEditProvision {
	return v.value
}

func (v *NullableServiceVOIPDeviceAddEditProvision) Set(val *ServiceVOIPDeviceAddEditProvision) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceVOIPDeviceAddEditProvision) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceVOIPDeviceAddEditProvision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceVOIPDeviceAddEditProvision(val *ServiceVOIPDeviceAddEditProvision) *NullableServiceVOIPDeviceAddEditProvision {
	return &NullableServiceVOIPDeviceAddEditProvision{value: val, isSet: true}
}

func (v NullableServiceVOIPDeviceAddEditProvision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceVOIPDeviceAddEditProvision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


