package gmopg

import "encoding/json"

type SaveCardArgs struct {
	MemberID    string       `json:"memberID"`
	SeqMode     *SeqMode     `json:"seqMode,omitempty"`
	CardSeq     int          `json:"cardSeq,omitempty"`
	DefaultFlag *DefaultFlag `json:"defaultFlag,omitempty"`
	CardName    string       `json:"cardName,omitempty"`
	UpdateType  int          `json:"updateType,omitempty"`
	Token       string       `json:"token"`
}

type SaveCardResult struct {
	Error                  ErrorResults
	CardSeq                string `json:"cardSeq"`
	CardNo                 string `json:"cardNo"`
	Forward                string `json:"forward"`
	Brand                  string `json:"brand"`
	DomesticFlag           int    `json:"domesticFlag"`
	IssuerCode             string `json:"issuerCode"`
	DebitPrepaidFlag       int    `json:"debitPrepaidFlag"`
	DebitPrepaidIssuerName string `json:"debitPrepaidIssuerName"`
	ForwardFinal           string `json:"forwardFinal"`
}

func (g *GMOPG) SaveCard(args *SaveCardArgs) (*SaveCardResult, error) {
	param := struct {
		SaveCardArgs
		SiteID   string `json:"siteID"`
		SitePass string `json:"sitePass"`
	}{
		SaveCardArgs: *args,
		SiteID:       g.siteID,
		SitePass:     g.sitePass,
	}

	paramsJSON, _ := json.Marshal(param)
	resp, err := g.client.Do("/payment/SaveCard.json", paramsJSON)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		var ok SaveCardResult
		if err := json.Unmarshal(resp.Body, &ok); err != nil {
			return nil, err
		}

		return &ok, nil
	}

	ng := &SaveCardResult{}
	if err = json.Unmarshal(resp.Body, &ng.Error); err != nil {
		return nil, err
	}

	return ng, nil
}
