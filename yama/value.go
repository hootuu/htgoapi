package yama

import (
	"github.com/hootuu/htgoapi/types/load"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/tome/vn/value"
)

type ValueLoadReq struct {
	VN   vn.ID      `json:"vn"`
	Page *load.Page `json:"page"`
}

type ValueLoadResp struct {
	Paging *load.Paging   `bson:"paging" json:"paging"`
	Items  []*value.Value `bson:"items" json:"items"`
}
