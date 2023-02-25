# gin-template

## 相关 CLI 工具

- [**Docker**](https://hub.docker.com/)
- [**golang-migrate**](https://github.com/golang-migrate/migrate)
- [**sqlc**](https://docs.sqlc.dev/en/stable/index.html)

## Go 相关依赖

- [**Gin**](https://github.com/gin-gonic/gin)
- [**Testify**](https://github.com/stretchr/testify)
- [**Viper**](https://github.com/spf13/viper)
- [**Paseto**](https://github.com/o1egl/paseto)

## 开发场景

#### 基本环境

- 安装上述工具
- 安装 Go 依赖 `go mod tidy`
- 修改 go.mod 中 module 名（项目文件 import 需要全局替换）

#### 数据库

1. 修改 Makefile 中数据库 URL
2. 执行`make mysql`，Docker 启动 mysql:8.0 容器
3. 执行`migrate_init`生成 schema
4. 使用 [**dbdiagram**](https://dbdiagram.io/home) 工具设计数据库，将 sql 语句复制到上一步的 schema 中
5. 执行`make migrate_up`创建数据表

#### sqlc

- 在 internal/db/query/ 下创建 表名.sql 文件，根据官网编写 sql 语句
- 执行`make sqlc`生成.go 文件

#### 编译运行

- 修改 app.env 文件，并设置 gitignore
- 执行`make server`，编译运行

## 部署场景

- 修改 Makefile 中 docker_build 命令的镜像名，并执行创建
- 修改 compose.yaml 中的镜像名即数据库 URL
- 在服务器中启动 Docker，并拉取 mysql:8.0
- 将生成的镜像与 compose.yaml 文件上传至服务器，执行`docker compose up`
