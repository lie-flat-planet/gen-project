package file

import (
	"fmt"
	"strings"

	"github.com/lie-flat-planet/gen-project/global"
)

type Config struct{}

func (*Config) Name() string {
	return "config.go"
}

func (*Config) Content() string {
	return strings.TrimSpace(fmt.Sprintf(`
package config

import (
	tool "github.com/lie-flat-planet/service-init-tool"
	"github.com/lie-flat-planet/service-init-tool/component/database"
	"github.com/lie-flat-planet/service-init-tool/component/redis"
	"github.com/lie-flat-planet/service-init-tool/component/prometheus"
)

func init() {
	if err := tool.Init("./", &Config); err != nil {
		panic(err)
	}
}

var Config = struct {
	Server *tool.Server
	Mysql  *database.Mysql
	Redis  *redis.Redis
	Prom   *prometheus.Prom
}{
	Server: &tool.Server{
		Name:     "%s",
		Code:     333 * 1e3,
		HttpPort: 8081,
		RunMode:  "debug",
	},
	Mysql: &database.Mysql{
		MysqlConf: database.MysqlConf{
			Host:        "127.0.0.1:3306",
			User:        "root",
			Password:    "",
			DbName:      "",
			MaxIdleConn: 1,
			MaxOpenConn: 2,
		},
	},
	Redis: &redis.Redis{
		Config: redis.Config{
			Host:     "127.0.0.1:6379",
			Password: "",
			DB:       0,
			PoolSize: 5,
			Timeout:  0,
		},
	},
	Prom: &prometheus.Prom{
		Addr: "1.2.3.4",
	},
}

`, global.GetProjectName()))
}
