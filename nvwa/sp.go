package nvwa

import (
	"github.com/hootuu/htgoapi/types/load"
	"github.com/hootuu/tome/fq"
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/tome/sp"
	"github.com/hootuu/tome/vn"
)

type SPCreateReq struct {
	Token string   `json:"token"`
	VN    string   `json:"vn"`
	Scope string   `json:"scope"`
	ID    string   `json:"id"`
	FQ    []*fq.FQ `json:"fq"`
}

type SPCreateResp struct {
	VN         vn.ID   `json:"vn"`
	ID         sp.ID   `json:"id"`
	Link       []sp.ID `json:"link"`
	Originator ki.Ki   `json:"originator"`
	Guardian   ki.Ki   `json:"guardian"`
	Dob        int64   `json:"dob"`
}

type SPGetReq struct {
	By load.ByCode `json:"by"`
	ID sp.ID       `json:"id"`
}

type SPLoadReq struct {
	IDs  []sp.ID    `json:"ids"`
	VNs  []vn.ID    `json:"vns"`
	Page *load.Page `json:"page"`
}

type SPLoadResp struct {
	Paging *load.Paging `bson:"paging" json:"paging"`
	Items  []*sp.SP     `bson:"items" json:"items"`
}
