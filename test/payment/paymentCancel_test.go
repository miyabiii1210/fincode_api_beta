package test

import (
	"context"
	"testing"

	"github.com/miyabiii1210/fincode_api_beta/go/pkg/fincode"
	"github.com/miyabiii1210/fincode_api_beta/go/pkg/util"
)

func TestPaymentCancel(t *testing.T) {
	type args struct {
		ctx     context.Context
		order   fincode.PaymentCancelRequest
		orderId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "payment cancel test",
			args: args{
				ctx: context.TODO(),
				order: fincode.PaymentCancelRequest{
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
			ret, err := fincode.PaymentCancel(tt.args.ctx, &tt.args.order, tt.args.orderId)
			if err != nil {
				t.Errorf("PaymentCancel error %v\n", err)
			}
			t.Log("result:", ret)
			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}
