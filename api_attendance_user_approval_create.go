// Code generated by lark_sdk_gen. DO NOT EDIT.

package lark

import (
	"context"
)

// CreateAttendanceUserApproval
//
// 由于部分企业使用的是自己的审批系统，而不是飞书审批系统，因此员工的请假、加班等数据无法流入到飞书考勤系统中，导致员工在请假时间段内依然收到打卡提醒，并且被记为缺卡。
// 对于这些只使用飞书考勤系统，而未使用飞书审批系统的企业，可以通过考勤开放接口的形式，将三方审批结果数据回写到飞书的考勤系统中。
// 目前支持加班、请假、出差和外出这四种审批结果的写入。
//
// doc: https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/Attendance//AddApprovalsInLarkAttendance
func (r *AttendanceService) CreateAttendanceUserApproval(ctx context.Context, request *CreateAttendanceUserApprovalReq, options ...MethodOptionFunc) (*CreateAttendanceUserApprovalResp, *Response, error) {
	if r.cli.mock.mockAttendanceCreateAttendanceUserApproval != nil {
		r.cli.log(ctx, LogLevelDebug, "[lark] Attendance#CreateAttendanceUserApproval mock enable")
		return r.cli.mock.mockAttendanceCreateAttendanceUserApproval(ctx, request, options...)
	}

	req := &RawRequestReq{
		Scope:                 "Attendance",
		API:                   "CreateAttendanceUserApproval",
		Method:                "POST",
		URL:                   "https://open.feishu.cn/open-apis/attendance/v1/user_approvals",
		Body:                  request,
		MethodOption:          newMethodOption(options),
		NeedTenantAccessToken: true,
	}
	resp := new(createAttendanceUserApprovalResp)

	response, err := r.cli.RawRequest(ctx, req, resp)
	return resp.Data, response, err
}

func (r *Mock) MockAttendanceCreateAttendanceUserApproval(f func(ctx context.Context, request *CreateAttendanceUserApprovalReq, options ...MethodOptionFunc) (*CreateAttendanceUserApprovalResp, *Response, error)) {
	r.mockAttendanceCreateAttendanceUserApproval = f
}

func (r *Mock) UnMockAttendanceCreateAttendanceUserApproval() {
	r.mockAttendanceCreateAttendanceUserApproval = nil
}

type CreateAttendanceUserApprovalReq struct {
	EmployeeType EmployeeType                                 `query:"employee_type" json:"-"` // 请求体中的 user_id 的员工工号类型，必选字段，可用值：【employee_id（员工employeeId），employee_no（员工工号）】，示例值："employee_id"
	UserApproval *CreateAttendanceUserApprovalReqUserApproval `json:"user_approval,omitempty"` // 审批信息
}

type CreateAttendanceUserApprovalReqUserApproval struct {
	UserID        string                                                     `json:"user_id,omitempty"`        // 审批用户
	Date          string                                                     `json:"date,omitempty"`           // 审批作用时间
	Outs          []*CreateAttendanceUserApprovalReqUserApprovalOut          `json:"outs,omitempty"`           // 外出信息
	Leaves        []*CreateAttendanceUserApprovalReqUserApprovalLeave        `json:"leaves,omitempty"`         // 请假信息
	OvertimeWorks []*CreateAttendanceUserApprovalReqUserApprovalOvertimeWork `json:"overtime_works,omitempty"` // 加班信息
	Trips         []*CreateAttendanceUserApprovalReqUserApprovalTrip         `json:"trips,omitempty"`          // 出差信息
}

type CreateAttendanceUserApprovalReqUserApprovalOut struct {
	UniqID        string     `json:"uniq_id,omitempty"`        // 外出类型唯一 ID，代表一种外出类型，长度小于 14
	Unit          int64      `json:"unit,omitempty"`           // 外出时长单位，可用值：【1（天），2（小时），3（半天），4（半小时）】
	Interval      int64      `json:"interval,omitempty"`       // 假期时长（单位秒）
	StartTime     string     `json:"start_time,omitempty"`     // 开始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime       string     `json:"end_time,omitempty"`       // 结束时间，时间格式为 yyyy-MM-dd HH:mm:ss
	I18nNames     *I18nNames `json:"i18n_names,omitempty"`     // 外出多语言展示，格式为 map，key 为["ch"、"en"、"ja"]，其中 ch 代表中文，en 代表英文、ja 代表日文
	DefaultLocale string     `json:"default_locale,omitempty"` // 默认语言类型，由于飞书客户端支持中、英、日三种语言，当用户切换语言时，如果外出名称没有所对应语言的名称，则会使用默认语言的名称
	Reason        string     `json:"reason,omitempty"`         // 外出理由
}

