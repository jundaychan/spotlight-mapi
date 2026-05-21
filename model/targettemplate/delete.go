package targettemplate

import "github.com/bububa/spotlight-mapi/util"

// DeleteRequest 删除定向包 API Request
type DeleteRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// TemplateIDs 定向包id列表，最多5个
	TemplateIDs []uint64 `json:"template_ids,omitempty"`
}

// Encode implement PostRequest interface
func (r DeleteRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
