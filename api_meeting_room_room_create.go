// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateMeetingRoomRoom 该接口用于创建会议室。
//
// doc: https://open.feishu.cn/document/ukTMukTMukTM/uITNwYjLyUDM24iM1AjN
func (r *MeetingRoomService) CreateMeetingRoomRoom(ctx context.Context, request *CreateMeetingRoomRoomReq, options ...MethodOptionFunc) (*CreateMeetingRoomRoomResp, *Response, error) {
	if r.cli.mock.mockMeetingRoomCreateMeetingRoomRoom != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] MeetingRoom#CreateMeetingRoomRoom mock enable")
		return r.cli.mock.mockMeetingRoomCreateMeetingRoomRoom(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "MeetingRoom",
		API:                   "CreateMeetingRoomRoom",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/meeting_room/room/create",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createMeetingRoomRoomResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockMeetingRoomCreateMeetingRoomRoom(f func(ctx context.Context, request *CreateMeetingRoomRoomReq, options ...MethodOptionFunc) (*CreateMeetingRoomRoomResp, *Response, error)) {
	r.mockMeetingRoomCreateMeetingRoomRoom = f
}

func (r *Mock) UnMockMeetingRoomCreateMeetingRoomRoom() {
	r.mockMeetingRoomCreateMeetingRoomRoom = nil
}

type CreateMeetingRoomRoomReq struct {
	BuildingID   string  `json:"building_id,omitempty"`    // 会议室所在的建筑ID
	Floor        string  `json:"floor,omitempty"`          // 会议室所在的建筑楼层
	Name         string  `json:"name,omitempty"`           // 会议室名称
	Capacity     int64   `json:"capacity,omitempty"`       // 容量
	IsDisabled   bool    `json:"is_disabled,omitempty"`    // 是否禁用
	CustomRoomID *string `json:"custom_room_id,omitempty"` // 租户自定义会议室ID
}

type createMeetingRoomRoomResp struct {
	Code int64                      `json:"code,omitempty"` // 返回码，非 0 表示失败
	Msg  string                     `json:"msg,omitempty"`  // 返回码的描述，"success" 表示成功，其他为错误提示信息
	Data *CreateMeetingRoomRoomResp `json:"data,omitempty"` // 返回业务信息
}

type CreateMeetingRoomRoomResp struct {
	RoomID string `json:"room_id,omitempty"` // 成功创建的会议室ID
}
