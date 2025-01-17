// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdateSheetConditionFormat 该接口用于更新已有的条件格式，单次最多支持更新10个条件格式，每个条件格式的更新会返回成功或者失败，失败的情况包括各种参数的校验。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/conditionformat/condition-format-update
func (r *DriveService) UpdateSheetConditionFormat(ctx context.Context, request *UpdateSheetConditionFormatReq, options ...MethodOptionFunc) (*UpdateSheetConditionFormatResp, *Response, error) {
	if r.cli.mock.mockDriveUpdateSheetConditionFormat != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#UpdateSheetConditionFormat mock enable")
		return r.cli.mock.mockDriveUpdateSheetConditionFormat(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "UpdateSheetConditionFormat",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/sheets/v2/spreadsheets/:spreadsheetToken/condition_formats/batch_update",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateSheetConditionFormatResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveUpdateSheetConditionFormat(f func(ctx context.Context, request *UpdateSheetConditionFormatReq, options ...MethodOptionFunc) (*UpdateSheetConditionFormatResp, *Response, error)) {
	r.mockDriveUpdateSheetConditionFormat = f
}

func (r *Mock) UnMockDriveUpdateSheetConditionFormat() {
	r.mockDriveUpdateSheetConditionFormat = nil
}

type UpdateSheetConditionFormatReq struct {
	SpreadSheetToken      string                                              `path:"spreadsheetToken" json:"-"`         // sheet 的 token，获取方式见 [在线表格开发指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/overview)
	SheetConditionFormats *UpdateSheetConditionFormatReqSheetConditionFormats `json:"sheet_condition_formats,omitempty"` // 表格的条件格式信息
}

type UpdateSheetConditionFormatReqSheetConditionFormats struct {
	SheetID         string                                                             `json:"sheet_id,omitempty"`         // sheet的id
	ConditionFormat *UpdateSheetConditionFormatReqSheetConditionFormatsConditionFormat `json:"condition_format,omitempty"` // 一个条件格式的详细信息
}

type UpdateSheetConditionFormatReqSheetConditionFormatsConditionFormat struct {
	CfID     string                                                                  `json:"cf_id,omitempty"`     // 需要更新的条件格式id，会校验id是否存在
	Ranges   []string                                                                `json:"ranges,omitempty"`    // 条件格式应用的范围，支持：sheetId（整表）；sheetId!1:2（整行）；sheetId!A:B（整列）；sheetId!A1:B2（普通范围）；sheetId!A1:C（应用至最后一行）。应用范围不能超过表格的行总数和列总数，sheetId要与参数的sheetId一致
	RuleType string                                                                  `json:"rule_type,omitempty"` // 条件格式规则类型，目前只有7种：***containsBlanks（为空）、notContainsBlanks（不为空）、duplicateValues（重复值）、uniqueValues（唯一值）、cellIs（限定值范围）、containsText（包含内容）、timePeriod（日期）***
	Attrs    *UpdateSheetConditionFormatReqSheetConditionFormatsConditionFormatAttrs `json:"attrs,omitempty"`     // rule_type对应的具体属性信息，详见 [条件格式指南](https://open.feishu.cn/document/ukTMukTMukTM/uATMzUjLwEzM14CMxMTN/conditionformat/condition-format-guide)
	Style    *UpdateSheetConditionFormatReqSheetConditionFormatsConditionFormatStyle `json:"style,omitempty"`     // 条件格式样式，只支持以下样式，以下样式每个参数都可选，但是不能设置空的style
}

type UpdateSheetConditionFormatReqSheetConditionFormatsConditionFormatAttrs struct {
	Operator   *string  `json:"operator,omitempty"`    // 操作方法
	TimePeriod *string  `json:"time_period,omitempty"` // 时间范围
	Formula    []string `json:"formula,omitempty"`     // 格式
	Text       *string  `json:"text,omitempty"`        // 文本
}

type UpdateSheetConditionFormatReqSheetConditionFormatsConditionFormatStyle struct {
	Font           *UpdateSheetConditionFormatReqSheetConditionFormatsConditionFormatStyleFont `json:"font,omitempty"`            // 字体样式
	TextDecoration *int64                                                                      `json:"text_decoration,omitempty"` // 文本装饰 ，0 默认，1 下划线，2 删除线 ，3 下划线和删除线
	ForeColor      *string                                                                     `json:"fore_color,omitempty"`      // 字体颜色
	BackColor      *string                                                                     `json:"back_color,omitempty"`      // 背景颜色
}

type UpdateSheetConditionFormatReqSheetConditionFormatsConditionFormatStyleFont struct {
	Bold   *bool `json:"bold,omitempty"`   // 加粗
	Italic *bool `json:"italic,omitempty"` // 斜体
}

type updateSheetConditionFormatResp struct {
	Code int64                           `json:"code,omitempty"`
	Msg  string                          `json:"msg,omitempty"`
	Data *UpdateSheetConditionFormatResp `json:"data,omitempty"`
}

type UpdateSheetConditionFormatResp struct {
	Responses []*UpdateSheetConditionFormatRespResponse `json:"responses,omitempty"` // 响应
}

type UpdateSheetConditionFormatRespResponse struct {
	SheetID string `json:"sheet_id,omitempty"` // sheet的Id
	CfID    string `json:"cf_id,omitempty"`    // 更新的条件格式id
	ResCode int64  `json:"res_code,omitempty"` // 条件格式更新状态码，0表示成功，非0表示失败
	ResMsg  string `json:"res_msg,omitempty"`  // 条件格式更新返回的状态信息，空表示成功，不空表示失败原因
}
