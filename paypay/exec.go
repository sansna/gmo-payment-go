package paypay

import (
	"encoding/json"

	"github.com/sansna/gmo-payment-go/gmopg"
)

type ExecTranPaypayArgs struct {
	ShopID         string `json:"shopID"`
	ShopPass       string `json:"shopPass"`
	AccessID       string `json:"accessID"`
	AccessPass     string `json:"accessPass"`
	OrderID        string `json:"orderID"`
	RetURL         string `json:"retURL"`
	PaymentTermSec int    `json:"paymentTermSec"`
	TransitionType int    `json:"transitionType"`
}

type ExecTranPaypayResult struct {
	Error          gmopg.ErrorResults
	AccessID       string `json:"accessID"`
	Token          string `json:"token"`
	StartURL       string `json:"startUrl"`
	StartLimitDate string `json:"startLimitDate"`
}

func (p Paypay) ExecTranPaypay(args *ExecTranPaypayArgs) (*ExecTranPaypayResult, error) {
	if args.ShopID == "" || args.ShopPass == "" {
		args.ShopID, args.ShopPass = p.GetShopIDPW()
	}
	paramsJSON, _ := json.Marshal(args)
	cli := p.GetClient()
	resp, err := cli.Do("/payment/ExecTranPaypay.json", paramsJSON)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		var ok ExecTranPaypayResult
		if err := json.Unmarshal(resp.Body, &ok); err != nil {
			return nil, err
		}

		return &ok, nil
	}

	ng := &ExecTranPaypayResult{}
	if err = json.Unmarshal(resp.Body, &ng.Error); err != nil {
		return nil, err
	}

	return ng, nil
}