type CreateAttendanceUserApprovalReqUserApprovalLeave struct {
	UniqID        string     `json:"uniq_id,omitempty"`        // 假期类型唯一 ID，代表一种假期类型，长度小于 14
	Unit          int64      `json:"unit,omitempty"`           // 假期时长单位，可用值：【1（天），2（小时），3（半天），4（半小时）】
	Interval      int64      `json:"interval,omitempty"`       // 假期时长（单位秒）
	StartTime     string     `json:"start_time,omitempty"`     // 开始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime       string     `json:"end_time,omitempty"`       // 结束时间，时间格式为 yyyy-MM-dd HH:mm:ss
	I18nNames     *I18nNames `json:"i18n_names,omitempty"`     // 假期多语言展示，格式为 map，key 为["ch"、"en"、"ja"]，其中 ch 代表中文，en 代表英文、ja 代表日文
	DefaultLocale string     `json:"default_locale,omitempty"` // 默认语言类型，由于飞书客户端支持中、英、日三种语言，当用户切换语言时，如果假期名称没有所对应语言的名称，则会使用默认语言的名称，可用值：【ch（中文），en（英文），ja（日文）】
	Reason        string     `json:"reason,omitempty"`         // 请假理由，必选字段
}

type CreateAttendanceUserApprovalReqUserApprovalOvertimeWork struct {
	Duration  float64 `json:"duration,omitempty"`   // 加班时长
	Unit      int64   `json:"unit,omitempty"`       // 加班时长单位，可用值：【1（天），2（小时）】
	Category  int64   `json:"category,omitempty"`   // 加班日期类型，可用值：【1（工作日），2（休息日），3（节假日）】
	Type      int64   `json:"type,omitempty"`       // 加班规则类型，可用值：【0（不关联加班规则），1（调休），2（加班费）】
	StartTime string  `json:"start_time,omitempty"` // 开始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime   string  `json:"end_time,omitempty"`   // 结束时间，时间格式为 yyyy-MM-dd HH:mm:ss
}

type CreateAttendanceUserApprovalReqUserApprovalTrip struct {
	StartTime        string `json:"start_time,omitempty"`         // 开始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime          string `json:"end_time,omitempty"`           // 结束时间，时间格式为 yyyy-MM-dd HH:mm:ss
	Reason           string `json:"reason,omitempty"`             // 出差理由
	ApprovePassTime  string `json:"approve_pass_time,omitempty"`  // 审批通过时间，时间格式为 yyyy-MM-dd HH:mm:ss
	ApproveApplyTime string `json:"approve_apply_time,omitempty"` // 审批申请时间，时间格式为 yyyy-MM-dd HH:mm:ss
}

type createAttendanceUserApprovalResp struct {
	Code int64                             `json:"code,omitempty"` // 错误码，非 0 表示失败
	Msg  string                            `json:"msg,omitempty"`  // 错误描述
	Data *CreateAttendanceUserApprovalResp `json:"data,omitempty"` // -
}

type CreateAttendanceUserApprovalResp struct {
	UserApprovals []*CreateAttendanceUserApprovalRespUserApproval `json:"user_approvals,omitempty"` // 审批结果列表
}

type CreateAttendanceUserApprovalRespUserApproval struct {
	UserID        string                                                      `json:"user_id,omitempty"`        // 审批用户 ID
	Date          string                                                      `json:"date,omitempty"`           // 审批作用时间
	Outs          []*CreateAttendanceUserApprovalRespUserApprovalOut          `json:"outs,omitempty"`           // 外出信息
	Leaves        []*CreateAttendanceUserApprovalRespUserApprovalLeave        `json:"leaves,omitempty"`         // 请假信息
	OvertimeWorks []*CreateAttendanceUserApprovalRespUserApprovalOvertimeWork `json:"overtime_works,omitempty"` // 加班信息
	Trips         []*CreateAttendanceUserApprovalRespUserApprovalTrip         `json:"trips,omitempty"`          // 出差信息
}

