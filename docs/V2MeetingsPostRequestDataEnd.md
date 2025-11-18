# V2MeetingsPostRequestDataEnd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **time.Time** | An ISO 8601 datetime indicating when a non-all day meeting ends. Note that this value is exclusive, meaning that the meeting ends before the specified time, not at it. For example, a one hour meeting starting at 14:00 would end at 15:00, not 15:59:59. | 
**Timezone** | Pointer to **string** | The IANA timezone the meeting ends in. If a datetime value is provided without an offset, this timezone will be used to convert the datetime to UTC using the timezone offset. If a datetime value is provided with an offset, this timezone will not be used to apply any additional offset to the datetime. Invalid timezones will be treated as UTC. | [optional] 
**Date** | **string** | An ISO 8601 date indicating when an all day meeting ends. Note that dates are exclusive, meaning that the meeting ends before the specified time, not at it. For example, a one day meeting on June 3rd would end on June 4th, not June 3rd. | 

## Methods

### NewV2MeetingsPostRequestDataEnd

`func NewV2MeetingsPostRequestDataEnd(datetime time.Time, date string, ) *V2MeetingsPostRequestDataEnd`

NewV2MeetingsPostRequestDataEnd instantiates a new V2MeetingsPostRequestDataEnd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MeetingsPostRequestDataEndWithDefaults

`func NewV2MeetingsPostRequestDataEndWithDefaults() *V2MeetingsPostRequestDataEnd`

NewV2MeetingsPostRequestDataEndWithDefaults instantiates a new V2MeetingsPostRequestDataEnd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *V2MeetingsPostRequestDataEnd) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *V2MeetingsPostRequestDataEnd) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *V2MeetingsPostRequestDataEnd) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.


### GetTimezone

`func (o *V2MeetingsPostRequestDataEnd) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *V2MeetingsPostRequestDataEnd) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *V2MeetingsPostRequestDataEnd) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *V2MeetingsPostRequestDataEnd) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetDate

`func (o *V2MeetingsPostRequestDataEnd) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *V2MeetingsPostRequestDataEnd) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *V2MeetingsPostRequestDataEnd) SetDate(v string)`

SetDate sets Date field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


