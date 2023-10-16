package nvwa

import (
	"github.com/hootuu/htgoapi/types/load"
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/tome/vn"
)

type VNCreateReq struct {
	Token string `json:"token"`
	ID    string `json:"id"`
}

type VNCreateResp struct {
	ID         string `json:"id"`
	Originator ki.Ki  `json:"originator"`
	Guardian   ki.Ki  `json:"guardian"`
	Dob        int64  `json:"dob"`
}

type VNGetReq struct {
	By load.ByCode `json:"by"`
	ID vn.ID       `json:"id"`
}

type VNLoadReq struct {
	IDs  []vn.ID    `json:"ids"`
	Page *load.Page `json:"page"`
}
