package keyword

import (
	"context"

	"github.com/jundaychan/spotlight-mapi/core"
	"github.com/jundaychan/spotlight-mapi/model/keyword"
)

// IndustryTaxonomyAttribute 行业类目属性
func IndustryTaxonomyAttribute(ctx context.Context, clt *core.SDKClient, req *keyword.IndustryTaxonomyAttributeRequest, accessToken string) (*keyword.IndustryTaxonomyAttributeResult, error) {
	var resp keyword.IndustryTaxonomyAttributeResponse
	if err := clt.Post(ctx, "/jg/keyword/industry/taxonomy/attribute", req, &resp, accessToken); err != nil {
		return nil, err
	}
	return resp.Data, nil
}
