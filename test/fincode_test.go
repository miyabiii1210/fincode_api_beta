package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/miyabiii1210/fincode_api_beta/go/pkg/fincode"
)

func Test(t *testing.T) {
	fmt.Println("test.")
}

func TestGetPaymentInfo(t *testing.T) {
	ctx := context.TODO()
	internal := fincode.GetPaymentRequest{
		PayType: "Card",
	}
	orderId := "o_CM******************Mg"

	ret, err := fincode.GetPaymentInfo(ctx, internal, orderId)
	if err != nil {
		t.Error(err)
	}
	t.Logf("result: %v", ret)
}
