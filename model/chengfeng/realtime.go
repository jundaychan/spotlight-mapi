// Package chengfeng 乘风(wind)开放平台数据报表模型。
//
// 乘风与聚光共用同一套 OAuth + 同一个 API 网关(adapi.xiaohongshu.com/api/open)，
// 仅报表网关前缀不同：聚光 /jg/...，乘风 /wind/...。
//
// 离线报表的请求/响应结构与聚光完全一致(data_list + DataReportDTO + Dimension)，
// 故 api/chengfeng 的离线接口直接复用 model/report/offline，本文件只补充
// 乘风实时报表的容错模型(乘风文档中创意实时的 base dto 拼写/列表 key 与聚光略有出入)。
package chengfeng

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/model/report"
	"github.com/jundaychan/spotlight-mapi/util"
)

// RealtimeRequest 乘风实时报表请求(账户/计划/创意层级共用)。
type RealtimeRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id"`
	// StartDate 开始时间，格式 yyyy-MM-dd
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间，格式 yyyy-MM-dd
	EndDate string `json:"end_date,omitempty"`
	// Columns 数据指标展示列，默认只返回基础指标
	Columns []string `json:"columns,omitempty"`
	// PageNum 页数，默认1
	PageNum int64 `json:"page_num,omitempty"`
	// PageSize 页大小，默认20,最大1000
	PageSize int64 `json:"page_size,omitempty"`
	// DeliveryModeList 投放模式过滤(仅计划层级实时支持)：自定义:0 / 托管(简单投):1
	DeliveryModeList []int `json:"delivery_mode_list,omitempty"`
}

// Encode implement PostRequest interface
func (r RealtimeRequest) Encode() []byte { return util.JSONMarshal(r) }

// RealtimeAccountResponse 账户层级实时数据：data 为单条汇总 DataReportDTO。
type RealtimeAccountResponse struct {
	model.BaseResponse
	Data *report.DataReportDTO `json:"data,omitempty"`
}

// RealtimeListResponse 计划/创意层级实时数据：data 含实体列表 + 汇总。
type RealtimeListResponse struct {
	model.BaseResponse
	Data *RealtimeList `json:"data,omitempty"`
}

// RealtimeList 实时报表列表数据。乘风文档中创意层级亦可能复用 campaign_dtos key，
// 这里两个 key 都接，由 Entities() 取非空者。
type RealtimeList struct {
	CampaignDTOs   []RealtimeEntityDTO   `json:"campaign_dtos,omitempty"`
	CreativityDTOs []RealtimeEntityDTO   `json:"creativity_dtos,omitempty"`
	TotalData      *report.DataReportDTO `json:"total_data,omitempty"`
	TotalCount     int64                 `json:"total_count,omitempty"`
}

// Entities 返回非空的实体列表(优先 creativity，其次 campaign)。
func (l *RealtimeList) Entities() []RealtimeEntityDTO {
	if l == nil {
		return nil
	}
	if len(l.CreativityDTOs) > 0 {
		return l.CreativityDTOs
	}
	return l.CampaignDTOs
}

// RealtimeEntityDTO 单个实体的实时数据 + 属性。
type RealtimeEntityDTO struct {
	Data *report.DataReportDTO `json:"data,omitempty"`
	// 乘风文档拼写为 base_creativity_dto，聚光 SDK 历史拼写为 base_creativty_dto，两者都接。
	BaseCampaignDTO   *RealtimeBaseDTO `json:"base_campaign_dto,omitempty"`
	BaseCreativityDTO *RealtimeBaseDTO `json:"base_creativity_dto,omitempty"`
	BaseCreativtyDTO  *RealtimeBaseDTO `json:"base_creativty_dto,omitempty"`
}

// Base 返回非空的属性 DTO(创意优先，否则计划)。
func (d *RealtimeEntityDTO) Base() *RealtimeBaseDTO {
	switch {
	case d.BaseCreativityDTO != nil:
		return d.BaseCreativityDTO
	case d.BaseCreativtyDTO != nil:
		return d.BaseCreativtyDTO
	default:
		return d.BaseCampaignDTO
	}
}

// RealtimeBaseDTO 实体属性(只取报表缓存需要的 id/name；宽松解析数字/字符串)。
type RealtimeBaseDTO struct {
	CampaignID     model.Uint64 `json:"campaign_id,omitempty"`
	CampaignName   string       `json:"campaign_name,omitempty"`
	CreativityID   model.Uint64 `json:"creativity_id,omitempty"`
	CreativityName string       `json:"creativity_name,omitempty"`
}
