package negativekeyword

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// BatchDeleteRequest 批量删除否定词 API Request
type BatchDeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UnitID 单元ID
	UnitID uint64 `json:"unit_id,omitempty"`
	// NegativeKeywordIDs 否定词id列表
	NegativeKeywordIDs []uint64 `json:"negative_keyword_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r BatchDeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// BatchDeleteResponse 批量删除否定词 API Response
type BatchDeleteResponse struct {
	model.BaseResponse
}
