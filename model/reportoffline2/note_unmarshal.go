package reportoffline2

import (
	"encoding/json"

	"github.com/jundaychan/spotlight-mapi/model"
)

// UnmarshalJSON 修复 Go 的一个真实坑：匿名内嵌 report.DataReportDTO（有自定义 UnmarshalJSON）
// 会让 NoteReport 自己声明的字段（note_id 等 84 个）在标准 struct 解码时被静默清空——
// 不报错、不少一个字段，就是全部拿到零值。2026-09-03 实测：note_id 变成空字符串后，
// fetchAdvertiserNoteRows 的 `if ext == "" { continue }` 把每一行都当无效数据跳过，
// 整条笔记线索同步表面「成功」、实际一行都没入库，还被误判成「账号历史真的没数据」。
//
// 修法：分两次 Unmarshal——一次单独喂给 DataReportDTO（继续走它自己的 camelCase 兼容逻辑），
// 一次喂给一个只含 NoteReport 自身字段、不内嵌任何东西的镜像结构体，再手动拼回来。
// 复现与验证见 model/report/embed_bug_repro_test.go 的 TestEmbedNoteIDBug(Fix)。
func (r *NoteReport) UnmarshalJSON(b []byte) error {
	if err := json.Unmarshal(b, &r.DataReportDTO); err != nil {
		return err
	}
	var o struct {
		NoteID                             string        `json:"note_id,omitempty"`
		NoteTitle                          string        `json:"note_title,omitempty"`
		NoteImage                          string        `json:"note_image,omitempty"`
		NoteJumpURL                        string        `json:"note_jump_url,omitempty"`
		Time                               string        `json:"time,omitempty"`
		PageID                             string        `json:"page_id,omitempty"`
		ItemID                             string        `json:"item_id,omitempty"`
		LiveRedID                          string        `json:"live_red_id,omitempty"`
		CountryName                        string        `json:"country_name,omitempty"`
		Province                           string        `json:"province,omitempty"`
		City                               string        `json:"city,omitempty"`
		MessageFstReplyTimeAvg             model.Float64 `json:"message_fst_reply_time_avg,omitempty"`
		IUserNum                           model.Int64   `json:"i_user_num,omitempty"`
		TiUserNum                          model.Int64   `json:"ti_user_num,omitempty"`
		IUserPrice                         model.Float64 `json:"i_user_price,omitempty"`
		TiUserPrice                        model.Float64 `json:"ti_user_price,omitempty"`
		LiveAverageOrderCost               model.Float64 `json:"live_average_order_cost,omitempty"`
		PhoneCallCnt                       model.Int64   `json:"phone_call_cnt,omitempty"`
		PhoneCallSuccCnt                   model.Int64   `json:"phone_call_succ_cnt,omitempty"`
		WechatCopyCnt                      model.Int64   `json:"wechat_copy_cnt,omitempty"`
		WechatCopySuccCnt                  model.Int64   `json:"wechat_copy_succ_cnt,omitempty"`
		IdentityCertiCnt                   model.Int64   `json:"identity_certi_cnt,omitempty"`
		CommodityBuyCnt                    model.Int64   `json:"commodity_buy_cnt,omitempty"`
		ExternalGoodsVisit7                model.Int64   `json:"external_goods_visit_7,omitempty"`
		ExternalGoodsVisitPrice7           model.Float64 `json:"external_goods_visit_price_7,omitempty"`
		ExternalGoodsVisitRate7            model.Float64 `json:"external_goods_visit_rate_7,omitempty"`
		JdActiveUserNum                    model.Int64   `json:"jd_active_user_num,omitempty"`
		JdActiveUserNumCvrNew              model.Float64 `json:"jd_active_user_num_cvr_new,omitempty"`
		JdActiveUserNumCpl                 model.Float64 `json:"jd_active_user_num_cpl,omitempty"`
		JdTaskFee                          model.Float64 `json:"jd_task_fee,omitempty"`
		JdTaskReadUserCnt                  model.Int64   `json:"jd_task_read_user_cnt,omitempty"`
		AppDownloadButtonClickCnt          model.Int64   `json:"app_download_button_click_cnt,omitempty"`
		AppDownloadButtonClickCtr          model.Float64 `json:"app_download_button_click_ctr,omitempty"`
		AppDownloadButtonClickCost         model.Float64 `json:"app_download_button_click_cost,omitempty"`
		AppActivateCnt                     model.Int64   `json:"app_activate_cnt,omitempty"`
		AppActivateCost                    model.Float64 `json:"app_activate_cost,omitempty"`
		AppActivateCtr                     model.Float64 `json:"app_activate_ctr,omitempty"`
		AppRegisterCnt                     model.Int64   `json:"app_register_cnt,omitempty"`
		AppRegisterCost                    model.Float64 `json:"app_register_cost,omitempty"`
		AppRegisterCtr                     model.Float64 `json:"app_register_ctr,omitempty"`
		FirstAppPayCnt                     model.Int64   `json:"first_app_pay_cnt,omitempty"`
		FirstAppPayCost                    model.Float64 `json:"first_app_pay_cost,omitempty"`
		FirstAppPayCtr                     model.Float64 `json:"first_app_pay_ctr,omitempty"`
		CurrentAppPayCnt                   model.Int64   `json:"current_app_pay_cnt,omitempty"`
		CurrentAppPayCost                  model.Float64 `json:"current_app_pay_cost,omitempty"`
		AppKeyActionCnt                    model.Int64   `json:"app_key_action_cnt,omitempty"`
		AppKeyActionCost                   model.Float64 `json:"app_key_action_cost,omitempty"`
		AppKeyActionCtr                    model.Float64 `json:"app_key_action_ctr,omitempty"`
		AppPayCnt7d                        model.Int64   `json:"app_pay_cnt_7d,omitempty"`
		AppPayCost7d                       model.Float64 `json:"app_pay_cost_7d,omitempty"`
		AppPayAmount                       model.Float64 `json:"app_pay_amount,omitempty"`
		AppPayRoi                          model.Float64 `json:"app_pay_roi,omitempty"`
		AppActivateAmount1d                model.Float64 `json:"app_activate_amount_1d,omitempty"`
		AppActivateAmount3d                model.Float64 `json:"app_activate_amount_3d,omitempty"`
		AppActivateAmount7d                model.Float64 `json:"app_activate_amount_7d,omitempty"`
		AppActivateAmount1dRoi             model.Float64 `json:"app_activate_amount_1d_roi,omitempty"`
		AppActivateAmount3dRoi             model.Float64 `json:"app_activate_amount_3d_roi,omitempty"`
		AppActivateAmount7dRoi             model.Float64 `json:"app_activate_amount_7d_roi,omitempty"`
		Retention1dCnt                     model.Int64   `json:"retention_1d_cnt,omitempty"`
		Retention3dCnt                     model.Int64   `json:"retention_3d_cnt,omitempty"`
		Retention7dCnt                     model.Int64   `json:"retention_7d_cnt,omitempty"`
		AddWechatCount                     model.Int64   `json:"add_wechat_count,omitempty"`
		AddWechatCost                      model.Float64 `json:"add_wechat_cost,omitempty"`
		AddWechatSucCount                  model.Int64   `json:"add_wechat_suc_count,omitempty"`
		AddWechatSucCost                   model.Float64 `json:"add_wechat_suc_cost,omitempty"`
		WechatTalkCount                    model.Int64   `json:"wechat_talk_count,omitempty"`
		WechatTalkCost                     model.Float64 `json:"wechat_talk_cost,omitempty"`
		ShopPoiClickNum                    model.Int64   `json:"shop_poi_click_num,omitempty"`
		ShopPoiPagePv                      model.Int64   `json:"shop_poi_page_pv,omitempty"`
		ShopPoiPageVisitPrice              model.Float64 `json:"shop_poi_page_visit_price,omitempty"`
		ShopPoiPageNavigateClick           model.Int64   `json:"shop_poi_page_navigate_click,omitempty"`
		WechatAppletsOpenCnt               model.Int64   `json:"wechat_applets_open_cnt,omitempty"`
		WechatAppletsPayCnt                model.Int64   `json:"wechat_applets_pay_cnt,omitempty"`
		WechatAppletsActivateCnt           model.Int64   `json:"wechat_applets_activate_cnt,omitempty"`
		WechatAppletsPayAmount             model.Float64 `json:"wechat_applets_pay_amount,omitempty"`
		WechatAppletsPayAmount3d           model.Float64 `json:"wechat_applets_pay_amount_3d,omitempty"`
		WechatAppletsPayAmount7d           model.Float64 `json:"wechat_applets_pay_amount_7d,omitempty"`
		WechatAppletsPayCnt3d              model.Int64   `json:"wechat_applets_pay_cnt_3d,omitempty"`
		WechatAppletsPayCnt7d              model.Int64   `json:"wechat_applets_pay_cnt_7d,omitempty"`
		CurrentWechatAppletsFirstPayCnt    model.Int64   `json:"currentWechatAppletsFirstPayCnt,omitempty"`
		CurrentWechatAppletsFirstPayAmount model.Float64 `json:"currentWechatAppletsFirstPayAmount,omitempty"`
		CurrentWechatAppletsPayCnt         model.Int64   `json:"currentWechatAppletsPayCnt,omitempty"`
		CurrentWechatAppletsPayAmount      model.Float64 `json:"currentWechatAppletsPayAmount,omitempty"`
		CurrentWechatAppletsActivateCnt    model.Int64   `json:"currentWechatAppletsActivateCnt,omitempty"`
	}
	if err := json.Unmarshal(b, &o); err != nil {
		return err
	}
	r.NoteID = o.NoteID
	r.NoteTitle = o.NoteTitle
	r.NoteImage = o.NoteImage
	r.NoteJumpURL = o.NoteJumpURL
	r.Time = o.Time
	r.PageID = o.PageID
	r.ItemID = o.ItemID
	r.LiveRedID = o.LiveRedID
	r.CountryName = o.CountryName
	r.Province = o.Province
	r.City = o.City
	r.MessageFstReplyTimeAvg = o.MessageFstReplyTimeAvg
	r.IUserNum = o.IUserNum
	r.TiUserNum = o.TiUserNum
	r.IUserPrice = o.IUserPrice
	r.TiUserPrice = o.TiUserPrice
	r.LiveAverageOrderCost = o.LiveAverageOrderCost
	r.PhoneCallCnt = o.PhoneCallCnt
	r.PhoneCallSuccCnt = o.PhoneCallSuccCnt
	r.WechatCopyCnt = o.WechatCopyCnt
	r.WechatCopySuccCnt = o.WechatCopySuccCnt
	r.IdentityCertiCnt = o.IdentityCertiCnt
	r.CommodityBuyCnt = o.CommodityBuyCnt
	r.ExternalGoodsVisit7 = o.ExternalGoodsVisit7
	r.ExternalGoodsVisitPrice7 = o.ExternalGoodsVisitPrice7
	r.ExternalGoodsVisitRate7 = o.ExternalGoodsVisitRate7
	r.JdActiveUserNum = o.JdActiveUserNum
	r.JdActiveUserNumCvrNew = o.JdActiveUserNumCvrNew
	r.JdActiveUserNumCpl = o.JdActiveUserNumCpl
	r.JdTaskFee = o.JdTaskFee
	r.JdTaskReadUserCnt = o.JdTaskReadUserCnt
	r.AppDownloadButtonClickCnt = o.AppDownloadButtonClickCnt
	r.AppDownloadButtonClickCtr = o.AppDownloadButtonClickCtr
	r.AppDownloadButtonClickCost = o.AppDownloadButtonClickCost
	r.AppActivateCnt = o.AppActivateCnt
	r.AppActivateCost = o.AppActivateCost
	r.AppActivateCtr = o.AppActivateCtr
	r.AppRegisterCnt = o.AppRegisterCnt
	r.AppRegisterCost = o.AppRegisterCost
	r.AppRegisterCtr = o.AppRegisterCtr
	r.FirstAppPayCnt = o.FirstAppPayCnt
	r.FirstAppPayCost = o.FirstAppPayCost
	r.FirstAppPayCtr = o.FirstAppPayCtr
	r.CurrentAppPayCnt = o.CurrentAppPayCnt
	r.CurrentAppPayCost = o.CurrentAppPayCost
	r.AppKeyActionCnt = o.AppKeyActionCnt
	r.AppKeyActionCost = o.AppKeyActionCost
	r.AppKeyActionCtr = o.AppKeyActionCtr
	r.AppPayCnt7d = o.AppPayCnt7d
	r.AppPayCost7d = o.AppPayCost7d
	r.AppPayAmount = o.AppPayAmount
	r.AppPayRoi = o.AppPayRoi
	r.AppActivateAmount1d = o.AppActivateAmount1d
	r.AppActivateAmount3d = o.AppActivateAmount3d
	r.AppActivateAmount7d = o.AppActivateAmount7d
	r.AppActivateAmount1dRoi = o.AppActivateAmount1dRoi
	r.AppActivateAmount3dRoi = o.AppActivateAmount3dRoi
	r.AppActivateAmount7dRoi = o.AppActivateAmount7dRoi
	r.Retention1dCnt = o.Retention1dCnt
	r.Retention3dCnt = o.Retention3dCnt
	r.Retention7dCnt = o.Retention7dCnt
	r.AddWechatCount = o.AddWechatCount
	r.AddWechatCost = o.AddWechatCost
	r.AddWechatSucCount = o.AddWechatSucCount
	r.AddWechatSucCost = o.AddWechatSucCost
	r.WechatTalkCount = o.WechatTalkCount
	r.WechatTalkCost = o.WechatTalkCost
	r.ShopPoiClickNum = o.ShopPoiClickNum
	r.ShopPoiPagePv = o.ShopPoiPagePv
	r.ShopPoiPageVisitPrice = o.ShopPoiPageVisitPrice
	r.ShopPoiPageNavigateClick = o.ShopPoiPageNavigateClick
	r.WechatAppletsOpenCnt = o.WechatAppletsOpenCnt
	r.WechatAppletsPayCnt = o.WechatAppletsPayCnt
	r.WechatAppletsActivateCnt = o.WechatAppletsActivateCnt
	r.WechatAppletsPayAmount = o.WechatAppletsPayAmount
	r.WechatAppletsPayAmount3d = o.WechatAppletsPayAmount3d
	r.WechatAppletsPayAmount7d = o.WechatAppletsPayAmount7d
	r.WechatAppletsPayCnt3d = o.WechatAppletsPayCnt3d
	r.WechatAppletsPayCnt7d = o.WechatAppletsPayCnt7d
	r.CurrentWechatAppletsFirstPayCnt = o.CurrentWechatAppletsFirstPayCnt
	r.CurrentWechatAppletsFirstPayAmount = o.CurrentWechatAppletsFirstPayAmount
	r.CurrentWechatAppletsPayCnt = o.CurrentWechatAppletsPayCnt
	r.CurrentWechatAppletsPayAmount = o.CurrentWechatAppletsPayAmount
	r.CurrentWechatAppletsActivateCnt = o.CurrentWechatAppletsActivateCnt
	return nil
}
