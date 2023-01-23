package fincode

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func PaymentConfirm(ctx context.Context, internal *PaymentConfirmRequest, orderId string) (*PaymentConfirmResponse, error) {
	path := fmt.Sprintf("/v1/payments/%s/capture", orderId)
	req := NewRequest(http.MethodPut, path, internal)

	res, err := sendRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	ret := new(PaymentConfirmResponse)
	if err = json.Unmarshal(res, ret); err != nil {
		return nil, err
	}

	return ret, nil
}

type PaymentConfirmRequest struct {
	PayType  string `json:"pay_type"`
	AccessID string `json:"access_id"`
	Method   string `json:"method"`
	PayTime  string `json:"pay_time"`
}

type PaymentConfirmResponse struct {
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
	Brand           string `json:"brand"`
	ErrorCode       string `json:"error_code"`
	Created         string `json:"created"`
	Updated         string `json:"updated"`
}
