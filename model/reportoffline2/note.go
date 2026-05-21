package reportoffline2

import (
	"github.com/jundaychan/spotlight-mapi/enum"
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/model/report"
	"github.com/jundaychan/spotlight-mapi/util"
)

// NoteFilterClause 笔记离线报表过滤条件
type NoteFilterClause struct {
	// Column 筛选列; noteId/fee/impression/click/interaction
	Column string `json:"column,omitempty"`
	// Operator 操作符; >：大于 <：小于 in：等于
	Operator string `json:"operator,omitempty"`
	// Values 值
	Values []string `json:"values,omitempty"`
}

// NoteRequest 笔记层级离线报表数据 API Request
type NoteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id"`
	// StartDate 开始时间，格式 yyyy-MM-dd
	StartDate string `json:"start_date,omitempty"`
	// EndDate 结束时间，格式 yyyy-MM-dd
	EndDate string `json:"end_date,omitempty"`
	// TimeUnit 时间维度："DAY"：分天 "HOUR"：分时 "SUMMARY"：汇总,默认分天
	TimeUnit enum.TimeUnit `json:"time_unit,omitempty"`
	// MarketingTarget 营销目标过滤条件
	MarketingTarget []int `json:"marketing_target,omitempty"`
	// BiddingStrategy 出价方式过滤条件：2：手动出价 3：自动出价 4：MCB 7：OCPX
	BiddingStrategy []int `json:"bidding_strategy,omitempty"`
	// OptimizeTarget 推广目标过滤条件
	OptimizeTarget []int `json:"optimize_target,omitempty"`
	// Placement 广告类型过滤条件 1：信息流 2：搜索 4：全站智投 7：视频流
	Placement []int `json:"placement,omitempty"`
	// PromotionTarget 推广标的类型过滤条件 1：笔记 2：商品 7：外链落地页 9：落地页 18：直播间 19：不限
	PromotionTarget []int `json:"promotion_target,omitempty"`
	// Programmatic 创意组合方式过滤条件 0：自定义创意 1：程序化创意
	Programmatic []int `json:"programmatic,omitempty"`
	// DeliveryMode 投放模式过滤条件 0：手动投放 1：自动投放
	DeliveryMode []int `json:"delivery_mode,omitempty"`
	// SplitColumns 细分条件(相当于group by)
	SplitColumns []string `json:"split_columns,omitempty"`
	// SortColumn 排序字段见附录column字段
	SortColumn string `json:"sort_column,omitempty"`
	// Sort 升降序asc：升序desc：降序
	Sort enum.SortType `json:"sort,omitempty"`
	// PageNum 页数，默认1
	PageNum int64 `json:"page_num,omitempty"`
	// PageSize 页大小，默认20,最大500
	PageSize int64 `json:"page_size,omitempty"`
	// DataCaliber 数据指标归因时间类型 0-点击时间 1-转化时间
	DataCaliber int `json:"data_caliber,omitempty"`
	// Filters 过滤条件
	Filters []NoteFilterClause `json:"filters,omitempty"`
}

