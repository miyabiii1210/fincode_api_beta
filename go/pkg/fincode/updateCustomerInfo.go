package fincode

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func UpdateCustomerInfo(ctx context.Context, internal *UpdateCustomerInfoRequest, customerId string) (*UpdateCustomerInfoResponse, error) {
	path := fmt.Sprintf("/v1/customers/%s", customerId)
	req := NewRequest(http.MethodPut, path, internal)

	res, err := sendRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	ret := new(UpdateCustomerInfoResponse)
	if err = json.Unmarshal(res, ret); err != nil {
		return nil, err
	}

	return ret, nil
}

type UpdateCustomerInfoRequest struct {
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

type UpdateCustomerInfoResponse struct {
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
