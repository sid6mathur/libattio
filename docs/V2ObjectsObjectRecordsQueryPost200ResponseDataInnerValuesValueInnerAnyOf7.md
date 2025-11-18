# V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveFrom** | **time.Time** | The point in time at which this value was made \&quot;active\&quot;. &#x60;active_from&#x60; can be considered roughly analogous to &#x60;created_at&#x60;. | 
**ActiveUntil** | **NullableTime** | The point in time at which this value was deactivated. If &#x60;null&#x60;, the value is active. | 
**CreatedByActor** | [**V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOfCreatedByActor**](V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOfCreatedByActor.md) |  | 
**InteractionType** | **string** | The type of interaction e.g. calendar or email. | 
**InteractedAt** | **time.Time** | When the interaction occurred. | 
**OwnerActor** | [**V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOfCreatedByActor**](V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOfCreatedByActor.md) |  | 
**AttributeType** | **string** | The attribute type of the value. | 

## Methods

### NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7

`func NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7(activeFrom time.Time, activeUntil NullableTime, createdByActor V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOfCreatedByActor, interactionType string, interactedAt time.Time, ownerActor V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOfCreatedByActor, attributeType string, ) *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7`

NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7 instantiates a new V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7WithDefaults

`func NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7WithDefaults() *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7`

NewV2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7WithDefaults instantiates a new V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveFrom

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7) GetActiveFrom() time.Time`

GetActiveFrom returns the ActiveFrom field if non-nil, zero value otherwise.

### GetActiveFromOk

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7) GetActiveFromOk() (*time.Time, bool)`

GetActiveFromOk returns a tuple with the ActiveFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveFrom

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7) SetActiveFrom(v time.Time)`

SetActiveFrom sets ActiveFrom field to given value.


### GetActiveUntil

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7) GetActiveUntil() time.Time`

GetActiveUntil returns the ActiveUntil field if non-nil, zero value otherwise.

### GetActiveUntilOk

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7) GetActiveUntilOk() (*time.Time, bool)`

GetActiveUntilOk returns a tuple with the ActiveUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveUntil

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7) SetActiveUntil(v time.Time)`

SetActiveUntil sets ActiveUntil field to given value.


### SetActiveUntilNil

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7) SetActiveUntilNil(b bool)`

 SetActiveUntilNil sets the value for ActiveUntil to be an explicit nil

### UnsetActiveUntil
`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7) UnsetActiveUntil()`

UnsetActiveUntil ensures that no value is present for ActiveUntil, not even an explicit nil
### GetCreatedByActor

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7) GetCreatedByActor() V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOfCreatedByActor`

GetCreatedByActor returns the CreatedByActor field if non-nil, zero value otherwise.

### GetCreatedByActorOk

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7) GetCreatedByActorOk() (*V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOfCreatedByActor, bool)`

GetCreatedByActorOk returns a tuple with the CreatedByActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByActor

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7) SetCreatedByActor(v V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOfCreatedByActor)`

SetCreatedByActor sets CreatedByActor field to given value.


### GetInteractionType

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7) GetInteractionType() string`

GetInteractionType returns the InteractionType field if non-nil, zero value otherwise.

### GetInteractionTypeOk

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7) GetInteractionTypeOk() (*string, bool)`

GetInteractionTypeOk returns a tuple with the InteractionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractionType

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7) SetInteractionType(v string)`

SetInteractionType sets InteractionType field to given value.


### GetInteractedAt

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7) GetInteractedAt() time.Time`

GetInteractedAt returns the InteractedAt field if non-nil, zero value otherwise.

### GetInteractedAtOk

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7) GetInteractedAtOk() (*time.Time, bool)`

GetInteractedAtOk returns a tuple with the InteractedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInteractedAt

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7) SetInteractedAt(v time.Time)`

SetInteractedAt sets InteractedAt field to given value.


### GetOwnerActor

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7) GetOwnerActor() V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOfCreatedByActor`

GetOwnerActor returns the OwnerActor field if non-nil, zero value otherwise.

### GetOwnerActorOk

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7) GetOwnerActorOk() (*V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOfCreatedByActor, bool)`

GetOwnerActorOk returns a tuple with the OwnerActor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnerActor

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7) SetOwnerActor(v V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOfCreatedByActor)`

SetOwnerActor sets OwnerActor field to given value.


### GetAttributeType

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7) GetAttributeType() string`

GetAttributeType returns the AttributeType field if non-nil, zero value otherwise.

### GetAttributeTypeOk

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7) GetAttributeTypeOk() (*string, bool)`

GetAttributeTypeOk returns a tuple with the AttributeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributeType

`func (o *V2ObjectsObjectRecordsQueryPost200ResponseDataInnerValuesValueInnerAnyOf7) SetAttributeType(v string)`

SetAttributeType sets AttributeType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


