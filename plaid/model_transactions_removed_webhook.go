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

// TransactionsRemovedWebhook Fired when transaction(s) for an Item are deleted. The deleted transaction IDs are included in the webhook payload. Plaid will typically check for deleted transaction data several times a day.
type TransactionsRemovedWebhook struct {
	// `TRANSACTIONS`
	WebhookType string `json:"webhook_type"`
	// `TRANSACTIONS_REMOVED`
	WebhookCode string `json:"webhook_code"`
	Error NullableError `json:"error,omitempty"`
	// An array of `transaction_ids` corresponding to the removed transactions
	RemovedTransactions []string `json:"removed_transactions"`
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	AdditionalProperties map[string]interface{}
}

type _TransactionsRemovedWebhook TransactionsRemovedWebhook

// NewTransactionsRemovedWebhook instantiates a new TransactionsRemovedWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsRemovedWebhook(webhookType string, webhookCode string, removedTransactions []string, itemId string) *TransactionsRemovedWebhook {
	this := TransactionsRemovedWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.RemovedTransactions = removedTransactions
	this.ItemId = itemId
	return &this
}

// NewTransactionsRemovedWebhookWithDefaults instantiates a new TransactionsRemovedWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsRemovedWebhookWithDefaults() *TransactionsRemovedWebhook {
	this := TransactionsRemovedWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *TransactionsRemovedWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *TransactionsRemovedWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *TransactionsRemovedWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *TransactionsRemovedWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *TransactionsRemovedWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *TransactionsRemovedWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetError returns the Error field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionsRemovedWebhook) GetError() Error {
	if o == nil || o.Error.Get() == nil {
		var ret Error
		return ret
	}
	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionsRemovedWebhook) GetErrorOk() (*Error, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// HasError returns a boolean if a field has been set.
func (o *TransactionsRemovedWebhook) HasError() bool {
	if o != nil && o.Error.IsSet() {
		return true
	}

	return false
}

// SetError gets a reference to the given NullableError and assigns it to the Error field.
func (o *TransactionsRemovedWebhook) SetError(v Error) {
	o.Error.Set(&v)
}
// SetErrorNil sets the value for Error to be an explicit nil
func (o *TransactionsRemovedWebhook) SetErrorNil() {
	o.Error.Set(nil)
}

// UnsetError ensures that no value is present for Error, not even an explicit nil
func (o *TransactionsRemovedWebhook) UnsetError() {
	o.Error.Unset()
}

// GetRemovedTransactions returns the RemovedTransactions field value
func (o *TransactionsRemovedWebhook) GetRemovedTransactions() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RemovedTransactions
}

// GetRemovedTransactionsOk returns a tuple with the RemovedTransactions field value
// and a boolean to check if the value has been set.
func (o *TransactionsRemovedWebhook) GetRemovedTransactionsOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.RemovedTransactions, true
}

// SetRemovedTransactions sets field value
func (o *TransactionsRemovedWebhook) SetRemovedTransactions(v []string) {
	o.RemovedTransactions = v
}

// GetItemId returns the ItemId field value
func (o *TransactionsRemovedWebhook) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *TransactionsRemovedWebhook) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *TransactionsRemovedWebhook) SetItemId(v string) {
	o.ItemId = v
}

func (o TransactionsRemovedWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if o.Error.IsSet() {
		toSerialize["error"] = o.Error.Get()
	}
	if true {
		toSerialize["removed_transactions"] = o.RemovedTransactions
	}
	if true {
		toSerialize["item_id"] = o.ItemId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TransactionsRemovedWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varTransactionsRemovedWebhook := _TransactionsRemovedWebhook{}

	if err = json.Unmarshal(bytes, &varTransactionsRemovedWebhook); err == nil {
		*o = TransactionsRemovedWebhook(varTransactionsRemovedWebhook)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "error")
		delete(additionalProperties, "removed_transactions")
		delete(additionalProperties, "item_id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTransactionsRemovedWebhook struct {
	value *TransactionsRemovedWebhook
	isSet bool
}

func (v NullableTransactionsRemovedWebhook) Get() *TransactionsRemovedWebhook {
	return v.value
}

func (v *NullableTransactionsRemovedWebhook) Set(val *TransactionsRemovedWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsRemovedWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsRemovedWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsRemovedWebhook(val *TransactionsRemovedWebhook) *NullableTransactionsRemovedWebhook {
	return &NullableTransactionsRemovedWebhook{value: val, isSet: true}
}

func (v NullableTransactionsRemovedWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsRemovedWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


