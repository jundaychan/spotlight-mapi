package reportoffline2

import (
	"encoding/json"
	"testing"
)

// TestNoteReportUnmarshalKeepsOwnFields 复现并验证修复：NoteReport 匿名内嵌
// report.DataReportDTO（自定义 UnmarshalJSON）时，note_id/note_title/time 等
// NoteReport 自己声明的字段不会被静默清空。
func TestNoteReportUnmarshalKeepsOwnFields(t *testing.T) {
	raw := []byte(`{
		"note_id": "682c4de7000000002300e845",
		"note_title": "鼻基底凹陷⁉️ 看完这篇直接上岸‼️",
		"note_image": "http://ci.xiaohongshu.com/spectrum/1040g0k0320m00c4c6o005pgokd41oju5v86chb8?imageView2/2/w/1080/format/jpg",
		"time": "2026-08-28",
		"fee": "2766.11",
		"impression": "5870",
		"click": "444",
		"msg_leads_num": "1",
		"message_consult": "4"
	}`)
	var r NoteReport
	if err := json.Unmarshal(raw, &r); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if r.NoteID != "682c4de7000000002300e845" {
		t.Errorf("NoteID = %q, want note_id", r.NoteID)
	}
	if r.NoteTitle == "" {
		t.Error("NoteTitle 丢了")
	}
	if r.Time != "2026-08-28" {
		t.Errorf("Time = %q, want 2026-08-28", r.Time)
	}
	if r.Fee.Value() != 2766.11 {
		t.Errorf("Fee = %v, want 2766.11 (来自内嵌 DataReportDTO，必须继续正常工作)", r.Fee.Value())
	}
	if r.MsgLeadsNum.Value() != 1 {
		t.Errorf("MsgLeadsNum = %v, want 1", r.MsgLeadsNum.Value())
	}
}

// TestNoteReportListUnmarshalMultiple 确认多条记录(data_list 数组)场景也正确——
// 生产代码是批量解析整个响应体，不是单条。
func TestNoteReportListUnmarshalMultiple(t *testing.T) {
	raw := []byte(`{"data_list":[
		{"note_id":"a1","fee":"10.5"},
		{"note_id":"a2","fee":"20.5"}
	]}`)
	var list NoteReportList
	if err := json.Unmarshal(raw, &list); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if len(list.List) != 2 {
		t.Fatalf("List 条数 = %d, want 2", len(list.List))
	}
	if list.List[0].NoteID != "a1" || list.List[1].NoteID != "a2" {
		t.Errorf("NoteID 顺序/值不对: %q, %q", list.List[0].NoteID, list.List[1].NoteID)
	}
}
