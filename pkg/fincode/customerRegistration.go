package fincode

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// 顧客登録
func CustomerRegistration(ctx context.Context, internal *CustomerRegistrationRequest) (*CustomerRegistrationResponse, error) {
	path := fmt.Sprint("/v1/customers")
	req := NewRequest(http.MethodPost, path, internal)

	res, err := sendRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	ret := new(CustomerRegistrationResponse)
	if err = json.Unmarshal(res, ret); err != nil {
		return nil, err
	}

	return ret, nil
}

type CustomerRegistrationRequest struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PhoneCc      string `json:"phone_cc"`
	PhoneNo      string `json:"phone_no"`
	AddrCity     string `json:"addr_city"`
	AddrCountry  string `json:"addr_country"`
	AddrLine1    string `json:"addr_line_1"`
	AddrLine2    string `json:"addr_line_2"`
	AddrLine3    string `json:"addr_line_3"`
	AddrPostCode string `json:"addr_post_code"`
	AddrState    string `json:"addr_state"`
}

type CustomerRegistrationResponse struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Email            string `json:"email"`
	PhoneCc          string `json:"phone_cc"`
	PhoneNo          string `json:"phone_no"`
	AddrCity         string `json:"addr_city"`
	AddrCountry      string `json:"addr_country"`
	AddrLine1        string `json:"addr_line_1"`
	AddrLine2        string `json:"addr_line_2"`
	AddrLine3        string `json:"addr_line_3"`
	AddrPostCode     string `json:"addr_post_code"`
	AddrState        string `json:"addr_state"`
	CardRegistration string `json:"card_registration"`
	Created          string `json:"created"`
	Updated          string `json:"updated"`
}
