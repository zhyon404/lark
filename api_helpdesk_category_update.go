// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdateHelpdeskCategory 该接口用于更新知识库分类详情。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/category/patch
func (r *HelpdeskService) UpdateHelpdeskCategory(ctx context.Context, request *UpdateHelpdeskCategoryReq, options ...MethodOptionFunc) (*UpdateHelpdeskCategoryResp, *Response, error) {
	if r.cli.mock.mockHelpdeskUpdateHelpdeskCategory != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#UpdateHelpdeskCategory mock enable")
		return r.cli.mock.mockHelpdeskUpdateHelpdeskCategory(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Helpdesk",
		API:                 "UpdateHelpdeskCategory",
		Method:              "PATCH",
		URL:                 "https://open.feishu.cn/open-apis/helpdesk/v1/categories/:id",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
		NeedHelpdeskAuth:    true,
	}
	resp := new(updateHelpdeskCategoryResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockHelpdeskUpdateHelpdeskCategory(f func(ctx context.Context, request *UpdateHelpdeskCategoryReq, options ...MethodOptionFunc) (*UpdateHelpdeskCategoryResp, *Response, error)) {
	r.mockHelpdeskUpdateHelpdeskCategory = f
}

func (r *Mock) UnMockHelpdeskUpdateHelpdeskCategory() {
	r.mockHelpdeskUpdateHelpdeskCategory = nil
}

type UpdateHelpdeskCategoryReq struct {
	ID       string  `path:"id" json:"-"`         // category id, 示例值："6948728206392295444"
	Name     *string `json:"name,omitempty"`      // 名称, 示例值："创建团队和邀请成员"
	ParentID *string `json:"parent_id,omitempty"` // 父知识库分类ID, 示例值："0"
}

type updateHelpdeskCategoryResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *UpdateHelpdeskCategoryResp `json:"data,omitempty"`
}

type UpdateHelpdeskCategoryResp struct{}
