package gmopg

import "encoding/json"

type SecureTran2Args struct {
	AccessID   string `json:"accessID"`
	AccessPass string `json:"accessPass"`
}

type SecureTran2Result struct {
	Error       ErrorResults
	OrderID     string `json:"orderID"`
	Forward     string `json:"forward"`
	Method      int    `json:"method"`
	PayTimes    string `json:"payTimes"`
	Approve     string `json:"approve"`
	TranID      string `json:"tranID"`
	TranDate    string `json:"tranDate"`
	CheckString string `json:"checkString"`
}

func (g *GMOPG) SecureTran2(args *SecureTran2Args) (*SecureTran2Result, error) {
	param := struct {
		SecureTran2Args
		//SiteID   string `json:"siteID"`
		//SitePass string `json:"sitePass"`
	}{
		SecureTran2Args: *args,
		//SiteID:          g.siteID,
		//SitePass:        g.sitePass,
	}

	paramsJSON, _ := json.Marshal(param)
	resp, err := g.client.Do("/payment/SecureTran2.json", paramsJSON)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		var ok SecureTran2Result
		if err := json.Unmarshal(resp.Body, &ok); err != nil {
			return nil, err
		}

		return &ok, nil
	}

	ng := &SecureTran2Result{}
	if err = json.Unmarshal(resp.Body, &ng.Error); err != nil {
		return nil, err
	}

	return ng, nil
}
