package gmopg

import "encoding/json"

type SearchCardArgs struct {
	MemberID  string   `json:"memberID"`
	SeqMode   *SeqMode `json:"seqMode,omitempty"`
	ValidFlag int      `json:"validFlag,omitempty"`
	CardSeq   int      `json:"cardSeq"`
}

type SearchCardResult struct {
	Error                  ErrorResults
	CardSeq                int    `json:"cardSeq"`
	DefaultFlag            int    `json:"defaultFlag"`
	CardName               string `json:"cardName"`
	CardNo                 string `json:"cardNo"`
	Expire                 string `json:"expire"`
	HolderName             string `json:"holderName"`
	DeleteFlag             int    `json:"deleteFlag"`
	Brand                  string `json:"brand"`
	DomesticFlag           int    `json:"domesticFlag"`
	IssuerCode             string `json:"issuerCode"`
	DebitPrepaidFlag       int    `json:"debitPrepaidFlag"`
	DebitPrepaidIssuerName string `json:"debitPrepaidIssuerName"`
	ForwardFinal           string `json:"forwardFinal"`
}

func (g *GMOPG) SearchCard(args *SearchCardArgs) (*SearchCardResult, error) {
	param := struct {
		SearchCardArgs
		SiteID   string `json:"siteID"`
		SitePass string `json:"sitePass"`
	}{
		SearchCardArgs: *args,
		SiteID:         g.siteID,
		SitePass:       g.sitePass,
	}

	paramsJSON, _ := json.Marshal(param)
	resp, err := g.client.Do("/payment/SearchCard.json", paramsJSON)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		var ok SearchCardResult
		if err := json.Unmarshal(resp.Body, &ok); err != nil {
			return nil, err
		}

		return &ok, nil
	}

	ng := &SearchCardResult{}
	if err = json.Unmarshal(resp.Body, &ng.Error); err != nil {
		return nil, err
	}

	return ng, nil
}
