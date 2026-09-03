package spu

import "encoding/json"

// UnmarshalJSON 容错解析 SPU 列表。
//
// 【为什么需要】聚光文档写的字段是 `spu` / `total`，但线上 /jg/spu/list 实际返回的是
// `spus` / `total_count`（2026-09-03 实测：某医美主体 can_bind=true 时上游
// `{"data":{"total_count":12,"spus":[...]}}`）。按文档的 tag 解析会静默得到 0 条，
// 表现为前端「没有可关联的 SPU」——明明后台有 SPU 却一个都选不了，且没有任何报错。
// 这里两种拼写都认，以实际返回为准。
func (r *ListResult) UnmarshalJSON(b []byte) error {
	var raw struct {
		Spu        []Spu `json:"spu"`
		Spus       []Spu `json:"spus"`
		Page       int64 `json:"page"`
		PageSize   int64 `json:"page_size"`
		Total      int64 `json:"total"`
		TotalCount int64 `json:"total_count"`
	}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	r.List = raw.Spu
	if len(r.List) == 0 {
		r.List = raw.Spus
	}
	r.Page = raw.Page
	r.PageSize = raw.PageSize
	r.Total = raw.Total
	if r.Total == 0 {
		r.Total = raw.TotalCount
	}
	return nil
}
