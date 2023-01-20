package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

// https://docs.fincode.jp/api
const apiKey = "Bearer ***"
const baseURL = "https://api.test.fincode.jp"
const method = "POST"
const id = "o_test_id"
const data = `{
 	"pay_type": "Card",
	"job_code": "CAPTURE",
	"amount": "500"
	}`

func main() {

	endpoint := "/v1/payments"
	URL := baseURL + endpoint

	req, err := http.NewRequest(
		method,
		URL,
		bytes.NewBuffer([]byte(data)),
	)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	bodyJsonBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("ERROR")
		fmt.Println(string(bodyJsonBytes))
		return
	}

	fmt.Println("SUCCESS")
	fmt.Println(string(bodyJsonBytes))
}
