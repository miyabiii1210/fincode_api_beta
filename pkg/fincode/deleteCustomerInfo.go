package fincode

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func DeleteCustomerInfo(ctx context.Context, customerId string) (*DeleteCustomerInfoResponse, error) {
	path := fmt.Sprintf("/v1/customers/%s", customerId)
	req := NewRequest(http.MethodDelete, path, nil)

	res, err := sendRequest(ctx, req)
	if err != nil {
		return nil, err
	}

	ret := new(DeleteCustomerInfoResponse)
	if err = json.Unmarshal(res, ret); err != nil {
		return nil, err
	}

	return ret, nil
}

type DeleteCustomerInfoRequest struct {
}

type DeleteCustomerInfoResponse struct {
}
