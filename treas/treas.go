package treas

import "github.com/hootuu/utils/errors"

type Service interface {
	CoinIssue(req CoinIssueReq) (*CoinIssueResp, *errors.Error)
	AccountCreate(req AccountCreateReq) (*AccountCreateResp, *errors.Error)
}
