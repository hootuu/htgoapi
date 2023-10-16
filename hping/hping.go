package hping

import (
	"github.com/hootuu/htgoapi/types/load"
	"github.com/hootuu/tome/sp"
	"github.com/hootuu/tome/vn"
	"github.com/hootuu/utils/errors"
)

type PingReq struct {
	VN     vn.ID  `json:"vn"`
	SP     sp.ID  `json:"sp"`
	PeerID string `json:"peer_id"`
}

type PingResp struct {
	Timestamp int64 `json:"timestamp"`
}

type NodeLoadReq struct {
	VN   vn.ID      `json:"vn"`
	Page *load.Page `json:"page"`
}

type NodeInfo struct {
	VN     vn.ID  `json:"vn"`
	SP     sp.ID  `json:"sp"`
	PeerID string `json:"peer_id"`
	IP     string `json:"ip"`
	LST    int64  `json:"lst"`
}

type Service interface {
	Ping(req PingReq) (*PingResp, *errors.Error)
	NPLoad(req NodeLoadReq) ([]*NodeInfo, *load.Paging, *errors.Error)
}
