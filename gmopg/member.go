package gmopg

import "encoding/json"

type SaveMemberResult struct {
	Error    ErrorResults
	MemberID string `json:"memberID"`
}

type SaveMemberArgs struct {
	MemberID   string `json:"memberID"`
	MemberName string `json:"memberName"`
}

func (g *GMOPG) SaveMember(args *SaveMemberArgs) (*SaveMemberResult, error) {
	param := struct {
		SaveMemberArgs
		SiteID   string `json:"siteID"`
		SitePass string `json:"sitePass"`
	}{
		SaveMemberArgs: *args,
		SiteID:         g.siteID,
		SitePass:       g.sitePass,
	}

	paramsJSON, _ := json.Marshal(param)
	resp, err := g.client.Do("/payment/SaveMember.json", paramsJSON)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		var ok SaveMemberResult
		if err := json.Unmarshal(resp.Body, &ok); err != nil {
			return nil, err
		}

		return &ok, nil
	}

	ng := &SaveMemberResult{}
	if err = json.Unmarshal(resp.Body, &ng.Error); err != nil {
		return nil, err
	}

	return ng, nil
}
