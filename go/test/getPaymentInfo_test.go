package test

import (
	"context"
	"testing"

	"github.com/miyabiii1210/fincode_api_beta/go/pkg/fincode"
	"github.com/miyabiii1210/fincode_api_beta/go/pkg/util"
)

func TestGetPaymentInfo(t *testing.T) {
	type args struct {
		ctx     context.Context
		order   fincode.GetPaymentInfoRequest
		orderId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "get payment information test",
			args: args{
				ctx: context.TODO(),
				order: fincode.GetPaymentInfoRequest{
					PayType: fincode.PAY_TYPE_CARD,
				},
				orderId: "",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret, err := fincode.GetPaymentInfo(tt.args.ctx, tt.args.order, tt.args.orderId)
			if err != nil {
				t.Errorf("GetPaymentInfo error: %v\n", err)
			}
			t.Log("result:", ret)
			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}
