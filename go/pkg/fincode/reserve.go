package fincode

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	baseURL   string
	ApiSecret string
)

const (
	PAY_TYPE_CARD = "Card"

	// CHECK:有効性確認, AUTH:仮売上, CAPTURE:売上確定
	JOB_CODE_CHECK   = "CHECK"
	JOB_CODE_AUTH    = "AUTH"
	JOB_CODE_CAPTURE = "CAPTURE"
)

func init() {
	baseURL = os.Getenv("FINCODE_BASE_URL")
	ApiSecret = os.Getenv("FINCODE_API_SECRET")
}

type Request struct {
	Method string
	Path   string
	Data   interface{}
}

func NewRequest(method, path string, data interface{}) *Request {
	return &Request{
		Method: method,
		Path:   path,
		Data:   data,
	}
}

func sendRequest(ctx context.Context, r *Request) ([]byte, error) {
	url := baseURL + r.Path

	var req *http.Request
	var err error
	if r.Method == http.MethodGet {
		req, err = http.NewRequestWithContext(ctx, r.Method, url, nil)
		if err != nil {
			return nil, err
		}
	} else {
		body, err := json.Marshal(r.Data)
		if err != nil {
			return nil, err
		}

		req, err = http.NewRequestWithContext(ctx, r.Method, url, bytes.NewBuffer(body))
		if err != nil {
			return nil, err
		}
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+ApiSecret)

	params := req.URL.Query()
	params.Add("pay_type", "Card")
	req.URL.RawQuery = params.Encode()

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	buf, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return buf, nil
}
