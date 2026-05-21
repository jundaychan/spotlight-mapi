package note

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// QueryIDRequest 获取笔记ID API Request
type QueryIDRequest struct {
	// TemporaryID 临时笔记ID
	TemporaryID string `json:"temporary_id,omitempty"`
}

// Encode implement PostRequest interface
func (r QueryIDRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// QueryIDResponse 获取笔记ID API Response
type QueryIDResponse struct {
	model.BaseResponse
	Data *QueryIDResult `json:"data,omitempty"`
}

// QueryIDResult 获取笔记ID结果
type QueryIDResult struct {
	// TemporaryID 笔记临时ID
	TemporaryID string `json:"temporary_id,omitempty"`
	// NoteID 笔记ID
	NoteID string `json:"note_id,omitempty"`
	// CreateTime 创建时间
	CreateTime int64 `json:"create_time,omitempty"`
}
