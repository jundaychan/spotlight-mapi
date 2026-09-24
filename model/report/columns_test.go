package report

import (
	"reflect"
	"strings"
	"testing"
)

// TestColumnsMapBackToDTO 每个列都必须能落回 DataReportDTO 的某个字段（CamelToSnake 命中 tag），
// 否则查回来的值没人接；同一层不许重复列。
func TestColumnsMapBackToDTO(t *testing.T) {
	tags := map[string]bool{}
	typ := reflect.TypeOf(DataReportDTO{})
	for i := 0; i < typ.NumField(); i++ {
		tags[strings.Split(typ.Field(i).Tag.Get("json"), ",")[0]] = true
	}
	for gw, cols := range reportColumns {
		seen := map[string]bool{}
		for _, c := range cols {
			if !tags[CamelToSnake(c)] {
				t.Errorf("%s: 列 %s → %s 不是 DataReportDTO 的字段", gw, c, CamelToSnake(c))
			}
			if seen[c] {
				t.Errorf("%s: 列 %s 重复", gw, c)
			}
			seen[c] = true
		}
		if len(cols) < 50 {
			t.Errorf("%s: 只有 %d 列，生成出错？", gw, len(cols))
		}
	}
	if Columns("/jg/data/report/offline/note") != nil {
		t.Error("非标准报表口应返回 nil")
	}
	c := Columns("/jg/data/report/offline/campaign")
	c[0] = "x"
	if reportColumns["/jg/data/report/offline/campaign"][0] == "x" {
		t.Error("Columns 必须返回副本")
	}
}
