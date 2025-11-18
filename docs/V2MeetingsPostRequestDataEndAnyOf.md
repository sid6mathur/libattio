# V2MeetingsPostRequestDataEndAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datetime** | **time.Time** | An ISO 8601 datetime indicating when a non-all day meeting ends. Note that this value is exclusive, meaning that the meeting ends before the specified time, not at it. For example, a one hour meeting starting at 14:00 would end at 15:00, not 15:59:59. | 
**Timezone** | Pointer to **NullableString** | The IANA timezone the meeting ends in. If a datetime value is provided without an offset, this timezone will be used to convert the datetime to UTC using the timezone offset. If a datetime value is provided with an offset, this timezone will not be used to apply any additional offset to the datetime. Invalid timezones will be treated as UTC. | [optional] 

## Methods

### NewV2MeetingsPostRequestDataEndAnyOf

`func NewV2MeetingsPostRequestDataEndAnyOf(datetime time.Time, ) *V2MeetingsPostRequestDataEndAnyOf`

NewV2MeetingsPostRequestDataEndAnyOf instantiates a new V2MeetingsPostRequestDataEndAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2MeetingsPostRequestDataEndAnyOfWithDefaults

`func NewV2MeetingsPostRequestDataEndAnyOfWithDefaults() *V2MeetingsPostRequestDataEndAnyOf`

NewV2MeetingsPostRequestDataEndAnyOfWithDefaults instantiates a new V2MeetingsPostRequestDataEndAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatetime

`func (o *V2MeetingsPostRequestDataEndAnyOf) GetDatetime() time.Time`

GetDatetime returns the Datetime field if non-nil, zero value otherwise.

### GetDatetimeOk

`func (o *V2MeetingsPostRequestDataEndAnyOf) GetDatetimeOk() (*time.Time, bool)`

GetDatetimeOk returns a tuple with the Datetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatetime

`func (o *V2MeetingsPostRequestDataEndAnyOf) SetDatetime(v time.Time)`

SetDatetime sets Datetime field to given value.


### GetTimezone

`func (o *V2MeetingsPostRequestDataEndAnyOf) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *V2MeetingsPostRequestDataEndAnyOf) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *V2MeetingsPostRequestDataEndAnyOf) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *V2MeetingsPostRequestDataEndAnyOf) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *V2MeetingsPostRequestDataEndAnyOf) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *V2MeetingsPostRequestDataEndAnyOf) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


