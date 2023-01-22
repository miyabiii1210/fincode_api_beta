package test

import (
	"context"
	"testing"

	"github.com/miyabiii1210/fincode_api_beta/go/pkg/fincode"
	"github.com/miyabiii1210/fincode_api_beta/go/pkg/util"
)

func TestCustomerRegistration(t *testing.T) {
	type args struct {
		ctx   context.Context
		order fincode.CustomerRegistrationRequest
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test format",
			args: args{
				ctx: context.TODO(),
				order: fincode.CustomerRegistrationRequest{
					ID:    "", // 未指定の場合はfincodeでユニークなIDを設定
					Name:  fincode.CUSTOMER_NAME,
					Email: fincode.CUSTOMER_EMAIL,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret, err := fincode.CustomerRegistration(tt.args.ctx, &tt.args.order)
			if err != nil {
				t.Errorf("CustomerRegistration error: %v\n", err)
				return
			}
			t.Log("result:", ret)
			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}
