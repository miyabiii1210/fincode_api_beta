package test

import (
	"context"
	"testing"

	"github.com/miyabiii1210/fincode_api_beta/go/pkg/fincode"
	"github.com/miyabiii1210/fincode_api_beta/go/pkg/util"
)

func TestPaymentConfirm(t *testing.T) {
	type args struct {
		ctx     context.Context
		order   fincode.PaymentConfirmRequest
		orderId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "payment confirm test",
			args: args{
				ctx: context.TODO(),
				order: fincode.PaymentConfirmRequest{
					PayType:  fincode.PAY_TYPE_CARD,
					AccessID: "",
				},
				orderId: "",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret, err := fincode.PaymentConfirm(tt.args.ctx, &tt.args.order, tt.args.orderId)
			if err != nil {
				t.Errorf("PaymentConfirm error: %v", err)
				return
			}
			t.Log("result:", ret)
			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}
