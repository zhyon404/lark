// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// GetAttendanceStatisticsData
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//task/query-statistics-data
func (r *AttendanceService) GetAttendanceStatisticsData(ctx context.Context, request *GetAttendanceStatisticsDataReq, options ...MethodOptionFunc) (*GetAttendanceStatisticsDataResp, *Response, error) {
	if r.cli.mock.mockAttendanceGetAttendanceStatisticsData != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#GetAttendanceStatisticsData mock enable")
		return r.cli.mock.mockAttendanceGetAttendanceStatisticsData(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "GetAttendanceStatisticsData",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/user_stats_datas/query",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(getAttendanceStatisticsDataResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockAttendanceGetAttendanceStatisticsData(f func(ctx context.Context, request *GetAttendanceStatisticsDataReq, options ...MethodOptionFunc) (*GetAttendanceStatisticsDataResp, *Response, error)) {
	r.mockAttendanceGetAttendanceStatisticsData = f
}

func (r *Mock) UnMockAttendanceGetAttendanceStatisticsData() {
	r.mockAttendanceGetAttendanceStatisticsData = nil
}

type GetAttendanceStatisticsDataReq struct {
	EmployeeType     EmployeeType `query:"employee_type" json:"-"`      // 用户 ID 类型, 可选值有: `employee_id`, `employee_no`
	Locale           string       `json:"locale,omitempty"`             // 语言类型, 可选值有: `en`：英文, `ja`：日文, `zh`：中文
	StatsType        string       `json:"stats_type,omitempty"`         // 统计类型,      , 可选值有: `daily`：日度统计, `month`：月度统计
	StartDate        int64        `json:"start_date,omitempty"`         // 开始时间, 示例值：20210316
	EndDate          int64        `json:"end_date,omitempty"`           // 结束时间, 示例值：20210323,      ,      （时间间隔不超过 40 天）
	UserIDs          []string     `json:"user_ids,omitempty"`           // 查询的用户 ID 列表,      ,      （用户数量不超过 20）
	NeedHistory      *bool        `json:"need_history,omitempty"`       // 是否包含历史数据, 示例值：true
	CurrentGroupOnly *bool        `json:"current_group_only,omitempty"` // 是否只包含当前考勤组, 示例值：true
}

type getAttendanceStatisticsDataResp struct {
	Code int64                            `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                           `json:"msg,omitempty"`  // 错误描述
	Data *GetAttendanceStatisticsDataResp `json:"data,omitempty"`
}

type GetAttendanceStatisticsDataResp struct {
	UserDatas []*GetAttendanceStatisticsDataRespUserData `json:"user_datas,omitempty"` // 用户统计数据
}

type GetAttendanceStatisticsDataRespUserData struct {
	Name   string                                         `json:"name,omitempty"`    // 姓名
	UserID string                                         `json:"user_id,omitempty"` // 用户 ID
	Datas  []*GetAttendanceStatisticsDataRespUserDataData `json:"datas,omitempty"`   // 数据
}

type GetAttendanceStatisticsDataRespUserDataData struct {
	Code     string                                                `json:"code,omitempty"`     // 字段编号
	Value    string                                                `json:"value,omitempty"`    // 数据值
	Features []*GetAttendanceStatisticsDataRespUserDataDataFeature `json:"features,omitempty"` // 数据属性
}

type GetAttendanceStatisticsDataRespUserDataDataFeature struct {
	Key   string `json:"key,omitempty"`   // 属性名
	Value string `json:"value,omitempty"` // 属性值
}
