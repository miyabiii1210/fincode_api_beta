package test

import (
	"context"
	"testing"

	"github.com/miyabiii1210/fincode_api_beta/go/pkg/fincode"
	"github.com/miyabiii1210/fincode_api_beta/go/pkg/util"
)

func TestDeleteCustomerInfo(t *testing.T) {
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
			name: "delete customer information test",
			args: args{
				ctx:        context.TODO(),
				customerId: "c_OoYczaFZRaS4hSNvdNTA-Q",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret, err := fincode.DeleteCustomerInfo(tt.args.ctx, tt.args.customerId)
			if err != nil {
				t.Errorf("DeleteCustomerInfo error: %v\n", err)
			}
			t.Log("result:", ret)
			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}
