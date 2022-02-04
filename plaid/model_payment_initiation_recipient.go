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

// PaymentInitiationRecipient PaymentInitiationRecipient defines a payment initiation recipient
type PaymentInitiationRecipient struct {
	// The ID of the recipient.
	RecipientId string `json:"recipient_id"`
	// The name of the recipient.
	Name string `json:"name"`
	Address NullablePaymentInitiationAddress `json:"address,omitempty"`
	// The International Bank Account Number (IBAN) for the recipient.
	Iban NullableString `json:"iban,omitempty"`
	Bacs NullableRecipientBACSNullable `json:"bacs,omitempty"`
	// The EMI (E-Money Institution) recipient that this recipient is associated with, if any. This EMI recipient is used as an intermediary account to enable Plaid to reconcile the settlement of funds for Payment Initiation requests.
	EmiRecipientId NullableString `json:"emi_recipient_id,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PaymentInitiationRecipient PaymentInitiationRecipient

// NewPaymentInitiationRecipient instantiates a new PaymentInitiationRecipient object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentInitiationRecipient(recipientId string, name string) *PaymentInitiationRecipient {
	this := PaymentInitiationRecipient{}
	this.RecipientId = recipientId
	this.Name = name
	return &this
}

// NewPaymentInitiationRecipientWithDefaults instantiates a new PaymentInitiationRecipient object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentInitiationRecipientWithDefaults() *PaymentInitiationRecipient {
	this := PaymentInitiationRecipient{}
	return &this
}

// GetRecipientId returns the RecipientId field value
func (o *PaymentInitiationRecipient) GetRecipientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RecipientId
}

// GetRecipientIdOk returns a tuple with the RecipientId field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRecipient) GetRecipientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RecipientId, true
}

// SetRecipientId sets field value
func (o *PaymentInitiationRecipient) SetRecipientId(v string) {
	o.RecipientId = v
}

// GetName returns the Name field value
func (o *PaymentInitiationRecipient) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PaymentInitiationRecipient) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PaymentInitiationRecipient) SetName(v string) {
	o.Name = v
}

// GetAddress returns the Address field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationRecipient) GetAddress() PaymentInitiationAddress {
	if o == nil || o.Address.Get() == nil {
		var ret PaymentInitiationAddress
		return ret
	}
	return *o.Address.Get()
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationRecipient) GetAddressOk() (*PaymentInitiationAddress, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Address.Get(), o.Address.IsSet()
}

// HasAddress returns a boolean if a field has been set.
func (o *PaymentInitiationRecipient) HasAddress() bool {
	if o != nil && o.Address.IsSet() {
		return true
	}

	return false
}

// SetAddress gets a reference to the given NullablePaymentInitiationAddress and assigns it to the Address field.
func (o *PaymentInitiationRecipient) SetAddress(v PaymentInitiationAddress) {
	o.Address.Set(&v)
}
// SetAddressNil sets the value for Address to be an explicit nil
func (o *PaymentInitiationRecipient) SetAddressNil() {
	o.Address.Set(nil)
}

// UnsetAddress ensures that no value is present for Address, not even an explicit nil
func (o *PaymentInitiationRecipient) UnsetAddress() {
	o.Address.Unset()
}

// GetIban returns the Iban field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationRecipient) GetIban() string {
	if o == nil || o.Iban.Get() == nil {
		var ret string
		return ret
	}
	return *o.Iban.Get()
}

// GetIbanOk returns a tuple with the Iban field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationRecipient) GetIbanOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Iban.Get(), o.Iban.IsSet()
}

// HasIban returns a boolean if a field has been set.
func (o *PaymentInitiationRecipient) HasIban() bool {
	if o != nil && o.Iban.IsSet() {
		return true
	}

	return false
}

// SetIban gets a reference to the given NullableString and assigns it to the Iban field.
func (o *PaymentInitiationRecipient) SetIban(v string) {
	o.Iban.Set(&v)
}
// SetIbanNil sets the value for Iban to be an explicit nil
func (o *PaymentInitiationRecipient) SetIbanNil() {
	o.Iban.Set(nil)
}

// UnsetIban ensures that no value is present for Iban, not even an explicit nil
func (o *PaymentInitiationRecipient) UnsetIban() {
	o.Iban.Unset()
}

// GetBacs returns the Bacs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationRecipient) GetBacs() RecipientBACSNullable {
	if o == nil || o.Bacs.Get() == nil {
		var ret RecipientBACSNullable
		return ret
	}
	return *o.Bacs.Get()
}

// GetBacsOk returns a tuple with the Bacs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationRecipient) GetBacsOk() (*RecipientBACSNullable, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Bacs.Get(), o.Bacs.IsSet()
}

// HasBacs returns a boolean if a field has been set.
func (o *PaymentInitiationRecipient) HasBacs() bool {
	if o != nil && o.Bacs.IsSet() {
		return true
	}

	return false
}

// SetBacs gets a reference to the given NullableRecipientBACSNullable and assigns it to the Bacs field.
func (o *PaymentInitiationRecipient) SetBacs(v RecipientBACSNullable) {
	o.Bacs.Set(&v)
}
// SetBacsNil sets the value for Bacs to be an explicit nil
func (o *PaymentInitiationRecipient) SetBacsNil() {
	o.Bacs.Set(nil)
}

// UnsetBacs ensures that no value is present for Bacs, not even an explicit nil
func (o *PaymentInitiationRecipient) UnsetBacs() {
	o.Bacs.Unset()
}

// GetEmiRecipientId returns the EmiRecipientId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PaymentInitiationRecipient) GetEmiRecipientId() string {
	if o == nil || o.EmiRecipientId.Get() == nil {
		var ret string
		return ret
	}
	return *o.EmiRecipientId.Get()
}

// GetEmiRecipientIdOk returns a tuple with the EmiRecipientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PaymentInitiationRecipient) GetEmiRecipientIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return o.EmiRecipientId.Get(), o.EmiRecipientId.IsSet()
}

// HasEmiRecipientId returns a boolean if a field has been set.
func (o *PaymentInitiationRecipient) HasEmiRecipientId() bool {
	if o != nil && o.EmiRecipientId.IsSet() {
		return true
	}

	return false
}

// SetEmiRecipientId gets a reference to the given NullableString and assigns it to the EmiRecipientId field.
func (o *PaymentInitiationRecipient) SetEmiRecipientId(v string) {
	o.EmiRecipientId.Set(&v)
}
// SetEmiRecipientIdNil sets the value for EmiRecipientId to be an explicit nil
func (o *PaymentInitiationRecipient) SetEmiRecipientIdNil() {
	o.EmiRecipientId.Set(nil)
}

// UnsetEmiRecipientId ensures that no value is present for EmiRecipientId, not even an explicit nil
func (o *PaymentInitiationRecipient) UnsetEmiRecipientId() {
	o.EmiRecipientId.Unset()
}

func (o PaymentInitiationRecipient) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["recipient_id"] = o.RecipientId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Address.IsSet() {
		toSerialize["address"] = o.Address.Get()
	}
	if o.Iban.IsSet() {
		toSerialize["iban"] = o.Iban.Get()
	}
	if o.Bacs.IsSet() {
		toSerialize["bacs"] = o.Bacs.Get()
	}
	if o.EmiRecipientId.IsSet() {
		toSerialize["emi_recipient_id"] = o.EmiRecipientId.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PaymentInitiationRecipient) UnmarshalJSON(bytes []byte) (err error) {
	varPaymentInitiationRecipient := _PaymentInitiationRecipient{}

	if err = json.Unmarshal(bytes, &varPaymentInitiationRecipient); err == nil {
		*o = PaymentInitiationRecipient(varPaymentInitiationRecipient)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "recipient_id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "address")
		delete(additionalProperties, "iban")
		delete(additionalProperties, "bacs")
		delete(additionalProperties, "emi_recipient_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePaymentInitiationRecipient struct {
	value *PaymentInitiationRecipient
	isSet bool
}

func (v NullablePaymentInitiationRecipient) Get() *PaymentInitiationRecipient {
	return v.value
}

func (v *NullablePaymentInitiationRecipient) Set(val *PaymentInitiationRecipient) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentInitiationRecipient) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentInitiationRecipient) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentInitiationRecipient(val *PaymentInitiationRecipient) *NullablePaymentInitiationRecipient {
	return &NullablePaymentInitiationRecipient{value: val, isSet: true}
}

func (v NullablePaymentInitiationRecipient) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentInitiationRecipient) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


