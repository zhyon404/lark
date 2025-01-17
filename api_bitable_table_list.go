// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetBitableTableList 根据  app_token，获取多维表格下的所有数据表
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/list
func (r *BitableService) GetBitableTableList(ctx context.Context, request *GetBitableTableListReq, options ...MethodOptionFunc) (*GetBitableTableListResp, *Response, error) {
	if r.cli.mock.mockBitableGetBitableTableList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#GetBitableTableList mock enable")
		return r.cli.mock.mockBitableGetBitableTableList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Bitable",
		API:                   "GetBitableTableList",
		Method:                "GET",
		URL:                   "https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(getBitableTableListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockBitableGetBitableTableList(f func(ctx context.Context, request *GetBitableTableListReq, options ...MethodOptionFunc) (*GetBitableTableListResp, *Response, error)) {
	r.mockBitableGetBitableTableList = f
}

func (r *Mock) UnMockBitableGetBitableTableList() {
	r.mockBitableGetBitableTableList = nil
}

type GetBitableTableListReq struct {
	PageToken *string `query:"page_token" json:"-"` // 分页标记，第一次请求不填，表示从头开始遍历；分页查询结果还有更多项时会同时返回新的 page_token，下次遍历可采用该 page_token 获取查询结果, 示例值："tblsRc9GRRXKqhvW"
	PageSize  *int64  `query:"page_size" json:"-"`  // 分页大小, 示例值：10, 最大值：`100`
	AppToken  string  `path:"app_token" json:"-"`   // bitable app token, 示例值："appbcbWCzen6D8dezhoCH2RpMAh"
}

type getBitableTableListResp struct {
	Code int64                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                   `json:"msg,omitempty"`  // 错误描述
	Data *GetBitableTableListResp `json:"data,omitempty"`
}

type GetBitableTableListResp struct {
	HasMore   bool                           `json:"has_more,omitempty"`   // 是否还有更多项
	PageToken string                         `json:"page_token,omitempty"` // 分页标记，当 has_more 为 true 时，会同时返回新的 page_token，否则不返回 page_token
	Items     []*GetBitableTableListRespItem `json:"items,omitempty"`      // 数据表信息
}

type GetBitableTableListRespItem struct {
	TableID  string `json:"table_id,omitempty"` // 表格表 id
	Revision int64  `json:"revision,omitempty"` // 数据表的版本号
	Name     string `json:"name,omitempty"`     // 数据表 名字
}
