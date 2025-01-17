// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetApprovalInstanceList
//
// 根据 approval_code 批量获取审批实例的 instance_code，用于拉取租户下某个审批定义的全部审批实例。
// 默认以审批创建时间排序。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uQDOyUjL0gjM14CN4ITN
func (r *ApprovalService) GetApprovalInstanceList(ctx context.Context, request *GetApprovalInstanceListReq, options ...MethodOptionFunc) (*GetApprovalInstanceListResp, *Response, error) {
	if r.cli.mock.mockApprovalGetApprovalInstanceList != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Approval#GetApprovalInstanceList mock enable")
		return r.cli.mock.mockApprovalGetApprovalInstanceList(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Approval",
		API:                   "GetApprovalInstanceList",
		Method:                "POST",
		URL:                   "https://www.feishu.cn/approval/openapi/v2/instance/list",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getApprovalInstanceListResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockApprovalGetApprovalInstanceList(f func(ctx context.Context, request *GetApprovalInstanceListReq, options ...MethodOptionFunc) (*GetApprovalInstanceListResp, *Response, error)) {
	r.mockApprovalGetApprovalInstanceList = f
}

func (r *Mock) UnMockApprovalGetApprovalInstanceList() {
	r.mockApprovalGetApprovalInstanceList = nil
}

type GetApprovalInstanceListReq struct {
	ApprovalCode string `json:"approval_code,omitempty"` // 审批定义唯一标识
	StartTime    int64  `json:"start_time,omitempty"`    // 审批实例创建时间区间（毫秒）
	EndTime      int64  `json:"end_time,omitempty"`      // 审批实例创建时间区间（毫秒）
	Offset       int64  `json:"offset,omitempty"`        // 查询偏移量
	Limit        int64  `json:"limit,omitempty"`         // 查询限制量 注:不得大于100
}

type getApprovalInstanceListResp struct {
	Code int64                        `json:"code,omitempty"` // 错误码，非0表示失败
	Msg  string                       `json:"msg,omitempty"`  // 返回码的描述
	Data *GetApprovalInstanceListResp `json:"data,omitempty"` // 返回业务信息
}

type GetApprovalInstanceListResp struct {
	InstanceCodeList []string `json:"instance_code_list,omitempty"` // 审批实例 Code
}
