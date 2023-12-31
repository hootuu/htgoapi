package nvwa

import (
	"github.com/hootuu/tome/sp"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/utils/errors"
)

type Service interface {
	VNCreate(req VNCreateReq) (*VNCreateResp, *errors.Error)
	SPCreate(req SPCreateReq) (*SPCreateResp, *errors.Error)

	VNGet(req VNGetReq) (*vn.VN, *errors.Error)
	VNLoad(req VNLoadReq) (*VNLoadResp, *errors.Error)
	SPGet(req SPGetReq) (*sp.SP, *errors.Error)
	SPLoad(req SPLoadReq) (*SPLoadResp, *errors.Error)
}
