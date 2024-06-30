package dir

func init() {
	Dirs = append(Dirs, NewStatusError())
}

type StatusError struct {
	projectName string
}

func NewStatusError() *StatusError {
	return new(StatusError)
}

func (p *StatusError) RootPath() string {
	return "pkg/constant/"
}

func (p *StatusError) Path() string {
	return "errors/"
}

func (p *StatusError) File() string {
	return "status_error.go"
}

func (p *StatusError) FileContent() string {
	return `package errors

import (
	"net/http"
)

//go:generate tools gen error StatusError
type StatusError int

func (StatusError) ServiceCode() int {
	return 999 * 1e3
}

const (
	// @errTalk 请求参数错误
	BadRequest StatusError = http.StatusBadRequest*1e6 + iota
)

const (
	// @errTalk 未找到
	NotFound StatusError = http.StatusNotFound*1e6 + iota
)

const (
	// @errTalk 未授权
	Unauthorized StatusError = http.StatusUnauthorized*1e6 + iota
)

const (
	// @errTalk 操作冲突
	Conflict StatusError = http.StatusConflict*1e6 + iota
)

const (
	// @errTalk 权限不够
	Forbidden StatusError = http.StatusForbidden*1e6 + iota
)
const (
	// @errTalk 内部处理错误
	InternalError StatusError = http.StatusInternalServerError*1e6 + iota
)
`
}

func (p *StatusError) SetProjectName(name string) {
	p.projectName = name
}
