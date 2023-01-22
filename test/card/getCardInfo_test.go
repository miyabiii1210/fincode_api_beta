package test

import (
	"context"
	"testing"

	"github.com/miyabiii1210/fincode_api_beta/go/pkg/fincode"
	"github.com/miyabiii1210/fincode_api_beta/go/pkg/util"
)

func TestGetCardInfo(t *testing.T) {
	type args struct {
		ctx        context.Context
		customerId string
		cardId     string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "get card information test",
			args: args{
				ctx:        context.TODO(),
				customerId: fincode.CUSTOMER_ID,
				cardId:     fincode.CARD_ID,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ret, err := fincode.GetCardInfo(tt.args.ctx, tt.args.customerId, tt.args.cardId)
			if err != nil {
				t.Errorf("GetCardInfo error: %v\n", err)
			}
			t.Log("result:", ret)
			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}
