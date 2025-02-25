# ServiceCallRecordingParameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional] 
**Format** | Pointer to **string** |  | [optional] 
**RecordMinSec** | Pointer to **int32** |  | [optional] 
**RecordOnAnswer** | Pointer to **bool** |  | [optional] 
**RecordOnBridge** | Pointer to **bool** |  | [optional] 
**RecordSampleRate** | Pointer to **int32** |  | [optional] 
**TimeLimit** | Pointer to **int32** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 

## Methods

### NewServiceCallRecordingParameters

`func NewServiceCallRecordingParameters() *ServiceCallRecordingParameters`

NewServiceCallRecordingParameters instantiates a new ServiceCallRecordingParameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceCallRecordingParametersWithDefaults

`func NewServiceCallRecordingParametersWithDefaults() *ServiceCallRecordingParameters`

NewServiceCallRecordingParametersWithDefaults instantiates a new ServiceCallRecordingParameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ServiceCallRecordingParameters) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ServiceCallRecordingParameters) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ServiceCallRecordingParameters) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ServiceCallRecordingParameters) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFormat

`func (o *ServiceCallRecordingParameters) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *ServiceCallRecordingParameters) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *ServiceCallRecordingParameters) SetFormat(v string)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *ServiceCallRecordingParameters) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetRecordMinSec

`func (o *ServiceCallRecordingParameters) GetRecordMinSec() int32`

GetRecordMinSec returns the RecordMinSec field if non-nil, zero value otherwise.

### GetRecordMinSecOk

`func (o *ServiceCallRecordingParameters) GetRecordMinSecOk() (*int32, bool)`

GetRecordMinSecOk returns a tuple with the RecordMinSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordMinSec

`func (o *ServiceCallRecordingParameters) SetRecordMinSec(v int32)`

SetRecordMinSec sets RecordMinSec field to given value.

### HasRecordMinSec

`func (o *ServiceCallRecordingParameters) HasRecordMinSec() bool`

HasRecordMinSec returns a boolean if a field has been set.

### GetRecordOnAnswer

`func (o *ServiceCallRecordingParameters) GetRecordOnAnswer() bool`

GetRecordOnAnswer returns the RecordOnAnswer field if non-nil, zero value otherwise.

### GetRecordOnAnswerOk

`func (o *ServiceCallRecordingParameters) GetRecordOnAnswerOk() (*bool, bool)`

GetRecordOnAnswerOk returns a tuple with the RecordOnAnswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordOnAnswer

`func (o *ServiceCallRecordingParameters) SetRecordOnAnswer(v bool)`

SetRecordOnAnswer sets RecordOnAnswer field to given value.

### HasRecordOnAnswer

`func (o *ServiceCallRecordingParameters) HasRecordOnAnswer() bool`

HasRecordOnAnswer returns a boolean if a field has been set.

### GetRecordOnBridge

`func (o *ServiceCallRecordingParameters) GetRecordOnBridge() bool`

GetRecordOnBridge returns the RecordOnBridge field if non-nil, zero value otherwise.

### GetRecordOnBridgeOk

`func (o *ServiceCallRecordingParameters) GetRecordOnBridgeOk() (*bool, bool)`

GetRecordOnBridgeOk returns a tuple with the RecordOnBridge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordOnBridge

`func (o *ServiceCallRecordingParameters) SetRecordOnBridge(v bool)`

SetRecordOnBridge sets RecordOnBridge field to given value.

### HasRecordOnBridge

`func (o *ServiceCallRecordingParameters) HasRecordOnBridge() bool`

HasRecordOnBridge returns a boolean if a field has been set.

### GetRecordSampleRate

`func (o *ServiceCallRecordingParameters) GetRecordSampleRate() int32`

GetRecordSampleRate returns the RecordSampleRate field if non-nil, zero value otherwise.

### GetRecordSampleRateOk

`func (o *ServiceCallRecordingParameters) GetRecordSampleRateOk() (*int32, bool)`

GetRecordSampleRateOk returns a tuple with the RecordSampleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordSampleRate

`func (o *ServiceCallRecordingParameters) SetRecordSampleRate(v int32)`

SetRecordSampleRate sets RecordSampleRate field to given value.

### HasRecordSampleRate

`func (o *ServiceCallRecordingParameters) HasRecordSampleRate() bool`

HasRecordSampleRate returns a boolean if a field has been set.

### GetTimeLimit

`func (o *ServiceCallRecordingParameters) GetTimeLimit() int32`

GetTimeLimit returns the TimeLimit field if non-nil, zero value otherwise.

### GetTimeLimitOk

`func (o *ServiceCallRecordingParameters) GetTimeLimitOk() (*int32, bool)`

GetTimeLimitOk returns a tuple with the TimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimit

`func (o *ServiceCallRecordingParameters) SetTimeLimit(v int32)`

SetTimeLimit sets TimeLimit field to given value.

### HasTimeLimit

`func (o *ServiceCallRecordingParameters) HasTimeLimit() bool`

HasTimeLimit returns a boolean if a field has been set.

### GetUrl

`func (o *ServiceCallRecordingParameters) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ServiceCallRecordingParameters) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ServiceCallRecordingParameters) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ServiceCallRecordingParameters) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


