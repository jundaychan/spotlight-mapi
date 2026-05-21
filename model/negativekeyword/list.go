package negativekeyword

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/util"
)

// ListRequest 查询否定词列表 API Request
type ListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UnitID 单元ID
	UnitID uint64 `json:"unit_id,omitempty"`
	// PageNum 页数，默认1
	PageNum int `json:"page_num,omitempty"`
	// PageSize 页大小，默认20
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement PostRequest interface
func (r ListRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// ListResponse 查询否定词列表 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListResult `json:"data,omitempty"`
}

// ListResult 查询否定词列表结果
type ListResult struct {
	// Page 分页信息
	Page model.PageRespDTO `json:"page,omitempty"`
	// NegativeKeywords 否定词详情列表
	NegativeKeywords []NegativeKeywordItemDTO `json:"negative_keywords,omitempty"`
}

// NegativeKeywordItemDTO 否定词详情
type NegativeKeywordItemDTO struct {
	// CampaignID 计划id
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// UnitID 单元id
	UnitID uint64 `json:"unit_id,omitempty"`
	// NegativeKeywordID 否定词id
	NegativeKeywordID uint64 `json:"negative_keyword_id,omitempty"`
	// Keyword 否定词
	Keyword string `json:"keyword,omitempty"`
	// PhraseMatchType 匹配方式，0-精确匹配，1-短语匹配
	PhraseMatchType int `json:"phrase_match_type,omitempty"`
}
