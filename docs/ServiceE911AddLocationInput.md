# ServiceE911AddLocationInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**ServiceE911LocationInput**](ServiceE911LocationInput.md) |  | [optional] 
**Uri** | Pointer to [**ServiceE911URIInput**](ServiceE911URIInput.md) |  | [optional] 

## Methods

### NewServiceE911AddLocationInput

`func NewServiceE911AddLocationInput() *ServiceE911AddLocationInput`

NewServiceE911AddLocationInput instantiates a new ServiceE911AddLocationInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceE911AddLocationInputWithDefaults

`func NewServiceE911AddLocationInputWithDefaults() *ServiceE911AddLocationInput`

NewServiceE911AddLocationInputWithDefaults instantiates a new ServiceE911AddLocationInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *ServiceE911AddLocationInput) GetLocation() ServiceE911LocationInput`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *ServiceE911AddLocationInput) GetLocationOk() (*ServiceE911LocationInput, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *ServiceE911AddLocationInput) SetLocation(v ServiceE911LocationInput)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *ServiceE911AddLocationInput) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetUri

`func (o *ServiceE911AddLocationInput) GetUri() ServiceE911URIInput`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *ServiceE911AddLocationInput) GetUriOk() (*ServiceE911URIInput, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *ServiceE911AddLocationInput) SetUri(v ServiceE911URIInput)`

SetUri sets Uri field to given value.

### HasUri

`func (o *ServiceE911AddLocationInput) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


