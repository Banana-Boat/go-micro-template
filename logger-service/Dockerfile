# 使用 multi-stage 进行构建，进一步减小镜像大小

# Build stage
FROM golang:1.19-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
# alpine没有curl，需要先下载
RUN apk add curl
RUN  curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY --from=builder  /app/migrate .
# <src> 是一个目录，则将目录下的所有文件写入<dest>中
COPY internal/db/migration ./migration

EXPOSE 8080

# 同时使用CMD和ENTRYPOINT，会将cmd作为entrypoint的默认参数进行执行，
# 即 ENTRYPOINT [ "/app/start.sh", "/app/main" ]
# 容器具有指定可执行文件，同时需要能方便地修改默认参数，可选择此中方式
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]