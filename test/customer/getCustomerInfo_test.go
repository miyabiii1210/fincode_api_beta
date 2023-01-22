package test

import (
	"context"
	"testing"

	"github.com/miyabiii1210/fincode_api_beta/go/pkg/fincode"
	"github.com/miyabiii1210/fincode_api_beta/go/pkg/util"
)

func TestGetCustomerInfo(t *testing.T) {
	type args struct {
		ctx        context.Context
		customerId string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "get customer information test",
			args: args{
				ctx:        context.TODO(),
				customerId: fincode.CUSTOMER_ID,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret, err := fincode.GetCustomerInfo(tt.args.ctx, tt.args.customerId)
			if err != nil {
				t.Errorf("GetCustomerInfo error : %v\n", err)
			}
			t.Log("result:", ret)
			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}
