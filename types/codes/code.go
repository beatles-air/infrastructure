// Package code to define some code
package codes

import (
	"reflect"
)

const (
	// CodeOK 0
	CodeOK = iota
	// CodeParamInvalid ...
	CodeParamInvalid
	// CodeSystemErr ...
	CodeSystemErr
	// CodeNoPermission ...
	CodeNoPermission
	// CodeServerTimeout ...
	CodeServerTimeout
	// CodeResourceNotFound ...
	CodeResourceNotFound
	// CodeIllegeOP ...
	CodeIllegeOP
	// ErrNoSuchCode ...
	ErrNoSuchCode = "错误码未定义"
)

var messages = map[int]string{
	CodeOK:               "成功",
	CodeParamInvalid:     "参数非法",
	CodeSystemErr:        "系统错误",
	CodeNoPermission:     "没有权限",
	CodeServerTimeout:    "服务超时",
	CodeResourceNotFound: "资源未找到",
	CodeIllegeOP:         "非法操作",
}

// Proto define a Proto type
type Proto struct {
	ErrCode    int    `json:"errcode"`
	ErrMessage string `json:"errmsg"`
}

// New create a new *Proto
func New(code int, message string) *Proto {
	if message == "" {
		message = GetMessage(code)
	}

	return &Proto{
		ErrCode:    code,
		ErrMessage: message,
	}
}

// Get get Proto with specified code
func Get(code int) *Proto {
	return &Proto{
		ErrCode:    code,
		ErrMessage: GetMessage(code),
	}
}

// GetMessage get code desc from messages
func GetMessage(code int) string {
	v, ok := messages[code]
	if !ok {
		return ErrNoSuchCode
	}
	return v
}

// Fill ... fill a response struct will *Proto
// TODO: validate v
func Fill(v interface{}, ci *Proto) interface{} {
	if reflect.TypeOf(v).Kind() != reflect.Ptr {
		panic("v must be ptr")
	}
	field := reflect.ValueOf(v).Elem().
		FieldByName("Proto")

	// set field
	field.FieldByName("ErrCode").SetInt(int64(ci.ErrCode))
	field.FieldByName("ErrMessage").SetString(ci.ErrMessage)

	return v
}
