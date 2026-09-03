package newcreate

// 新创编里有一批「0 是合法枚举值」的字段（优化目标 0=点击量、投放模式 0=手动、
// 预算类型 0=不限、组件类型 0=无组件、关键词匹配 0=精确匹配……）。这些字段一律用
// 指针 + omitempty：
//
//	nil  → 不进报文，聚光按它自己的默认值处理（search_flag 是 -1、keyword_gen_type 是 -1
//	       这种「默认≠0」的字段，正是靠这个区分出来的）
//	&0   → 显式发 0，表达「就是要这个枚举值」
//
// 改成指针之前它们是裸 int + omitempty，零值被 json 静默丢掉：建计划时
// 「点击量 / 手动投放 / 不限预算 / 精确匹配」全部发不出去，聚光按默认值建，
// 而调用方代码里明明写着 0。见 zero_value_test.go。

// Int 返回 v 的指针，用于显式传 0 这类有语义的零值。
func Int(v int) *int { return &v }

// Int64 返回 v 的指针。
func Int64(v int64) *int64 { return &v }

// Bool 返回 v 的指针，用于显式传 false。
func Bool(v bool) *bool { return &v }
