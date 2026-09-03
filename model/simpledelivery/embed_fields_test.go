package simpledelivery

import (
	"encoding/json"
	"testing"
)

// TestSimpleDeliveryReportsKeepDimensions 简单投三个报表是「维度结构体 + DataReportDTO」
// 双内嵌：DataReportDTO 的 UnmarshalJSON 被提升后，**维度那层也会被一起吃掉**
// （维度结构体自己没有 UnmarshalJSON，但它的字段同样解不出来）。
func TestSimpleDeliveryReportsKeepDimensions(t *testing.T) {
	t.Run("NoteReport", func(t *testing.T) {
		var r NoteReport
		mustJSON(t, `{"time":"2026-09-01","note_id":"n1","note_title":"标题",
			"campaign_id":"77","fee":"5.5"}`, &r)
		if r.Time != "2026-09-01" || r.NoteID != "n1" || r.NoteTitle != "标题" || r.CampaignID.Value() != 77 {
			t.Errorf("维度丢了: %+v", r.NoteReportDimension)
		}
		if r.Fee.Value() != 5.5 {
			t.Errorf("fee=%v", r.Fee.Value())
		}
	})

	t.Run("CampaignReport", func(t *testing.T) {
		var r CampaignReport
		mustJSON(t, `{"time":"2026-09-01","campaign_id":"88","campaign_name":"计划A","fee":"6.6"}`, &r)
		if r.Time != "2026-09-01" || r.CampaignID.Value() != 88 || r.CampaignName != "计划A" {
			t.Errorf("维度丢了: %+v", r.CampaignReportDimension)
		}
		if r.Fee.Value() != 6.6 {
			t.Errorf("fee=%v", r.Fee.Value())
		}
	})

	t.Run("TargetReport", func(t *testing.T) {
		var r TargetReport
		mustJSON(t, `{"time":"2026-09-01","ube_view_id":"u1","campaign_group_name":"标的A","fee":"7.7"}`, &r)
		if r.Time != "2026-09-01" || r.UbeViewID != "u1" || r.CampaignGroupName != "标的A" {
			t.Errorf("维度丢了: %+v", r.TargetReportDimension)
		}
		if r.Fee.Value() != 7.7 {
			t.Errorf("fee=%v", r.Fee.Value())
		}
	})
}

// TestSimpleDeliveryCamelCase 简单投(ube/*)正是返回 camelCase 指标键的那条链路——
// DataReportDTO 的兼容逻辑必须继续生效。
func TestSimpleDeliveryCamelCase(t *testing.T) {
	var r NoteReport
	mustJSON(t, `{"time":"2026-09-01","note_id":"n1","fee":"5.5","actionButtonCtr":"0.25"}`, &r)
	if r.NoteID != "n1" || r.Fee.Value() != 5.5 {
		t.Errorf("noteId=%q fee=%v", r.NoteID, r.Fee.Value())
	}
	if r.ActionButtonCtr.Value() != 0.25 {
		t.Errorf("camelCase 指标键没解出来: %v", r.ActionButtonCtr.Value())
	}
}

func mustJSON(t *testing.T, raw string, v any) {
	t.Helper()
	if err := json.Unmarshal([]byte(raw), v); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
}
