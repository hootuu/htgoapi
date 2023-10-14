package htapi

import (
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/utils/errors"
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

type VNService interface {
	Create(req VNCreateReq) (*VNCreateResp, *errors.Error)
}
