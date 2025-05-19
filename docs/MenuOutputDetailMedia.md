# MenuOutputDetailMedia

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Greeting** | Pointer to **string** | The ID of a media object that should be used as the menu greeting | [optional] 
**InvalidMedia** | Pointer to **map[string]interface{}** | When the collected digits don&#39;t result in a match or hunt this media can be played | [optional] 
**TransferMedia** | Pointer to **map[string]interface{}** | When a call is transferred from the menu, either after all retries exhausted or a successful hunt, this media can be played | [optional] 

## Methods

### NewMenuOutputDetailMedia

`func NewMenuOutputDetailMedia() *MenuOutputDetailMedia`

NewMenuOutputDetailMedia instantiates a new MenuOutputDetailMedia object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMenuOutputDetailMediaWithDefaults

`func NewMenuOutputDetailMediaWithDefaults() *MenuOutputDetailMedia`

NewMenuOutputDetailMediaWithDefaults instantiates a new MenuOutputDetailMedia object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGreeting

`func (o *MenuOutputDetailMedia) GetGreeting() string`

GetGreeting returns the Greeting field if non-nil, zero value otherwise.

### GetGreetingOk

`func (o *MenuOutputDetailMedia) GetGreetingOk() (*string, bool)`

GetGreetingOk returns a tuple with the Greeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreeting

`func (o *MenuOutputDetailMedia) SetGreeting(v string)`

SetGreeting sets Greeting field to given value.

### HasGreeting

`func (o *MenuOutputDetailMedia) HasGreeting() bool`

HasGreeting returns a boolean if a field has been set.

### GetInvalidMedia

`func (o *MenuOutputDetailMedia) GetInvalidMedia() map[string]interface{}`

GetInvalidMedia returns the InvalidMedia field if non-nil, zero value otherwise.

### GetInvalidMediaOk

`func (o *MenuOutputDetailMedia) GetInvalidMediaOk() (*map[string]interface{}, bool)`

GetInvalidMediaOk returns a tuple with the InvalidMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidMedia

`func (o *MenuOutputDetailMedia) SetInvalidMedia(v map[string]interface{})`

SetInvalidMedia sets InvalidMedia field to given value.

### HasInvalidMedia

`func (o *MenuOutputDetailMedia) HasInvalidMedia() bool`

HasInvalidMedia returns a boolean if a field has been set.

### GetTransferMedia

`func (o *MenuOutputDetailMedia) GetTransferMedia() map[string]interface{}`

GetTransferMedia returns the TransferMedia field if non-nil, zero value otherwise.

### GetTransferMediaOk

`func (o *MenuOutputDetailMedia) GetTransferMediaOk() (*map[string]interface{}, bool)`

GetTransferMediaOk returns a tuple with the TransferMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferMedia

`func (o *MenuOutputDetailMedia) SetTransferMedia(v map[string]interface{})`

SetTransferMedia sets TransferMedia field to given value.

### HasTransferMedia

`func (o *MenuOutputDetailMedia) HasTransferMedia() bool`

HasTransferMedia returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


