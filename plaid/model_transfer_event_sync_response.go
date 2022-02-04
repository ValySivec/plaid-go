/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.64.9
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// TransferEventSyncResponse Defines the response schema for `/transfer/event/sync`
type TransferEventSyncResponse struct {
	TransferEvents []TransferEvent `json:"transfer_events"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferEventSyncResponse TransferEventSyncResponse

// NewTransferEventSyncResponse instantiates a new TransferEventSyncResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferEventSyncResponse(transferEvents []TransferEvent, requestId string) *TransferEventSyncResponse {
	this := TransferEventSyncResponse{}
	this.TransferEvents = transferEvents
	this.RequestId = requestId
	return &this
}

// NewTransferEventSyncResponseWithDefaults instantiates a new TransferEventSyncResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferEventSyncResponseWithDefaults() *TransferEventSyncResponse {
	this := TransferEventSyncResponse{}
	return &this
}

// GetTransferEvents returns the TransferEvents field value
func (o *TransferEventSyncResponse) GetTransferEvents() []TransferEvent {
	if o == nil {
		var ret []TransferEvent
		return ret
	}

	return o.TransferEvents
}

// GetTransferEventsOk returns a tuple with the TransferEvents field value
// and a boolean to check if the value has been set.
func (o *TransferEventSyncResponse) GetTransferEventsOk() (*[]TransferEvent, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferEvents, true
}

// SetTransferEvents sets field value
func (o *TransferEventSyncResponse) SetTransferEvents(v []TransferEvent) {
	o.TransferEvents = v
}

// GetRequestId returns the RequestId field value
func (o *TransferEventSyncResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferEventSyncResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferEventSyncResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferEventSyncResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transfer_events"] = o.TransferEvents
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferEventSyncResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferEventSyncResponse := _TransferEventSyncResponse{}

	if err = json.Unmarshal(bytes, &varTransferEventSyncResponse); err == nil {
		*o = TransferEventSyncResponse(varTransferEventSyncResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "transfer_events")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferEventSyncResponse struct {
	value *TransferEventSyncResponse
	isSet bool
}

func (v NullableTransferEventSyncResponse) Get() *TransferEventSyncResponse {
	return v.value
}

func (v *NullableTransferEventSyncResponse) Set(val *TransferEventSyncResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferEventSyncResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferEventSyncResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferEventSyncResponse(val *TransferEventSyncResponse) *NullableTransferEventSyncResponse {
	return &NullableTransferEventSyncResponse{value: val, isSet: true}
}

func (v NullableTransferEventSyncResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferEventSyncResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


