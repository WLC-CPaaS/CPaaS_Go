# MenuOutputDetailData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the menu | [optional] 
**Media** | Pointer to [**MenuOutputDetailMedia**](MenuOutputDetailMedia.md) | The media (prompt) parameters | [optional] 
**Name** | Pointer to **string** | A friendly name for the menu | [optional] 
**Timeout** | Pointer to **int32** | The amount of time (in milliseconds) to wait for the caller to begin entering digits | [optional] 

## Methods

### NewMenuOutputDetailData

`func NewMenuOutputDetailData() *MenuOutputDetailData`

NewMenuOutputDetailData instantiates a new MenuOutputDetailData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuOutputDetailDataWithDefaults

`func NewMenuOutputDetailDataWithDefaults() *MenuOutputDetailData`

NewMenuOutputDetailDataWithDefaults instantiates a new MenuOutputDetailData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MenuOutputDetailData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MenuOutputDetailData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MenuOutputDetailData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MenuOutputDetailData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMedia

`func (o *MenuOutputDetailData) GetMedia() MenuOutputDetailMedia`

GetMedia returns the Media field if non-nil, zero value otherwise.

### GetMediaOk

`func (o *MenuOutputDetailData) GetMediaOk() (*MenuOutputDetailMedia, bool)`

GetMediaOk returns a tuple with the Media field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedia

`func (o *MenuOutputDetailData) SetMedia(v MenuOutputDetailMedia)`

SetMedia sets Media field to given value.

### HasMedia

`func (o *MenuOutputDetailData) HasMedia() bool`

HasMedia returns a boolean if a field has been set.

### GetName

`func (o *MenuOutputDetailData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MenuOutputDetailData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MenuOutputDetailData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MenuOutputDetailData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTimeout

`func (o *MenuOutputDetailData) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *MenuOutputDetailData) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *MenuOutputDetailData) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *MenuOutputDetailData) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


