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

// checks if the ServiceVoicemailOutputFull type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceVoicemailOutputFull{}

// ServiceVoicemailOutputFull struct for ServiceVoicemailOutputFull
type ServiceVoicemailOutputFull struct {
	Id *string `json:"id,omitempty"`
	Mailbox *string `json:"mailbox,omitempty"`
	Media *ServiceVoicemailMedia `json:"media,omitempty"`
	MediaExtension *string `json:"media_extension,omitempty"`
	Messages *int32 `json:"messages,omitempty"`
	Name *string `json:"name,omitempty"`
	OwnerId *string `json:"owner_id,omitempty"`
	Pin *string `json:"pin,omitempty"`
	RequirePin *bool `json:"require_pin,omitempty"`
	Timezone *string `json:"timezone,omitempty"`
}

// NewServiceVoicemailOutputFull instantiates a new ServiceVoicemailOutputFull object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceVoicemailOutputFull() *ServiceVoicemailOutputFull {
	this := ServiceVoicemailOutputFull{}
	return &this
}

// NewServiceVoicemailOutputFullWithDefaults instantiates a new ServiceVoicemailOutputFull object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceVoicemailOutputFullWithDefaults() *ServiceVoicemailOutputFull {
	this := ServiceVoicemailOutputFull{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceVoicemailOutputFull) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVoicemailOutputFull) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceVoicemailOutputFull) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceVoicemailOutputFull) SetId(v string) {
	o.Id = &v
}

// GetMailbox returns the Mailbox field value if set, zero value otherwise.
func (o *ServiceVoicemailOutputFull) GetMailbox() string {
	if o == nil || IsNil(o.Mailbox) {
		var ret string
		return ret
	}
	return *o.Mailbox
}

// GetMailboxOk returns a tuple with the Mailbox field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVoicemailOutputFull) GetMailboxOk() (*string, bool) {
	if o == nil || IsNil(o.Mailbox) {
		return nil, false
	}
	return o.Mailbox, true
}

// HasMailbox returns a boolean if a field has been set.
func (o *ServiceVoicemailOutputFull) HasMailbox() bool {
	if o != nil && !IsNil(o.Mailbox) {
		return true
	}

	return false
}

// SetMailbox gets a reference to the given string and assigns it to the Mailbox field.
func (o *ServiceVoicemailOutputFull) SetMailbox(v string) {
	o.Mailbox = &v
}

// GetMedia returns the Media field value if set, zero value otherwise.
func (o *ServiceVoicemailOutputFull) GetMedia() ServiceVoicemailMedia {
	if o == nil || IsNil(o.Media) {
		var ret ServiceVoicemailMedia
		return ret
	}
	return *o.Media
}

// GetMediaOk returns a tuple with the Media field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVoicemailOutputFull) GetMediaOk() (*ServiceVoicemailMedia, bool) {
	if o == nil || IsNil(o.Media) {
		return nil, false
	}
	return o.Media, true
}

// HasMedia returns a boolean if a field has been set.
func (o *ServiceVoicemailOutputFull) HasMedia() bool {
	if o != nil && !IsNil(o.Media) {
		return true
	}

	return false
}

// SetMedia gets a reference to the given ServiceVoicemailMedia and assigns it to the Media field.
func (o *ServiceVoicemailOutputFull) SetMedia(v ServiceVoicemailMedia) {
	o.Media = &v
}

// GetMediaExtension returns the MediaExtension field value if set, zero value otherwise.
func (o *ServiceVoicemailOutputFull) GetMediaExtension() string {
	if o == nil || IsNil(o.MediaExtension) {
		var ret string
		return ret
	}
	return *o.MediaExtension
}

// GetMediaExtensionOk returns a tuple with the MediaExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVoicemailOutputFull) GetMediaExtensionOk() (*string, bool) {
	if o == nil || IsNil(o.MediaExtension) {
		return nil, false
	}
	return o.MediaExtension, true
}

// HasMediaExtension returns a boolean if a field has been set.
func (o *ServiceVoicemailOutputFull) HasMediaExtension() bool {
	if o != nil && !IsNil(o.MediaExtension) {
		return true
	}

	return false
}

// SetMediaExtension gets a reference to the given string and assigns it to the MediaExtension field.
func (o *ServiceVoicemailOutputFull) SetMediaExtension(v string) {
	o.MediaExtension = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *ServiceVoicemailOutputFull) GetMessages() int32 {
	if o == nil || IsNil(o.Messages) {
		var ret int32
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVoicemailOutputFull) GetMessagesOk() (*int32, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *ServiceVoicemailOutputFull) HasMessages() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given int32 and assigns it to the Messages field.
func (o *ServiceVoicemailOutputFull) SetMessages(v int32) {
	o.Messages = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ServiceVoicemailOutputFull) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVoicemailOutputFull) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ServiceVoicemailOutputFull) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ServiceVoicemailOutputFull) SetName(v string) {
	o.Name = &v
}

