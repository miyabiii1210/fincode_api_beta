package fincode

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// 決済登録
func PaymentRegistration(ctx context.Context, internal *PaymentRegistrationRequest) (*PaymentRegistrationResponse, error) {
	path := fmt.Sprint("/v1/payments")
	req := NewRequest(http.MethodPost, path, internal)

	res, err := sendRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	ret := new(PaymentRegistrationResponse)
	if err = json.Unmarshal(res, ret); err != nil {
		return nil, err
	}

	return ret, nil
}

type PaymentRegistrationRequest struct {
	PayType        string `json:"pay_type"`
	ID             string `json:"id"`
	JobCode        string `json:"job_code"`
	Amount         string `json:"amount"`
	Tax            string `json:"tax"`
	ItemCode       string `json:"item_code"`
	ClientField1   string `json:"client_field_1"`
	ClientField2   string `json:"client_field_2"`
	ClientField3   string `json:"client_field_3"`
	SendURL        string `json:"send_url"`
	TdsType        string `json:"tds_type"`
	TdTenantName   string `json:"td_tenant_name"`
	SubscriptionID string `json:"subscription_id"`
}

type PaymentRegistrationResponse struct {
	ShopID          string `json:"shop_id"`
	ID              string `json:"id"`
	PayType         string `json:"pay_type"`
	Status          string `json:"status"`
	AccessID        string `json:"access_id"`
	ProcessDate     string `json:"process_date"`
	JobCode         string `json:"job_code"`
	ItemCode        string `json:"item_code"`
	Amount          int    `json:"amount"`
	Tax             int    `json:"tax"`
	TotalAmount     int    `json:"total_amount"`
	CustomerGroupID string `json:"customer_group_id"`
	CustomerID      string `json:"customer_id"`
	CardNo          string `json:"card_no"`
	CardID          string `json:"card_id"`
	Expire          string `json:"expire"`
	HolderName      string `json:"holder_name"`
	CardNoHash      string `json:"card_no_hash"`
	Method          string `json:"method"`
	PayTimes        string `json:"pay_times"`
	Forward         string `json:"forward"`
	Issuer          string `json:"issuer"`
	TransactionID   string `json:"transaction_id"`
	Approve         string `json:"approve"`
	AuthMaxDate     string `json:"auth_max_date"`
	ClientField1    string `json:"client_field_1"`
	ClientField2    string `json:"client_field_2"`
	ClientField3    string `json:"client_field_3"`
	TdsType         string `json:"tds_type"`
	Tds2Type        string `json:"tds2_type"`
	Tds2RetURL      string `json:"tds2_ret_url"`
	Tds2Status      string `json:"tds2_status"`
	MerchantName    string `json:"merchant_name"`
	SendURL         string `json:"send_url"`
	SubscriptionID  string `json:"subscription_id"`
	ErrorCode       string `json:"error_code"`
	Created         string `json:"created"`
	Updated         string `json:"updated"`
}
