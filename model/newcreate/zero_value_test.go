package newcreate

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/jundaychan/spotlight-mapi/util"
)

// TestCreateRequestKeepsMeaningfulZeros 回归：新创编里 0 是合法枚举值的字段必须真的发得出去。
//
// 改成指针之前它们是裸 int + omitempty，下面这一整批全被 json 静默丢掉——
// 建计划时「点击量 / 手动投放 / 不限预算 / 长期投放 / 搜索快投关闭 / 手动选词 /
// 搜索定向 / 无组件 / 精确匹配」写了等于没写，聚光按自己的默认值建（其中
// search_flag、keyword_gen_type、phrase_match_type_upgrade 的默认是 -1，
// phrase_match_type 的默认是 1-短语匹配，都不是 0）。
func TestCreateRequestKeepsMeaningfulZeros(t *testing.T) {
	r := CreateRequest{
		AdvertiserID: 123, CreateType: 1,
		CreateCascadeInfoList: []CreateCascadeInfo{{
			Campaign: &CreateCampaign{
				CampaignName: "回归用例", MarketingTarget: 4, Placement: 2,
				DeliveryMode: util.Ptr(0), OptimizeObjective: util.Ptr(0), TimeType: util.Ptr(0),
				TimePeriodType: util.Ptr(0), LimitDayBudget: util.Ptr(0), SmartSwitch: util.Ptr(0),
				ExploreState: util.Ptr(0), SearchFlag: util.Ptr(0), AgreedRedStarFee: util.Ptr(0),
			},
			UnitWithCreativeList: []CreateUnitWithCreative{{
				Unit: &CreateUnit{
					UnitName: "u", TargetType: util.Ptr(0), KeywordGenType: util.Ptr(0),
					KeywordWithBid: []KeywordWithBidDTO{{Keyword: "医美", Bid: 100, PhraseMatchType: util.Ptr(0)}},
					TargetInfo:     &CreateTargetInfo{SearchTargetCityIntent: util.Ptr(0)},
				},
				CreativityList: []CreateCreativity{{
					CreativityName: "c", NoteID: "n1",
					ConversionType: util.Ptr(0), ComponentConvNumIsShow: util.Ptr(false),
				}},
			}},
		}},
	}
	got := string(r.Encode())
	for _, want := range []string{
		`"delivery_mode":0`, `"optimize_objective":0`, `"time_type":0`, `"time_period_type":0`,
		`"limit_day_budget":0`, `"smart_switch":0`, `"explore_state":0`, `"search_flag":0`,
		`"agreed_red_star_fee":0`, `"target_type":0`, `"keyword_gen_type":0`,
		`"phrase_match_type":0`, `"search_target_city_intent":0`,
		`"conversion_type":0`, `"component_conv_num_is_show":false`,
	} {
		if !strings.Contains(got, want) {
			t.Errorf("报文里缺 %s\n实际: %s", want, got)
		}
	}
}

// TestCreateRequestOmitsUnsetPointers nil 仍然不进报文——「不传」和「显式传 0」必须分得开，
// 否则每个没设的开关都会被当成显式关闭发给聚光。
func TestCreateRequestOmitsUnsetPointers(t *testing.T) {
	r := CreateRequest{
		AdvertiserID: 1,
		CreateCascadeInfoList: []CreateCascadeInfo{{
			Campaign: &CreateCampaign{CampaignName: "x", MarketingTarget: 4, Placement: 2},
		}},
	}
	got := string(r.Encode())
	for _, absent := range []string{
		"delivery_mode", "optimize_objective", "search_flag", "agreed_red_star_fee",
		"smart_switch", "explore_state", "horse_race",
	} {
		if strings.Contains(got, absent) {
			t.Errorf("未设置的 %s 不该进报文\n实际: %s", absent, got)
		}
	}
}

// TestUpdateRequestKeepsMeaningfulZeros 编辑接口同理：Field Mask 的全部意义就是
// 「把字段改回默认值/清空」，mask 里写了 limit_day_budget 而报文里发不出 0，等于白写。
func TestUpdateRequestKeepsMeaningfulZeros(t *testing.T) {
	r := UpdateRequest{
		AdvertiserID: 1, ModifyType: 1,
		ModCascadeInfoList: []ModCascadeInfo{{
			Campaign: &ModifyCampaign{
				CampaignID:   9,
				TimeType:     util.Ptr(0),
				ExploreState: util.Ptr(0),
				SearchFlag:   util.Ptr(0),
				UpdateFields: []string{"time_type", "explore_state", "search_flag", "limit_day_budget"},
			},
		}},
	}
	got := string(r.Encode())
	for _, want := range []string{`"time_type":0`, `"explore_state":0`, `"search_flag":0`} {
		if !strings.Contains(got, want) {
			t.Errorf("报文里缺 %s\n实际: %s", want, got)
		}
	}
}

// TestCreateRequestRoundTripsFrontendZeros 前端把整张报文按 4877 拼好发给后端，
// 后端反序列化进本结构体再转发。显式的 0 必须一路活到聚光，不能在中转时蒸发。
func TestCreateRequestRoundTripsFrontendZeros(t *testing.T) {
	raw := []byte(`{"advertiser_id":1,"create_type":1,"create_cascade_info_list":[{"campaign":{
		"campaign_name":"前端拼的","marketing_target":4,"placement":2,
		"optimize_objective":0,"delivery_mode":0,"limit_day_budget":0,"search_flag":0}}]}`)
	var r CreateRequest
	if err := json.Unmarshal(raw, &r); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	got := string(r.Encode())
	for _, want := range []string{`"optimize_objective":0`, `"delivery_mode":0`, `"limit_day_budget":0`, `"search_flag":0`} {
		if !strings.Contains(got, want) {
			t.Errorf("中转后丢了 %s\n实际: %s", want, got)
		}
	}
}
