package test

import (
	"context"
	"fmt"
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
					// JobCode: fincode.JOB_CODE_CAPTURE, // 売上確定
					JobCode: fincode.JOB_CODE_AUTH, // 仮売上
					Amount:  "500",
					Tax:     "0",
					TdsType: "0",
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
