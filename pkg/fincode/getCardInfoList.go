package fincode

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetCardInfoList(ctx context.Context, customerId string) (*GetCardInfoListResponse, error) {
	path := fmt.Sprintf("/v1/customers/%s/cards", customerId)
	req := NewRequest(http.MethodGet, path, nil)

	res, err := sendRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	ret := new(GetCardInfoListResponse)
	if err = json.Unmarshal(res, ret); err != nil {
		return nil, err
	}

	return ret, nil
}

type GetCardInfoListRequest struct{}

type GetCardInfoListResponse struct {
	List []struct {
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
	} `json:"list"`
}
