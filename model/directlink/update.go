package directlink

import "github.com/jundaychan/spotlight-mapi/util"

// UpdateRequest 编辑直达链接 API Request
type UpdateRequest struct {
	// AdvertiserID 广告主id
	AdvertiserID uint64 `json:"advertiser_id,omitempty"`
	// DirectLink 直达链接信息
	DirectLink *UpdateDirectLink `json:"direct_link,omitempty"`
}

// UpdateDirectLink 编辑直达链接信息
type UpdateDirectLink struct {
	// ID 直达链接的id
	ID uint64 `json:"id,omitempty"`
	// URL url内容
	URL string `json:"url,omitempty"`
	// RemarkName 备注名称
	RemarkName string `json:"remark_name,omitempty"`
}

// Encode implement PostRequest interface
func (r UpdateRequest) Encode() []byte {
	return util.JSONMarshal(r)
}
