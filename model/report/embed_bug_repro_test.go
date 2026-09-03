package report

import (
	"encoding/json"
	"testing"
)

// TestEmbedNoteIDBug 复现：外层结构体有自己的 tagged 字段 + 匿名内嵌一个带自定义
// UnmarshalJSON 的类型时，外层字段会不会被吃掉。
func TestEmbedNoteIDBug(t *testing.T) {
	type Outer struct {
		NoteID string `json:"note_id,omitempty"`
		DataReportDTO
	}
	raw := `{"note_id":"abc123","fee":"1.89","impression":"10"}`
	var o Outer
	if err := json.Unmarshal([]byte(raw), &o); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	t.Logf("NoteID=%q Fee=%v Impression=%v", o.NoteID, o.Fee.Value(), o.Impression.Value())
	if o.NoteID != "abc123" {
		t.Errorf("NoteID 丢了: got %q, want %q", o.NoteID, "abc123")
	}
}

// TestEmbedNoteIDBugFix 验证修复思路：分两次 Unmarshal（一次进内嵌类型，一次进外层专属字段的镜像 struct）。
func TestEmbedNoteIDBugFix(t *testing.T) {
	type Outer struct {
		NoteID string `json:"note_id,omitempty"`
		DataReportDTO
	}
	raw := []byte(`{"note_id":"abc123","fee":"1.89","impression":"10"}`)
	var o Outer
	if err := json.Unmarshal(raw, &o.DataReportDTO); err != nil {
		t.Fatalf("unmarshal DataReportDTO: %v", err)
	}
	type outerOnly struct {
		NoteID string `json:"note_id,omitempty"`
	}
	var oo outerOnly
	if err := json.Unmarshal(raw, &oo); err != nil {
		t.Fatalf("unmarshal outerOnly: %v", err)
	}
	o.NoteID = oo.NoteID
	t.Logf("修复后: NoteID=%q Fee=%v Impression=%v", o.NoteID, o.Fee.Value(), o.Impression.Value())
	if o.NoteID != "abc123" || o.Fee.Value() != 1.89 {
		t.Fatalf("修复无效: NoteID=%q Fee=%v", o.NoteID, o.Fee.Value())
	}
}
