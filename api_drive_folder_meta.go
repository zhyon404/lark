// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetDriveFolderMeta
//
// **该篇文档有两个接口，一个接口用于获取指定文件夹的元信息，一个用于获取我的空间文件夹的元信息**
// 该接口用于根据 folderToken 获取该文件夹的元信息。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uAjNzUjLwYzM14CM2MTN
func (r *DriveService) GetDriveFolderMeta(ctx context.Context, request *GetDriveFolderMetaReq, options ...MethodOptionFunc) (*GetDriveFolderMetaResp, *Response, error) {
	if r.cli.mock.mockDriveGetDriveFolderMeta != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Drive#GetDriveFolderMeta mock enable")
		return r.cli.mock.mockDriveGetDriveFolderMeta(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:               "Drive",
		API:                 "GetDriveFolderMeta",
		Method:              "GET",
		URL:                 "https://open.feishu.cn/open-apis/drive/explorer/v2/folder/:folderToken/meta",
		Body:                request,
		MethodOption:        newMethodOption(options),
		NeedUserAccessToken: true,
	}
	resp := new(getDriveFolderMetaResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockDriveGetDriveFolderMeta(f func(ctx context.Context, request *GetDriveFolderMetaReq, options ...MethodOptionFunc) (*GetDriveFolderMetaResp, *Response, error)) {
	r.mockDriveGetDriveFolderMeta = f
}

func (r *Mock) UnMockDriveGetDriveFolderMeta() {
	r.mockDriveGetDriveFolderMeta = nil
}

type GetDriveFolderMetaReq struct {
	FolderToken string `path:"folderToken" json:"-"` // 文件夹的 token，获取方式见 [对接前说明](https://open.feishu.cn/document/ukTMukTMukTM/uczNzUjL3czM14yN3MTN) 的第 3 项
}

type getDriveFolderMetaResp struct {
	Code int64                   `json:"code,omitempty"`
	Msg  string                  `json:"msg,omitempty"`
	Data *GetDriveFolderMetaResp `json:"data,omitempty"`
}

type GetDriveFolderMetaResp struct {
	ID        string `json:"id,omitempty"`        // 文件夹的 id
	Name      string `json:"name,omitempty"`      // 文件夹的标题
	Token     string `json:"token,omitempty"`     // 文件夹的 token
	CreateUid string `json:"createUid,omitempty"` // 文件夹的创建者 id
	EditUid   string `json:"editUid,omitempty"`   // 文件夹的最后编辑者 id
	ParentID  string `json:"parentId,omitempty"`  // 文件夹的上级目录 id
	OwnUid    string `json:"ownUid,omitempty"`    // 文件夹为个人文件夹时，为文件夹的所有者 id；文件夹为共享文件夹时，为文件夹树id
}
