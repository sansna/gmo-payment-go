package paypay

import (
	"encoding/json"

	"github.com/sansna/gmo-payment-go/gmopg"
)

type EntryTranPaypayArgs struct {
	ShopID   string      `json:"shopID"`
	ShopPass string      `json:"shopPass"`
	OrderID  string      `json:"orderID"`
	JobCd    gmopg.JobCd `json:"jobCd"`
	Amount   int         `json:"amount"`
	Tax      int         `json:"tax"`
}

type EntryTranPaypayResult struct {
	Error      gmopg.ErrorResults
	AccessID   string `json:"accessID"`
	AccessPass string `json:"accessPass"`
}

type Paypay struct {
	gmopg.GMOPG
}

func NewPaypayClient(g gmopg.GMOPG) Paypay {
	return Paypay{
		GMOPG: g,
	}
}

func (p Paypay) EntryTranPaypay(args *EntryTranPaypayArgs) (*EntryTranPaypayResult, error) {
	if args.ShopID == "" || args.ShopPass == "" {
		args.ShopID, args.ShopPass = p.GetShopIDPW()
	}
	paramsJSON, _ := json.Marshal(args)
	cli := p.GetClient()
	resp, err := cli.Do("/payment/EntryTranPaypay.json", paramsJSON)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		var ok EntryTranPaypayResult
		if err := json.Unmarshal(resp.Body, &ok); err != nil {
			return nil, err
		}

		return &ok, nil
	}

	ng := &EntryTranPaypayResult{}
	if err = json.Unmarshal(resp.Body, &ng.Error); err != nil {
		return nil, err
	}

	return ng, nil
}
