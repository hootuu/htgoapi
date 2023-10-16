package treas

import (
	"github.com/hootuu/htgoapi/types/load"
	"github.com/hootuu/tome/uc"
	"github.com/hootuu/tome/vn"
)

type CoinIssueReq struct {
	Token    string  `json:"token"`
	Issuer   vn.ID   `json:"issuer"`
	Coin     uc.Code `json:"coin"`
	Wei      uc.WEI  `json:"wei"`
	Issuance int64   `json:"issuance"`
}

type CoinIssueResp struct {
	Coin *uc.Coin `json:"coin"`
}

type CoinGetReq struct {
	By   load.ByCode `json:"by"`
	Coin uc.Code     `json:"coin"`
}
