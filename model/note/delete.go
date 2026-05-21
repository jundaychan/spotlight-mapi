package note

import (
	"github.com/jundaychan/spotlight-mapi/model"
	"github.com/jundaychan/spotlight-mapi/util"
)

// DeleteRequest 删除笔记 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// NoteID 笔记ID
	NoteID string `json:"note_id,omitempty"`
}

// Encode implement PostRequest interface
func (r DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}

// DeleteResponse 删除笔记 API Response
type DeleteResponse struct {
	model.BaseResponse
}
