// 由实测聚光各报表口生成（2026-09-24）；改列表前先重新实测。

package report

// reportColumns 各报表口可查询的指标列（camelCase）。
//
// 2026-09-20 起离线/实时报表支持 columns 按需查询，10-01 之后不传只回基础指标——
// 线索、私信、互动这些列会整列变 0 且不报错。列名大小写敏感（reservePV 不是 reservePv），
// 传了任何一个该层不认的列整个请求 10002 失败，所以按层分开列、只收实测通过的。
// DataReportDTO 里有、这里没有的字段 = 该层上游不支持（平铺报文里也没有，本来就是 0）。
var reportColumns = map[string][]string{
	"/jg/data/report/offline/account": {
		"fee", "impression", "click", "ctr", "acp", "cpm", "like", "comment", "collect", "follow", "share",
		"interaction", "cpi", "actionButtonClick", "actionButtonCtr", "screenshot", "picSave", "reservePV",
		"clkLiveEntryPv", "clkLiveEntryPvCost", "clkLiveAvgViewTime", "clkLiveAllFollow", "clkLive5sEntryPv",
		"clkLive5sEntryUvCost", "clkLiveComment", "searchCmtClick", "searchCmtClickCvr", "searchCmtAfterRead",
		"searchCmtAfterReadAvg", "goodsVisit", "goodsVisitPrice", "sellerVisit", "sellerVisitPrice",
		"shoppingCartAdd", "addCartPrice", "presaleOrderNum7d", "presaleOrderGmv7d", "goodsOrder",
		"goodsOrderPrice", "rgmv", "roi", "successGoodsOrder", "clickOrderCvr", "purchaseOrderPrice7d",
		"purchaseOrderGmv7d", "purchaseOrderRoi7d", "clkLiveRoomOrderNum", "liveAverageOrderCost",
		"clkLiveRoomRgmv", "clkLiveRoomRoi", "leads", "leadsCpl", "landingPageVisit", "leadsButtonImpression",
		"validLeads", "validLeadsCpl", "leadsCvr", "messageUser", "message", "messageConsult", "initiativeMessage",
		"messageConsultCpl", "initiativeMessageCpl", "msgLeadsNum", "msgLeadsCost", "externalGoodsOrder7",
		"externalRgmv7", "externalGoodsOrderPrice7", "externalGoodsOrderRate7", "externalRoi7",
		"externalGoodsOrder15", "externalRgmv15", "externalGoodsOrderPrice15", "externalGoodsOrderRate15",
		"externalRoi15", "externalGoodsOrder30", "externalRgmv30", "externalGoodsOrderPrice30",
		"externalGoodsOrderRate30", "externalRoi30", "invokeAppOpenCnt", "invokeAppOpenCost",
		"invokeAppEnterStoreCnt", "invokeAppEnterStoreCost", "invokeAppEngagementCnt", "invokeAppEngagementCost",
		"invokeAppPaymentCnt", "invokeAppPaymentCost", "searchInvokeButtonClickCnt", "searchInvokeButtonClickCost",
	},
	"/jg/data/report/offline/campaign": {
		"fee", "impression", "click", "ctr", "acp", "cpm", "like", "comment", "collect", "follow", "share",
		"interaction", "cpi", "actionButtonClick", "actionButtonCtr", "screenshot", "picSave", "reservePV",
		"clkLiveEntryPv", "clkLiveEntryPvCost", "clkLiveAvgViewTime", "clkLiveAllFollow", "clkLive5sEntryPv",
		"clkLive5sEntryUvCost", "clkLiveComment", "searchCmtClick", "searchCmtClickCvr", "searchCmtAfterRead",
		"searchCmtAfterReadAvg", "goodsVisit", "goodsVisitPrice", "sellerVisit", "sellerVisitPrice",
		"shoppingCartAdd", "addCartPrice", "presaleOrderNum7d", "presaleOrderGmv7d", "goodsOrder",
		"goodsOrderPrice", "rgmv", "roi", "successGoodsOrder", "clickOrderCvr", "purchaseOrderPrice7d",
		"purchaseOrderGmv7d", "purchaseOrderRoi7d", "clkLiveRoomOrderNum", "liveAverageOrderCost",
		"clkLiveRoomRgmv", "clkLiveRoomRoi", "leads", "leadsCpl", "landingPageVisit", "leadsButtonImpression",
		"validLeads", "validLeadsCpl", "leadsCvr", "messageUser", "message", "messageConsult", "initiativeMessage",
		"messageConsultCpl", "initiativeMessageCpl", "msgLeadsNum", "msgLeadsCost", "externalGoodsOrder7",
		"externalRgmv7", "externalGoodsOrderPrice7", "externalGoodsOrderRate7", "externalRoi7",
		"externalGoodsOrder15", "externalRgmv15", "externalGoodsOrderPrice15", "externalGoodsOrderRate15",
		"externalRoi15", "externalGoodsOrder30", "externalRgmv30", "externalGoodsOrderPrice30",
		"externalGoodsOrderRate30", "externalRoi30", "invokeAppOpenCnt", "invokeAppOpenCost",
		"invokeAppEnterStoreCnt", "invokeAppEnterStoreCost", "invokeAppEngagementCnt", "invokeAppEngagementCost",
		"invokeAppPaymentCnt", "invokeAppPaymentCost", "searchInvokeButtonClickCnt", "searchInvokeButtonClickCost",
	},
	"/jg/data/report/offline/creative": {
		"fee", "impression", "click", "ctr", "acp", "cpm", "like", "comment", "collect", "follow", "share",
		"interaction", "cpi", "actionButtonClick", "actionButtonCtr", "screenshot", "picSave", "reservePV",
		"clkLiveEntryPv", "clkLiveEntryPvCost", "clkLiveAvgViewTime", "clkLiveAllFollow", "clkLive5sEntryPv",
		"clkLive5sEntryUvCost", "clkLiveComment", "searchCmtClick", "searchCmtClickCvr", "searchCmtAfterRead",
		"searchCmtAfterReadAvg", "goodsVisit", "goodsVisitPrice", "sellerVisit", "sellerVisitPrice",
		"shoppingCartAdd", "addCartPrice", "presaleOrderNum7d", "presaleOrderGmv7d", "goodsOrder",
		"goodsOrderPrice", "rgmv", "roi", "successGoodsOrder", "clickOrderCvr", "purchaseOrderPrice7d",
		"purchaseOrderGmv7d", "purchaseOrderRoi7d", "clkLiveRoomOrderNum", "liveAverageOrderCost",
		"clkLiveRoomRgmv", "clkLiveRoomRoi", "leads", "leadsCpl", "landingPageVisit", "leadsButtonImpression",
		"validLeads", "validLeadsCpl", "leadsCvr", "messageUser", "message", "messageConsult", "initiativeMessage",
		"messageConsultCpl", "initiativeMessageCpl", "msgLeadsNum", "msgLeadsCost", "externalGoodsOrder7",
		"externalRgmv7", "externalGoodsOrderPrice7", "externalGoodsOrderRate7", "externalRoi7",
		"externalGoodsOrder15", "externalRgmv15", "externalGoodsOrderPrice15", "externalGoodsOrderRate15",
		"externalRoi15", "externalGoodsOrder30", "externalRgmv30", "externalGoodsOrderPrice30",
		"externalGoodsOrderRate30", "externalRoi30", "invokeAppOpenCnt", "invokeAppOpenCost",
		"invokeAppEnterStoreCnt", "invokeAppEnterStoreCost", "invokeAppEngagementCnt", "invokeAppEngagementCost",
		"invokeAppPaymentCnt", "invokeAppPaymentCost", "searchInvokeButtonClickCnt", "searchInvokeButtonClickCost",
	},
	"/jg/data/report/offline/keyword": {
		"fee", "impression", "click", "ctr", "acp", "cpm", "like", "comment", "collect", "follow", "share",
		"interaction", "cpi", "actionButtonClick", "actionButtonCtr", "screenshot", "picSave", "reservePV",
		"clkLiveEntryPv", "clkLiveEntryPvCost", "goodsVisit", "goodsVisitPrice", "sellerVisit", "sellerVisitPrice",
		"shoppingCartAdd", "addCartPrice", "presaleOrderNum7d", "presaleOrderGmv7d", "goodsOrder",
		"goodsOrderPrice", "rgmv", "roi", "successGoodsOrder", "clickOrderCvr", "purchaseOrderPrice7d",
		"purchaseOrderGmv7d", "purchaseOrderRoi7d", "clkLiveRoomOrderNum", "liveAverageOrderCost", "leads",
		"leadsCpl", "landingPageVisit", "leadsButtonImpression", "validLeads", "validLeadsCpl", "leadsCvr",
		"messageUser", "message", "messageConsult", "initiativeMessage", "messageConsultCpl",
		"initiativeMessageCpl", "msgLeadsNum", "msgLeadsCost", "externalGoodsOrder7", "externalRgmv7",
		"externalGoodsOrderPrice7", "externalGoodsOrderRate7", "externalRoi7", "externalGoodsOrder15",
		"externalRgmv15", "externalGoodsOrderPrice15", "externalGoodsOrderRate15", "externalRoi15",
		"externalGoodsOrder30", "externalRgmv30", "externalGoodsOrderPrice30", "externalGoodsOrderRate30",
		"externalRoi30", "wordAvgLocation", "invokeAppOpenCnt", "invokeAppOpenCost", "invokeAppEnterStoreCnt",
		"invokeAppEnterStoreCost", "invokeAppEngagementCnt", "invokeAppEngagementCost", "invokeAppPaymentCnt",
		"invokeAppPaymentCost", "searchInvokeButtonClickCnt", "searchInvokeButtonClickCost",
	},
	"/jg/data/report/offline/unit": {
		"fee", "impression", "click", "ctr", "acp", "cpm", "like", "comment", "collect", "follow", "share",
		"interaction", "cpi", "actionButtonClick", "actionButtonCtr", "screenshot", "picSave", "reservePV",
		"clkLiveEntryPv", "clkLiveEntryPvCost", "clkLiveAvgViewTime", "clkLiveAllFollow", "clkLive5sEntryPv",
		"clkLive5sEntryUvCost", "clkLiveComment", "searchCmtClick", "searchCmtClickCvr", "searchCmtAfterRead",
		"searchCmtAfterReadAvg", "goodsVisit", "goodsVisitPrice", "sellerVisit", "sellerVisitPrice",
		"shoppingCartAdd", "addCartPrice", "presaleOrderNum7d", "presaleOrderGmv7d", "goodsOrder",
		"goodsOrderPrice", "rgmv", "roi", "successGoodsOrder", "clickOrderCvr", "purchaseOrderPrice7d",
		"purchaseOrderGmv7d", "purchaseOrderRoi7d", "clkLiveRoomOrderNum", "liveAverageOrderCost",
		"clkLiveRoomRgmv", "clkLiveRoomRoi", "leads", "leadsCpl", "landingPageVisit", "leadsButtonImpression",
		"validLeads", "validLeadsCpl", "leadsCvr", "messageUser", "message", "messageConsult", "initiativeMessage",
		"messageConsultCpl", "initiativeMessageCpl", "msgLeadsNum", "msgLeadsCost", "externalGoodsOrder7",
		"externalRgmv7", "externalGoodsOrderPrice7", "externalGoodsOrderRate7", "externalRoi7",
		"externalGoodsOrder15", "externalRgmv15", "externalGoodsOrderPrice15", "externalGoodsOrderRate15",
		"externalRoi15", "externalGoodsOrder30", "externalRgmv30", "externalGoodsOrderPrice30",
		"externalGoodsOrderRate30", "externalRoi30", "invokeAppOpenCnt", "invokeAppOpenCost",
		"invokeAppEnterStoreCnt", "invokeAppEnterStoreCost", "invokeAppEngagementCnt", "invokeAppEngagementCost",
		"invokeAppPaymentCnt", "invokeAppPaymentCost", "searchInvokeButtonClickCnt", "searchInvokeButtonClickCost",
	},
	"/jg/data/report/realtime/account": {
		"fee", "impression", "click", "ctr", "acp", "cpm", "like", "comment", "collect", "follow", "share",
		"interaction", "cpi", "actionButtonClick", "actionButtonCtr", "screenshot", "picSave", "reservePV",
		"clkLiveEntryPv", "clkLiveEntryPvCost", "clkLiveAvgViewTime", "clkLiveAllFollow", "clkLive5sEntryPv",
		"clkLive5sEntryUvCost", "clkLiveComment", "searchCmtClick", "searchCmtClickCvr", "searchCmtAfterRead",
		"searchCmtAfterReadAvg", "goodsVisit", "goodsVisitPrice", "sellerVisit", "sellerVisitPrice",
		"shoppingCartAdd", "addCartPrice", "presaleOrderNum7d", "presaleOrderGmv7d", "goodsOrder",
		"goodsOrderPrice", "rgmv", "roi", "successGoodsOrder", "clickOrderCvr", "purchaseOrderPrice7d",
		"purchaseOrderGmv7d", "purchaseOrderRoi7d", "clkLiveRoomOrderNum", "liveAverageOrderCost",
		"clkLiveRoomRgmv", "clkLiveRoomRoi", "leads", "leadsCpl", "landingPageVisit", "leadsButtonImpression",
		"validLeads", "validLeadsCpl", "leadsCvr", "phoneCallCnt", "phoneCallSuccCnt", "wechatCopyCnt",
		"wechatCopySuccCnt", "commodityBuyCnt", "messageUser", "message", "messageConsult", "initiativeMessage",
		"messageConsultCpl", "initiativeMessageCpl", "msgLeadsNum", "msgLeadsCost", "externalGoodsVisit24h",
		"externalGoodsVisitRate24h", "externalGoodsOrder7", "externalRgmv7", "externalGoodsOrderPrice7",
		"externalGoodsOrderRate7", "externalRoi7", "externalGoodsOrder15", "externalRgmv15",
		"externalGoodsOrderPrice15", "externalGoodsOrderRate15", "externalRoi15", "externalGoodsOrder30",
		"externalRgmv30", "externalGoodsOrderPrice30", "externalGoodsOrderRate30", "externalRoi30",
		"externalLeads", "externalLeadsCpl", "invokeAppOpenCnt", "invokeAppOpenCost", "invokeAppEnterStoreCnt",
		"invokeAppEnterStoreCost", "invokeAppEngagementCnt", "invokeAppEngagementCost", "invokeAppPaymentCnt",
		"invokeAppPaymentCost", "searchInvokeButtonClickCnt", "searchInvokeButtonClickCost",
	},
	"/jg/data/report/realtime/campaign": {
		"fee", "impression", "click", "ctr", "acp", "cpm", "like", "comment", "collect", "follow", "share",
		"interaction", "cpi", "actionButtonClick", "actionButtonCtr", "screenshot", "picSave", "reservePV",
		"clkLiveEntryPv", "clkLiveEntryPvCost", "clkLiveAvgViewTime", "clkLiveAllFollow", "clkLive5sEntryPv",
		"clkLive5sEntryUvCost", "clkLiveComment", "searchCmtClick", "searchCmtClickCvr", "searchCmtAfterRead",
		"searchCmtAfterReadAvg", "goodsVisit", "goodsVisitPrice", "sellerVisit", "sellerVisitPrice",
		"shoppingCartAdd", "addCartPrice", "presaleOrderNum7d", "presaleOrderGmv7d", "goodsOrder",
		"goodsOrderPrice", "rgmv", "roi", "successGoodsOrder", "clickOrderCvr", "purchaseOrderPrice7d",
		"purchaseOrderGmv7d", "purchaseOrderRoi7d", "clkLiveRoomOrderNum", "liveAverageOrderCost",
		"clkLiveRoomRgmv", "clkLiveRoomRoi", "leads", "leadsCpl", "landingPageVisit", "leadsButtonImpression",
		"validLeads", "validLeadsCpl", "leadsCvr", "phoneCallCnt", "phoneCallSuccCnt", "wechatCopyCnt",
		"wechatCopySuccCnt", "commodityBuyCnt", "messageUser", "message", "messageConsult", "initiativeMessage",
		"messageConsultCpl", "initiativeMessageCpl", "msgLeadsNum", "msgLeadsCost", "externalGoodsVisit24h",
		"externalGoodsVisitRate24h", "externalGoodsOrder7", "externalRgmv7", "externalGoodsOrderPrice7",
		"externalGoodsOrderRate7", "externalRoi7", "externalGoodsOrder15", "externalRgmv15",
		"externalGoodsOrderPrice15", "externalGoodsOrderRate15", "externalRoi15", "externalGoodsOrder30",
		"externalRgmv30", "externalGoodsOrderPrice30", "externalGoodsOrderRate30", "externalRoi30",
		"externalLeads", "externalLeadsCpl", "invokeAppOpenCnt", "invokeAppOpenCost", "invokeAppEnterStoreCnt",
		"invokeAppEnterStoreCost", "invokeAppEngagementCnt", "invokeAppEngagementCost", "invokeAppPaymentCnt",
		"invokeAppPaymentCost", "searchInvokeButtonClickCnt", "searchInvokeButtonClickCost",
	},
	"/jg/data/report/realtime/creativity": {
		"fee", "impression", "click", "ctr", "acp", "cpm", "like", "comment", "collect", "follow", "share",
		"interaction", "cpi", "actionButtonClick", "actionButtonCtr", "screenshot", "picSave", "reservePV",
		"clkLiveEntryPv", "clkLiveEntryPvCost", "clkLiveAvgViewTime", "clkLiveAllFollow", "clkLive5sEntryPv",
		"clkLive5sEntryUvCost", "clkLiveComment", "searchCmtClick", "searchCmtClickCvr", "searchCmtAfterRead",
		"searchCmtAfterReadAvg", "goodsVisit", "goodsVisitPrice", "sellerVisit", "sellerVisitPrice",
		"shoppingCartAdd", "addCartPrice", "presaleOrderNum7d", "presaleOrderGmv7d", "goodsOrder",
		"goodsOrderPrice", "rgmv", "roi", "successGoodsOrder", "clickOrderCvr", "purchaseOrderPrice7d",
		"purchaseOrderGmv7d", "purchaseOrderRoi7d", "clkLiveRoomOrderNum", "liveAverageOrderCost",
		"clkLiveRoomRgmv", "clkLiveRoomRoi", "leads", "leadsCpl", "landingPageVisit", "leadsButtonImpression",
		"validLeads", "validLeadsCpl", "leadsCvr", "phoneCallCnt", "phoneCallSuccCnt", "wechatCopyCnt",
		"wechatCopySuccCnt", "commodityBuyCnt", "messageUser", "message", "messageConsult", "initiativeMessage",
		"messageConsultCpl", "initiativeMessageCpl", "msgLeadsNum", "msgLeadsCost", "externalGoodsVisit24h",
		"externalGoodsVisitRate24h", "externalGoodsOrder7", "externalRgmv7", "externalGoodsOrderPrice7",
		"externalGoodsOrderRate7", "externalRoi7", "externalGoodsOrder15", "externalRgmv15",
		"externalGoodsOrderPrice15", "externalGoodsOrderRate15", "externalRoi15", "externalGoodsOrder30",
		"externalRgmv30", "externalGoodsOrderPrice30", "externalGoodsOrderRate30", "externalRoi30",
		"externalLeads", "externalLeadsCpl", "invokeAppOpenCnt", "invokeAppOpenCost", "invokeAppEnterStoreCnt",
		"invokeAppEnterStoreCost", "invokeAppEngagementCnt", "invokeAppEngagementCost", "invokeAppPaymentCnt",
		"invokeAppPaymentCost", "searchInvokeButtonClickCnt", "searchInvokeButtonClickCost",
	},
	"/jg/data/report/realtime/keyword": {
		"fee", "impression", "click", "ctr", "acp", "cpm", "like", "comment", "collect", "follow", "share",
		"interaction", "cpi", "actionButtonClick", "actionButtonCtr", "screenshot", "picSave", "reservePV",
		"goodsVisit", "goodsVisitPrice", "sellerVisit", "sellerVisitPrice", "shoppingCartAdd", "addCartPrice",
		"presaleOrderNum7d", "presaleOrderGmv7d", "goodsOrder", "goodsOrderPrice", "rgmv", "roi",
		"successGoodsOrder", "clickOrderCvr", "purchaseOrderGmv7d", "purchaseOrderRoi7d", "leads", "leadsCpl",
		"landingPageVisit", "leadsButtonImpression", "validLeads", "validLeadsCpl", "leadsCvr", "phoneCallCnt",
		"phoneCallSuccCnt", "wechatCopyCnt", "wechatCopySuccCnt", "commodityBuyCnt", "messageUser", "message",
		"messageConsult", "initiativeMessage", "messageConsultCpl", "initiativeMessageCpl", "msgLeadsNum",
		"msgLeadsCost", "externalGoodsVisit24h", "externalGoodsVisitRate24h", "externalGoodsOrder7",
		"externalRgmv7", "externalGoodsOrderPrice7", "externalGoodsOrderRate7", "externalRoi7",
		"externalGoodsOrder15", "externalRgmv15", "externalGoodsOrderPrice15", "externalGoodsOrderRate15",
		"externalRoi15", "externalGoodsOrder30", "externalRgmv30", "externalGoodsOrderPrice30",
		"externalGoodsOrderRate30", "externalRoi30", "externalLeads", "externalLeadsCpl",
	},
	"/jg/data/report/realtime/unit": {
		"fee", "impression", "click", "ctr", "acp", "cpm", "like", "comment", "collect", "follow", "share",
		"interaction", "cpi", "actionButtonClick", "actionButtonCtr", "screenshot", "picSave", "reservePV",
		"clkLiveEntryPv", "clkLiveEntryPvCost", "clkLiveAvgViewTime", "clkLiveAllFollow", "clkLive5sEntryPv",
		"clkLive5sEntryUvCost", "clkLiveComment", "searchCmtClick", "searchCmtClickCvr", "searchCmtAfterRead",
		"searchCmtAfterReadAvg", "goodsVisit", "goodsVisitPrice", "sellerVisit", "sellerVisitPrice",
		"shoppingCartAdd", "addCartPrice", "presaleOrderNum7d", "presaleOrderGmv7d", "goodsOrder",
		"goodsOrderPrice", "rgmv", "roi", "successGoodsOrder", "clickOrderCvr", "purchaseOrderPrice7d",
		"purchaseOrderGmv7d", "purchaseOrderRoi7d", "clkLiveRoomOrderNum", "liveAverageOrderCost",
		"clkLiveRoomRgmv", "clkLiveRoomRoi", "leads", "leadsCpl", "landingPageVisit", "leadsButtonImpression",
		"validLeads", "validLeadsCpl", "leadsCvr", "phoneCallCnt", "phoneCallSuccCnt", "wechatCopyCnt",
		"wechatCopySuccCnt", "commodityBuyCnt", "messageUser", "message", "messageConsult", "initiativeMessage",
		"messageConsultCpl", "initiativeMessageCpl", "msgLeadsNum", "msgLeadsCost", "externalGoodsVisit24h",
		"externalGoodsVisitRate24h", "externalGoodsOrder7", "externalRgmv7", "externalGoodsOrderPrice7",
		"externalGoodsOrderRate7", "externalRoi7", "externalGoodsOrder15", "externalRgmv15",
		"externalGoodsOrderPrice15", "externalGoodsOrderRate15", "externalRoi15", "externalGoodsOrder30",
		"externalRgmv30", "externalGoodsOrderPrice30", "externalGoodsOrderRate30", "externalRoi30",
		"externalLeads", "externalLeadsCpl", "invokeAppOpenCnt", "invokeAppOpenCost", "invokeAppEnterStoreCnt",
		"invokeAppEnterStoreCost", "invokeAppEngagementCnt", "invokeAppEngagementCost", "invokeAppPaymentCnt",
		"invokeAppPaymentCost", "searchInvokeButtonClickCnt", "searchInvokeButtonClickCost",
	},
}

// Columns 返回该报表口（如 "/jg/data/report/offline/campaign"）全部可查指标列的副本；
// 不是标准报表口返回 nil（调用方据此不传 columns）。
func Columns(gateway string) []string {
	cols, ok := reportColumns[gateway]
	if !ok {
		return nil
	}
	return append([]string(nil), cols...)
}

// splitUnsupported 带 split_columns（按性别/年龄/地域…细分）时上游不认的列（离线账户层实测，
// 六个维度一致）：应用唤起一组不支持细分，混进去整请求 10002。
var splitUnsupported = map[string]bool{
	"invokeAppOpenCnt": true, "invokeAppOpenCost": true,
	"invokeAppEnterStoreCnt": true, "invokeAppEnterStoreCost": true,
	"invokeAppEngagementCnt": true, "invokeAppEngagementCost": true,
	"invokeAppPaymentCnt": true, "invokeAppPaymentCost": true,
}

// SplitColumns 同 Columns，但去掉带 split_columns 时不支持的列。
func SplitColumns(gateway string) []string {
	cols := Columns(gateway)
	out := cols[:0]
	for _, c := range cols {
		if !splitUnsupported[c] {
			out = append(out, c)
		}
	}
	if len(out) == 0 {
		return nil
	}
	return out
}
