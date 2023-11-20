package treas

import (
	"github.com/hootuu/tome/uc"
	"github.com/hootuu/utils/errors"
)

type Service interface {
	CoinIssue(req CoinIssueReq) (*CoinIssueResp, *errors.Error)
	CoinGet(req CoinGetReq) (*uc.Coin, *errors.Error)
	AccountCreate(req AccountCreateReq) (*AccountCreateResp, *errors.Error)
	AccountGet(req AccountGetReq) (*uc.Account, *errors.Error)
	AccountLoad(req AccountLoadReq) (*AccountLoadResp, *errors.Error)
	AlterLoad(req AlterLoadReq) (*AlterLoadResp, *errors.Error)

	NftIssue(req NftIssueReq) (*NftIssueResp, *errors.Error)
	NftAirdrop(req NftAirdropReq) (*NftAirdropResp, *errors.Error)
	NftPledge(req NftPledgeReq) (*NftPledgeResp, *errors.Error)
	NftRent(req NftRentReq) (*NftRentResp, *errors.Error)
	NftTrans(req NftTransReq) (*NftTransResp, *errors.Error)
	NftMine(req NftMineReq) (*NftMineResp, *errors.Error)
}
