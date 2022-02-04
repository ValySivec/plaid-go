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

// TransferIntentCreateResponse Defines the response schema for `/transfer/intent/create`
type TransferIntentCreateResponse struct {
	TransferIntent TransferIntentCreate `json:"transfer_intent"`
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _TransferIntentCreateResponse TransferIntentCreateResponse

// NewTransferIntentCreateResponse instantiates a new TransferIntentCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransferIntentCreateResponse(transferIntent TransferIntentCreate, requestId string) *TransferIntentCreateResponse {
	this := TransferIntentCreateResponse{}
	this.TransferIntent = transferIntent
	this.RequestId = requestId
	return &this
}

// NewTransferIntentCreateResponseWithDefaults instantiates a new TransferIntentCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferIntentCreateResponseWithDefaults() *TransferIntentCreateResponse {
	this := TransferIntentCreateResponse{}
	return &this
}

// GetTransferIntent returns the TransferIntent field value
func (o *TransferIntentCreateResponse) GetTransferIntent() TransferIntentCreate {
	if o == nil {
		var ret TransferIntentCreate
		return ret
	}

	return o.TransferIntent
}

// GetTransferIntentOk returns a tuple with the TransferIntent field value
// and a boolean to check if the value has been set.
func (o *TransferIntentCreateResponse) GetTransferIntentOk() (*TransferIntentCreate, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TransferIntent, true
}

// SetTransferIntent sets field value
func (o *TransferIntentCreateResponse) SetTransferIntent(v TransferIntentCreate) {
	o.TransferIntent = v
}

// GetRequestId returns the RequestId field value
func (o *TransferIntentCreateResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *TransferIntentCreateResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *TransferIntentCreateResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o TransferIntentCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["transfer_intent"] = o.TransferIntent
	}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransferIntentCreateResponse) UnmarshalJSON(bytes []byte) (err error) {
	varTransferIntentCreateResponse := _TransferIntentCreateResponse{}

	if err = json.Unmarshal(bytes, &varTransferIntentCreateResponse); err == nil {
		*o = TransferIntentCreateResponse(varTransferIntentCreateResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "transfer_intent")
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransferIntentCreateResponse struct {
	value *TransferIntentCreateResponse
	isSet bool
}

func (v NullableTransferIntentCreateResponse) Get() *TransferIntentCreateResponse {
	return v.value
}

func (v *NullableTransferIntentCreateResponse) Set(val *TransferIntentCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTransferIntentCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTransferIntentCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransferIntentCreateResponse(val *TransferIntentCreateResponse) *NullableTransferIntentCreateResponse {
	return &NullableTransferIntentCreateResponse{value: val, isSet: true}
}

func (v NullableTransferIntentCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransferIntentCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


