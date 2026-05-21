package simpledelivery

import (
	"github.com/bububa/spotlight-mapi/model"
	"github.com/bububa/spotlight-mapi/model/report"
	"github.com/bububa/spotlight-mapi/util"
)

// KeywordRealtimeRequest 简单投关键词层级实时数据 API Request
type KeywordRealtimeRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// CampaignGroupIDList 广告组ID，必传且只能传一个
	CampaignGroupIDList []uint64 `json:"campaign_group_id_list,omitempty"`
	// PageNum 页数，默认1
	PageNum int64 `json:"page_num,omitempty"`
	// PageSize 页大小，默认20，最大100
	PageSize int64 `json:"page_size,omitempty"`
	// StartDate 开始时间，格式 yyyy-MM-dd
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间，格式 yyyy-MM-dd
	EndDate string `json:"end_date,omitempty"`
	// DataCaliber 数据指标归因时间类型，0-计费时间、1-转化时间
	DataCaliber int `json:"data_caliber,omitempty"`
	// KeywordFilterState 关键词状态，见枚举说明
	KeywordFilterState int `json:"keyword_filter_state,omitempty"`
	// PhraseMatchType 关键词匹配类型，0-精准匹配、1-短语匹配
	PhraseMatchType int `json:"phrase_match_type,omitempty"`
	// KeywordName 关键词
	KeywordName string `json:"keyword_name,omitempty"`
	// MatchType 搜索匹配类型，与 keyword_name 配合，0-包含、1-等于
	MatchType int `json:"match_type,omitempty"`
}

// Encode implement PostRequest interface
func (r KeywordRealtimeRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// KeywordRealtimeResponse 简单投关键词层级实时数据 API Response
type KeywordRealtimeResponse struct {
	model.BaseResponse
	// Data 分页数据
	Data *KeywordRealtimeData `json:"data,omitempty"`
}

// KeywordRealtimeData 简单投关键词层级实时数据
type KeywordRealtimeData struct {
	// Page 分页信息
	Page *PageDTO `json:"page,omitempty"`
	// TotalData 总计数据
	TotalData *report.DataReportDTO `json:"total_data,omitempty"`
	// List 当前页数据
	List []KeywordDTO `json:"keyword_dtos,omitempty"`
}

// KeywordDTO 关键词数据
type KeywordDTO struct {
	// BaseKeywordDTO 关键词实体对象
	BaseKeywordDTO *BaseKeywordDTO `json:"base_keyword_dto,omitempty"`
	// Data 数据报表
	Data *report.DataReportDTO `json:"data,omitempty"`
}

// BaseKeywordDTO 关键词实体对象
type BaseKeywordDTO struct {
	// KeywordID 关键词ID
	KeywordID string `json:"keyword_id,omitempty"`
	// Keyword 关键词
	Keyword string `json:"keyword,omitempty"`
	// CampaignGroupID 广告组ID
	CampaignGroupID uint64 `json:"campaign_group_id,omitempty"`
	// CampaignGroupName 标的名称
	CampaignGroupName string `json:"campaign_group_name,omitempty"`
	// CampaignID 计划ID
	CampaignID uint64 `json:"campaign_id,omitempty"`
	// CampaignName 计划名称
	CampaignName string `json:"campaign_name,omitempty"`
	// UnitID 单元ID
	UnitID uint64 `json:"unit_id,omitempty"`
	// UnitName 单元名称
	UnitName string `json:"unit_name,omitempty"`
	// KeywordEnable 关键词上线状态，0-下线、1-上线
	KeywordEnable int `json:"keyword_enable,omitempty"`
	// KeywordFilterState 关键词状态，见枚举说明
	KeywordFilterState int `json:"keyword_filter_state,omitempty"`
	// CreateTime 创建时间，格式 yyyy-MM-dd HH:mm:ss
	CreateTime string `json:"create_time,omitempty"`
	// PhraseMatchType 匹配类型，0-精准匹配、1-短语匹配
	PhraseMatchType int `json:"phrase_match_type,omitempty"`
	// Bid 出价
	Bid int64 `json:"bid,omitempty"`
}
