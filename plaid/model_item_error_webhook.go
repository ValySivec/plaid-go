/*
 * The Plaid API
 *
 * The Plaid REST API. Please see https://plaid.com/docs/api for more details.
 *
 * API version: 2020-09-14_1.33.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plaid

import (
	"encoding/json"
)

// ItemErrorWebhook Fired when an error is encountered with an Item. The error can be resolved by having the user go through Link’s update mode.
type ItemErrorWebhook struct {
	// `ITEM`
	WebhookType string `json:"webhook_type"`
	// `ERROR`
	WebhookCode string `json:"webhook_code"`
	// The `item_id` of the Item associated with this webhook, warning, or error
	ItemId string `json:"item_id"`
	Error NullableError `json:"error"`
	AdditionalProperties map[string]interface{}
}

type _ItemErrorWebhook ItemErrorWebhook

// NewItemErrorWebhook instantiates a new ItemErrorWebhook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewItemErrorWebhook(webhookType string, webhookCode string, itemId string, error_ NullableError) *ItemErrorWebhook {
	this := ItemErrorWebhook{}
	this.WebhookType = webhookType
	this.WebhookCode = webhookCode
	this.ItemId = itemId
	this.Error = error_
	return &this
}

// NewItemErrorWebhookWithDefaults instantiates a new ItemErrorWebhook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemErrorWebhookWithDefaults() *ItemErrorWebhook {
	this := ItemErrorWebhook{}
	return &this
}

// GetWebhookType returns the WebhookType field value
func (o *ItemErrorWebhook) GetWebhookType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookType
}

// GetWebhookTypeOk returns a tuple with the WebhookType field value
// and a boolean to check if the value has been set.
func (o *ItemErrorWebhook) GetWebhookTypeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookType, true
}

// SetWebhookType sets field value
func (o *ItemErrorWebhook) SetWebhookType(v string) {
	o.WebhookType = v
}

// GetWebhookCode returns the WebhookCode field value
func (o *ItemErrorWebhook) GetWebhookCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebhookCode
}

// GetWebhookCodeOk returns a tuple with the WebhookCode field value
// and a boolean to check if the value has been set.
func (o *ItemErrorWebhook) GetWebhookCodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.WebhookCode, true
}

// SetWebhookCode sets field value
func (o *ItemErrorWebhook) SetWebhookCode(v string) {
	o.WebhookCode = v
}

// GetItemId returns the ItemId field value
func (o *ItemErrorWebhook) GetItemId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ItemId
}

// GetItemIdOk returns a tuple with the ItemId field value
// and a boolean to check if the value has been set.
func (o *ItemErrorWebhook) GetItemIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ItemId, true
}

// SetItemId sets field value
func (o *ItemErrorWebhook) SetItemId(v string) {
	o.ItemId = v
}

// GetError returns the Error field value
// If the value is explicit nil, the zero value for Error will be returned
func (o *ItemErrorWebhook) GetError() Error {
	if o == nil || o.Error.Get() == nil {
		var ret Error
		return ret
	}

	return *o.Error.Get()
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ItemErrorWebhook) GetErrorOk() (*Error, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Error.Get(), o.Error.IsSet()
}

// SetError sets field value
func (o *ItemErrorWebhook) SetError(v Error) {
	o.Error.Set(&v)
}

func (o ItemErrorWebhook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["webhook_type"] = o.WebhookType
	}
	if true {
		toSerialize["webhook_code"] = o.WebhookCode
	}
	if true {
		toSerialize["item_id"] = o.ItemId
	}
	if true {
		toSerialize["error"] = o.Error.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ItemErrorWebhook) UnmarshalJSON(bytes []byte) (err error) {
	varItemErrorWebhook := _ItemErrorWebhook{}

	if err = json.Unmarshal(bytes, &varItemErrorWebhook); err == nil {
		*o = ItemErrorWebhook(varItemErrorWebhook)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "webhook_type")
		delete(additionalProperties, "webhook_code")
		delete(additionalProperties, "item_id")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableItemErrorWebhook struct {
	value *ItemErrorWebhook
	isSet bool
}

func (v NullableItemErrorWebhook) Get() *ItemErrorWebhook {
	return v.value
}

func (v *NullableItemErrorWebhook) Set(val *ItemErrorWebhook) {
	v.value = val
	v.isSet = true
}

func (v NullableItemErrorWebhook) IsSet() bool {
	return v.isSet
}

func (v *NullableItemErrorWebhook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItemErrorWebhook(val *ItemErrorWebhook) *NullableItemErrorWebhook {
	return &NullableItemErrorWebhook{value: val, isSet: true}
}

func (v NullableItemErrorWebhook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItemErrorWebhook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


