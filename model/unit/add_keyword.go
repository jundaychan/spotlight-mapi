package unit

import (
	"github.com/jundaychan/spotlight-mapi/enum"
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// AddKeywordRequest 修改单元关键词 API Request
type AddKeywordRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// UnitID 单元id
	UnitID uint64 `json:"unit_id,omitempty"`
	// AddType 添加关键词的方式
	// 0-如果已有词不包括本次添加的词，则将原来已有词删除、1-只是进行添加操作，不删除 默认值0
	AddType int `json:"add_type,omitempty"`
	// PhraseMatchTypeUpgrade 关键词匹配方式升级开关(精确匹配->精确包含)
	// 0-关闭、1-开启、默认值-1
	PhraseMatchTypeUpgrade int `json:"phrase_match_type_upgrade,omitempty"`
	// KeywordWithBid 关键词列表
	KeywordWithBid []AddKeywordItem `json:"keyword_with_bid,omitempty"`
}

// AddKeywordItem 关键词出价信息
type AddKeywordItem struct {
	// Keyword 关键词
	Keyword string `json:"keyword,omitempty"`
	// Bid 出价
	Bid int64 `json:"bid,omitempty"`
	// PhraseMatchType 关键词匹配方式 0-精确匹配、1-短语匹配
	PhraseMatchType enum.PhraseMatchType `json:"phrase_match_type,omitempty"`
}

// Encode implement PostRequest interface
func (r AddKeywordRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// AddKeywordResponse 修改单元关键词 API Response
type AddKeywordResponse struct {
	model.BaseResponse
	Data *AddKeywordResult `json:"data,omitempty"`
}

// AddKeywordResult 修改单元关键词结果
type AddKeywordResult struct {
	// UploadKeywordNum 上传成功的关键词的个数
	UploadKeywordNum int `json:"upload_keyword_num,omitempty"`
}
