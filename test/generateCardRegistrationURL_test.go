package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/miyabiii1210/fincode_api_beta/go/pkg/fincode"
	"github.com/miyabiii1210/fincode_api_beta/go/pkg/util"
)

// test-card URL: https://docs.fincode.jp/develop_support/testcard

func TestGenerateCardRegistrationURL(t *testing.T) {
	type args struct {
		ctx   context.Context
		order fincode.GenerateCardRegistrationURLRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "generate card registration url test",
			args: args{
				ctx: context.TODO(),
				order: fincode.GenerateCardRegistrationURLRequest{
					SuccessURL:      fincode.SUCCESS_URL,
					CancelURL:       fincode.CANCEL_URL,
					ShopServiceName: fincode.SHOP_SERVER_NAME,
					CustomerID:      "c_vkEjbdODRYeryK2_P1ZeCw",
					CustomerName:    "Miyabii1210",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret, err := fincode.GenerateCardRegistrationURL(tt.args.ctx, &tt.args.order)
			if err != nil {
				t.Errorf("GenerateCardRegistrationURL error: %v\n", err)
			}
			t.Log("result:", ret)
			fmt.Printf("card registration URL: %s", ret.LinkURL)
			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}
