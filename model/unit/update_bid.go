package unit

import (
	"github.com/bububa/spotlight-mapi/util"
)

// UpdateBidRequest 修改单元出价 API Request
type UpdateBidRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// EventBidList 调价单元列表，最多100个
	EventBidList []UnitBidInfo `json:"event_bid_list,omitempty"`
}

// UnitBidInfo 单元出价信息
type UnitBidInfo struct {
	// UnitID 单元ID
	UnitID uint64 `json:"unit_id,omitempty"`
	// EventBid 出价，单位分
	EventBid int64 `json:"event_bid,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateBidRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
