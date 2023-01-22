package fincode

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func CardRegistration(ctx context.Context, internal CardRegistrationRequest, customerId string) (*CardRegistrationResponse, error) {
	path := fmt.Sprintf("/v1/customers/%s/cards", customerId)
	req := NewRequest(http.MethodPost, path, internal)

	res, err := sendRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	ret := new(CardRegistrationResponse)
	if err = json.Unmarshal(res, ret); err != nil {
		return nil, err
	}

	return ret, nil
}

type CardRegistrationRequest struct {
	DefaultFlag  string `json:"default_flag"`
	HolderName   string `json:"holder_name"`
	SecurityCode string `json:"security_code"`
	Token        string `json:"token"`
}

type CardRegistrationResponse struct {
	CustomerID  string `json:"customer_id"`
	ID          string `json:"id"`
	DefaultFlag string `json:"default_flag"`
	CardNo      string `json:"card_no"`
	Expire      string `json:"expire"`
	HolderName  string `json:"holder_name"`
	CardNoHash  string `json:"card_no_hash"`
	Created     string `json:"created"`
	Updated     string `json:"updated"`
	Type        string `json:"type"`
	Brand       string `json:"brand"`
}
