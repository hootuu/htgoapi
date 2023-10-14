package nvwa

import "github.com/hootuu/tome/ki"

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
