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

// TransactionsGetRequestOptions An optional object to be used with the request. If specified, `options` must not be `null`.
type TransactionsGetRequestOptions struct {
	// A list of `account_ids` to retrieve for the Item  Note: An error will be returned if a provided `account_id` is not associated with the Item.
	AccountIds *[]string `json:"account_ids,omitempty"`
	// The number of transactions to fetch.
	Count *int32 `json:"count,omitempty"`
	// The number of transactions to skip. The default value is 0.
	Offset *int32 `json:"offset,omitempty"`
	// Include the raw unparsed transaction description from the financial institution. This field is disabled by default. If you need this information in addition to the parsed data provided, contact your Plaid Account Manager.
	IncludeOriginalDescription NullableBool `json:"include_original_description,omitempty"`
	// Include the `personal_finance_category` object in the response. This feature is currently in beta – to request access, contact transactions-feedback@plaid.com.
	IncludePersonalFinanceCategoryBeta *bool `json:"include_personal_finance_category_beta,omitempty"`
}

// NewTransactionsGetRequestOptions instantiates a new TransactionsGetRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransactionsGetRequestOptions() *TransactionsGetRequestOptions {
	this := TransactionsGetRequestOptions{}
	var count int32 = 100
	this.Count = &count
	var offset int32 = 0
	this.Offset = &offset
	var includeOriginalDescription bool = false
	this.IncludeOriginalDescription = *NewNullableBool(&includeOriginalDescription)
	var includePersonalFinanceCategoryBeta bool = false
	this.IncludePersonalFinanceCategoryBeta = &includePersonalFinanceCategoryBeta
	return &this
}

// NewTransactionsGetRequestOptionsWithDefaults instantiates a new TransactionsGetRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransactionsGetRequestOptionsWithDefaults() *TransactionsGetRequestOptions {
	this := TransactionsGetRequestOptions{}
	var count int32 = 100
	this.Count = &count
	var offset int32 = 0
	this.Offset = &offset
	var includeOriginalDescription bool = false
	this.IncludeOriginalDescription = *NewNullableBool(&includeOriginalDescription)
	var includePersonalFinanceCategoryBeta bool = false
	this.IncludePersonalFinanceCategoryBeta = &includePersonalFinanceCategoryBeta
	return &this
}

// GetAccountIds returns the AccountIds field value if set, zero value otherwise.
func (o *TransactionsGetRequestOptions) GetAccountIds() []string {
	if o == nil || o.AccountIds == nil {
		var ret []string
		return ret
	}
	return *o.AccountIds
}

// GetAccountIdsOk returns a tuple with the AccountIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsGetRequestOptions) GetAccountIdsOk() (*[]string, bool) {
	if o == nil || o.AccountIds == nil {
		return nil, false
	}
	return o.AccountIds, true
}

// HasAccountIds returns a boolean if a field has been set.
func (o *TransactionsGetRequestOptions) HasAccountIds() bool {
	if o != nil && o.AccountIds != nil {
		return true
	}

	return false
}

// SetAccountIds gets a reference to the given []string and assigns it to the AccountIds field.
func (o *TransactionsGetRequestOptions) SetAccountIds(v []string) {
	o.AccountIds = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *TransactionsGetRequestOptions) GetCount() int32 {
	if o == nil || o.Count == nil {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsGetRequestOptions) GetCountOk() (*int32, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *TransactionsGetRequestOptions) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *TransactionsGetRequestOptions) SetCount(v int32) {
	o.Count = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *TransactionsGetRequestOptions) GetOffset() int32 {
	if o == nil || o.Offset == nil {
		var ret int32
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsGetRequestOptions) GetOffsetOk() (*int32, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *TransactionsGetRequestOptions) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int32 and assigns it to the Offset field.
func (o *TransactionsGetRequestOptions) SetOffset(v int32) {
	o.Offset = &v
}

// GetIncludeOriginalDescription returns the IncludeOriginalDescription field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TransactionsGetRequestOptions) GetIncludeOriginalDescription() bool {
	if o == nil || o.IncludeOriginalDescription.Get() == nil {
		var ret bool
		return ret
	}
	return *o.IncludeOriginalDescription.Get()
}

