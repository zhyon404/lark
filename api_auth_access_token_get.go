// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetAccessToken 通过此接口获取登录预授权码 code 对应的登录用户身份。
//
// 该接口仅适用于通过网页应用登录方式获取的预授权码，小程序登录中用户身份的获取，请使用[小程序 code2session 接口](/ssl:ttdoc/uYjL24iN/ukjM04SOyQjL5IDN)
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uEDO4UjLxgDO14SM4gTN
func (r *AuthService) GetAccessToken(ctx context.Context, request *GetAccessTokenReq, options ...MethodOptionFunc) (*GetAccessTokenResp, *Response, error) {
	if r.cli.mock.mockAuthGetAccessToken != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Auth#GetAccessToken mock enable")
		return r.cli.mock.mockAuthGetAccessToken(ctx, request, options...)
	}

	r.cli.log(ctx, LogLevelInfo, "[lark] Auth#GetAccessToken call api")
	r.cli.log(ctx, LogLevelDebug, "[lark] Auth#GetAccessToken request: %s", jsonString(request))

	req := &RawRequestReq{
		Method:       "POST",
		URL:          "https://open.feishu.cn/open-apis/authen/v1/access_token",
		Body:         request,
		MethodOption: newMethodOption(options),

		NeedAppAccessToken:  true,
		NeedUserAccessToken: true,
	}
	resp := new(getAccessTokenResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	requestID, statusCode := getResponseRequestID(response)
	if err != nil {
		r.cli.log(ctx, LogLevelError, "[lark] Auth#GetAccessToken POST https://open.feishu.cn/open-apis/authen/v1/access_token failed, request_id: %s, status_code: %d, error: %s", requestID, statusCode, err)
		return nil, response, err
	} else if resp.Code != 0 {
		r.cli.log(ctx, LogLevelError, "[lark] Auth#GetAccessToken POST https://open.feishu.cn/open-apis/authen/v1/access_token failed, request_id: %s, status_code: %d, code: %d, msg: %s", requestID, statusCode, resp.Code, resp.Msg)
		return nil, response, NewError("Auth", "GetAccessToken", resp.Code, resp.Msg)
	}

	r.cli.log(ctx, LogLevelDebug, "[lark] Auth#GetAccessToken success, request_id: %s, status_code: %d, response: %s", requestID, statusCode, jsonString(resp.Data))

	return resp.Data, response, nil
}

func (r *Mock) MockAuthGetAccessToken(f func(ctx context.Context, request *GetAccessTokenReq, options ...MethodOptionFunc) (*GetAccessTokenResp, *Response, error)) {
	r.mockAuthGetAccessToken = f
}

func (r *Mock) UnMockAuthGetAccessToken() {
	r.mockAuthGetAccessToken = nil
}

type GetAccessTokenReq struct {
	GrantType string `json:"grant_type,omitempty"` // 在本流程中，此值为 authorization_code
	Code      string `json:"code,omitempty"`       // 来自[请求身份验证(新)](/ssl:ttdoc/ukTMukTMukTM/ukzN4UjL5cDO14SO3gTN)流程，用户扫码登录后会自动302到redirect_uri并带上此参数
}

type getAccessTokenResp struct {
	Code int64               `json:"code,omitempty"` // 返回码，0表示请求成功，其他表示请求失败
	Msg  string              `json:"msg,omitempty"`  // 返回信息
	Data *GetAccessTokenResp `json:"data,omitempty"` // 返回业务数据
}

type GetAccessTokenResp struct {
	AccessToken      string `json:"access_token,omitempty"`       // user_access_token，用于获取用户资源
	AvatarURL        string `json:"avatar_url,omitempty"`         // 用户头像
	AvatarThumb      string `json:"avatar_thumb,omitempty"`       // 用户头像 72x72
	AvatarMiddle     string `json:"avatar_middle,omitempty"`      // 用户头像 240x240
	AvatarBig        string `json:"avatar_big,omitempty"`         // 用户头像 640x640
	ExpiresIn        int64  `json:"expires_in,omitempty"`         // access_token 的有效期，单位: 秒
	Name             string `json:"name,omitempty"`               // 用户姓名
	EnName           string `json:"en_name,omitempty"`            // 用户英文名称
	OpenID           string `json:"open_id,omitempty"`            // 用户在应用内的唯一标识
	UnionID          string `json:"union_id,omitempty"`           // 用户统一ID
	Email            string `json:"email,omitempty"`              // 申请了"获取用户邮箱"权限的应用返回该字段
	UserID           string `json:"user_id,omitempty"`            // 申请了"获取用户 user_id"权限的应用返回该字段
	Mobile           string `json:"mobile,omitempty"`             // 申请了"获取用户手机号"权限的应用返回该字段
	TenantKey        string `json:"tenant_key,omitempty"`         // 当前企业标识
	RefreshExpiresIn int64  `json:"refresh_expires_in,omitempty"` // refresh_token 的有效期，单位: 秒
	RefreshToken     string `json:"refresh_token,omitempty"`      // 刷新用户 access_token 时使用的 token
	TokenType        string `json:"token_type,omitempty"`         // 此处为 Bearer
}