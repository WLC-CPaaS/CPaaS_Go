# ServiceVoicemailMessageOutput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallId** | Pointer to **string** |  | [optional] 
**CallerIdName** | Pointer to **string** |  | [optional] 
**CallerIdNumber** | Pointer to **string** |  | [optional] 
**Folder** | Pointer to **string** |  | [optional] 
**From** | Pointer to **string** |  | [optional] 
**Length** | Pointer to **int32** |  | [optional] 
**MediaId** | Pointer to **string** |  | [optional] 
**Succeeded** | Pointer to **[]string** |  | [optional] 
**Timestamp** | Pointer to **int32** |  | [optional] 
**To** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceVoicemailMessageOutput

`func NewServiceVoicemailMessageOutput() *ServiceVoicemailMessageOutput`

NewServiceVoicemailMessageOutput instantiates a new ServiceVoicemailMessageOutput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceVoicemailMessageOutputWithDefaults

`func NewServiceVoicemailMessageOutputWithDefaults() *ServiceVoicemailMessageOutput`

NewServiceVoicemailMessageOutputWithDefaults instantiates a new ServiceVoicemailMessageOutput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallId

`func (o *ServiceVoicemailMessageOutput) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *ServiceVoicemailMessageOutput) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *ServiceVoicemailMessageOutput) SetCallId(v string)`

SetCallId sets CallId field to given value.

### HasCallId

`func (o *ServiceVoicemailMessageOutput) HasCallId() bool`

HasCallId returns a boolean if a field has been set.

### GetCallerIdName

`func (o *ServiceVoicemailMessageOutput) GetCallerIdName() string`

GetCallerIdName returns the CallerIdName field if non-nil, zero value otherwise.

### GetCallerIdNameOk

`func (o *ServiceVoicemailMessageOutput) GetCallerIdNameOk() (*string, bool)`

GetCallerIdNameOk returns a tuple with the CallerIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdName

`func (o *ServiceVoicemailMessageOutput) SetCallerIdName(v string)`

SetCallerIdName sets CallerIdName field to given value.

### HasCallerIdName

`func (o *ServiceVoicemailMessageOutput) HasCallerIdName() bool`

HasCallerIdName returns a boolean if a field has been set.

### GetCallerIdNumber

`func (o *ServiceVoicemailMessageOutput) GetCallerIdNumber() string`

GetCallerIdNumber returns the CallerIdNumber field if non-nil, zero value otherwise.

### GetCallerIdNumberOk

`func (o *ServiceVoicemailMessageOutput) GetCallerIdNumberOk() (*string, bool)`

GetCallerIdNumberOk returns a tuple with the CallerIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdNumber

`func (o *ServiceVoicemailMessageOutput) SetCallerIdNumber(v string)`

SetCallerIdNumber sets CallerIdNumber field to given value.

### HasCallerIdNumber

`func (o *ServiceVoicemailMessageOutput) HasCallerIdNumber() bool`

HasCallerIdNumber returns a boolean if a field has been set.

### GetFolder

`func (o *ServiceVoicemailMessageOutput) GetFolder() string`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *ServiceVoicemailMessageOutput) GetFolderOk() (*string, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *ServiceVoicemailMessageOutput) SetFolder(v string)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *ServiceVoicemailMessageOutput) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

### GetFrom

`func (o *ServiceVoicemailMessageOutput) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ServiceVoicemailMessageOutput) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ServiceVoicemailMessageOutput) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ServiceVoicemailMessageOutput) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetLength

`func (o *ServiceVoicemailMessageOutput) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *ServiceVoicemailMessageOutput) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *ServiceVoicemailMessageOutput) SetLength(v int32)`

SetLength sets Length field to given value.

### HasLength

`func (o *ServiceVoicemailMessageOutput) HasLength() bool`

HasLength returns a boolean if a field has been set.

### GetMediaId

`func (o *ServiceVoicemailMessageOutput) GetMediaId() string`

GetMediaId returns the MediaId field if non-nil, zero value otherwise.

### GetMediaIdOk

`func (o *ServiceVoicemailMessageOutput) GetMediaIdOk() (*string, bool)`

GetMediaIdOk returns a tuple with the MediaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaId

`func (o *ServiceVoicemailMessageOutput) SetMediaId(v string)`

SetMediaId sets MediaId field to given value.

### HasMediaId

`func (o *ServiceVoicemailMessageOutput) HasMediaId() bool`

HasMediaId returns a boolean if a field has been set.

### GetSucceeded

`func (o *ServiceVoicemailMessageOutput) GetSucceeded() []string`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *ServiceVoicemailMessageOutput) GetSucceededOk() (*[]string, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *ServiceVoicemailMessageOutput) SetSucceeded(v []string)`

SetSucceeded sets Succeeded field to given value.

### HasSucceeded

`func (o *ServiceVoicemailMessageOutput) HasSucceeded() bool`

HasSucceeded returns a boolean if a field has been set.

### GetTimestamp

`func (o *ServiceVoicemailMessageOutput) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ServiceVoicemailMessageOutput) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ServiceVoicemailMessageOutput) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *ServiceVoicemailMessageOutput) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetTo

`func (o *ServiceVoicemailMessageOutput) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ServiceVoicemailMessageOutput) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ServiceVoicemailMessageOutput) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ServiceVoicemailMessageOutput) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


