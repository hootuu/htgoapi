package yama

import (
	"github.com/hootuu/htgoapi/types/load"
	"github.com/hootuu/tome/sp"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/tome/vn/value"
)

type ValueLoadReq struct {
	VN   vn.ID      `json:"vn"`
	Page *load.Page `json:"page"`
}

type ValueSpItem struct {
	SP    sp.ID       `json:"sp"`
	TTM   int64       `json:"ttm"`
	Value value.Value `json:"value"`
	LstDt string      `json:"lst_dt"`
}

type ValueLoadResp struct {
	TTM    int64          `json:"ttm"`
	Items  []*ValueSpItem `bson:"items" json:"items"`
	Paging *load.Paging   `bson:"paging" json:"paging"`
}
