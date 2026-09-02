package report

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"
	"unicode"
)

// snakeToCamel 按上游习惯把 tag 还原成 camelCase（数字开头的段原样拼接：5s / 7d / 24h）。
func snakeToCamel(s string) string {
	parts := strings.Split(s, "_")
	var b strings.Builder
	for i, p := range parts {
		if p == "" {
			continue
		}
		if i == 0 || !unicode.IsLetter(rune(p[0])) {
			b.WriteString(p)
			continue
		}
		b.WriteString(strings.ToUpper(p[:1]) + p[1:])
	}
	return b.String()
}

// TestCamelToSnakeCoversEveryTag 结构体里每一个 tag 从 camelCase 转回来都必须命中原 tag，
// 否则那个指标在 ube/* 实时口上会静默为 0。
func TestCamelToSnakeCoversEveryTag(t *testing.T) {
	typ := reflect.TypeOf(DataReportDTO{})
	for i := 0; i < typ.NumField(); i++ {
		tag := strings.Split(typ.Field(i).Tag.Get("json"), ",")[0]
		if tag == "" || tag == "-" {
			continue
		}
		camel := snakeToCamel(tag)
		if got := CamelToSnake(camel); got != tag {
			t.Errorf("%s: camel=%s → %s (want %s)", typ.Field(i).Name, camel, got, tag)
		}
	}
}

func TestCamelToSnakeUpstreamSpellings(t *testing.T) {
	cases := map[string]string{
		"messageConsult":         "message_consult",
		"initiativeMessage":      "initiative_message",
		"msgLeadsNum":            "msg_leads_num",
		"goodsOrder":             "goods_order",
		"clkLive5sEntryPv":       "clk_live_5s_entry_pv",
		"presaleOrderNum7d":      "presale_order_num_7d",
		"externalGoodsOrder7":    "external_goods_order_7",
		"currentAppPayROI":       "current_app_pay_roi",
		"appActivateAmount1dROI": "app_activate_amount_1d_roi",
		"fee":                    "fee",
		"already_snake":          "already_snake",
	}
	for in, want := range cases {
		if got := CamelToSnake(in); got != want {
			t.Errorf("%s → %s, want %s", in, got, want)
		}
	}
}

func TestDataReportDTOUnmarshalBothStyles(t *testing.T) {
	camel := `{"fee":"364.07","impression":"5655","click":"203","interaction":"3","messageConsult":"10",
	  "initiativeMessage":"7","msgLeadsNum":"9","leads":"1","goodsOrder":"2","rgmv":"88.5","ctr":"3.59%","id":"3762674"}`
	snake := `{"fee":364.07,"impression":5655,"click":203,"interaction":3,"message_consult":10,
	  "initiative_message":7,"msg_leads_num":9,"leads":1,"goods_order":2,"rgmv":88.5,"ctr":"3.59%"}`
	for name, src := range map[string]string{"camel": camel, "snake": snake} {
		var d DataReportDTO
		if err := json.Unmarshal([]byte(src), &d); err != nil {
			t.Fatalf("%s: %v", name, err)
		}
		if d.Fee.Value() != 364.07 || d.Impression.Value() != 5655 || d.Click.Value() != 203 ||
			d.MessageConsult.Value() != 10 || d.InitiativeMessage.Value() != 7 || d.MsgLeadsNum.Value() != 9 ||
			d.Leads.Value() != 1 || d.GoodsOrder.Value() != 2 || d.Rgmv.Value() != 88.5 {
			t.Errorf("%s: 解析结果不对 %+v", name, d)
		}
	}
	// 嵌在上层结构里也要生效（实时响应的 data / total_data）
	var wrap struct {
		Data *DataReportDTO `json:"data"`
	}
	if err := json.Unmarshal([]byte(`{"data":`+camel+`}`), &wrap); err != nil || wrap.Data.MsgLeadsNum.Value() != 9 {
		t.Fatalf("嵌套解析失败 err=%v data=%+v", err, wrap.Data)
	}
}
