package nvwa

import (
	"github.com/hootuu/utils/errors"
)

type Service interface {
	VNCreate(req VNCreateReq) (*VNCreateResp, *errors.Error)
	SPCreate(req SPCreateReq) (*SPCreateReq, *errors.Error)
}
