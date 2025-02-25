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

// checks if the ServiceVoicemailMessageOutput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceVoicemailMessageOutput{}

// ServiceVoicemailMessageOutput struct for ServiceVoicemailMessageOutput
type ServiceVoicemailMessageOutput struct {
	CallId *string `json:"call_id,omitempty"`
	CallerIdName *string `json:"caller_id_name,omitempty"`
	CallerIdNumber *string `json:"caller_id_number,omitempty"`
	Folder *string `json:"folder,omitempty"`
	From *string `json:"from,omitempty"`
	Length *int32 `json:"length,omitempty"`
	MediaId *string `json:"media_id,omitempty"`
	Succeeded []string `json:"succeeded,omitempty"`
	Timestamp *int32 `json:"timestamp,omitempty"`
	To *string `json:"to,omitempty"`
}

// NewServiceVoicemailMessageOutput instantiates a new ServiceVoicemailMessageOutput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceVoicemailMessageOutput() *ServiceVoicemailMessageOutput {
	this := ServiceVoicemailMessageOutput{}
	return &this
}

// NewServiceVoicemailMessageOutputWithDefaults instantiates a new ServiceVoicemailMessageOutput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceVoicemailMessageOutputWithDefaults() *ServiceVoicemailMessageOutput {
	this := ServiceVoicemailMessageOutput{}
	return &this
}

// GetCallId returns the CallId field value if set, zero value otherwise.
func (o *ServiceVoicemailMessageOutput) GetCallId() string {
	if o == nil || IsNil(o.CallId) {
		var ret string
		return ret
	}
	return *o.CallId
}

// GetCallIdOk returns a tuple with the CallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVoicemailMessageOutput) GetCallIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallId) {
		return nil, false
	}
	return o.CallId, true
}

// HasCallId returns a boolean if a field has been set.
func (o *ServiceVoicemailMessageOutput) HasCallId() bool {
	if o != nil && !IsNil(o.CallId) {
		return true
	}

	return false
}

// SetCallId gets a reference to the given string and assigns it to the CallId field.
func (o *ServiceVoicemailMessageOutput) SetCallId(v string) {
	o.CallId = &v
}

// GetCallerIdName returns the CallerIdName field value if set, zero value otherwise.
func (o *ServiceVoicemailMessageOutput) GetCallerIdName() string {
	if o == nil || IsNil(o.CallerIdName) {
		var ret string
		return ret
	}
	return *o.CallerIdName
}

// GetCallerIdNameOk returns a tuple with the CallerIdName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVoicemailMessageOutput) GetCallerIdNameOk() (*string, bool) {
	if o == nil || IsNil(o.CallerIdName) {
		return nil, false
	}
	return o.CallerIdName, true
}

// HasCallerIdName returns a boolean if a field has been set.
func (o *ServiceVoicemailMessageOutput) HasCallerIdName() bool {
	if o != nil && !IsNil(o.CallerIdName) {
		return true
	}

	return false
}

// SetCallerIdName gets a reference to the given string and assigns it to the CallerIdName field.
func (o *ServiceVoicemailMessageOutput) SetCallerIdName(v string) {
	o.CallerIdName = &v
}

// GetCallerIdNumber returns the CallerIdNumber field value if set, zero value otherwise.
func (o *ServiceVoicemailMessageOutput) GetCallerIdNumber() string {
	if o == nil || IsNil(o.CallerIdNumber) {
		var ret string
		return ret
	}
	return *o.CallerIdNumber
}

// GetCallerIdNumberOk returns a tuple with the CallerIdNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVoicemailMessageOutput) GetCallerIdNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CallerIdNumber) {
		return nil, false
	}
	return o.CallerIdNumber, true
}

// HasCallerIdNumber returns a boolean if a field has been set.
func (o *ServiceVoicemailMessageOutput) HasCallerIdNumber() bool {
	if o != nil && !IsNil(o.CallerIdNumber) {
		return true
	}

	return false
}

// SetCallerIdNumber gets a reference to the given string and assigns it to the CallerIdNumber field.
func (o *ServiceVoicemailMessageOutput) SetCallerIdNumber(v string) {
	o.CallerIdNumber = &v
}

// GetFolder returns the Folder field value if set, zero value otherwise.
func (o *ServiceVoicemailMessageOutput) GetFolder() string {
	if o == nil || IsNil(o.Folder) {
		var ret string
		return ret
	}
	return *o.Folder
}

// GetFolderOk returns a tuple with the Folder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVoicemailMessageOutput) GetFolderOk() (*string, bool) {
	if o == nil || IsNil(o.Folder) {
		return nil, false
	}
	return o.Folder, true
}

// HasFolder returns a boolean if a field has been set.
func (o *ServiceVoicemailMessageOutput) HasFolder() bool {
	if o != nil && !IsNil(o.Folder) {
		return true
	}

	return false
}

// SetFolder gets a reference to the given string and assigns it to the Folder field.
func (o *ServiceVoicemailMessageOutput) SetFolder(v string) {
	o.Folder = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *ServiceVoicemailMessageOutput) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVoicemailMessageOutput) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *ServiceVoicemailMessageOutput) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *ServiceVoicemailMessageOutput) SetFrom(v string) {
	o.From = &v
}

