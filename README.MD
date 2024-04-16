# 项目架构

## 项目简介
- 基于go-zero的关联后台项目。给一些新手参考。

## 框架选择
- go-zero go微服务框架

## 目录说明
- project 项目目录
  - app 应用目录
    - admin 后台管理模块
      - api 接口对接层,相关接口都在这里
      - rpc 由于管理后台不需要rpc 暂时为空
    - article 文章模块
      - api 接口对接层,相关接口都在这里
      - rpc 其他模块调用时 使用的接口
    - user 用户模块
      - api 接口对接层,相关接口都在这里
      - rpc 其他模块调用时 使用的接口
  - common 公用组件库
    - helper 助手函数包，自定义的方法
    - response 公用返回包
  - database 数据库管理库
    - main 数据库迁移 
    - model 数据库实体
  - build 部署相关库
  - doc 项目文档
    - [管理后台模块](./doc/admin.md)
    - [文章模块](./doc/article.md)
  - template goctl 模板库 
    - 默认模板 执行命令 ```goctl --version``` 找到版本号,我的是1.6.3 查看默认模板 ```cd ~/.goctl/1.6.3```
    - 复制模板到项目来，使所有开发人员使用一套模板 ```cp -R ~/.goctl/1.6.3/* ./template/```
  - go.mod 包管理