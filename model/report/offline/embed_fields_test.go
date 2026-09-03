package offline

import (
	"encoding/json"
	"testing"
)

// TestOfflineReportKeepsDimension Report 是「Dimension + DataReportDTO」双内嵌，
// 维度字段（时间 / 计划 / 单元 / 创意）会被内嵌层的 UnmarshalJSON 一起吃掉。
func TestOfflineReportKeepsDimension(t *testing.T) {
	var r Report
	raw := `{"time":"2026-09-01","campaign_id":"11","campaign_name":"计划A",
		"unit_id":"22","unit_name":"单元A","creativity_id":"33","fee":"8.8"}`
	if err := json.Unmarshal([]byte(raw), &r); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if r.Time != "2026-09-01" || r.CampaignID.Value() != 11 || r.CampaignName != "计划A" ||
		r.UnitID.Value() != 22 || r.UnitName != "单元A" || r.CreativityID.Value() != 33 {
		t.Errorf("维度丢了: %+v", r.Dimension)
	}
	if r.Fee.Value() != 8.8 {
		t.Errorf("fee=%v", r.Fee.Value())
	}
}
