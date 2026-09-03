package spu

import (
	"encoding/json"
	"testing"
)

// 线上真实返回用的是 spus/total_count，文档写的是 spu/total，两种都要认。
func TestListResultUnmarshalBothSpellings(t *testing.T) {
	cases := []struct {
		name      string
		body      string
		wantLen   int
		wantTotal int64
		wantFirst string
	}{
		{
			"线上实际返回 spus/total_count",
			`{"total_count":12,"spus":[{"spu_id":1768616,"spu_name":"涓雅玻尿酸填充内轮廓（上海）"},{"spu_id":2,"spu_name":"B"}]}`,
			2, 12, "涓雅玻尿酸填充内轮廓（上海）",
		},
		{
			"文档口径 spu/total",
			`{"total":5,"spu":[{"spu_id":9,"spu_name":"文档口径"}]}`,
			1, 5, "文档口径",
		},
		{"空列表", `{"total_count":0,"spus":[]}`, 0, 0, ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var r ListResult
			if err := json.Unmarshal([]byte(c.body), &r); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if len(r.List) != c.wantLen {
				t.Fatalf("len(List)=%d want %d", len(r.List), c.wantLen)
			}
			if r.Total != c.wantTotal {
				t.Fatalf("Total=%d want %d", r.Total, c.wantTotal)
			}
			if c.wantFirst != "" && r.List[0].SpuName != c.wantFirst {
				t.Fatalf("List[0].SpuName=%q want %q", r.List[0].SpuName, c.wantFirst)
			}
		})
	}
}