type CreateAttendanceUserApprovalRespUserApprovalOut struct {
	UniqID           string     `json:"uniq_id,omitempty"`            // 外出类型唯一 ID，代表一种外出类型，长度小于 14
	Unit             int64      `json:"unit,omitempty"`               // 外出时长单位，可用值：【1（天），2（小时），3（半天），4（半小时）】
	Interval         int64      `json:"interval,omitempty"`           // 假期时长（单位秒）
	StartTime        string     `json:"start_time,omitempty"`         // 开始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime          string     `json:"end_time,omitempty"`           // 结束时间，时间格式为 yyyy-MM-dd HH:mm:ss
	I18nNames        *I18nNames `json:"i18n_names,omitempty"`         // 外出多语言展示，格式为 map，key 为["ch"、"en"、"ja"]，其中 ch 代表中文，en 代表英文、ja 代表日文
	DefaultLocale    string     `json:"default_locale,omitempty"`     // 默认语言类型，由于飞书客户端支持中、英、日三种语言，当用户切换语言时，如果外出名称没有所对应语言的名称，则会使用默认语言的名称
	Reason           string     `json:"reason,omitempty"`             // 外出理由
	ApprovePassTime  string     `json:"approve_pass_time,omitempty"`  // 审批通过时间
	ApproveApplyTime string     `json:"approve_apply_time,omitempty"` // 审批申请时间
}

type CreateAttendanceUserApprovalRespUserApprovalLeave struct {
	UniqID           string     `json:"uniq_id,omitempty"`            // 假期类型唯一 ID，代表一种假期类型，长度小于 14
	Unit             int64      `json:"unit,omitempty"`               // 假期时长单位，可用值：【1（天），2（小时），3（半天），4（半小时）】
	Interval         int64      `json:"interval,omitempty"`           // 假期时长（单位秒）
	StartTime        string     `json:"start_time,omitempty"`         // 开始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime          string     `json:"end_time,omitempty"`           // 结束时间，时间格式为 yyyy-MM-dd HH:mm:ss
	I18nNames        *I18nNames `json:"i18n_names,omitempty"`         // 假期多语言展示，格式为 map，key 为["ch"、"en"、"ja"]，其中 ch 代表中文，en 代表英文、ja 代表日文
	DefaultLocale    string     `json:"default_locale,omitempty"`     // 默认语言类型，由于飞书客户端支持中、英、日三种语言，当用户切换语言时，如果假期名称没有所对应语言的名称，则会使用默认语言的名称，可用值：【ch（中文），en（英文），ja（日文）】
	Reason           string     `json:"reason,omitempty"`             // 请假理由
	ApprovePassTime  string     `json:"approve_pass_time,omitempty"`  // 审批通过时间，时间格式为 yyyy-MM-dd HH:mm:ss
	ApproveApplyTime string     `json:"approve_apply_time,omitempty"` // 审批申请时间，时间格式为 yyyy-MM-dd HH:mm:ss
}

type CreateAttendanceUserApprovalRespUserApprovalOvertimeWork struct {
	Duration  float64 `json:"duration,omitempty"`   // 加班时长
	Unit      int64   `json:"unit,omitempty"`       // 加班时长单位，可用值：【1（天），2（小时）】
	Category  int64   `json:"category,omitempty"`   // 加班日期类型，可用值：【1（工作日），2（休息日），3（节假日）】
	Type      int64   `json:"type,omitempty"`       // 加班规则类型，可用值：【0（不关联加班规则），1（调休），2（加班费），3（关联加班规则，没有调休或加班费）】
	StartTime string  `json:"start_time,omitempty"` // 开始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime   string  `json:"end_time,omitempty"`   // 结束时间，时间格式为 yyyy-MM-dd HH:mm:ss
}

type CreateAttendanceUserApprovalRespUserApprovalTrip struct {
	StartTime        string `json:"start_time,omitempty"`         // 开始时间，时间格式为 yyyy-MM-dd HH:mm:ss
	EndTime          string `json:"end_time,omitempty"`           // 结束时间，时间格式为 yyyy-MM-dd HH:mm:ss
	Reason           string `json:"reason,omitempty"`             // 出差理由
	ApprovePassTime  string `json:"approve_pass_time,omitempty"`  // 审批通过时间，时间格式为 yyyy-MM-dd HH:mm:ss
	ApproveApplyTime string `json:"approve_apply_time,omitempty"` // 审批申请时间，时间格式为 yyyy-MM-dd HH:mm:ss
}
