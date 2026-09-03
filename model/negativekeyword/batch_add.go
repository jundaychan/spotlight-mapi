package negativekeyword

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// BatchAddRequest 批量添加否定词 API Request
type BatchAddRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UnitID 单元ID
	UnitID uint64 `json:"unit_id,omitempty"`
	// Keywords 否定词列表，同一个单元下精确匹配和短语匹配的否定词的个数分别最多为200个
	Keywords []NegativeKeywordAddItemDTO `json:"keywords,omitempty"`
}

// NegativeKeywordAddItemDTO 否定词添加项
type NegativeKeywordAddItemDTO struct {
	// Keyword 否定词
	Keyword string `json:"keyword,omitempty"`
	// PhraseMatchType 匹配方式，0-精确匹配，1-短语匹配。
	// 指针：0 必须显式传，否则否词按短语匹配生效，屏蔽范围比预期大
	PhraseMatchType *int `json:"phrase_match_type,omitempty"`
}

// Encode implement PostRequest interface
func (r BatchAddRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// BatchAddResponse 批量添加否定词 API Response
type BatchAddResponse struct {
	model.BaseResponse
}
