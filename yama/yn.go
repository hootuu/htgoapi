package yama

import "github.com/hootuu/tome/yn"

type YinPlantReq struct {
	Yin yn.Yin `json:"yin"`
}

type YinPlantResp struct {
	YID yn.YID `json:"yid"`
}
