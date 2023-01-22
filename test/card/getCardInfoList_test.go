package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/miyabiii1210/fincode_api_beta/go/pkg/fincode"
	"github.com/miyabiii1210/fincode_api_beta/go/pkg/util"
)

func TestGetCardInfoList(t *testing.T) {
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
			name: "get card information list format",
			args: args{
				ctx:        context.TODO(),
				customerId: fincode.CUSTOMER_ID,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret, err := fincode.GetCardInfoList(tt.args.ctx, tt.args.customerId)
			if err != nil {
				t.Errorf("GetCardInfoList error: %v\n", err)
			}
			// t.Log("result:", ret)
			for i, v := range ret.List {
				fmt.Printf("[%d] %s\n", i, v)
			}
			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}
