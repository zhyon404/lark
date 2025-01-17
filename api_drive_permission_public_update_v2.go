// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// UpdateDrivePublicPermissionV2 该接口用于根据 filetoken 更新文档的公共设置。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uITN5UjLyUTO14iM1kTN
func (r *DriveService) UpdateDrivePublicPermissionV2(ctx context.Context, request *UpdateDrivePublicPermissionV2Req, options ...MethodOptionFunc) (*UpdateDrivePublicPermissionV2Resp, *Response, error) {
	if r.cli.mock.mockDriveUpdateDrivePublicPermissionV2 != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#UpdateDrivePublicPermissionV2 mock enable")
		return r.cli.mock.mockDriveUpdateDrivePublicPermissionV2(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Drive",
		API:                   "UpdateDrivePublicPermissionV2",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/drive/permission/v2/public/update/",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
		NeedUserAccessToken:   true,
	}
	resp := new(updateDrivePublicPermissionV2Resp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveUpdateDrivePublicPermissionV2(f func(ctx context.Context, request *UpdateDrivePublicPermissionV2Req, options ...MethodOptionFunc) (*UpdateDrivePublicPermissionV2Resp, *Response, error)) {
	r.mockDriveUpdateDrivePublicPermissionV2 = f
}

func (r *Mock) UnMockDriveUpdateDrivePublicPermissionV2() {
	r.mockDriveUpdateDrivePublicPermissionV2 = nil
}

type UpdateDrivePublicPermissionV2Req struct {
	Token           string  `json:"token,omitempty"`             // 文件的 token，获取方式见 [对接前说明](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN)的第 4 项
	Type            string  `json:"type,omitempty"`              // 文档类型  "doc"  or  "sheet" or "file"
	SecurityEntity  *string `json:"security_entity,omitempty"`   // 可创建副本/打印/导出/复制设置（不传则保持原值）：<br>"anyone_can_view" - 所有可访问此文档的用户<br>"anyone_can_edit" - 有编辑权限的用户
	CommentEntity   *string `json:"comment_entity,omitempty"`    // 可评论设置（不传则保持原值）：<br>"anyone_can_view" - 所有可访问此文档的用户<br>"anyone_can_edit" - 有编辑权限的用户
	ShareEntity     *string `json:"share_entity,omitempty"`      // 谁可以添加和管理协作者（不传则保持原值）：<br>"anyone"-所有可阅读或编辑此文档的用户<br>"same_tenant"-组织内所有可阅读或编辑此文档的用户<br>"off"-只有我可以
	LinkShareEntity *string `json:"link_share_entity,omitempty"` // 链接共享（不传则保持原值）：<br>"off" - 关闭链接分享<br>"tenant_readable" - 组织内获得链接的人可阅读<br>"tenant_editable" - 组织内获得链接的人可编辑<br>"anyone_readable" - 获得链接的任何人可阅读<br>"anyone_editable" - 获得链接的任何人可编辑
	ExternalAccess  *bool   `json:"external_access,omitempty"`   // 是否允许分享到租户外开关（不传则保持原值）
	InviteExternal  *bool   `json:"invite_external,omitempty"`   // 非owner是否允许邀请外部人（不传则保持原值）
}

type updateDrivePublicPermissionV2Resp struct {
	Code int64                              `json:"code,omitempty"`
	Msg  string                             `json:"msg,omitempty"`
	Data *UpdateDrivePublicPermissionV2Resp `json:"data,omitempty"`
}

type UpdateDrivePublicPermissionV2Resp struct{}
