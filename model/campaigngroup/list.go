package campaigngroup

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// ListRequest 查询广告组列表 API Request
type ListRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// IDs 广告组ID列表，name传值后ids不可传值
	IDs []uint64 `json:"ids,omitempty"`
	// Name 广告组名称，ids传值后name不可传值
	Name string `json:"name,omitempty"`
	// PageNum 页码
	PageNum int `json:"page_num,omitempty"`
	// PageSize 页大小
	PageSize int `json:"page_size,omitempty"`
}

// Encode implement PostRequest interface
func (r ListRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// ListResponse 查询广告组列表 API Response
type ListResponse struct {
	model.BaseResponse
	Data *ListResult `json:"data,omitempty"`
}

// ListResult 查询广告组列表结果
type ListResult struct {
	// PageNum 页码
	PageNum int `json:"page_num,omitempty"`
	// PageSize 页大小
	PageSize int `json:"page_size,omitempty"`
	// TotalPage 总页数
	TotalPage int `json:"total_page,omitempty"`
	// TotalCount 总记录数
	TotalCount int `json:"total_count,omitempty"`
	// DataList 广告组列表
	DataList []GroupDTO `json:"data_list,omitempty"`
}

// GroupDTO 广告组
type GroupDTO struct {
	// GroupParadigm 广告组信息
	GroupParadigm *GroupParadigm `json:"group_paradigm,omitempty"`
}

// GroupParadigm 广告组详情
type GroupParadigm struct {
	// CampaignGroupID 广告组ID
	CampaignGroupID uint64 `json:"campaign_group_id,omitempty"`
	// Enable 启停状态，0-未开启、1-开启
	Enable int `json:"enable,omitempty"`
	// CampaignGroupName 广告组名称
	CampaignGroupName string `json:"campaign_group_name,omitempty"`
	// LimitDayBudget 是否限制日预算，0-不限预算、1-限制预算
	LimitDayBudget int `json:"limit_day_budget,omitempty"`
	// OriginGroupDayBudget 广告组日预算，单位分
	OriginGroupDayBudget int64 `json:"origin_group_day_budget,omitempty"`
	// CreateTime 创建时间，yyyy-MM-dd HH:mm:ss
	CreateTime string `json:"create_time,omitempty"`
}
