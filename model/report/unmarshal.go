package report

import (
	"bytes"
	"encoding/json"
	"unicode"
)

// UnmarshalJSON 兼容两种指标键风格。
//
// 标准投的离线/实时报表回 snake_case（message_consult / msg_leads_num），本结构体的 tag 按它写；
// 而简单投的 ube/* 实时口（/jg/data/report/realtime/ube/{group,campaign,note,keyword}）回 camelCase
// （messageConsult / msgLeadsNum）。不做归一，后者除 fee/impression/click 这类单词键外全部解成 0，
// 且不报错——进线/开口/留资整列静默归零。
//
// 键里出现大写字母才走归一路径；纯 snake 报文直接解码，不多花一次 map 往返。
func (d *DataReportDTO) UnmarshalJSON(b []byte) error {
	type plain DataReportDTO
	if !bytes.ContainsAny(b, "ABCDEFGHIJKLMNOPQRSTUVWXYZ") {
		return json.Unmarshal(b, (*plain)(d))
	}
	var raw map[string]json.RawMessage
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	norm := make(map[string]json.RawMessage, len(raw))
	for k, v := range raw {
		sk := CamelToSnake(k)
		// 同一指标两种写法都在时以 snake 为准（那是本结构体的原生键）
		if _, dup := norm[sk]; dup && sk != k {
			continue
		}
		norm[sk] = v
	}
	nb, err := json.Marshal(norm)
	if err != nil {
		return err
	}
	return json.Unmarshal(nb, (*plain)(d))
}

// CamelToSnake 把 camelCase 指标键转成本结构体使用的 snake_case tag。
//
// 规则对齐 tag 的既有写法（clk_live_5s_entry_pv / presale_order_num_7d / current_app_pay_roi）：
//   - 大写字母前加下划线，但连续大写视为一个词（ROI → roi）
//   - 数字段前加下划线（Num7d → num_7d），数字后紧跟的小写字母不再分词（5s / 7d / 24h 保持整体）
func CamelToSnake(s string) string {
	rs := []rune(s)
	out := make([]rune, 0, len(rs)+8)
	for i, r := range rs {
		switch {
		case unicode.IsUpper(r):
			if i > 0 {
				prev := rs[i-1]
				nextLower := i+1 < len(rs) && unicode.IsLower(rs[i+1])
				if !unicode.IsUpper(prev) || nextLower {
					out = append(out, '_')
				}
			}
			out = append(out, unicode.ToLower(r))
		case unicode.IsDigit(r):
			if i > 0 && unicode.IsLetter(rs[i-1]) {
				out = append(out, '_')
			}
			out = append(out, r)
		default:
			out = append(out, r)
		}
	}
	return string(out)
}
