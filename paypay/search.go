package paypay

import (
	"net/url"
	"strconv"

	"github.com/go-playground/form"
	"github.com/google/go-querystring/query"
	"github.com/sansna/gmo-payment-go/gmopg"
)

type SearchTradeMultiArgs struct {
	ShopID   string `form:"ShopID"`
	ShopPass string `form:"ShopPass"`
	OrderID  string `form:"OrderID"`
	PayType  int    `form:"PayType"`
}

type ErrorSt struct {
	ErrCode []string `form:"ErrCode"`
	ErrInfo []string `form:"ErrInfo"`
}

type SearchTradeMultiResult struct {
	Error              gmopg.ErrorResults
	Status             string `form:"Status"`
	ProcessDate        string `form:"ProcessDate"`
	JobCd              string `form:"JobCd"`
	AccessID           string `form:"AccessID"`
	AccessPass         string `form:"AccessPass"`
	Amount             string `form:"Amount"`
	Tax                string `form:"Tax"`
	PayType            string `form:"PayType"`
	PayPayCancelAmount string `form:"PayPayCancelAmount"`
	PayPayCancelTax    string `form:"PayPayCancelTax"`
	PayPayTrackingID   string `form:"PayPayTrackingID"`
	PayPayAcceptCode   string `form:"PayPayAcceptCode"`
	PayPayOrderID      string `form:"PayPayOrderID"`
}

func atoi(s string) int {
	i, _ := strconv.ParseInt(s, 10, 64)
	return int(i)
}

// only idPass version
func (p Paypay) SearchTradeMulti(args *SearchTradeMultiArgs) (*SearchTradeMultiResult, error) {
	if args.ShopID == "" || args.ShopPass == "" {
		args.ShopID, args.ShopPass = p.GetShopIDPW()
	}
	args.PayType = atoi(string(gmopg.Paypay))
	v, _ := query.Values(args)
	path := "/payment/SearchTradeMulti.idPass"
	path += "?" + v.Encode()
	cli := p.GetClient()
	resp, err := cli.Do(path, nil)
	if err != nil {
		return nil, err
	}

	vs, err := url.ParseQuery(string(resp.Body))
	if err != nil {
		return nil, err
	}
	d := form.NewDecoder()
	var ok SearchTradeMultiResult
	if resp.StatusCode == 200 {
		if err := d.Decode(&ok, vs); err != nil {
			return nil, err
		}
	}

	var est ErrorSt
	if err = d.Decode(&est, vs); err != nil {
		return nil, err
	}
	for i := 0; i < len(est.ErrCode); i++ {
		ok.Error = append(ok.Error, gmopg.ErrorResult{
			ErrCode: est.ErrCode[i],
			ErrInfo: est.ErrInfo[i],
		})
	}

	return &ok, nil
}
