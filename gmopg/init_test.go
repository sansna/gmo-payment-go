package gmopg

import "github.com/sansna/gmo-payment-go/httpc"

var (
	gmopg GMOPG
)

func setup(baseURL string) {
	gmopg = Init(&Config{
		BaseURL:  &baseURL,
		SiteID:   "siteid",
		ShopID:   "shopid",
		SitePass: "sitepass",
		ShopPass: "shoppass",
	})

	gmopg.client = &httpc.Client{
		BaseURL: gmopg.baseURL,
	}
}
