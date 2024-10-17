# gen-project
自动生成工程目录结构工具

## 使用方式
1. 安装工具
```shell
go install github.com/lie-flat-planet/gen-project@latest
```
2. 生成目录, 例如
```shell
gen-project github.com/your-group/your-project-name
```
更多命令细节请执行 `gen-project -h`

3. 生成工程后进入工程目录, 按如下步骤即可启动工程

第一步: 下载依赖到本地：
```shell
go mod tidy
```

第二步: 修改数据库，redis的配置信息:
```shell
修改你本地 config.Config 中的配置信息
```

第三步: 启动服务:
```shell
make local-run
```