// Encode implement PostRequest interface
func (r NoteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// NoteResponse 笔记层级离线报表数据 API Response
type NoteResponse struct {
	model.BaseResponse
	// Data 数据
	Data *NoteReportList `json:"data,omitempty"`
}

// NoteReportList 笔记层级离线数据列表
type NoteReportList struct {
	// TotalCount 总条数
	TotalCount int64 `json:"total_count,omitempty"`
	// List 详细数据
	List []NoteReport `json:"data_list,omitempty"`
	// AggregationData 汇总数据
	AggregationData *NoteReport `json:"aggregation_data,omitempty"`
}

// NoteReport 笔记层级离线数据报表
type NoteReport struct {
	// NoteID 笔记id
	NoteID string `json:"note_id,omitempty"`
	// NoteTitle 笔记标题
	NoteTitle string `json:"note_title,omitempty"`
	// NoteImage 笔记图片
	NoteImage string `json:"note_image,omitempty"`
	// NoteJumpURL 笔记链接
	NoteJumpURL string `json:"note_jump_url,omitempty"`
	// Time 时间
	Time string `json:"time,omitempty"`
	// PageID 落地页id
	PageID string `json:"page_id,omitempty"`
	// ItemID 商品id
	ItemID string `json:"item_id,omitempty"`
	// LiveRedID 直播间id
	LiveRedID string `json:"live_red_id,omitempty"`
	// CountryName 国家; 细分条件选择国家才有
	CountryName string `json:"country_name,omitempty"`
	// Province 省份; 细分条件选择省份才有
	Province string `json:"province,omitempty"`
	// City 城市; 细分条件选择城市才有
	City string `json:"city,omitempty"`
	report.DataReportDTO
	// MessageFstReplyTimeAvg 平均响应时长(分); 账户下单次咨询响应时长加和/总咨询次数
	MessageFstReplyTimeAvg model.Float64 `json:"message_fst_reply_time_avg,omitempty"`
	// IUserNum 新增种草人群
	IUserNum model.Int64 `json:"i_user_num,omitempty"`
	// TiUserNum 新增深度种草人群
	TiUserNum model.Int64 `json:"ti_user_num,omitempty"`
	// IUserPrice 新增种草人群成本
	IUserPrice model.Float64 `json:"i_user_price,omitempty"`
	// TiUserPrice 新增深度种草人群成本
	TiUserPrice model.Float64 `json:"ti_user_price,omitempty"`
	// LiveAverageOrderCost 直播间支付订单成本; 推广消费/直播间支付订单量
	LiveAverageOrderCost model.Float64 `json:"live_average_order_cost,omitempty"`
	// PhoneCallCnt 电话拨打
	PhoneCallCnt model.Int64 `json:"phone_call_cnt,omitempty"`
	// PhoneCallSuccCnt 电话接通
	PhoneCallSuccCnt model.Int64 `json:"phone_call_succ_cnt,omitempty"`
	// WechatCopyCnt 微信复制
	WechatCopyCnt model.Int64 `json:"wechat_copy_cnt,omitempty"`
	// WechatCopySuccCnt 微信加为好友
	WechatCopySuccCnt model.Int64 `json:"wechat_copy_succ_cnt,omitempty"`
	// IdentityCertiCnt 身份认证
	IdentityCertiCnt model.Int64 `json:"identity_certi_cnt,omitempty"`
	// CommodityBuyCnt 付费
	CommodityBuyCnt model.Int64 `json:"commodity_buy_cnt,omitempty"`
	// ExternalGoodsVisit7 行业商品点击量
	ExternalGoodsVisit7 model.Int64 `json:"external_goods_visit_7,omitempty"`
	// ExternalGoodsVisitPrice7 行业商品点击成本
	ExternalGoodsVisitPrice7 model.Float64 `json:"external_goods_visit_price_7,omitempty"`
	// ExternalGoodsVisitRate7 行业商品点击转化率
	ExternalGoodsVisitRate7 model.Float64 `json:"external_goods_visit_rate_7,omitempty"`
	// JdActiveUserNum 小红盟站外活跃行为UV
	JdActiveUserNum model.Int64 `json:"jd_active_user_num,omitempty"`
	// JdActiveUserNumCvrNew 小红盟站外活跃行为率
	JdActiveUserNumCvrNew model.Float64 `json:"jd_active_user_num_cvr_new,omitempty"`
	// JdActiveUserNumCpl 小红盟站外活跃成本
	JdActiveUserNumCpl model.Float64 `json:"jd_active_user_num_cpl,omitempty"`
	// JdTaskFee 小红盟任务期消费
	JdTaskFee model.Float64 `json:"jd_task_fee,omitempty"`
	// JdTaskReadUserCnt 小红盟任务期阅读UV
	JdTaskReadUserCnt model.Int64 `json:"jd_task_read_user_cnt,omitempty"`
	// AppDownloadButtonClickCnt APP下载按钮点击
	AppDownloadButtonClickCnt model.Int64 `json:"app_download_button_click_cnt,omitempty"`
	// AppDownloadButtonClickCtr APP下载按钮点击率
	AppDownloadButtonClickCtr model.Float64 `json:"app_download_button_click_ctr,omitempty"`
	// AppDownloadButtonClickCost APP下载按钮点击成本
	AppDownloadButtonClickCost model.Float64 `json:"app_download_button_click_cost,omitempty"`
	// AppActivateCnt 激活数
	AppActivateCnt model.Int64 `json:"app_activate_cnt,omitempty"`
	// AppActivateCost 激活成本
	AppActivateCost model.Float64 `json:"app_activate_cost,omitempty"`
	// AppActivateCtr 激活率
	AppActivateCtr model.Float64 `json:"app_activate_ctr,omitempty"`
	// AppRegisterCnt 注册数
	AppRegisterCnt model.Int64 `json:"app_register_cnt,omitempty"`
	// AppRegisterCost 注册成本
	AppRegisterCost model.Float64 `json:"app_register_cost,omitempty"`
	// AppRegisterCtr 注册率
	AppRegisterCtr model.Float64 `json:"app_register_ctr,omitempty"`
	// FirstAppPayCnt 首次付费数
	FirstAppPayCnt model.Int64 `json:"first_app_pay_cnt,omitempty"`
	// FirstAppPayCost 首次付费成本
	FirstAppPayCost model.Float64 `json:"first_app_pay_cost,omitempty"`
	// FirstAppPayCtr 首次付费率
	FirstAppPayCtr model.Float64 `json:"first_app_pay_ctr,omitempty"`
	// CurrentAppPayCnt 当日付费次数
	CurrentAppPayCnt model.Int64 `json:"current_app_pay_cnt,omitempty"`
	// CurrentAppPayCost 当日付费成本
	CurrentAppPayCost model.Float64 `json:"current_app_pay_cost,omitempty"`
	// AppKeyActionCnt 关键行为数
	AppKeyActionCnt model.Int64 `json:"app_key_action_cnt,omitempty"`
	// AppKeyActionCost 关键行为成本
	AppKeyActionCost model.Float64 `json:"app_key_action_cost,omitempty"`
	// AppKeyActionCtr 关键行为率
	AppKeyActionCtr model.Float64 `json:"app_key_action_ctr,omitempty"`
	// AppPayCnt7d 7日付费次数
	AppPayCnt7d model.Int64 `json:"app_pay_cnt_7d,omitempty"`
	// AppPayCost7d 7日付费成本
	AppPayCost7d model.Float64 `json:"app_pay_cost_7d,omitempty"`
	// AppPayAmount 付费金额
	AppPayAmount model.Float64 `json:"app_pay_amount,omitempty"`
	// AppPayRoi 付费ROI
	AppPayRoi model.Float64 `json:"app_pay_roi,omitempty"`
	// AppActivateAmount1d 当日LTV
	AppActivateAmount1d model.Float64 `json:"app_activate_amount_1d,omitempty"`
	// AppActivateAmount3d 三日LTV
	AppActivateAmount3d model.Float64 `json:"app_activate_amount_3d,omitempty"`
	// AppActivateAmount7d 七日LTV
	AppActivateAmount7d model.Float64 `json:"app_activate_amount_7d,omitempty"`
	// AppActivateAmount1dRoi 当日广告付费ROI
	AppActivateAmount1dRoi model.Float64 `json:"app_activate_amount_1d_roi,omitempty"`
	// AppActivateAmount3dRoi 三日广告付费ROI
	AppActivateAmount3dRoi model.Float64 `json:"app_activate_amount_3d_roi,omitempty"`
	// AppActivateAmount7dRoi 七日广告付费ROI
	AppActivateAmount7dRoi model.Float64 `json:"app_activate_amount_7d_roi,omitempty"`
	// Retention1dCnt 次留
	Retention1dCnt model.Int64 `json:"retention_1d_cnt,omitempty"`
	// Retention3dCnt 3日留存
	Retention3dCnt model.Int64 `json:"retention_3d_cnt,omitempty"`
	// Retention7dCnt 7日留存
	Retention7dCnt model.Int64 `json:"retention_7d_cnt,omitempty"`
	// AddWechatCount 添加企微量
	AddWechatCount model.Int64 `json:"add_wechat_count,omitempty"`
	// AddWechatCost 添加企微成本
	AddWechatCost model.Float64 `json:"add_wechat_cost,omitempty"`
	// AddWechatSucCount 成功添加企微量
	AddWechatSucCount model.Int64 `json:"add_wechat_suc_count,omitempty"`
	// AddWechatSucCost 成功添加企微成本
	AddWechatSucCost model.Float64 `json:"add_wechat_suc_cost,omitempty"`
	// WechatTalkCount 企微开口量
	WechatTalkCount model.Int64 `json:"wechat_talk_count,omitempty"`
	// WechatTalkCost 企微开口成本
	WechatTalkCost model.Float64 `json:"wechat_talk_cost,omitempty"`
	// ShopPoiClickNum 门店组件点击量
	ShopPoiClickNum model.Int64 `json:"shop_poi_click_num,omitempty"`
	// ShopPoiPagePv 门店页面访问量
	ShopPoiPagePv model.Int64 `json:"shop_poi_page_pv,omitempty"`
	// ShopPoiPageVisitPrice 门店页面访问成本
	ShopPoiPageVisitPrice model.Float64 `json:"shop_poi_page_visit_price,omitempty"`
	// ShopPoiPageNavigateClick 门店页面导航栏按钮点击量
	ShopPoiPageNavigateClick model.Int64 `json:"shop_poi_page_navigate_click,omitempty"`
	// WechatAppletsOpenCnt 微信打开按钮点击次数
	WechatAppletsOpenCnt model.Int64 `json:"wechat_applets_open_cnt,omitempty"`
	// WechatAppletsPayCnt 微信小程序付费次数
	WechatAppletsPayCnt model.Int64 `json:"wechat_applets_pay_cnt,omitempty"`
	// WechatAppletsActivateCnt 微信小程序激活次数
	WechatAppletsActivateCnt model.Int64 `json:"wechat_applets_activate_cnt,omitempty"`
	// WechatAppletsPayAmount 微信小程序付费总金额
	WechatAppletsPayAmount model.Float64 `json:"wechat_applets_pay_amount,omitempty"`
	// WechatAppletsPayAmount3d 微信小程序激活后三日付费总金额
	WechatAppletsPayAmount3d model.Float64 `json:"wechat_applets_pay_amount_3d,omitempty"`
	// WechatAppletsPayAmount7d 微信小程序激活后七日付费总金额
	WechatAppletsPayAmount7d model.Float64 `json:"wechat_applets_pay_amount_7d,omitempty"`
	// WechatAppletsPayCnt3d 微信小程序激活后三日付费次数
	WechatAppletsPayCnt3d model.Int64 `json:"wechat_applets_pay_cnt_3d,omitempty"`
	// WechatAppletsPayCnt7d 微信小程序激活后七日付费次数
	WechatAppletsPayCnt7d model.Int64 `json:"wechat_applets_pay_cnt_7d,omitempty"`
	// CurrentWechatAppletsFirstPayCnt 微信小程序当日首次付费次数
	CurrentWechatAppletsFirstPayCnt model.Int64 `json:"currentWechatAppletsFirstPayCnt,omitempty"`
	// CurrentWechatAppletsFirstPayAmount 微信小程序当日首次付费金额
	CurrentWechatAppletsFirstPayAmount model.Float64 `json:"currentWechatAppletsFirstPayAmount,omitempty"`
	// CurrentWechatAppletsPayCnt 微信小程序当日付费次数
	CurrentWechatAppletsPayCnt model.Int64 `json:"currentWechatAppletsPayCnt,omitempty"`
	// CurrentWechatAppletsPayAmount 微信小程序当日付费总金额
	CurrentWechatAppletsPayAmount model.Float64 `json:"currentWechatAppletsPayAmount,omitempty"`
	// CurrentWechatAppletsActivateCnt 微信小程序当日激活次数
	CurrentWechatAppletsActivateCnt model.Int64 `json:"currentWechatAppletsActivateCnt,omitempty"`
}
