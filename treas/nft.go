package treas

import (
	"github.com/hootuu/htgoapi/types/load"
	"github.com/hootuu/tome/ki"
	"github.com/hootuu/tome/uc/nft"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/utils/errors"
)

type NftIssueReq struct {
	VN          vn.ID        `bson:"vn" json:"vn"`
	Category    nft.Category `bson:"category" json:"category"`
	Token       nft.Token    `bson:"token" json:"token"`
	Tag         []nft.Tag    `bson:"tag" json:"tag"`
	Link        nft.Link     `bson:"link" json:"link,omitempty"`
	Title       string       `bson:"title" json:"title,omitempty"`
	Description string       `bson:"desc" json:"desc,omitempty"`
	Meta        nft.Meta     `bson:"meta" json:"meta,omitempty"`
}

func (req *NftIssueReq) Verify() *errors.Error {
	if err := vn.IDVerify(req.VN.S()); err != nil {
		return err
	}
	if err := nft.CategoryVerify(req.Category.S()); err != nil {
		return err
	}
	if err := nft.TokenVerify(req.Token.S()); err != nil {
		return err
	}
	if err := nft.TagArrVerify(req.Tag); err != nil {
		return err
	}
	return nil
}

func (req *NftIssueReq) ToNFT() (*nft.NFT, *errors.Error) {
	nftM := &nft.NFT{
		VN:          req.VN,
		Category:    req.Category,
		Token:       req.Token,
		Tag:         req.Tag,
		Link:        req.Link,
		Title:       req.Title,
		Description: req.Description,
		Meta:        req.Meta,
	}
	if err := nftM.Verify(); err != nil {
		return nil, err
	}
	return nftM, nil
}

type NftIssueResp struct {
	NID nft.NID `bson:"nid" json:"nid"`
}

type NftAirdropReq struct {
	NID   nft.NID `bson:"nid" json:"nid"`
	Owner ki.ADR  `bson:"owner" json:"owner"`
}

func (req *NftAirdropReq) Verify() *errors.Error {
	if len(req.NID) == 0 {
		return errors.Verify("require nid")
	}
	if err := ki.ADRVerify(req.Owner.S()); err != nil {
		return errors.Verify("invalid owner: " + err.Error())
	}
	return nil
}

type NftAirdropResp struct {
	Tx string `bson:"tx" json:"tx"`
}

type NftTransReq struct {
	NID       nft.NID `bson:"nid" json:"nid"`
	Owner     ki.ADR  `bson:"owner" json:"owner"`
	Receiver  ki.ADR  `bson:"receiver" json:"receiver"`
	Signature string  `bson:"signature" json:"signature"`
}

type NftTransResp struct {
	Tx string `bson:"tx" json:"tx"`
}

type NftPledgeReq struct {
	NID       nft.NID `bson:"nid" json:"nid"`
	Owner     ki.ADR  `bson:"owner" json:"owner"`
	Receiver  ki.ADR  `bson:"receiver" json:"receiver"`
	Signature string  `bson:"signature" json:"signature"`
}

func (req *NftPledgeReq) Verify() *errors.Error {
	if len(req.NID) == 0 {
		return errors.Verify("require nid")
	}
	if err := ki.ADRVerify(req.Owner.S()); err != nil {
		return errors.Verify("invalid owner: " + err.Error())
	}
	if err := ki.ADRVerify(req.Receiver.S()); err != nil {
		return errors.Verify("invalid receiver: " + err.Error())
	}
	if len(req.Signature) == 0 {
		return errors.Verify("require signature")
	}
	return nil
}

type NftPledgeResp struct {
	Tx string `bson:"tx" json:"tx"`
}

type NftRentReq struct {
	NID       nft.NID `bson:"nid" json:"nid"`
	Owner     ki.ADR  `bson:"owner" json:"owner"`
	Receiver  ki.ADR  `bson:"receiver" json:"receiver"`
	Signature string  `bson:"signature" json:"signature"`
}

func (req *NftRentReq) Verify() *errors.Error {
	if len(req.NID) == 0 {
		return errors.Verify("require nid")
	}
	if err := ki.ADRVerify(req.Owner.S()); err != nil {
		return errors.Verify("invalid owner: " + err.Error())
	}
	if err := ki.ADRVerify(req.Receiver.S()); err != nil {
		return errors.Verify("invalid receiver: " + err.Error())
	}
	if len(req.Signature) == 0 {
		return errors.Verify("require signature")
	}
	return nil
}

type NftRentResp struct {
	Tx string `bson:"tx" json:"tx"`
}

const (
	NftByOwner load.ByCode = "owner"
	NftByUser  load.ByCode = "user"
)

type NftMineReq struct {
	By    load.ByCode `bson:"by" json:"by"`
	Owner ki.ADR      `bson:"owner" json:"owner"`
	User  ki.ADR      `bson:"user" json:"user"`
	Page  *load.Page  `bson:"page" json:"page"`
}

func (req *NftMineReq) Verify() *errors.Error {
	if len(req.By) == 0 {
		req.By = "owner"
	}
	switch req.By {
	case NftByOwner:
		if err := ki.ADRVerify(req.Owner.S()); err != nil {
			return errors.Verify("invalid owner: " + err.Error())
		}
	case NftByUser:
		if err := ki.ADRVerify(req.User.S()); err != nil {
			return errors.Verify("invalid user: " + err.Error())
		}
	default:
		return errors.Verify("invalid by")
	}
	return nil
}

type NftMineResp struct {
	Paging *load.Paging `bson:"paging" json:"paging"`
	Items  []*nft.Wrap  `bson:"items" json:"items"`
}
