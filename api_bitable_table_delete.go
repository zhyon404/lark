// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// DeleteBitableTable 删除一个数据表
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/bitable-v1/app-table/delete
func (r *BitableService) DeleteBitableTable(ctx context.Context, request *DeleteBitableTableReq, options ...MethodOptionFunc) (*DeleteBitableTableResp, *Response, error) {
	if r.cli.mock.mockBitableDeleteBitableTable != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#DeleteBitableTable mock enable")
		return r.cli.mock.mockBitableDeleteBitableTable(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Bitable#DeleteBitableTable call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#DeleteBitableTable request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:       "DELETE",
		URL:          "https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id",
		Body:         request,
		MethodOption: newMethodOption(options),

		NeedUserAccessToken: true,
	}
	resp := new(deleteBitableTableResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Bitable#DeleteBitableTable DELETE https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Bitable#DeleteBitableTable DELETE https://open.feishu.cn/open-apis/bitable/v1/apps/:app_token/tables/:table_id failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Bitable", "DeleteBitableTable", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Bitable#DeleteBitableTable success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockBitableDeleteBitableTable(f func(ctx context.Context, request *DeleteBitableTableReq, options ...MethodOptionFunc) (*DeleteBitableTableResp, *Response, error)) {
	r.mockBitableDeleteBitableTable = f
}

func (r *Mock) UnMockBitableDeleteBitableTable() {
	r.mockBitableDeleteBitableTable = nil
}

type DeleteBitableTableReq struct {
	AppToken string `path:"app_token" json:"-"` // bitable app token, 示例值："appbcbWCzen6D8dezhoCH2RpMAh"
	TableID  string `path:"table_id" json:"-"`  // table id, 示例值："tblsRc9GRRXKqhvW"
}

type deleteBitableTableResp struct {
	Code int64                   `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                  `json:"msg,omitempty"`  // 错误描述
	Data *DeleteBitableTableResp `json:"data,omitempty"`
}

type DeleteBitableTableResp struct{}