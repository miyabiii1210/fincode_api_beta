package test

import (
	"context"
	"fmt"
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
			ret, err := fincode.GetPaymentInfo(tt.args.ctx, &tt.args.order, tt.args.orderId)
			if err != nil {
				t.Errorf("GetPaymentInfo error: %v\n", err)
				// fmt.Println("ErrorCode      :", ret.ErrorCode)
				return
			}
			t.Log("result:", ret)
			fmt.Println("==============")
			fmt.Println("ID             :", ret.ID)
			fmt.Println("Status         :", ret.Status)
			fmt.Println("AccessID       :", ret.AccessID)
			fmt.Println("TransactionID  :", ret.TransactionID)
			fmt.Println("ShopID         :", ret.ShopID)
			fmt.Println("CustomerID     :", ret.CustomerID)
			fmt.Println("CardID         :", ret.CardID)
			fmt.Println("CardNo         :", ret.CardNo)
			fmt.Println("CardNoHash     :", ret.CardNoHash)
			fmt.Println("AuthMaxDate    :", ret.AuthMaxDate)
			fmt.Println("Approve        :", ret.Approve)
			fmt.Println("JobCode        :", ret.JobCode)
			fmt.Println("Amount         :", ret.Amount)
			fmt.Println("TotalAmount    :", ret.TotalAmount)
			fmt.Println("Method         :", ret.Method)
			fmt.Println("PayType        :", ret.PayType)
			fmt.Println("Created        :", ret.Created)
			fmt.Println("Updated        :", ret.Updated)
			fmt.Println("ProcessDate    :", ret.ProcessDate)
			fmt.Println("==============")
			util.Sleep(1)
			t.Logf("%s fin.\n", tt.name)
		})
	}
}
