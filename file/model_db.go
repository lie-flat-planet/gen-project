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
	"github.com/lie-flat-planet/service-init-tool/component/database"
	// "gorm.io/plugin/soft_delete"
)

type ID struct {
	ID uint ` + "`gorm:\"primarykey\"`" + `
}

type TimeAt struct {
	CreatedAt database.Time  ` + "`json:\"createdAt\" gorm:\"default:CURRENT_TIMESTAMP\"`" + `
	UpdatedAt database.Time  ` + "`json:\"updatedAt\" gorm:\"default:CURRENT_TIMESTAMP\"`" + `
	// DeletedAt soft_delete.DeletedAt ` + "`json:\"deletedAt\"`" + `
}

type TimestampAt struct {
	CreatedAt   int64 ` + "`json:\"createdAt\" gorm:\"autoCreateTime:true;not null\"`" + `
	UpdatedAt   int64 ` + "`json:\"updatedAt\" gorm:\"autoCreateTime:true;not null\"`" + `
	CreatedTime   string ` + "`json:\"createdTime\" gorm:\"-\"` // 时间字符串" + `
	UpdatedTime   string ` + "`json:\"updatedTime\" gorm:\"-\"` // 时间字符串" + `
}
`)
}
