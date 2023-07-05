package gmopg

import "encoding/json"

type DeleteCardArgs struct {
	MemberID string   `json:"memberID"`
	SeqMode  *SeqMode `json:"seqMode,omitempty"`
	CardSeq  int      `json:"cardSeq"`
}

type DeleteCardResult struct {
	Error   ErrorResults
	CardSeq string `json:"cardSeq"`
}

func (g *GMOPG) DeleteCard(args *DeleteCardArgs) (*DeleteCardResult, error) {
	param := struct {
		DeleteCardArgs
		SiteID   string `json:"siteID"`
		SitePass string `json:"sitePass"`
	}{
		DeleteCardArgs: *args,
		SiteID:         g.siteID,
		SitePass:       g.sitePass,
	}

	paramsJSON, _ := json.Marshal(param)
	resp, err := g.client.Do("/payment/DeleteCard.json", paramsJSON)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		var ok DeleteCardResult
		if err := json.Unmarshal(resp.Body, &ok); err != nil {
			return nil, err
		}

		return &ok, nil
	}

	ng := &DeleteCardResult{}
	if err = json.Unmarshal(resp.Body, &ng.Error); err != nil {
		return nil, err
	}

	return ng, nil
}
