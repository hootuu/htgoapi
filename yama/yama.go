package yama

import "github.com/hootuu/utils/errors"

type Service interface {
	ContractDeploy() *errors.Error
	YinPlant(req YinPlantReq) (*YinPlantResp, *errors.Error)
}
