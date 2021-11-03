/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.44.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// BankTransferCancelResponse Defines the response schema for `/bank_transfer/cancel`
type BankTransferCancelResponse struct {
	// A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive.
	RequestId string `json:"request_id"`
	AdditionalProperties map[string]interface{}
}

type _BankTransferCancelResponse BankTransferCancelResponse

// NewBankTransferCancelResponse instantiates a new BankTransferCancelResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferCancelResponse(requestId string) *BankTransferCancelResponse {
	this := BankTransferCancelResponse{}
	this.RequestId = requestId
	return &this
}

// NewBankTransferCancelResponseWithDefaults instantiates a new BankTransferCancelResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferCancelResponseWithDefaults() *BankTransferCancelResponse {
	this := BankTransferCancelResponse{}
	return &this
}

// GetRequestId returns the RequestId field value
func (o *BankTransferCancelResponse) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *BankTransferCancelResponse) GetRequestIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *BankTransferCancelResponse) SetRequestId(v string) {
	o.RequestId = v
}

func (o BankTransferCancelResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request_id"] = o.RequestId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BankTransferCancelResponse) UnmarshalJSON(bytes []byte) (err error) {
	varBankTransferCancelResponse := _BankTransferCancelResponse{}

	if err = json.Unmarshal(bytes, &varBankTransferCancelResponse); err == nil {
		*o = BankTransferCancelResponse(varBankTransferCancelResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "request_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBankTransferCancelResponse struct {
	value *BankTransferCancelResponse
	isSet bool
}

func (v NullableBankTransferCancelResponse) Get() *BankTransferCancelResponse {
	return v.value
}

func (v *NullableBankTransferCancelResponse) Set(val *BankTransferCancelResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferCancelResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferCancelResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferCancelResponse(val *BankTransferCancelResponse) *NullableBankTransferCancelResponse {
	return &NullableBankTransferCancelResponse{value: val, isSet: true}
}

func (v NullableBankTransferCancelResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferCancelResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


