/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.39.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
	"time"
)

// TransferEvent Represents an event in the Transfers API.
type TransferEvent struct {
	// Plaid’s unique identifier for this event. IDs are sequential unsigned 64-bit integers.
	EventId int32 `json:"event_id"`
	// The datetime when this event occurred. This will be of the form `2006-01-02T15:04:05Z`.
	Timestamp time.Time `json:"timestamp"`
	EventType TransferEventType `json:"event_type"`
	// The account ID associated with the transfer.
	AccountId string `json:"account_id"`
	// Plaid’s unique identifier for a transfer.
	TransferId string `json:"transfer_id"`
	// The ID of the origination account that this balance belongs to.
	OriginationAccountId NullableString `json:"origination_account_id"`
	TransferType TransferType `json:"transfer_type"`
	// The amount of the transfer (decimal string with two digits of precision e.g. “10.00”).
	TransferAmount string `json:"transfer_amount"`
	FailureReason NullableTransferFailure `json:"failure_reason"`
	AdditionalProperties map[string]interface{}
}

type _TransferEvent TransferEvent

// NewTransferEvent instantiates a new TransferEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferEvent(eventId int32, timestamp time.Time, eventType TransferEventType, accountId string, transferId string, originationAccountId NullableString, transferType TransferType, transferAmount string, failureReason NullableTransferFailure) *TransferEvent {
	this := TransferEvent{}
	this.EventId = eventId
	this.Timestamp = timestamp
	this.EventType = eventType
	this.AccountId = accountId
	this.TransferId = transferId
	this.OriginationAccountId = originationAccountId
	this.TransferType = transferType
	this.TransferAmount = transferAmount
	this.FailureReason = failureReason
	return &this
}

// NewTransferEventWithDefaults instantiates a new TransferEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferEventWithDefaults() *TransferEvent {
	this := TransferEvent{}
	return &this
}

// GetEventId returns the EventId field value
func (o *TransferEvent) GetEventId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetEventIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *TransferEvent) SetEventId(v int32) {
	o.EventId = v
}

// GetTimestamp returns the Timestamp field value
func (o *TransferEvent) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetTimestampOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *TransferEvent) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetEventType returns the EventType field value
func (o *TransferEvent) GetEventType() TransferEventType {
	if o == nil {
		var ret TransferEventType
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetEventTypeOk() (*TransferEventType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *TransferEvent) SetEventType(v TransferEventType) {
	o.EventType = v
}

// GetAccountId returns the AccountId field value
func (o *TransferEvent) GetAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AccountId, true
}

// SetAccountId sets field value
func (o *TransferEvent) SetAccountId(v string) {
	o.AccountId = v
}

// GetTransferId returns the TransferId field value
func (o *TransferEvent) GetTransferId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransferId
}

// GetTransferIdOk returns a tuple with the TransferId field value
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetTransferIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferId, true
}

// SetTransferId sets field value
func (o *TransferEvent) SetTransferId(v string) {
	o.TransferId = v
}

// GetOriginationAccountId returns the OriginationAccountId field value
// If the value is explicit nil, the zero value for string will be returned
func (o *TransferEvent) GetOriginationAccountId() string {
	if o == nil || o.OriginationAccountId.Get() == nil {
		var ret string
		return ret
	}

	return *o.OriginationAccountId.Get()
}

// GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferEvent) GetOriginationAccountIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.OriginationAccountId.Get(), o.OriginationAccountId.IsSet()
}

// SetOriginationAccountId sets field value
func (o *TransferEvent) SetOriginationAccountId(v string) {
	o.OriginationAccountId.Set(&v)
}

// GetTransferType returns the TransferType field value
func (o *TransferEvent) GetTransferType() TransferType {
	if o == nil {
		var ret TransferType
		return ret
	}

	return o.TransferType
}

// GetTransferTypeOk returns a tuple with the TransferType field value
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetTransferTypeOk() (*TransferType, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferType, true
}

// SetTransferType sets field value
func (o *TransferEvent) SetTransferType(v TransferType) {
	o.TransferType = v
}

// GetTransferAmount returns the TransferAmount field value
func (o *TransferEvent) GetTransferAmount() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TransferAmount
}

// GetTransferAmountOk returns a tuple with the TransferAmount field value
// and a boolean to check if the value has been set.
func (o *TransferEvent) GetTransferAmountOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferAmount, true
}

// SetTransferAmount sets field value
func (o *TransferEvent) SetTransferAmount(v string) {
	o.TransferAmount = v
}

// GetFailureReason returns the FailureReason field value
// If the value is explicit nil, the zero value for TransferFailure will be returned
func (o *TransferEvent) GetFailureReason() TransferFailure {
	if o == nil || o.FailureReason.Get() == nil {
		var ret TransferFailure
		return ret
	}

	return *o.FailureReason.Get()
}

// GetFailureReasonOk returns a tuple with the FailureReason field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransferEvent) GetFailureReasonOk() (*TransferFailure, bool) {
	if o == nil  {
		return nil, false
	}
	return o.FailureReason.Get(), o.FailureReason.IsSet()
}

// SetFailureReason sets field value
func (o *TransferEvent) SetFailureReason(v TransferFailure) {
	o.FailureReason.Set(&v)
}

func (o TransferEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["event_id"] = o.EventId
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if true {
		toSerialize["event_type"] = o.EventType
	}
	if true {
		toSerialize["account_id"] = o.AccountId
	}
	if true {
		toSerialize["transfer_id"] = o.TransferId
	}
	if true {
		toSerialize["origination_account_id"] = o.OriginationAccountId.Get()
	}
	if true {
		toSerialize["transfer_type"] = o.TransferType
	}
	if true {
		toSerialize["transfer_amount"] = o.TransferAmount
	}
	if true {
		toSerialize["failure_reason"] = o.FailureReason.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferEvent) UnmarshalJSON(bytes []byte) (err error) {
	varTransferEvent := _TransferEvent{}

	if err = json.Unmarshal(bytes, &varTransferEvent); err == nil {
		*o = TransferEvent(varTransferEvent)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "event_id")
		delete(additionalProperties, "timestamp")
		delete(additionalProperties, "event_type")
		delete(additionalProperties, "account_id")
		delete(additionalProperties, "transfer_id")
		delete(additionalProperties, "origination_account_id")
		delete(additionalProperties, "transfer_type")
		delete(additionalProperties, "transfer_amount")
		delete(additionalProperties, "failure_reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferEvent struct {
	value *TransferEvent
	isSet bool
}

func (v NullableTransferEvent) Get() *TransferEvent {
	return v.value
}

func (v *NullableTransferEvent) Set(val *TransferEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferEvent(val *TransferEvent) *NullableTransferEvent {
	return &NullableTransferEvent{value: val, isSet: true}
}

func (v NullableTransferEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