// GetLength returns the Length field value if set, zero value otherwise.
func (o *ServiceVoicemailMessageOutput) GetLength() int32 {
	if o == nil || IsNil(o.Length) {
		var ret int32
		return ret
	}
	return *o.Length
}

// GetLengthOk returns a tuple with the Length field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVoicemailMessageOutput) GetLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.Length) {
		return nil, false
	}
	return o.Length, true
}

// HasLength returns a boolean if a field has been set.
func (o *ServiceVoicemailMessageOutput) HasLength() bool {
	if o != nil && !IsNil(o.Length) {
		return true
	}

	return false
}

// SetLength gets a reference to the given int32 and assigns it to the Length field.
func (o *ServiceVoicemailMessageOutput) SetLength(v int32) {
	o.Length = &v
}

// GetMediaId returns the MediaId field value if set, zero value otherwise.
func (o *ServiceVoicemailMessageOutput) GetMediaId() string {
	if o == nil || IsNil(o.MediaId) {
		var ret string
		return ret
	}
	return *o.MediaId
}

// GetMediaIdOk returns a tuple with the MediaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVoicemailMessageOutput) GetMediaIdOk() (*string, bool) {
	if o == nil || IsNil(o.MediaId) {
		return nil, false
	}
	return o.MediaId, true
}

// HasMediaId returns a boolean if a field has been set.
func (o *ServiceVoicemailMessageOutput) HasMediaId() bool {
	if o != nil && !IsNil(o.MediaId) {
		return true
	}

	return false
}

// SetMediaId gets a reference to the given string and assigns it to the MediaId field.
func (o *ServiceVoicemailMessageOutput) SetMediaId(v string) {
	o.MediaId = &v
}

// GetSucceeded returns the Succeeded field value if set, zero value otherwise.
func (o *ServiceVoicemailMessageOutput) GetSucceeded() []string {
	if o == nil || IsNil(o.Succeeded) {
		var ret []string
		return ret
	}
	return o.Succeeded
}

// GetSucceededOk returns a tuple with the Succeeded field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVoicemailMessageOutput) GetSucceededOk() ([]string, bool) {
	if o == nil || IsNil(o.Succeeded) {
		return nil, false
	}
	return o.Succeeded, true
}

// HasSucceeded returns a boolean if a field has been set.
func (o *ServiceVoicemailMessageOutput) HasSucceeded() bool {
	if o != nil && !IsNil(o.Succeeded) {
		return true
	}

	return false
}

// SetSucceeded gets a reference to the given []string and assigns it to the Succeeded field.
func (o *ServiceVoicemailMessageOutput) SetSucceeded(v []string) {
	o.Succeeded = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *ServiceVoicemailMessageOutput) GetTimestamp() int32 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVoicemailMessageOutput) GetTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *ServiceVoicemailMessageOutput) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int32 and assigns it to the Timestamp field.
func (o *ServiceVoicemailMessageOutput) SetTimestamp(v int32) {
	o.Timestamp = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *ServiceVoicemailMessageOutput) GetTo() string {
	if o == nil || IsNil(o.To) {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVoicemailMessageOutput) GetToOk() (*string, bool) {
	if o == nil || IsNil(o.To) {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *ServiceVoicemailMessageOutput) HasTo() bool {
	if o != nil && !IsNil(o.To) {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *ServiceVoicemailMessageOutput) SetTo(v string) {
	o.To = &v
}

func (o ServiceVoicemailMessageOutput) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceVoicemailMessageOutput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallId) {
		toSerialize["call_id"] = o.CallId
	}
	if !IsNil(o.CallerIdName) {
		toSerialize["caller_id_name"] = o.CallerIdName
	}
	if !IsNil(o.CallerIdNumber) {
		toSerialize["caller_id_number"] = o.CallerIdNumber
	}
	if !IsNil(o.Folder) {
		toSerialize["folder"] = o.Folder
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.Length) {
		toSerialize["length"] = o.Length
	}
	if !IsNil(o.MediaId) {
		toSerialize["media_id"] = o.MediaId
	}
	if !IsNil(o.Succeeded) {
		toSerialize["succeeded"] = o.Succeeded
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.To) {
		toSerialize["to"] = o.To
	}
	return toSerialize, nil
}

type NullableServiceVoicemailMessageOutput struct {
	value *ServiceVoicemailMessageOutput
	isSet bool
}

func (v NullableServiceVoicemailMessageOutput) Get() *ServiceVoicemailMessageOutput {
	return v.value
}

func (v *NullableServiceVoicemailMessageOutput) Set(val *ServiceVoicemailMessageOutput) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceVoicemailMessageOutput) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceVoicemailMessageOutput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceVoicemailMessageOutput(val *ServiceVoicemailMessageOutput) *NullableServiceVoicemailMessageOutput {
	return &NullableServiceVoicemailMessageOutput{value: val, isSet: true}
}

func (v NullableServiceVoicemailMessageOutput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceVoicemailMessageOutput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


