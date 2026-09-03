package reportoffline2

import (
	"encoding/json"
	"testing"
)

// TestOffline2ReportsKeepOwnFields 本包三个报表结构体都匿名内嵌了 report.DataReportDTO，
// 没有自己的 UnmarshalJSON 时维度字段（note_id / search_word / spu_id …）会被静默留空。
// 每加一个内嵌 DataReportDTO 的报表类型，都要在这里补一条。
func TestOffline2ReportsKeepOwnFields(t *testing.T) {
	t.Run("NoteReport", func(t *testing.T) {
		var r NoteReport
		mustJSON(t, `{"note_id":"682c4de7000000002300e845","note_title":"鼻基底凹陷",
			"time":"2026-08-28","fee":"2766.11","msg_leads_num":"1"}`, &r)
		if r.NoteID != "682c4de7000000002300e845" || r.NoteTitle == "" || r.Time != "2026-08-28" {
			t.Errorf("维度字段丢了: id=%q title=%q time=%q", r.NoteID, r.NoteTitle, r.Time)
		}
		if r.Fee.Value() != 2766.11 || r.MsgLeadsNum.Value() != 1 {
			t.Errorf("指标丢了: fee=%v leads=%v", r.Fee.Value(), r.MsgLeadsNum.Value())
		}
	})

	t.Run("SearchWordReport", func(t *testing.T) {
		var r SearchWordReport
		mustJSON(t, `{"search_word":"鼻综合","campaign_id":"123","unit_name":"单元A",
			"note_id":"n1","time":"2026-08-28","fee":"12.5"}`, &r)
		if r.SearchWord != "鼻综合" || r.CampaignID.Value() != 123 || r.UnitName != "单元A" ||
			r.NoteID != "n1" || r.Time != "2026-08-28" {
			t.Errorf("维度字段丢了: %+v", r)
		}
		if r.Fee.Value() != 12.5 {
			t.Errorf("fee=%v", r.Fee.Value())
		}
	})

	t.Run("SPUReport", func(t *testing.T) {
		var r SPUReport
		// spu 的自有字段分布在内嵌层的前后两侧，摊平必须两边都覆盖到
		mustJSON(t, `{"spu_id":"s1","spu_name":"玻尿酸","time":"2026-08-28",
			"fee":"9.9","i_user_num":"42","ti_user_price":"3.3"}`, &r)
		if r.SpuID != "s1" || r.SpuName != "玻尿酸" || r.Time != "2026-08-28" {
			t.Errorf("内嵌层之前的字段丢了: %+v", r)
		}
		if r.IUserNum.Value() != 42 || r.TiUserPrice.Value() != 3.3 {
			t.Errorf("内嵌层之后的字段丢了: i=%v tiPrice=%v", r.IUserNum.Value(), r.TiUserPrice.Value())
		}
		if r.Fee.Value() != 9.9 {
			t.Errorf("fee=%v", r.Fee.Value())
		}
	})
}

func mustJSON(t *testing.T, raw string, v any) {
	t.Helper()
	if err := json.Unmarshal([]byte(raw), v); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
}
