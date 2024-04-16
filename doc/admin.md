# admin 模块管理

## 生成吗命令
```bash
# 进入项目目录
mkdir -p app/admin/api
cd app/admin/api
# 创建api.api文件
# 编辑api.api 语法参考@see  https://go-zero.dev/docs/tutorials

# 执行.api文件
# -dir 指定生成目录  
# --style 生成代码风格 goZero 小驼峰写法
# --home 生成代码采用的模板 
goctl api go --api api.api -dir ./  --style="goZero" --home="../../../template"

# 启动项目
go run api.go
```