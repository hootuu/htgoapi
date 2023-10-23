package treas

import (
	"github.com/hootuu/htgoapi/types/load"
	"github.com/hootuu/tome/uc"
)

type AccountCreateReq struct {
}

type AccountCreateResp struct {
}

type AccountGetReq struct {
	By   load.ByCode     `json:"by"`
	Lead *uc.AccountLead `json:"lead"`
}

type AccountLoadReq struct {
	Owners []uc.AccountOwner `json:"owners"`
	Page   *load.Page        `json:"page"`
}

type AccountLoadResp struct {
	Paging *load.Paging  `bson:"paging" json:"paging"`
	Items  []*uc.Account `bson:"items" json:"items"`
}