// GetOwnerId returns the OwnerId field value if set, zero value otherwise.
func (o *ServiceVoicemailOutputFull) GetOwnerId() string {
	if o == nil || IsNil(o.OwnerId) {
		var ret string
		return ret
	}
	return *o.OwnerId
}

// GetOwnerIdOk returns a tuple with the OwnerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVoicemailOutputFull) GetOwnerIdOk() (*string, bool) {
	if o == nil || IsNil(o.OwnerId) {
		return nil, false
	}
	return o.OwnerId, true
}

// HasOwnerId returns a boolean if a field has been set.
func (o *ServiceVoicemailOutputFull) HasOwnerId() bool {
	if o != nil && !IsNil(o.OwnerId) {
		return true
	}

	return false
}

// SetOwnerId gets a reference to the given string and assigns it to the OwnerId field.
func (o *ServiceVoicemailOutputFull) SetOwnerId(v string) {
	o.OwnerId = &v
}

// GetPin returns the Pin field value if set, zero value otherwise.
func (o *ServiceVoicemailOutputFull) GetPin() string {
	if o == nil || IsNil(o.Pin) {
		var ret string
		return ret
	}
	return *o.Pin
}

// GetPinOk returns a tuple with the Pin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVoicemailOutputFull) GetPinOk() (*string, bool) {
	if o == nil || IsNil(o.Pin) {
		return nil, false
	}
	return o.Pin, true
}

// HasPin returns a boolean if a field has been set.
func (o *ServiceVoicemailOutputFull) HasPin() bool {
	if o != nil && !IsNil(o.Pin) {
		return true
	}

	return false
}

// SetPin gets a reference to the given string and assigns it to the Pin field.
func (o *ServiceVoicemailOutputFull) SetPin(v string) {
	o.Pin = &v
}

// GetRequirePin returns the RequirePin field value if set, zero value otherwise.
func (o *ServiceVoicemailOutputFull) GetRequirePin() bool {
	if o == nil || IsNil(o.RequirePin) {
		var ret bool
		return ret
	}
	return *o.RequirePin
}

// GetRequirePinOk returns a tuple with the RequirePin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVoicemailOutputFull) GetRequirePinOk() (*bool, bool) {
	if o == nil || IsNil(o.RequirePin) {
		return nil, false
	}
	return o.RequirePin, true
}

// HasRequirePin returns a boolean if a field has been set.
func (o *ServiceVoicemailOutputFull) HasRequirePin() bool {
	if o != nil && !IsNil(o.RequirePin) {
		return true
	}

	return false
}

// SetRequirePin gets a reference to the given bool and assigns it to the RequirePin field.
func (o *ServiceVoicemailOutputFull) SetRequirePin(v bool) {
	o.RequirePin = &v
}

// GetTimezone returns the Timezone field value if set, zero value otherwise.
func (o *ServiceVoicemailOutputFull) GetTimezone() string {
	if o == nil || IsNil(o.Timezone) {
		var ret string
		return ret
	}
	return *o.Timezone
}

// GetTimezoneOk returns a tuple with the Timezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceVoicemailOutputFull) GetTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.Timezone) {
		return nil, false
	}
	return o.Timezone, true
}

// HasTimezone returns a boolean if a field has been set.
func (o *ServiceVoicemailOutputFull) HasTimezone() bool {
	if o != nil && !IsNil(o.Timezone) {
		return true
	}

	return false
}

// SetTimezone gets a reference to the given string and assigns it to the Timezone field.
func (o *ServiceVoicemailOutputFull) SetTimezone(v string) {
	o.Timezone = &v
}

func (o ServiceVoicemailOutputFull) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceVoicemailOutputFull) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Mailbox) {
		toSerialize["mailbox"] = o.Mailbox
	}
	if !IsNil(o.Media) {
		toSerialize["media"] = o.Media
	}
	if !IsNil(o.MediaExtension) {
		toSerialize["media_extension"] = o.MediaExtension
	}
	if !IsNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OwnerId) {
		toSerialize["owner_id"] = o.OwnerId
	}
	if !IsNil(o.Pin) {
		toSerialize["pin"] = o.Pin
	}
	if !IsNil(o.RequirePin) {
		toSerialize["require_pin"] = o.RequirePin
	}
	if !IsNil(o.Timezone) {
		toSerialize["timezone"] = o.Timezone
	}
	return toSerialize, nil
}

type NullableServiceVoicemailOutputFull struct {
	value *ServiceVoicemailOutputFull
	isSet bool
}

func (v NullableServiceVoicemailOutputFull) Get() *ServiceVoicemailOutputFull {
	return v.value
}

func (v *NullableServiceVoicemailOutputFull) Set(val *ServiceVoicemailOutputFull) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceVoicemailOutputFull) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceVoicemailOutputFull) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceVoicemailOutputFull(val *ServiceVoicemailOutputFull) *NullableServiceVoicemailOutputFull {
	return &NullableServiceVoicemailOutputFull{value: val, isSet: true}
}

func (v NullableServiceVoicemailOutputFull) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceVoicemailOutputFull) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


