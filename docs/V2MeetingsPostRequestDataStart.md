# V2MeetingsPostRequestDataStart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **time.Time** | An ISO 8601 datetime indicating when a non-all day meeting starts. | 
**Timezone** | Pointer to **string** | The IANA timezone the meeting starts in. If a datetime value is provided without an offset, this timezone will be used to convert the datetime to UTC using the timezone offset. If a datetime value is provided with an offset, this timezone will not be used to apply any additional offset to the datetime. Invalid timezones will be treated as UTC. | [optional] 
**Date** | **string** | An ISO 8601 date indicating when an all day meeting starts. | 

## Methods

### NewV2MeetingsPostRequestDataStart

`func NewV2MeetingsPostRequestDataStart(datetime time.Time, date string, ) *V2MeetingsPostRequestDataStart`

NewV2MeetingsPostRequestDataStart instantiates a new V2MeetingsPostRequestDataStart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MeetingsPostRequestDataStartWithDefaults

`func NewV2MeetingsPostRequestDataStartWithDefaults() *V2MeetingsPostRequestDataStart`

NewV2MeetingsPostRequestDataStartWithDefaults instantiates a new V2MeetingsPostRequestDataStart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *V2MeetingsPostRequestDataStart) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *V2MeetingsPostRequestDataStart) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *V2MeetingsPostRequestDataStart) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.


### GetTimezone

`func (o *V2MeetingsPostRequestDataStart) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *V2MeetingsPostRequestDataStart) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *V2MeetingsPostRequestDataStart) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *V2MeetingsPostRequestDataStart) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetDate

`func (o *V2MeetingsPostRequestDataStart) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *V2MeetingsPostRequestDataStart) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *V2MeetingsPostRequestDataStart) SetDate(v string)`

SetDate sets Date field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


