package file

import (
	"fmt"
	"github.com/lie-flat-planet/gen-project/global"
	"strings"
)

type ModelUser struct{}

func (*ModelUser) Name() string {
	return "user.go"
}

func (*ModelUser) Content() string {
	return strings.TrimSpace(fmt.Sprintf(`
package model

import (
	"%s/config"
)

func init() {
	config.Config.Mysql.AppendModel(&User{})
}

type User struct {
	ID
	Name  string
	Age   int
	Phone string
	TimeAt
}

func (User) TableName() string {
	return "user"
}
`, global.GetModuleName()))
}
