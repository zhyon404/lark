// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdateHelpdeskTicketCustomizedField
//
// 该接口用于更新自定义字段。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/helpdesk-v1/ticket_customized_field/update-ticket-customized-field
func (r *HelpdeskService) UpdateHelpdeskTicketCustomizedField(ctx context.Context, request *UpdateHelpdeskTicketCustomizedFieldReq, options ...MethodOptionFunc) (*UpdateHelpdeskTicketCustomizedFieldResp, *Response, error) {
	if r.cli.mock.mockHelpdeskUpdateHelpdeskTicketCustomizedField != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Helpdesk#UpdateHelpdeskTicketCustomizedField mock enable")
		return r.cli.mock.mockHelpdeskUpdateHelpdeskTicketCustomizedField(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Helpdesk",
		API:                 "UpdateHelpdeskTicketCustomizedField",
		Method:              "PATCH",
		URL:                 "https://open.feishu.cn/open-apis/helpdesk/v1/ticket_customized_fields/:ticket_customized_field_id",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
		NeedHelpdeskAuth:    true,
	}
	resp := new(updateHelpdeskTicketCustomizedFieldResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockHelpdeskUpdateHelpdeskTicketCustomizedField(f func(ctx context.Context, request *UpdateHelpdeskTicketCustomizedFieldReq, options ...MethodOptionFunc) (*UpdateHelpdeskTicketCustomizedFieldResp, *Response, error)) {
	r.mockHelpdeskUpdateHelpdeskTicketCustomizedField = f
}

func (r *Mock) UnMockHelpdeskUpdateHelpdeskTicketCustomizedField() {
	r.mockHelpdeskUpdateHelpdeskTicketCustomizedField = nil
}

type UpdateHelpdeskTicketCustomizedFieldReq struct {
	TicketCustomizedFieldID string                  `path:"ticket_customized_field_id" json:"-"` // 工单自定义字段ID, 示例值："6948728206392295444"
	DisplayName             *string                 `json:"display_name,omitempty"`              // 名称, 示例值："test dropdown"
	Position                *string                 `json:"position,omitempty"`                  // 字段在列表后台管理列表中的位置, 示例值："3"
	Description             *string                 `json:"description,omitempty"`               // 描述, 示例值："下拉示例"
	Visible                 *bool                   `json:"visible,omitempty"`                   // 是否可见, 示例值：true
	Required                *bool                   `json:"required,omitempty"`                  // 是否必填, 示例值：false
	DropdownOptions         *HelpdeskDropdownOption `json:"dropdown_options,omitempty"`          // 下拉列表选项
}

type updateHelpdeskTicketCustomizedFieldResp struct {
	Code int64                                    `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                                   `json:"msg,omitempty"`  // 错误描述
	Data *UpdateHelpdeskTicketCustomizedFieldResp `json:"data,omitempty"`
}

type UpdateHelpdeskTicketCustomizedFieldResp struct{}
