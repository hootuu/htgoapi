package treas

import (
	"github.com/hootuu/htgoapi/types/load"
	"github.com/hootuu/tome/uc"
	"github.com/hootuu/utils/errors"
)

type Service interface {
	CoinIssue(req CoinIssueReq) (*CoinIssueResp, *errors.Error)
	AccountCreate(req AccountCreateReq) (*AccountCreateResp, *errors.Error)
	AccountGet(req AccountGetReq) (*uc.Account, *errors.Error)
	AccountLoad(req AccountLoadReq) ([]*uc.Account, *load.Paging, *errors.Error)
}
