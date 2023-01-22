package test

import (
	"context"
	"testing"

	"github.com/miyabiii1210/fincode_api_beta/go/pkg/fincode"
	"github.com/miyabiii1210/fincode_api_beta/go/pkg/util"
)

func TestUpdateCustomerInfo(t *testing.T) {
	type args struct {
		ctx        context.Context
		order      fincode.UpdateCustomerInfoRequest
		customerId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "update customer information test",
			args: args{
				ctx: context.TODO(),
				order: fincode.UpdateCustomerInfoRequest{
					Name:  fincode.CUSTOMER_NAME,
					Email: fincode.CUSTOMER_EMAIL,
				},
				customerId: fincode.CUSTOMER_ID,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret, err := fincode.UpdateCustomerInfo(tt.args.ctx, &tt.args.order, tt.args.customerId)
			if err != nil {
				t.Errorf("UpdateCustomerInfo error: %v\n", err)
			}
			t.Log("result:", ret)
			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}
