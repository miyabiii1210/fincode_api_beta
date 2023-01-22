package fincode

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func GenerateCardRegistrationURL(ctx context.Context,
	internal *GenerateCardRegistrationURLRequest) (*GenerateCardRegistrationURLResponse, error) {

	path := fmt.Sprint("/v1/card_sessions")
	req := NewRequest(http.MethodPost, path, internal)

	res, err := sendRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	ret := new(GenerateCardRegistrationURLResponse)
	if err = json.Unmarshal(res, ret); err != nil {
		return nil, err
	}

	return ret, nil
}

type GenerateCardRegistrationURLRequest struct {
	SuccessURL      string `json:"success_url"`
	CancelURL       string `json:"cancel_url"`
	ShopServiceName string `json:"shop_service_name"`
	CustomerID      string `json:"customer_id"`
	CustomerName    string `json:"customer_name"`
}

type GenerateCardRegistrationURLResponse struct {
	ID              string `json:"id"`
	LinkURL         string `json:"link_url"`
	SuccessURL      string `json:"success_url"`
	CancelURL       string `json:"cancel_url"`
	Status          string `json:"status"`
	Expire          string `json:"expire"`
	ShopServiceName string `json:"shop_service_name"`
	CustomerID      string `json:"customer_id"`
	CustomerName    string `json:"customer_name"`
	Created         string `json:"created"`
	Updated         string `json:"updated"`
}
