package file

import (
	"strings"
)

type ModelDb struct{}

func (*ModelDb) Name() string {
	return "db.go"
}

func (*ModelDb) Content() string {
	return strings.TrimSpace(`
package model

import (
	"github.com/lie-flat-planet/service-init-tool/component/mysql"
)

type ID struct {
	ID uint ` + "`gorm:\"primarykey\"`" + `
}

type TimeAt struct {
	CreatedAt mysql.Time ` + "`gorm:\"default:CURRENT_TIMESTAMP\"`" + `
	UpdatedAt mysql.Time ` + "`gorm:\"default:CURRENT_TIMESTAMP\"`" + `
	DeletedAt mysql.DeletedTime
}
`)
}