// GetIncludeOriginalDescriptionOk returns a tuple with the IncludeOriginalDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TransactionsGetRequestOptions) GetIncludeOriginalDescriptionOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return o.IncludeOriginalDescription.Get(), o.IncludeOriginalDescription.IsSet()
}

// HasIncludeOriginalDescription returns a boolean if a field has been set.
func (o *TransactionsGetRequestOptions) HasIncludeOriginalDescription() bool {
	if o != nil && o.IncludeOriginalDescription.IsSet() {
		return true
	}

	return false
}

// SetIncludeOriginalDescription gets a reference to the given NullableBool and assigns it to the IncludeOriginalDescription field.
func (o *TransactionsGetRequestOptions) SetIncludeOriginalDescription(v bool) {
	o.IncludeOriginalDescription.Set(&v)
}
// SetIncludeOriginalDescriptionNil sets the value for IncludeOriginalDescription to be an explicit nil
func (o *TransactionsGetRequestOptions) SetIncludeOriginalDescriptionNil() {
	o.IncludeOriginalDescription.Set(nil)
}

// UnsetIncludeOriginalDescription ensures that no value is present for IncludeOriginalDescription, not even an explicit nil
func (o *TransactionsGetRequestOptions) UnsetIncludeOriginalDescription() {
	o.IncludeOriginalDescription.Unset()
}

// GetIncludePersonalFinanceCategoryBeta returns the IncludePersonalFinanceCategoryBeta field value if set, zero value otherwise.
func (o *TransactionsGetRequestOptions) GetIncludePersonalFinanceCategoryBeta() bool {
	if o == nil || o.IncludePersonalFinanceCategoryBeta == nil {
		var ret bool
		return ret
	}
	return *o.IncludePersonalFinanceCategoryBeta
}

// GetIncludePersonalFinanceCategoryBetaOk returns a tuple with the IncludePersonalFinanceCategoryBeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransactionsGetRequestOptions) GetIncludePersonalFinanceCategoryBetaOk() (*bool, bool) {
	if o == nil || o.IncludePersonalFinanceCategoryBeta == nil {
		return nil, false
	}
	return o.IncludePersonalFinanceCategoryBeta, true
}

// HasIncludePersonalFinanceCategoryBeta returns a boolean if a field has been set.
func (o *TransactionsGetRequestOptions) HasIncludePersonalFinanceCategoryBeta() bool {
	if o != nil && o.IncludePersonalFinanceCategoryBeta != nil {
		return true
	}

	return false
}

// SetIncludePersonalFinanceCategoryBeta gets a reference to the given bool and assigns it to the IncludePersonalFinanceCategoryBeta field.
func (o *TransactionsGetRequestOptions) SetIncludePersonalFinanceCategoryBeta(v bool) {
	o.IncludePersonalFinanceCategoryBeta = &v
}

func (o TransactionsGetRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountIds != nil {
		toSerialize["account_ids"] = o.AccountIds
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.IncludeOriginalDescription.IsSet() {
		toSerialize["include_original_description"] = o.IncludeOriginalDescription.Get()
	}
	if o.IncludePersonalFinanceCategoryBeta != nil {
		toSerialize["include_personal_finance_category_beta"] = o.IncludePersonalFinanceCategoryBeta
	}
	return json.Marshal(toSerialize)
}

type NullableTransactionsGetRequestOptions struct {
	value *TransactionsGetRequestOptions
	isSet bool
}

func (v NullableTransactionsGetRequestOptions) Get() *TransactionsGetRequestOptions {
	return v.value
}

func (v *NullableTransactionsGetRequestOptions) Set(val *TransactionsGetRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionsGetRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionsGetRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionsGetRequestOptions(val *TransactionsGetRequestOptions) *NullableTransactionsGetRequestOptions {
	return &NullableTransactionsGetRequestOptions{value: val, isSet: true}
}

func (v NullableTransactionsGetRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionsGetRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


