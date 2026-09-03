package util

// Ptr 返回 v 的指针。
//
// 用于请求结构体里「0 / false 是合法枚举值」的字段：这类字段一律 *T + omitempty，
//
//	nil     → 不进报文，聚光按自己的默认值处理
//	Ptr(0)  → 显式发 0，表达「就是要这个枚举值」
//
// 之所以必须区分：聚光有一批字段的默认值**不是 0**——search_flag / keyword_gen_type /
// phrase_match_type_upgrade 默认 -1，phrase_match_type 默认 1-短语匹配。用裸 int +
// omitempty 时零值被 json 静默丢掉，代码里写着「精确匹配」，投出去是短语匹配。
func Ptr[T any](v T) *T { return &v }
