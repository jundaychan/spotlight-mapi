package report

import (
	"encoding/json"
	"testing"
)

// TestEmbeddedUnmarshalerSwallowsOuterFields 把这个 Go 陷阱本身钉在这里：
// 匿名内嵌一个自带 UnmarshalJSON 的类型（DataReportDTO），标准 json.Unmarshal
// 只会走内嵌类型的方法，**外层自己声明的字段被静默留空**——不报错，值是空的。
// 这条用例断言的就是"坏行为仍然存在"，一旦哪天 Go 或 DataReportDTO 改了语义，
// 它会失败并提醒我们 UnmarshalEmbedded 这层补丁可以撤了。
func TestEmbeddedUnmarshalerSwallowsOuterFields(t *testing.T) {
	type Outer struct {
		NoteID string `json:"note_id,omitempty"`
		DataReportDTO
	}
	var o Outer
	if err := json.Unmarshal([]byte(`{"note_id":"abc123","fee":"1.89"}`), &o); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if o.NoteID != "" {
		t.Fatalf("陷阱不复现了？NoteID=%q——若 Go/DataReportDTO 语义已变，UnmarshalEmbedded 可以撤了", o.NoteID)
	}
	if o.Fee.Value() != 1.89 {
		t.Errorf("内嵌层本身应该正常: Fee=%v", o.Fee.Value())
	}
}

// TestUnmarshalEmbedded 修复后：外层字段与内嵌指标两边都在。
func TestUnmarshalEmbedded(t *testing.T) {
	type Fixed struct {
		NoteID string `json:"note_id,omitempty"`
		DataReportDTO
	}
	var f Fixed
	if err := UnmarshalEmbedded([]byte(`{"note_id":"abc123","fee":"1.89","impression":"10"}`), &f); err != nil {
		t.Fatalf("UnmarshalEmbedded: %v", err)
	}
	if f.NoteID != "abc123" {
		t.Errorf("NoteID = %q, want abc123", f.NoteID)
	}
	if f.Fee.Value() != 1.89 || f.Impression.Value() != 10 {
		t.Errorf("内嵌指标丢了: Fee=%v Impression=%v", f.Fee.Value(), f.Impression.Value())
	}
}

// TestUnmarshalEmbeddedKeepsCamelCase 内嵌层的 camelCase 兼容（ube/* 用）必须继续生效——
// 它正是 DataReportDTO 当初要自定义 UnmarshalJSON 的原因，不能因为这层补丁失效。
func TestUnmarshalEmbeddedKeepsCamelCase(t *testing.T) {
	type Fixed struct {
		NoteID string `json:"note_id,omitempty"`
		DataReportDTO
	}
	var f Fixed
	if err := UnmarshalEmbedded([]byte(`{"note_id":"n1","fee":"2.5","actionButtonCtr":"0.12"}`), &f); err != nil {
		t.Fatalf("UnmarshalEmbedded: %v", err)
	}
	if f.NoteID != "n1" || f.Fee.Value() != 2.5 {
		t.Errorf("NoteID=%q Fee=%v", f.NoteID, f.Fee.Value())
	}
	if f.ActionButtonCtr.Value() != 0.12 {
		t.Errorf("camelCase 键 actionButtonCtr 没解出来: %v", f.ActionButtonCtr.Value())
	}
}

// TestUnmarshalEmbeddedFlattensPlainEmbeds 双内嵌（维度结构体 + DataReportDTO）——
// simpledelivery 的三个报表就是这个形状。维度字段在普通内嵌里，一样会被吃掉，
// 摊平后必须解得出来。
func TestUnmarshalEmbeddedFlattensPlainEmbeds(t *testing.T) {
	type Dim struct {
		Time string `json:"time,omitempty"`
		Name string `json:"name,omitempty"`
	}
	type Fixed struct {
		Dim
		DataReportDTO
	}
	var f Fixed
	if err := UnmarshalEmbedded([]byte(`{"time":"2026-09-01","name":"计划A","fee":"3.5"}`), &f); err != nil {
		t.Fatalf("UnmarshalEmbedded: %v", err)
	}
	if f.Time != "2026-09-01" || f.Name != "计划A" {
		t.Errorf("普通内嵌的维度字段丢了: Time=%q Name=%q", f.Time, f.Name)
	}
	if f.Fee.Value() != 3.5 {
		t.Errorf("Fee=%v", f.Fee.Value())
	}
}

// TestUnmarshalEmbeddedShallowWins 同名字段按 encoding/json 的规则：层级浅的赢。
func TestUnmarshalEmbeddedShallowWins(t *testing.T) {
	type Dim struct {
		Time string `json:"time,omitempty"`
	}
	type Fixed struct {
		Time string `json:"time,omitempty"` // 浅层，应该赢
		Dim
		DataReportDTO
	}
	var f Fixed
	if err := UnmarshalEmbedded([]byte(`{"time":"2026-09-01"}`), &f); err != nil {
		t.Fatalf("UnmarshalEmbedded: %v", err)
	}
	if f.Time != "2026-09-01" {
		t.Errorf("浅层 Time 没填上: %q", f.Time)
	}
	if f.Dim.Time != "" {
		t.Errorf("深层 Time 不该被填: %q", f.Dim.Time)
	}
}

// TestUnmarshalEmbeddedKeepsAbsentFields 报文里没出现的字段保持原值，与标准库一致。
func TestUnmarshalEmbeddedKeepsAbsentFields(t *testing.T) {
	type Fixed struct {
		NoteID string `json:"note_id,omitempty"`
		Title  string `json:"title,omitempty"`
		DataReportDTO
	}
	f := Fixed{NoteID: "旧值", Title: "旧标题"}
	if err := UnmarshalEmbedded([]byte(`{"note_id":"新值"}`), &f); err != nil {
		t.Fatalf("UnmarshalEmbedded: %v", err)
	}
	if f.NoteID != "新值" {
		t.Errorf("NoteID = %q", f.NoteID)
	}
	if f.Title != "旧标题" {
		t.Errorf("报文里没有的 Title 不该被清空: %q", f.Title)
	}
}

// TestUnmarshalEmbeddedRejectsNonPointer 传错类型要报错，别静默什么都不做。
func TestUnmarshalEmbeddedRejectsNonPointer(t *testing.T) {
	type Fixed struct{ DataReportDTO }
	if err := UnmarshalEmbedded([]byte(`{}`), Fixed{}); err == nil {
		t.Error("传非指针应该报错")
	}
	var p *Fixed
	if err := UnmarshalEmbedded([]byte(`{}`), p); err == nil {
		t.Error("传 nil 指针应该报错")
	}
}
