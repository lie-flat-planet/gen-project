package file

import "strings"

type CursorRules struct{}

func (*CursorRules) Name() string {
	return ".cursorrules"
}

func (*CursorRules) Content() string {
	return strings.TrimSpace(`
# Cursor Rules Configuration

# 前置
- 你是一个经验丰富的Go语言开发工程师，熟悉N9E的代码结构和业务逻辑
- 你现在已经在一个Go语言项目中，因此不需要帮我生成项目目录结构
- 我需要一个Go语言的代码补全工具，因此需要你帮我生成代码补全规则
- Go语言版本为1.24.0
- 精通各种设计模式，特别是领域驱动(DDD)，并能灵活运用

# 指定项目的主要语言
language: go

# 定义项目结构
structure:
  # 源代码目录
  src:
    - api/
    - cmd/
    - config/
    - internal/
    - pkg/
  
  # 测试目录
  test:
    - "**/*_test.go"
  
  # 配置文件目录
  config:
    - config/
  
  # 文档目录
  docs:
    - docs/
    - README.md

# 定义代码分析规则
analysis:
  # 启用的分析器
  enabled:
    - go
    - markdown
    - yaml
    - json
  
  # 分析深度
  depth: 3
  
  # 忽略特定的警告或错误
  ignore:
    - "unused parameter"
    - "unused variable"

# 定义智能提示规则
intellisense:
  # 自动导入包
  auto_import: true
  
  # 代码补全
  completion:
    enabled: true
    
  # 代码导航
  navigation:
    enabled: true
    
  # 代码格式化
  formatting:
    enabled: true
    
# 定义搜索规则
search:
  # 最大搜索结果数
  max_results: 100
  
  # 搜索排除模式（与 .cursorignore 配合使用）
  exclude:
    - "vendor/**"
    - "bin/**"
    - "front/statik/**" 
`)
}
