package nvwa

import (
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/tome/sp"
	"github.com/hootuu/tome/vn"
)

type SPCreateReq struct {
	Token string `json:"token"`
	VN    string `json:"vn"`
	Scope string `json:"scope"`
	ID    string `json:"id"`
}

type SPCreateResp struct {
	VN         vn.ID   `json:"vn"`
	ID         sp.ID   `json:"id"`
	Link       []sp.ID `json:"link"`
	Originator ki.Ki   `json:"originator"`
	Guardian   ki.Ki   `json:"guardian"`
	Dob        int64   `json:"dob"`
}
