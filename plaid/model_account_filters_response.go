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

// AccountFiltersResponse The `account_filters` specified in the original call to `/link/token/create`. 
type AccountFiltersResponse struct {
	Depository *DepositoryFilter `json:"depository,omitempty"`
	Credit *CreditFilter `json:"credit,omitempty"`
	Loan *LoanFilter `json:"loan,omitempty"`
	Investment *InvestmentFilter `json:"investment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountFiltersResponse AccountFiltersResponse

// NewAccountFiltersResponse instantiates a new AccountFiltersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountFiltersResponse() *AccountFiltersResponse {
	this := AccountFiltersResponse{}
	return &this
}

// NewAccountFiltersResponseWithDefaults instantiates a new AccountFiltersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountFiltersResponseWithDefaults() *AccountFiltersResponse {
	this := AccountFiltersResponse{}
	return &this
}

// GetDepository returns the Depository field value if set, zero value otherwise.
func (o *AccountFiltersResponse) GetDepository() DepositoryFilter {
	if o == nil || o.Depository == nil {
		var ret DepositoryFilter
		return ret
	}
	return *o.Depository
}

// GetDepositoryOk returns a tuple with the Depository field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountFiltersResponse) GetDepositoryOk() (*DepositoryFilter, bool) {
	if o == nil || o.Depository == nil {
		return nil, false
	}
	return o.Depository, true
}

// HasDepository returns a boolean if a field has been set.
func (o *AccountFiltersResponse) HasDepository() bool {
	if o != nil && o.Depository != nil {
		return true
	}

	return false
}

// SetDepository gets a reference to the given DepositoryFilter and assigns it to the Depository field.
func (o *AccountFiltersResponse) SetDepository(v DepositoryFilter) {
	o.Depository = &v
}

// GetCredit returns the Credit field value if set, zero value otherwise.
func (o *AccountFiltersResponse) GetCredit() CreditFilter {
	if o == nil || o.Credit == nil {
		var ret CreditFilter
		return ret
	}
	return *o.Credit
}

// GetCreditOk returns a tuple with the Credit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountFiltersResponse) GetCreditOk() (*CreditFilter, bool) {
	if o == nil || o.Credit == nil {
		return nil, false
	}
	return o.Credit, true
}

// HasCredit returns a boolean if a field has been set.
func (o *AccountFiltersResponse) HasCredit() bool {
	if o != nil && o.Credit != nil {
		return true
	}

	return false
}

// SetCredit gets a reference to the given CreditFilter and assigns it to the Credit field.
func (o *AccountFiltersResponse) SetCredit(v CreditFilter) {
	o.Credit = &v
}

// GetLoan returns the Loan field value if set, zero value otherwise.
func (o *AccountFiltersResponse) GetLoan() LoanFilter {
	if o == nil || o.Loan == nil {
		var ret LoanFilter
		return ret
	}
	return *o.Loan
}

// GetLoanOk returns a tuple with the Loan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountFiltersResponse) GetLoanOk() (*LoanFilter, bool) {
	if o == nil || o.Loan == nil {
		return nil, false
	}
	return o.Loan, true
}

// HasLoan returns a boolean if a field has been set.
func (o *AccountFiltersResponse) HasLoan() bool {
	if o != nil && o.Loan != nil {
		return true
	}

	return false
}

// SetLoan gets a reference to the given LoanFilter and assigns it to the Loan field.
func (o *AccountFiltersResponse) SetLoan(v LoanFilter) {
	o.Loan = &v
}

// GetInvestment returns the Investment field value if set, zero value otherwise.
func (o *AccountFiltersResponse) GetInvestment() InvestmentFilter {
	if o == nil || o.Investment == nil {
		var ret InvestmentFilter
		return ret
	}
	return *o.Investment
}

// GetInvestmentOk returns a tuple with the Investment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountFiltersResponse) GetInvestmentOk() (*InvestmentFilter, bool) {
	if o == nil || o.Investment == nil {
		return nil, false
	}
	return o.Investment, true
}

// HasInvestment returns a boolean if a field has been set.
func (o *AccountFiltersResponse) HasInvestment() bool {
	if o != nil && o.Investment != nil {
		return true
	}

	return false
}

// SetInvestment gets a reference to the given InvestmentFilter and assigns it to the Investment field.
func (o *AccountFiltersResponse) SetInvestment(v InvestmentFilter) {
	o.Investment = &v
}

func (o AccountFiltersResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Depository != nil {
		toSerialize["depository"] = o.Depository
	}
	if o.Credit != nil {
		toSerialize["credit"] = o.Credit
	}
	if o.Loan != nil {
		toSerialize["loan"] = o.Loan
	}
	if o.Investment != nil {
		toSerialize["investment"] = o.Investment
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccountFiltersResponse) UnmarshalJSON(bytes []byte) (err error) {
	varAccountFiltersResponse := _AccountFiltersResponse{}

	if err = json.Unmarshal(bytes, &varAccountFiltersResponse); err == nil {
		*o = AccountFiltersResponse(varAccountFiltersResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "depository")
		delete(additionalProperties, "credit")
		delete(additionalProperties, "loan")
		delete(additionalProperties, "investment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountFiltersResponse struct {
	value *AccountFiltersResponse
	isSet bool
}

func (v NullableAccountFiltersResponse) Get() *AccountFiltersResponse {
	return v.value
}

func (v *NullableAccountFiltersResponse) Set(val *AccountFiltersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountFiltersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountFiltersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountFiltersResponse(val *AccountFiltersResponse) *NullableAccountFiltersResponse {
	return &NullableAccountFiltersResponse{value: val, isSet: true}
}

func (v NullableAccountFiltersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountFiltersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


