package test

import (
	"context"
	"testing"

	"github.com/miyabiii1210/fincode_api_beta/go/pkg/fincode"
	"github.com/miyabiii1210/fincode_api_beta/go/pkg/util"
)

func TestPaymentRegistration(t *testing.T) {
	type args struct {
		ctx   context.Context
		order fincode.PaymentRegistrationRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "payment registration test",
			args: args{
				ctx: context.TODO(),
				order: fincode.PaymentRegistrationRequest{
					PayType: fincode.PAY_TYPE_CARD,
					JobCode: fincode.JOB_CODE_CAPTURE,
					Amount:  "2000",
					Tax:     "0",
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret, err := fincode.PaymentRegistration(tt.args.ctx, &tt.args.order)
			if err != nil {
				t.Errorf("PaymentRegistration error: %v\n", err)
			}
			t.Log("result:", ret)
			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}
