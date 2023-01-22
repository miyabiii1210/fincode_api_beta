package test

import (
	"context"
	"testing"

	"github.com/miyabiii1210/fincode_api_beta/go/pkg/fincode"
	"github.com/miyabiii1210/fincode_api_beta/go/pkg/util"
)

func TestPaymentExecution(t *testing.T) {
	type args struct {
		ctx     context.Context
		order   fincode.PaymentExecutionRequest
		orderId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "payment execution test",
			args: args{
				ctx: context.TODO(),
				order: fincode.PaymentExecutionRequest{
					PayType:    fincode.PAY_TYPE_CARD,
					AccessID:   "",
					CustomerID: "",
					CardID:     "",
					Method:     "",
				},
				orderId: "",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret, err := fincode.PaymentExecution(tt.args.ctx, &tt.args.order, tt.args.orderId)
			if err != nil {
				t.Errorf("PaymentExecution error: %v\n", err)
			}
			t.Log("result:", ret)
			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}
