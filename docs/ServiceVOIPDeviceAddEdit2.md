

# ServiceVOIPDeviceAddEdit2


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**CallForward** | [**ServiceCallForward**](ServiceCallForward.md) |  |  [optional] |
|**CallRecording** | [**ServiceCallRecordingSettings**](ServiceCallRecordingSettings.md) |  |  [optional] |
|**CallerId** | [**ServiceVOIPDeviceAddEdit3c**](ServiceVOIPDeviceAddEdit3c.md) |  |  [optional] |
|**DeviceType** | [**DEVICE_TYPE**](#DEVICE_TYPE) |  |  [optional] |
|**DoNotDisturb** | [**ServiceVOIPSharedDoNotDisturb**](ServiceVOIPSharedDoNotDisturb.md) |  |  [optional] |
|**Enabled** | **bool** | cannot use required, else it has to be true and false is not allowed |  [optional] |
|**MacAddress** | **string** | dont use mac, it enforces :, which voip does not like |  [optional] |
|**MusicOnHold** | [**ServiceMusicOnHold**](ServiceMusicOnHold.md) |  |  [optional] |
|**Name** | **string** |  |  |
|**OwnerId** | **string** | json omitempty is needed else voip fails on \&quot;\&quot; for owner_id |  [optional] |
|**Sip** | [**ServiceVOIPDeviceAddEdit3a**](ServiceVOIPDeviceAddEdit3a.md) |  |  |



## Enum: DEVICE_TYPE

| Name | Value |
|---- | -----|
| SOFTPHONE | &quot;softphone&quot; |
| SIP_URI | &quot;sip_uri&quot; |



