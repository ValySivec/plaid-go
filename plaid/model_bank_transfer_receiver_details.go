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
)

// BankTransferReceiverDetails The receiver details if the type of this event is `reciever_pending` or `reciever_posted`. Null value otherwise.
type BankTransferReceiverDetails struct {
	// The sign of the available balance for the receiver bank account associated with the receiver event at the time the matching transaction was found. Can be `positive`, `negative`, or null if the balance was not available at the time.
	AvailableBalance NullableString `json:"available_balance"`
	AdditionalProperties map[string]interface{}
}

type _BankTransferReceiverDetails BankTransferReceiverDetails

// NewBankTransferReceiverDetails instantiates a new BankTransferReceiverDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBankTransferReceiverDetails(availableBalance NullableString) *BankTransferReceiverDetails {
	this := BankTransferReceiverDetails{}
	this.AvailableBalance = availableBalance
	return &this
}

// NewBankTransferReceiverDetailsWithDefaults instantiates a new BankTransferReceiverDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBankTransferReceiverDetailsWithDefaults() *BankTransferReceiverDetails {
	this := BankTransferReceiverDetails{}
	return &this
}

// GetAvailableBalance returns the AvailableBalance field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BankTransferReceiverDetails) GetAvailableBalance() string {
	if o == nil || o.AvailableBalance.Get() == nil {
		var ret string
		return ret
	}

	return *o.AvailableBalance.Get()
}

// GetAvailableBalanceOk returns a tuple with the AvailableBalance field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BankTransferReceiverDetails) GetAvailableBalanceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.AvailableBalance.Get(), o.AvailableBalance.IsSet()
}

// SetAvailableBalance sets field value
func (o *BankTransferReceiverDetails) SetAvailableBalance(v string) {
	o.AvailableBalance.Set(&v)
}

func (o BankTransferReceiverDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["available_balance"] = o.AvailableBalance.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BankTransferReceiverDetails) UnmarshalJSON(bytes []byte) (err error) {
	varBankTransferReceiverDetails := _BankTransferReceiverDetails{}

	if err = json.Unmarshal(bytes, &varBankTransferReceiverDetails); err == nil {
		*o = BankTransferReceiverDetails(varBankTransferReceiverDetails)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "available_balance")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBankTransferReceiverDetails struct {
	value *BankTransferReceiverDetails
	isSet bool
}

func (v NullableBankTransferReceiverDetails) Get() *BankTransferReceiverDetails {
	return v.value
}

func (v *NullableBankTransferReceiverDetails) Set(val *BankTransferReceiverDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableBankTransferReceiverDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableBankTransferReceiverDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBankTransferReceiverDetails(val *BankTransferReceiverDetails) *NullableBankTransferReceiverDetails {
	return &NullableBankTransferReceiverDetails{value: val, isSet: true}
}

func (v NullableBankTransferReceiverDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBankTransferReceiverDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


