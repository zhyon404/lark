// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdateSearchDataSource 更新一个已经存在的数据源
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/search-v2/data_source/patch
func (r *SearchService) UpdateSearchDataSource(ctx context.Context, request *UpdateSearchDataSourceReq, options ...MethodOptionFunc) (*UpdateSearchDataSourceResp, *Response, error) {
	if r.cli.mock.mockSearchUpdateSearchDataSource != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Search#UpdateSearchDataSource mock enable")
		return r.cli.mock.mockSearchUpdateSearchDataSource(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Search",
		API:                   "UpdateSearchDataSource",
		Method:                "PATCH",
		URL:                   "https://open.feishu.cn/open-apis/search/v2/data_sources/:data_source_id",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(updateSearchDataSourceResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockSearchUpdateSearchDataSource(f func(ctx context.Context, request *UpdateSearchDataSourceReq, options ...MethodOptionFunc) (*UpdateSearchDataSourceResp, *Response, error)) {
	r.mockSearchUpdateSearchDataSource = f
}

func (r *Mock) UnMockSearchUpdateSearchDataSource() {
	r.mockSearchUpdateSearchDataSource = nil
}

type UpdateSearchDataSourceReq struct {
	DataSourceID string  `path:"data_source_id" json:"-"` // 数据源的唯一标识, 示例值："service_ticket"
	Name         *string `json:"name,omitempty"`          // 数据源的展示名称, 示例值："客服工单"
	State        *int64  `json:"state,omitempty"`         // 数据源状态，0-未上线，1-已上线, 示例值：0, 可选值有: `0`：未上线, `1`：已上线
	Description  *string `json:"description,omitempty"`   // 对于数据源的描述, 示例值："搜索客服工单"
}

type updateSearchDataSourceResp struct {
	Code int64                       `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                      `json:"msg,omitempty"`  // 错误描述
	Data *UpdateSearchDataSourceResp `json:"data,omitempty"` //
}

type UpdateSearchDataSourceResp struct {
	DataSource *UpdateSearchDataSourceRespDataSource `json:"data_source,omitempty"` // 数据源
}

type UpdateSearchDataSourceRespDataSource struct {
	ID            string `json:"id,omitempty"`              // 数据源的唯一标识
	Name          string `json:"name,omitempty"`            // data_source的展示名称
	State         int64  `json:"state,omitempty"`           // 数据源状态，0-未上线，1-已上线, 可选值有: `0`：未上线, `1`：已上线
	Description   string `json:"description,omitempty"`     // 对于数据源的描述
	CreateTime    string `json:"create_time,omitempty"`     // 创建时间，使用Unix时间戳，秒
	UpdateTime    string `json:"update_time,omitempty"`     // 更新时间，使用Unix时间戳，秒
	IsExceedQuota bool   `json:"is_exceed_quota,omitempty"` // 是否超限
}