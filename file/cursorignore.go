package file

import (
	"strings"
)

type CursorIgnore struct{}

func (*CursorIgnore) Name() string {
	return ".cursorignore"
}

func (*CursorIgnore) Content() string {
	return strings.TrimSpace(`
# 编译后的二进制文件和目录
bin/
*.exe
*.exe~
*.dll
*.dylib
*.so
*.o
*.a

# 大型静态资源文件
front/statik/statik.go

# 构建输出目录
build/
dist/
out/
pub/
tmp/
run/
tarball/

# 数据和日志目录
data*/
log*/
*.log

# 临时文件
*.sw[po]
*.tar.gz
.DS_Store
.cache-loader
.payload

# 版本控制相关
.git/
.github/

# IDE 相关
.idea/
.vscode/
.index/

# 第三方依赖
vendor/

# 配置文件
etc/*.local.yml
etc/*.local.conf
etc/rsa/*
etc/plugins/*.local.yml
etc.local*/

# Docker 相关
docker/pub/
docker/compose-bridge/mysqldata/
docker/compose-host-network/mysqldata/
docker/compose-host-network-metric-log/mysqldata/
docker/compose-host-network-metric-log/n9e-logs/
docker/compose-postgres/pgdata/
`)
}
