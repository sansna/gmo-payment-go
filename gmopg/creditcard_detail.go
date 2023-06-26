package gmopg

import "encoding/json"

type SearchCardDetailArgs struct {
	MemberID string   `json:"memberID"`
	SeqMode  *SeqMode `json:"seqMode,omitempty"`
	CardSeq  int      `json:"cardSeq,omitempty"`
}

type SearchCardDetailResult struct {
	Error                  ErrorResults
	CardNo                 string `json:"cardNo"`
	Brand                  string `json:"brand"`
	DomesticFlag           int    `json:"domesticFlag"`
	IssuerCode             string `json:"issuerCode"`
	DebitPrepaidFlag       int    `json:"debitPrepaidFlag"`
	DebitPrepaidIssuerName string `json:"debitPrepaidIssuerName"`
	ForwardFinal           string `json:"forwardFinal"`
}

func (g *GMOPG) SearchCardDetail(args *SearchCardDetailArgs) (*SearchCardDetailResult, error) {
	param := struct {
		SearchCardDetailArgs
		SiteID   string `json:"siteID"`
		SitePass string `json:"sitePass"`
		ShopID   string `json:"shopID"`
		ShopPass string `json:"shopPass"`
	}{
		SearchCardDetailArgs: *args,
		SiteID:               g.siteID,
		SitePass:             g.sitePass,
		ShopID:               g.shopID,
		ShopPass:             g.shopPass,
	}

	paramsJSON, _ := json.Marshal(param)
	resp, err := g.client.Do("/payment/SearchCardDetail.json", paramsJSON)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		var ok SearchCardDetailResult
		if err := json.Unmarshal(resp.Body, &ok); err != nil {
			return nil, err
		}

		return &ok, nil
	}

	ng := &SearchCardDetailResult{}
	if err = json.Unmarshal(resp.Body, &ng.Error); err != nil {
		return nil, err
	}

	return ng, nil
}
