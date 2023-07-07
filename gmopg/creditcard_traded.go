package gmopg

import "encoding/json"

type TradedCardArgs struct {
	OrderID  string   `json:"orderID"`
	MemberID string   `json:"memberID"`
	SeqMode  *SeqMode `json:"seqMode,omitempty"`

	CardSeq     int          `json:"cardSeq,omitempty"`
	DefaultFlag *DefaultFlag `json:"defaultFlag,omitempty"`
	CardName    string       `json:"cardName,omitempty"`
	HolderName  string       `json:"holderName"`
	CardPass    string       `json:"cardPass"`
}

type TradedCardResult struct {
	Error   ErrorResults
	CardSeq string `json:"cardSeq"`
	CardNo  string `json:"cardNo"`
	Forward string `json:"forward"`
}

func (g *GMOPG) TradedCard(args *TradedCardArgs) (*TradedCardResult, error) {
	param := struct {
		TradedCardArgs
		SiteID   string `json:"siteID"`
		SitePass string `json:"sitePass"`
		ShopID   string `json:"shopID"`
		ShopPass string `json:"shopPass"`
	}{
		TradedCardArgs: *args,
		SiteID:         g.siteID,
		SitePass:       g.sitePass,
		ShopID:         g.shopID,
		ShopPass:       g.shopPass,
	}

	paramsJSON, _ := json.Marshal(param)
	resp, err := g.client.Do("/payment/TradedCard.json", paramsJSON)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		var ok TradedCardResult
		if err := json.Unmarshal(resp.Body, &ok); err != nil {
			return nil, err
		}

		return &ok, nil
	}

	ng := &TradedCardResult{}
	if err = json.Unmarshal(resp.Body, &ng.Error); err != nil {
		return nil, err
	}

	return ng, nil
}
